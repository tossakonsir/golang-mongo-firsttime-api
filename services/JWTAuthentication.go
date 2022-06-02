package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
)

type JWTService interface {
	GenerateToken(username string, isUser bool) error
	StoreJWTAuthToRedis(username string, token string) error
	GetJWTAuthFromRedis(username string) string
	VerifyToken(token string) error
}
type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Bikash",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(username string, isUser bool) error {
	claims := &authCustomClaims{
		username,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	service.StoreJWTAuthToRedis(username, t)
	return nil
}

func (service *jwtServices) StoreJWTAuthToRedis(username string, token string) error {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.TODO()
	err := rdb.Set(ctx, username, token, 0).Err()

	if err != nil {

		log.Fatal(err)

	}

	return nil
}

func (service *jwtServices) GetJWTAuthFromRedis(username string) string {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	ctx := context.TODO()
	val, err := rdb.Get(ctx, username).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val
}

func (maker *jwtServices) VerifyToken(token string) error {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("InvalidToken")
		}
		return []byte(maker.secretKey), nil
	}

	_, err := jwt.Parse(token, keyFunc)
	if err != nil {
		return err
	}
	return nil
}

func maintest() int {
	a := 1
	b := 2
	c := a
	d := b
	return c + d
}
