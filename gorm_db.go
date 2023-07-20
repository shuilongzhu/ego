package ego

import (
	"errors"
)

type DbMethod interface {
	CommonCreate(tableName string, object interface{}) error
	CommonDelete(field Field, args ...interface{}) error
	CommonSelect(result interface{}, field Field, args ...interface{}) error
	CommonUpdateDb(field Field, mapInfo map[string]interface{}, args ...interface{}) error
}

var (
	GormMethodParamErr = errors.New("gorm method parameter error")
)

// CommonCreate @description: gorm公共创建方法
// @receiver gorm
// @parameter tableName
// @parameter object
// @return error
func (gorm *Gorm) CommonCreate(tableName string, object interface{}) error {

	if "" == tableName {
		return GormMethodParamErr
	}

	return gorm.Db.Table(tableName).Create(object).Error
}

// CommonDelete @description: gorm公共删除方法
// @receiver gorm
// @parameter field
// @parameter args
// @return error
func (gorm *Gorm) CommonDelete(field Field, args ...interface{}) error {

	if !gormParamCheck(field) {
		return GormMethodParamErr
	}

	operate := gorm.Db.Table(field.TableName)
	if "" != field.Where {
		operate = operate.Where(field.Where, args...)
	}

	operate = operate.Delete(nil)

	return operate.Error
}

// CommonSelect @description: gorm公共查询方法
// @receiver gorm
// @parameter result
// @parameter field
// @return error
func (gorm *Gorm) CommonSelect(result interface{}, field Field, args ...interface{}) error {

	if !gormParamCheck(field) {
		return GormMethodParamErr
	}

	if "" == field.Pluck {
		field.Pluck = "*"
	}
	operate := gorm.Db.Table(field.TableName).Select(field.Pluck)
	if "" != field.Joins {
		operate = operate.Joins(field.Joins)
	}

	//条件拼接
	if "" != field.Where {
		operate = operate.Where(field.Where, args...)
	}

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

// CommonUpdateDb @description: gorm公共更新方法
// @receiver gorm
// @parameter field
// @parameter mapInfo
// @parameter args
// @return error
func (gorm *Gorm) CommonUpdateDb(field Field, mapInfo map[string]interface{}, args ...interface{}) error {

	if !gormParamCheck(field) {
		return GormMethodParamErr
	}

	operate := gorm.Db.Table(field.TableName)
	if "" != field.Where {
		operate = operate.Where(field.Where, args...)
	}

	operate = operate.Updates(mapInfo)
	return operate.Error
}

// gormParamCheck @description: gorm方法入参校验
// @parameter field
// @return bool(是否合规) true:合规；false:不合规
func gormParamCheck(field Field) bool {
	if "" == field.TableName {
		return false
	}
	return true
}
