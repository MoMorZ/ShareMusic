package api

import (
	"sharemusic/models/util"
	"strings"
)

// 歌曲详情

func SongDetail(query map[string]interface{}) map[string]interface{} {
	idArray := strings.Split(query["ids"].(string), ",")
	c := func(idArray *[]string) string {
		var ret string
		for i := 0; i < len(*idArray); i++ {
			(*idArray)[i] = "{\"id\":" + (*idArray)[i] + "}"
		}
		ret = strings.Join(*idArray, ",")
		ret = "[" + ret + "]"
		return ret
	}(&idArray)
	ids := "[" + query["ids"].(string) + "]"
	data := map[string]interface{}{
		"c":   c,
		"ids": ids,
	}
	options := map[string]interface{}{
		"crypto": "weapi",
	}
	//fmt.Println(data)
	return util.CreateRequest("https://music.163.com/weapi/v3/song/detail", data, options)
}
