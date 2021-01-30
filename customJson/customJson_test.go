package customjson

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestMarshal(t *testing.T) {
	u := &User{
		Id:    3,
		Name:  "peter",
		Xtime: time.Now(),
	}
	b, err := json.MarshalIndent(u, "", "    ")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf(string(b))
}
