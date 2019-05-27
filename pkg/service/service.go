package service

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"

	compb "github.com/bitstored/compression-service/pb"
	crypb "github.com/bitstored/crypto-service/pb"
	watpb "github.com/bitstored/watermarking-service/pb"

	"github.com/bitstored/file-service/pb"
	"github.com/bitstored/file-service/pkg/repo"
	"github.com/bitstored/file-service/pkg/service/commands"
	"github.com/bitstored/file-service/pkg/service/events"
)

const (
	CRYPTO_GRPC_PORT       = "localhost:4004"
	COMPRESSION_GRPC_PORT  = "localhost:4003"
	WATERMARKING_GRPC_PORT = "localhost:4008"
)

type Service struct {
	Repo *repo.Repository
}

func NewService() *Service {
	repo, err := repo.NewFileRepository()
	if err != nil {
		return nil
	}
	return &Service{Repo: repo}
}

// CreateDrive gets the user identifier and creates a drive associated with user, returns the drive ID
func (s *Service) CreateDrive(ctx context.Context, userID string) (string, error) {
	command := new(commands.CreateDrive)
	command.UserID = userID

	event := s.LaunchCommand(ctx, command).(*events.DriveCreated)

	if event.Error != nil {
		return "", event.Error
	}
	return event.DriveID, nil
}

// CreateNewFile creates a file associated with user and containing the following data, returns the file ID
func (s *Service) CreateNewFile(ctx context.Context, userID, fileName, parentID, fileType string, writable, private bool, content, secretKey []byte) (string, error) {
	//compress file according to file type
	compressed, err := s.compress(ctx, fileType, content)
	if err != nil {
		return "", err
	}
	// if file is private, encrypt file content
	encrypted, err := s.encrypt(ctx, compressed, secretKey, userID)
	if err != nil {
		return "", err
	}
	command := new(commands.CreateNewFile)
	command.UserID = userID
	command.ParentID = parentID
	command.CreationDate = time.Now().String()
	command.Name = fileName
	command.FileType = fileType
	command.Writable = writable
	command.Private = private
	command.Content = encrypted

	event := s.LaunchCommand(ctx, command).(*events.NewFileCreated)

	if event.Error != nil {
		return "", event.Error
	}
	return event.FileID, nil
}

func (s *Service) CreateNewFolder(ctx context.Context, userID, folderName, parentID string) (string, error) {
	command := new(commands.CreateNewFolder)
	command.UserID = userID
	command.Name = folderName
	command.ParentID = parentID
	command.CreationDate = time.Now().String()

	event := s.LaunchCommand(ctx, command).(*events.NewFolderCreated)
	if event.Error != nil {
		return "", event.Error
	}
	return event.FolderID, nil
}

func (s *Service) UpdateFileContent(ctx context.Context, userID, fileID string, fileType pb.Type, secretKey, content []byte) error {
	command := new(commands.UpdateFileContent)
	command.UserID = userID
	if content == nil {
		return fmt.Errorf("File content can't be nil")
	}
	var fType string
	if fileType == pb.Type_IMAGE {
		fType = "PNG"
	} else if fileType == pb.Type_PDF {
		fType = "PDF"
	} else {
		fType = "TXT"
	}
	compressed, err := s.compress(ctx, fType, content)
	if err != nil {
		return err
	}
	// if file is private, encrypt file content
	encrypted, err := s.encrypt(ctx, compressed, secretKey, userID)
	if err != nil {
		return err
	}
	command.FileID = fileID
	command.Data = encrypted
	event := s.LaunchCommand(ctx, command).(events.FileContentUpdated)
	return event.Error
}

func (s *Service) DeleteFile(ctx context.Context, userID, fileID string) error {
	command := new(commands.DeleteFile)
	command.UserID = userID
	command.FileID = fileID
	event := s.LaunchCommand(ctx, command).(events.FileDeleted)
	return event.Error
}

func (s *Service) RenameFile(ctx context.Context, userID, fileID, newName string) error {
	command := new(commands.RenameFile)
	command.UserID = userID
	command.FileID = fileID
	command.Name = newName
	event := s.LaunchCommand(ctx, command).(events.FileRenamed)
	return event.Error
}

func (s *Service) MoveFile(ctx context.Context, userID, fileID, sourceID, destinationID string, copy bool) error {
	command := new(commands.MoveFile)
	command.UserID = userID
	command.FileID = fileID
	command.DestinationID = destinationID
	command.SourceID = sourceID
	command.Copy = copy
	if sourceID == destinationID {
		return fmt.Errorf("File can not be moved to the same directory")
	}
	event := s.LaunchCommand(ctx, command).(events.FileRenamed)
	return event.Error
}

