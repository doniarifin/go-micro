package helper

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type AuthUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func GenerateToken(s *AuthUser) (string, error) {
	mySigningKey := []byte("AllYourBase")
	if s.Role == "" {
		s.Role = "user"
	}
	//create claims
	claims := &CustomClaims{
		s.Email,
		s.Role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(1516239022, 0)),
			Issuer:    s.ID,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	if err != nil {
		return "error generate jwt: ", err
	}

	return ss, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(tokenString *jwt.Token) (any, error) {
		return []byte("AllYourBase"), nil
	})

	if err != nil {
		log.Fatal(err)
		return err
	} else if _, ok := token.Claims.(*CustomClaims); ok {
		// fmt.Println(claims.Email, claims.Issuer)
		return nil
	} else {
		return errors.New("unknown claims type, failed proceed")
	}
}
