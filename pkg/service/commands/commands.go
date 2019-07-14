package commands

type CreateDrive struct {
	UserID string
}

func (c CreateDrive) IsValid() bool {
	return len(c.UserID) > 0
}

type CreateNewFile struct {
	UserID         string
	ParentID       string
	CreationDate   string
	Name           string
	FileType       string
	Writable       bool
	Private        bool
	Content        []byte
	InitialSize    int64
	CompressedSize int64
}

func (c CreateNewFile) IsValid() bool {
	return len(c.UserID) > 0 && len(c.ParentID) > 0 && len(c.Name) > 0 && len(c.FileType) != 0
}

type CreateNewFolder struct {
	UserID       string
	Name         string
	ParentID     string
	CreationDate string
}

func (c CreateNewFolder) IsValid() bool {
	return len(c.UserID) > 0 && len(c.ParentID) > 0 && len(c.Name) > 0
}

type UpdateFileContent struct {
	UserID         string
	FileID         string
	Data           []byte
	InitialSize    int64
	CompressedSize int64
}

func (c UpdateFileContent) IsValid() bool {
	return len(c.UserID) > 0 && len(c.FileID) > 0 && c.Data != nil
}

type DeleteFile struct {
	UserID string
	FileID string
}

func (c DeleteFile) IsValid() bool {
	return len(c.UserID) > 0 && len(c.FileID) > 0
}

type RenameFile struct {
	UserID string
	FileID string
	Name   string
}

func (c RenameFile) IsValid() bool {
	return len(c.UserID) > 0 && len(c.FileID) > 0 && len(c.Name) > 0
}

type MoveFile struct {
	UserID        string
	FileID        string
	SourceID      string
	DestinationID string
	Copy          bool
}

func (c MoveFile) IsValid() bool {
	return len(c.UserID) > 0 && len(c.FileID) > 0 && len(c.SourceID) > 0 && len(c.DestinationID) > 0
}

type UploadFile struct {
	UserID         string
	ParentID       string
	CreationDate   string
	Name           string
	FileType       string
	Writable       bool
	Private        bool
	Content        []byte
	InitialSize    int64
	CompressedSize int64
}

func (c UploadFile) IsValid() bool {
	return len(c.UserID) > 0 && len(c.ParentID) > 0 && len(c.Name) > 0 && len(c.Content) > 0
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

func (c DownloadFile) IsValid() bool {
	return len(c.UserID) > 0 && len(c.FileID) > 0
}
