package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
    "server/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v10"
	"golang.org/x/crypto/bcrypt"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var userCtxKey = &contextKey{"user"}
type contextKey struct {
	name string
}

// A stand-in for our database backed user object

func CreateToken(userId string) (string, error) {
  var err error
  //Creating Access Token
  atClaims := jwt.MapClaims{}
  atClaims["authorized"] = true
  atClaims["user_id"] = userId
  atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
  at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
  token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
  if err != nil {
     return "", err
  }
  return token, nil
}

func ExtractToken(r *http.Request) string {
  bearToken := r.Header.Get("Authorization")
  //normally Authorization the_token_xxx
  strArr := strings.Split(bearToken, " ")
  if len(strArr) == 2 {
     return strArr[1]
  }
  return ""
}

func ComparePasswords (hashedpass string, plainPass string) bool{
     byteHash := []byte(hashedpass)
     err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPass))
    if err != nil {
        fmt.Println(err)
        return false
    }
    
    return true
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
  tokenString := ExtractToken(r)
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
     //Make sure that the token method conform to "SigningMethodHMAC"
     if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
     }
     return []byte(os.Getenv("ACCESS_SECRET")), nil
  })
  if err != nil {
     return nil, err
  }
  return token, nil
}

func UserCtx(db *pg.DB) func(http.Handler) http.Handler {
  return func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token, err := VerifyToken(r)
        if err != nil || token == nil{
            next.ServeHTTP(w, r) 
            return
        } 

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            userid := claims["user_id"]
            user := new(models.User)
            err := db.Model(user).Where("ID = ?", userid).Select()
            if err != nil{
                fmt.Println("A database error occured")
                next.ServeHTTP(w, r) 
                return
            }
            ctx := context.WithValue(r.Context(), userCtxKey, user)
            next.ServeHTTP(w, r.WithContext(ctx)) 
        } else {
            next.ServeHTTP(w, r) 
            return
        }
    })
  }
}

func ForContext(ctx context.Context) *models.User {
	raw, _ := ctx.Value(userCtxKey).(*models.User)
	return raw
}

