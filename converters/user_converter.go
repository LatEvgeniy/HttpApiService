package converters

import (
	"HttpApiService/models"
	"HttpApiService/proto"
)

func ConvertProtoUserToUserModel(protoUser *proto.User) *models.UserModel {
	userModel := &models.UserModel{
		Id:   protoUser.Id,
		Name: protoUser.Name,
	}

	for _, protoBalance := range protoUser.Balances {
		balanceModel := &models.BalanceModel{
			CurrencyName:  protoBalance.CurrencyName,
			Balance:       protoBalance.Balance,
			LockedBalance: protoBalance.LockedBalance,
		}
		userModel.Balances = append(userModel.Balances, balanceModel)
	}

	return userModel
}
