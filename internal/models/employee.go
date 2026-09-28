package models

type Employee struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

type EmployeeStore interface {
	Create(emp *Employee) error
	List() ([]Employee, error)
	Get(id int) (Employee, error)
	Update(id int, emp *Employee) error
	Delete(id int) error
}
