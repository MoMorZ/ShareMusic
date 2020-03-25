package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"sharemusic/models"
	"strconv"
)

type MusicController struct {
	beego.Controller
}

func (c *MusicController) Heart() {
	id, _ := c.GetInt64("id")
	time := c.GetString("time")
	mState, _ := c.GetInt64("mstate")
	nowMusic, _ := c.GetInt64("nowmusic")

	nowTime, _ := strconv.ParseInt(time, 10, 64)
	res, _ := models.GetOneById(id)
	oldTime, _ := strconv.ParseInt(res.Time, 10, 64)
	newBody := map[string]interface{}{
		"code": "200",
	}

	fmt.Println("Heart 数据 ", id, time, mState, nowMusic)
	fmt.Println("sql 结果 ", res)

	if nowTime > oldTime && (res.State != mState || res.NowNum != nowMusic) {
		c.Ctx.Output.SetStatus(200)
		newBody = map[string]interface{}{
			"code":      "200",
			"nowmusic":  -1,
			"playstate": -1,
			"time":      time,
		}
		_ = models.UpdateRoomInfo(id, time, mState, nowMusic)
	} else {
		c.Ctx.Output.SetStatus(200)
		newBody = map[string]interface{}{
			"code":      "200",
			"nowmusic":  res.NowNum,
			"playstate": res.State,
			"time":      res.Time,
		}
	}
	fmt.Println(newBody)
	c.Data["json"] = newBody
	c.ServeJSON()
}
