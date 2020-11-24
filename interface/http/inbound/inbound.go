package inbound

import (
	"encoding/json"
	"fmt"
	"golang-clean-architecture/domain"
	"golang-clean-architecture/service"
	"net/http"
)

func InitServer(){
	http.HandleFunc("/saveStudent", saveStudentHandler)
	http.HandleFunc("/getAllStudent",getAllStudentHandler)

	err:=http.ListenAndServe(":8080",nil);if err!=nil{
		panic(err)
	}

}
func saveStudentHandler(w http.ResponseWriter, r *http.Request)  {
	var student domain.Student
	err:=json.NewDecoder(r.Body).Decode(&student);if err!=nil{
		fmt.Println("Decoding error:",err)
		w.Write([]byte(err.Error()))
		return
	}
	s:=service.RegisterStudent(student);if s{
		w.Write([]byte("successfully saved"))
		return
	}
	w.Write([]byte("failed to save"))
}
func getAllStudentHandler(w http.ResponseWriter, r *http.Request){
	re,_:=json.Marshal(service.GetAllStudent())
	w.Write(re)
}
