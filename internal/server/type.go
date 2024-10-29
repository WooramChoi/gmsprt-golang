package server

type Config struct {
	Server struct {
		Port int
	}
	Database struct {
		Type     string
		Host     string
		Port     int
		Dbname   string
		Username string
		Password string
	}
}
