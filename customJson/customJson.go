package customjson

import (
	"encoding/json"
	"time"
)

type Temple struct {
	ID int64 `json:"id,string"`
	Lv int64 `json:"lv"`
	// 升级所需时间（秒） map[Lv]Seconds
	LvSeconds map[int64]int64 `json:"lv_seconds"`
	ChildID   []int64         `json:"child_id"`
}

type User struct {
	Id    int       `json:"id"`
	Name  string    `json:"name"`
	Xtime time.Time `json:"-"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return json.Marshal(struct {
		Alias
		Xtime int64
	}{
		Alias: Alias(*u),
		Xtime: u.Xtime.Unix(),
	})
}

type pencil struct{}

func (p pencil) Write([]byte) (int, error) { return 0, nil }
