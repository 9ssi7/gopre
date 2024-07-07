package validation

var (
	userNameRegexp             = "^[a-zA-Z0-9]+([_.]?[a-zA-Z0-9])*$" // allow _ and . only between characters and numbers
	userNameAndCodeRegexp      = "^[a-zA-Z0-9]+([_.]?[a-zA-Z0-9])*(-|)?[0-9]{4}$"
	passwordRegexp             = "^[a-zA-Z0-9._].{8,16}$" // least one upper case, least one lower case, least one digit, least one special character, min 8 length, max 32 length
	localeRegexp               = "^[A-Za-z]{2}"           //"^[A-Za-z]{2,4}([_-][A-Za-z]{4})?([_-]([A-Za-z]{2}|[0-9]{3}))?$"
	slugRegexp                 = "^[a-z0-9]+(?:-[a-z0-9]+)*$"
	genderRegexp               = "^male$|^female$|^none$"
	emailRegexp                = "^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$"
	phoneWithCountryCodeRegexp = "^\\+[0-9]{1,3}[0-9]{3}[0-9]{3}[0-9]{4}$" // +1-123-456-7890, +90-123-456-7890, +123-123-456-7890
	locationRegexp             = "^[0-9]{1,3}\\.[0-9]{1,3}$"               // 123.123
)
