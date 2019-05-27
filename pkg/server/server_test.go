package server

import (
	"context"
	"reflect"
	"testing"

	"github.com/bitstored/file-service/pb"
	"github.com/bitstored/file-service/pkg/service"
	"github.com/stretchr/testify/require"
)

func TestNewServer(t *testing.T) {
	type args struct {
		s *service.Service
	}
	tests := []struct {
		name string
		args args
		want *Server
	}{
		{
			"Nil service",
			args{nil},
			&Server{nil},
		},
		{
			"Good service",
			args{&service.Service{}},
			&Server{&service.Service{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServer(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_CreateDrive(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.CreateDriveRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CreateDriveResponse
		wantErr bool
	}{
		{
			name: "create ok",
			fields: fields{
				Service: service.NewService(),
			},
			args: args{
				ctx: context.TODO(),
				in: &pb.CreateDriveRequest{
					UserId: "uid:dianaaremere",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.CreateDrive(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateDrive() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.NotEmpty(t, got.GetDriveId())
		})
	}
}

func TestServer_CreateNewFolder(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.CreateNewFolderRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CreateNewFolderResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.CreateNewFolder(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateNewFolder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CreateNewFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_CreateNewFile(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.CreateNewFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CreateNewFileResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.CreateNewFile(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.CreateNewFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.CreateNewFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetFolderContent(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.GetFolderContentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetFolderContentResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.GetFolderContent(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetFolderContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetFolderContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetFileContent(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.GetFileContentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetFileContentResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.GetFileContent(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetFileContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetFileContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_GetFileTree(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.GetFileTreeRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.GetFileTreeResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.GetFileTree(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.GetFileTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.GetFileTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_UpdateFileContent(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.UpdateFileContentRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UpdateFileContentResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.UpdateFileContent(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.UpdateFileContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.UpdateFileContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DeleteFile(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.DeleteFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DeleteFileResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.DeleteFile(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DeleteFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_RenameFile(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.RenameFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.RenameFileResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.RenameFile(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.RenameFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.RenameFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_MoveFile(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.MoveFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UploadFileResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.MoveFile(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.MoveFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.MoveFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_UploadFile(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		ctx context.Context
		in  *pb.UploadFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.UploadFileResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.UploadFile(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.UploadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.UploadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_ShareFile(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		in0 context.Context
		in1 *pb.ShareFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ShareFileResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.ShareFile(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.ShareFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.ShareFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_DownloadFile(t *testing.T) {
	type fields struct {
		Service *service.Service
	}
	type args struct {
		in0 context.Context
		in1 *pb.DownloadFileRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.DownloadFileResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Service: tt.fields.Service,
			}
			got, err := s.DownloadFile(tt.args.in0, tt.args.in1)
			if (err != nil) != tt.wantErr {
				t.Errorf("Server.DownloadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Server.DownloadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
