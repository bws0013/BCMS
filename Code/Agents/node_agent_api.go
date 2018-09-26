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
  log.Fatal(http.ListenAndServe(":8889", router))
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

func CreateFile(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  var file File
  _ = json.NewDecoder(r.Body).Decode(&file)
  file.ID = params["id"]
  files = append(files, file)
  json.NewEncoder(w).Encode(files)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  for index, item := range files {
    if item.ID == params["global_id"] {
      files = append(files[:index], files[index+1:]...)
      break
    }
    json.NewEncoder(w).Encode(files)
  }
}

func addFiles() {
  files = append(files, File{ID: "1", Filename: "hello.java", MD5: "598a9fc1486631bb148e511c4ea879b8", Content: &Content{Location: "/usr/local/java/cool/", Data: "hello.java content stuff"}})
  files = append(files, File{ID: "2", Filename: "haproxy.cfg", MD5: "3edde1531eec964a3617e4b7855489ac", Content: &Content{Location: "/home/", Data: "haproxy.cfb content stuff"}})
  files = append(files, File{ID: "3", Filename: "world.go", MD5: "5e399fe8c999e687153cda50dce4280f"})
}
