package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/robfig/cron"

	"github.com/labstack/echo"
	"github.com/spf13/viper"

	"github.com/reinhardjs/sayakaya/domain"
	_userHttpDelivery "github.com/reinhardjs/sayakaya/user/delivery/http"
	_userRepo "github.com/reinhardjs/sayakaya/user/repository/mysql"
	_userUcase "github.com/reinhardjs/sayakaya/user/usecase"

	_bulkemailsendRepo "github.com/reinhardjs/sayakaya/bulkemailsend/repository/smtp"
	_bulkemailsendUcase "github.com/reinhardjs/sayakaya/bulkemailsend/usecase"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	username := viper.GetString(`smtp.username`)
	password := viper.GetString(`smtp.password`)
	host := viper.GetString(`smtp.host`)
	port := viper.GetString(`smtp.port`)
	sender := viper.GetString(`smtp.sender`)

	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
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
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userRepository := _userRepo.NewMysqlUserRepository(dbConn)
	userUsecase := _userUcase.NewUserUsecase(userRepository, timeoutContext)
	bulkEmailSendRepository := _bulkemailsendRepo.NewBulkEmailSendRepository(username, password, host, port, sender)
	bulkEmailSendUsecase := _bulkemailsendUcase.NewBulkEmailSendUsecase(bulkEmailSendRepository)

	// ---------------------- CRON START ---------------------- //
	c := cron.New()
	// Schedule the cron job to run every day at 1:00 using 0 1 * * *
	// You can use */7 * * * * * to run every 7 seconds
	err = c.AddFunc("*/7 * * * * *", func() {
		sendBirthdayWishes(userUsecase, bulkEmailSendUsecase)
	})
	if err != nil {
		fmt.Println("Error adding cron job:", err)
		return
	}
	c.Start()
	// ---------------------- CRON END -------------------------- //

	// ---------------------- HTTP HANDLER START ---------------------- //
	_userHttpDelivery.NewUserHandler(e, userUsecase)
	log.Fatal(e.Start(viper.GetString("server.address"))) //nolint
	// ---------------------- HTTP HANDLER END ------------------------ //
}

func sendBirthdayWishes(userUsecase domain.UserUsecase, bulkEmailSendUsecase domain.BulkEmailSendUsecase) {
	// Get the current time
	now := time.Now()

	// Print the current time in hours and seconds
	fmt.Printf("%02d:%02d:%02d - Start sending birthday wishes to: \n", now.Hour(), now.Minute(), now.Second())

	// Connect to the database and retrieve users with birthdays today.
	// Send birthday wishes to each user.
	// Parse the string into a time.Time object
	date, err := time.Parse("2006-01-02", time.Now().Format("2006-01-02"))
	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}

	// run in main context
	ctx, _ := context.WithCancel(context.Background())

	// Example code assumes a function getTodayBirthdaysFromDB that retrieves users with birthdays today.
	users, err := userUsecase.FetchByBirthDay(ctx, date)
	if err != nil {
		fmt.Println("Error retrieving birthdays:", err)
		return
	}

	recipientEmails := make([]string, len(users))
	for i, user := range users {
		recipientEmails[i] = user.Email
	}

	fmt.Println(recipientEmails)

	// Example code assumes a function sendBirthdayMessage that sends birthday wishes to a user.
	bulkEmailSendUsecase.BulkSend(&domain.BulkEmailSend{Recipients: recipientEmails, Subject: "Happy Birthday", Message: "Hii, happy birthday"})
}
