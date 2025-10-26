package helpers

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func getSecretKey() (key string) {
	key = os.Getenv("SECRET_KEY")
	return key
}

func GenerateToken(username, apiKey string) (res string, err error) {
	// ambil secret key dari function
	key := getSecretKey()

	// masukkan username dan apikey
	claims := jwt.MapClaims{
		"username": username,
		"api_key":  apiKey,
		"ll":       time.Now(),
	}

	// encrypt claims ke sha256 (token belum dittd)
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// ttd token yang sudah diencrypt dengan key(secret key)
	signedToken, err := parseToken.SignedString([]byte(key))
	if err != nil {
		log.Println("Error parse")
		return
	}

	res = signedToken

	return
}

func VerifyToken(ctx *gin.Context) (res interface{}, err error) {
	// ambil secret key
	key := getSecretKey()

	// get header token dengan nama Authorization
	headerToken := ctx.Request.Header.Get("Authorization")

	// ambil prefix/ index pertama dari string dengan nama Bearer
	bearer := strings.HasPrefix(headerToken, "Bearer")
	// jika bukan Bearer akan error
	if !bearer {
		err = errors.New("Invalid header authorization")
		return
	}

	// ambil index ke 1 dari header token => "Bearer token"
	stringToken := strings.Split(headerToken, " ")[1]

	// check string token apakah asli dan validitas jwt
	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (res interface{}, err error) {
		// t.Method adalah objek jwt.SigningMethod yang merepresentasikan algoritma yang disebut di header JWT (alg).
		// t.Method.(*jwt.SigningMethodHMAC) melakukan type assertion untuk melihat apakah method itu tipe HMAC (yang dipakai untuk HS256/HS384/HS512).
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			err = errors.New("Invalid method header alg")
			return
		}

		// jika ok maka nilai res digunakan oleh jwt.Parse sebagai key untuk memverifikasi signature token (untuk HMAC, key harus berupa byte slice).
		// Jika verifikasi signature cocok, jwt.Parse akan mengembalikan token yang valid; jika tidak cocok jwt.Parse mengembalikan error.
		res, err = []byte(key), nil
		return
	})

	// apakah token masih valid
	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		err = errors.New("token invalid")
		return
	}

	// ambil data claims nya
	res = token.Claims.(jwt.MapClaims)
	return
}
