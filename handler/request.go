package handler

import (
	"fmt"
)

func errParamRequired(param string) error {
	return fmt.Errorf("%s is required", param)
}

// CreateOpeningHandler is the handler for creating an opening

type CreateOpeningHandlerRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *CreateOpeningHandlerRequest) Validate() error {
	if r.Role == "" {
		return errParamRequired("role")
	}

	if r.Company == "" {
		return errParamRequired("company")
	}

	if r.Location == "" {
		return errParamRequired("location")
	}

	if r.Link == "" {
		return errParamRequired("link")
	}

	if r.Remote == nil {
		return errParamRequired("remote")
	}

	if r.Salary <= 0 {
		return errParamRequired("salary")
	}

	return nil
}

type UpdateOpeningHandlerRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateOpeningHandlerRequest) Validate() error {
	// if any of fields provided its valid

	if r.Role == "" || r.Company == "" || r.Location == "" || r.Link == "" || r.Remote == nil || r.Salary == 0 {
		return nil
	}

	return fmt.Errorf("at least one field must be provided")
}
