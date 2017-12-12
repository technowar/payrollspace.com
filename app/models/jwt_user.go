package models

import (
	"github.com/XanderDwyl/payrollspace.com/app/config"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTUser ...
type JWTUser struct {
	ID             int64  `json:"id"  gorm:"AUTO_INCREMENT"`
	UserID         int64  `json:"user_id" gorm:"unique_index"`
	ExpiresAt      int64  `json:"expires_at,omitempty"`
	Username       string `json:"username,omitempty"`
	FullName       string `json:"full_name,omitempty"`
	ProfilePicture string `json:"profile_picture,omitempty"`
	AccessToken    string `json:"access_token,omitempty"`
	Email          string `json:"email,omitempty"`
	jwt.StandardClaims
}

// CreateJWToken ...
func (u *JWTUser) CreateJWToken() (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), u)
	tokenString, err := token.SignedString([]byte(config.GetJWTSalt()))

	return tokenString, err
}
