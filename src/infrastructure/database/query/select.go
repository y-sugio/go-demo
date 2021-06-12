package query

import (
	"database/sql"
	"domain/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

func Select(table string, id int64) (*model.User, error){
	db, err := sql.Open("mysql", "user:@/go_demo")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("select * from " + table + " where id = " + strconv.Itoa(int(id)))

	var user model.User
	err = row.Scan(&user.Id, &user.Name, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}
