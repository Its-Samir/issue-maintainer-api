package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

/* jwt secret key */
const secretKey = "secret key for jwt"

func GenerateJwtToken(email string, userId int64) (string, error) {
	/* add data to MapClaims{} to extract and use them later */
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})

	signedToken, err := token.SignedString([]byte(secretKey))

	return signedToken, err
}

func VerifyJwtToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Signin method Invalid")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("Could not parse the token")
	}

	/* check if token is valid or not */
	isTokenValid := parsedToken.Valid

	if !isTokenValid {
		return 0, errors.New("Invalid token")
	}

	/* extract the claims from parsed token */
	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Invalid claims")
	}

	/* extract the userId from the claims */
	userId := int64(claims["userId"].(float64))

	/* we need this userId in order to perform certain tasks conditionally inside the handlers */
	return userId, nil
}
