package main

import (
	"cosmart-library/database"
	"cosmart-library/entity/model"
	"cosmart-library/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Author struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type Work struct {
	Key               string   `json:"key"`
	Title             string   `json:"title"`
	Edition_count     int      `json:"edition_count"`
	Cover_id          int      `json:"cover_id"`
	Cover_edition_key string   `json:"cover_edition_key"`
	Authors           []Author `json:"authors"`
}

type OpenLibrary struct {
	Key         string `json:"key"`
	Name        string `json:"name"`
	SubjectType string `json:"subject_type"`
	WorkCount   int    `json:"work_count"`
	Works       []Work `json:"works"`
}

func main() {

	utils.LoadEnv()
	database.InitDB()

	fmt.Println("opening json file...")

	json_file, err := os.Open("openlibrary.json")
	if err != nil {
		fmt.Println(err)
	}
	defer json_file.Close()

	byteValue, _ := ioutil.ReadAll(json_file)
	var open_library OpenLibrary
	json.Unmarshal(byteValue, &open_library)

	fmt.Println("Starting seeding")

	err = database.Transaction(func() error {
		for i := 0; i < len(open_library.Works); i++ {
			book := model.Book{
				Title:   open_library.Works[i].Title,
				Author:  open_library.Works[i].Authors[0].Name,
				Edition: open_library.Works[i].Edition_count,
			}
			if result := database.GetConn().Create(&book); result.Error != nil {
				return result.Error
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Done seeding, goodluck!")

}
