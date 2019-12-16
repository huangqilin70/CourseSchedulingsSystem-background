package main

import "net/http"

func testPing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
	return
}

func init() {
	http.HandleFunc("/Courses", testPing)
	http.HandleFunc("/Courses/create", testPing)
	http.HandleFunc("/Courses/Delete", testPing)
	http.HandleFunc("/Courses/Edit", testPing)
	http.HandleFunc("/Enrollments/Delete", testPing)
	http.HandleFunc("/Enrollments/Create", testPing)
	http.HandleFunc("/Students", testPing)
	http.HandleFunc("/Students/Delete", testPing)
	http.HandleFunc("/Students/create", testPing)
	http.HandleFunc("/Students/Edit", testPing)
}

func main() {
	http.ListenAndServe("127.0.0.1:8004", nil)
}
