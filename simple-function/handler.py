import requests


def get_coordinates(city: str):
    url = "https://geocoding-api.open-meteo.com/v1/search"
    params = {
        "name": city,
        "count": 1,
        "language": "en",
        "format": "json"
    }

    resp = requests.get(url, params=params, timeout=5)
    data = resp.json()

    if "results" not in data or not data["results"]:
        return None

    result = data["results"][0]
    return result["latitude"], result["longitude"], result["name"]


def get_weather(lat: float, lon: float):
    url = "https://api.open-meteo.com/v1/forecast"
    params = {
        "latitude": lat,
        "longitude": lon,
        "current_weather": True
    }

    resp = requests.get(url, params=params, timeout=5)
    data = resp.json()

    return data.get("current_weather")


def handler(event, context):
    query = event.get("queryStringParameters") or {}
    city = query.get("city")

    if not city:
        return {
            "statusCode": 400,
            "body": {
                "error": "city query param is required"
            }
        }

    coords = get_coordinates(city)
    if not coords:
        return {
            "statusCode": 404,
            "body": {
                "error": "city not found"
            }
        }

    lat, lon, normalized_name = coords
    weather = get_weather(lat, lon)

    if not weather:
        return {
            "statusCode": 500,
            "body": {
                "error": "failed to fetch weather"
            }
        }

    return {
        "statusCode": 200,
        "body": {
            "city": normalized_name,
            "temperature": weather["temperature"],
            "windspeed": weather["windspeed"]
        }
    }