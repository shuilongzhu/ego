package ego

import "gorm.io/gorm"

type Field struct {
	TableName string        `json:"table_name"` //表名称
	Pluck     string        `json:"pluck"`      //要查询的列，*：查询全部字段
	Joins     string        `json:"joins"`      //关联的表和关联条件
	Group     string        `json:"group"`      //分组
	Sort      string        `json:"sort"`       //排序,例:id desc
	Offset    int           `json:"offset"`     //分页从第几个开始查
	Limit     int           `json:"limit"`      //分页一页查几条
	Where     string        `json:"where"`      //where条件 user_name = ? and password = ?
	Args      []interface{} `json:"args"`       //where字段 张三 1111
}

type Gorm struct {
	Db *gorm.DB
}
