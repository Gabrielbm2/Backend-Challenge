package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Article struct {
	ID          int      `json:"id"`
	Featured    bool     `json:"featured"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
	ImageURL    string   `json:"imageurl"`
	NewsSite    string   `json:"newsSite"`
	Summary     string   `json:"summary"`
	PublishedAt string   `json:"publishedAt"`
	Launches    []Launch `json:"launches"`
	Events      []Event  `json:"events"`
}

type Launch struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
}

type Event struct {
	ID       string `json:"id"`
	Provider string `json:"provider"`
}

func CreateArticleTable() error {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=gbm158545 dbname=Article sslmode=disable")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	sql := `
	CREATE TABLE IF NOT EXISTS articles (
		id SERIAL PRIMARY KEY,
		featured BOOLEAN,
		title TEXT,
		url TEXT,
		image_url TEXT,
		news_site TEXT,
		summary TEXT,
		published_at TEXT
	)
	`
	_, err = db.Exec(sql)
	if err != nil {
		return err
	}

	return nil
}

type ArticleRepository struct {
	db *sql.DB
}

func NewArticleRepository() *ArticleRepository {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=gbm158545 dbname=Article sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	return &ArticleRepository{
		db: db,
	}
}
