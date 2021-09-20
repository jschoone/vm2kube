package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	port := os.Args[4]
	http.HandleFunc("/", HelloHandler)
	fmt.Printf("Server started at port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	appname := path.Base(os.Args[0])
	log.Default()
	fmt.Fprintf(w, "Moin, ich bin %s auf %s\n", appname, hostname)

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", os.Args[1], os.Args[2], os.Args[5], os.Args[3]))
	if err != nil {
		panic(err)
	} else if err = db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS moin (id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, hostname TEXT NOT NULL, appname TEXT NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		panic(err)
	}

	// Create
	_, err = db.Exec("INSERT INTO moin (hostname, appname) VALUES (?,?)", hostname, appname)
	if err != nil {
		panic(err)
	}

	// Read
	rows, err := db.Query("SELECT created_at FROM moin WHERE hostname = ? AND appname = ?", hostname, appname)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var (
			created string
		)
		if err := rows.Scan(&created); err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%s\n", created)
	}
}
