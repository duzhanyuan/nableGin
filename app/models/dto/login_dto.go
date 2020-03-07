package dto

import (
	"github.com/gin-gonic/gin"
	"nable.gin/libraries/validator"
	"net/http"
)

type Login_validate struct {
	Username       string    `form:"username" json:"用户名" validate:"required,min=5,max=12"`
	Password       string    `form:"password" json:"密码" validate:"required,min=6"`
	Capt_request   string    `form:"captcha" json:"验证码" validate:"required"`
}

func (r *Login_validate) Bind(ctx *gin.Context) string {

	////读取表单数据绑定到结构体
	if err := ctx.ShouldBind(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	//表单验证
	if err, errMsg := validator.Check(r); err != nil {
		return errMsg
	}

	return ""
}
