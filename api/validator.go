package api

import (
	"github.com/IvanRoussev/autocare/util"
	"github.com/go-playground/validator/v10"
)

var validMaintenanceType validator.Func = func(fieldLevel validator.FieldLevel) bool {
	maintenanceType, ok := fieldLevel.Field().Interface().(string)
	if ok {
		isValid := util.IsSupportedMaintenanceType(maintenanceType)
		return isValid
	}
	return false
}
