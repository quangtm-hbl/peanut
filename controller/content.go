package controller

import (
	"net/http"
	"peanut/domain"
	"peanut/pkg/gcloudservices"
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

// ListContent godoc
//
//	@Summary		content
//	@Description	content
//	@Tags			Content
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]domain.Content
//	@Failure		400	{object}	domain.ErrorResponse
//	@Failure		500	{object}	domain.ErrorResponse
//	@Router			/contents [get]
func (c *ContentController) GetContents(ctx *gin.Context) {
	contents, err := c.Usecase.GetContents()
	if err != nil {
		ctx.JSON(http.StatusNoContent, gin.H{
			"message": "Not found any record",
		})
		return
	}
	ctx.JSON(http.StatusOK, domain.Response{
		Success: true, Data: contents, Message: "Get data successful",
	})
}

// CreateContent godoc
//
//	@Summary		content
//	@Description	content
//	@Tags			Content
//	@Accept			json
//	@Produce		json
//	@Param			Thumbnail	formData	file	true	"file"
//	@Param			Media		formData	file	true	"file"
//	@Param			Name		formData	string	true	"string"	minlength(1)	maxlength(30)
//	@Param			Description	formData	string	true	"string"	minlength(0)	maxlength(500)
//	@Param			PlayTime	formData	int		true	"int"
//	@Param			Resolution	formData	int		true	"int"
//	@Param			ARheight	formData	int		true	"int"
//	@Param			ARwidth		formData	int		true	"int"
//	@Param			Fever		formData	boolean	false	"boolean"
//	@Param			Ondemand	formData	boolean	false	"boolean"
//	@Success		200			{object}	domain.Content
//	@Failure		400			{object}	domain.ErrorResponse
//	@Failure		500			{object}	domain.ErrorResponse
//	@Router			/contents [post]
func (c *ContentController) CreateContent(ctx *gin.Context) {
	var req domain.CreateContentRequest
	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
	}

	uploader := gcloudservices.NewClientUploader()
	uploader.SetDstPath("thumbnail/")
	thumbnailURL, err := uploader.UploadFile(req.Thumbnail)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
	}
	uploader.SetDstPath("media/")
	mediaURL, err := uploader.UploadFile(req.Media)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
	}

	newContent := domain.Content{
		Thumbnail:   thumbnailURL,
		Media:       mediaURL,
		Name:        req.Name,
		Description: req.Description,
		Playtime:    req.Playtime,
		Resolution:  req.Resolution,
		ARwidth:     req.ARwidth,
		ARheight:    req.ARheight,
		Ondemand:    *req.Ondemand,
		Fever:       *req.Fever,
	}
	content, err := c.Usecase.CreateContent(newContent)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.Response{false, nil, err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, domain.Response{true, content, "create content successfully"})
}
