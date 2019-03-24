package repo

import (
	"context"
	"testing"
)

func TestGetFolderContent(t *testing.T) {
	t.Skip("not implemented yet")
}
func TestGetFileContent(t *testing.T) {
	t.Skip("not implemented yet")
}
func TestGetFileTree(t *testing.T) {
	t.Skip("not implemented yet")
}
func TestUpdateFileContent(t *testing.T) {
	t.Skip("not implemented yet")
}
func TestCreateDrive(t *testing.T) {
	r, _ := NewFileRepository()
	_, err := r.CreateDrive(context.Background(), "uid: dididii")
	if err != nil {
		panic(err)
	}
}
func TestCreateNewFile(t *testing.T) {
	t.Skip("not implemented yet")
}
func TestDeleteFile(t *testing.T) {
	t.Skip("not implemented yet")
}
func TestRenameFile(t *testing.T) {
	t.Skip("not implemented yet")
}
func TestMoveFile(t *testing.T) {
	t.Skip("not implemented yet")
}
func TestUploadFile(t *testing.T) {
	t.Skip("not implemented yet")
}
