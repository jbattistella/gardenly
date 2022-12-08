package database

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/joho/godotenv"
)

var DB *pg.DB

func ConnectDB() error {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	opts := &pg.Options{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Addr:     os.Getenv("DB_ADDR"),
		Database: os.Getenv("DB_DATABASE"),
	}

	DB = pg.Connect(opts)
	if DB == nil {
		log.Printf("Database connection failed.\n")
		os.Exit(100)
	} else {
		log.Printf("connected")
	}
	return nil
}

func UpdateDateDataBase() {
	if err := createVegetableSchema(DB); err != nil {
		log.Fatal(err)
	}

	data, err := loadDataBaseWithCsv()
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Model(&data).Insert()
	if err != nil {
		log.Println(err)
	}
}

func createVegetableSchema(db *pg.DB) error {
	models := []interface{}{
		(*Vegetable)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func loadDataBaseWithCsv() ([]Vegetable, error) {

	csvFile, _ := os.Open("vegetables2.csv")
	reader := csv.NewReader(csvFile)
	reader.TrimLeadingSpace = true

	var veg []Vegetable

	for {
		line, error := reader.Read()

		if error == io.EOF {
			break

		} else if error != nil {
			log.Fatal(error)
		}

		dtm, err := strconv.Atoi(line[1])
		if err != nil {
			return []Vegetable{}, err
		}
		dtt, err := strconv.Atoi(line[2])
		if err != nil {
			return []Vegetable{}, err
		}

		veg = append(veg, Vegetable{
			CommonName: line[0],
			DTM:        dtm,
			DownToTemp: dtt,
		})

	}
	return veg, nil
}
