package main

import (
	"encoding/json"
	"time"
)

func main() {}

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
