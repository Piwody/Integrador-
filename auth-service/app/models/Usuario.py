from typing import Optional  
from sqlmodel import SQLModel, Field, create_engine

sqlite_file = "users.db"
engine = create_engine(f"sqlite:///{sqlite_file}", echo=True)

class User(SQLModel, table=True):
    id: Optional[int] = Field(default=None, primary_key=True) 
    username: str = Field(unique=True, index=True)
    hashed_password: str


class UserIn(SQLModel):
    username: str
    password: str

def create_db():
    SQLModel.metadata.create_all(engine)