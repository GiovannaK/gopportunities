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
