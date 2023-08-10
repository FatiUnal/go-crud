package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type person struct {
	id   int
	name string
}

func main() {

	urlExample := "postgres://postgres:Fatih.2606@localhost:5432/person"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	//getbyıd
	person := getbyid(conn, 2)
	fmt.Println(person)

	fmt.Println("----------------------")
	getall(conn)

	fmt.Println("----------------------")
	delete(conn, 4)

	fmt.Println("----------------------")
	getall(conn)

	fmt.Println("----------------------")

	addPerson(conn, 7, "disla")
	getall(conn)

}

// add object
func addPerson(conn *pgx.Conn, id int, name string) {

	_, err := conn.Exec(context.Background(), "INSERT INTO person (id,name) VALUES ($1,$2)", id, name)

	if err != nil {
		fmt.Println("Sorgu hatası:", err)
		os.Exit(1)
	}

	fmt.Println("Veri başarıyla eklendi.")
}

// delete by id
func delete(conn *pgx.Conn, idx int) {

	_, err := conn.Exec(context.Background(), "delete from person where id=$1", idx)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Kayıt Silindi")

}

// Get By Id
func getbyid(conn *pgx.Conn, idx int) person {

	var personid int
	var name string

	err := conn.QueryRow(context.Background(), "select * from person where id=$1", idx).Scan(&personid, &name)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	getPerson := person{
		id:   personid,
		name: name,
	}

	return getPerson

}

// get All List
func getall(conn *pgx.Conn) {
	rows, err := conn.Query(context.Background(), "select * from person")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		values, err := rows.Values()

		if err != nil {
			log.Fatal(err)
		}

		id := values[0].(int32)
		name := values[1].(string)

		log.Println("[id: ", id, ", name: ", name, "]")

	}
}
