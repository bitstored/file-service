package server

import "github.com/dianabejan/file-service/server/pb"

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetFolderContent(ctx context.Context, *pb.GetFolderContentRequest) (*GetFolderContentResponse, error) {
	return nil, nil
}
func (s *Server) GetFileContent(ctx context.Context, *pb.GetFileContentRequest) (*GetFileContentResponse, error) {
	return nil, nil
}
func (s *Server) GetFileTree(ctx context.Context, *pb.GetFileTreeRequest) (*GetFileTreeResponse, error) {
	return nil, nil
}
func (s *Server) UpdateFileContent(ctx context.Context, *pb.UpdateFileContentRequest) (*UpdateFileContentResponse, error) {
	return nil, nil
}
func (s *Server) CreateDrive(ctx context.Context, *pb.CreateDriveRequest) (*CreateDriveResponse, error) {
	return nil, nil
}
func (s *Server) CreateNewFile(ctx context.Context, *pb.CreateNewFileRequest) (*CreateNewFileResponse, error) {
	return nil, nil
}
func (s *Server) DeleteFile(ctx context.Context, *pb.DeleteFileRequest) (*DeleteFileResponse, error) {
	return nil, nil
}
func (s *Server) RenameFile(ctx context.Context, *pb.RenameFileRequest) (*RenameFileResponse, error) {
	return nil, nil
}
func (s *Server) MoveFile(ctx context.Context, *pb.MoveFileRequest) (*UploadFileResponse, error) {
	return nil, nil
}
func (s *Server) UploadFile(ctx context.Context, *pb.UploadFileRequest) (*UploadFileResponse, error) {
	return nil, nil
}
