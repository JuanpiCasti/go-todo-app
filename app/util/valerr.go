package util

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// TODO: Make error messages *
// TODO: Make input error struct instead of map[string]string *
// TODO: implement own ShouldBindJSON to collect all unmarshalling errors ***

func BuildErrorResponse(err error, bodyType reflect.Type) []map[string]string {
	var bindingErrors validator.ValidationErrors
	var jsonErr *json.UnmarshalTypeError
	if errors.As(err, &bindingErrors) {
		var errorMessages []map[string]string
		for _, valerr := range bindingErrors {
			fieldName := valerr.Field()
			field, _ := bodyType.Elem().FieldByName(fieldName)
			fieldJSONName, _ := field.Tag.Lookup("json")
			errorMessages = append(errorMessages,
				map[string]string{
					"field": fieldJSONName,
					"msg":   valerr.ActualTag(),
				})
		}
		return errorMessages
	} else if errors.As(err, &jsonErr) {
		return []map[string]string{{
			"field": jsonErr.Field,
			"msg":   jsonErr.Type.String(),
		},
		}
	}

	return []map[string]string{{
		"error": err.Error(),
		"msg":   "JSON unmarshalling error",
	},
	}
}

func BindJsonWithErrs(c *gin.Context, body interface{}) []map[string]string {
	if err := c.ShouldBindJSON(body); err != nil {
		return BuildErrorResponse(err, reflect.TypeOf(body))
	}
	return nil
}
