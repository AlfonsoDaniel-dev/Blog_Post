package config

import (
	"fmt"
	"os"
)

var envDbVarsRequired = []string{"DB_USER", "DB_PASSWORD", "DB_PORT", "DB_HOST", "DB_NAME"}

func CreateStrConn() string {
	vars := getEnvDbVars()

	user, password, port, host, dbName := vars[0], vars[1], vars[2], vars[3], vars[4]

	fmt.Println(vars[2])

	connStr := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, dbName)

	return connStr
}

func getEnvDbVars() []string {
	values := make([]string, len(envDbVarsRequired))
	for i, envVar := range envDbVarsRequired {
		value := os.Getenv(envVar)
		if value == "" {
			fmt.Printf("Variable de entorno %s no está definida\n", envVar)
			return []string{}
		}
		values[i] = value
	}
	return values
}
