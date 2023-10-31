package common

var JWT_SECRET_KEY = []byte("hehehehehehheh_moah_moah")

// regex
const EMAIL_REGEX = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

// ERORR TAG
const ROLE_USER_DENIED = "The role of the user don't have permissions"
const FULL_NAME_BLANK = "Full Name can't be blank"
const EMAILL_BLANK = "Email can't be blank"
const PASS_WORD_BLANK = "Password can't be blank"
const EMAILL_WRONG_REGEX = "Email wrong format"
const EMAILL_DUPLICATE = "Email duplicate"
const USER_ID_NOT_FOUND = "user id not exist"
const USER_ID_BLANK = "User can't be blank"

// ERORR TAG JWT
const TOKEN_NOT_APPROPRIATE = "Token not appropriate"
