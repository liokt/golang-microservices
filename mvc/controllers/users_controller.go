package controllers

import (
	"encoding/json"
	"github.com/liomazza/golang-microservices/mvc/services"
	"github.com/liomazza/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request){

	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		apiError := &utils.ApplicationError{
			Message: "user id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}
		jsonValue,_ := json.Marshal(apiError)
		resp.WriteHeader(apiError.StatusCode)
		resp.Write(jsonValue)
		//Just return the bad request to the client
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue,_ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write([]byte(jsonValue))
		return
	}

	//return user to client
	jsonValue,_ := json.Marshal(user)
	resp.Write(jsonValue)

}
