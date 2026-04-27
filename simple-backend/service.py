from storage import favorites


def add_favorite(city: str):
    favorites.append({"city": city})
    return {"city": city}


def get_favorites():
    return favorites