package grpc_merchant

import (
	"cc-callback/protogen/merchant"
	"cc-callback/utils"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type(
	merchantGrpc struct{
		merchantConn merchant.InquiryServicesClient
		// callbackHost		string
		// inquiryItems		string
		// inquiryDiscounts	string
	}
	MerchantInterface interface{
		InquiryItems()(*merchant.InquiryMerchantItemsModel,error)
	}
)

func (g *merchantGrpc)InquiryItems()(*merchant.InquiryMerchantItemsModel,error){
	res,err:=g.merchantConn.InquiryItems(context.Background(),&emptypb.Empty{})
	if err != nil {
		log.Println("Error on grpc merchant :", err)
		return res,err
	}

	return res,nil
}

func InitGrpcMerchant()MerchantInterface{
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	fmt.Println("merchant server:",utils.GetEnv("MERCHANT_HOST_GRPC"))
	conn, err := grpc.Dial(utils.GetEnv("MERCHANT_HOST_GRPC"),opts...)
	if err!=nil{
		fmt.Println("failed to dial grpc merchant:",err)
	}
	
	merchantConn:=merchant.NewInquiryServicesClient(conn)
	fmt.Println("grpc merchant connected")
	return &merchantGrpc{
		merchantConn:merchantConn,
	}
}