package server

import "fmt"

// Common API types

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

type Client struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// Universal responses

type Id struct {
	Id string `json:"id"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type CountResponse struct {
	Count int64 `json:"count"`
}

// Particular Request/Response's

type SigninRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
