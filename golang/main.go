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

	// file, err := os.Open("csv/MMA.csv")
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
	// 		English:     record[3],
	// 		Arabic:      record[0],
	// 		French:      "",
	// 		German:      "",
	// 		URL:         "https://www.stats.gov.sa/glossary-of-statistical-terms",
	// 		Description: record[1],
	// 	}
	// 	term.DomainID = 43
	// 	term.SourceID = 13

	// 	db.Create(&term)
	// }

	// source := models.Source{
	// 	ID:          13,
	// 	Name:      		"General Authority for Statistics",
	// 	NameAr:        	"الهيئة العامة للإحصاء",
	//  	Description: "تعــد الإحصاءات أحــد أهــم العناصــر الأساسية لعمليــة التخطيــط ودعــم اتخــاذ القــرارات لكافــة القطاعــات. ومــن هــذا المنطلــق تم العمل على هــذا المعجم الذي يحتوي على قائمة كاملة بجميع المصطلحات والتعريفات الإحصائية المستخدمة في كافة المنهجيات والتقارير والنشرات الإحصائية الصادرة عن الهيئة العامة للإحصاء باللغتين العربية والانجليزية. تأتي أهمية هذا المعجم في كونه يوفر جميع المصطلحات المستخدمة في الهيئة العامة للإحصاء، مما يجعل من السهل توحيد استخدام هذه المصطلحات أينما وردت، إضافة إلى نشره على البوابة الإلكترونية الرسمية للهيئة، لتستفيد منه كافة الجهات والمنظمات ذوات العلاقة واستخدام المصطلحات متى دعت الحاجة إلى ذلك. نأمــل بــأن يســاهم هذا المعجم فــي تعزيــز الشــفافية في العمــل الإحصائي وزيــادة الوعــي الإحصائي لجميــع الفئــات المســتخدمة للبيانـات والمعلومــات والمؤشــرات الإحصائية، وستقوم الهيئة بتطويــره وتحديثــه بشكل مستمر حسب الحاجة والمتغيرات. يتضمن المعجم اكثر من 850 مصطلح إحصائي, وسيتم تحديثه دورياً لمواكبة التطورات والمستجدات في المجال الإحصائي.",
	//  	URL:         "https://www.stats.gov.sa/glossary-of-statistical-terms",
	//  }
	//  db.Create(&source)

	// domain := models.Domain{
	// 	Name:   "Statistics",
	// 	NameAr: "الإحصاء",
	// 	NameDE: "Statistik",
	// 	NameFr: "Statistiques",
	// }
	//  db.Create(&domain)

	app := fiber.New()
	app.Static("/", "./public")
	app.Use(cors.New())
	controller := &controllers.App{DB: db}
	routes.SetRoutes(app, controller)

	app.Use(routes.NotFoundMiddleware)

	log.Fatal(app.Listen(port))
}
