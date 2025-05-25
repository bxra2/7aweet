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
		port = ":4000"
	} else {
		port = ":" + port
	}

	db, err := db.OpenDatabase()
	if err != nil {
		log.Fatalf("failed to Open database: %v", err)
	}
	err = db.AutoMigrate(&models.Source{}, &models.Term{}, &models.Domain{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// file, err := os.Open("csv/sdaia/sdaia.csv")
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
	// 		English:     record[0],
	// 		Arabic:      record[1],
	// 		French:      "",
	// 		German:      "",
	// 		URL:         "https://sdaia.gov.sa/en/MediaCenter/KnowledgeCenter/ResearchLibrary/SDAIAPublications15.pdf",
	// 		Description: record[2],
	// 	}
	// 	term.DomainID = 12
	// 	term.SourceID = 9

	// 	db.Create(&term)
	// }

	// source := models.Source{
	// 	ID:          9,
	// 	Name:        "الهيئة السعودية للبيانات والذكاء الاصطناعي",
	// 	NameAr:      "SDAIA",
	// 	Description: "الهيئة السعودية للبيانات والذكاء الاصطناعي (​سدايا) هي الجهة المختصة في المملكة بالبيانات والذكاء الاصطناعي وتشمل: البيانات الضخمة، وهي المرجع الوطني في كل ما يتعلق بهما من تنظيم وتطوير وتعامل، وهي صاحبة الاختصاص الأصيل في كل ما يتعلق بالتشغيل والأبحاث والابتكار في قطاع البيانات والذكاء الاصطناعي",
	// 	URL:         "https://sdaia.gov.sa/ar/default.aspx",
	// }
	// db.Create(&source)

	app := fiber.New()
	app.Static("/", "./public")
	app.Use(cors.New())
	controller := &controllers.App{DB: db}
	routes.SetRoutes(app, controller)

	app.Use(routes.NotFoundMiddleware)

	log.Fatal(app.Listen(port))
}
