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

	// file, err := os.Open("csv/dff-full.csv")
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
	// 		URL:         record[2],
	// 		Description: record[3],
	// 	}
	// 	term.DomainID = 12
	// 	term.SourceID = 11

	// 	db.Create(&term)
	// }

	// source := models.Source{
	// 	ID:          11,
	// 	Name:      "Dubai Future Foundation",
	// 	NameAr:        "مؤسسة دبي للمستقبل",
	//  	Description: "مؤسسة دبي للمستقبل هي أول مؤسسة وقف بحثي في المنطقة العربية تم تأسيسها في أبريل 2016 بهدف تعزيز مكانة دولة الإمارات ومدينة دبي كمركز عالمي لاستشراف وصناعة المستقبل. وتشرف المؤسسة على تنفيذ أجندة دبي المستقبل والتي تندرج تحتها مجموعة واسعة من المبادرات الهادفة لتعزيز قدرات الأفراد وبناء الثقافة في مجال استشراف المستقبل إضافة إلى زيادة جاهزية المؤسسات للمتغيرات المستقبلية وأخيراً بناء قطاعات مستقبلية ذات تأثير إيجابي على الاقتصاد الوطني. تم إطلاق “مؤسسة دبي للمستقبل” بهدف أداء دور محوري في استشراف المستقبل بإمارة دبي تماشياً مع إطلاق أجندة دبي المستقبل لتكون خارطة طريق تسترشد بها المؤسسة في استشراف مستقبل القطاعات الاستراتيجية على المدى متوسط وطويل الأجل، وذلك بالتعاون مع العديد من الجهات والمؤسسات في القطاعين العام والخاص.",
	//  	URL:         "https://www.dubaifuture.ae/ar/",
	//  }
	//  db.Create(&source)

	app := fiber.New()
	app.Static("/", "./public")
	app.Use(cors.New())
	controller := &controllers.App{DB: db}
	routes.SetRoutes(app, controller)

	app.Use(routes.NotFoundMiddleware)

	log.Fatal(app.Listen(port))
}
