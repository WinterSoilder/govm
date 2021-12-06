package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// SignInUser Used for Signing In the Users
func SignInUser(response http.ResponseWriter, request *http.Request, mongoSession *mongo.Client) {
	var loginRequest LoginParams
	var result UserDetails
	var errorResponse = ErrorResponse{
		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
	}

	fmt.Print("im in")

	decoder := json.NewDecoder(request.Body)
	decoderErr := decoder.Decode(&loginRequest)
	defer request.Body.Close()

	if decoderErr != nil {
		ReturnErrorResponse(response, request, errorResponse)
	} else {
		errorResponse.Code = http.StatusBadRequest
		if loginRequest.Email == "" {
			errorResponse.Message = "Last Name can't be empty"
			ReturnErrorResponse(response, request, errorResponse)
		} else if loginRequest.Password == "" {
			errorResponse.Message = "Password can't be empty"
			ReturnErrorResponse(response, request, errorResponse)
		} else {

			collection := mongoSession.Database("govm").Collection("users")

			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			var err = collection.FindOne(ctx, bson.M{
				"email":    loginRequest.Email,
				"password": loginRequest.Password,
			}).Decode(&result)

			defer cancel()

			if err != nil {
				ReturnErrorResponse(response, request, errorResponse)
			} else {
				tokenString, _ := CreateJWT(loginRequest.Email)

				if tokenString == "" {
					ReturnErrorResponse(response, request, errorResponse)
				}

				var successResponse = SuccessResponse{
					Code:    http.StatusOK,
					Message: "You are registered, login again",
					Response: SuccessfulLoginResponse{
						AuthToken: tokenString,
						Email:     loginRequest.Email,
					},
				}

				successJSONResponse, jsonError := json.Marshal(successResponse)

				if jsonError != nil {
					ReturnErrorResponse(response, request, errorResponse)
				}
				response.Header().Set("Content-Type", "application/json")
				response.Write(successJSONResponse)
			}
		}
	}
}
