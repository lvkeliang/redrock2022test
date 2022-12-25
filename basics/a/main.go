package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type user struct {
	username string    `json:"username"`
	nickname string    `json:"nickname"`
	sex      uint8     `json:"sex"`
	birthday time.Time `json:"birthday"`
}

// time.Time类型无法被json序列化
func main() {
	u := user{
		username: "坤坤",
		nickname: "阿坤",
		sex:      20,
		birthday: time.Now(),
	}
	birthday := u.birthday.Format("2006-01-02")
	bs, err := json.Marshal(&struct {
		Username string
		Nickname string
		Sex      uint8
		Birthday string
	}{
		Username: u.username,
		Nickname: u.nickname,
		Sex:      u.sex,
		Birthday: birthday,
	})
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bs))
}
