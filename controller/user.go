package controller

import (
	"net/http"
	"peanut/domain"
	"peanut/usecase"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Usecase usecase.UserUsecase
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		Usecase: usecase.NewUserUsecase(db),
	}
}

func (c *UserController) GetUsers(ctx *gin.Context) {

}

func (c *UserController) GetUser(ctx *gin.Context) {

}

// Register godoc
//
//	@Summary		register
//	@Description	register
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			json	body		domain.User	true	"Body"
//	@Success		200		{object}	domain.Response
//	@Failure		400		{object}	domain.ErrorResponse
//	@Failure		500		{object}	domain.Response
//	@Router			/signup [post]
func (c *UserController) CreateUser(ctx *gin.Context) {
	user := domain.User{}

	if !bindJSON(ctx, &user) {
		return
	}

	err := c.Usecase.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			domain.Response{false, nil, "create user failed"},
		)
	}

	ctx.JSON(http.StatusOK,
		domain.Response{true, user, "create user successfuly"},
	)
}

// Login godoc
//
//	@Summary		login
//	@Description	login
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			json	body		domain.LoginForm	true	"Body"
//	@Success		200		{object}	domain.Response
//	@Failure		400		{object}	domain.Response
//	@Failure		500		{object}	domain.Response
//	@Router			/login [post]
func (c *UserController) Login(ctx *gin.Context) {
	var loginForm domain.LoginForm
	if !bindJSON(ctx, &loginForm) {
		return
	}
	tokenString, errRes := c.Usecase.Login(ctx, loginForm)
	if errRes != nil {
		if errRes.Code == "400" {
			ctx.JSON(http.StatusBadRequest,
				domain.Response{false, nil, errRes.DebugMessage},
			)
			return
		} else if errRes.Code == "500" {
			ctx.JSON(http.StatusInternalServerError,
				domain.Response{false, nil, errRes.DebugMessage},
			)
			return
		}

	}
	ctx.JSON(http.StatusOK,
		domain.Response{true, tokenString, "Login success"})
}
