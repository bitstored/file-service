package server

import (
	"context"
	"fmt"
	"strings"

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

	var name string
	name = strings.Trim(file.GetName(), " ")
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "File name can't be empty")
	}

	if len(in.GetSecretPhrase()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Secret phrase can't be empty")
	}

	var Type string
	switch file.GetFileType() {
	case pb.Type_IMAGE:
		Type = "PNG"
	case pb.Type_PDF:
		Type = "PDF"
	default:
		Type = "TXT"
	}

	identifier, err := s.Service.CreateNewFile(ctx, in.GetUserId(), name, file.GetParentIdentifier(), Type, file.Writable, file.Private, file.Content, in.SecretPhrase)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.CreateNewFileResponse{
		ResponseCode:    codes.OK.String(),
		ResponseMessage: "File created with success",
		FileId:          identifier,
	}

	return resp, nil
}

func (s *Server) GetFolderContent(ctx context.Context, in *pb.GetFolderContentRequest) (*pb.GetFolderContentResponse, error) {
	var uid string
	uid = strings.Trim(in.GetUserId(), " ")
	if uid == "" {
		return nil, status.Error(codes.InvalidArgument, "UserID can't be empty")
	}
	var fid string
	fid = strings.Trim(in.GetIdentifier(), " ")
	if fid == "" {
		return nil, status.Error(codes.InvalidArgument, "FolderID can't be empty")
	}

	level, err := s.Service.GetFolderContent(ctx, uid, fid)
	if err != nil {
		return nil, err
	}

	rsp := new(pb.GetFolderContentResponse)
	rsp.Content = level
	rsp.ResponseCode = codes.OK.String()
	rsp.ResponseMessage = "Folder listed"

	return rsp, nil
}

func (s *Server) GetFileContent(ctx context.Context, in *pb.GetFileContentRequest) (*pb.GetFileContentResponse, error) {
	var uid string
	uid = strings.Trim(in.GetUserId(), " ")
	if uid == "" {
		return nil, status.Error(codes.InvalidArgument, "UserID can't be empty")
	}
	var fid string
	fid = strings.Trim(in.GetIdentifier(), " ")
	if fid == "" {
		return nil, status.Error(codes.InvalidArgument, "FileID can't be empty")
	}

	file, err := s.Service.GetFileContent(ctx, uid, fid, in.GetType(), in.GetSecretKey())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	rsp := new(pb.GetFileContentResponse)
	rsp.File = file

	return rsp, nil
}

func (s *Server) GetFileTree(ctx context.Context, in *pb.GetFileTreeRequest) (*pb.GetFileTreeResponse, error) {
	// TODO: implement
	return nil, nil
}

func (s *Server) UpdateFileContent(ctx context.Context, in *pb.UpdateFileContentRequest) (*pb.UpdateFileContentResponse, error) {
	var uid string
	uid = strings.Trim(in.GetUserId(), " ")
	if uid == "" {
		return nil, status.Error(codes.InvalidArgument, "UserID can't be empty")
	}
	var fid string
	fid = strings.Trim(in.GetIdentifier(), " ")
	if fid == "" {
		return nil, status.Error(codes.InvalidArgument, "FileID can't be empty")
	}

	err := s.Service.UpdateFileContent(ctx, uid, fid, in.GetFileType(), in.GetSecretKey(), in.GetNewContent())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := new(pb.UpdateFileContentResponse)
	rsp.ResponseCode = codes.Internal.String()
	rsp.ResponseMessage = "File update successful"

	return rsp, err
}

func (s *Server) DeleteFile(ctx context.Context, in *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	var uid string
	uid = strings.Trim(in.GetUserId(), " ")
	if uid == "" {
		return nil, status.Error(codes.InvalidArgument, "UserID can't be empty")
	}
	var fid string
	fid = strings.Trim(in.GetIdentifier(), " ")
	if fid == "" {
		return nil, status.Error(codes.InvalidArgument, "FileID can't be empty")
	}

	err := s.Service.DeleteFile(ctx, uid, fid)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	rsp := new(pb.DeleteFileResponse)
	rsp.ResponseCode = codes.OK.String()
	rsp.ResponseMessage = "File deleted successfully"
	return rsp, nil
}

func (s *Server) RenameFile(ctx context.Context, in *pb.RenameFileRequest) (*pb.RenameFileResponse, error) {
	var uid string
	uid = strings.Trim(in.GetUserId(), " ")
	if uid == "" {
		return nil, status.Error(codes.InvalidArgument, "UserID can't be empty")
	}
	var fid string
	fid = strings.Trim(in.GetIdentifier(), " ")
	if fid == "" {
		return nil, status.Error(codes.InvalidArgument, "FileID can't be empty")
	}
	var name string
	name = strings.Trim(in.GetName(), " ")
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "New file name can't be empty")
	}

	err := s.Service.RenameFile(ctx, uid, fid, name)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	rsp := new(pb.RenameFileResponse)
	rsp.ResponseCode = codes.OK.String()
	rsp.ResponseMessage = "File renamed successfully"
	return rsp, nil
}

