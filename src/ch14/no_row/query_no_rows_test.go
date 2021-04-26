package no_rows_error

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/pkg/errors"
)

type User struct {
	id   int64
	name string
	age  int
}

func queryById(id int) (User, error) {

	sqlStr := "select * from user where id=?"
	var u User
	err := db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)

	// no user
	if err == sql.ErrNoRows {
		return u, nil
	}
	// find user error
	if err != nil {
		return u, errors.WithMessage(err, fmt.Sprintf("query.users : find id=%v user error. \n", id))
	}
	return u, nil
}

func TestQueryNoRows(t *testing.T) {
	u, err := queryById(1)
	if err != nil {
		fmt.Printf("origin error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n")
		os.Exit(1)
	}
}
