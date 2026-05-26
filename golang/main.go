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

	// file, err := os.Open("csv/مصطلحات اللسانيات و المعاجم.csv")
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
	// 		URL:         "https://archive.org/details/NAH2012ENFRAR/",
	// 		Description: "",
	// 	}
	// 	term.DomainID = 16
	// 	term.SourceID = 23
	// 	term.CollectionID = 16
	// 	db.Create(&term)
	// }

	// source := models.Source{
	// 	ID:           23,
	// 	NameAr:       "مصطلحات اللسانيات و المعاجم",
	// 	Name:         "Linguistics and Lexicography Terminology",
	// 	Description:  "",
	// 	URL:          "https://archive.org/details/NAH2012ENFRAR",
	// 	CollectionID: 16,
	// }
	// db.Create(&source)

	// collection := models.Collection{
	// 	ID:          16,
	// 	NameAr:      "المنظمة العربية للترجمة",
	// 	Name:        "Arab Organization for Translation",
	// 	Description: "المنظمة العربية للترجمة منظمة دولية متخصصة، غير حكومية ، مستقلة ولا تهدف إلى الربح، مقرها بيروت - لبنان. وهي تتمتع 'بكافة المزايا والحصانات اللازمة لأداء مهامها، أسوة بالمنظمات الدولية والإقليمية العاملة في إطار الأمم المتحدة'، اعتماداً على المرسوم رقم 2803 الصادر عن الحكومة اللبنانية بتاريخ 14 أبريل 2000. ",
	// 	URL:         "http://www.aot-arab.org/",
	// }
	// db.Create(&collection)

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

	app.Use(routes.SPAFallback("./public/index.html"))

	log.Fatal(app.Listen(port))
}
