package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)
func OpenConnection() (*sql.DB, error) {
	viper.SetConfigFile("../envs/.env")
	viper.ReadInConfig()

	host := viper.Get("HOST").(string)
	dbPort := "5432"
	user := viper.Get("USER").(string)
	pass := viper.Get("PASS").(string)
	dbName := viper.Get("NAME").(string)

	url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	host, dbPort, user, pass, dbName)

	fmt.Println(url)

conn, err := sql.Open("postgres", url)
if err != nil {
	panic(err)
}

err = conn.Ping()

return conn, err
}