package mysql

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	mysqlDB *sql.DB
}

func ConnectToMySQL() *MySQL {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")

	//defer db.Close()

	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("MySQL Connected !\n")
	}

	return &MySQL{
		mysqlDB: db,
	}
}

func (my *MySQL) GetData(http.ResponseWriter, *http.Request) {
	res, err := my.mysqlDB.Query("SELECT * FROM admin")

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Printf("%v\n --> gelen data", res)

}
