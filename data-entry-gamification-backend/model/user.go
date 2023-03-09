package model

// We keep the JSON notation bcause on Register we use this struct to pass the pwd
type User struct {
	ID        int64  `json:"ID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

type UserDTO struct {
	ID        int64  `json:"ID"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type UserDAO struct {
	ID        int64  
	FirstName string 
	LastName  string 
	Password  string 
	Email     string 
}
