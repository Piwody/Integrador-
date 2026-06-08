from fastapi import FastAPI, HTTPException, Header
from fastapi.middleware.cors import CORSMiddleware
from sqlmodel import Session, select
from pydantic import BaseModel
from app.models import engine, User, UserIn, create_db
from app.core.seguridad import (
    hash_password, verify_password,
    create_access_token, verify_token,
    encrypt_aes_cbc, decrypt_aes_cbc
)

app = FastAPI(title="Autenticacion de Proyecto Integrador")

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)



class AESRequest(BaseModel):
    text: str

class AESDecryptRequest(BaseModel):
    encrypted: str



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

@app.get("/api/auth/verify")
def verify(authorization: str = Header(...)):
    token = authorization.replace("Bearer ", "").replace("bearer ", "")
    username = verify_token(token)
    if not username:
        raise HTTPException(status_code=401, detail="Token inválido o expirado")
    return {"valid": True, "username": username}



@app.post("/api/aes/encrypt")
def aes_encrypt(req: AESRequest, authorization: str = Header(...)):
    """Encripta un texto usando AES-CBC. Requiere token válido."""
    token = authorization.replace("Bearer ", "").replace("bearer ", "")
    if not verify_token(token):
        raise HTTPException(status_code=401, detail="Token inválido o expirado")

    encrypted = encrypt_aes_cbc(req.text)
    return {
        "original": req.text,
        "encrypted": encrypted,
        "algorithm": "AES-256-CBC",
        "note": "Formato: IV (16 bytes) || Ciphertext en hexadecimal"
    }

@app.post("/api/aes/decrypt")
def aes_decrypt(req: AESDecryptRequest, authorization: str = Header(...)):
    """Desencripta un texto cifrado con AES-CBC. Requiere token válido."""
    token = authorization.replace("Bearer ", "").replace("bearer ", "")
    if not verify_token(token):
        raise HTTPException(status_code=401, detail="Token inválido o expirado")

    try:
        decrypted = decrypt_aes_cbc(req.encrypted)
        return {
            "encrypted": req.encrypted,
            "decrypted": decrypted,
            "algorithm": "AES-256-CBC"
        }
    except Exception:
        raise HTTPException(status_code=400, detail="Error al desencriptar: datos inválidos")