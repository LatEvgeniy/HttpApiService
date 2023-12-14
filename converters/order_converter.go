package converters

import (
	"HttpApiService/models"
	"HttpApiService/proto"
	"time"
)

func ConvertProtoOrderToOrderModel(protoOrder *proto.Order) *models.OrderModel {
	return &models.OrderModel{
		UserId:         protoOrder.UserId,
		OrderId:        protoOrder.OrderId,
		Pair:           protoOrder.Pair.String(),
		Direction:      protoOrder.Direction.String(),
		Type:           protoOrder.Type.String(),
		InitVolume:     protoOrder.InitVolume,
		InitPriice:     protoOrder.InitPrice,
		FilledVolume:   protoOrder.FilledVolume,
		FilledPrice:    protoOrder.FilledPrice,
		LockedVolume:   protoOrder.LockedVolume,
		ExpirationDate: time.Unix(protoOrder.ExpirationDate, 0).String(),
		CreationDate:   time.Unix(protoOrder.CreationDate, 0).String(),
		UpdatedDate:    time.Unix(protoOrder.UpdatedDate, 0).String(),
	}
}
