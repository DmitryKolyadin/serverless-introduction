from fastapi import APIRouter
from schemas import FavoriteCreate
from service import add_favorite, get_favorites

router = APIRouter(prefix="/api")


@router.post("/favorites")
def create_favorite(data: FavoriteCreate):
    return add_favorite(data.city)


@router.get("/favorites")
def list_favorites():
    return get_favorites()