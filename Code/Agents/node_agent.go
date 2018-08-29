package main

import (
  // "fmt"
  "encoding/json"
  "log"
  "net/http"
  "github.com/gorilla/mux"
)

var (
  files []File
)

func main() {
  router := mux.NewRouter()
  addFiles()
  router.HandleFunc("/files", GetFiles).Methods("GET")
  router.HandleFunc("/files/{id}", GetFile).Methods("GET")
  router.HandleFunc("/files/{id}", CreateFile).Methods("POST")
  router.HandleFunc("/files/{id}", DeleteFile).Methods("DELETE")
  log.Fatal(http.ListenAndServe(":8000", router))
}

func GetFiles(w http.ResponseWriter, r *http.Request) {
  json.NewEncoder(w).Encode(files)
}
func GetFile(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for _, item := range files {
    if item.ID == params["id"] {
      json.NewEncoder(w).Encode(item)
    }
  }
}
func CreateFile(w http.ResponseWriter, r *http.Request) {}
func DeleteFile(w http.ResponseWriter, r *http.Request) {}

func addFiles() {
  files = append(files, File{ID: "1", Filename: "hello.java", MD5: "598a9fc1486631bb148e511c4ea879b8", Content: &Content{Location: "/usr/local/java/cool/", Data: "hello.java content stuff"}})
  files = append(files, File{ID: "2", Filename: "haproxy.cfg", MD5: "3edde1531eec964a3617e4b7855489ac", Content: &Content{Location: "/home/", Data: "haproxy.cfb content stuff"}})
  files = append(files, File{ID: "3", Filename: "world.go", MD5: "5e399fe8c999e687153cda50dce4280f"})
}

// This is a node that runs on a host machine
// It monitors for changes in files that it is supposed to monitor
// It also keeps a persistent list of changes that needs to be made

/*
  Change process
    Recieve change to be added
    Add change to change list
    Every N seconds close the list and announce that changes need to wait
    For each entry in the list
      Create directories if the option has been set
      Retrieve intended file by communicating with one of the master nodes slaves
      Delete/Change existing file
      Store the md5 in persistent and dynamic storage
*/