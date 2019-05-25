package server

import (
	"context"

	"github.com/bitstored/file-service/pb"
	"github.com/bitstored/file-service/pkg/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	Service *service.Service
}

func NewServer(s *service.Service) *Server {
	return &Server{s}
}
func (s *Server) CreateDrive(ctx context.Context, in *pb.CreateDriveRequest) (*pb.CreateDriveResponse, error) {

	// TODO Verify auth

	rsp, err := s.Service.CreateDrive(ctx, in.UserId)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	response := &pb.CreateDriveResponse{DriveId: rsp}

	return response, nil
}

func (s *Server) CreateNewFolder(ctx context.Context, in *pb.CreateNewFolderRequest) (*pb.CreateNewFolderResponse, error) {
	folder := in.GetFolder()
	if folder.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "folder name is invalid")
	}
	if folder.GetParentIdentifier() == "" {
		return nil, status.Error(codes.InvalidArgument, "destination folder is invalid")
	}
	folderID, err := s.Service.CreateNewFolder(ctx, in.GetUserId(), folder.GetName(), folder.GetParentIdentifier())
	if err != nil {
		return nil, err
	}
	rsp := &pb.CreateNewFolderResponse{
		ResponseCode:    codes.OK.String(),
		ResponseMessage: "Folder created with success.",
		FolderId:        folderID,
	}
	return rsp, nil
}
func (s *Server) CreateNewFile(ctx context.Context, in *pb.CreateNewFileRequest) (*pb.CreateNewFileResponse, error) {
	file := in.File
	if file.Content == nil {
		return nil, status.Error(codes.InvalidArgument, "File content can't be nil")
	}

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
func (s *Server) ShareFile(ctx context.Context, in *pb.ShareFileRequest) (*pb.ShareFileResponse, error) {
	return nil, nil
}
func (s *Server) DownloadFile(ctx context.Context, in *pb.DownloadFileRequest) (*pb.DownloadFileResponse, error) {
	return nil, nil
}
