package models

func usersKey() string {
	return "linx:users"
}

func UserExists(id string) bool {

	return false
}

func UserDBId(id string) int64 {

	return 0
}

func (m *User) SaveUser() {
}