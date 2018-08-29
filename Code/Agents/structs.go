package main

type File struct {
    ID        string   `json:"id,omitempty"`
    Filename string   `json:"filename,omitempty"`
    MD5  string   `json:"md5,omitempty"`
    Content   *Content `json:"content,omitempty"`
}
type Content struct {
    Location  string `json:"location,omitempty"`
    Data string `json:"data,omitempty"`
}
