package ego

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

//ArrayIntersection 取两个string数组的交集(a∩b）
func ArrayIntersection(list1, list2 []string) []string {
	result := make([]string, 0)
	mapTemp := make(map[string]bool)
	for _, str := range list1 {
		mapTemp[str] = true
	}
	for _, str := range list2 {
		if _, ok := mapTemp[str]; ok {
			result = append(result, str)
		}
	}
	return result
}

// ArrayNotExistAnother @description: 取出a数组中不在b数组中的元素(a-a∩b)
// @parameter list1
// @parameter list2
// @return []string
func ArrayNotExistAnother(list1, list2 []string) []string {
	result := make([]string, 0)
	mapTemp := make(map[string]bool)
	for _, str := range list2 {
		mapTemp[str] = true
	}
	for _, str := range list1 {
		if _, ok := mapTemp[str]; ok {
			continue
		}
		result = append(result, str)
	}
	return result
}

func ListToInString(list []string) string {
	var str strings.Builder
	str.WriteString("(")
	for index, s := range list {
		str.WriteString("'")
		str.WriteString(s)
		str.WriteString("'")
		if index+1 != len(list) {
			str.WriteString(",")
		}
	}
	str.WriteString(")")
	return str.String()
}

func IntListToInString(list []int) string {
	var str strings.Builder
	str.WriteString("(")
	for index, s := range list {
		str.WriteString("'")
		str.WriteString(fmt.Sprintf("%v", s))
		str.WriteString("'")
		if index+1 != len(list) {
			str.WriteString(",")
		}
	}
	str.WriteString(")")
	return str.String()
}

func Int64ListToInString(list []int64) string {
	var str strings.Builder
	str.WriteString("(")
	for index, s := range list {
		str.WriteString("'")
		str.WriteString(fmt.Sprintf("%v", s))
		str.WriteString("'")
		if index+1 != len(list) {
			str.WriteString(",")
		}
	}
	str.WriteString(")")
	return str.String()
}

//IsExistStrList string是否存在在list字符数组中
func IsExistStrList(list []string, str string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

func IsNotNullStr(str string) bool {
	if "" == str || 0 == len(str) {
		return false
	}
	return true
}

// JsonStrToStruct @description: jsonstring --> struct
// @parameter jsonStr
// @parameter eventStruct(*指针类型)
// @return int
func JsonStrToStruct(jsonStr string, eventStruct interface{}) (err error) {
	if err = json.Unmarshal([]byte(jsonStr), &eventStruct); err != nil {
		log.Errorln(err)
	}
	return
}

/**
map --> struct
*/
func MapToStruct(mapBean map[string]interface{}, eventStruct interface{}) {
	//将 map 转换为指定的结构体
	str, err := MapToJsonStr(mapBean)
	if err != nil {
		log.Errorln(err)
	}
	JsonStrToStruct(str, &eventStruct)
}

/**
map --> jsonstring
*/

func MapToJsonStr(mapBean map[string]interface{}) (str string, err error) {
	bytes, err := json.Marshal(mapBean)
	if err != nil {
		log.Errorln(err)
		return
	}
	return string(bytes), err
}

/**
struct -> jsonString
*/
func StructToJsonStr(eventStruct interface{}) (str string, err error) {
	buf, err := json.Marshal(eventStruct) //格式化编码
	if err != nil {
		log.Errorln(err)
		return
	}
	return string(buf), err
}

/**
struct -> map
*/
func StructToMap(obj interface{}) map[string]interface{} {
	//获取参数o的类型
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	field := t.NumField()

	var data = make(map[string]interface{})
	for i := 0; field > i; i++ {

		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// ObjectAToObjectB @description: ObjectA -> ObjectB(包含数组转换)
// @parameter objA
// @parameter objB
// @return code
func ObjectAToObjectB(objA, objB interface{}) (err error) {
	tempStr, err := json.Marshal(objA)
	if err != nil {
		log.Errorln("error:", err)
		return
	}
	err = json.Unmarshal(tempStr, objB)
	if err != nil {
		log.Errorln("error:", err)
		return
	}
	return
}

// MapKStrToList @description: map key string转list string
// @parameter inputMap
// @return []string
func MapKStrToList(inputMap map[string]int) []string {
	var resultList []string
	for k, _ := range inputMap {
		resultList = append(resultList, k)
	}
	return resultList
}
