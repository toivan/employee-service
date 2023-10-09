package repository

import (
	"database/sql"
	"employee-service/internal/app/database"
	"employee-service/internal/app/model"
)

type EmployeeRepository struct {
	DB *sql.DB
}

func NewEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{
		DB: database.DB,
	}
}

func (er *EmployeeRepository) CreateEmployee(employee model.Employee) error {
	_, err := database.DB.Exec(
		"INSERT INTO employees (full_name, phone, gender, age, email, address) VALUES (?, ?, ?, ?, ?, ?)",
		employee.FullName, employee.Phone, employee.Gender, employee.Age, employee.Email, employee.Address,
	)
	return err
}

func (er *EmployeeRepository) DeleteEmployee(id int) error {
	_, err := database.DB.Exec("DELETE FROM employees WHERE id = ?", id)
	return err
}

func (er *EmployeeRepository) GetEmployeeVacationDays(id int) (int, error) {
	return 0, nil
}

func (er *EmployeeRepository) FindEmployeeByName(name string) ([]model.Employee, error) {
	rows, err := database.DB.Query("SELECT * FROM employees WHERE full_name LIKE ?", "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []model.Employee
	for rows.Next() {
		var e model.Employee
		if err := rows.Scan(&e.ID, &e.FullName, &e.Phone, &e.Gender, &e.Age, &e.Email, &e.Address); err != nil {
			return nil, err
		}
		employees = append(employees, e)
	}

	return employees, nil
}
