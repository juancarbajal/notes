package main

import (
	"flag"
	"fmt"
	"log"

	db "github.com/juancarbajal/notes/pkg"
)

// main ...
func main()  {
	notesDb := db.NewDb()
	text := flag.String("a", "", "Note to save.")
	search := flag.String("s", "", "Text to search.")
	flag.Parse()
	if (*text != "") {
		ok := notesDb.SaveNote(*text);
		if (ok){
			fmt.Print("Note saved")
		} else {
			fmt.Print("Error to save note")
		}
	}
	if (*search != "") {
		data, ok := notesDb.SearchNote(*search);
		defer data.Close()
		if (ok){
			for data.Next(){
				var id int
				var note string
				err := data.Scan(&id, &note)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("%d | %s\n", id, note)
			} 
		} else {
			fmt.Print("No data")
		}
		
	}
}
