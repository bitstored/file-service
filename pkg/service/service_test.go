package service

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/bitstored/file-service/pb"
	"github.com/bitstored/file-service/pkg/repo"
)

func TestService_CreateNewFile(t *testing.T) {
	repo, _ := repo.NewFileRepository()

	type args struct {
		ctx       context.Context
		userID    string
		fileName  string
		parentID  string
		fileType  string
		writable  bool
		private   bool
		content   []byte
		secretKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TESTOK",
			args: args{
				ctx:       context.Background(),
				userID:    "uid:1",
				fileName:  "file.txt",
				parentID:  "fdid:vKN+caR6QBIz3ldzzzx7og",
				fileType:  "txt",
				writable:  true,
				private:   false,
				content:   []byte("ana are mere"),
				secretKey: "ana",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: repo,
			}
			got, err := s.CreateNewFile(tt.args.ctx, tt.args.userID, tt.args.fileName, tt.args.parentID, tt.args.fileType, tt.args.writable, tt.args.private, tt.args.content, tt.args.secretKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.CreateNewFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestService_DownloadFile(t *testing.T) {
	repo, _ := repo.NewFileRepository()
	type args struct {
		ctx                 context.Context
		userID              string
		fileID              string
		secretKey           string
		watermarkingPhrase  string
		steganographyPhrase string
		watermarkingImage   string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.File
		wantErr bool
	}{
		{
			name: "TESTOK",
			args: args{
				ctx:       context.Background(),
				userID:    "uid:1",
				fileID:    "flid:oYMr34YkTi1/ZCxTqlaGTA",
				secretKey: "ana",
			},
			want: &pb.File{
				Content: []byte("ana are mere"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: repo,
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
