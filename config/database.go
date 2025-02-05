package config

func GetDatabaseHost() string {
	return GetStringOrPanic("DATABASE_HOST")
}

func GetDatabaseUser() string {
	return GetStringOrPanic("DATABASE_USER")
}


func GetDatabasePassword() string {
	return GetStringOrPanic("DATABASE_PASSWORD")
}

func GetDatabaseName() string {
	return GetStringOrPanic("DATABASE_NAME")
}

func GetDatabasePort() int {
	return GetIntWithDefault("DATABASE_PORT", 5432)
}
