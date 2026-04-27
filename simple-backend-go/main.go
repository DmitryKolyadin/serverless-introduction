package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Favorite struct {
	City string `json:"city"`
}

type FavoriteCreate struct {
	City string `json:"city"`
}

type FavoriteStore struct {
	mu    sync.RWMutex
	items []Favorite
}

func NewFavoriteStore() *FavoriteStore {
	return &FavoriteStore{items: make([]Favorite, 0)}
}

func (s *FavoriteStore) Add(city string) Favorite {
	favorite := Favorite{City: city}

	s.mu.Lock()
	s.items = append(s.items, favorite)
	s.mu.Unlock()

	return favorite
}

func (s *FavoriteStore) List() []Favorite {
	s.mu.RLock()
	defer s.mu.RUnlock()

	items := make([]Favorite, len(s.items))
	copy(items, s.items)
	return items
}

var store = NewFavoriteStore()
var router = setupRouter()

func setupRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/favorites", createFavorite)
		api.GET("/favorites", listFavorites)
	}

	return router
}

func createFavorite(c *gin.Context) {
	var input FavoriteCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	c.JSON(http.StatusOK, store.Add(input.City))
}

func listFavorites(c *gin.Context) {
	c.JSON(http.StatusOK, store.List())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

func main() {
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
