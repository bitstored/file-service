package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/bitstored/file-service/pkg/service/commands"
	"github.com/bitstored/file-service/pkg/service/events"
)

func TestService_createNewFolder(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx     context.Context
		command *commands.CreateNewFolder
		c       chan *events.NewFolderCreated
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			s.createNewFolder(tt.args.ctx, tt.args.command, tt.args.c)
		})
	}
}

func TestService_LaunchCommand(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx     context.Context
		command interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			if got := s.LaunchCommand(tt.args.ctx, tt.args.command); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.LaunchCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_createDrive(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx     context.Context
		command *commands.CreateDrive
		c       chan *events.DriveCreated
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			s.createDrive(tt.args.ctx, tt.args.command, tt.args.c)
		})
	}
}

func TestService_createNewFile(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx     context.Context
		command *commands.CreateNewFile
		c       chan *events.NewFileCreated
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			s.createNewFile(tt.args.ctx, tt.args.command, tt.args.c)
		})
	}
}

func TestService_updateFileContent(t *testing.T) {
	type fields struct {
		Repo *repo.Repository
	}
	type args struct {
		ctx     context.Context
		command *commands.UpdateFileContent
		c       chan *events.FileContentUpdated
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				Repo: tt.fields.Repo,
			}
			s.updateFileContent(tt.args.ctx, tt.args.command, tt.args.c)
		})
	}
}
