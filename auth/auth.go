package auth

import (
    "context"
    "os"
    "fmt"
    "time"
    "net/http"
    "strings"

    "github.com/dgrijalva/jwt-go"
    "golang.org/x/crypto/bcrypt"
)

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

func UserCtx(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    token, err := VerifyToken(r)
    if err != nil{
        token = nil
    }
    ctx := context.WithValue(r.Context(), "token", token)
    next.ServeHTTP(w, r.WithContext(ctx))
    })
}
