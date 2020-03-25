package api

import (
	"fmt"
	"sharemusic/models/util"
)

func SongUrl(query map[string]interface{}) map[string]interface{} {
	ids := "[" + query["id"].(string) + "]"
	fmt.Println(query["id"])
	data := map[string]interface{}{
		"ids": ids,
		"br":  query["br"],
	}
	if data["br"] == nil {
		data["br"] = 999000
	}
	options := map[string]interface{}{
		"crypto": "linuxapi",
	}
	return util.CreateRequest("https://music.163.com/api/song/enhance/player/url", data, options)
}
