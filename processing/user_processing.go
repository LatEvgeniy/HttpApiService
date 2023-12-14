package processing

import (
	"HttpApiService/converters"
	"HttpApiService/models"
	"HttpApiService/proto"
	"HttpApiService/providers"

	googleProto "google.golang.org/protobuf/proto"
)

var (
	balanceServiceExchangeName = "ex.BalanceService"

	addMoneyResponseListenerQueueName    = "q.HttpApiService.AddMoneyResponse.Listener"
	createUserReponseListenerQueueName   = "q.HttpApiService.CreateUserResponse.Listener"
	getUserByIdResponseListenerQueueName = "q.HttpApiService.GetUserByIdResponse.Listener"

	addMoneyResponseRk    = "rk.AddMoneyResponse"
	createUserResponseRk  = "rk.CreateUserResponse"
	getUserByIdResponseRk = "rk.GetUserByIdResponse"

	addMoneyRequestRk   = "rk.AddMoneyRequest"
	createUserRequestRk = "rk.CreateUserRequest"
	getUserRequestRk    = "rk.GetUserByIdRequest"
)

type UserProcessor struct {
	RabbitProvider *providers.RabbitProvider
}

func (u *UserProcessor) GetCreateUserResponse(request string) (*models.UserModel, *proto.ErrorDto) {
	var response proto.CreateUserResponse
	if err := u.RabbitProvider.SendMessageAndHandleResponse(balanceServiceExchangeName, createUserRequestRk,
		createUserResponseRk, createUserReponseListenerQueueName, []byte(request), &response); err != nil {
		return nil, &proto.ErrorDto{Code: proto.ErrorCode_ERROR_INTERNAL, Message: err.Error()}
	}

	if err := response.GetError(); err != nil {
		return nil, err
	}

	return converters.ConvertProtoUserToUserModel(response.CreatedUser), nil
}

func (u *UserProcessor) GetAddMoneyResponse(request *proto.AddMoneyRequest) (*models.UserModel, *proto.ErrorDto) {
	sendBody, err := googleProto.Marshal(request)
	if err != nil {
		return nil, &proto.ErrorDto{Code: proto.ErrorCode_ERROR_INTERNAL, Message: err.Error()}
	}

	var response proto.AddMoneyResponse
	if err := u.RabbitProvider.SendMessageAndHandleResponse(balanceServiceExchangeName, addMoneyRequestRk,
		addMoneyResponseRk, addMoneyResponseListenerQueueName, sendBody, &response); err != nil {
		return nil, &proto.ErrorDto{Code: proto.ErrorCode_ERROR_INTERNAL, Message: err.Error()}
	}

	if err := response.GetError(); err != nil {
		return nil, err
	}

	return converters.ConvertProtoUserToUserModel(response.User), nil
}

func (u *UserProcessor) GetUserByIdResponse(request string) (*models.UserModel, *proto.ErrorDto) {
	var response proto.GetUserByIdResponse
	if err := u.RabbitProvider.SendMessageAndHandleResponse(balanceServiceExchangeName, getUserRequestRk,
		getUserByIdResponseRk, getUserByIdResponseListenerQueueName, []byte(request), &response); err != nil {
		return nil, &proto.ErrorDto{Code: proto.ErrorCode_ERROR_INTERNAL, Message: err.Error()}
	}

	if err := response.GetError(); err != nil {
		return nil, err
	}

	return converters.ConvertProtoUserToUserModel(response.User), nil
}
