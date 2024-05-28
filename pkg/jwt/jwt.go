package jwt

import (
	"errors"
	"time"

	driJWT "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type JWTService interface {
	GenerateToken(userId uuid.UUID, roleId uuid.UUID) (t string, err error)
	ParseToken(tokenString string) (claims JwtCustomClaim, err error)
}

type JwtCustomClaim struct {
	UserID uuid.UUID `json:"userId"`
	RoleID uuid.UUID `json:"roleId"`
	driJWT.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
	expired   int
}

func NewJWTService(secretKey, issuer string, expired int) JWTService {
	return &jwtService{
		issuer:    issuer,
		secretKey: secretKey,
		expired:   expired,
	}
}

func (j *jwtService) GenerateToken(userId uuid.UUID, roleId uuid.UUID) (t string, err error) {
	claims := &JwtCustomClaim{
		userId,
		roleId,
		driJWT.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.expired)).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := driJWT.NewWithClaims(driJWT.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ParseToken(tokenString string) (claims JwtCustomClaim, err error) {
	if token, err := driJWT.ParseWithClaims(tokenString, &claims, func(token *driJWT.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	}); err != nil || !token.Valid {
		return JwtCustomClaim{}, errors.New("token is not valid")
	}

	return
}
