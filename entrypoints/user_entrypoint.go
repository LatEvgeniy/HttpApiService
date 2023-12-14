package entrypoints

import (
	"HttpApiService/processing"
	"HttpApiService/proto"
	"HttpApiService/providers"
	"HttpApiService/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
)

var (
	apiGotCreateUserRequestMsg  = "HttpApiService got CreateUserRequest with name: %s"
	apiGotAddMoneyRequestMsg    = "HttpApiService got AddMoneyRequest: %s"
	apiGotGetUserByIdRequestMsg = "HttpApiService got GetUserByIdRequest with id: %s"
)

type UserEntrypoint struct {
	RabbitProvider *providers.RabbitProvider
	UserProcessor  *processing.UserProcessor
}

func (u *UserEntrypoint) CreateUser(c *gin.Context) {
	request, err := validators.GetValidatedCreateUserRequest(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, proto.ErrorDto{Code: proto.ErrorCode_ERROR_INVALID_REQUEST, Message: err.Error()})
		return
	}

	logger.Infof(apiGotCreateUserRequestMsg, request.Name)

	userModel, responseErr := u.UserProcessor.GetCreateUserResponse(request.Name)
	if responseErr != nil {
		c.IndentedJSON(http.StatusBadRequest, responseErr)
	} else {
		c.IndentedJSON(http.StatusOK, userModel)
	}
}

func (u *UserEntrypoint) AddMoney(c *gin.Context) {
	var request proto.AddMoneyRequest
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, proto.ErrorDto{Code: proto.ErrorCode_ERROR_INVALID_REQUEST, Message: err.Error()})
		return
	}

	logger.Infof(apiGotAddMoneyRequestMsg, request.String())

	userModel, responseErr := u.UserProcessor.GetAddMoneyResponse(&request)
	if responseErr != nil {
		c.IndentedJSON(http.StatusBadRequest, responseErr)
	} else {
		c.IndentedJSON(http.StatusOK, userModel)
	}
}

func (u *UserEntrypoint) GetUserById(c *gin.Context) {
	request, err := validators.GetValidatedGetUserByIdRequest(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, proto.ErrorDto{Code: proto.ErrorCode_ERROR_INVALID_REQUEST, Message: err.Error()})
		return
	}

	logger.Infof(apiGotGetUserByIdRequestMsg, request.Id)

	userModel, responseErr := u.UserProcessor.GetUserByIdResponse(request.Id)
	if responseErr != nil {
		c.IndentedJSON(http.StatusBadRequest, responseErr)
	} else {
		c.IndentedJSON(http.StatusOK, userModel)
	}
}
