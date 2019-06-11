package service

import (
	"context"
	"fmt"

	"github.com/bitstored/file-service/pkg/service/commands"
	"github.com/bitstored/file-service/pkg/service/events"
)

func (s *Service) LaunchCommand(ctx context.Context, command interface{}) interface{} {

	switch command.(type) {
	case *commands.CreateDrive:
		{
			cdCommand := command.(*commands.CreateDrive)
			if !cdCommand.IsValid() {
				return nil
			}

			channel := make(chan *events.DriveCreated)

			go s.createDrive(ctx, cdCommand, channel)

			return <-channel
		}
	case *commands.CreateNewFile:
		{
			cfCommand := command.(*commands.CreateNewFile)
			if !cfCommand.IsValid() {
				return nil
			}

			channel := make(chan *events.NewFileCreated)

			go s.createNewFile(ctx, cfCommand, channel)

			return <-channel
		}
	case *commands.CreateNewFolder:
		{
			cfCommand := command.(*commands.CreateNewFolder)
			if !cfCommand.IsValid() {
				return nil
			}

			channel := make(chan *events.NewFolderCreated)

			go s.createNewFolder(ctx, cfCommand, channel)

			return <-channel
		}
	case *commands.UpdateFileContent:
		{
			cfCommand := command.(*commands.UpdateFileContent)
			if !cfCommand.IsValid() {
				return nil
			}

			channel := make(chan *events.FileContentUpdated)

			go s.updateFileContent(ctx, cfCommand, channel)

			return <-channel
		}
	case *commands.DeleteFile:
		{
			cfCommand := command.(*commands.DeleteFile)
			if !cfCommand.IsValid() {
				return nil
			}

			channel := make(chan *events.FileDeleted)

			go s.deleteFile(ctx, cfCommand, channel)

			return <-channel
		}
	case *commands.RenameFile:
		{
			cfCommand := command.(*commands.RenameFile)
			if !cfCommand.IsValid() {
				return nil
			}

			channel := make(chan *events.FileRenamed)

			go s.renameFile(ctx, cfCommand, channel)

			return <-channel
		}
	case *commands.MoveFile:
		{
			cfCommand := command.(*commands.MoveFile)
			if !cfCommand.IsValid() {
				return nil
			}

			channel := make(chan *events.FileMoved)

			go s.moveFile(ctx, cfCommand, channel)

			return <-channel
		}
	case *commands.UploadFile:
		{
			cfCommand := command.(*commands.UploadFile)
			if !cfCommand.IsValid() {
				return nil
			}

			channel := make(chan *events.FileUploaded)

			go s.uploadFile(ctx, cfCommand, channel)

			return <-channel
		}
	case *commands.DownloadFile:
		{
			cfCommand := command.(*commands.DownloadFile)
			if !cfCommand.IsValid() {
				return nil
			}

			channel := make(chan *events.FileDownloaded)

			go s.downloadFile(ctx, cfCommand, channel)

			return <-channel
		}
	default:
		return &events.Failed{Error: fmt.Errorf("Unknown command type")}
	}
}

func (s *Service) createDrive(ctx context.Context, command *commands.CreateDrive, c chan *events.DriveCreated) {
	rsp, err := s.Repo.CreateDrive(ctx, command.UserID)
	event := events.DriveCreated{DriveID: string(rsp), UserID: command.UserID, Error: err}
	c <- &event
}

func (s *Service) createNewFile(ctx context.Context, command *commands.CreateNewFile, c chan *events.NewFileCreated) {
	rsp, err := s.Repo.CreateNewFile(ctx,
		command.UserID,
		command.Name,
		command.ParentID,
		command.FileType,
		command.Writable,
		command.Private,
		command.Content,
		command.InitialSize,
		command.CompressedSize)
	event := events.NewFileCreated{FileID: string(rsp), UserID: command.UserID, Error: err}
	c <- &event
}

func (s *Service) createNewFolder(ctx context.Context, command *commands.CreateNewFolder, c chan *events.NewFolderCreated) {
	rsp, err := s.Repo.CreateNewFolder(ctx,
		command.UserID,
		command.Name,
		command.ParentID)
	event := events.NewFolderCreated{FolderID: string(rsp), UserID: command.UserID, Error: err}
	c <- &event
}

func (s *Service) updateFileContent(ctx context.Context, command *commands.UpdateFileContent, c chan *events.FileContentUpdated) {
	rsp, err := s.Repo.UpdateFileContent(ctx,
		command.UserID,
		command.FileID,
		command.Data)
	event := events.FileContentUpdated{FileID: string(rsp), UserID: command.UserID, Error: err}
	c <- &event
}

func (s *Service) deleteFile(ctx context.Context, command *commands.DeleteFile, c chan *events.FileDeleted) {
	err := s.Repo.DeleteFile(ctx,
		command.UserID,
		command.FileID)
	event := events.FileDeleted{FileID: command.FileID, UserID: command.UserID, Error: err}
	c <- &event
}

func (s *Service) renameFile(ctx context.Context, command *commands.RenameFile, c chan *events.FileRenamed) {
	err := s.Repo.RenameFile(ctx,
		command.UserID,
		command.FileID,
		command.Name)
	event := events.FileRenamed{FileID: command.FileID, UserID: command.UserID, Error: err}
	c <- &event
}

func (s *Service) moveFile(ctx context.Context, command *commands.MoveFile, c chan *events.FileMoved) {
	err := s.Repo.MoveFile(ctx,
		command.FileID,
		command.SourceID,
		command.DestinationID,
		command.Copy)
	event := events.FileMoved{FileID: command.FileID, UserID: command.UserID, Error: err}
	c <- &event
}

func (s *Service) uploadFile(ctx context.Context, command *commands.UploadFile, c chan *events.FileUploaded) {
	rsp, err := s.Repo.UploadFile(ctx,
		command.UserID,
		command.Name,
		command.ParentID,
		command.FileType,
		command.Writable,
		command.Private,
		command.Content)
	event := events.FileUploaded{FileID: string(rsp), UserID: command.UserID, Error: err}
	c <- &event
}

func (s *Service) downloadFile(ctx context.Context, command *commands.DownloadFile, c chan *events.FileDownloaded) {
	meta, content, err := s.Repo.DownloadFile(ctx,
		command.UserID,
		command.FileID)
	event := events.FileDownloaded{FileID: command.FileID,
		UserID:       command.UserID,
		CreationDate: meta.Created,
		Name:         meta.Name,
		FileType:     meta.Type,
		Writable:     meta.IsWritable,
		Private:      meta.IsPrivate,
		Content:      content.Data,
		Error:        err}
	c <- &event
}
