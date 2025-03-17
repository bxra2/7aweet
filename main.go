package main

import (
	"log"
	"os"

	"github.com/bxra2/7aweet/controllers"
	"github.com/bxra2/7aweet/db"
	"github.com/bxra2/7aweet/models"
	"github.com/bxra2/7aweet/routes"
	"github.com/gofiber/fiber/v2"
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

	// file, err := os.Open("information_tech.csv")
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
	// 		French:      record[1],
	// 		German:      record[2],
	// 		Arabic:      record[3],
	// 		Description: record[4],
	// 		URL:         record[5],
	// 	}
	// 	domainID, err := strconv.ParseUint(record[6], 10, 64)
	// 	if err != nil {
	// 		fmt.Println("Error parsing DomainID:", err)
	// 		continue
	// 	}
	// 	term.DomainID = uint(domainID)

	// 	sourceID, err := strconv.ParseUint(record[7], 10, 64)
	// 	if err != nil {
	// 		fmt.Println("Error parsing DomainID:", err)
	// 		continue
	// 	}
	// 	term.SourceID = uint(sourceID)

	// 	db.Create(&term)
	// }

	app := fiber.New()
	controller := &controllers.App{DB: db}
	routes.SetRoutes(app, controller)

	app.Use(routes.NotFoundMiddleware)

	log.Fatal(app.Listen(port))
}
