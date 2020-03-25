package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"sharemusic/models"
	"sharemusic/models/api"
	"sharemusic/models/tool"
)

func (c *MusicController) GetListDetail() {
	id := c.GetString("id")
	query := map[string]interface{}{
		"id": id,
	}
	rNum := tool.HashCode(id + tool.GetTime(false))
	answer := api.PlayListDetail(query)
	body := answer["body"].(string)
	m := make(map[string]interface{})
	_ = json.Unmarshal([]byte(body), &m)
	newBody := map[string]interface{}{
		"list":   m["privileges"],
		"code":   m["code"],
		"number": rNum,
	}
	//fmt.Println(id)
	//fmt.Println(query["id"])
	//fmt.Println(answer)
	//fmt.Println(m["privileges"])
	o := orm.NewOrm()
	room := models.RoomInfo{int64(rNum), id, -1, -1, -1, -1, 0, 1, "1"}
	_, err := o.Insert(&room)
	if err != nil {
		fmt.Println("插入发生错误 ", err)
	}
	c.Ctx.Output.SetStatus(200)
	fmt.Println(answer)
	c.Data["json"] = newBody
	c.ServeJSON()
}
