package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index1(data map[string]interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFiles("./src/hello_user.html")
		if err != nil {
			fmt.Printf("Error parse HTML file! \n")
			fmt.Println(err)
			return
		}

		err = t.Execute(w, data)
		if err != nil {
			fmt.Printf("Error execute template! \n")
			fmt.Println(err)
			return
		}
	}
}

func main() {
	var data = map[string]interface{}{
		"Name":     "Verdiansyah",
		"Age":      21,
		"Messages": "Iam Love Golang and Python",
	}

	var port string = fmt.Sprintf(":%d", 5050)
	http.HandleFunc("/", index1(data))

	fmt.Printf("Listening server at: http://localhost%s", port)
	http.ListenAndServe(port, nil)
}
