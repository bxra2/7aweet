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

	// file, err := os.Open("csv/yelp_cleaned.csv")
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
	// 		URL:         "https://gitlab.gnome.org/GNOME",
	// 		Description: "",
	// 	}
	// 	term.DomainID = 12
	// 	term.SourceID = 6

	// 	db.Create(&term)
	// }

	// domain := models.Source{
	// 	ID:          6,
	// 	Name:        "Gnome",
	// 	NameAr:      "جنوم",
	// 	Description: "جنوم أو غنوم (بالإنجليزية: GNOME) (تلفظ /ˈnoʊm/ أو /ɡəˈnoʊm/) هو مشروع عالمي يهدف إلى إنشاء بيئة سطح مكتب سهلة الاستعمال مبنية حصرا على البرمجيات الحرة. يدخل في هذا المسعى تكوين أدوات تساعد المطورين على كتابة تطبيقات برمجية قائمة بذاتها، وانتقاء وترشيح أحسن البرامج لتكون ضمن الإصدار الرسمي، والتركيز على ما يسمى «البيئة المكتبية»، أي جميع البرامج المتدخلة في إطلاق البرامج وتدبير المستندات وتنظيم النوافذ والأشغال الجارية. جنوم جزء من مشروع جنو ويمكن استخدامها مع مختلف أنظمة التشغيل الشبيهة بيونيكس، وبالأخص لينكس، وكذلك كجزء من سطح مكتب أوبن سولاريس. يستضيف مشروع جنوم باقة من البرامج المختلفة، والتي تُنتخب منها مجموعة محدودة تصدر تحت الاسم الشامل جنوم. يعمل جنوم على نظام تشغيل مثل لينكس أو سولاريس لشركة صن ميكروسيستيمز مكونين نظامًا حاسوبيًا مكتملَ الأوصاف. يعد جنوم جزءًا من نظام تشغيل جنو، وهو بيئته المكتبية الرسمية. تنطق كلمة جنوم رسميا /گْنُوْمْ/، غير أن نطقها /نُوْمْ/ شائع أيضا. يعرب اسمه رسميًا «جنوم».",
	// 	URL:         "https://www.gnome.org/",
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
