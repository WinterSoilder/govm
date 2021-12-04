package db

import (
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

func WithDB(target func(http.ResponseWriter, *http.Request, *mongo.Client), mongoSession *mongo.Client) func(http.ResponseWriter, *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		target(res, req, mongoSession)
	}
}
