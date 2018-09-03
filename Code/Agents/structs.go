package main

type File struct {
    ID        string   `json:"id,omitempty"`
    Filename  string   `json:"filename,omitempty"`
    MD5       string   `json:"md5,omitempty"`
    Content   *Content `json:"content,omitempty"`
}
type Content  struct {
    Location  string `json:"location,omitempty"`
    Data      string `json:"data,omitempty"`
}

// -------------- Better implementation --------------

type File2 struct {
    Global_ID     string        `json:"global_id,omitempty"`
    Filename      string        `json:"filename,omitempty"`
    Filepath      string        `json:"filepath,omitempty"`
    Storage_Loc   string        `json:"storage_loc,omitempty"`
    MD5           string        `json:"md5,omitempty"`
    Permissions   string        `json:"permissions,omitempty"`
    Overwrite     bool          `json:"overwrite,omitempty"`
    Create_Dir    bool          `json:"create_dir,omitempty"`
    Has_Sym_Link  bool          `json:"has_sym_link,omitempty"`
    Sym_Link_Info Sym_Link_Info `json:"Sym_Link_Info,omitempty"`
}

type Sym_Link_Info  struct {
    Overwrite_Existing    bool    `json:"overwrite,omitempty"`
    Filename              string  `json:"filename,omitempty"`
    Filepath              string  `json:"filepath,omitempty"`
    Create_Dir            bool    `json:"create_dir,omitempty"`
}

/*

What can we need to know:
  Filename (just file name, not file path)
  Filepath (We dont actually need the content)
  Md5
  Permission/owner/group
  Overwrite existing file flag
  Create non-existent directory flag
  A link to a place to obtain the file (unsafe to use md5 for this
  Some kind of id local to the filesystem (we may not need this is we know filename)

All of the files are to be held in an in-memory and in-disk list format
should an incident occur while operations are occuring, the same operations can re-occur
*/
