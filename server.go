package microservice_users_status

import (
	"github.com/Ascemme/microservice.users.status/pkg/handler"
	"github.com/Ascemme/microservice.users.status/protos/ascemme/grpcProto"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"time"
)

func Run(r *mux.Router) {
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s := grpc.NewServer()

	rpcSrv := &handler.GRPCServer{}
	grpcProto.RegisterStatusGrpcServer(s, rpcSrv)

	listen, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal(err)
		return
	}
	go s.Serve(listen)

	srv.ListenAndServe()
}
