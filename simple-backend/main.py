from fastapi import FastAPI
from routes import router

app = FastAPI(title="Favorites API", docs_url="/api/docs", redoc_url="/api/redocs", openapi_url="/api/openapi.json")

app.include_router(router)