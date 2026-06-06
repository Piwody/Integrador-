import os
import bcrypt
from dotenv import load_dotenv
import jwt  
from datetime import datetime, timedelta, timezone

load_dotenv()
PEPPER = os.getenv("PEPPER", "pepper_de_respaldo_por_si_acaso")
JWT_SECRET = os.getenv("JWT_SECRET", "clave-secreta-desarrollo-cambiar-en-produccion")

def hash_password(password: str) -> str:
    password_peppered = (password + PEPPER).encode()
    
    hashed = bcrypt.hashpw(password_peppered, bcrypt.gensalt())
    
    return hashed.decode()

def verify_password(password: str, hashed_password: str) -> bool:
    
    password_peppered = (password + PEPPER).encode()
    return bcrypt.checkpw(password_peppered, hashed_password.encode())

def create_access_token(username: str):
    #30min de accesdp
    expire = datetime.now(timezone.utc) + timedelta(minutes=30)
    
    payload = {
        "sub": username,
        "exp": expire
    }
    
    #has
    encoded_jwt = jwt.encode(payload, JWT_SECRET, algorithm="HS256")
    return encoded_jwt