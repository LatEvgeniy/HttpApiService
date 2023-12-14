package processing

import (
	"HttpApiService/converters"
	"HttpApiService/models"
	"HttpApiService/proto"
	"HttpApiService/providers"

	googleProto "google.golang.org/protobuf/proto"
)

var (
	orderProcessingServiceExName = "ex.OrderProcessingService"

	createOrderResponseListenerQueueName   = "q.HttpApiService.CreateOrderResponse.Listener"
	getUserOrdersResponseListenerQueueName = "q.HttpApiService.GetUserOrdersResponse.Listener"
	removeOrderResponseListenerQueueName   = "q.HttpApiService.RemoveOrderResponse.Listener"

	createOrderResponseRk     = "rk.CreateOrderResponse"
	getUserOrdersResponseRk   = "rk.GetUserOrdersResponse"
	removeOrderResponseRkName = "rk.RemoveOrderResponse"

	createOrderRequestRk       = "rk.CreateOrdeRequest"
	getUserOrdersRequestRkName = "rk.GetUserOrdersRequest"
	removeOrderRequestRkName   = "rk.RemoveOrderRequest"
)

type OrderProcessor struct {
	RabbitProvider *providers.RabbitProvider
}

func (o *OrderProcessor) GetCreateOrderResponse(request *proto.CreateOrderRequest) (*models.OrderModel, *proto.ErrorDto) {
	sendBody, err := googleProto.Marshal(request)
	if err != nil {
		return nil, &proto.ErrorDto{Code: proto.ErrorCode_ERROR_INTERNAL, Message: err.Error()}
	}

	var response proto.CreateOrderResponse
	if err := o.RabbitProvider.SendMessageAndHandleResponse(orderProcessingServiceExName, createOrderRequestRk,
		createOrderResponseRk, createOrderResponseListenerQueueName, sendBody, &response); err != nil {
		return nil, &proto.ErrorDto{Code: proto.ErrorCode_ERROR_INTERNAL, Message: err.Error()}
	}

	if err := response.GetError(); err != nil {
		return nil, err
	}

	return converters.ConvertProtoOrderToOrderModel(response.CreatedOrder), nil
}

func (o *OrderProcessor) GetGetUserOrdersResponse(request *proto.GetUserOrdersRequest) ([]models.OrderModel, *proto.ErrorDto) {
	sendBody, err := googleProto.Marshal(request)
	if err != nil {
		return nil, &proto.ErrorDto{Code: proto.ErrorCode_ERROR_INTERNAL, Message: err.Error()}
	}

	var response proto.GetUserOrdersResponse
	if err := o.RabbitProvider.SendMessageAndHandleResponse(orderProcessingServiceExName, getUserOrdersRequestRkName,
		getUserOrdersResponseRk, getUserOrdersResponseListenerQueueName, sendBody, &response); err != nil {
		return nil, &proto.ErrorDto{Code: proto.ErrorCode_ERROR_INTERNAL, Message: err.Error()}
	}

	if err := response.GetError(); err != nil {
		return nil, err
	}
	var orderModels []models.OrderModel
	for _, order := range response.Orders {
		orderModel := converters.ConvertProtoOrderToOrderModel(order)
		orderModels = append(orderModels, *orderModel)
	}

	return orderModels, nil
}

func (o *OrderProcessor) GetRemoveOrderResponse(request *proto.RemoveOrderRequest) *proto.ErrorDto {
	sendBody, err := googleProto.Marshal(request)
	if err != nil {
		return &proto.ErrorDto{Code: proto.ErrorCode_ERROR_INTERNAL, Message: err.Error()}
	}

	var response proto.RemoveOrderResponse
	if err := o.RabbitProvider.SendMessageAndHandleResponse(orderProcessingServiceExName, removeOrderRequestRkName,
		removeOrderResponseRkName, removeOrderResponseListenerQueueName, sendBody, &response); err != nil {
		return &proto.ErrorDto{Code: proto.ErrorCode_ERROR_INTERNAL, Message: err.Error()}
	}

	if err := response.GetError(); err != nil {
		return err
	}

	return nil
}
