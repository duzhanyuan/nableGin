package dto

import (
	"github.com/gin-gonic/gin"
	"nable.gin/libraries/validator"
	"net/http"
)

type Roleadd_validate struct {
	Name       string    `form:"name" json:"角色名称" validate:"required"`
}


func (r *Roleadd_validate) Bind(ctx *gin.Context) string {

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
