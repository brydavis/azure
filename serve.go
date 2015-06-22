package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Results struct {
	Query,
	Data string
}

func ListenAndServe(port int, conn *sql.DB) error {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		RootHandler(res, req, conn)
	})

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		return err
	}

	return nil
}

func RootHandler(res http.ResponseWriter, req *http.Request, conn *sql.DB) {
	b, _ := ioutil.ReadFile("views/index.html")
	t := template.New("")
	t, _ = t.Parse(string(b))

	if req.Method == "POST" {
		req.ParseForm()
		query := req.FormValue("query")

		rows, err := conn.Query(query)
		if err != nil {
			fmt.Println(err.Error(), "\n")
		}
		defer rows.Close()

		columns, _ := rows.Columns()
		count := len(columns)
		values := make([]interface{}, count)
		valuePtrs := make([]interface{}, count)

		var store []map[string]interface{}

		for rows.Next() {
			for i, _ := range columns {
				valuePtrs[i] = &values[i]
			}

			rows.Scan(valuePtrs...)

			row := make(map[string]interface{})
			for i, col := range columns {
				var v interface{}
				val := values[i]
				b, ok := val.([]byte)

				if ok {
					v = string(b)
				} else {
					v = val
				}
				row[col] = v
			}

			store = append(store, row)
		}

		b, _ := json.Marshal(store)
		t.Execute(res, Results{Query: query, Data: string(b)})

	} else {
		t.Execute(res, Results{Query: "", Data: ""})
	}

}
