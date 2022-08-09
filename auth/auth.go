package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

func TokenCreate(userID uint64) (string, error) {
	perms := jwt.MapClaims{}
	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 6).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)
	return token.SignedString([]byte("foobar"))
}

func TokenValidate(r *http.Request) error {
	stringToken := tokenExtract(r)
	token, err := jwt.Parse(stringToken, returnKey)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token inválido")
}

func UserIdExtract(r *http.Request) (uint64, error) {
	stringToken := tokenExtract(r)
	token, err := jwt.Parse(stringToken, returnKey)
	if err != nil {
		return 0, err
	}
	if perms, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", perms["userId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return usuarioID, nil
	}

	return 0, errors.New("token inválido")
}

func tokenExtract(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func returnKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("método de assinatura inesperado! %v", token.Header["alg"])
	}
	return "foobar", nil
}
