package routes

import (
	"api-chat/utils"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Error bool `json:"error"`
	Data utils.User `json:"data"`
}
const baseUrl= "https://staging.l-earnapp.com"
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
		// type user float64
		var responseObject Response
		if claims, ok := tok.Claims.(jwt.MapClaims); ok && tok.Valid {
			user:=claims["userId"].(float64)
			if user!=0{
				fmt.Println("validate: user:", user)
				id:=fmt.Sprintf("%v",user)
				response,err:=http.Get(baseUrl+"/api/auth/users/"+id)
				if err!=nil{
					w.WriteHeader(http.StatusUnauthorized)
					fmt.Fprintf(w, "Unauthorized")
					return
				}
				responseData,err:=ioutil.ReadAll(response.Body)
				if err!=nil{
					w.WriteHeader(http.StatusUnauthorized)
					fmt.Fprintf(w, "Unauthorized")
					return
				}
				fmt.Println("validate: response:", string(responseData))
				json.Unmarshal(responseData,&responseObject)
			}
		}
		fmt.Println(responseObject.Data.Name,"usser")
		ctx:=context.WithValue(r.Context(), "user", responseObject.Data)
		fmt.Println(tok.Claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}