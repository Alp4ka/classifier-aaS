package db

type Config struct {
	MaxOpenConns          int
	MaxIdleConns          int
	DSN                   string
	MaxConnectionAttempts int
}
