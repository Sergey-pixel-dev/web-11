package api

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type envelope map[string]string

func (srv *Server) PostCount(e echo.Context) error {
	i2 := envelope{"count": "0"}
	err := e.Bind(&i2)
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	err = srv.uc.FetchHelloMessage(i2["count"])
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, "OK")
}
func (srv *Server) GetCount(e echo.Context) error {
	msg, err := srv.uc.GetHelloMessage()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, msg)
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var jwtSecret = []byte("uraaaa") // Секретный ключ
// Генерация токена
func generateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Срок действия токена — 24 часа
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// Проверка токена
func validateToken(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Проверка метода подписи токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
}

// Обработчик логина
func (srv *Server) Login(c echo.Context) error {
	req := new(LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}

	// Пример проверки логина и пароля
	if req.Username != "admin" || req.Password != "password" {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
	}

	// Генерация токена
	token, err := generateToken(1) // userID = 1 (пример)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

// Middleware для проверки токена
func JWTMiddleware() echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "missing token"})
			}

			// Проверка токена
			token, err := validateToken(authHeader)
			if err != nil || !token.Valid {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid token"})
			}

			// Передача токена в контекст
			c.Set("user", token)
			return next(c)
		}
	})
}
