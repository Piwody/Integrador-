from fastapi import FastAPI, HTTPException
from sqlmodel import Session, select
from app.models import engine, User, UserIn, create_db
from app.core.seguridad import hash_password, verify_password, create_access_token, JWT_SECRET

app = FastAPI(title="Autenticacion de Proyectointergrador")

@app.on_event("startup")
def on_startup():
    create_db()

@app.post("/register")
def register(user: UserIn):
    with Session(engine) as session:
        existing_user = session.exec(
            select(User).where(User.username == user.username)
        ).first()

        if existing_user:
            raise HTTPException(status_code=409, detail="El usuario ya existe")

        new_user = User(
            username=user.username,
            hashed_password=hash_password(user.password)
        )

        session.add(new_user)
        session.commit()
        session.refresh(new_user)

        return {
            "message": "Usuario registrado correctamente",
            "id": new_user.id,
            "username": new_user.username
        }

@app.post("/login")
def login(user: UserIn):
    with Session(engine) as session:
        db_user = session.exec(
            select(User).where(User.username == user.username)
        ).first()

        if not db_user or not verify_password(user.password, db_user.hashed_password):
            raise HTTPException(status_code=401, detail="Credenciales incorrectas")

        token = create_access_token(db_user.username)
        return {
            "message": "Login exitoso",
            "access_token": token,
            "token_type": "bearer"
        }