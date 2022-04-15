package models

var Db []Book

type Book struct {
	Id          int
	Title       string `json:"title"`
	CreatedYear int16  `json:"created_year"`
	Author      Author `json:"author"`
}

type Author struct {
	Name      string `json:"name"`
	BirthYear int16  `json:"birth_year"`
}

func init() {
	book1 := Book{
		1,
		"Искусство программирования",
		1968,
		Author{
			"Knuth",
			1938,
		},
	}

	Db = append(Db, book1)
}

func GetBookById(id int) (*Book, bool) {
	for key, bookInDb := range Db {
		if bookInDb.Id == id {
			return &Db[key], true
		}
	}

	return nil, false
}
