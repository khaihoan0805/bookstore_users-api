package users

import (
	"github.com/khaihoan0805/bookstore_users-api/utils/mysql_utils"

	"github.com/khaihoan0805/bookstore_users-api/datasource/mysql/users_db"
	"github.com/khaihoan0805/bookstore_users-api/utils/date_utils"

	"github.com/khaihoan0805/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser = "DELETE FROM users WHERE id=?;"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(saveErr)
	}
	user.Id = userId

	// current := userDB[user.Id]
	// if current != nil {
	// 	if current.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	// }
	// userDB[user.Id] = user

	userDB[user.Id] = user

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}
	user.DateCreated = date_utils.GetNowString()
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	_, deleteErr := stmt.Exec(user.Id)
	if deleteErr != nil {
		return mysql_utils.ParseError(deleteErr)
	}
	return nil
}
