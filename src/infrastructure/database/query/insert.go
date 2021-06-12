package query

import (
	"database/sql"
	"domain/model"
	_ "github.com/go-sql-driver/mysql"
	)

func Ins(table string, user *model.User) (err error){
	db, err := sql.Open("mysql", "user:@/go_demo")
	if err != nil {
		return
	}
	defer db.Close()

	rows, err := db.Prepare(makeInsQuery(table, user))
	if err != nil {
		return
	}
	result ,err := rows.Exec(user.Name, user.Email, user.Password)

	if err != nil {
		return
	}

	id, err := result.LastInsertId()
	user.Id = id
	return nil
}

func makeInsQuery(table string, user *model.User) (query string){
	query = "insert into " + table + "(name, email, password) values(?, ?, ?)"
	return
}
