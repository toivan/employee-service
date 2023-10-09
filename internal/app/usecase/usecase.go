package usecase

import (
	"employee-service/internal/app/model"
	"employee-service/internal/app/repository"
)

type EmployeeUsecase struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeUsecase(repo *repository.EmployeeRepository) *EmployeeUsecase {
	return &EmployeeUsecase{
		Repo: repo,
	}
}

func (eu *EmployeeUsecase) CreateEmployee(e model.Employee) error {
	return eu.Repo.CreateEmployee(e)
}

func (eu *EmployeeUsecase) DeleteEmployee(id int) error {
	return eu.Repo.DeleteEmployee(id)
}

func (eu *EmployeeUsecase) GetEmployeeVacationDays(id int) (int, error) {
	return eu.Repo.GetEmployeeVacationDays(id)
}

func (eu *EmployeeUsecase) FindEmployeeByName(name string) ([]model.Employee, error) {
	return eu.Repo.FindEmployeeByName(name)
}
