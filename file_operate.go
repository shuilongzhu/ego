package ego

import (
	"encoding/base64"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"mime/multipart"
	"os"
)

// IsExistLocalFile @description: 判断一个本地文件或文件夹是否存在
// @parameter path(文件路径)
// @return bool(true:存在；false:不存在)
func IsExistLocalFile(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// ReadLocalFile @description: 读取本地文件
// @parameter localPath(绝对路径)
// @parameter bytes
// @return int
func ReadLocalFile(localPath string, bytes *[]byte) (err error) {
	file, err := os.Open(localPath)
	if nil != err {
		return
	}
	defer file.Close()
	*bytes, err = ioutil.ReadAll(file)
	if nil != err {
		return
	}
	return
}

// RemoveFile @description: 删除本地文件,不判断是否存在
// @parameter fileName
// @return int
func RemoveFile(fileName string) (err error) {
	if 0 == len(fileName) {
		return
	}
	err = os.RemoveAll(fileName)
	if nil != err {
		return
	}
	return
}

// CreateFolders @description: 创建多个文件夹
// @parameter filePaths(相对文件路径)
// @return int
func CreateFolders(filePaths []string) (err error) {
	for _, filePath := range filePaths {
		err = CreateFolder(filePath)
		if nil != err {
			return
		}
	}
	return
}

// CreateFolder @description: 创建单个文件夹
// @parameter filePath(绝对文件路径)
// @return int
func CreateFolder(filePath string) (err error) {
	if !IsExistLocalFile(filePath) {
		err = os.MkdirAll(filePath, os.ModePerm)
		if nil != err {
			return
		}
	}
	return
}

// GetFileFormat @description: 获取文件格式
// @parameter fileName
// @return string
func GetFileFormat(fileName string) string {
	var result string
	if 0 == len(fileName) {
		return result
	}
	var index = -1
	for i := 0; i < len(fileName); i++ {
		if '.' == fileName[i] {
			index = i
		}
	}
	if -1 != index {
		result = fileName[index:]
	}
	return result
}

// Base64DecodeString @description: base64 string解码
// @parameter base64Str
// @return byteList
func Base64DecodeString(base64Str string) (byteList []byte) {
	byteList, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		log.Error("base64 decode error:", err)
		return
	}
	return
}

// FileToByte @description:multipart.FileHeader file to []byte
// @parameter fileHeader
// @return err
// @return bytes
func FileToByte(fileHeader multipart.FileHeader) (err error, bytes []byte) {
	file, err := fileHeader.Open()
	if nil != err {
		return
	}

	bytes, err = ioutil.ReadAll(file)
	return
}

// OpenLocalFileToByte @description: read local file to []byte
// @parameter filePath
// @return err
// @return bytes
func OpenLocalFileToByte(filePath string) (err error, bytes []byte) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	bytes, err = ioutil.ReadAll(file)
	return
}
