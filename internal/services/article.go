package services

import (
	"blog/internal/models"
	"blog/internal/requests"
	"blog/internal/responses"
	"blog/internal/respositories"
	"blog/pkg/sessions"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleService struct {
	articleRespository respositories.ArticleRespositoryInterface
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		articleRespository: respositories.NewArticleRespository(),
	}
}

func (articleService *ArticleService) GetFeaturedArticles() responses.Articles {
	articles := articleService.articleRespository.List(4)
	return responses.ToArticles(articles)
}

func (articleService *ArticleService) GetStoriesArticles() responses.Articles {
	articles := articleService.articleRespository.List(8)
	return responses.ToArticles(articles)
}

func (articleService *ArticleService) Find(id int) (responses.Article, error) {
	var response responses.Article

	article := articleService.articleRespository.Find(id)

	if article.ID == 0 {
		return response, errors.New("article not found")
	}

	return responses.ToArticle(article), nil
}

func (articleService *ArticleService) CreateArticle(c *gin.Context, request requests.CreateArticle) (responses.Article, error) {
	var response responses.Article

	authID := sessions.Get(c, "auth")
	userID, _ := strconv.Atoi(authID)
	user := respositories.NewUserRespository().FindByID(userID)

	if user.ID == 0 {
		return response, errors.New("article could not be created")
	}

	newArticle := models.Article{
		Title:   request.Title,
		Content: request.Content,
		UserId:  user.ID,
	}

	article := articleService.articleRespository.Create(newArticle)

	if article.ID == 0 {
		return response, errors.New("article could not be created")
	}

	return responses.ToArticle(article), nil
}
