package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"math/rand"


	_ "github.com/go-sql-driver/mysql"
)

type cars struct {
	Brand, Model, HorsePower string
	ID                       int
}

func main() {

	http.HandleFunc("/services/v1/cars", postCars)
	direccion := ":8080"
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
	fmt.Println("hola")
}

func postCars(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/services/v1/cars" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {

	case "POST":
		var c cars

		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			panic(err)
		}

		log.Println(c.Brand, c.Model, c.HorsePower)
		correcto := insertar(c, w)

		if correcto != nil {
			fmt.Printf("Error insertando: %v", correcto)
		} else {
			fmt.Println("Insertado correctamente")
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := "roger"
	host := "tcp(database)"
	nombreBaseDeDatos := "entrust"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func insertar(c cars, w http.ResponseWriter) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()
	var ID = rand.Intn(10000)
	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("INSERT INTO coche (ID, Brand_victoria2, Model, HorsePower) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(ID, c.Brand, c.Model, c.HorsePower)
	c.ID = ID;
	json.NewEncoder(w).Encode(c)
	if err != nil {
		return err
	}
	return nil
}
