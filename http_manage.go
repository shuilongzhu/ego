package ego

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

var ClientIp = GetClientIp() //获取本机真正IP

// CommonPostCall @description:公共post请求调用
// @parameter auth
// @parameter uri
// @parameter reqObject
// @parameter respObject(*指针类型)
// @return int
func CommonPostCall(auth string, uri string, reqObject interface{}, respObject interface{}) (err error) {
	jsonData := make([]byte, 0, 0)
	if str, ok := reqObject.(string); ok {
		jsonData = []byte(str)
	} else {
		jsonData, err = json.Marshal(reqObject)
		if err != nil {
			return
		}
	}
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003c"), []byte("<"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u003e"), []byte(">"))
	jsonData = bytes.ReplaceAll(jsonData, []byte("\\u0026"), []byte("&"))
	//设置post请求,第三个参数传byte类型,很关键！
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return
	}
	//设置请求数据格式
	req.Header.Set("Content-Type", "application/json")
	if 0 != len(auth) {
		//授权！！！
		req.Header.Set("authorization", auth)
	}
	//获取客户端对象，发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(err)
		return
	}
	defer resp.Body.Close()
	//读取返回值
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
		return
	}
	log.Println(string(res))
	return JsonStrToStruct(string(res), &respObject)
}

// CommonGetCall @description: 公共get请求调用
// @parameter auth
// @parameter uri
// @parameter reqObject
// @parameter respObject
// @return int
func CommonGetCall(auth string, uri string, respObject interface{}) (err error) {
	jsonData := make([]byte, 0, 0)
	//设置get请求,第三个参数传byte类型,很关键！
	req, err := http.NewRequest(http.MethodGet, uri, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return
	}
	//设置请求数据格式
	req.Header.Set("Content-Type", "application/json")
	if 0 != len(auth) {
		//授权！！！
		req.Header.Set("authorization", auth)
	}
	//获取客户端对象，发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(err)
		return
	}
	defer resp.Body.Close()
	//读取返回值
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorln(err)
		return
	}
	return JsonStrToStruct(string(res), &respObject)
}

// GetClientIp @description: 获取服务IP地址
// @return string
func GetClientIp() string {
	var str string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return str
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						var tempStr = ipnet.IP.String()
						var strList = strings.Split(tempStr, ".")
						if "1" != strList[len(strList)-1] { //服务器ip地址最后一位不为1
							str = tempStr
						}
					}
				}
			}
		}
	}
	log.Println("service ip is ", str)
	return str
}

func GenSOAPXml(nameSpace string, methodName string, valueStr string) string {
	soapBody := `<?xml version="1.0" encoding="utf-16"?>`
	soapBody += `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">`
	soapBody += " <soap:Body>"
	soapBody += "<" + methodName + " xmlns=\"" + nameSpace + "\">"
	//以下是具体参数
	soapBody += valueStr
	soapBody += "</" + methodName + "> </soap:Body></soap:Envelope>"
	return soapBody
}
