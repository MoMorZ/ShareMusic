package api

import (
	"sharemusic/models/util"
)

func PlayListDetail(query map[string]interface{}) map[string]interface{} {
	data := map[string]interface{}{
		"id": query["id"],
		"n":  100000,
		//"n": 3,
		"s": query["s"],
	}
	if query["s"] == nil {
		data["s"] = 8
	}
	options := map[string]interface{}{
		"crypto": "linuxapi",
		"cookie": query["cookie"],
		"proxy":  query["proxy"],
	}
	return util.CreateRequest("https://music.163.com/weapi/v3/playlist/detail", data, options)
}
