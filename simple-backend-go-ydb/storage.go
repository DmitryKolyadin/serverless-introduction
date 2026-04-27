package main

import (
	"context"

	ydb "github.com/ydb-platform/gorm-driver"
	yc "github.com/ydb-platform/ydb-go-yc"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Favorite struct {
	City string `json:"city" gorm:"column:city;primaryKey"`
}

func (Favorite) TableName() string {
	return "favorites"
}

type Storage struct {
	db        *gorm.DB
	tableName string
}

func NewStorage(cfg Config) (*Storage, error) {
	db, err := gorm.Open(
		ydb.Open(cfg.DSN, ydb.With(yc.WithInternalCA(), yc.WithServiceAccountKeyFileCredentials(cfg.PathToKey))),
		&gorm.Config{},
	)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db, tableName: cfg.TableName}, nil
}

func (s *Storage) AddFavorite(ctx context.Context, city string) (Favorite, error) {
	f := Favorite{City: city}
	err := s.db.WithContext(ctx).
		Table(s.tableName).
		Clauses(clause.OnConflict{DoNothing: true}).
		Create(&f).Error
	if err != nil {
		return Favorite{}, err
	}
	return f, nil
}

func (s *Storage) ListFavorites(ctx context.Context) ([]Favorite, error) {
	var items []Favorite
	err := s.db.WithContext(ctx).Table(s.tableName).Find(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}
