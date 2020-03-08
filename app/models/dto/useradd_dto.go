package dto

import (
	"github.com/gin-gonic/gin"
	"nable.gin/libraries/validator"
	"nable.gin/libraries/helper"
	"net/http"
)

type Useradd_validate struct {
	Truename       string    `form:"truename" json:"姓名" validate:"required"`
	Username       string    `form:"username" json:"用户名" validate:"required,min=5,max=12"`
	Password       string    `form:"password" json:"密码" validate:"required,min=6"`
	Email          string    `form:"email" json:"邮箱" validate:"required,email"`
	Mobile         string    `form:"mobile" json:"手机号码" validate:"required,min=11"`
}


func (r *Useradd_validate) Bind(ctx *gin.Context) string {

	//读取表单数据绑定到结构体
	if err := ctx.ShouldBind(&r); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	r.Username = helper.Strim(r.Username)//去掉左右两边空格
	r.Email = helper.Strim(r.Email)//去掉左右两边空格

	//表单验证
	if err, errMsg := validator.Check(r); err != nil {
		return errMsg
	}


	return ""
}