func (s *Server) MoveFile(ctx context.Context, in *pb.MoveFileRequest) (*pb.MoveFileResponse, error) {
	var uid string
	uid = strings.Trim(in.GetUserId(), " ")
	if uid == "" {
		return nil, status.Error(codes.InvalidArgument, "UserID can't be empty")
	}
	var fid string
	fid = strings.Trim(in.GetIdentifier(), " ")
	if fid == "" {
		return nil, status.Error(codes.InvalidArgument, "FileID can't be empty")
	}
	var did string
	did = strings.Trim(in.GetDestination(), " ")
	if did == "" {
		return nil, status.Error(codes.InvalidArgument, "Destination identifier can't be empty")
	}
	var sid string
	sid = strings.Trim(in.GetSource(), " ")
	if sid == "" {
		return nil, status.Error(codes.InvalidArgument, "Source identifier can't be empty")
	}
	err := s.Service.MoveFile(ctx, uid, fid, sid, did, in.GetCopy())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	rsp := new(pb.MoveFileResponse)
	rsp.ResponseCode = codes.OK.String()
	rsp.ResponseMessage = "File renamed successfully"
	return rsp, nil
}

func (s *Server) UploadFile(ctx context.Context, in *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	file := in.File
	if file.Content == nil {
		return nil, status.Error(codes.InvalidArgument, "File content can't be nil")
	}
	var uid string
	uid = strings.Trim(in.GetUserId(), " ")
	if uid == "" {
		return nil, status.Error(codes.InvalidArgument, "UserID can't be empty")
	}
	var name string
	name = strings.Trim(file.GetName(), " ")
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "File name can't be empty")
	}
	if len(in.GetSecretPhrase()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Secret phrase can't be empty")
	}
	var did string
	did = strings.Trim(file.GetParentIdentifier(), " ")
	if did == "" {
		return nil, status.Error(codes.InvalidArgument, "Parent identifier can't be empty")
	}

	var Type string
	switch file.GetFileType() {
	case pb.Type_IMAGE:
		Type = "PNG"
	case pb.Type_PDF:
		Type = "PDF"
	default:
		Type = "TXT"
	}

	identifier, message, err := s.Service.UploadFile(ctx, uid, name, did, Type, file.Writable, file.Private, file.Content, in.SecretPhrase)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &pb.UploadFileResponse{
		ResponseCode:     codes.OK.String(),
		ResponseMessage:  "File created with success",
		FileId:           identifier,
		WatermarkMessage: message,
	}

	return resp, nil
}

func (s *Server) ShareFile(ctx context.Context, in *pb.ShareFileRequest) (*pb.ShareFileResponse, error) {
	// TODO: implement methods
	return nil, nil
}

func (s *Server) DownloadFile(ctx context.Context, in *pb.DownloadFileRequest) (*pb.DownloadFileResponse, error) {
	var uid string
	uid = strings.Trim(in.GetUserId(), " ")
	if uid == "" {
		return nil, status.Error(codes.InvalidArgument, "UserID can't be empty")
	}
	var fid string
	fid = strings.Trim(in.GetIdentifier(), " ")
	if fid == "" {
		return nil, status.Error(codes.InvalidArgument, "File identifier can't be empty")
	}
	if len(in.GetSecretPhrase()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Secret phrase can't be empty")
	}
	var steganoMessage []byte = nil
	if len(in.GetSteganoMessage()) != 0 {
		steganoMessage = in.GetSteganoMessage()
	}
	var watermarkImage []byte = nil
	if len(in.GetWatermarkImage()) != 0 {
		watermarkImage = in.GetWatermarkImage()
	}
	var watermarkMessage string = ""
	if len(in.GetWatermarkMessage()) != 0 {
		watermarkMessage = in.GetWatermarkMessage()
	}
	file, err := s.Service.DownloadFile(ctx, uid, fid, in.GetSecretPhrase(), watermarkMessage, steganoMessage, watermarkImage)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	fmt.Printf("%v\n\n", file)

	rsp := new(pb.DownloadFileResponse)
	rsp.File = file
	rsp.ResponseCode = codes.OK.String()
	rsp.ResponseMessage = "File downloaded"
	return rsp, nil
}

func (s *Server) ComputeSize(ctx context.Context, in *pb.ComputeSizeRequest) (*pb.ComputeSizeResponse, error) {
	iSize, cSize, err := s.Service.ComputeSize(ctx, in.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &pb.ComputeSizeResponse{
		InitialSize:    iSize,
		CompressedSize: cSize,
	}
	return rsp, nil
}

func (s *Server) GetMyDriveId(ctx context.Context, in *pb.GetMyDriveIdRequest) (*pb.GetMyDriveIdResponse, error) {
	id, err := s.Service.GetMyDriveId(ctx, in.GetUserId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	rsp := &pb.GetMyDriveIdResponse{
		DriveId: id,
	}
	return rsp, nil
}
