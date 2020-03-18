package dto

import (
	"github.com/gin-gonic/gin"
	"nable.gin/libraries/validator"
	"net/http"
)

type Nodeadd_validate struct {
	Name       string    `form:"name" json:"节点名称" validate:"required"`
	IsType     uint8      `form:"istype" json:"权限类型" validate:"gte=0,lte=1"`
	Reorder    string      `form:"reorder" json:"排序" validate:"numeric,gte=0"`
}


func (r *Nodeadd_validate) Bind(ctx *gin.Context) string {

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
