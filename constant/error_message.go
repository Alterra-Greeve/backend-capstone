package constant

import "errors"

const Unauthorized = "Unauthorized"
const InternalServerError = "Internal Server Error"

// Users Errors
var ErrEmptyLogin = errors.New("Email or Password cannot be empty")
var UserNotFound = errors.New("User not found")
var ErrLoginIncorrectPassword = errors.New("Incorrect password")
var ErrEmptyEmailRegister = errors.New("Email cannot be empty")
var ErrEmptyPasswordRegister = errors.New("Password cannot be empty")
var ErrEmptyAddressRegister = errors.New("Address cannot be empty")
var ErrEmptyNameRegister = errors.New("Name cannot be empty")
var ErrEmptyGenderRegister = errors.New("Gender cannot be empty")
var ErrEmailAlreadyExist = errors.New("Email already exist")
var ErrUsernameAlreadyExist = errors.New("Username already exist")
var ErrEmptyPhoneRegister = errors.New("Phone cannot be empty")
var ErrRegister = errors.New("Failed to register user")

var ErrGenerateJWT = errors.New("failed to generate jwt token")
var ErrValidateJWT = errors.New("failed to validate jwt token")

var ErrHashPassword = errors.New("failed to hash password")
