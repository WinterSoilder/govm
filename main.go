package main

import (
	"context"
	"example/user/govm/api"
	"example/user/govm/db"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load("project.env")
	DB := os.Getenv("DB")

	clientOptions := options.Client().ApplyURI(DB)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(nil)
	}

	r := mux.NewRouter()

	r.HandleFunc("/userlogin", db.WithDB(api.SignInUser, client)).Methods(http.MethodPost)

	r.HandleFunc("/vm_config", db.WithDB(api.CreateVMConfig, client)).Methods(http.MethodPost)
	r.HandleFunc("/vm_config", db.WithDB(api.GetVMConfigs, client)).Methods(http.MethodGet)
	r.HandleFunc("/vm_config/{id}", db.WithDB(api.GetVMConfig, client)).Methods(http.MethodGet)

	r.PathPrefix("/admin").Handler(http.FileServer(http.Dir("./web/build/")))

	log.Fatal(http.ListenAndServe(":8030", &api.CORSRouterDecorator{r}))
}
