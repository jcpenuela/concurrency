package main

import "fmt"

// Book record format
type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	Book{ID: 1, Title: "The hobbit", Author: "Tolkien", YearPublished: 1967},
	Book{ID: 2, Title: "1984", Author: "Orwell", YearPublished: 1936},
	Book{ID: 3, Title: "Multiverse", Author: "Kaku", YearPublished: 1998},
	Book{ID: 4, Title: "The doctor", Author: "Gordon", YearPublished: 1994},
	Book{ID: 5, Title: "ACE HIGH", Author: "Michael Atamanov", YearPublished: 2020},
	Book{ID: 6, Title: "THE APOCALYPSE STRAIN", Author: "Jason Parent", YearPublished: 2020},
	Book{ID: 7, Title: "A BEGINNING AT THE END", Author: "Mike Chen", YearPublished: 2020},
	Book{ID: 8, Title: "ABRACADAVER", Author: "Laura Resnick", YearPublished: 2014},
	Book{ID: 9, Title: "ALIAS HOOK", Author: "Lisa Jensen", YearPublished: 2014},
	Book{ID: 10, Title: "ACE IN THE HOLE", Author: "George R. R. Martin", YearPublished: 1990},
	Book{ID: 11, Title: "THE BARRENS", Author: "F. Paul Wilson", YearPublished: 1990},
	Book{ID: 12, Title: "THE CAIRENE PURSE", Author: "Michael Moorcock", YearPublished: 1990},
}
