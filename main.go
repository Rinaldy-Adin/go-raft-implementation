package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"tubes.sister/raft/hello"
)

var (
	server     = flag.Bool("is_server", true, "GRPC Server or Client")
	port       = flag.Int("port", 50051, "GRPC Server Port")
	serverAddr = flag.String("server_address", "localhost:50051", "GRPC Server Address")
	clientPort = flag.Int("client_port", 3000, "HTTP Client Port")
)

type ClientRequest struct {
	Message string `json:"message"`
}

type ClientRes struct {
	Result string `json:"result"`
}

func main() {
	flag.Parse()

	if *server {
		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))

		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		var opts []grpc.ServerOption
		grpcServer := grpc.NewServer(opts...)
		hello.RegisterHelloServer(grpcServer, &hello.HelloServerImpl{})

		log.Printf("Server started at port %d", *port)
		grpcServer.Serve(lis)
	} else {
		var opts []grpc.DialOption
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

		conn, err := grpc.Dial(*serverAddr, opts...)
		if err != nil {
			log.Fatalf("Failed to dial server: %v", err)
		}

		defer conn.Close()

		client := hello.NewHelloClient(conn)

		router := chi.NewRouter()
		router.Use(middleware.Logger)
		router.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"POST"},
		}))

		router.Post("/", func(w http.ResponseWriter, r *http.Request) {
			var req ClientRequest

			err := json.NewDecoder(r.Body).Decode(&req)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}

			response, err := client.SayHello(context.Background(), &hello.HelloMsg{Name: req.Message})
			if err != nil {
				log.Printf("Failed to call SayHello: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			res := &ClientRes{Result: response.Message}
			w.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(w).Encode(res)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		})

		log.Printf("Client started at port %d", *clientPort)
		http.ListenAndServe(fmt.Sprintf(":%d", *clientPort), router)
	}
}