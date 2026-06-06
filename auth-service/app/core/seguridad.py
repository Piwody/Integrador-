import os
import bcrypt
from dotenv import load_dotenv
import jwt
from datetime import datetime, timedelta, timezone

load_dotenv()
PEPPER = os.getenv("PEPPER", "pepper_de_respaldo_por_si_acaso")
JWT_SECRET = os.getenv("JWT_SECRET", "clave-secreta-desarrollo-cambiar-en-produccion")
JWT_ALGORITHM = "HS256"

def hash_password(password: str) -> str:
    password_peppered = (password + PEPPER).encode()
    hashed = bcrypt.hashpw(password_peppered, bcrypt.gensalt())
    return hashed.decode()

def verify_password(password: str, hashed_password: str) -> bool:
    password_peppered = (password + PEPPER).encode()
    return bcrypt.checkpw(password_peppered, hashed_password.encode())

def create_access_token(username: str) -> str:
    expire = datetime.now(timezone.utc) + timedelta(minutes=30)
    payload = {
        "sub": username,
        "exp": expire
    }
    return jwt.encode(payload, JWT_SECRET, algorithm=JWT_ALGORITHM)

def verify_token(token: str):
    try:
        payload = jwt.decode(token, JWT_SECRET, algorithms=[JWT_ALGORITHM])
        return payload.get("sub")
    except jwt.ExpiredSignatureError:
        return None
    except jwt.InvalidTokenError:
        return None

        