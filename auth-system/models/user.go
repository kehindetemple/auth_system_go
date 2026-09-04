package models

type User struct {
    ID        string
    Email     string
    Password  string
    CreatedAt string
}

type LoginRequest struct {
	Email string
	Password string
}