package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AuthResponse struct {
	User                 User      `json:"user"`
	AccessToken          string    `json:"access_token"`
	RefreshToken         string    `json:"refresh_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

type SignUpRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (srv *Server) signUp(ctx *gin.Context) {
	var signUpRequest SignUpRequest

	if err := ctx.ShouldBindJSON(&signUpRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	hashedPw, err := utils.HashPassword(signUpRequest.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	user, err := srv.store.CreateUser(ctx, db.CreateUserParams{
		Username: signUpRequest.Username,
		Email:    signUpRequest.Email,
		Password: hashedPw,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	aPayload, accessTkn, err := srv.tokenMaker.CreateToken(user.Username, user.ID, user.Email, time.Hour)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	_, refreshTkn, err := srv.tokenMaker.CreateToken(user.Username, user.ID, user.Email, time.Hour*24)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	resCreateUser := AuthResponse{
		User: User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
		AccessToken:          accessTkn,
		RefreshToken:         refreshTkn,
		AccessTokenExpiresAt: aPayload.ExpiredAt,
	}

	ctx.JSON(http.StatusCreated, resCreateUser)
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (srv *Server) login(ctx *gin.Context) {
	var authRequest LoginRequest

	if err := ctx.ShouldBindJSON(&authRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(err))
		return
	}

	user, err := srv.store.FindUserByEmail(ctx, authRequest.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errJSON(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	err = utils.IsPassword(authRequest.Password, user.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errJSON(err))
		return
	}

	aPayload, accessTkn, err := srv.tokenMaker.CreateToken(user.Username, user.ID, user.Email, time.Hour)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	_, refreshTkn, err := srv.tokenMaker.CreateToken(user.Username, user.ID, user.Email, time.Hour*24)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}
	loginResponse := AuthResponse{
		User: User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
		AccessToken:          accessTkn,
		RefreshToken:         refreshTkn,
		AccessTokenExpiresAt: aPayload.ExpiredAt,
	}

	ctx.JSON(http.StatusOK, loginResponse)
}
