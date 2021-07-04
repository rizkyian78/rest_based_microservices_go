package service

import (
	"rest_based_micro_go/src/domain"
	"rest_based_micro_go/src/utils"
)

func GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(userId)
}
