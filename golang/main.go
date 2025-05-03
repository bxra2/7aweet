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

	// file, err := os.Open("csv/nemo-extensions_cleaned.csv")
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

	// for _, record := range records {
	// 	term := models.Term{
	// 		English:     record[0],
	// 		Arabic:      record[3],
	// 		French:      record[1],
	// 		German:      record[2],
	// 		URL:         "https://github.com/linuxmint/cinnamon-translations/tree/master/po-export",
	// 		Description: "",
	// 	}
	// 	term.DomainID = 12
	// 	term.SourceID = 7

	// 	db.Create(&term)
	// }

	// domain := models.Source{
	// 	ID:          8,
	// 	Name:        "KDE",
	// 	NameAr:      "كيدي",
	// 	Description: "كِيدِي أو كدي (بالإنجليزية: KDE)‏ هي منظمة برمجيات حرة تختص بتطوير برمجيات حرة مفتوحة المصدر. من أبرز منتجاتها بيئة سطح المكتب بلازما، وأطر عمل كيدي، والعديد من البرمجيات الأخرى مثل كيت، ودجي‌كام، وكريتا. تستخدم منتجات كيدي مكتبة كيوت، وبعض برمجياتها متعددة المنصات وتعمل على أنظمة يونكس، وشبيه يونكس، ومايكروسوفت ويندوز، وأندرويد.شعار كِيدِي هو حرف K فوق عجلة ترس ملونان بالأبيض داخل خلفية زرقاء. وتميمة المشروع هو تنين أخضر يدعى كونكي.",
	// 	URL:         "https://kde.org/",
	// }
	// db.Create(&domain)

	app := fiber.New()
	app.Static("/", "./public")
	app.Use(cors.New())
	controller := &controllers.App{DB: db}
	routes.SetRoutes(app, controller)

	app.Use(routes.NotFoundMiddleware)

	log.Fatal(app.Listen(port))
}
