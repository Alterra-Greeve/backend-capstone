package route

const BasePath = "/api/v1"

const UserPath = BasePath + "/user"
const UserByUsername = UserPath + "/:username"
const UserLogin = UserPath + "/login"
const UserRegister = UserPath + "/register"
const UserForgotPassword = UserPath + "/forgot-password"
const UserVerifyOTP = UserPath + "/verify-otp"
const UserResetPassword = UserPath + "/reset-password"

// Leaderboard
const Leaderboard = BasePath + "/leaderboard"

// Forum
const ForumPath = BasePath + "/forums"
const ForumByID = ForumPath + "/:id"
const GetForumByUserID = ForumPath + "/user"

// Forum Message
const ForumMessagePath = BasePath + "/message" + "/:id"
const ForumMessage = ForumPath + "/message"
const ForumMessageByID = ForumMessage + "/:id"

// Admin
const AdminPath = BasePath + "/admin"
const AdminLogin = AdminPath + "/login"

// const AdminGetData = AdminPath + "/data"
const AdminEdit = AdminPath + "/edit/:id"
const AdminDelete = AdminPath + "/delete"

// Admin Manage User
const AdminManageUserPath = AdminPath + "/users"
const AdminManageUserByID = AdminManageUserPath + "/:id"

// Impact Category Path
const ImpactCategoryPath = AdminPath + "/impact"
const ImpactCategoryByID = ImpactCategoryPath + "/:id"

// Voucher
const VoucherPath = AdminPath + "/voucher"
const VoucherByID = VoucherPath + "/:id"

// Product Path
const ProductPath = BasePath + "/products"
const ProductByID = ProductPath + "/:id"
const ProductByName = ProductPath + "?name=:name"

// Challenge Path
const ChallengePath = BasePath + "/challenges"
const AdminChallengePath = AdminPath + "/challenges"
const AdminChallengeByID = AdminChallengePath + "/:id"
const ChallengeParticipate = ChallengePath + "/participate"
const ChallengeByIDForUser = ChallengeParticipate + "/:id"
const ChallengeByID = ChallengePath + "/:id"
const ChallengeAdminEdit = AdminChallengePath + "/:id"
