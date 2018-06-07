package util

import "database/sql"

var dbs = make(map[string]*sql.DB, 1)

func GetMysqlConn() (conn *sql.DB, err error) {
	if conn, ok := dbs["mysql"]; ok {
		conn.Ping()
		return conn, nil
	}

	conn, err = sql.Open("mysql", BlogConfig.DB.GenDsn())
	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(-1)
	conn.SetMaxIdleConns(10)
	conn.SetMaxOpenConns(20)
	dbs["mysql"] = conn

	return conn, nil
}