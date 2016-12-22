package dao

import (
	"database/sql"
	"log"
)

//获取ID type 值
func GetITValue(userID int64, tp int32) (value int64, err error) {
	var timestamp int64
	var row *sql.Rows
	row, err = Tdb.Query("SELECT TIMESTAMP FROM IM_TIMESTAMP_IDTYPE_4 WHERE ID =? AND TYPE= ? limit 1", userID, tp)
	defer row.Close()
	if err != nil {
		log.Printf("GetITValue err:%v\n", err)
		return timestamp, err
	}
	row.Next()
	row.Scan(&timestamp)
	log.Printf("%v", timestamp)
	return timestamp, nil
}

func AddITValue(userID int64, tp int32, value int64) (err error) {
	_, err = Tdb.Exec("INSERT INTO IM_TIMESTAMP_IDTYPE_4(ID,TYPE,TIMESTAMP) VALUES(?,?,?)", userID, tp, value)
	if err != nil {
		log.Printf("AddITValue err:%v\n", err)
		return err
	}
	return nil
}
func UpdateITValue(userID int64, tp int32, value int64) error {
	_, err := Tdb.Exec("UPDATE  IM_TIMESTAMP_IDTYPE_4 SET TIMESTAMP=? WHERE ID=? AND TYPE=?", value, userID, tp)
	if err != nil {
		log.Printf("UpdateITValue err:%v\n", err)
		return err
	}
	return nil
}
