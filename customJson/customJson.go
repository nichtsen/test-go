package customjson

import (
	"encoding/json"
	"time"
)

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
