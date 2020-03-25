package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type RoomInfo struct {
	ID      int64  `orm:"column(id)"`
	ListNum string `orm:"column(listnum)"`
	NowNum  int64  `orm:"column(nownum)"`
	Min     int64  `orm:"column(min)"`
	Second  int64  `orm:"column(second)"`
	Ms      int64  `orm:"column(ms)"`
	State   int64  `orm:"column(state)"`
	UserNum int64  `orm:"column(usernum)"`
	Time    string `orm:"column(time)"`
}

func (r *RoomInfo) TableName() string {
	return "roominfo"
}

func init() {
	orm.RegisterModel(new(RoomInfo))
}

func UpdateRoomInfo(id int64, time string, mState int64, nowMusic int64) error {
	o := orm.NewOrm()
	t := RoomInfo{ID: id}
	if o.Read(&t) == nil {
		t.Time = time
		t.State = mState
		t.NowNum = nowMusic
		if num, err := o.Update(&t); err == nil {
			fmt.Println("影响:", num, "更新1成功")
		} else {
			return fmt.Errorf("更新1失败 %e", err)
		}
	}
	return nil
}

func UpdateRoomInfoUserName(id int64, userNum int64) error {
	o := orm.NewOrm()
	t := RoomInfo{ID: id}
	if o.Read(&t) == nil {
		t.UserNum = userNum
		if num, err := o.Update(&t); err == nil {
			fmt.Println("影响:", num, "更新2成功")
		} else {
			return fmt.Errorf("更新2失败 %e", err)
		}
	}
	return nil
}

func GetOneById(id int64) (res RoomInfo, err error) {
	o := orm.NewOrm()
	res = RoomInfo{ID: id}
	err = o.Read(&res)
	return
}

func DelRoomInfo(id int64) error {
	o := orm.NewOrm()
	t := RoomInfo{ID: id}
	if o.Read(&t) == nil {
		if _, err := o.Delete(&t); err == nil {
			fmt.Println("删除成功")
		} else {
			return fmt.Errorf("删除失败 %e", err)
		}
	}
	return nil
}

func GetAll() {
	o := orm.NewOrm()
	qs := o.QueryTable(new(RoomInfo))
	var res []RoomInfo
	fmt.Println(qs.All(&res))
	fmt.Println(res)
}
