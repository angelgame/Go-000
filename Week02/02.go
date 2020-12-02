package main

import (
	"database/sql"
	"error"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Service struct {
	d *Dao
}
type Dao struct{}

func (s *Service) getUserToken(userId int64) (string, error) {
	userId, err := s.d.getTokenById(userId)
	if err == sql.ErrNoRows {
		//获取为空的时候一些特殊的逻辑
	}
	if err != nil {
		return "", errors.Errorf("get user token failed, userId=%d,err=%w", userId, err)
	}
	return userId, nil
}

func (d *Dao) getTokenById() (userToken string, err error) {
	row, err := db.QueryRow("select token from users where id=?", 1)
	err := row.Scan(userId)
	return userToken, err
}

func main() {
	fmt.println("hello world")
}
