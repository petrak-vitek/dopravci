package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var tmpl = template.Must(template.ParseFiles("home.html"))

type State struct {
	ID   int
	Name string
}

type Service struct {
	Dopravce string
	Name     string
	Price    string
}

type PageData struct {
	States   []State
	Services []Service
}

func home(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.Query(`SELECT "id", "nazev" FROM "stat" ORDER BY "nazev" ASC `)
	states := []State{}
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		states = append(states, State{ID: id, Name: name})
	}

	state, _ := strconv.Atoi(r.FormValue("state"))
	weight, _ := strconv.Atoi(r.FormValue("weight"))
	log.Println("state:", state, "weight:", weight)
	services := []Service{}
	if state != 0 {
		rows2, _ := db.Query(`
WITH vyber AS (
    SELECT
        c.sluzba_id,
        s.nazev AS sluzba,
        d.jmeno AS dopravce,
        c.cena,
        c.max_hmotnost,
        ROW_NUMBER() OVER (
            PARTITION BY c.sluzba_id
            ORDER BY c.max_hmotnost ASC
        ) AS rn
    FROM public.cena c
    JOIN public.sluzba s
        ON s.id = c.sluzba_id
    JOIN public.dopravce d
        ON d.id = s.dopravce_id
    WHERE c.stat_id = $1
      AND c.max_hmotnost >= $2
)
SELECT
    dopravce,
    sluzba,
    cena
FROM vyber
WHERE rn = 1
ORDER BY sluzba;`, state, weight)
		log.Println("query error:", rows2)

		for rows2.Next() {
			var dopravce string
			var name string
			var price string
			rows2.Scan(&dopravce, &name, &price)
			log.Println(rows2)
			services = append(services, Service{Dopravce: dopravce, Name: name, Price: price})
		}
	}
	data := PageData{
		States:   states,
		Services: services,
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Chyba šablony", http.StatusInternalServerError)
	}
}
