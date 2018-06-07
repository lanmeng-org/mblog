package model

import (
	"MBlog/util"
	"errors"
	"time"
)

type User struct {
	ID        int
	Name      string
	Password  string
	Nickname  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetUserByName(name string) (*User, error) {
	db, err := util.GetMysqlConn()
	if err != nil {
		return nil, err
	}

	user := new(User)
	err = db.QueryRow("SELECT id,name,password, nickname, created_at, updated_at FROM users where name=?", name).
		Scan(&user.ID, &user.Name, &user.Password, &user.Nickname, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user *User) error {
	db, err := util.GetMysqlConn()
	if err != nil {
		return err
	}

	curTime := time.Now()
	curTimeStr := curTime.Format("2006-01-02 15:04:05")

	result, err := db.Exec("INSERT INTO users(name,password,nickname,created_at,updated_at) VALUES(?,?,?,?,?)",
		user.Name, user.Password, user.Nickname, curTimeStr, curTimeStr)
	if err != nil {
		return err
	}

	i, err := result.LastInsertId()
	if err != nil {
		return nil
	}

	if i == 0 {
		return errors.New("create fail")
	}

	user.ID = int(i)
	user.CreatedAt = curTime
	user.UpdatedAt = curTime

	return nil
}

func SaveUser(user *User) error {
	db, err := util.GetMysqlConn()
	if err != nil {
		return err
	}

	if user.ID <= 0 {
		return CreateUser(user)
	}

	result, err := db.Exec("UPDATE users SET name=?,password=?,nickname=?,created_at=?,updated_at=? WHERE id=?",
		user.Name, user.Password, user.Nickname, user.CreatedAt, user.UpdatedAt, user.ID)

	if err != nil {
		return err
	}

	i, err := result.RowsAffected()
	if err != nil {
		return nil
	}

	if i == 0 {
		return errors.New("update fail")
	}

	return nil
}
