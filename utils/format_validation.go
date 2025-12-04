package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func HandleValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationError {
			switch e.Tag() {
			case "gt":
				errors[e.Field()] = e.Field() + "phải lớn hơn giá trị tối thiểu"
			case "uuid":
				errors[e.Field()] = e.Field() + "phải là UUID hợp lệ"
			case "slug":
				errors[e.Field()] = e.Field() + "chỉ chứa chữ thường , số gạch đầu dòng "
			}

		}

		return gin.H{"error": errors}
	}
	return gin.H{"error": "Yêu cầu không hợp lệ"}
}