func (s *Service) UploadFile(ctx context.Context, userID, fileName, parentID, fileType string, writable, private bool, content, secretKey []byte) (string, []byte, error) {
	//compress file according to file type
	compressed, err := s.compress(ctx, fileType, content)
	if err != nil {
		return "", nil, err
	}
	msg, err := s.decodeMessage(ctx, content)
	if err != nil {
		msg = []byte{}
	}
	// if file is private, encrypt file content
	encrypted, err := s.encrypt(ctx, compressed, secretKey, userID)
	if err != nil {
		return "", nil, err
	}
	command := new(commands.UploadFile)
	command.UserID = userID
	command.ParentID = parentID
	command.CreationDate = time.Now().String()
	command.Name = fileName
	command.FileType = fileType
	command.Writable = writable
	command.Private = private
	command.Content = encrypted

	event := s.LaunchCommand(ctx, command).(*events.FileUploaded)

	if event.Error != nil {
		return "", nil, event.Error
	}

	return event.FileID, msg, nil
}

func (s *Service) DownloadFile(ctx context.Context, userID, fileID string, secretKey []byte, watermarkingPhrase string, steganographyPhrase []byte, watermarkingImage []byte) (*pb.File, error) {
	command := new(commands.DownloadFile)
	command.UserID = userID
	command.FileID = fileID
	event := s.LaunchCommand(ctx, command).(*events.FileDownloaded)
	if event.Error != nil {
		return nil, event.Error
	}
	content := event.Content
	content, err := s.decrypt(ctx, content, secretKey, userID)
	if err != nil {
		return nil, fmt.Errorf("Could not decrypt the content")
	}
	content, err = s.decompress(ctx, event.FileType, content)
	if err != nil {
		return nil, fmt.Errorf("Could not decompress the content")
	}
	file := new(pb.File)
	file.Name = event.Name
	file.CreationDate = event.CreationDate
	file.FileType = pb.Type_OTHER
	if event.Writable && event.FileType == "PNG" {
		if watermarkingImage != nil {
			content, err = s.watermarkWithImage(ctx, content, watermarkingImage)
		}
		if watermarkingPhrase != "" {
			content, err = s.watermarkWithText(ctx, content, watermarkingPhrase)
		}
	}
	if event.Private && event.FileType == "PNG" {
		content, err = s.encodeMessage(ctx, content, steganographyPhrase)
	}
	return file, nil
}

func (s *Service) GetFileContent(ctx context.Context, userID, fileID string, fileType pb.Type, secretKey []byte) (*pb.File, error) {
	cont, err := s.Repo.GetFileContent(ctx, userID, fileID)
	if err != nil {
		return nil, fmt.Errorf("Could not find the file")
	}
	content := cont.Data
	content, err = s.decrypt(ctx, content, secretKey, userID)
	if err != nil {
		return nil, fmt.Errorf("Could not decrypt the content")
	}
	var fType string
	if fileType == pb.Type_IMAGE {
		fType = "PNG"
	} else if fileType == pb.Type_PDF {
		fType = "PDF"
	} else {
		fType = "TXT"
	}
	content, err = s.decompress(ctx, fType, content)
	if err != nil {
		return nil, fmt.Errorf("Could not decompress the content")
	}

	file := new(pb.File)
	file.Content = content
	file.FileType = pb.Type_OTHER

	return file, nil
}

func (s *Service) GetFolderContent(ctx context.Context, userID, folderID string) (*pb.FSLevel, error) {
	cont, err := s.Repo.GetFolderContent(ctx, userID, folderID)
	if err != nil {
		return nil, fmt.Errorf("Could not find the file")
	}
	fsLevel := new(pb.FSLevel)
	for _, fr := range cont.Folders {
		fp := new(pb.Folder)
		fp.CreationDate = fr.Created
		fp.Identifier = fr.ID
		fp.Name = fr.Name
		fp.ParentIdentifier = folderID
		fsLevel.Folders = append(fsLevel.Folders, fp)
	}
	for _, fr := range cont.Files {
		fp := new(pb.File)
		fp.CreationDate = fr.Created
		fp.Identifier = fr.ID
		fp.Name = fr.Name
		fp.Content = []byte("unknown")
		if fr.Type == "PNG" {
			fp.FileType = pb.Type_IMAGE
		} else if fr.Type == "PDF" {
			fp.FileType = pb.Type_PDF
		} else {
			fp.FileType = pb.Type_TXT
		}
		fp.Private = fr.IsPrivate
		fp.Writable = fr.IsWritable
		fp.ParentIdentifier = folderID
		fsLevel.Files = append(fsLevel.Files, fp)
	}
	return fsLevel, nil
}

