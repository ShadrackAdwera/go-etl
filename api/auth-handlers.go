package api

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/ShadrackAdwera/go-etl/db/sqlc"
	"github.com/ShadrackAdwera/go-etl/token"
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
	User                  User      `json:"user"`
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	SessionID             uuid.UUID `json:"session_id"`
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

	ctx.JSON(http.StatusCreated, User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	})
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

	aPayload, accessTkn, err := srv.tokenMaker.CreateToken(user.Username, user.ID, user.Email, time.Minute*15)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	rPayload, refreshTkn, err := srv.tokenMaker.CreateToken(user.Username, user.ID, user.Email, time.Hour*24)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	session, err := srv.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           rPayload.TokenId,
		Username:     rPayload.Username,
		UserID:       rPayload.ID,
		RefreshToken: refreshTkn,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    rPayload.ExpiredAt,
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
		AccessToken:           accessTkn,
		RefreshToken:          refreshTkn,
		AccessTokenExpiresAt:  aPayload.ExpiredAt,
		RefreshTokenExpiresAt: rPayload.ExpiredAt,
		SessionID:             session.ID,
	}

	ctx.JSON(http.StatusOK, loginResponse)
}

type RefreshTokenArgs struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

func (srv *Server) renewAccessToken(ctx *gin.Context) {
	var refreshTokenArgs RefreshTokenArgs

	if err := ctx.ShouldBindJSON(&refreshTokenArgs); err != nil {
		ctx.JSON(http.StatusBadRequest, errJSON(errors.New("invalid request")))
		return
	}

	aTknPayload := ctx.MustGet(authPayload).(*token.TokenPayload)

	rTknPayload, err := srv.tokenMaker.VerifyToken(refreshTokenArgs.RefreshToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errJSON(fmt.Errorf("%s. login to create a new session", err)))
		return
	}

	if aTknPayload.ID != rTknPayload.ID {
		ctx.JSON(http.StatusUnauthorized, errJSON(fmt.Errorf("you are not authorized to make this request")))
		return
	}

	foundSess, err := srv.store.GetSession(ctx, rTknPayload.TokenId)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errJSON(fmt.Errorf("this session is not valid")))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	if foundSess.IsBlocked {
		ctx.JSON(http.StatusUnauthorized, errJSON(fmt.Errorf("your account has been blocc'd")))
		return
	}

	if foundSess.ID != rTknPayload.TokenId {
		ctx.JSON(http.StatusUnauthorized, errJSON(fmt.Errorf("your account has been blocc'd")))
		return
	}

	accessTknPayload, tkn, err := srv.tokenMaker.CreateToken(aTknPayload.Username, aTknPayload.ID, aTknPayload.Email, time.Minute*15)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errJSON(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":        tkn,
		"access_token_expiry": accessTknPayload.ExpiredAt,
	})

}
