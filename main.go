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
	title := flag.String("t", "", "Title of the note.")
	text := flag.String("m", "", "Note to save.")
	search := flag.String("s", "", "Text to search.")
	flag.Parse()
	if (*text != "") {
		ok := notesDb.SaveNote(*title, *text);
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
				var note,title string
				err := data.Scan(&id, &title, &note)
				if err != nil {
					log.Fatal(err)
				}
				if len(title)>17 {
					title = title[0:17]+"..."
				}
				fmt.Printf("%d | %-20s | %s\n", id, title, note)
			} 
		} else {
			fmt.Print("No data")
		}
		
	}
}
