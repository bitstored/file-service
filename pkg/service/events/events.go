package events

type DriveCreated struct {
	UserID  string
	DriveID string
}
type NewFileCreated struct {
	UserID string
	FileID string
}
type NewFolderCreated struct {
	UserID   string
	FolderID string
}
type FileContentUpdated struct {
	UserID string
	FileID string
}
type FileDeleted struct {
	UserID string
	FileID string
}
type FileRenamed struct {
	UserID string
	FileID string
}
type FileMoved struct {
	UserID string
	FileID string
}
type FileUploaded struct {
	UserID string
	FileID string
}
type FileShared struct {
	UserID string
	FileID string
}
type FileDownloaded struct {
	UserID       string
	FileID       string
	CreationDate string
	FileType     string
	Writable     bool
	Private      bool
	Content      string
}
