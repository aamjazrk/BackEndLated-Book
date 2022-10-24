package main

import (
	"github.com/aamjazrk/week5/controller"
	"github.com/aamjazrk/week5/entity"

	"github.com/aamjazrk/week5/middlewares"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	router := r.Group("/")
	{
		router.Use(middlewares.Authorizes())
		{
			// shelf Routes
			router.GET("/shelves", controller.ListShelf)
			router.GET("/shelve/:id", controller.GetShelf)
			//router.GET("/shelves/:type", controller.ListShelfByBookType)
			router.POST("/shelves", controller.CreateShelf)
			router.PATCH("/shelves", controller.UpdateShelf)
			router.DELETE("/shelves/:id", controller.DeleteShelf)

			// BOOK_TYPE Routes
			router.GET("/book_types", controller.ListBookType)
			router.GET("/book_type/:id", controller.GetBookType)
			router.POST("/book_types", controller.CreateBookType)
			router.PATCH("/book_types", controller.UpdateBookType)
			router.DELETE("/book_types/:id", controller.DeleteBookType)

			// Role Routes
			router.GET("/roles", controller.ListRole)
			router.GET("/role/:id", controller.GetRole)
			router.POST("/roles", controller.CreateRole)
			router.PATCH("/roles", controller.UpdateRole)
			router.DELETE("/roles/:id", controller.DeleteRole)

			router.GET("/users", controller.ListUser)
			router.GET("/user/:id", controller.GetUser)
			router.POST("/users", controller.CreateUser)
			router.PATCH("/users", controller.UpdateUser)
			router.DELETE("/users/:id", controller.DeleteUser)

			// book Routes
			router.GET("/books", controller.ListBook)
			router.GET("/book/:id", controller.GetBook)
			router.POST("/createbooks", controller.CreateBook)
			router.PATCH("/books", controller.UpdateBook)
			router.DELETE("/books/:id", controller.DeleteBook)

			//memberClass routes
			router.GET("/memberclasses", controller.ListMemberClass)
			router.GET("/memberclass/:id", controller.GetMemberClass)
			router.POST("/memberclasses", controller.CreateMemberClass)
			router.PATCH("/memberclasses", controller.UpdateMemberclass)
			router.DELETE("/memberclasses/:id", controller.DeleteMemberClass)

			//province routes
			router.GET("/provinces", controller.ListProvince)
			router.GET("/province/:id", controller.GetProvince)
			router.POST("/provinces", controller.CreateProvince)
			router.PATCH("/provinces", controller.UpdateProvince)
			router.DELETE("/provinces/:id", controller.DeleteMemberClass)

			// Research_Room Routes
			router.GET("/researchrooms", controller.ListResearchRooms)
			router.GET("/researchroom/:id", controller.GetResearchRoom)
			router.POST("/researchrooms", controller.CreateResearchRoom)
			router.PATCH("/researchrooms", controller.UpdateResearchRoom)
			router.DELETE("/researchrooms/:id", controller.DeleteResearchRoom)

			// Equipment Routes
			router.GET("/equipments", controller.ListEquipments)
			router.GET("/equipment/:id", controller.GetEquipment)
			router.POST("/equipments", controller.CreateEquipment)
			router.PATCH("/equipments", controller.UpdateEquipment)
			router.DELETE("/equipments/:id", controller.DeleteEquipment)

			// Room_Type Routes
			router.GET("/roomtypes", controller.ListRoomTypes)
			router.GET("/roomtype/:id", controller.GetRoomType)
			router.POST("/roomtypes", controller.CreateRoomType)
			router.PATCH("/roomtypes", controller.UpdateRoomType)
			router.DELETE("/roomtypes/:id", controller.DeleteRoomType)

			// AddOn Routes
			router.GET("/addons", controller.ListAddOns)
			router.GET("/addon/:id", controller.GetAddOn)
			router.POST("/addons", controller.CreateAddOn)
			router.PATCH("/addons", controller.UpdateAddOn)
			router.DELETE("/addons/:id", controller.DeleteAddOn)

			// Timeroom Routes
			router.GET("/timerooms", controller.ListTimes)
			router.GET("/timeroom/:id", controller.GetTime)
			router.POST("/timerooms", controller.CreateTime)
			router.PATCH("/timerooms", controller.UpdateTime)
			router.DELETE("/timerooms/:id", controller.DeleteTime)

			// Research_Room_Reservation_Record Routes
			router.GET("/researchroomreservationrecords", controller.ListResearchRoomReservationRecords)
			router.GET("/researchroomreservationrecord/:id", controller.GetResearchRoomReservationRecord)
			router.POST("researchroomreservationrecords", controller.CreateResearchRoomReservationRecord)
			router.PATCH("/researchroomreservationrecords", controller.UpdateResearchRoomReservationRecord)
			router.DELETE("/researchroomreservationrecords/:id", controller.DeleteResearchRoomReservationRecord)

			// Computer_os Routes
			router.GET("/computer_oss", controller.ListComputer_oss)
			router.GET("/computer_os/:id", controller.GetComputer_os)
			router.POST("/computer_oss", controller.CreateComputer_os)
			router.PATCH("/computer_oss", controller.UpdateComputer_os)
			router.DELETE("/computer_oss/:id", controller.DeleteComputer_os)

			// Computer_reservation Routes
			router.GET("/computer_reservations", controller.ListComputer_reservations)
			router.GET("/computer_reservation/:id", controller.GetComputer_reservation)
			// router.GET("/playlist/watched/user/:id", controller.GetPlaylistWatchedByUser)
			router.POST("/computer_reservation", controller.CreateComputer_reservation)
			router.PATCH("/computer_reservations", controller.UpdateComputer_reservation)
			router.DELETE("/computer_reservations/:id", controller.DeleteComputer_reservation)

			// Computer Routes
			router.GET("/computers", controller.ListComputers)
			router.GET("/computer/:id", controller.GetComputer)
			router.POST("/computers", controller.CreateComputer)
			router.PATCH("/computers", controller.UpdateComputer)
			router.DELETE("/computers/:id", controller.DeleteComputer)

			// Time_com Routes
			router.GET("/time_coms", controller.ListTime_coms)
			router.GET("/time_com/:id", controller.GetTime_com)
			router.POST("/time_coms", controller.CreateTime_com)
			router.PATCH("/time_coms", controller.UpdateTime_com)
			router.DELETE("/time_coms/:id", controller.DeleteTime_com)
		}
	}

	//Signup User Route
	r.POST("/signup", controller.CreateLoginUser)
	//login User Route
	r.POST("/login", controller.Login)

	//Run the server go run main.go
	r.Run("0.0.0.0:8080")
	//r.Run("localhost: " + PORT)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") //สำหรับใส่ ip server
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
