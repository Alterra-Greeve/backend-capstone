package constant

import "errors"

const Unauthorized = "Unauthorized"
const InternalServerError = "Internal Server Error"
const BadInput = "Format data not valid"

var BadRequest = errors.New("Format data not valid")
var ErrPageInvalid = errors.New("Page is invalid")

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
var ErrUpdateUser = errors.New("Failed to update user")
var ErrEmptyUpdate = errors.New("One or more fields for update cannot be empty")
var ErrEmailUsernameAlreadyExist = errors.New("Email or Username already exist")
var ErrEmptyEmail = errors.New("Email cannot be empty")
var ErrEmailNotFound = errors.New("Email not found")
var ErrForgotPassword = errors.New("Failed to forgot password")
var ErrOTPNotValid = errors.New("OTP not valid")
var ErrOTPExpired = errors.New("OTP expired")
var ErrEmptyOTP = errors.New("OTP cannot be empty")
var ErrResetPassword = errors.New("Failed to reset password")
var ErrDeleteUser = errors.New("Failed to delete user")
var ErrEmptyResetPassword = errors.New("Email, password and confirmation password cannot be empty")
var ErrPasswordNotMatch = errors.New("Password not match")
var ErrInvalidEmail = errors.New("Email is not valid")

var ErrGenerateJWT = errors.New("failed to generate jwt token")

var ErrValidateJWT = errors.New("failed to validate jwt token")

var ErrHashPassword = errors.New("failed to hash password")

// Forum
var ErrCreateForum = errors.New("Failed to create forum")
var ErrGetForum = errors.New("Failed to get forum")
var ErrGetForumByID = errors.New("Failed to get forum by ID")
var ErrEditForum = errors.New("Field Description Cannot be Empty")
var ErrForumNotFound = errors.New("error, id Forum Not Found")

var ErrGetMessage = errors.New("Failed to get Message")
var ErrGetMessageByID = errors.New("Failed to get Message by ID")
var ErrMessgaeNotFound = errors.New("error, id Message Not Found")
var ErrEditMessage = errors.New("Field Message Cannot be Empty")
var ErrCreateMessage = errors.New("Failed to create Message")
var ErrDeleteForum = errors.New("Failed to delete forum")
var ErrUpdateMessage = errors.New("Failed to update Message")
var ErrDeleteMessage = errors.New("Failed to delete Message")
var UnatuhorizeForumAndMessage = errors.New("you are not authorized to update this forum")

// Admin Errors
var ErrGetDataAdmin = errors.New("Failed to get data admin")
var ErrEmptyEmailandPasswordAdmin = errors.New("Email and Password cannot be empty")
var ErrNotFoundEmailAdmin = errors.New("Email is incorrect")
var ErrEditAdmin = errors.New("Field Cannot be Empty")
var ErrAdminEmailUsernameAlreadyExist = errors.New("Email or Username already taken by another admin")

// Impact Category Errors
var ErrImpactCategoryNotFound = errors.New("Impact Category Not Found")
var ErrCreateImpactCategory = errors.New("Failed to create impact category")
var ErrUpdateImpactCategory = errors.New("Failed to update impact category")
var ErrDeleteImpactCategory = errors.New("Failed to delete impact category")
var ErrImpactCategoryField = errors.New("Field cannot be empty")
var ErrImpactCategoryFieldUpdate = errors.New("One or more fields for update cannot be empty")

// Product Errors
var ErrProductNotFound = errors.New("Product Not Found")
var ErrGetProduct = errors.New("Failed to get product")
var ErrProductEmpty = errors.New("There is no product")
var ErrCreateProduct = errors.New("Failed to create product")
var ErrUpdateProduct = errors.New("Failed to update product")
var ErrDeleteProduct = errors.New("Failed to delete product")
var ErrProductField = errors.New("Field cannot be empty")
var ErrProductFieldUpdate = errors.New("One or more fields for update cannot be empty")
var ErrProductNameEmpty = errors.New("Name cannot be empty")
var ErrProductDescriptionEmpty = errors.New("Description cannot be empty")
var ErrProductCoinEmpty = errors.New("Coin cannot be empty")
var ErrProductPriceEmpty = errors.New("Price cannot be empty")
var ErrProductImagesEmpty = errors.New("Images cannot be empty")
var ErrProductImpactCategoriesEmpty = errors.New("Impact Categories cannot be empty")
var ErrProductIDEmpty = errors.New("ID cannot be empty")
var ErrProductUpdateEmpty = errors.New("One or more fields for update cannot be empty")
var ErrProductDelete = errors.New("Failed to delete product")

// Challenges Errors
var ErrUserAlreadyParticipate = errors.New("User already participate")
var ErrUserNotParticipate = errors.New("User not participate")
var ErrCreateChallenge = errors.New("Failed to create challenge")
var ErrGetChallenge = errors.New("Failed to get challenge")
var ErrGetChallengeByID = errors.New("Failed to get challenge by ID")
var ErrEditChallenge = errors.New("Field Description Cannot be Empty")
var ErrChallengeNotFound = errors.New("error, id Challenge Not Found")
var ErrUpdateChallenge = errors.New("Failed to update challenge")
var ErrDeleteChallenge = errors.New("Failed to delete challenge")
var ErrChallengeField = errors.New("Field cannot be empty")
var ErrChallengeFieldUpdate = errors.New("One or more fields for update cannot be empty")
var ErrChallengeFieldSwipe = errors.New("Field cannot be empty")
var ErrChallengeType = errors.New("Challenge type must be accept or decline")
var ErrChallengeFieldCreate = errors.New("Field cannot be empty")

