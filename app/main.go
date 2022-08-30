package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/labstack/echo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nelsonomoi/invoice/client/delivery"
	_clientRepo "github.com/nelsonomoi/invoice/client/repository/mysql"
	_clientUseCase "github.com/nelsonomoi/invoice/client/usecase"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil{
		panic(err)
	}

	if viper.GetBool(`debug`){
		log.Println("Server RUN on DEBUG mode")
	}
}

func main() {

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Africa/Nairobi")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)

	if err != nil {
		log.Fatal(err)
	}
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	
	e := echo.New()
	client := _clientRepo.NewMysqlClientRepository(dbConn)

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	au := _clientUseCase.NewClientUsecase(client, timeoutContext)
	delivery.NewClientHandler(e, au)

	log.Fatal(e.Start(viper.GetString("server.address")))
}