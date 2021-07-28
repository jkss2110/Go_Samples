package models

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf("Title:\t\t%q\n"+"Auther:\t\t%q\n"+"Published:\t\t%v\n", b.Title, b.Author, b.YearPublished)
}

var Books = []Book{
	{
		ID:            1,
		Title:         "Algorithms Core",
		Author:        "Anupa",
		YearPublished: 1992,
	},
	{
		ID:            2,
		Title:         "THe Hobbit",
		Author:        "JK Rowling",
		YearPublished: 1937,
	},
	{
		ID:            3,
		Title:         "Data Structures",
		Author:        "JK",
		YearPublished: 2003,
	},
	{
		ID:            4,
		Title:         "Computer Architecture",
		Author:        "Some guy",
		YearPublished: 1978,
	},
	{
		ID:            5,
		Title:         "Sapians",
		Author:        "Harare",
		YearPublished: 2001,
	},
	{
		ID:            6,
		Title:         "HomoSapiens",
		Author:        "JK1",
		YearPublished: 2002,
	},
	{
		ID:            7,
		Title:         "HP Lap",
		Author:        "Mine",
		YearPublished: 2018,
	},
	{
		ID:            8,
		Title:         "Dell Lap",
		Author:        "Company",
		YearPublished: 2019,
	},
	{
		ID:            9,
		Title:         "Green Leafs",
		Author:        "Bottle",
		YearPublished: 2020,
	},
	{
		ID:            10,
		Title:         "The darkness",
		Author:        "lime",
		YearPublished: 2003,
	},
}
