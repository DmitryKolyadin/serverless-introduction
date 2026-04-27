import React, { useEffect, useMemo, useState } from "react";
import { createRoot } from "react-dom/client";
import "nes.css/css/nes.min.css";
import "./styles.css";

type WeatherResponse = {
  city: string;
  temperature: number;
  windspeed: number;
};

type Favorite = {
  city: string;
};

type StatusType = "idle" | "loading" | "success" | "error";

type UiStatus = {
  type: StatusType;
  message: string;
};

const WEATHER_URL = "/weather";
const FAVORITES_URL = "/api/favorites";

function App() {
  const [city, setCity] = useState("Orenburg");
  const [weather, setWeather] = useState<WeatherResponse | null>(null);
  const [favorites, setFavorites] = useState<Favorite[]>([]);
  const [status, setStatus] = useState<UiStatus>({
    type: "idle",
    message: "Type a city and check weather.",
  });
  const [loadingWeather, setLoadingWeather] = useState(false);
  const [loadingFavorites, setLoadingFavorites] = useState(false);

  const canSave = useMemo(() => Boolean(weather?.city), [weather]);

  const fetchWeather = async () => {
    const normalizedCity = city.trim();
    if (!normalizedCity) {
      setStatus({ type: "error", message: "City is required" });
      return;
    }

    setLoadingWeather(true);
    setStatus({ type: "loading", message: "Fetching weather..." });

    try {
      const targetUrl = `${WEATHER_URL}?city=${encodeURIComponent(normalizedCity)}`;
      const response = await fetch(targetUrl);
      const data = (await response.json()) as
        | WeatherResponse
        | { error?: string };

      if (!response.ok) {
        const errorMessage =
          "error" in data && data.error
            ? data.error
            : "Failed to fetch weather";
        throw new Error(errorMessage);
      }

      setWeather(data as WeatherResponse);
      setStatus({ type: "success", message: "Weather loaded" });
    } catch (error) {
      setWeather(null);
      setStatus({
        type: "error",
        message:
          error instanceof Error ? error.message : "Unknown network error",
      });
    } finally {
      setLoadingWeather(false);
    }
  };

  const loadFavorites = async () => {
    setLoadingFavorites(true);
    setStatus({ type: "loading", message: "Loading favorites..." });

    try {
      const response = await fetch(FAVORITES_URL);
      const data = (await response.json()) as Favorite[] | { detail?: string };

      if (!response.ok) {
        const detail =
          "detail" in data && data.detail
            ? data.detail
            : "Failed to load favorites";
        throw new Error(detail);
      }

      setFavorites(Array.isArray(data) ? data : []);
      setStatus({ type: "success", message: "Favorites loaded" });
    } catch (error) {
      setStatus({
        type: "error",
        message:
          error instanceof Error ? error.message : "Unknown network error",
      });
    } finally {
      setLoadingFavorites(false);
    }
  };

  const addFavorite = async () => {
    if (!weather?.city) {
      setStatus({ type: "error", message: "Load weather before saving" });
      return;
    }

    setStatus({ type: "loading", message: "Saving favorite..." });

    try {
      const response = await fetch(FAVORITES_URL, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ city: weather.city }),
      });

      const data = (await response.json()) as Favorite | { detail?: string };

      if (!response.ok) {
        const detail =
          "detail" in data && data.detail
            ? data.detail
            : "Failed to save favorite";
        throw new Error(detail);
      }

      setStatus({ type: "success", message: "City saved to favorites" });
      await loadFavorites();
    } catch (error) {
      setStatus({
        type: "error",
        message:
          error instanceof Error ? error.message : "Unknown network error",
      });
    }
  };

  useEffect(() => {
    void loadFavorites();
  }, []);

  const statusClass = `status ${status.type === "error" ? "error" : status.type === "success" ? "success" : ""}`;

  return (
    <main className="pixel-app">
      <section className="header">
        <h1>Pixel Weather Radar</h1>
        <p>Fast weather lookup with a tiny favorites board</p>
      </section>

      <section className="grid">
        <article className="panel nes-container with-title is-rounded">
          <p className="title">Weather</p>

          <div className="field">
            <label htmlFor="city">City</label>
            <input
              id="city"
              className="nes-input"
              value={city}
              onChange={(event) => setCity(event.target.value)}
              placeholder="Orenburg"
              onKeyDown={(event) => {
                if (event.key === "Enter") {
                  void fetchWeather();
                }
              }}
            />
          </div>

          <div className="actions">
            <button
              type="button"
              className="nes-btn is-primary"
              onClick={() => void fetchWeather()}
              disabled={loadingWeather}
            >
              Get weather
            </button>
            <button
              type="button"
              className="nes-btn is-success"
              onClick={() => void addFavorite()}
              disabled={!canSave}
            >
              Add favorite
            </button>
          </div>

          <div className={statusClass}>{status.message}</div>

          {weather ? (
            <section className="weather-box">
              <div className="weather-top">
                <div className="weather-main">{weather.city}</div>
                <i className="nes-icon star is-medium" aria-hidden="true" />
              </div>
              <div className="metric-row">
                <span className="metric-label">Temp</span>
                <span className="metric-value">{weather.temperature} C</span>
              </div>
              <div className="metric-row">
                <span className="metric-label">Wind</span>
                <span className="metric-value">{weather.windspeed} km/h</span>
              </div>
            </section>
          ) : (
            <p className="hint">No weather yet. Press Get weather.</p>
          )}
        </article>

        <section className="panel nes-container with-title is-rounded">
          <p className="title">Favorites</p>
          <div className="actions">
            <button
              type="button"
              className="nes-btn"
              onClick={() => void loadFavorites()}
              disabled={loadingFavorites}
            >
              Refresh list
            </button>
          </div>
          {favorites.length === 0 ? (
            <p className="hint">
              No favorites yet. Save a city from weather card.
            </p>
          ) : (
            <ul className="favorites-list">
              {favorites.map((item, index) => (
                <li key={`${item.city}-${index}`} className="favorite-item">
                  <i className="nes-icon is-small heart" aria-hidden="true" />
                  <span>{item.city}</span>
                </li>
              ))}
            </ul>
          )}
        </section>
      </section>
    </main>
  );
}

const rootNode = document.getElementById("root");
if (!rootNode) {
  throw new Error("Root element not found");
}

createRoot(rootNode).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
);
