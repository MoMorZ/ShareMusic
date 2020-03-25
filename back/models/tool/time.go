package tool

import (
	"fmt"
	"time"
)

func GetTime(offset bool) string {
	date := time.Now()
	if offset == true {
		h, _ := time.ParseDuration("-12h")
		date.Add(h)
	}
	t1 := date.Year()   //年
	t2 := date.Month()  //月
	t3 := date.Day()    //日
	t4 := date.Hour()   //小时
	t5 := date.Minute() //分钟
	t6 := date.Second() //秒
	//t7:=date.Nanosecond()  //纳秒
	now := fmt.Sprintf("%d%d%d%d%d%d", t1, t2, t3, t4, t5, t6)
	return now
}
