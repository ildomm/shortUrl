package users

import (
	"github.com/go-sql-driver/mysql"
	"github.com/ildomm/linx_challenge/db"
	"github.com/ildomm/linx_challenge/models"
	"log"
	"time"
)

var tableName = "users"
var tableFields = "id, identify, created_at"

func ById(id int64) *models.User {
	var entry *models.User

	session := db.Mysql()
	session.Select("*").From(tableName).
		Where("Id = ?", id).
		LoadOne(&entry)
	return entry
}

func UserId(identify string) int64 {
	var id int64 = 0

	session := db.Mysql()
	session.Select("id").From(tableName).
		Where("identify = ?", identify).
		LoadOne(&id)

	return id
}

func ByIdentify(identify string) *models.User {
	var entry *models.User

	session := db.Mysql()
	session.Select("*").From(tableName).
		Where("identify = ?", identify).
		LoadOne(&entry)
	return entry
}

func IdentifyExists(identify string) bool {
	var total int

	session := db.Mysql()
	err :=
		session.Select("COUNT(*) as total").
			From(tableName).
			Where("identify = ?", identify).
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

func Insert(identify string) (*models.User, *mysql.MySQLError) {
	entry := new(models.User)

	session := db.Mysql()
	stmt := session.InsertInto(tableName).
		Pair("identify", identify).
		Pair("created_at", time.Now().Format(db.TIME_FORMAT))

	var entryID int64
	_, err := stmt.Exec()

	if err != nil {
		errDB := err.(*mysql.MySQLError)
		return nil, errDB
	} else {
		_ = session.Select("LAST_INSERT_ID() as ID").
			From(tableName).LoadOne(&entryID)
		entry.ID = string(entryID)
		return entry, nil
	}

	return nil, nil
}

func Delete(id int64) bool {
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