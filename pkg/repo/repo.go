package repo

import (
	"context"
	cdb "github.com/bitstored/repository/pkg/couchbaserepo"
	"github.com/couchbase/gocb"
)

const (
	connSpecStr    = "http://localhost"
	username       = "Administrator"
	password       = "password"
	bucketName     = "bitstored"
	bucketPassword = ""
)

type Repository struct {
	next cdb.Repository
}

func FileRepository() (*Repository, error) {
	// connect to cluster
	cluster, err := gocb.Connect(connSpecStr)
	if err != nil {
		return nil, err
	}
	cluster.Authenticate(gocb.PasswordAuthenticator{
		Username: username,
		Password: password,
	})
	// open bucket
	bucket, _ := cluster.OpenBucket(bucketName, bucketPassword)
	return &Repository{next: cdb.NewRepository(bucket)}, nil
}

func (r *Repository) GetFolderContent(ctx context.Context, in *pb.GetFolderContentRequest) (*pb.GetFolderContentResponse, error) {
	return nil, nil
}
func (r *Repository) GetFileContent(ctx context.Context, in *pb.GetFileContentRequest) (*pb.GetFileContentResponse, error) {
	return nil, nil
}
func (r *Repository) GetFileTree(ctx context.Context, in *pb.GetFileTreeRequest) (*pb.GetFileTreeResponse, error) {
	return nil, nil
}
func (r *Repository) UpdateFileContent(ctx context.Context, in *pb.UpdateFileContentRequest) (*pb.UpdateFileContentResponse, error) {
	return nil, nil
}
func (r *Repository) CreateDrive(ctx context.Context, in *pb.CreateDriveRequest) (*pb.CreateDriveResponse, error) {
	return nil, nil
}
func (r *Repository) CreateNewFile(ctx context.Context, in *pb.CreateNewFileRequest) (*pb.CreateNewFileResponse, error) {
	return nil, nil
}
func (r *Repository) DeleteFile(ctx context.Context, in *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	return nil, nil
}
func (r *Repository) RenameFile(ctx context.Context, in *pb.RenameFileRequest) (*pb.RenameFileResponse, error) {
	return nil, nil
}
func (r *Repository) MoveFile(ctx context.Context, in *pb.MoveFileRequest) (*pb.UploadFileResponse, error) {
	return nil, nil
}
func (r *Repository) UploadFile(ctx context.Context, in *pb.UploadFileRequest) (*pb.UploadFileResponse, error) {
	return nil, nil
}
