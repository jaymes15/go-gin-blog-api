package respositories

import (
	"blog/internal/models"
	"blog/pkg/database"

	"gorm.io/gorm"
)

type ArticleRespository struct {
	DB *gorm.DB
}

func NewArticleRespository() *ArticleRespository {
	return &ArticleRespository{
		DB: database.Connection(),
	}

}

func (articleRespository *ArticleRespository) Create(article models.Article) models.Article {
	var newArticle models.Article

	articleRespository.DB.Create(&article).Scan(&newArticle)

	return newArticle
}

func (articleRespository *ArticleRespository) List(limit int) []models.Article {
	var articles []models.Article

	articleRespository.DB.Limit(limit).Order("created_at").Find(&articles)

	return articles
}

func (articleRepository *ArticleRespository) Find(id int) models.Article {
	var article models.Article

	articleRepository.DB.Joins("User").First(&article, id)

	return article
}
