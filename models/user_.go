package models

import (
	"github.com/ildomm/linx_challenge/db"
	"math"
	"time"
)

func userskey() string {
	return "linx:users"
}

func UserDBId(id string) bool {
	result := int64(0)
	conn := db.Redis()
	value := conn.Exists(userskey(), id)
	if value != nil {
		result = value.Val()
	}
	return result > 0
}

func (m *User) Save() {
	conn := db.Redis()
	conn.Set(userskey(), m.ID,(time.Duration(math.MaxInt16) * time.Hour))
}