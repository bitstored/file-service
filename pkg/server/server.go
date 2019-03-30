package server

import (
	"context"
	"github.com/bitstored/file-service/pb"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}
func (s *Server) CreateDrive(ctx context.Context, in *pb.CreateDriveRequest) (*pb.CreateDriveResponse, error) {
	return nil, nil
}
func (s *Server) CreateNewFolder(context.Context, *pb.CreateNewFolderRequest) (*pb.CreateNewFolderResponse, error) {
	return nil, nil
}
func (s *Server) CreateNewFile(ctx context.Context, in *pb.CreateNewFileRequest) (*pb.CreateNewFileResponse, error) {
	return nil, nil
}
func (s *Server) GetFolderContent(ctx context.Context, in *pb.GetFolderContentRequest) (*pb.GetFolderContentResponse, error) {
	return nil, nil
}
func (s *Server) GetFileContent(ctx context.Context, in *pb.GetFileContentRequest) (*pb.GetFileContentResponse, error) {
	return nil, nil
}
func (s *Server) GetFileTree(ctx context.Context, in *pb.GetFileTreeRequest) (*pb.GetFileTreeResponse, error) {
	return nil, nil
}
func (s *Server) UpdateFileContent(ctx context.Context, in *pb.UpdateFileContentRequest) (*pb.UpdateFileContentResponse, error) {
	return nil, nil
}
func (s *Server) DeleteFile(ctx context.Context, in *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	return nil, nil
}
func (s *Server) RenameFile(ctx context.Context, in *pb.RenameFileRequest) (*pb.RenameFileResponse, error) {
	return nil, nil
}
func (s *Server) MoveFile(ctx context.Context, in *pb.MoveFileRequest) (*pb.UploadFileResponse, error) {
	return nil, nil
}
func (s *Server) UploadFile(ctx context.Context, in *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	return nil, nil
}
func (s *Server) ShareFile(context.Context, *pb.ShareFileRequest) (*pb.ShareFileResponse, error) {
	return nil, nil
}
func (s *Server) DownloadFile(context.Context, *pb.DownloadFileRequest) (*pb.DownloadFileResponse, error) {
	return nil, nil
}
