package sqlconn

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB = nil
var dialect goqu.DialectWrapper

func init() {
	if db != nil {
		return
	}

	err := Connect()
	if err != nil {
		panic(err)
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
}

func Connect() error {
	connStr := os.Getenv("MYSQL_CONN_STR")
	if connStr == "" {
		connStr = "root:@tcp(127.0.0.1:3306)/Flightening?parseTime=true"
	}

	conn, err := sql.Open("mysql", connStr)
	if err == nil {
		db = conn
	}
	dialect = goqu.Dialect("mysql")
	return err
}

func Close() error {
	return db.Close()
}
