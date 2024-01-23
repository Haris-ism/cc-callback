package usecase_grpc

import (
	"cc-callback/protogen/merchant"
	"fmt"

	postgre "cc-callback/databases/postgresql"
	redis_db "cc-callback/databases/redis"
	grpc_client "cc-callback/grpc/client"
)

type (
	usecaseGrpc struct {
		postgre postgre.PostgreInterface
		redis   redis_db.RedisInterface
		host	grpc_client.GrpcInterface
	}
	UsecaseGrpcInterface interface {
		InquiryItems()([]*merchant.InquiryItemsModel, error)
		InquiryDiscounts()([]*merchant.InquiryDiscountsModel, error)
	}
)

func InitUsecaseGrpc(postgre postgre.PostgreInterface, redis redis_db.RedisInterface, host grpc_client.GrpcInterface) UsecaseGrpcInterface {
	fmt.Println("init uc grpc")
	return &usecaseGrpc{
		postgre: postgre,
		redis:   redis,
		host: host,
	}
}
