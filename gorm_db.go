package ego

import (
	"errors"
)

type DbMethod interface {
	CommonSelect(result interface{}, field Field, args ...interface{}) error
	CommonDelete(tableName string, args ...interface{}) error
}

var (
	MethodParamErr = errors.New("函数方法入参格式错误")
)

// CommonSelect @description: 公共gorm查询方法
// @receiver gorm
// @parameter result
// @parameter field
// @return error
func (gorm *Gorm) CommonSelect(result interface{}, field Field, args ...interface{}) error {

	if "" == field.TableName {
		return MethodParamErr
	}

	if "" == field.Pluck {
		field.Pluck = "*"
	}
	operate := gorm.Db.Table(field.TableName).Select(field.Pluck)
	if "" != field.Joins {
		operate = operate.Joins(field.Joins)
	}

	//条件拼接
	operate = operate.Where(field.Where, args...)

	if "" != field.Group {
		operate = operate.Group(field.Group)
	}
	if "" != field.Sort {
		operate = operate.Order(field.Sort)
	}
	if 0 != field.Offset {
		operate = operate.Offset(field.Offset)
	}
	if 0 != field.Limit {
		operate = operate.Limit(field.Limit)
	}

	operate = operate.Scan(result)
	return operate.Error
}

func (gorm *Gorm) CommonDelete(tableName string, args ...interface{}) error {
	return nil
}
