package route

const BasePath = "/api/v1"

const UserPath = BasePath + "/user"
const UserByUsername = UserPath + "/:username"
const UserLogin = UserPath + "/login"
const UserRegister = UserPath + "/register"
const UserForgotPassword = UserPath + "/forgot-password"
const UserVerifyOTP = UserPath + "/verify-otp"
const UserResetPassword = UserPath + "/reset-password"

// Forum
const ForumPath = BasePath + "/forums"
const ForumByID = ForumPath + "/:id"
const ForumUpdate = ForumPath + "/update/:id"
const ForumDelete = ForumPath + "/delete/:id"

// Forum Message
const ForumMessage = ForumPath + "/message"
const ForumMessageByID = ForumMessage + "/:id"
const ForumMessageUpdate = ForumMessage + "/update/:id"
const ForumMessageDelete = ForumMessage + "/delete/:id"

// Admin
const AdminPath = BasePath + "/admin"
const AdminLogin = AdminPath + "/login"

// const AdminGetData = AdminPath + "/data"
const AdminEdit = AdminPath + "/edit/:id"
const AdminDelete = AdminPath + "/delete"

// Impact Category Path
const ImpactCategoryPath = AdminPath + "/impact"
const ImpactCategoryByID = ImpactCategoryPath + "/:id"

// Product Path
const ProductPath = BasePath + "/products"
const ProductByID = ProductPath + "/:id"
const ProductByName = ProductPath + "?name=:name"

// Challenge Path
const ChallengePath = BasePath + "/challenges"
const ChallengeByID = ChallengePath + "/:id"
