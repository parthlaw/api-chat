package routes

import (
	"api-chat/utils"
	"net/http"
	"strings"

	"fmt"

	"github.com/dgrijalva/jwt-go"
)
func Middleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token:=r.Header.Get("Authorization")
		if token==""{
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
			return
		}
		if strings.Contains(token, "Bearer ") {
			token=strings.Split(token,"Bearer ")[1]
		}
		
		reqToken:=token
		key, er:=jwt.ParseRSAPublicKeyFromPEM([]byte(utils.PublicKey))
		if er!=nil{
			fmt.Println("validate: parse key:", er)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
			return
		}
		tok,err:=jwt.Parse(reqToken,func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected method: %s", token.Header["alg"])
			}
			return key, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized")
			fmt.Println("validate: %w", err)
			return
		}
		fmt.Println(tok.Claims)
		next.ServeHTTP(w, r)
	})
}