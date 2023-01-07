package controller

import (
	"net/http"
	"peanut/domain"
	"peanut/repository"
	"peanut/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ContentController struct {
	Usecase usecase.ContentUsecase
}

func NewContentController(db *gorm.DB) *ContentController {
	return &ContentController{
		Usecase: usecase.NewContentUsecase(repository.NewContentRepo(db)),
	}
}

func (c *ContentController) GetContents(ctx *gin.Context) {
	contents, err := c.Usecase.GetContents()
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{
			"message": "not found any record",
		})
		return
	}
	ctx.JSON(http.StatusOK, domain.Response{
		Success: true, Data: contents, Message: "Get data successful",
	})
}

func (c *ContentController) CreateContent(ctx *gin.Context) {
	newContent := domain.Content{}
	if !bindJSON(ctx, &newContent) {
		return
	}

	book, err := c.Usecase.CreateContent(newContent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.Response{true, book, "create book success"})
}
