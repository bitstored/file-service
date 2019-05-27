package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/bitstored/file-service/pb"
	"github.com/bitstored/file-service/pkg/repo"
)

func TestNewService(t *testing.T) {
	tests := []struct {
		name string
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateDrive(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx    context.Context
		userID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.CreateDrive(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateDrive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.CreateDrive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateNewFile(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx       context.Context
		userID    string
		fileName  string
		parentID  string
		fileType  string
		writable  bool
		private   bool
		content   []byte
		secretKey []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.CreateNewFile(tt.args.ctx, tt.args.userID, tt.args.fileName, tt.args.parentID, tt.args.fileType, tt.args.writable, tt.args.private, tt.args.content, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateNewFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.CreateNewFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_CreateNewFolder(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx        context.Context
		userID     string
		folderName string
		parentID   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.CreateNewFolder(tt.args.ctx, tt.args.userID, tt.args.folderName, tt.args.parentID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateNewFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.CreateNewFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_UpdateFileContent(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx       context.Context
		userID    string
		fileID    string
		fileType  string
		secretKey []byte
		content   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			if err := s.UpdateFileContent(tt.args.ctx, tt.args.userID, tt.args.fileID, tt.args.fileType, tt.args.secretKey, tt.args.content); (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateFileContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_DeleteFile(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx    context.Context
		userID string
		fileID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			if err := s.DeleteFile(tt.args.ctx, tt.args.userID, tt.args.fileID); (err != nil) != tt.wantErr {
				t.Errorf("Service.DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_RenameFile(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx     context.Context
		userID  string
		fileID  string
		newName string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			if err := s.RenameFile(tt.args.ctx, tt.args.userID, tt.args.fileID, tt.args.newName); (err != nil) != tt.wantErr {
				t.Errorf("Service.RenameFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_MoveFile(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx           context.Context
		userID        string
		fileID        string
		sourceID      string
		destinationID string
		copy          bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			if err := s.MoveFile(tt.args.ctx, tt.args.userID, tt.args.fileID, tt.args.sourceID, tt.args.destinationID, tt.args.copy); (err != nil) != tt.wantErr {
				t.Errorf("Service.MoveFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestService_UploadFile(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx       context.Context
		userID    string
		fileName  string
		parentID  string
		fileType  string
		writable  bool
		private   bool
		content   []byte
		secretKey []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, got1, err := s.UploadFile(tt.args.ctx, tt.args.userID, tt.args.fileName, tt.args.parentID, tt.args.fileType, tt.args.writable, tt.args.private, tt.args.content, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.UploadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Service.UploadFile() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Service.UploadFile() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestService_DownloadFile(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx                 context.Context
		userID              string
		fileID              string
		secretKey           []byte
		watermarkingPhrase  string
		steganographyPhrase string
		watermarkingImage   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.DownloadFile(tt.args.ctx, tt.args.userID, tt.args.fileID, tt.args.secretKey, tt.args.watermarkingPhrase, tt.args.steganographyPhrase, tt.args.watermarkingImage)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.DownloadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.DownloadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_encodeMessage(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx     context.Context
		target  []byte
		message []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.encodeMessage(tt.args.ctx, tt.args.target, tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.encodeMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.encodeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_decodeMessage(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx    context.Context
		target []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.decodeMessage(tt.args.ctx, tt.args.target)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.decodeMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.decodeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_watermarkWithImage(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx    context.Context
		target []byte
		data   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.watermarkWithImage(tt.args.ctx, tt.args.target, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.watermarkWithImage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.watermarkWithImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_watermarkWithText(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx    context.Context
		target []byte
		data   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.watermarkWithText(tt.args.ctx, tt.args.target, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.watermarkWithText() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.watermarkWithText() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_compress(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx      context.Context
		fileType string
		data     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.compress(tt.args.ctx, tt.args.fileType, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.compress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.compress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_decompress(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx      context.Context
		fileType string
		data     []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.decompress(tt.args.ctx, tt.args.fileType, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.decompress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.decompress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_encrypt(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx       context.Context
		data      []byte
		secretKey []byte
		userID    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.encrypt(tt.args.ctx, tt.args.data, tt.args.secretKey, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_decrypt(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx       context.Context
		data      []byte
		secretKey []byte
		userID    string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.decrypt(tt.args.ctx, tt.args.data, tt.args.secretKey, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetFolderContent(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx      context.Context
		userID   string
		folderID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.FSLevel
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.GetFolderContent(tt.args.ctx, tt.args.userID, tt.args.folderID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetFolderContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetFolderContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetFileContent(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx       context.Context
		userID    string
		fileID    string
		fileType  string
		secretKey []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.File
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			got, err := s.GetFileContent(tt.args.ctx, tt.args.userID, tt.args.fileID, tt.args.fileType, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.GetFileContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetFileContent() = %v, want %v", got, tt.want)
			}
		})
	}
}
