package controller_grpc

import (
	"cc-callback/constants"
	"cc-callback/protogen/merchant"
	"context"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (g *ControllerGrpc) InquiryItems(context.Context, *emptypb.Empty) (*merchant.InquiryMerchantItemsModel, error){
	fmt.Println("masuk con")
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