package main

import (
	"fmt"
	"github.com/iproduct/coursego/08-databases/entities"
	"github.com/iproduct/coursego/08-databases/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm_projects?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Company{})
	db.AutoMigrate(&entities.Project{})

	users := []entities.User{
		{FirstName: "Linus", LastName: "Torvalds", Email: "linus@linux.com", Username: "linus", Password: "linus",
			Active: true, Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		{FirstName: "James", LastName: "Gosling", Email: "gosling@java.com", Username: "james", Password: "james",
			Active: true, Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		{FirstName: "Rob", LastName: "Pike", Email: "pike@golang.com", Username: "rob", Password: "rob",
			Active: true, Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
		{FirstName: "Kamel", LastName: "Founadi", Email: "kamel@docker.com", Username: "kamel", Password: "kamel",
			Active: true, Model: gorm.Model{CreatedAt: time.Now(), UpdatedAt: time.Now()}},
	}

	result := db.Create(&users) // pass pointer of data to Create
	// batch size 100
	//db.CreateInBatches(users, 100)
	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("New users created with IDs: ")
	for _, user := range users {
		fmt.Printf("%d, ", user.ID)
	}
	fmt.Println()

	utils.PrintUsers(users)

	//Get all users
	result = db.Find(&users) 	// SELECT * FROM users;

	if result.Error != nil {
		log.Fatal(result.Error) // returns error
	}
	fmt.Printf("Number of users: %d\n", result.RowsAffected)// returns found records count, equals `len(users)`
	utils.PrintUsers(users)
}