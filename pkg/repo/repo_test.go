package repo

import (
	"context"
	"testing"
	
	"github.com/stretchr/testify/require"
)


func TestCreateDrive(t *testing.T) {
	r, _ := NewFileRepository()
	_, err := r.CreateDrive(context.Background(), "uid:TestCreateDrive")
	if err != nil {
		panic(err)
	}
}

func TestGetFolderContent(t *testing.T) {
	userID := "uid:TestGetFolderContent"
	r, _ := NewFileRepository()
	driveID, err := r.CreateDrive(context.Background(), userID)
	require.NoError(t, err)
	folderID, err := r.CreateNewFolder(context.Background(), userID, "folder1", string(driveID))
	require.NoError(t, err)

	fileID, err := r.CreateNewFile(context.Background(), userID, "file1", folderID, "txt", true, true, []byte("some text"))
	require.NoError(t, err)

	fsLevel, err := r.GetFolderContent(context.Background(), userID, folderID)
	require.NoError(t, err)

	require.NotNil(t, fsLevel)
	require.Equal(t,0,  len(fsLevel.Folders) )
	require.Equal(t,1,  len(fsLevel.Files))

	require.Equal(t, fileID , fsLevel.Files[0].ID)
		
}
func TestGetFileContent(t *testing.T) {
	userID := "uid:TestGetFileContent"
	content := []byte("some text")
	r, _ := NewFileRepository()
	driveID, err := r.CreateDrive(context.Background(), userID)
	require.NoError(t, err)
	
	fileID, err := r.CreateNewFile(context.Background(), userID, "file1", string(driveID), "txt", true, true, content )
	require.NoError(t, err)

	cont, err := r.GetFileContent(context.Background(), userID, fileID)
	require.NoError(t, err)

	require.NotNil(t, cont)

	require.EqualValues(t, content , cont.Data)
}

func TestUpdateFileContent(t *testing.T) {
	userID := "uid:TestUpdateFileContent"
	content1 := []byte("some text")
	content2 := []byte("some other text")
	r, _ := NewFileRepository()
	driveID, err := r.CreateDrive(context.Background(), userID)
	require.NoError(t, err)
	
	fileID, err := r.CreateNewFile(context.Background(), userID, "file1", string(driveID), "txt", true, true, content1)
	require.NoError(t, err)

	_, err = r.UpdateFileContent(context.Background(), userID, fileID, content2)
	require.NoError(t, err)

	cont, err := r.GetFileContent(context.Background(), userID, fileID)
	require.NoError(t, err)

	require.NotNil(t, cont)

	require.EqualValues(t, content2 , cont.Data)
}

func TestDeleteFile(t *testing.T) {
	userID := "uid:TestDeleteFile"
	content1 := []byte("some text")
	r, _ := NewFileRepository()
	driveID, err := r.CreateDrive(context.Background(), userID)
	require.NoError(t, err)
	
	fileID, err := r.CreateNewFile(context.Background(), userID, "file1", string(driveID), "txt", true, true, content1)
	require.NoError(t, err)

	err = r.DeleteFile(context.Background(), userID, fileID)
	require.NoError(t, err)

	cont, err := r.GetFileContent(context.Background(), userID, fileID)
	require.Error(t, err)

	require.Nil(t, cont)
}

func TestCreateFile(t *testing.T) {
	userID := t.Name()

	r, _ := NewFileRepository()
	driveID, err := r.CreateDrive(context.Background(), userID)
	require.NoError(t, err)
	folderID, err := r.CreateNewFolder(context.Background(), userID, "folder1", string(driveID))
	require.NoError(t, err)

	fileID1, err := r.CreateNewFile(context.Background(), userID, "file1", folderID, "txt", true, true, []byte("some text"))
	require.NoError(t, err)
	
	fileID2, err := r.CreateNewFile(context.Background(), userID, "file2", folderID, "txt", true, true, []byte("some text1"))
	require.NoError(t, err)
	fsLevel, err := r.GetFolderContent(context.Background(), userID, folderID)
	require.NoError(t, err)

	require.NotNil(t, fsLevel)
	require.Equal(t,0,  len(fsLevel.Folders) )
	require.Equal(t,2,  len(fsLevel.Files))

	require.Equal(t, fileID1 , fsLevel.Files[0].ID)
	require.Equal(t, fileID2 , fsLevel.Files[1].ID)
}

func TestUploadFile(t *testing.T) {
	userID := t.Name()

	r, _ := NewFileRepository()
	driveID, err := r.CreateDrive(context.Background(), userID)
	require.NoError(t, err)
	folderID, err := r.CreateNewFolder(context.Background(), userID, "folder1", string(driveID))
	require.NoError(t, err)

	fileID1, err := r.CreateNewFile(context.Background(), userID, "file1", folderID, "txt", true, true, []byte("some text"))
	require.NoError(t, err)
	
	fileID2, err := r.UploadFile(context.Background(), userID, "file2", folderID, "txt", true, true, []byte("some text1"))
	require.NoError(t, err)
	fsLevel, err := r.GetFolderContent(context.Background(), userID, folderID)
	require.NoError(t, err)

	require.NotNil(t, fsLevel)
	require.Equal(t,0,  len(fsLevel.Folders) )
	require.Equal(t,2,  len(fsLevel.Files))

	require.Equal(t, fileID1 , fsLevel.Files[0].ID)
	require.Equal(t, fileID2 , fsLevel.Files[1].ID)
}

func TestDownloadFile(t *testing.T) {
userID := t.Name()

	r, _ := NewFileRepository()
	driveID, err := r.CreateDrive(context.Background(), userID)
	require.NoError(t, err)
	folderID, err := r.CreateNewFolder(context.Background(), userID, "folder1", string(driveID))
	require.NoError(t, err)

	fileID, err := r.CreateNewFile(context.Background(), userID, "file1", folderID, "txt", true, true, []byte("some text"))
	require.NoError(t, err)
	
	_, cont, err := r.DownloadFile(context.Background(), userID, fileID)
	require.NoError(t, err)

	require.NotNil(t, cont)
}

func TestRenameFile(t *testing.T) {
	t.Skip("not implemented yet")
}

func TestMoveFile(t *testing.T) {
	t.Skip("not implemented yet")
}

