package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ReturnErrorResponse(response http.ResponseWriter, request *http.Request, errorResponse ErrorResponse) {
	response.Header().Set("Content-Type", "application/json")
	errorJSONResponse, _ := json.Marshal(errorResponse)
	response.Write(errorJSONResponse)
}

func GetUserId(email string, mongoSession *mongo.Client) string {
	var user UserDetails
	mongoSession.Database("govm").Collection("users").FindOne(context.TODO(), bson.D{{"Email", email}}).Decode(&user)
	return user.ID.String()
}

// CORSRouterDecorator applies CORS headers to a mux.Router
type CORSRouterDecorator struct {
	R *mux.Router
}

// ServeHTTP wraps the HTTP server enabling CORS headers.
// For more info about CORS, visit https://www.w3.org/TR/cors/
func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		rw.WriteHeader(http.StatusOK)
	} else {
		c.R.ServeHTTP(rw, req)
	}

}
