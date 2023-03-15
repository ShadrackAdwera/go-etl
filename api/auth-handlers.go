package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type AuthResponse struct {
	User                 User      `json:"user,omitempty"`
	AccessToken          string    `json:"access_token,omitempty"`
	RefreshToken         string    `json:"refresh_token,omitempty"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at,omitempty"`
	SessionID            uuid.UUID `json:"session_id,omitempty"`
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

	_, err = srv.store.FindUserByEmail(ctx, signUpRequest.Email)

	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusInternalServerError, errJSON(err))
			return
		}
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

	resCreateUser := AuthResponse{
		User: User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
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

	pL, refreshTkn, err := srv.tokenMaker.CreateToken(user.Username, user.ID, user.Email, time.Hour*24)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	sess, err := srv.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           pL.TokenId,
		Username:     user.Username,
		UserID:       user.ID,
		RefreshToken: refreshTkn,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    pL.ExpiredAt,
	})

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
		SessionID:            sess.ID,
	}

	ctx.JSON(http.StatusOK, loginResponse)
}
