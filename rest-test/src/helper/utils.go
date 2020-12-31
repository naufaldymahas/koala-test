package helper

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func HSHA256Handler(p string, s []byte) string {
	h := hmac.New(sha256.New, s)
	h.Write([]byte(p))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateRandomSalt(saltSize int) []byte {
	var salt = make([]byte, saltSize)

	_, err := rand.Read(salt[:])

	if err != nil {
		panic(err)
	}

	return salt
}

// tokenType 1 refresh token else normal token
func GenerateJWT(id string, tokenType int32) string {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	EXP_NORMAL, err1 := strconv.Atoi(os.Getenv("EXP_NORMAL_TOKEN_JWT"))
	EXP_REFRESH, err2 := strconv.Atoi(os.Getenv("EXP_REFRESH_TOKEN_JWT"))

	if err1 != nil || err2 != nil {
		log.Panicln("Wrong format expired token jwt at .env")
	}

	var exp int64
	if tokenType == 1 {
		exp = time.Now().Add(time.Hour * time.Duration(EXP_REFRESH)).Unix()
	} else {
		exp = time.Now().Add(time.Minute * time.Duration(EXP_NORMAL)).Unix()
	}

	SECRET := os.Getenv("SALT_JWT")

	claims := &jwt.StandardClaims{
		ExpiresAt: exp,
		Subject:   id,
		Issuer:    "naufaldymahas@gmail.com",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(SECRET))
	if err != nil {
		log.Println(err)
	}
	return ss
}
