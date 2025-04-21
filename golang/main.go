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

	// file, err := os.Open("ds1.csv")
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
	// 		Arabic:      record[2],
	// 		French:      record[1],
	// 		German:      "",
	// 		URL:         "https://arabacademy-sy.org//uploads/academy_publication/books/civilization-1.pdf",
	// 		Description: record[3],
	// 	}
	// 	term.DomainID = 16
	// 	term.SourceID = 5

	// 	db.Create(&term)
	// }

	// domain := models.Source{
	// 	ID:          5,
	// 	Name:        "Arab Academy of Damascus",
	// 	NameAr:      "مجمع اللغة العربية بدمشق",
	// 	Description: "مَجْمَعُ اللُّغةِ العَرَبيَّةِ بدِمَشق، عُرف سابقاً باسم المجمع العلمي العربي حتى سنة 1960م، هو أقدم مَجمَع للُّغة العربيَّة في الوطن العربي؛ تأسَّس في عهد حكومة الملك فيصل سنة 1919م في الثامن من حَزِيران في سوريا للنهوض باللُّغة العربيَّة. وكان له أثرٌ كبير في تعريب مؤسسات الدَّولة وهيئاتها، وتعريب التَّعليم، وإنشاء المدارس الأولى في سوريا والدُّول العربيَّة. وهو مَجمَع أكاديمي يتألف من عشرين عضوًا من علماء اللُّغة العربيَّة في سوريا والمتخصِّصين بها والمَعنيِّين بها من ذوي التخصُّصات العلمية، يشكلون عدَّة لجان؛ كلجنة المخطوطات وإحياء التراث، ولجنة المصطلحات، ولجنة اللَّهَجات العربيَّة المُعاصِرة. وهو عضو في اتحاد المجامع اللُّغوية العِلمية العربيَّة.",
	// 	URL:         "https://arabacademy-sy.org/",
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
