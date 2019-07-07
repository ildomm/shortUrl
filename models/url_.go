package models

import (
	"fmt"
	"github.com/ildomm/linx_challenge/config"
	"math/rand"
	"strconv"
	"time"
)

func urlsKey(user string) string {
	return "linx:urls:" + strconv.FormatInt(UserDBId(user), 10)
}

func UrlExists(user string, url string) bool {
	return UrlDBId(user, url) > 0
}

func UrlDBId(user string, url string) int64 {

	return 0
}

func (m *URL) SaveUrlKey(user string) {

}

func (m *URL) GenerateShort() {
	key1 := 10 + rand.Float32()*(99-10)
	key2 := time.Now().Unix()
	short := fmt.Sprintf("%f", key1) + fmt.Sprintf("%f", key2)

	m.ShortURL = config.App.Runtime.Url + short
}

func (m *URL) SaveUrlDetail(user string) {
}

func (m *URL) LoadUrlDetail(user string) URL {
	m.Hits = 0
	m.ShortURL = "sdsd"
	return *m
}