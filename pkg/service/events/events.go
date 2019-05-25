package events

type DriveCreated struct {
	UserID  string
	DriveID string
	Error   error
}
type NewFileCreated struct {
	UserID string
	FileID string
	Error  error
}
type NewFolderCreated struct {
	UserID   string
	FolderID string
	Error    error
}
type FileContentUpdated struct {
	UserID string
	FileID string
	Error  error
}
type FileDeleted struct {
	UserID string
	FileID string
	Error  error
}
type FileRenamed struct {
	UserID string
	FileID string
	Error  error
}
type FileMoved struct {
	UserID string
	FileID string
	Error  error
}
type FileUploaded struct {
	UserID string
	FileID string
	Error  error
}
type FileShared struct {
	UserID string
	FileID string
	Error  error
}
type FileDownloaded struct {
	UserID       string
	FileID       string
	Name         string
	CreationDate string
	FileType     string
	Writable     bool
	Private      bool
	Content      []byte
	Error        error
}

type Failed struct {
	Error error
}
