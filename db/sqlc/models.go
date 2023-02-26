// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0

package db

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID          int64     `json:"id"`
	FileUrl     string    `json:"file_url"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedByID int64     `json:"created_by_id"`
}

type MatchDatum struct {
	ID          int64     `json:"id"`
	HomeScored  int32     `json:"home_scored"`
	AwayScored  int32     `json:"away_scored"`
	HomeTeam    string    `json:"home_team"`
	AwayTeam    string    `json:"away_team"`
	MatchDate   time.Time `json:"match_date"`
	Referee     string    `json:"referee"`
	Winner      string    `json:"winner"`
	Season      string    `json:"season"`
	CreatedAt   time.Time `json:"created_at"`
	CreatedByID int64     `json:"created_by_id"`
	FileID      int64     `json:"file_id"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type User struct {
	ID                int64     `json:"id"`
	Username          string    `json:"username"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
