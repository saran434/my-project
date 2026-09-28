package validation

import (
	"testing"

	"employeejwt/internal/models"
)

func TestValidateRegisterRequest(t *testing.T) {
	req := models.RegisterRequest{
		Name:     "Jane Doe",
		Email:    "  JANE@EXAMPLE.COM  ",
		Password: "secret123",
	}

	if err := ValidateRegisterRequest(req); err != nil {
		t.Fatalf("expected valid register request, got error: %v", err)
	}
}

func TestValidateRegisterRequestRejectsMissingFields(t *testing.T) {
	req := models.RegisterRequest{Name: "", Email: "", Password: ""}

	if err := ValidateRegisterRequest(req); err == nil {
		t.Fatal("expected validation error for missing fields")
	}
}
