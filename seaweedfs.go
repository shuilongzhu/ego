package ego

import (
	"archive/tar"
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
)

const (
	postSuccess   = 201
	deleteSuccess = 204
	notFound      = 404
	successfully  = 200
)

var (
	DfsMethodParamErr = errors.New("seaweedfs method parameter error")
)

type SeaweedfsMethod interface {
	Post(filePath string, byteList []byte) error
	Delete(filePath string) error
	Get(filePath string) (err error, bytes []byte)
	Head(filePath string) (error, bool)
	PostAppend(filePath string, byteList []byte) error
	UploadZip(ZipFiles []ZipFile, zipPath string) error
}

// Post @description: 分布式存储seaweedfs客户端post请求方法
// @receiver dfsClient
// @parameter filePath(文件路径，最前面以/开头，例如:/data_center/images/1.jpg)
// @parameter byteList(字节数组)
// @return error
func (dfsClient *DfsHTTPClient) Post(filePath string, byteList []byte) error {

	if !dfsParamCheck(filePath) {
		return DfsMethodParamErr
	}

	var body = &bytes.Buffer{}
	var writer = multipart.NewWriter(body)

	// 写入文件到multipart/form-data请求体
	var part, err = writer.CreateFormFile("file", "")
	if nil != err {
		return errors.New(fmt.Sprintf("Post writer.CreateFormFile() filePath:%s err:%v", filePath, err))
	}

	_, err = part.Write(byteList)
	if nil != err {
		return errors.New(fmt.Sprintf("Post part.Write() filePath:%s err:%v", filePath, err))
	}

	err = writer.Close()
	if nil != err {
		return errors.New(fmt.Sprintf("Post writer.Close() filePath:%s err:%v", filePath, err))
	}

	var req, err2 = http.NewRequest(http.MethodPost, dfsClient.DfsFilerAddress+filePath, body)
	if nil != err2 {
		return errors.New(fmt.Sprintf("Post http.NewRequest() filePath:%s err:%v", filePath, err2))
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	var resp, err3 = dfsClient.Client.Do(req)
	if nil != err3 {
		return errors.New(fmt.Sprintf("Post dfsClient.Client.Do() filePath:%s err:%v", filePath, err3))
	}
	defer resp.Body.Close()

	//resp.StatusCode为201,创建文件成功
	if postSuccess == resp.StatusCode {
		return nil
	}

	return errors.New(fmt.Sprintf("Post dfsClient.Client.Do() filePath:%s code:%d err:%v", filePath, resp.StatusCode, resp.Status))
}

// Delete @description: 分布式存储seaweedfs客户端delete请求方法
// @receiver dfsClient
// @parameter filePath(文件路径，最前面以/开头，例如:/data_center/images/1.jpg)
// @return error
func (dfsClient *DfsHTTPClient) Delete(filePath string) error {

	if !dfsParamCheck(filePath) {
		return DfsMethodParamErr
	}

	var req, err = http.NewRequest(http.MethodDelete, dfsClient.DfsFilerAddress+filePath+"?recursive=true", nil) //可以强制删除非空的文件
	if nil != err {
		return errors.New(fmt.Sprintf("Delete http.NewRequest() filePath:%s err:%v", filePath, err))
	}

	var resp, err2 = dfsClient.Client.Do(req)
	if nil != err2 {
		return errors.New(fmt.Sprintf("Delete c.Client.Do() filePath:%s err:%v", filePath, err2))
	}
	defer resp.Body.Close()
	//resp.StatusCode:204代表删除成功;200代表已经被删除过了
	if deleteSuccess == resp.StatusCode || successfully == resp.StatusCode {
		return nil
	}

	return errors.New(fmt.Sprintf("Delete dfsClient.Client.Do() filePath:%s code:%d err:%s", filePath, resp.StatusCode, resp.Status))
}

// Get @description: 分布式存储seaweedfs客户端get请求方法
// @receiver dfsClient
// @parameter filePath(文件路径，最前面以/开头，例如:/data_center/images/1.jpg)
// @return error
// @return bytes
func (dfsClient *DfsHTTPClient) Get(filePath string) (err error, bytes []byte) {

	if !dfsParamCheck(filePath) {
		err = DfsMethodParamErr
		return
	}

	req, err1 := http.NewRequest(http.MethodGet, dfsClient.DfsFilerAddress+filePath, nil)
	if nil != err {
		err = errors.New(fmt.Sprintf("Get http.NewRequest() filePath:%s err:%v", filePath, err1))
		return
	}

	resp, err2 := dfsClient.Client.Do(req)
	if nil != err2 {
		err = errors.New(fmt.Sprintf("Get dfsClient.Client.Do() filePath:%s err:%v", filePath, err2))
		return
	}
	defer resp.Body.Close()

	if successfully == resp.StatusCode {
		//读取请求响应
		bytes, err = io.ReadAll(resp.Body)
		if nil != err {
			err = errors.New(fmt.Sprintf("Get io.ReadAll() filePath:%s err:%v", filePath, err))
		}
		return
	}

	if notFound == resp.StatusCode {
		fmt.Printf("Get() filePath:%s 文件不存在", filePath)
		return
	}
	err = errors.New(fmt.Sprintf("Get dfsClient.Client.Do() filePath:%s code:%d err:%v", filePath, resp.StatusCode, err2))
	return
}

// Head @description: 分布式存储seaweedfs客户端head请求方法,判断文件是否存在
// @receiver dfsClient
// @parameter filePath
// @return error
// @return bool
func (dfsClient *DfsHTTPClient) Head(filePath string) (error, bool) {

	if !dfsParamCheck(filePath) {
		return DfsMethodParamErr, false
	}

	var req, err = http.NewRequest(http.MethodHead, dfsClient.DfsFilerAddress+filePath, nil)
	if nil != err {
		err = errors.New(fmt.Sprintf("Head http.NewRequest() filePath:%s err:%v", filePath, err))
		return err, false
	}

	var resp, err2 = dfsClient.Client.Do(req)
	if nil != err2 {
		err = errors.New(fmt.Sprintf("Head dfsClient.Client.Do() filePath:%s err:%v", filePath, err2))
		return err, false
	}
	defer resp.Body.Close()

	//resp.StatusCode返回200，响应成功并存在
	if successfully == resp.StatusCode {
		return nil, true
	}
	//resp.StatusCode返回404，请求成功但文件不存在存在
	if notFound == resp.StatusCode {
		return nil, false
	}

	err = errors.New(fmt.Sprintf("Head filePath:%s resp.StatusCode:%d,resp.Status:%s", filePath, resp.StatusCode, resp.Status))
	return err, false
}

// PostAppend @description: 附加到文件中（附加到文件中的单个文件资源会按照资源名称升序排列）
// @receiver dfsClient
// @parameter filePath
// @parameter byteList
// @return error
func (dfsClient *DfsHTTPClient) PostAppend(filePath string, byteList []byte) error {

	if !dfsParamCheck(filePath) {
		return DfsMethodParamErr
	}

	var body = &bytes.Buffer{}
	var writer = multipart.NewWriter(body)

	// 写入文件到multipart/form-data请求体
	var part, err = writer.CreateFormFile("file", "")
	if nil != err {
		return errors.New(fmt.Sprintf("Post writer.CreateFormFile() filePath:%s err:%v", filePath, err))
	}

	_, err = part.Write(byteList)
	if nil != err {
		return errors.New(fmt.Sprintf("PostAppend part.Write() filePath:%s err:%v", filePath, err))
	}

	err = writer.Close()
	if nil != err {
		return errors.New(fmt.Sprintf("PostAppend writer.Close() filePath:%s err:%v", filePath, err))
	}

	var req, err2 = http.NewRequest(http.MethodPost, dfsClient.DfsFilerAddress+filePath+"?op=append", body)
	if nil != err2 {
		return errors.New(fmt.Sprintf("PostAppend http.NewRequest() filePath:%s err:%v", filePath, err2))
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	var resp, err3 = dfsClient.Client.Do(req)
	if nil != err3 {
		return errors.New(fmt.Sprintf("PostAppend dfsClient.Client.Do() filePath:%s err:%v", filePath, err3))
	}
	defer resp.Body.Close()

	//resp.StatusCode为201,创建文件成功
	if postSuccess == resp.StatusCode {
		return nil
	}
	return errors.New(fmt.Sprintf("PostAppend dfsClient.Client.Do() filePath:%s code:%d err:%v", filePath, resp.StatusCode, resp.Status))
}

// UploadZip @description: 从seaweedfs获取资源打包压缩分片上传到seaweedfs
// @receiver dfsClient
// @parameter ZipFiles
// @parameter zipPath
// @return err
func (dfsClient *DfsHTTPClient) UploadZip(ZipFiles []ZipFile, zipPath string) (err error) {

	if 0 == len(ZipFiles) {
		return DfsMethodParamErr
	}

	// 压缩多个文件流到单个tar文件
	var buf bytes.Buffer
	var w, file = tar.NewWriter(&buf), make([]byte, 0)

	for i, zipFile := range ZipFiles {

		if !dfsParamCheck(zipFile.FilePath) {
			return DfsMethodParamErr
		}

		//获取资源二进制数组
		if err, file = dfsClient.Get(zipFile.FilePath); nil != err && 0 == i {
			return err
		}
		//为空，继续
		if 0 == len(file) {
			continue
		}

		//压缩资源
		if err = addFileToZip(w, zipFile.FileName, file); nil != err {
			return err
		}

		//每10张资源上传
		if i%10 == 0 {
			//分区上传
			err = dfsClient.PostAppend(zipPath, buf.Bytes())
			if nil != err {
				return err
			}
			buf.Reset() //重置buf
		}
	}

	if err = w.Close(); nil != err {
		return
	}

	//上传最后部分
	return dfsClient.PostAppend(zipPath, buf.Bytes())
}

// addFileToZip @description: 单个文件流压缩
// @parameter w
// @parameter filename
// @parameter fileContent
// @return int
func addFileToZip(w *tar.Writer, filename string, fileContent []byte) (err error) {
	if err = w.WriteHeader(&tar.Header{
		Name: filename,
		//Mode: 0600,
		Size: int64(len(fileContent)),
	}); err != nil {
		return
	}
	if _, err = w.Write(fileContent); err != nil {
		return
	}
	return
}

// dfsParamCheck @description: seaweedfs分布式存储操作方法入参校验
// @parameter field
// @return bool(是否合规) true:合规；false:不合规
func dfsParamCheck(filePath string) bool {
	if "" == filePath || '/' != filePath[0] {
		return false
	}
	return true
}
