package util

import (
	"fmt"
	"github.com/astaxie/beego/httplib"
	"regexp"
)

// 只写了POST请求的
// 因为GET只有登录用到
// 返回的BODY为JSON格式的string
// 使用按需转换
// 好吧，实际上是因为我也不知道返回什么类型比较好⊙﹏⊙∥

func CreateRequest(url string, data map[string]interface{}, options map[string]interface{}) map[string]interface{} {
	//req := httplib.Post(url)
	var req *httplib.BeegoHTTPRequest
	crypto := fmt.Sprintf("%v", options["crypto"])
	if crypto == "weapi" {
		data = WeApi(data)
		//data= map[string]interface{}{
		//	"params":"IwU5M99eDDrMgVGg+ACtDzrwjYDx6QE/h0DD675p7d+XoLpWfCMUUXoiQD+AMPna",
		//	"encSecKey":"35967b04a9f3407a80ca4f16503ac162dfae073cd36e97b12b7c9daa13755cc5871f4585454dae627d280fe5ce5295832ed6f63c6c7b583b2a7e464a8efa1d013634681af6eea4cfc248a7518404b5e8cf558f0f8de4ddc4451496299a2ec88cebff39f3a398ae7256ec35d59990cb8608d17939bd4cfa40eb18c20708f0a8d2",
		//}
		re, _ := regexp.Compile(`\w*api`)
		url = re.ReplaceAllString(url, "weapi")
	} else if crypto == "linuxapi" {
		re, _ := regexp.Compile(`\w*api`)
		tmpUrl := re.ReplaceAllString(url, "api")
		tmpData := map[string]interface{}{
			"method": "POST",
			"url":    tmpUrl,
			"params": data,
		}
		data = LinuxApi(tmpData)
		//data= map[string]interface{}{
		//	"eparams":"A0D9583F4C5FF68DE851D2893A49DE98FAFB24399F27B4F7E74C64B6FC49A965CFA972FA5EA3D6247CD6247C8198CB878588AFAA94B0E0C872FF37D6781726EEF1A7727DF25C2EA3128E3EEDFC533649EB2EED728D9F7CD04FE8162B0D2BC8030E0B2B4275BCDB24CC212BE79672C8EF233ECCFBFCC7ACB34AB513EBBA50B1EF",
		//}
		url = "https://music.163.com/api/linux/forward"
	}
	req = httplib.Post(url)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	var key, value string
	// 把data的参数添加到请求中
	for k, v := range data {
		key = fmt.Sprintf("%v", k)
		value = fmt.Sprintf("%v", v)
		//fmt.Printf("%s\n%s\n", key, value)
		req.Param(key, value)
	}
	answer := map[string]interface{}{
		"status": 500,
		"body":   map[string]interface{}{},
	}
	resp, err := req.String()
	//var resp interface{}
	//err :=req.ToJSON(resp)
	//resp,err:=req.Response()

	if err != nil {
		answer["status"] = 502
		answer["body"] = map[string]interface{}{
			"code": 502,
			"msg":  err.Error(),
		}
	}
	answer["status"] = 200
	answer["body"] = resp
	return answer
}
