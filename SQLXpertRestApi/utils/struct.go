package utils

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Port     int
	DbName   string
}
