package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func DatabaseInit() *sql.DB {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)

	//psqlconn := "host=postgresql.ad.vejtek.cz database=dopravci user=konici password=Fl1jiAmyQ0Gx9hl3bMf8Yz7PI9DAS9H5Ngm1"
	db, _ := sql.Open("postgres", psqlconn)
	//defer db.Close()
	return db
}


func DatabaseReCreateTables() bool {
	_, err := db.Exec(`

		DROP TABLE IF EXISTS public.cena CASCADE;
		DROP TABLE IF EXISTS public.sluzba CASCADE;
		DROP TABLE IF EXISTS public.stat CASCADE;
		DROP TABLE IF EXISTS public.dopravce CASCADE;
		CREATE TABLE public.dopravce
        (
			id integer NOT NULL,
			jmeno text,
            palivovy_priplatek double precision,
			CONSTRAINT dopravce_pkey PRIMARY KEY (id)
		);

		CREATE TABLE public.stat (
			id integer NOT NULL,
			nazev text,
			CONSTRAINT stat_pkey PRIMARY KEY (id)
		);

		CREATE TABLE public.sluzba (
			id integer NOT NULL,
			nazev text,
			dopravce_id integer NOT NULL,
            cena_dobirky double precision,
			CONSTRAINT sluzby_pkey PRIMARY KEY (id),
			CONSTRAINT sluzba_dopravce_fk
				FOREIGN KEY (dopravce_id)
				REFERENCES public.dopravce(id)
		);

		CREATE TABLE public.cena (
			id integer NOT NULL,
			dopravce_id integer NOT NULL,
			stat_id integer NOT NULL,
			sluzba_id integer NOT NULL,
			cena double precision,
			max_hmotnost double precision,
			CONSTRAINT cena_pkey PRIMARY KEY (id),
			CONSTRAINT cena_dopravce_fk
				FOREIGN KEY (dopravce_id)
				REFERENCES public.dopravce(id),
			CONSTRAINT cena_stat_fk
				FOREIGN KEY (stat_id)
				REFERENCES public.stat(id),
			CONSTRAINT cena_sluzba_fk
				FOREIGN KEY (sluzba_id)
				REFERENCES public.sluzba(id)
		);

	`)
	if err != nil {
		fmt.Println("Chyba při vytváření tabulek:", err)
		return false
	}
	return true
}
