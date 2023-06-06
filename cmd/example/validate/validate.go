package validate

// #[Validate(email), ValidateMsg(zh-CN = "邮箱格式不正确", en = "Email format is incorrect")]
type SignupData struct {
	// #[validate(email)]
	// #[validate_msg(zh-CN = "邮箱格式不正确", en = "Email format is incorrect")]
	Mail string
	// #[validate(phone)]
	Phone string
	// #[validate(url)]
	Site string
}
