package respositories

import "blog/internal/models"

type ArticleRespositoryInterface interface {
	List(limit int) []models.Article
	Find(id int) models.Article
	Create(article models.Article) models.Article
}
