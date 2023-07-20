package ego

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

var dfsClient = NewDfsClient()

func NewDfsClient() *DfsHTTPClient {
	var client = &http.Client{
		Transport: &http.Transport{
			//Proxy:               http.ProxyFromEnvironment, //HTTP代理
			TLSHandshakeTimeout: 15 * time.Second, //TLS握手的超时时间
			MaxIdleConns:        200,              //最大空闲数(默认100)
			MaxIdleConnsPerHost: 20,               //最大并发数(默认2)
		},
		Timeout: 40 * time.Second, //整个请求的超时时间
	}
	return &DfsHTTPClient{Client: client, DfsMasterAddress: "http://172.10.50.239:9333", DfsFilerAddress: "http://172.10.50.239:8888"}
}

func TestDfs(t *testing.T) {
	err, bytes := OpenLocalFileToByte("C:\\Users\\xxx\\Desktop\\test.jpg")
	if nil != err {
		fmt.Errorf(err.Error())
		return
	}
	err = dfsClient.Post("/Test/test.jpg", bytes)
}
