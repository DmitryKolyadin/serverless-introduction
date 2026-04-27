package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type FavoriteCreate struct {
	City string `json:"city"`
}

var (
	appRouter *gin.Engine
	appCfg    Config
	initOnce  sync.Once
	initErr   error
)

func initApp() error {
	initOnce.Do(func() {
		cfg, err := LoadConfig()
		if err != nil {
			initErr = err
			return
		}

		store, err := NewStorage(cfg)
		if err != nil {
			initErr = err
			return
		}

		appCfg = cfg
		appRouter = setupRouter(store)
	})

	return initErr
}

func setupRouter(store *Storage) *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	{
		api.POST("/favorites", func(c *gin.Context) {
			var input FavoriteCreate
			if err := c.ShouldBindJSON(&input); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
				return
			}
			if input.City == "" {
				c.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
				return
			}

			item, err := store.AddFavorite(c.Request.Context(), input.City)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save favorite"})
				return
			}

			c.JSON(http.StatusOK, item)
		})

		api.GET("/favorites", func(c *gin.Context) {
			items, err := store.ListFavorites(c.Request.Context())
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list favorites"})
				return
			}
			c.JSON(http.StatusOK, items)
		})
	}

	return r
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if err := initApp(); err != nil {
		log.Printf("init error: %v", err)
		http.Error(w, "initialization failed", http.StatusInternalServerError)
		return
	}

	appRouter.ServeHTTP(w, r)
}

func main() {
	if err := initApp(); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":"+appCfg.Port, appRouter); err != nil {
		log.Fatal(err)
	}
}
