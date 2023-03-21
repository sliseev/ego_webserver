package server

import "fmt"

///////////////////
// Common API types

// @Description	Driver data
type Driver struct {
	Name    string `json:"name"`
	License string `json:"license"`
}

func (d Driver) Validate() error {
	if len(d.Name) == 0 {
		return fmt.Errorf("Name field must be NOT NULL")
	}
	if len(d.License) == 0 {
		return fmt.Errorf("License field must be NOT NULL")
	}
	return nil
}

// @Description Service client data
type Client struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//////////////////////
// Universal responses

// @Description	Success response with object ID
type Id struct {
	Id string `json:"id"`
}

// @Description	Error response with details
type ErrorResponse struct {
	Error string `json:"error"`
}

// @Description	Success response with objects count
type CountResponse struct {
	Count int64 `json:"count"`
}

////////////////////////////////
// Particular Request/Response's

// @Description	Signin Request data
type SigninRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Description	Login request data
type LoginRequest struct {
	Email    string `json:"login"`
	Password string `json:"password"`
}

// @Description	Success login response with token
type LoginResponse struct {
	Token string `json:"token"`
}

// @Description Test content generator data
type GeneratorRequest struct {
	Count   int  `json:"count"`
	Cleanup bool `json:"cleanup"`
}
