package handler

import (
	"log"
	"net/http"
	"test-praweda-app-eng/internal/model"
	"test-praweda-app-eng/internal/repository"
	"test-praweda-app-eng/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	ar *repository.AuthRepository
}

func NewAuthHandler(ar *repository.AuthRepository) *AuthHandler {
	return &AuthHandler{ar: ar}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var body model.LoginBody

	if err := ctx.ShouldBindJSON(&body); err != nil || body.Username == "" {
		log.Println("failed binding cause invalid username:", err)
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code:      http.StatusBadRequest,
			IsSuccess: false,
			Err:       "Invalid username",
		})
		return
	}

	user, err := h.ar.FindOrCreate(ctx, body.Username)
	if err != nil {
		log.Println("db error:", err)
		ctx.JSON(http.StatusInternalServerError, model.Response{
			IsSuccess: false,
			Code:      http.StatusInternalServerError,
			Err:       "db error",
		})
		return
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		log.Println("error generate token:", err)
		ctx.JSON(http.StatusInternalServerError, model.Response{
			IsSuccess: false,
			Code:      http.StatusInternalServerError,
			Err:       "error generate token",
		})
		return
	}

	// Set HTTP-only cookie
	ctx.SetCookie(
		body.Username,
		token,
		86400,
		"/",
		"",
		false,
		true,
	)

	ctx.JSON(http.StatusOK, model.ResponseData{
		Response: model.Response{
			IsSuccess: true,
			Code:      http.StatusOK,
			Msg:       "Login successful",
		},
		Data: gin.H{
			"user":  user,
			"token": token,
		},
	})
}

func (h *AuthHandler) Me(ctx *gin.Context) {
	var body model.LoginBody
	if err := ctx.ShouldBindJSON(&body); err != nil || body.Username == "" {
		log.Println("failed binding cause invalid username:", err)
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code:      http.StatusBadRequest,
			IsSuccess: false,
			Err:       "Invalid username",
		})
		return
	}

	cookie, err := ctx.Cookie(body.Username)
	if err != nil {
		log.Println("no cookie found:", err)
		ctx.JSON(http.StatusUnauthorized, model.Response{
			IsSuccess: false,
			Code:      http.StatusUnauthorized,
			Err:       "unauthorized",
		})
		return
	}

	ctx.JSON(http.StatusOK, model.ResponseData{
		Response: model.Response{
			IsSuccess: true,
			Code:      http.StatusOK,
			Msg:       "ok",
		},
		Data: gin.H{
			"cookie": cookie,
		},
	})
}

func (h *AuthHandler) Logout(ctx *gin.Context) {
	var body model.LoginBody
	if err := ctx.ShouldBindJSON(&body); err != nil || body.Username == "" {
		log.Println("failed binding cause invalid username:", err)
		ctx.JSON(http.StatusBadRequest, model.Response{
			Code:      http.StatusBadRequest,
			IsSuccess: false,
			Err:       "Invalid username",
		})
		return
	}

	ctx.SetCookie(body.Username, "", -1, "/", "", false, true)
	ctx.JSON(http.StatusOK, model.Response{
		IsSuccess: true,
		Code:      http.StatusOK,
		Msg:       "logged out",
	})

}
