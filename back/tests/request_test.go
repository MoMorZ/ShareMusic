package test

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"reflect"
	"sharemusic/models/api"
	"sharemusic/models/util"
	"testing"
)

// 专辑详细信息 weapi
func TestRequest1(t *testing.T) {
	data := map[string]interface{}{
		"id": "32311",
	}
	url := "https://music.163.com/weapi/album/detail/dynamic"
	options := map[string]interface{}{
		"crypto": "weapi",
	}
	answer := util.CreateRequest(url, data, options)
	fmt.Println(answer["body"])
	fmt.Println(reflect.TypeOf(answer["body"]))
}

// 歌曲url linuxapi
func TestRequest2(t *testing.T) {
	data := map[string]interface{}{
		"ids": "[3894312]",
		"br":  999000,
	}
	url := "https://music.163.com/api/song/enhance/player/url"
	options := map[string]interface{}{
		"crypto": "linuxapi",
	}
	answer := util.CreateRequest(url, data, options)
	fmt.Println(answer)
}

// 歌单详情
func TestRequest3(t *testing.T) {
	query := map[string]interface{}{}
	var str []string
	str = append(str, "374012053")
	query["id"] = str
	tt := api.PlayListDetail(query)
	res := tt["body"].(string)
	m := make(map[string]interface{})
	_ = json.Unmarshal([]byte(res), &m)
	fmt.Println(m["privileges"])
	fmt.Println(reflect.TypeOf(m["privileges"]))
	a := m["privileges"].([]interface{})
	fmt.Println(a[0])
	fmt.Println("type of m is ", reflect.TypeOf(m))

}

func TestUrl(t *testing.T) {
	url := "https://music.163.com/api/album/detail/dynamic?id=32311"
	req := httplib.Post(url)
	orignData := map[string]interface{}{
		"id": "32311",
	}
	cryptoData := util.WeApi(orignData)
	fmt.Println(orignData)
	fmt.Println(cryptoData["params"])
	fmt.Println(cryptoData["encSecKey"])
	fmt.Println("---------------------------------------")
	req.Param("params", fmt.Sprintf("%v", cryptoData["params"]))
	req.Param("encSecKey", fmt.Sprintf("%v", cryptoData["encSecKey"]))
	//req.Param("params",cryptoData["params"])
	//req.Param("encSecKey",cryptoData["encSecKey"])
	fmt.Println(req.String())
	//req.Header("Accept", "*/*")
	//req.Header("Accept-Language", "zh-CN,zh;q=0.8,gl;q=0.6,zh-TW;q=0.4")
	//req.Header("Content-Type", "application/x-www-form-urlencoded")
	//req.Header("Cookie", "appver=2.0.2")
	//req.Header("cache-control", "no-cache")
}
