package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&Employee{},
		&Book{},
		&BookType{},
		&Shelf{},
		&Role{},
	)

	db = database
	// password1, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	// password2, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	//password3, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)

	// //Employee
	// db.Model(&Employee{}).Create(&Employee{
	// 	Name:     "Sirinya",
	// 	//Password: string(password1),
	// 	Password: "123456",
	// 	Email:    "sirin@gmail.com",
	// })

	// db.Model(&Employee{}).Create(&Employee{
	// 	Name:     "Attawit",
	// 	Password: "123456",
	// 	Email:    "attawit@example.com",
	// })password1, err := bcrypt.GenerateFromPassword([]byte("zaq1@wsX"), 14)
	// password1, err := bcrypt.GenerateFromPassword([]byte("zxvseta"), 14)
	// password2, err := bcrypt.GenerateFromPassword([]byte("1111111111111"), 14)

	//add example data
	//emp

	password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	db.Model(&Employee{}).Create(&Employee{
		Name:     "Sirinya",
		Email:    "sirinya@mail.com",
		Password: string(password),
	})

	db.Model(&Employee{}).Create(&Employee{
		Name:     "Attawit",
		Email:    "attawit@mail.com",
		Password: string(password),
	})

	var sirin Employee
	var attawit Employee
	db.Raw("SELECT * FROM employees WHERE email = ?", "sirinya@mail.com").Scan(&sirin)
	db.Raw("SELECT * FROM employees WHERE email = ?", "attawit@mail.com").Scan(&attawit)

	//Role
	student := Role{
		Name:       "Student",
		BorrowDay:  3,
		BookRoomHR: 3,
		BookComHR:  4,
	}

	db.Model(&Role{}).Create(&student)

	teacher := Role{
		Name:       "Teacher",
		BorrowDay:  7,
		BookRoomHR: 12,
		BookComHR:  12,
	}
	db.Model(&Role{}).Create(&teacher)

	//Shelf
	S1 := Shelf{
		Type:  "SCIENCE",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S1)
	S2 := Shelf{
		Type:  "ENGINEERING",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S2)
	S3 := Shelf{
		Type:  "ENVIRRONMENT",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S3)
	S4 := Shelf{
		Type:  "HISTORY",
		Floor: 1,
	}
	db.Model(&Shelf{}).Create(&S4)
	S5 := Shelf{
		Type:  "FICTION",
		Floor: 2,
	}
	db.Model(&Shelf{}).Create(&S5)
	S6 := Shelf{
		Type:  "FANTASY",
		Floor: 2,
	}
	db.Model(&Shelf{}).Create(&S6)
	S7 := Shelf{
		Type:  "HORROR",
		Floor: 2,
	}
	db.Model(&Shelf{}).Create(&S7)

	//Book Type
	BT1 := BookType{
		Type: "COMPUTER ENGINEERING",
	}
	db.Model(&BookType{}).Create(&BT1)

	BT2 := BookType{
		Type: "ELECTRIC ENGINEERING",
	}
	db.Model(&BookType{}).Create(&BT2)

	BT3 := BookType{
		Type: "SUPERHERO FANTASY",
	}
	db.Model(&BookType{}).Create(&BT3)

	BT4 := BookType{
		Type: "HORROR FICTION",
	}
	db.Model(&BookType{}).Create(&BT4)

	BT5 := BookType{
		Type: "DARK AND GRIMDARK FANTASY",
	}
	db.Model(&BookType{}).Create(&BT5)
	BT6 := BookType{
		Type: "CONTEMPORARY FANTASY",
	}
	db.Model(&BookType{}).Create(&BT6)

	//Book
	db.Model(&Book{}).Create(&Book{
		Name:     "Python 1",
		Employee: sirin,
		Booktype: BT1,
		Shelf:    S2,
		Role:     student,
		Author:   "Sirin",
		Page:     500,
		Quantity: 20,
		Price:    300,
		Date:     time.Now(),
	})
	db.Model(&Book{}).Create(&Book{
		Name:     "Java",
		Employee: attawit,
		Booktype: BT1,
		Shelf:    S2,
		Role:     teacher,
		Author:   "AJ",
		Page:     350,
		Quantity: 10,
		Price:    200,
		Date:     time.Now(),
	})
	//
	// === Query
	//

	// var target Employee
	// db.Model(&Employee{}).Find(&target, db.Where("email = ?", "Sirin@gmail.com"))

	// // var role ROLE
	// // db.Model(&ROLE{}).Find(&role, db.Where("name = ?", "Student"))

	// var book []*Book
	// db.Model(&Book{}).
	// 	Joins("Role").
	// 	Joins("Shelf").
	// 	Joins("Book_type").
	// 	Joins("Employee").
	// 	Find(&book, db.Where("id = ?", target.ID))

	// for _, wl := range book {
	// 	fmt.Printf("book: %v\n", wl.ID)
	// 	fmt.Printf("%v\n", wl.Role.NAME)
	// 	fmt.Printf("%v\n", wl.Shelf.Type)
	// 	fmt.Printf("%v\n", wl.Booktype.Type)
	// 	fmt.Printf("%v\n", wl.Emp.ID)
	// 	fmt.Println("====")
	// }

}
