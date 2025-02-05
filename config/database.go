package config

func GetDatabaseHost() string {
	return GetStringOrPanic("database.host")
}

func GetDatabaseUser() string {
	return GetStringOrPanic("database.user")
}

func GetDatabasePassword() string {
	return GetStringOrPanic("database.password")
}

func GetDatabaseName() string {
	return GetStringOrPanic("database.dbname")
}

func GetDatabasePort() int {
	return GetIntWithDefault("database.port", 5432)
}
