package validators

import (
	"HttpApiService/proto"
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetValidatedCreateUserRequest(c *gin.Context) (*proto.User, error) {
	var newUser proto.User
	if err := c.BindJSON(&newUser); err != nil {
		return nil, err
	}

	newUser.Name = strings.TrimSpace(newUser.Name)
	if newUser.Name == "" {
		return nil, errors.New("name cannot be empty string")
	}

	return &newUser, nil
}

func GetValidatedGetUserByIdRequest(c *gin.Context) (*proto.User, error) {
	var user proto.User
	if err := c.BindJSON(&user); err != nil {
		return nil, err
	}

	_, err := uuid.Parse(user.Id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetValidatedCreateOrderRequest(c *gin.Context) (*proto.CreateOrderRequest, error) {
	var createOrderRequest proto.CreateOrderRequest
	if err := c.BindJSON(&createOrderRequest); err != nil {
		return nil, err
	}

	_, err := uuid.Parse(createOrderRequest.UserId)
	if err != nil {
		return nil, err
	}

	return &createOrderRequest, nil
}

func GetValidatedGetUserOrdersRequest(c *gin.Context) (*proto.GetUserOrdersRequest, error) {
	var getUserOrdersRequest proto.GetUserOrdersRequest
	if err := c.BindJSON(&getUserOrdersRequest); err != nil {
		return nil, err
	}

	_, err := uuid.Parse(getUserOrdersRequest.UserId)
	if err != nil {
		return nil, err
	}

	return &getUserOrdersRequest, nil
}

func GetValidatedRemoveOrderRequest(c *gin.Context) (*proto.RemoveOrderRequest, error) {
	var removeOrderRequest proto.RemoveOrderRequest
	if err := c.BindJSON(&removeOrderRequest); err != nil {
		return nil, err
	}

	_, err := uuid.Parse(removeOrderRequest.UserId)
	if err != nil {
		return nil, err
	}

	_, err = uuid.Parse(removeOrderRequest.OrderId)
	if err != nil {
		return nil, err
	}

	return &removeOrderRequest, nil
}
