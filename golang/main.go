package main

import (
	"log"
	"os"

	"github.com/bxra2/7aweet/controllers"
	"github.com/bxra2/7aweet/db"
	"github.com/bxra2/7aweet/models"
	"github.com/bxra2/7aweet/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	env := godotenv.Load()
	if env != nil {
		panic("cannot find environment variables")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	} else {
		port = ":" + port
	}

	db, err := db.OpenDatabase()
	if err != nil {
		log.Fatalf("failed to Open database: %v", err)
	}
	err = db.AutoMigrate(&models.Source{}, &models.Term{}, &models.Domain{}, &models.Collection{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// file, err := os.Open("csv/قائمة مصطلحات المعلوماتية.csv")
	// if err != nil {
	// 	fmt.Println("Error opening CSV file:", err)
	// 	return
	// }
	// defer file.Close()

	// reader := csv.NewReader(file)
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	fmt.Println("Error reading CSV file:", err)
	// 	return
	// }

	// for _, record := range records[1:] {
	// 	term := models.Term{
	// 		English:     record[2],
	// 		Arabic:      record[1],
	// 		French:      "",
	// 		German:      "",
	// 		URL:         "https://arabacademy-sy.org//uploads/academy_publication/books/informatic-list.pdf",
	// 		Description: "",
	// 	}
	// 	term.DomainID = 12
	// 	term.SourceID = 3
	// 	term.CollectionID = 4
	// 	db.Create(&term)
	// }

	// source := models.Source{
	// 	ID:           3,
	// 	NameAr:       "قائمة مصطلحات المعلوماتية",
	// 	Name:         "informatic list",
	// 	Description:  "",
	// 	URL:          "https://arabacademy-sy.org/uploads/academy_publication/books/informatic-list.pdf",
	// 	CollectionID: 4,
	// }
	// db.Create(&source)

	// domain := models.Domain{
	// 	Name:   "Statistics",
	// 	NameAr: "الإحصاء",
	// 	NameDE: "Statistik",
	// 	NameFr: "Statistiques",
	// }
	//  db.Create(&domain)

	app := fiber.New()
	app.Static("/", "./public")
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3005, http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	controller := &controllers.App{DB: db}
	routes.SetRoutes(app, controller)

	app.Use(routes.NotFoundMiddleware)

	log.Fatal(app.Listen(port))
}
