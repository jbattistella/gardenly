package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	var url = "postgresql://" + os.Getenv("PGUSER") + os.Getenv("PGPASS") + "@" + os.Getenv("PGHOST") + ":" + os.Getenv("PGPORT") + "/railway"

	DB, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Sucessfully created the PostgreSQL server!")

	// postgresql://${{ PGUSER }}:${{ PGPASSWORD }}@${{ PGHOST }}:${{ PGPORT }}/${{ PGDATABASE }}
	// url := "postgresql://postgres:Ra4uTQDKbj5mZNyDsMfn@containers-us-west-89.railway.app:6607/railway"

	// psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
	// 	user, password, host, port, dbname)
	// fmt.Println(psqlInfo)

	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	fmt.Println("here")
	// }
	// defer db.Close()

	// err = db.Ping()
	// if err != nil {
	// fmt.Println("there")
	// }

	return DB, nil
}

// func UpdateDateDataBase() {
// 	if err := createVegetableSchema(DB); err != nil {
// 		log.Fatal(err)
// 	}

// 	data, err := loadDataBaseWithCsv()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = DB.Model(&data).Insert()
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createVegetableSchema(db *pg.DB) error {
// 	models := []interface{}{
// 		(*Vegetable)(nil),
// 	}

// 	for _, model := range models {
// 		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
// 			IfNotExists: true,
// 		})

// 		if err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func loadDataBaseWithCsv() ([]Vegetable, error) {

// 	csvFile, _ := os.Open("vegetables2.csv")
// 	reader := csv.NewReader(csvFile)
// 	reader.TrimLeadingSpace = true

// 	var veg []Vegetable

// 	for {
// 		line, error := reader.Read()

// 		if error == io.EOF {
// 			break

// 		} else if error != nil {
// 			log.Fatal(error)
// 		}

// 		dtm, err := strconv.Atoi(line[1])
// 		if err != nil {
// 			return []Vegetable{}, err
// 		}
// 		dtt, err := strconv.Atoi(line[2])
// 		if err != nil {
// 			return []Vegetable{}, err
// 		}

// 		veg = append(veg, Vegetable{
// 			CommonName: line[0],
// 			DTM:        dtm,
// 			DownToTemp: dtt,
// 		})

// 	}
// 	return veg, nil
// }
