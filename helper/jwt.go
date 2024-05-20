package helper

import (
	"backendgreeve/constant"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTInterface interface {
	GenerateUserJWT(user UserJWT) (string, error)
	GenerateAdminJWT(user AdminJWT) (string, error)
	GenerateUserToken(user UserJWT) string
	GenerateAdminToken(user AdminJWT) string
	ExtractUserToken(token *jwt.Token) map[string]interface{}
	ExtractAdminToken(token *jwt.Token) map[string]interface{}
	ValidateToken(token string) (*jwt.Token, error)
}

type UserJWT struct {
	ID       string
	Name     string
	Email    string
	Username string
	Address  string
	Role     string
}

type AdminJWT struct {
	ID       string
	Name     string
	Email    string
	Username string
	Role     string
}
type JWT struct {
	signKey string
}

func NewJWT(signKey string) JWTInterface {
	return &JWT{
		signKey: signKey,
	}
}

func (j *JWT) GenerateUserJWT(user UserJWT) (string, error) {
	var accessToken = j.GenerateUserToken(user)
	if accessToken == "" {
		return "", constant.ErrGenerateJWT
	}

	return accessToken, nil
}

func (j *JWT) GenerateUserToken(user UserJWT) string {
	var claims = jwt.MapClaims{}
	claims[constant.JWT_ID] = user.ID
	claims[constant.JWT_NAME] = user.Name
	claims[constant.JWT_EMAIL] = user.Email
	claims[constant.JWT_USERNAME] = user.Username
	claims[constant.JWT_ROLE] = constant.RoleUser
	claims[constant.JWT_ADDRESS] = user.Address
	claims[constant.JWT_IAT] = time.Now().Unix()
	claims[constant.JWT_EXP] = time.Now().Add(time.Hour * 24 * 31).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(j.signKey))

	if err != nil {
		return ""
	}

	return validToken
}

func (j *JWT) ExtractUserToken(token *jwt.Token) map[string]interface{} {
	if token.Valid {
		var claims = token.Claims
		expTime, _ := claims.GetExpirationTime()
		if expTime.Time.Compare(time.Now()) > 0 {
			var mapClaim = claims.(jwt.MapClaims)
			var result = map[string]interface{}{}
			result[constant.JWT_ID] = mapClaim[constant.JWT_ID]
			result[constant.JWT_NAME] = mapClaim[constant.JWT_NAME]
			result[constant.JWT_EMAIL] = mapClaim[constant.JWT_EMAIL]
			result[constant.JWT_USERNAME] = mapClaim[constant.JWT_USERNAME]
			result[constant.JWT_ADDRESS] = mapClaim[constant.JWT_ADDRESS]
			result[constant.JWT_ROLE] = mapClaim[constant.JWT_ROLE]
			return result
		}
		return nil
	}
	return nil
}

func (j *JWT) GenerateAdminJWT(user AdminJWT) (string, error) {
	var accessToken = j.GenerateAdminToken(user)
	if accessToken == "" {
		return "", constant.ErrGenerateJWT
	}

	return accessToken, nil
}

func (j *JWT) GenerateAdminToken(user AdminJWT) string {
	var claims = jwt.MapClaims{}
	claims[constant.JWT_ID] = user.ID
	claims[constant.JWT_NAME] = user.Name
	claims[constant.JWT_EMAIL] = user.Email
	claims[constant.JWT_USERNAME] = user.Username
	claims[constant.JWT_ROLE] = constant.RoleAdmin
	claims[constant.JWT_IAT] = time.Now().Unix()
	claims[constant.JWT_EXP] = time.Now().Add(time.Hour * 24 * 31).Unix()

	var sign = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := sign.SignedString([]byte(j.signKey))

	if err != nil {
		return ""
	}

	return validToken
}

func (j *JWT) ExtractAdminToken(token *jwt.Token) map[string]interface{} {
	if token.Valid {
		var claims = token.Claims
		expTime, _ := claims.GetExpirationTime()
		if expTime.Time.Compare(time.Now()) > 0 {
			var mapClaim = claims.(jwt.MapClaims)
			var result = map[string]interface{}{}
			result[constant.JWT_ID] = mapClaim[constant.JWT_ID]
			result[constant.JWT_NAME] = mapClaim[constant.JWT_NAME]
			result[constant.JWT_EMAIL] = mapClaim[constant.JWT_EMAIL]
			result[constant.JWT_USERNAME] = mapClaim[constant.JWT_USERNAME]
			result[constant.JWT_ROLE] = mapClaim[constant.JWT_ROLE]
			return result
		}
		return nil
	}
	return nil
}

func (j *JWT) ValidateToken(token string) (*jwt.Token, error) {

	var authHeader = token[7:]
	parsedToken, err := jwt.Parse(authHeader, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(j.signKey), nil
	})
	if err != nil {
		return nil, constant.ErrValidateJWT
	}
	return parsedToken, nil
}
