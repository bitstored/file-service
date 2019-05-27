package commands_test

import "testing"

func TestCreateDrive_IsValid(t *testing.T) {
	type fields struct {
		UserID string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateDrive{
				UserID: tt.fields.UserID,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("CreateDrive.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateNewFile_IsValid(t *testing.T) {
	type fields struct {
		UserID       string
		ParentID     string
		CreationDate string
		Name         string
		FileType     string
		Writable     bool
		Private      bool
		Content      []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateNewFile{
				UserID:       tt.fields.UserID,
				ParentID:     tt.fields.ParentID,
				CreationDate: tt.fields.CreationDate,
				Name:         tt.fields.Name,
				FileType:     tt.fields.FileType,
				Writable:     tt.fields.Writable,
				Private:      tt.fields.Private,
				Content:      tt.fields.Content,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("CreateNewFile.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateNewFolder_IsValid(t *testing.T) {
	type fields struct {
		UserID       string
		Name         string
		ParentID     string
		CreationDate string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CreateNewFolder{
				UserID:       tt.fields.UserID,
				Name:         tt.fields.Name,
				ParentID:     tt.fields.ParentID,
				CreationDate: tt.fields.CreationDate,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("CreateNewFolder.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateFileContent_IsValid(t *testing.T) {
	type fields struct {
		UserID string
		FileID string
		Data   []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := UpdateFileContent{
				UserID: tt.fields.UserID,
				FileID: tt.fields.FileID,
				Data:   tt.fields.Data,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("UpdateFileContent.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFile_IsValid(t *testing.T) {
	type fields struct {
		UserID string
		FileID string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := DeleteFile{
				UserID: tt.fields.UserID,
				FileID: tt.fields.FileID,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("DeleteFile.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRenameFile_IsValid(t *testing.T) {
	type fields struct {
		UserID string
		FileID string
		Name   string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := RenameFile{
				UserID: tt.fields.UserID,
				FileID: tt.fields.FileID,
				Name:   tt.fields.Name,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("RenameFile.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMoveFile_IsValid(t *testing.T) {
	type fields struct {
		UserID        string
		FileID        string
		SourceID      string
		DestinationID string
		Copy          bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := MoveFile{
				UserID:        tt.fields.UserID,
				FileID:        tt.fields.FileID,
				SourceID:      tt.fields.SourceID,
				DestinationID: tt.fields.DestinationID,
				Copy:          tt.fields.Copy,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("MoveFile.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUploadFile_IsValid(t *testing.T) {
	type fields struct {
		UserID       string
		ParentID     string
		CreationDate string
		Name         string
		FileType     string
		Writable     bool
		Private      bool
		Content      []byte
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := UploadFile{
				UserID:       tt.fields.UserID,
				ParentID:     tt.fields.ParentID,
				CreationDate: tt.fields.CreationDate,
				Name:         tt.fields.Name,
				FileType:     tt.fields.FileType,
				Writable:     tt.fields.Writable,
				Private:      tt.fields.Private,
				Content:      tt.fields.Content,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("UploadFile.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDownloadFile_IsValid(t *testing.T) {
	type fields struct {
		UserID string
		FileID string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := DownloadFile{
				UserID: tt.fields.UserID,
				FileID: tt.fields.FileID,
			}
			if got := c.IsValid(); got != tt.want {
				t.Errorf("DownloadFile.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
