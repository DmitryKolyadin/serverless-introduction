from pydantic import BaseModel


class FavoriteCreate(BaseModel):
    city: str


class Favorite(BaseModel):
    city: str