package validation

import (
	"errors"
	"strings"

	"employeejwt/internal/constants"
	"employeejwt/internal/models"
	"employeejwt/internal/utils"
)

func ValidateRegisterRequest(req models.RegisterRequest) error {
	if utils.IsBlank(req.Name) || utils.IsBlank(req.Email) || utils.IsBlank(req.Password) {
		return errors.New(constants.NameEmailPasswordRequired)
	}

	req.Email = strings.TrimSpace(req.Email)
	if !strings.Contains(req.Email, "@") {
		return errors.New(constants.ErrInvalidEmailAddress)
	}

	return nil
}

func ValidateLoginRequest(req models.LoginRequest) error {
	if utils.IsBlank(req.Email) || utils.IsBlank(req.Password) {
		return errors.New(constants.EmailPasswordRequired)
	}
	return nil
}

func ValidateEmployee(emp models.Employee) error {
	if utils.IsBlank(emp.Name) || utils.IsBlank(emp.Email) || utils.IsBlank(emp.Department) {
		return errors.New(constants.NameEmailDepartmentRequired)
	}
	return nil
}
