package usecase

import (
	"cc-callback/constants"
	cModels "cc-callback/controllers/models"
	"cc-callback/hosts/merchant/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (uc *usecase)TransItem(req models.TransactionItems, headers cModels.ReqHeader)(models.ResponseItems,error){
	result:=models.ResponseTransactionItems{}
	header := make(http.Header)
	header.Add("Accept", "*/*")
	header.Add("Content-Type", "application/json")
	header.Add("TimeStamp", headers.TimeStamp)
	header.Add("Signature", headers.Signature)
	res,bytes,err:=uc.host.Merchant().Send(constants.TRANSACTION_ITEMS,req,header)
	if err!=nil{
		return result.Data, errors.New(constants.ERROR_DB)
	}
	if res.StatusCode!=200{
		fmt.Println("code:",res.StatusCode)
		return result.Data, errors.New(constants.ERROR_INQUIRY)
	}
	err=json.Unmarshal(bytes,&result)
	if err!=nil{
		fmt.Println("err:",err)
		return result.Data, errors.New(constants.ERROR_INQUIRY)
	}
	return result.Data,nil
}