package sessiondao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var dsn = "wc:wcuser@sviluppo.mtl.it:3306/wc?parseTime=true"
var logger *log.Logger

func openConn() *sql.DB {
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.Fatalf("Database link error! %T %s", err, err)
	}
	return conn
}
func GetUserIdBySession(sessionid string) (int64, error) {
	conn := openConn()
	defer conn.Close()
	var id int64
	err := conn.QueryRow("SELECT user_id FROM wc_session WHERE session_id = ?", sessionid).Scan(&id)
	return id, err
}

func InsertSession(sessionid string, userid int64) error {
	conn := openConn()
	defer conn.Close()
	now := time.Now().Local().UnixNano()
	_, err := conn.Exec("INSERT INTO wc_session VALUES(NULL,?,?,?)", now, userid, sessionid)
	return err
}
