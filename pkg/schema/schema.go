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
	ID             string
	ParentID       string
	OwnerID        string
	Name           string
	Created        string
	Type           string
	IsWritable     bool
	IsPrivate      bool
	ContentID      string
	InitialSize    int64
	CompressedSize int64
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
	FileID   string
	Key      string
	EntityID string
}

type SharedFile struct {
	ID       string
	OwnerID  string
	TargetID string
}

type UserData struct {
	DriveID string
}
