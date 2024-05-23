package route

const BasePath = "/api/v1"

const UserPath = BasePath + "/user"
const UserByUsername = UserPath + "/:username"
const UserLogin = UserPath + "/login"
const UserRegister = UserPath + "/register"
const UserForgotPassword = UserPath + "/forgot-password"
const UserVerifyOTP = UserPath + "/verify-otp"
const UserResetPassword = UserPath + "/reset-password"

// Admin
const AdminPath = BasePath + "/admin"
const AdminLogin = AdminPath + "/login"

// const AdminGetData = AdminPath + "/data"
const AdminEdit = AdminPath + "/edit/:id"
const AdminDelete = AdminPath + "/delete"
