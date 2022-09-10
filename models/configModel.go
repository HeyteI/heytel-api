package models

type DatabaseConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	Database     string
	Ssl          string
	TimeZone     string
	MaxDbConns   int
	MaxIdleConns int
}
