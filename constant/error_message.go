package constant

import "errors"

const Unauthorized = "Unauthorized"
const InternalServerError = "Internal Server Error"

var ErrGenerateJWT = errors.New("failed to generate jwt token")
var ErrValidateJWT = errors.New("failed to validate jwt token")

var ErrHashPassword = errors.New("failed to hash password")
