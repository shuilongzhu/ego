package ego

import (
	"gorm.io/gorm"
	"net/http"
)

type Field struct {
	TableName string `json:"table_name"` //表名称
	Pluck     string `json:"pluck"`      //要查询的列，*：查询全部字段
	Joins     string `json:"joins"`      //关联的表和关联条件
	Group     string `json:"group"`      //分组
	Sort      string `json:"sort"`       //排序,例:id desc
	Offset    int    `json:"offset"`     //分页从第几个开始查
	Limit     int    `json:"limit"`      //分页一页查几条
	Where     string `json:"where"`      //where条件 user_name = ? and password = ?
}

type Gorm struct {
	Db       *gorm.DB `json:"db"`
	Ip       string   `json:"ip"`
	DataBase string   `json:"dataBase"`
}

// DfsHTTPClient
// @Description: seaweedfs分布式存储相关的配置信息结构体
type DfsHTTPClient struct {
	DfsMasterAddress string       `json:"dfs_master_address"`
	DfsFilerAddress  string       `json:"dfs_filer_address"`
	Client           *http.Client `json:"client"`
}

type ZipFile struct {
	FilePath string `json:"file_path"`
	FileName string `json:"file_name"`
}
