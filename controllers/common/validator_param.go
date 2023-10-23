package common

import (
	valid "gin_demo/pkg/validator"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
)

// Compare this snippet from controllers/common/validator_param.go:
// package common

func Parameter(c *gin.Context, p any) (err error) {
	if err = c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			// 非validator.ValidationErrors类型错误直接返回
			ResponseError(c, CodeInvalidParams)
			return
		}
		// validator.ValidationErrors类型错误则进行翻译
		ResponseErrorWithMsg(c, CodeInvalidParams, valid.RemoveTopStruct(errs.Translate(valid.Trans)))
		return
	}
	return
}
