package main

import (
  // "fmt"
  "io/ioutil"
  "log"
  "strings"
)

// func main() {
//   fmt.Println("test")
// }

// Check if a file exists at a location and it is symlinked to the file in question
func contains_sym_link(destination_sym_link_path, destination_file_path string) bool {
  return false
}

// Maybe include the names of the files also
func create_sym_link(source_file_path, destination_file_path string) error {
  return nil
}

// Check if a file exists at the location
func contains_existing_file(destination_file_path string) bool {
  return false
}

// This will write to a file, check existence/permission prior
func add_file(file []byte, destination_file_path string) error {
  return nil
}

/*
  Contains information such as:
    The global_ids of the files currently on the system
*/
func get_static_table_of_contents(filepath string) {

  var files []string
  files = make([]string, 0)
  file_contents, err := read_file(filepath)
  c(err)
  file_lines := strings.Split(file_contents, "\n")
  for _, line := range file_lines {
    files = append(files, line)
  }

  // return files
}

func get_in_memory_table_of_contents() error {
  return nil
}

// func read_file(filename string) []string {
// 	fileBytes, err := ioutil.ReadFile(filename)
// 	c(err)
// 	//fmt.Println("HTML:\n\n", string(fileBytes))
// 	return string(fileBytes)
// }

func read_file(filename string) (string, error) {
  fileBytes, err := ioutil.ReadFile(filename)
	if err != nil {
    return "", err
  }
	//fmt.Println("HTML:\n\n", string(fileBytes))
	return string(fileBytes), nil
}

func c(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

// func get
