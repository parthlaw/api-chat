package main

import (
	"api-chat/rabbit"
	"api-chat/routes"
	"api-chat/websocket"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)
func helloHandler(name string) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        responseText := "<h1>Hello " + name + "</h1>"

        if requestTime := r.Context().Value("requestTime"); requestTime != nil {
            if str, ok := requestTime.(string); ok {
                responseText = responseText + "\n<small>Generated at: " + str + "</small>"
            }
        }
        w.Write([]byte(responseText))
    })
}
func websocketServer (wsServer *websocket.WsServer) http.Handler {
    fmt.Println("Websocket server started")
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {websocket.ServeWs(wsServer,w,r)})
}
func main() {
    if os.Getenv("ENV") != "production" {
        er:=godotenv.Load(".env")
        if er!=nil {
            panic(er)
        } 
    }
    
    conn,err:=amqp.Dial(os.Getenv("RABBIT_URI"))
    if err!=nil {
        fmt.Println(err)
        panic(1)
    }
    defer conn.Close()
    fmt.Println("Connected to RabbitMQ instance")
    con:=rabbit.NewConn(conn)
    go con.HandleMessages()
    wsServer:=websocket.NewWebsocketServer(conn)
    go wsServer.Run()
	fmt.Println("Working ok")
	mux:=http.NewServeMux()
	mux.Handle("/",routes.Middleware(helloHandler("World")))
    mux.Handle("/ws",routes.Middleware(websocketServer(wsServer)))
    port:=os.Getenv("PORT")
    if port=="" {
        port="8080"
    }
	http.ListenAndServe(":"+port,mux)
}