package controller

import (
	"cc-callback/constants"
	"cc-callback/controllers/models"
	mModels "cc-callback/hosts/merchant/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *controller)TransItem(ctx *gin.Context){
	res:=models.GeneralResponse{
		Message: constants.SUCCESS,
		Code:http.StatusOK,
	}
	req:=mModels.TransactionItems{}
	if err:=ctx.BindJSON(&req);err!=nil{
		res.Message=constants.INVALID_INPUT
		res.Code=http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest,res)
		return
	}
	data,err:=c.usecase.TransItem(req)
	if err!=nil{
		res.Message=err.Error()
		res.Code=http.StatusInternalServerError
		ctx.JSON(http.StatusInternalServerError,res)
		return
	}
	
	res.Data=data
	ctx.JSON(res.Code,res)
}