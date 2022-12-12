package mysqldb

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func NewClient(host, port, username, password, database, net string) (mysqldb *sql.DB, err error) {
	mySqlAddr := fmt.Sprintf("%s:%s", host, port)

	cfg := mysql.Config{
		User:                 username,
		Passwd:               password,
		Net:                  net,
		Addr:                 mySqlAddr,
		DBName:               database,
		AllowNativePasswords: true,
	}

	client, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("failed connect to mysqldb: %v", err)
	}
	pingErr := client.Ping()
	if pingErr != nil {
		return nil, fmt.Errorf("failed ping to mysqldb: %v", err)
	}

	return client, nil
}
