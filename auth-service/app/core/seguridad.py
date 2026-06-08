import os
import bcrypt
from dotenv import load_dotenv
import jwt
from datetime import datetime, timedelta, timezone
from Crypto.Cipher import AES
from Crypto.Util.Padding import pad, unpad
import secrets

load_dotenv()
PEPPER = os.getenv("PEPPER", "Mamacita_Alista_El_Higado")
JWT_SECRET = os.getenv("JWT_SECRET", "Aqui_Andamos_Haciendo_solo_yo_xd")
JWT_ALGORITHM = "HS256"


AES_KEY = os.getenv("AES_KEY", "MateoCantos2026ProyectoIntegra").encode()[:32]



def hash_password(password: str) -> str:
    password_peppered = (password + PEPPER).encode()
    hashed = bcrypt.hashpw(password_peppered, bcrypt.gensalt())
    return hashed.decode()

def verify_password(password: str, hashed_password: str) -> bool:
    password_peppered = (password + PEPPER).encode()
    return bcrypt.checkpw(password_peppered, hashed_password.encode())


def create_access_token(username: str) -> str:
    expire = datetime.now(timezone.utc) + timedelta(minutes=30)
    payload = {"sub": username, "exp": expire}
    return jwt.encode(payload, JWT_SECRET, algorithm=JWT_ALGORITHM)

def verify_token(token: str):
    try:
        payload = jwt.decode(token, JWT_SECRET, algorithms=[JWT_ALGORITHM])
        return payload.get("sub")
    except jwt.ExpiredSignatureError:
        return None
    except jwt.InvalidTokenError:
        return None



def encrypt_aes_cbc(plaintext: str) -> str:
    iv = secrets.token_bytes(16)                        
    cipher = AES.new(AES_KEY, AES.MODE_CBC, iv)
    padded = pad(plaintext.encode(), AES.block_size)    
    ciphertext = cipher.encrypt(padded)
    return (iv + ciphertext).hex()                      

def decrypt_aes_cbc(encrypted_hex: str) -> str:
    raw = bytes.fromhex(encrypted_hex)
    iv = raw[:16]                                        
    ciphertext = raw[16:]
    cipher = AES.new(AES_KEY, AES.MODE_CBC, iv)
    plaintext = unpad(cipher.decrypt(ciphertext), AES.block_size)
    return plaintext.decode()