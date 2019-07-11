package urls

import (
	"github.com/go-sql-driver/mysql"
	"github.com/ildomm/linx_challenge/db"
	"github.com/ildomm/linx_challenge/models"
	"log"
	"time"
)

var tableName = "urls"
var tableFields = "id, user_id, url, short, hits, created_at"

func ById(id string) *models.URL {
	var entry *models.URL

	session := db.Mysql()
	session.Select("*").From(tableName).
		Where("id = ?", id).
		LoadOne(&entry)
	return entry
}

func ByUrl(url string) *models.URL {
	var entry *models.URL

	session := db.Mysql()
	session.Select("*").From(tableName).
		Where("url = ?", url).
		LoadOne(&entry)
	return entry
}

func ByShort(short string) *models.URL {
	var entry *models.URL

	session := db.Mysql()
	session.Select("*").From(tableName).
		Where("short = ?", short).
		LoadOne(&entry)
	return entry
}

func UrlExists(url string) bool {
	var total int

	session := db.Mysql()
	err :=
		session.Select("COUNT(*) as total").
			From(tableName).
			Where("url = ?", url).
			LoadOne(&total)

	if err != nil {
		log.Printf("Error select from table %s : %s", tableName, err.Error())
	}
	if total > 0 {
		return true
	} else {
		return false
	}
}

func Insert(entry models.URL, userId int64) (*models.URL, *mysql.MySQLError) {

	entry.ShortURL = entry.GenerateShortURL()

	session := db.Mysql()
	stmt := session.InsertInto(tableName).
		Pair("user_id", userId).
		Pair("url", entry.URL).
		Pair("short", entry.ShortURL).
		Pair("hits", 0).
		Pair("created_at", time.Now().Format(db.TIME_FORMAT))

	var entryID int64
	_, err := stmt.Exec()

	if err != nil {
		errDB := err.(*mysql.MySQLError)
		return nil, errDB
	} else {
		_ = session.Select("LAST_INSERT_ID() as ID").
			From(tableName).LoadOne(&entryID)
		entry.ID = entryID
		return &entry, nil
	}

	return nil, nil
}

func Delete(id string) bool {
	var total int

	session := db.Mysql()
	stmt := session.DeleteFrom(tableName).
		Where("id = ?", id)

	_, err := stmt.Exec()
	if err != nil {
		log.Printf("Error deleting from table %s : %s", tableName, err.Error())
	}
	if total > 0 {
		return true
	} else {
		return false
	}
}

func DeleteByUser(userId int64) bool {
	var total int

	session := db.Mysql()
	stmt := session.DeleteFrom(tableName).
		Where("user_id = ?", userId)

	_, err := stmt.Exec()
	if err != nil {
		log.Printf("Error deleting from table %s : %s", tableName, err.Error())
	}
	if total > 0 {
		return true
	} else {
		return false
	}
}

func IncreaseHits(id int64) *mysql.MySQLError {
	session := db.Mysql()
	_, err := session.UpdateBySql(" UPDATE urls " +
		" SET hits = (hits + 1) " +
		" WHERE id = '" + string(id) + "'").
		Exec()

	if err != nil {
		errDB := err.(*mysql.MySQLError)
		return errDB
	}

	return nil
}