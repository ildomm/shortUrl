package stats

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/ildomm/linx_challenge/db"
	"github.com/ildomm/linx_challenge/models"
	"log"
)

var tableName = "urls"
var tableFields = "id, user_id, url, short, hits, created_at"

func Top10SummaryHits(userId *int64) (*models.Stats, *mysql.MySQLError) {
	var entry *models.Stats

	session := db.Mysql()
	statement := session.Select("SUM(hits) as Hits, COUNT(*) as URLCount").
		From(tableName)

	if userId != nil {
		statement.Where("user_id = ?", userId)
	}

	rows, err := statement.Rows()
	if err != nil {
		log.Println(err)
		errDB := err.(*mysql.MySQLError)
		return nil, errDB
	}

	for rows.Next() {
		entry = statsScan(rows)
	}

	urls, err := top10Urls(userId)
	if err != nil {
		log.Println(err)
		errDB := err.(*mysql.MySQLError)
		return nil, errDB
	}
	entry.TopUrls = urls

	return entry, nil
}

func top10Urls(userId *int64) ([]*models.URL, *mysql.MySQLError) {
	var entries []*models.URL

	session := db.Mysql()
	statement := session.Select("*").From(tableName).
		OrderDesc("hits").
		Limit(10)

	if userId != nil {
		statement.Where("user_id = ?", userId)
	}

	_, err := statement.Load(&entries)
	if err != nil {
		log.Println(err)
		errDB := err.(*mysql.MySQLError)
		return entries, errDB
	}

	return entries, nil
}

func statsScan(rows *sql.Rows) *models.Stats {
	var entryTmp models.Stats

	err := rows.Scan(
		&entryTmp.Hits,
		&entryTmp.URLCount,
	)
	if err != nil {
		log.Println(err)
	}

	return &entryTmp
}