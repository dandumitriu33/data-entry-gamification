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

// Maps the UserDTO struct data to User struct data to use internally.
// Should be used on operation when other Roles pass user data not the user.
// The user on password update would use the User struct.
func mapFromUserDTOToModel(userDTO UserDTO, user *User) {
	user.ID = userDTO.ID
	user.FirstName = userDTO.FirstName
	user.LastName = userDTO.LastName
	user.Email = userDTO.Email
}

// Maps the User struct data to UserDTO struct data to use externally.
// Should be used on operation when other Roles pass user data not the user.
// The user on password update would use the User struct.
func mapFromUserModelToDTO(user User, userDTO *UserDTO) {
	userDTO.ID = user.ID
	userDTO.FirstName = user.FirstName
	userDTO.LastName = user.LastName
	userDTO.Email = user.Email
}

// Maps the User with DB data to the User struct
func mapFromUserDAOToModel(userDAO UserDAO, user *User) {
	user.ID = userDAO.ID
	user.Password = userDAO.Password
	user.FirstName = userDAO.FirstName
	user.LastName = userDAO.LastName
	user.Email = userDAO.Email
}

// Maps the UserDAO with User data to insert in the DB 
func mapFromUserModelToDao(user User, userDAO *UserDAO) {
	userDAO.ID = user.ID
	userDAO.Password = user.Password
	userDAO.FirstName = user.FirstName
	userDAO.LastName = user.LastName
	userDAO.Email = user.Email
}