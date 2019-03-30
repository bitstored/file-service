package commands

type CreateDrive struct {
	UserID  string
	DriveID string
}
type CreateNewFile struct {
	UserID       string
	FileID       string
	ParentID     string
	CreationDate string
	FileType     string
	Writable     bool
	Private      bool
}
type CreateNewFolder struct {
	UserID       string
	FolderID     string
	ParentID     string
	CreationDate string
}
type UpdateFileContent struct {
	UserID string
	FileID string
	Data   string
}
type DeleteFile struct {
	UserID string
	FileID string
}
type RenameFile struct {
	UserID string
	FileID string
	Name   string
}
type MoveFile struct {
	UserID        string
	FileID        string
	DestinationID string
}
type UploadFile struct {
	UserID       string
	FileID       string
	ParentID     string
	CreationDate string
	FileType     string
	Writable     bool
	Private      bool
	Content      string
}
type ShareFile struct {
	UserID        string
	FileID        string
	BeneficiarIDs []string
}
type DownloadFile struct {
	UserID string
	FileID string
}
