package dto

import (
	"github.com/gin-gonic/gin"
	"nable.gin/libraries/validator"
	"net/http"
)

type Useredit_validate struct {
	Truename       string    `form:"truename" json:"姓名" validate:"required"`
	Password       string    `form:"password" json:"密码" validate:"required,min=6"`
	Mobile         string    `form:"mobile" json:"手机号码" validate:"required,min=11"`
}


func (r *Useredit_validate) Bind(ctx *gin.Context) string {

	//读取表单数据绑定到结构体
	if err := ctx.ShouldBind(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//表单验证
	if err, errMsg := validator.Check(r); err != nil {
		return errMsg
	}


	return ""
}
