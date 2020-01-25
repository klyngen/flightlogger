package repository

import "github.com/klyngen/flightlogger/common"

// CreateRole Create a new Role
func (f *MySQLRepository) CreateRole(role *common.Role) error {
	stmt, err := f.db.Prepare(`
		INSERT INTO Flightlog.Role
			(ResourceName, ResourceDescription)
			VALUES(?, ?)
	`)

	if err != nil {
		return BadSqlError.NewFromException(err, "INSERT", "Role")
	}

	defer stmt.Close()

	res, err := stmt.Exec(role.Name, role.Description)

	if err != nil {
		return StatementExecutionError.NewFromException(err, "INSERT", "Role")
	}

	id, err := res.LastInsertId()

	if err != nil {
		return LastInsertionIDExtractionError.NewFromException(err, "INSERT", "Role")
	}

	role.ID = uint(id)

	return nil
}

// DeleteRole deletes a role
func (f *MySQLRepository) DeleteRole(ID uint) error {
	stmt, err := f.db.Prepare(`DELETE FROM Flightlog.Role
		WHERE Id=? `)

	if err != nil {
		return BadSqlError.NewFromException(err, "Delete", "Role")
	}

	defer stmt.Close()

	_, err = stmt.Exec(ID)

	if err != nil {
		return StatementExecutionError.NewFromException(err, "INSERT", "Role")
	}

	return nil
}

// GetUserRole gets a user role
func (f *MySQLRepository) GetUserRole(userID string, role *common.Role) error {
	stmt, err := f.db.Prepare(`SELECT 
	r.Id, 
	r.ResourceName, 
	r.ResourceDescription 
	from Flightlog.User u 
		INNER JOIN Flightlog.Role r on r.Id = u.RoleId
	where u.Id = ? `)

	if err != nil {
		return BadSqlError.NewFromException(err, "Delete", "Role")
	}

	defer stmt.Close()

	row := stmt.QueryRow(userID)

	return mapRole(role, row)

}

func mapRole(role *common.Role, scanner rowScanner) error {
	if role == nil {
		role = &common.Role{}
	}

	if err := scanner.Scan(&role.ID, &role.Name, &role.Description); err != nil {
		return SerilizationError.NewFromException(err, "SELECT", "Role")
	}

	return nil
}
