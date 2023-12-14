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
	apiGotCreateOrderRequestMsg   = "HttpApiService got CreateOrderRequest: %+v"
	apiGotGetUserOrdersRequestMsg = "HttpApiService got GetUserOrdersRequest: %+v"
)

type OrderEntrypoint struct {
	RabbitProvider *providers.RabbitProvider
	OrderPorcessor *processing.OrderProcessor
}

func (o *OrderEntrypoint) CreateOrder(c *gin.Context) {
	request, err := validators.GetValidatedCreateOrderRequest(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, proto.ErrorDto{Code: proto.ErrorCode_ERROR_INVALID_REQUEST, Message: err.Error()})
		return
	}

	logger.Infof(apiGotCreateOrderRequestMsg, request.String())

	createdOrder, responseErr := o.OrderPorcessor.GetCreateOrderResponse(request)
	if responseErr != nil {
		c.IndentedJSON(http.StatusBadRequest, responseErr)
	} else {
		c.IndentedJSON(http.StatusOK, createdOrder)
	}
}

func (o *OrderEntrypoint) GetUserOrders(c *gin.Context) {
	request, err := validators.GetValidatedGetUserOrdersRequest(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, proto.ErrorDto{Code: proto.ErrorCode_ERROR_INVALID_REQUEST, Message: err.Error()})
		return
	}

	logger.Infof(apiGotGetUserOrdersRequestMsg, &request)

	orders, responseErr := o.OrderPorcessor.GetGetUserOrdersResponse(request)
	if responseErr != nil {
		c.IndentedJSON(http.StatusBadRequest, responseErr)
	} else {
		c.IndentedJSON(http.StatusOK, orders)
	}
}

func (o *OrderEntrypoint) RemoveOrder(c *gin.Context) {
	request, err := validators.GetValidatedRemoveOrderRequest(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, proto.ErrorDto{Code: proto.ErrorCode_ERROR_INVALID_REQUEST, Message: err.Error()})
		return
	}

	logger.Infof(apiGotGetUserOrdersRequestMsg, &request)

	responseErr := o.OrderPorcessor.GetRemoveOrderResponse(request)
	if responseErr != nil {
		c.IndentedJSON(http.StatusBadRequest, responseErr)
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully deleted"})
	}
}
