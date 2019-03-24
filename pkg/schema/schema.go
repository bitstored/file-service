package schema

type Folder struct {
	ID          string
	ParentID    string
	OwnerID     string
	Name        string
	Created     string
	IsDeletable bool
}

type File struct {
	ID         string
	ParentID   string
	OwnerID    string
	Name       string
	Created    string
	Type       string
	IsWritable bool
	IsPrivate  bool
	ContentID  string
}

type Content struct {
	ID   string
	Data []byte
}

type FSLevel struct {
	RootID    string
	FilesID   []string
	FoldersID []string
}

type FSLevelDetailed struct {
	RootID  string
	Folders []Folder
	Files   []File
}

type EncriptionKey struct {
	Key      string
	EntityID string
}
