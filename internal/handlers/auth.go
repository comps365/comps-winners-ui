package handlers

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

var (
	errUnauthorised = errors.New("you shall not pass")
)

type AuthHandler struct {
	signingKey []byte
	username   string
	password   string
}

func NewAuthHandler(k []byte, u, p string) *AuthHandler {
	return &AuthHandler{
		signingKey: k,
		username:   u,
		password:   p,
	}
}

type JWTCustomClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type authRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type authResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Login(c echo.Context) error {
	var req authRequest
	if err := c.Bind(&req); err != nil {
		return notOk(c, http.StatusBadRequest, err)
	}

	// username and password are passed from the config
	// this is a single user application so ...
	// there is no point in connecting to a db, do user management etc.
	// keep it simple, stupid :)
	if req.Username != h.username || req.Password != h.password {
		return notOk(c, http.StatusUnauthorized, errUnauthorised)
	}

	claims := &JWTCustomClaims{
		req.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(h.signingKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &authResponse{
		Token: t,
	})
}
