package models

import (
	"log"

	"gorm.io/gorm"
)

type Domain struct {
	ID     uint   `gorm:"primaryKey"`
	Name   string `gorm:"unique;column:name"`
	NameAr string `gorm:"unique;column:name_ar"`
	NameDE string `gorm:"unique;column:name_de"`
	NameFr string `gorm:"unique;column:name_fr"`
}
type DomainCount struct {
	DomainID uint
	Cnt      int
	Name     string
	NameAr   string
}

type DomainWTermCount struct {
	Domain
	TermCount int64 `gorm:"column:term_count"`
}

func GetAllDomains(db *gorm.DB) ([]Domain, error) {
	var domains []Domain
	if err := db.Find(&domains).Error; err != nil {
		return nil, err
	}
	log.Printf("Found %d domains in the database", len(domains))
	return domains, nil
}

// func GetDomainsWithTermCounts(db *gorm.DB) ([]Domain, error) {
// 	var domains []Domain

// 	err := db.Table("domains").
// 		Select("domains.*, COUNT(terms.id) as term_count").
// 		Joins("LEFT JOIN terms ON terms.domain_id = domains.id").
// 		Group("domains.id").
// 		Scan(&domains).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return domains, nil
// }

func GetDomainsWithTermCounts(db *gorm.DB) ([]DomainWTermCount, error) {
	var results []DomainWTermCount

	err := db.Preload("Term").
		Table("domains").
		Select(`domains.id, domains.name, domains.name_ar, domains.name_de, domains.name_fr, 
		        COUNT(terms.id) AS term_count`).
		Joins("LEFT JOIN terms ON terms.domain_id = domains.id").
		Group("domains.id, domains.name, domains.name_ar, domains.name_de, domains.name_fr").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	return results, nil
}

func GetDomainsWithTermCountsRaw(db *gorm.DB) ([]DomainWTermCount, error) {
	var domains []DomainWTermCount

	query := `
		SELECT 
			d.id,
			d.name,
			d.name_ar,
			d.name_de,
			d.name_fr,
			COUNT(t.id) AS term_count
		FROM domains d
		LEFT JOIN terms t ON t.domain_id = d.id
		GROUP BY d.id, d.name, d.name_ar, d.name_de, d.name_fr;
	`

	err := db.Raw(query).Scan(&domains).Error
	if err != nil {
		return nil, err
	}

	return domains, nil
}

