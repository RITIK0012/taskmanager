package repository


var users = map[string]string{}

func CreateUser(username,password string)bool{
	if _,exists:= users[username]; exists{
		return false
	}
	users[username] = password
	return true
}

func ValidateUser(username,password string) bool{
	pass,exists := users[username]
	return exists && pass==password
}