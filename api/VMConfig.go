package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VM_Config struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	VM_name  string             `bson:"VM_name,omitempty"`
	cpus     int                `bson:"cpus,omitempty"`
	Disk     int                `bson:"Disk,omitempty"`
	Memory   int                `bson:"Memory,omitempty"`
	Template string             `bson:"Template,omitempty"`
}

type User_Config struct {
	name string `name:"VM_name,omitempty"`
}

func CreateVMConfig(res http.ResponseWriter, req *http.Request, mongoSession *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	var vm VM_Config
	err = json.Unmarshal(body, &vm)
	if err != nil {
		panic(err)
	}
	log.Println(vm)

	// return vm.Template
	coll := mongoSession.Database("govm").Collection("vm_config")
	data, insertErr := coll.InsertOne(ctx, vm)
	if insertErr != nil {
		panic(insertErr)
	} else if data != nil {
		fmt.Print(data)
	}

	// time.Sleep(2)
	// go process.LaunchVM()
}

func GetVMConfigs(res http.ResponseWriter, req *http.Request, mongoSession *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opts := options.Find()
	filter := bson.D{{"user", req.URL.Query().Get("user")}}
	coll := mongoSession.Database("govm").Collection("vm_config")
	cursor, err := coll.Find(ctx, filter, opts)
	if err != nil {
		panic(err)
	}
	var results []bson.M
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

	var result []bson.D
	coll.FindOne(ctx, filter, opts).Decode(&result)

	jData, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(jData)
}