// db.Create(&models.Domain{Name: "Automotive Engineering", NameFr: "Technique automobile", NameDE: "Kfz-Technik", NameAr: "هندسة وتكنولوجيا السيارات"})
// db.Create(&models.Domain{Name: "Water Engineering", NameFr: "Technologie de l’eau", NameDE: "Wassertechnik", NameAr: "هندسة المياه"})
// db.Create(&models.Domain{Name: "Renewable Energy", NameFr: "Energies Renouvelables", NameDE: "Erneuerbare Energien", NameAr: "الطاقات المتجددة"})
// db.Create(&models.Domain{Name: "Electrical Engineering", NameFr: "Génie Electrique", NameDE: "Elektrotechnik", NameAr: "الهندسة الكهربائية"})
// db.Create(&models.Domain{Name: "Transport and Infrastructure", NameFr: "Transport et Infrastructure", NameDE: "Transport und Infrastruktur", NameAr: "النقل والبنية التحتية"})
// db.Create(&models.Domain{Name: "Textiles Industries", NameFr: "l’Industrie Textile", NameDE: "Textilindustrie", NameAr: "صناعة النسيج"})
// db.Create(&models.Domain{Name: "Civil Engineering", NameFr: "Génie Civil", NameDE: "Bauingenieurwesen", NameAr: "الهندسة المدنية"})
// db.Create(&models.Domain{Name: "Information and Communication", NameFr: "Technologie de l’Information", NameDE: "Informationstechnologie", NameAr: "تقانة المعلومات"})
// db.Create(&models.Domain{Name: "Climate, Environment and Solid Waste management", NameFr: "Climat, l’Environment, et la Gestion des déchets solides", NameDE: "Klima, Umwelt und Abfallwirtschaft", NameAr: "المناخ والبيئة وإدارة النفايات الصلبة"})
// db.Create(&models.Domain{Name: "Educational and Computer Techniques", NameFr: "Techniques Pédagogiques et Informatiques", NameDE: "Pädagogik und Informatiktechniken", NameAr: "التقنيات التربوية والحاسوبية"})
// db.Create(&models.Domain{Name: "Education", NameFr: "Education", NameDE: "Erziehungswissenschaft", NameAr: "التربية"})
// db.Create(&models.Domain{Name: "Sociology and Anthropology", NameFr: "Sociologie et Anthropologie", NameDE: "Soziologie und Anthropologie", NameAr: "علم الاجتماع والأنثروبولوجيا"})
// db.Create(&models.Domain{Name: "Economics", NameFr: "Economie", NameDE: "Wirtschaft", NameAr: "الاقتصاد"})
// db.Create(&models.Domain{Name: "Commerce and Accounting", NameFr: "Commerce et Comptabilité", NameDE: "Handel und Rechnungswesen", NameAr: "التجارة والمحاسبة"})
// db.Create(&models.Domain{Name: "Law", NameFr: "Droit", NameDE: "Rechtswissenschaft", NameAr: "القانون"})
// db.Create(&models.Domain{Name: "Mathematics and Astronomy", NameFr: "Mathématiques et Astronomie", NameDE: "Mathematik und Astronomie", NameAr: "الرياضيات والفلك"})
// db.Create(&models.Domain{Name: "Physics", NameFr: "Physique", NameDE: "Physik", NameAr: "الفيزياء"})
// db.Create(&models.Domain{Name: "Chemistry", NameFr: "Chimie", NameDE: "Chemie", NameAr: "الكيمياء"})
// db.Create(&models.Domain{Name: "Geology", NameFr: "Géologie", NameDE: "Geologie", NameAr: "الجيولوجيا"})
// db.Create(&models.Domain{Name: "Seismology", NameFr: "Séismologie", NameDE: "Seismologie", NameAr: "علم الزلازل"})
// db.Create(&models.Domain{Name: "Meteorology", NameFr: "Météorologie", NameDE: "Meteorologie", NameAr: "الأرصاد الجوية"})
// db.Create(&models.Domain{Name: "Oceanology", NameFr: "Océanographie", NameDE: "Ozeanographie", NameAr: "علوم البحار"})
// db.Create(&models.Domain{Name: "Petroleum", NameFr: "Pétrole", NameDE: "Erdöl", NameAr: "النفط"})
// db.Create(&models.Domain{Name: "Biology", NameFr: "Biologie", NameDE: "Biologie", NameAr: "علم الأحياء"})
// db.Create(&models.Domain{Name: "Hygienics and Human Body", NameFr: "Santé et Corps Humain", NameDE: "Hygiene und Menschlicher Körper", NameAr: "الصحة وجسم الإنسان"})
// db.Create(&models.Domain{Name: "Genetics", NameFr: "Génétique", NameDE: "Genetik", NameAr: "علم الوراثة"})
// db.Create(&models.Domain{Name: "Pharmacy", NameFr: "Pharmacie", NameDE: "Pharmazeutik", NameAr: "علم الصيدلة"})
// db.Create(&models.Domain{Name: "Electronic Warfare", NameFr: "Guerre éléctronique", NameDE: "Elektronische Kriegführung", NameAr: "الحرب الإلكترونية"})
// db.Create(&models.Domain{Name: "Remote Sensing", NameFr: "Télédétection", NameDE: "Fernerkundung", NameAr: "الاستشعار عن بعد"})
// db.Create(&models.Domain{Name: "Veterinary Medicine", NameFr: "Médecine Vétérinaire", NameDE: "Veterinärmedizin", NameAr: "الطب البيطري"})
// db.Create(&models.Domain{Name: "Gross Anatomy", NameFr: "Anatomie Macroscopique", NameDE: "Mikroskopische Anatomie", NameAr: "التشريح العياني"})
// db.Create(&models.Domain{Name: "Masonry - Carpentry", NameFr: "Maçonnerie - Charpenterie", NameDE: "Maurerhandwerk - Zimmerhandwerk", NameAr: "البناء - النجارة"})
// db.Create(&models.Domain{Name: "Printing - Electricity", NameFr: "Imprimerie - Electricité", NameDE: "Buchdruck - Elektrizität", NameAr: "الطباعة - الكهرباء"})
// db.Create(&models.Domain{Name: "Nutrition Technologies", NameFr: "Technologies Alimentaires", NameDE: "Nahrungsmitteltechnologie", NameAr: "تقانات الأغذية"})
// db.Create(&models.Domain{Name: "Information and Communication", NameFr: "Information et Communication", NameDE: "Information und Kommunikation", NameAr: "الإعلام والتواصل"})
// db.Create(&models.Domain{Name: "Philosophy and Psychology", NameFr: "Philosophie et Psychologie", NameDE: "Philosophie und Psychologie", NameAr: "الفلسفة وعلم النفس"})
// db.Create(&models.Domain{Name: "Arts, Recreation and Sports", NameFr: "Art, Divertissement et sports", NameDE: "Kunst, Vergnügung und Sport", NameAr: "الفن، التسلية والرياضة"})
// db.Create(&models.Domain{Name: "Language and Literature", NameFr: "Langue et Littérature", NameDE: "Sprache und Literatur", NameAr: "اللغة والأدب"})
// db.Create(&models.Domain{Name: "Geography and History", NameFr: "Géographie et Histoire", NameDE: "Geographie und Geschichte", NameAr: "الجغرافيا والتاريخ"})
