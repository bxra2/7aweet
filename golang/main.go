package main

import (
	"encoding/csv"
	"fmt"
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

	file, err := os.Open("csv/tdra.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	for _, record := range records[1:] {
		term := models.Term{
			English:     record[0],
			Arabic:      record[1],
			French:      "",
			German:      "",
			URL:         "https://tdra.gov.ae/ar/Pages/ict-dictionary#P",
			Description: record[2],
		}
		term.DomainID = 12
		term.SourceID = 10

		db.Create(&term)
	}

	source := models.Source{
		ID:          10,
	 	Name:        "هيئة تنظيم الاتصالات والحكومة الرقمية",
	 	NameAr:      "TDRA",
	 	Description: "يتلخص دور الهيئة التنظيمي في ضمان تأمين خدمات اتصالات متميزة، وتطوير قطاع الاتصالات، ورعاية مصالح الأطراف، وتطبيق أطر السياسات والأنظمة ذات الصلة، وتطوير الموارد البشرية، وتشجيع البحوث والتطوير، بما يضمن للإمارات العربية مكانة إقليمية وعالمية رائدة في قطاع الاتصالات. وفي مجال تمكين التحول الرقمي، تتولى الهيئة مسؤولية الإشراف على الحكومة الرقمية الاتحادية بموجب القانون رقم 3 لسنة 2011. ومنذ ذلك الوقت أصبحت الهيئة مسؤولة عن التحول الرقمي على المستوى الوطني من خلال هدفين استراتيجيين هما: تعزيز أسلوب الحياة الذكي، والريادة في البنية التحتية الرقمية في دولة الإمارات العربية المتحدة. ",
	 	URL:         "https://tdra.gov.ae/",
	 }
	 db.Create(&source)

	app := fiber.New()
	app.Static("/", "./public")
	app.Use(cors.New())
	controller := &controllers.App{DB: db}
	routes.SetRoutes(app, controller)

	app.Use(routes.NotFoundMiddleware)

	log.Fatal(app.Listen(port))
}