func (s *Service) encodeMessage(ctx context.Context, target []byte, message []byte) ([]byte, error) {
	conn, err := grpc.Dial(WATERMARKING_GRPC_PORT, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	client := watpb.NewWatermarkingClient(conn)
	req := &watpb.EncodeMessageRequest{
		Image: target,
		Text:  message,
	}
	rsp, err := client.EncodeMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.Image, nil
}

func (s *Service) decodeMessage(ctx context.Context, target []byte) ([]byte, error) {
	conn, err := grpc.Dial(WATERMARKING_GRPC_PORT, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	client := watpb.NewWatermarkingClient(conn)
	req := &watpb.DecodeMessageRequest{
		Image: target,
	}
	rsp, err := client.DecodeMessage(ctx, req)
	if err != nil {
		return nil, err
	}
	return rsp.Text, nil
}

func (s *Service) watermarkWithImage(ctx context.Context, target []byte, data []byte) ([]byte, error) {
	conn, err := grpc.Dial(WATERMARKING_GRPC_PORT, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	client := watpb.NewWatermarkingClient(conn)
	req := watpb.WatermarkImageWithImageRequest{target, data}
	rsp, err := client.WatermarkImageWithImage(ctx, &req)
	if err != nil {
		return nil, err
	}
	return rsp.Image, nil
}

func (s *Service) watermarkWithText(ctx context.Context, target []byte, data string) ([]byte, error) {
	conn, err := grpc.Dial(WATERMARKING_GRPC_PORT, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	client := watpb.NewWatermarkingClient(conn)
	req := watpb.WatermarkImageWithTextRequest{target, data}
	rsp, err := client.WatermarkImageWithText(ctx, &req)
	if err != nil {
		return nil, err
	}
	return rsp.Image, nil
}

func (s *Service) compress(ctx context.Context, fileType string, data []byte) ([]byte, error) {
	compressionConn, err := grpc.Dial(COMPRESSION_GRPC_PORT, grpc.WithInsecure())
	defer compressionConn.Close()

	if err != nil {
		return nil, err
	}
	compressionClient := compb.NewCompressionClient(compressionConn)

	switch fileType {
	case "PNG", "png":
		{
			req := &compb.CompressImageRequest{Image: data,
				Level: compb.CompressionLevel_BestCompression,
				Type:  compb.ImageType_PNG}
			resp, err := compressionClient.CompressImage(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp.Image, nil
		}

	case "JPEG", "JPG", "jpeg", "jpg":
		{
			req := &compb.CompressImageRequest{Image: data,
				Level: compb.CompressionLevel_BestCompression,
				Type:  compb.ImageType_JPEG}
			resp, err := compressionClient.CompressImage(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp.Image, nil
		}
	case "TXT", "txt", "code", "PDF", "pdf":
		{
			req := &compb.CompressTextRequest{Text: data,
				Level: compb.CompressionLevel_BestCompression}
			resp, err := compressionClient.CompressText(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp.Text, nil
		}
	default:
		return nil, fmt.Errorf("Unknown file type")
	}
}

func (s *Service) decompress(ctx context.Context, fileType string, data []byte) ([]byte, error) {
	compressionConn, err := grpc.Dial(COMPRESSION_GRPC_PORT, grpc.WithInsecure())
	defer compressionConn.Close()

	if err != nil {
		return nil, err
	}
	compressionClient := compb.NewCompressionClient(compressionConn)

	switch fileType {
	case "PNG", "png":
		{
			req := &compb.DecompressImageRequest{Image: data,
				Level: compb.CompressionLevel_BestCompression,
				Type:  compb.ImageType_PNG}
			resp, err := compressionClient.DecompressImage(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp.Image, nil
		}

	case "JPEG", "JPG", "jpeg", "jpg":
		{
			req := &compb.DecompressImageRequest{Image: data,
				Level: compb.CompressionLevel_BestCompression,
				Type:  compb.ImageType_JPEG}
			resp, err := compressionClient.DecompressImage(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp.Image, nil
		}
	case "TXT", "txt", "code", "PDF", "pdf":
		{
			req := &compb.DecompressTextRequest{Text: data}
			resp, err := compressionClient.DecompressText(ctx, req)
			if err != nil {
				return nil, err
			}
			return resp.Text, nil
		}
	default:
		return nil, fmt.Errorf("Unknown file type")
	}
}

func (s *Service) encrypt(ctx context.Context, data []byte, secretKey []byte, userID string) ([]byte, error) {
	conn, err := grpc.Dial(CRYPTO_GRPC_PORT, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	client := crypb.NewCryptoClient(conn)

	file := &crypb.File{data, secretKey}
	req := &crypb.EncryptFileRequest{file, userID}

	resp, err := client.EncryptFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.EncryptedData, nil
}

func (s *Service) decrypt(ctx context.Context, data []byte, secretKey []byte, userID string) ([]byte, error) {
	conn, err := grpc.Dial(CRYPTO_GRPC_PORT, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		return nil, err
	}
	client := crypb.NewCryptoClient(conn)

	file := &crypb.File{data, secretKey}
	req := &crypb.DecryptFileRequest{EncryptedFile: file, SecretPhrase: secretKey}

	resp, err := client.DecryptFile(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.OriginalData, nil
}
