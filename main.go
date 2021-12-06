package main

import (
	"context"
	"example/user/govm/api"
	"example/user/govm/db"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://shashankmadan:railway999@cluster0.gjkve.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(nil)
	}

	r := mux.NewRouter()

	r.PathPrefix("").Handler(http.FileServer(http.Dir("./web/build/")))

	r.HandleFunc("/userlogin", db.WithDB(api.SignInUser, client)).Methods(http.MethodPost, http.MethodOptions)

	r.HandleFunc("/vm_config", db.WithDB(api.CreateVMConfig, client)).Methods(http.MethodPost, http.MethodOptions)
	r.HandleFunc("/vm_config", db.WithDB(api.GetVMConfigs, client)).Methods(http.MethodGet)
	r.HandleFunc("/vm_config/{id}", db.WithDB(api.GetVMConfig, client)).Methods(http.MethodGet)

	// http.HandleFunc("/vm_config")

	log.Fatal(http.ListenAndServe(":8030", r))
}
