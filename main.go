package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
)

var debug = flag.Bool("debug", false, "enable debugging")

func main() {
	byts, _ := ioutil.ReadFile("config.json")
	var config map[string]interface{}
	if err := json.Unmarshal(byts, &config); err != nil {
		panic(err)
	}

	connString := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%d;database=%s",
		config["server"],
		config["user"],
		config["password"],
		int(config["port"].(float64)),
		config["database"],
	)

	flag.Parse()
	if *debug {
		fmt.Printf(" password:%s\n", config["password"])
		fmt.Printf(" port:%d\n", int(config["port"].(float64)))
		fmt.Printf(" server:%s\n", config["server"])
		fmt.Printf(" user:%s\n", config["user"])
		fmt.Printf(" connString:%s\n", connString)
	}

	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	go ListenAndServe(8080, conn)

	silent := false
	for !silent {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("azure ~> ")
		scanner.Scan()

		text := scanner.Text()
		textArray := strings.Split(text, " ")

		switch strings.ToLower(textArray[0]) {
		case "quit", "exit":
			os.Exit(1)
		case "clear":
			cmd, _ := exec.Command("clear").Output()
			fmt.Println(string(cmd))
		case "run":
			file, _ := ioutil.ReadFile("./sql/" + textArray[1])
			rows, err := conn.Query(string(file))
			if err != nil {
				// log.Fatal("Query Error: ", err.Error())
				fmt.Println(err.Error(), "\n")
				continue
			}
			defer rows.Close()

			columns, _ := rows.Columns()
			count := len(columns)
			values := make([]interface{}, count)
			valuePtrs := make([]interface{}, count)

			for rows.Next() {
				for i, _ := range columns {
					valuePtrs[i] = &values[i]
				}

				rows.Scan(valuePtrs...)
				store := make(map[string]interface{})
				for i, col := range columns {
					var v interface{}
					val := values[i]
					b, ok := val.([]byte)

					if ok {
						v = string(b)
					} else {
						v = val
					}
					store[col] = v
				}
				fmt.Println(store)
			}
			fmt.Println("\n")

		case "export":
			file, _ := ioutil.ReadFile("./sql/" + textArray[1])
			rows, err := conn.Query(string(file))
			if err != nil {
				// log.Fatal("Query Error: ", err.Error())
				fmt.Println(err.Error(), "\n")
				continue
			}
			defer rows.Close()

			columns, _ := rows.Columns()
			count := len(columns)
			values := make([]interface{}, count)
			valuePtrs := make([]interface{}, count)

			megastore := make([]map[string]interface{})

			for rows.Next() {
				for i, _ := range columns {
					valuePtrs[i] = &values[i]
				}

				rows.Scan(valuePtrs...)
				store := make(map[string]interface{})
				for i, col := range columns {
					var v interface{}
					val := values[i]
					b, ok := val.([]byte)

					if ok {
						v = string(b)
					} else {
						v = val
					}
					store[col] = v
				}
				megastore = append(megastore, store)
				// fmt.Println(store)
			}
			fmt.Println("\n")

			j, err := json.Marshal(megastore)
			ioutil.WriteFile(textArray[2], j, 0777)

		default:
			rows, err := conn.Query(text)
			if err != nil {
				// log.Fatal("Query Error: ", err.Error())
				fmt.Println(err.Error(), "\n")
				continue
			}
			defer rows.Close()

			columns, _ := rows.Columns()
			count := len(columns)
			values := make([]interface{}, count)
			valuePtrs := make([]interface{}, count)

			for rows.Next() {
				for i, _ := range columns {
					valuePtrs[i] = &values[i]
				}

				rows.Scan(valuePtrs...)
				store := make(map[string]interface{})
				for i, col := range columns {
					var v interface{}
					val := values[i]
					b, ok := val.([]byte)

					if ok {
						v = string(b)
					} else {
						v = val
					}
					store[col] = v
				}
				fmt.Println(store)
			}
			fmt.Println("\n")
		}
	}
}
