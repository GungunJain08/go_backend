package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type Student struct {
	Name   string
	RollNo int32
}

func main() {
	Port := ":8000"
	var uri string
	uri = "MONGODB_URI"
	if uri == "" {
		fmt.Println("uri is empty")
	}
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	coll := client.Database("students_data").Collection("Students")
	newStudent := Student{Name: "Gunnu", RollNo: 82}
	result, err := coll.InsertOne(context.TODO(), newStudent)
	if err != nil {
		fmt.Println("error occured")
	}
	fmt.Println(result, result.InsertedID)
	http.HandleFunc("/gunnu", gunnu_diva)
	http.HandleFunc("/file", createFile)
	fmt.Println("Server running on", Port)
	http.ListenAndServe(Port, nil)

}
func createFile(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create("gunnu.txt")
	if err != nil {
		fmt.Fprintln(w, "there is a error")
	}
	fmt.Fprintln(w, "file created successfully"+file.Name())
}

func gunnu_diva(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "gunnu is the diva")
}
