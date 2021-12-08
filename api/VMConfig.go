package api

import (
	"context"
	"encoding/json"
	"example/user/govm/process"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User_Config struct {
	name string `name:"VM_name,omitempty"`
}

func CreateVMConfig(res http.ResponseWriter, req *http.Request, mongoSession *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Server Error",
	}
	var vm VM_Config

	defer cancel()

	bearerToken := req.Header.Get("Authorization")
	var authorizationToken = strings.Split(bearerToken, " ")[1]

	email, _ := VerifyToken(authorizationToken)
	if email == "" {
		ReturnErrorResponse(res, req, errorResponse)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		ReturnErrorResponse(res, req, errorResponse)
	}

	err = json.Unmarshal(body, &vm)
	if err != nil {
		ReturnErrorResponse(res, req, errorResponse)
	}

	coll := mongoSession.Database("govm").Collection("vm_config")
	vm.User_Id = GetUserId(email, mongoSession)
	data, insertErr := coll.InsertOne(ctx, vm)
	if insertErr != nil {
		errorResponse.Code = http.StatusNoContent
		errorResponse.Message = "Data Not Inserted"
		ReturnErrorResponse(res, req, errorResponse)
	} else if data != nil {
		var result map[string]interface{}
		coll.FindOne(ctx, bson.D{{"_id", data.InsertedID}}, options.FindOne()).Decode(&result)
		process.LaunchVM(result)
	}
}

func GetVMConfigs(res http.ResponseWriter, req *http.Request, mongoSession *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Server Error",
	}
	var results []bson.M

	defer cancel()

	opts := options.Find()

	bearerToken := req.Header.Get("Authorization")
	var authorizationToken = strings.Split(bearerToken, " ")[1]
	email, _ := VerifyToken(authorizationToken)
	if email == "" {
		ReturnErrorResponse(res, req, errorResponse)
	}

	filter := bson.D{{"User_Id", GetUserId(email, mongoSession)}}
	coll := mongoSession.Database("govm").Collection("vm_config")
	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		errorResponse.Code = http.StatusNotFound
		errorResponse.Message = "Data Invalid Or Not Found"
		ReturnErrorResponse(res, req, errorResponse)
	}
	if err = cursor.All(ctx, &results); err != nil {
		errorResponse.Code = http.StatusNotFound
		errorResponse.Message = "Data Invalid Or Not Found"
		ReturnErrorResponse(res, req, errorResponse)
	}

	jData, err := json.Marshal(results)
	if err != nil {
		ReturnErrorResponse(res, req, errorResponse)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jData)
}

func GetVMConfig(res http.ResponseWriter, req *http.Request, mongoSession *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "Server Error",
	}
	var result []bson.M

	defer cancel()

	objectId, err := primitive.ObjectIDFromHex(mux.Vars(req)["id"])
	if err != nil {
		ReturnErrorResponse(res, req, errorResponse)
	}

	filter := bson.D{{"_id", objectId}}
	coll := mongoSession.Database("govm").Collection("vm_config")

	coll.FindOne(ctx, filter, options.FindOne()).Decode(&result)

	jData, err := json.Marshal(result)
	if err != nil {
		ReturnErrorResponse(res, req, errorResponse)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jData)
}
