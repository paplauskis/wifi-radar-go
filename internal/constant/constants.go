package constant

type envKeys struct {
	ServerAddress      string
	CorsAllowedOrigin  string
	DBDriver           string
	DBConnectionString string
}

var EnvKeys = envKeys{
	ServerAddress:      "SERVER_ADDRESS",
	CorsAllowedOrigin:  "CORS_ALLOWED_ORIGIN",
	DBDriver:           "DB_DRIVER",
	DBConnectionString: "DB_CONNECTION_STRING",
}
