package services

import (
	"Backend-Challenge/db"
	"Backend-Challenge/models"
	"context"
	"fmt"
)

type ArticleService struct {
	ArticleRepository interface{}
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		ArticleRepository: models.NewArticleRepository(),
	}
}

func GetArticleByID(id int) (*models.Article, error) {
	db := db.GetDatabase()

	sql := `SELECT id, featured, title, url, imageUrl, newsSite, summary, publishedAt FROM articles WHERE id = $1`

	article := &models.Article{}
	err := db.QueryRow(context.Background(), sql, id).Scan(
		&article.ID,
		&article.Featured,
		&article.Title,
		&article.URL,
		&article.ImageURL,
		&article.NewsSite,
		&article.Summary,
		&article.PublishedAt,
	)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error getting the article: %s", err)
	}

	return article, nil
}

func GetAllArticles() ([]*models.Article, error) {
	db := db.GetDatabase()

	var articles []*models.Article

	sql := `SELECT id, featured, title, url, imageUrl, newsSite, summary, publishedAt FROM articles`
	rows, err := db.Query(context.Background(), sql)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error getting the articles: %s", err)
	}
	defer rows.Close()

	for rows.Next() {
		article := &models.Article{}
		err := rows.Scan(
			&article.ID,
			&article.Featured,
			&article.Title,
			&article.URL,
			&article.ImageURL,
			&article.NewsSite,
			&article.Summary,
			&article.PublishedAt,
		)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("error getting the articles: %s", err)
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func AddArticle(article *models.Article) error {
	db := db.GetDatabase()

	sql := `INSERT INTO articles (featured, title, url, imageUrl, newsSite, summary, publishedAt) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := db.QueryRow(context.Background(), sql,
		article.Featured,
		article.Title,
		article.URL,
		article.ImageURL,
		article.NewsSite,
		article.Summary,
		article.PublishedAt,
	).Scan(&article.ID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error adding the article: %s", err)
	}

	return nil
}

func UpdateArticle(article *models.Article) error {
	db := db.GetDatabase()

	sql := `UPDATE articles SET featured = $1, title = $2, url = $3, imageUrl = $4, newsSite = $5, summary = $6, publishedAt = $7 WHERE id = $8`
	_, err := db.Exec(context.Background(), sql,
		article.Featured,
		article.Title,
		article.URL,
		article.ImageURL,
		article.NewsSite,
		article.Summary,
		article.PublishedAt,
		article.ID,
	)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error updating the article: %s", err)
	}

	return nil
}

func DeleteArticle(id int) error {
	db := db.GetDatabase()

	sql := `DELETE FROM articles WHERE id = $1`
	_, err := db.Exec(context.Background(), sql, id)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error deleting the article: %s", err)
	}

	return nil
}
