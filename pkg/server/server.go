package server

import (
	"context"
	"github.com/bitstored/file-service/src/server/pb"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetFolderContent(ctx context.Context, in *pb.GetFolderContentRequest) (*GetFolderContentResponse, error) {
	return nil, nil
}
func (s *Server) GetFileContent(ctx context.Context, in *pb.GetFileContentRequest) (*GetFileContentResponse, error) {
	return nil, nil
}
func (s *Server) GetFileTree(ctx context.Context, in *pb.GetFileTreeRequest) (*GetFileTreeResponse, error) {
	return nil, nil
}
func (s *Server) UpdateFileContent(ctx context.Context, in *pb.UpdateFileContentRequest) (*UpdateFileContentResponse, error) {
	return nil, nil
}
func (s *Server) CreateDrive(ctx context.Context, in *pb.CreateDriveRequest) (*CreateDriveResponse, error) {
	return nil, nil
}
func (s *Server) CreateNewFile(ctx context.Context, in *pb.CreateNewFileRequest) (*CreateNewFileResponse, error) {
	return nil, nil
}
func (s *Server) DeleteFile(ctx context.Context, in *pb.DeleteFileRequest) (*DeleteFileResponse, error) {
	return nil, nil
}
func (s *Server) RenameFile(ctx context.Context, in *pb.RenameFileRequest) (*RenameFileResponse, error) {
	return nil, nil
}
func (s *Server) MoveFile(ctx context.Context, in *pb.MoveFileRequest) (*UploadFileResponse, error) {
	return nil, nil
}
func (s *Server) UploadFile(ctx context.Context, in *pb.UploadFileRequest) (*UploadFileResponse, error) {
	return nil, nil
}
