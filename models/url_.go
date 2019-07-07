package models

import (
	"fmt"
	"github.com/ildomm/linx_challenge/config"
	"math/rand"
	"time"
)


func (m *URL) GenerateShortURL() string {
	key1 := 10 + rand.Float32()*(99-10)
	key2 := time.Now().Unix()
	short := fmt.Sprintf("%f", key1) + fmt.Sprintf("%f", key2)

	m.ShortURL = config.App.Runtime.Url + short
	return m.ShortURL
}
