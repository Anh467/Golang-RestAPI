package common

import (
	"entities"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user *entities.UserCreateModel) (string, error) {
	// create claims
	claims := &entities.Claims{
		UserID:   user.UserID,
		Email:    user.Email,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Hạn sử dụng JWT trong 1 ngày.
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// create token sign
	tokenString, err := token.SignedString(JWT_SECRET_KEY)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func PraseToken(tokenString string) (*entities.Claims, error) {
	// Parse token từ chuỗi tokenString.
	token, err := jwt.ParseWithClaims(tokenString, &entities.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET_KEY, nil
	})
	if err != nil {
		return nil, err
	}
	// Kiểm tra lỗi và xác minh token.
	if claims, ok := token.Claims.(*entities.Claims); ok && token.Valid {
		// Xử lý dữ liệu từ claims (email và password).
		return claims, nil
	}
	// Xác thực không thành công.
	return nil, errors.New(TOKEN_NOT_APPROPRIATE)

}
