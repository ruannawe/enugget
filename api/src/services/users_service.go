package services

import (
	"database/sql"
	// "github.com/cockroachdb/cockroach-go/crdb"
    _ "github.com/lib/pq"
	"log"
	"fmt"
)

func GetUsers() []byte {
	db, err := sql.Open("postgres", "postgres://admin01:Bleat@34781294@free-tier.gcp-us-central1.cockroachlabs.cloud:26257/crud?sslmode=verify-full&sslrootcert=/home/ubuntu/.postgresql/root.crt&options=--cluster%3Dtwin-mule-3790")
    if err != nil {
        log.Fatal("error connecting to the database: ", err)
    }

    defer db.Close()

	rows, err := db.Query(`
		select
			json_agg(u.id),
			json_agg(u.name),
			json_agg(u.age)
		from
			users u
	`)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
        var response string

        if err := rows.Scan(&response); err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%s\n", response)
    }

	jsonData := []byte(`[{"id": 1, "name": "Ruan", "age": "12"}, {"id": "2", "name": "Ediane", "age": "37"}]`)
	return jsonData
}

func GetUser() string {
	return "{id: 1, name: 'Ruan', age: 12}"
}

func CreateUser() string {
	return "{id: 1, name: 'Ruan', age: 12}"
}

func UpdateUser() string {
	return "{id: 1, name: 'Ruan', age: 12}"
}

func DestroyUser() {

}
