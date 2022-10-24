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
		//Book
		&Book{},
		&BookType{},
		&Shelf{},
		//User
		&Province{},
		&MemberClass{},
		&User{},
		&Role{},
		//Borrow
		&Borrow{},
		//reseachroom
		&RoomType{},
		&Equipment{},
		&ResearchRoom{},
		&TimeRoom{},
		&AddOn{},
		&ResearchRoomReservationRecord{},
		//Bill
		&Bill{},
		//Com_reser
		&Computer_os{},
		&Computer_reservation{},
		&Computer{},
		&Time_com{},
		//Problem
		&Place_Class{},
		&Relation{},
		&Toilet{},
		&ReadingZone{},
		&ProblemReport{},
	)

	db = database
	password1, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	//password2, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password3, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
	password4, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)
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

	employee := Role{
		Name:       "Employee",
		BorrowDay:  5,
		BookRoomHR: 6,
		BookComHR:  6,
	}
	db.Model(&Role{}).Create(&employee)

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
	//province
	korat := Province{
		Name: "Nakhon Ratchasima",
	}
	db.Model(&Province{}).Create(&korat)

	chon := Province{
		Name: "Chonburi",
	}
	db.Model(&Province{}).Create(&chon)

	bangkok := Province{
		Name: "Bangkok",
	}
	db.Model(&Province{}).Create(&bangkok)

	//member
	classic := MemberClass{
		Name:     "classic",
		Discount: 0,
	}
	db.Model(&MemberClass{}).Create(&classic)

	silver := MemberClass{
		Name:     "silver",
		Discount: 5,
	}
	db.Model(&MemberClass{}).Create(&silver)

	gold := MemberClass{
		Name:     "gold",
		Discount: 10,
	}
	db.Model(&MemberClass{}).Create(&gold)

	plat := MemberClass{
		Name:     "platinum",
		Discount: 20,
	}
	db.Model(&MemberClass{}).Create(&plat)

	//user
	user1 := User{
		Pin:       "B6111111",
		FirstName: "preecha",
		LastName:  "anpa",
		Civ:       "1111111111111",
		Phone:     "0811111111",
		Email:     "preechapat@mail.com",
		Password:  string(password3),
		Address:   "ถนน a อำเภอ v",
		//FK
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	}
	db.Model(&User{}).Create(&user1)

	user2 := User{
		Pin:       "E123456",
		FirstName: "Sirinya",
		LastName:  "kot",
		Civ:       "1234567890123",
		Phone:     "0899999999",
		Email:     "sirinya@mail.com",
		Password:  string(password1),
		Address:   "ถนน c อำเภอ z",
		//FK
		Role:        employee,
		Province:    bangkok,
		MemberClass: plat,
	}
	db.Model(&User{}).Create(&user2)

	db.Model(&User{}).Create(&User{
		Pin:       "T654321",
		FirstName: "Wichai",
		LastName:  "Micro",
		Civ:       "3210987654321",
		Phone:     "0823456789",
		Email:     "wichai@mail.com",
		Password:  string(password3),
		Address:   "ถนน c อำเภอ z",
		//FK
		Role:        teacher,
		Province:    bangkok,
		MemberClass: plat,
	})

	db.Model(&User{}).Create(&User{
		Pin:       "B6222222",
		FirstName: "kawin",
		LastName:  "anpa",
		Civ:       "22222222222222",
		Phone:     "0811111111",
		Email:     "kawin@mail.com",
		Password:  string(password4),
		Address:   "ถนน a อำเภอ v",
		//FK
		Role:        student,
		Province:    korat,
		MemberClass: classic,
	})
	var preecha User
	db.Raw("SELECT * FROM users WHERE email = ?", "preechapat@mail.com").Scan(&preecha)

	var kawin User
	db.Raw("SELECT * FROM users WHERE email = ?", "kawin@mail.com").Scan(&kawin)
	//Book
	db.Model(&Book{}).Create(&Book{
		Name:     "Python 1",
		User:     user2,
		Booktype: BT1,
		Shelf:    S2,
		Role:     student,
		Author:   "Sirin",
		Page:     500,
		Quantity: 20,
		Price:    300.0,
		Date:     time.Now(),
	})
	db.Model(&Book{}).Create(&Book{
		Name:     "Java",
		User:     user2,
		Booktype: BT1,
		Shelf:    S2,
		Role:     teacher,
		Author:   "AJ",
		Page:     350,
		Quantity: 10,
		Price:    200.5,
		Date:     time.Now(),
	})

	//Equipment data
	monitor := Equipment{
		Name: "จอ monitor สำหรับการนำเสนอ",
	}
	db.Model(&Equipment{}).Create(&monitor)

	printer := Equipment{
		Name: "เครื่องปริ้นท์",
	}
	db.Model(&Equipment{}).Create(&printer)

	printerMoniter := Equipment{
		Name: "เครื่องปริ้นท์ + จอ monitor สำหรับการนำเสนอ",
	}
	db.Model(&Equipment{}).Create(&printerMoniter)

	//Room_type data
	single_room := RoomType{
		Type: "ห้องเดี่ยว",
	}
	db.Model(&RoomType{}).Create(&single_room)

	group_room := RoomType{
		Type: "ห้องกลุ่ม",
	}
	db.Model(&RoomType{}).Create(&group_room)

	tutor_room := RoomType{
		Type: "ห้องสำหรับติว",
	}
	db.Model(&RoomType{}).Create(&tutor_room)

	//Research room
	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		//Name:      "RR001",
		RoomType:  group_room,
		Equipment: monitor,
	})

	db.Model(&ResearchRoom{}).Create(&ResearchRoom{
		//Name:      "RR002",
		RoomType:  group_room,
		Equipment: printerMoniter,
	})
	//ดึง Data ของ User มาเก็บไว้ในตัวแปรก่อน
	// var preecha User
	// db.Raw("SELECT * FROM users WHERE email = ?", "preechapat@mail.com").Scan(&preecha)

	// var kawin User
	// db.Raw("SELECT * FROM users WHERE email = ?", "kawin@mail.com").Scan(&kawin)

	//ดึง Data ของ researchroom มาเก็บไว้ในตัวแปรก่อน
	var Room1 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE id = ?", 1).Scan(&Room1)
	var Room2 ResearchRoom
	db.Raw("SELECT * FROM research_rooms WHERE id = ?", 2).Scan(&Room2)

	//Addon data
	powerPlug := AddOn{
		Name: "ปลั๊กพ่วง",
	}
	db.Model(&AddOn{}).Create(&powerPlug)

	Adapter := AddOn{
		Name: "สายชาร์จ",
	}
	db.Model(&AddOn{}).Create(&Adapter)

	Pillow := AddOn{
		Name: "หมอน",
	}
	db.Model(&AddOn{}).Create(&Pillow)

	powerPlugAndAdapter := AddOn{
		Name: "ปลั๊กพ่วง + สายชาร์จ",
	}
	db.Model(&AddOn{}).Create(&powerPlugAndAdapter)

	adapterAndPillow := AddOn{
		Name: "สายชาร์จ + หมอน",
	}
	db.Model(&AddOn{}).Create(&adapterAndPillow)

	powerPlugAndAdapterAndPillow := AddOn{
		Name: "ปลั๊กพ่วง + สายชาร์จ + หมอน",
	}
	db.Model(&AddOn{}).Create(&powerPlugAndAdapterAndPillow)

	//Time data
	timeMorning := TimeRoom{
		Period: "08:00 - 12:00",
	}
	db.Model(&TimeRoom{}).Create(&timeMorning)

	timeAfternoon := TimeRoom{
		Period: "13:00 - 17:00",
	}
	db.Model(&TimeRoom{}).Create(&timeAfternoon)

	//RRRR 1
	db.Model(&ResearchRoomReservationRecord{}).Create(&ResearchRoomReservationRecord{
		ResearchRoom: Room1,
		User:         preecha,
		AddOn:        powerPlugAndAdapter,
		BookDate:     time.Now(),
		TimeRoom:     timeMorning,
	})

	//RRRR 2
	db.Model(&ResearchRoomReservationRecord{}).Create(&ResearchRoomReservationRecord{
		ResearchRoom: Room2,
		User:         kawin,
		AddOn:        powerPlugAndAdapterAndPillow,
		BookDate:     time.Now(),
		TimeRoom:     timeAfternoon,
	})

	//Computer_os data
	comp_os_name1 := Computer_os{
		Name: "Windows",
	}
	db.Model(&Computer_os{}).Create(&comp_os_name1)

	//Computer data
	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W01A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W02A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W03A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W04A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	db.Model(&Computer{}).Create(&Computer{
		Comp_name:   "W05A",
		Comp_room:   "ROOM A",
		Computer_os: comp_os_name1,
	})

	//TIME data
	time_p1 := Time_com{
		Time_com_period: "08:00 - 09:00",
	}
	db.Model(&Time_com{}).Create(&time_p1)

	time_p2 := Time_com{
		Time_com_period: "09:00 - 10:00",
	}
	db.Model(&Time_com{}).Create(&time_p2)

	time_p3 := Time_com{
		Time_com_period: "10:00 - 11:00",
	}
	db.Model(&Time_com{}).Create(&time_p3)

	time_p4 := Time_com{
		Time_com_period: "11:00 - 12:00",
	}
	db.Model(&Time_com{}).Create(&time_p4)

	time_p5 := Time_com{
		Time_com_period: "12:00 - 13:00",
	}
	db.Model(&Time_com{}).Create(&time_p5)

	time_p6 := Time_com{
		Time_com_period: "13:00 - 14:00",
	}
	db.Model(&Time_com{}).Create(&time_p6)

	time_p7 := Time_com{
		Time_com_period: "14:00 - 15:00",
	}
	db.Model(&Time_com{}).Create(&time_p7)

	time_p8 := Time_com{
		Time_com_period: "15:00 - 16:00",
	}
	db.Model(&Time_com{}).Create(&time_p8)

	//ดึง Data ของ User มาเก็บไว้ในตัวแปรก่อน


	//ดึง Data ของ COMPUTER มาเก็บไว้ในตัวแปรก่อน
	// cn = comp_name ที่มาจาก Comp_name ใน Entity Computer
	var cn1 Computer
	db.Raw("SELECT * FROM computers WHERE Comp_name = ?", "W01A").Scan(&cn1)
	var cn2 Computer
	db.Raw("SELECT * FROM computers WHERE Comp_name = ?", "W02A").Scan(&cn2)
	var cn3 Computer
	db.Raw("SELECT * FROM computers WHERE Comp_name = ?", "W03A").Scan(&cn3)
	var cn4 Computer
	db.Raw("SELECT * FROM computers WHERE Comp_name = ?", "W04A").Scan(&cn4)
	var cn5 Computer
	db.Raw("SELECT * FROM computers WHERE Comp_name = ?", "W05A").Scan(&cn5)

	//Computer_reservation
	db.Model(&Computer_reservation{}).Create(&Computer_reservation{

		Date:     time.Now(),
		Computer: cn1,
		Time_com: time_p1,
		User:     preecha,
	})

	db.Model(&Computer_reservation{}).Create(&Computer_reservation{

		Date:     time.Now(),
		Computer: cn2,
		Time_com: time_p2,
		User:     kawin,
	})


}
