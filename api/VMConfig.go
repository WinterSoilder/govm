package api

import (
	"context"
	"encoding/json"
	"fmt"
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
	defer cancel()

	var vm VM_Config
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}

	bearerToken := req.Header.Get("Authorization")
	var authorizationToken = strings.Split(bearerToken, " ")[1]

	email, _ := VerifyToken(authorizationToken)
	if email == "" {
		ReturnErrorResponse(res, req, errorResponse)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	vm.user_id = GetUserId(email, mongoSession)
	err = json.Unmarshal(body, &vm)
	if err != nil {
		panic(err)
	}

	coll := mongoSession.Database("govm").Collection("vm_config")
	data, insertErr := coll.InsertOne(ctx, vm)
	if insertErr != nil {
		panic(insertErr)
	} else if data != nil {
		fmt.Print(data)
	}
}

func GetVMConfigs(res http.ResponseWriter, req *http.Request, mongoSession *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var results []bson.M
	opts := options.Find()
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}

	bearerToken := req.Header.Get("Authorization")
	var authorizationToken = strings.Split(bearerToken, " ")[1]

	email, _ := VerifyToken(authorizationToken)
	if email == "" {
		ReturnErrorResponse(res, req, errorResponse)
	}

	filter := bson.D{{"user_id", GetUserId(email, mongoSession)}}
	coll := mongoSession.Database("govm").Collection("vm_config")
	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &results); err != nil {
		panic(err)
	}

	// var vm VM_Config
	jData, err := json.Marshal(results)
	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jData)
}

func GetVMConfig(res http.ResponseWriter, req *http.Request, mongoSession *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.FindOne()
	fmt.Print(mux.Vars(req)["id"])
	objectId, err := primitive.ObjectIDFromHex(mux.Vars(req)["id"])
	if err != nil {
		panic(err)
	}

	filter := bson.D{{"_id", objectId}}
	coll := mongoSession.Database("govm").Collection("vm_config")

	var result []bson.M
	coll.FindOne(ctx, filter, opts).Decode(&result)

	jData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jData)
}
