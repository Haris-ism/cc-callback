package controller_grpc

import (
	"cc-callback/constants"
	"cc-callback/protogen/merchant"
	"context"
	"fmt"
	"net/http"

	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (g *ControllerGrpc) InquiryItems(ctx context.Context, empt *emptypb.Empty) (*merchant.InquiryMerchantItemsModel, error){
	fmt.Println("masuk con")
	md, ok := metadata.FromIncomingContext(ctx)
	fmt.Println("ieu md:",md.Get("tes")[0])
	fmt.Println("ieu ok:",ok)
	res:=merchant.InquiryMerchantItemsModel{
		Message:constants.SUCCESS,
		Code:http.StatusOK,
	}
	result,err:=g.uc.InquiryItems()
	
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
	}
	res.Data=result
	return &res,nil
}

func (g *ControllerGrpc) InquiryDiscounts(context.Context, *emptypb.Empty) (*merchant.InquiryMerchantDiscountsModel, error){
	fmt.Println("masuk con")
	res:=merchant.InquiryMerchantDiscountsModel{
		Message:constants.SUCCESS,
		Code:http.StatusOK,
	}
	result,err:=g.uc.InquiryDiscounts()
	
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
	}
	res.Data=result
	return &res,nil
}