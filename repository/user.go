package repository

import (
	"log"

	"github.com/klyngen/flightlogger/common"
	"github.com/pkg/errors"
)

// CreateUser creates an unique user
func (f *MySQLRepository) CreateUser(user *common.User) error {
	stmt, err := f.db.Prepare("INSERT INTO User (ID, Firstname, Lastname, Email, PasswordHash, PasswordSalt)  VALUES (UUID(), ?, ?, ?, ?, ?)")
	defer stmt.Close()

	if err != nil {
		return errors.Wrap(err, "Could not understand the statement")
	}

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.PasswordHash, user.PasswordSalt)

	if err != nil {
		return errors.Wrap(err, "Unable to insert the user")
	}

	// Get the ID from the newly inserted row
	f.db.QueryRow("SELECT ID FROM User WHERE Email = ? LIMIT 1", user.Email).Scan(&user.ID)
	return nil

}

// GetAllUsers gets all the users implements paging
func (f *MySQLRepository) GetAllUsers(limit int, page int) ([]common.User, error) {
	stmt, err := f.db.Prepare("SELECT ID, Firstname, Lastname, Email, PasswordHash, PasswordSalt FROM User LIMIT ?,?")

	defer stmt.Close()
	if err != nil {
		return nil, errors.Wrap(err, "Unable to prepare statement")
	}

	result, err := stmt.Query(page*limit, limit)
	defer result.Close()

	var users []common.User

	// Loop the result
	for result.Next() {
		user := common.User{}

		err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.PasswordSalt)

		if err != nil {
			log.Printf("Skipping marshaling of a row due to the following error: \n %v", err)
		}

		users = append(users, user)
	}

	return users, errors.Wrap(err, "Unable to fetch the requested users")
}

// GetUser gets a singular user
func (f *MySQLRepository) GetUser(ID string, user *common.User) error {
	stmt, err := f.db.Prepare("SELECT ID, Firstname, Lastname, Email, PasswordHash, PasswordSalt FROM User where ID = ? LIMIT 1")

	defer stmt.Close()

	if err != nil {
		return errors.Wrap(err, "Could not understand the statement")
	}

	// Instantiate the object
	if user == nil {
		user = &common.User{}
	}
	// Map the rows if possible
	err = stmt.QueryRow(ID).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.PasswordSalt)

	return errors.Wrap(err, "Could not feth the given user")
}

// UpdateUser does just that
func (f *MySQLRepository) UpdateUser(ID string, user *common.User) error {
	stmt, err := f.db.Prepare("UPDATE User SET Firstname = ?, Lastname = ?, Email = ?, PasswordHash = ?, PasswordSalt = ? WHERE ID = ? LIMIT 1")

	defer stmt.Close()

	if err != nil {
		return errors.Wrap(err, "Could not understand the statement")
	}

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.PasswordHash, user.PasswordSalt, ID)

	return errors.Wrap(err, "Unable to update the user due to unexpected error")
}

// DeleteUser should be hardDelete
func (f *MySQLRepository) DeleteUser(ID string) error {
	stmt, err := f.db.Prepare("DELETE FROM User where ID = ? LIMIT 1")
	defer stmt.Close()

	if err != nil {
		return errors.Wrap(err, "Could not understand the statement")
	}

	_, err = stmt.Exec(ID)

	return errors.Wrap(err, "Could not delete the user")
}

// GetUserByEmail is used in the authentication-process
func (f *MySQLRepository) GetUserByEmail(Email string, user *common.User) error {
	stmt, err := f.db.Prepare("SELECT ID, Firstname, Lastname, Email, PasswordHash, PasswordSalt FROM User where Email = ?")

	defer stmt.Close()

	if err != nil {
		return errors.Wrap(err, "Could not understand the statement")
	}

	err = stmt.QueryRow(Email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.PasswordHash, &user.PasswordSalt)

	return errors.Wrap(err, "Could not feth the given user")
}

// UserGroup CRUD and search
func (f *MySQLRepository) CreateUserGroup(userGroup *common.UserGroup, scopes []int) error {
	panic("not implemented")
}

func (f *MySQLRepository) UpdateUserGroup(groupID uint, userGroup *common.UserGroup, permissions []int) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetAllUserGroups(limit int, page int, userGroups []common.UserGroup) error {
	panic("not implemented")
}

func (f *MySQLRepository) GetUserGroup(ID uint, userGroup *common.UserGroup) error {
	panic("not implemented")
}

func (f *MySQLRepository) UserGroupSearchByName(name string, userGroups []common.UserGroup) error {
	panic("not implemented")
}

// TODO: We need to add permissions to the user somehow
