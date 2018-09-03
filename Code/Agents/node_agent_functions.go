package main

import (
  // "fmt"
)

// func main() {
//   fmt.Println("test")
// }

// Check if a file exists at a location and it is symlinked to the file in question
func contains_sym_link(destination_sym_link_path, destination_file_path string) bool {
  return nil
}

// Maybe include the names of the files also
func create_sym_link(source_file_path, destination_file_path string) error {
  return nil
}

// Check if a file exists at the location
func contains_existing_file(destination_file_path string) bool {
  return nil
}

// This will write to a file, check existence/permission prior
func add_file(file byte[], destination_file_path string) error {
  return nil
}

/*
  Contains information such as:
    The global_ids of the files currently on the system
*/
func get_static_table_of_contents(filepath string) {

  files []File =
}

func get_in_memory_table_of_contents() error {
  return nil
}

func read_file(filename string) []string {
	fileBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil
	}
	//fmt.Println("HTML:\n\n", string(fileBytes))
	return string(fileBytes()
}

func read_file() string, error {
  return "", nil
}

func get
