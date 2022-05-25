package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"regexp"
)

type titles_info struct {
	Title string
	Info  string
}

type combo_field struct {
	value string
}

func main() {
	http.HandleFunc("/", main_page)
	http.ListenAndServe(":8080", nil)
}

func main_page(w http.ResponseWriter, r *http.Request) {
	reg, _ := regexp.Compile("^[0-9]")

	info := titles_info{"This is the combo cracker! By: Bryce James", "Based on a genetic algorithm.\nHTML rendered by golang html templates."}
	fp := path.Join("templates", "index.html")
	tmpl := template.Must(template.ParseFiles(fp))

	details := combo_field{
		value: r.FormValue("combination"),
	}

	if r.Method != http.MethodPost {
		var data = struct {
			Submitted bool
			Title     string
			Info      string
			Error     string
		}{
			Submitted: false,
			Title:     info.Title,
			Info:      info.Info,
			Error:     "",
		}
		err := tmpl.Execute(w, &data)
		if err != nil {
			fmt.Println(err)
		}
		_ = details
		return

	} else if !reg.MatchString(details.value) || len(details.value) != 10 {

		var data = struct {
			Submitted bool
			Title     string
			Info      string
			Error     string
		}{
			Submitted: false,
			Title:     info.Title,
			Info:      info.Info,
			Error:     "Something went wrong with submission. Please enter 10 digits 0-9, no letters.",
		}
		err := tmpl.Execute(w, &data)
		if err != nil {
			fmt.Println(err)
		}
		_ = details
		return
	} else {
		// /fmt.Printf("\n%+v\n", details)
		answer := Naive_Hill_Climb(details.value)
		var data = struct {
			Submitted bool
			Mutations string
			Best_Fit  string
			Count     string
			Title     string
			Info      string
			Error     string
		}{
			Submitted: true,
			Mutations: answer[0],
			Best_Fit:  answer[1],
			Count:     answer[2],
			Title:     info.Title,
			Info:      info.Info,
			Error:     "",
		}
		err := tmpl.Execute(w, &data)
		if err != nil {
			fmt.Println(err)
		}
		_ = details
		return
	}

}

/*func ShowBooks(w http.ResponseWriter, r *http.Request) {
	book := Book{"Building Web Apps with Go", "Jeremy Saenz"}

	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}*/