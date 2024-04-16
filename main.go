package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Age int `json:"age"`
	Phone string `json:"phone"`
}



func main ()  {

	var users = []User{}
	var idIncrement = 1
	
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		switch r.Method {
		case "GET":

			jsonData, err := json.Marshal(users)
			
			if err != nil {
				fmt.Println("Error Parse")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}

			w.Header().Set("Content-type", "application/json")
			w.Write(jsonData)

		case "POST":
			var user User

			err := json.NewDecoder(r.Body).Decode(&user)
			
			if err != nil {
				fmt.Println("Error Decode")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}

			user.Id = idIncrement

			idIncrement+=1

			users = append(users, user)

			jsonData, err := json.Marshal(user)
			
			if err != nil {
				fmt.Println("Error Parse")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}

			w.Header().Set("Content-type", "application/json")
			w.Write(jsonData)
		}

	})

	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)

	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}


}