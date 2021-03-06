package repo

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"github.com/bitstored/file-service/pkg/schema"
	cdb "github.com/bitstored/repository/pkg/couchbaserepo"
	"github.com/couchbase/gocb"
)

type DriveID string

const (
	connSpecStr     = "http://localhost"
	username        = "Administrator"
	password        = "password"
	bucketName      = "bitstored"
	bucketPassword  = ""
	idLen           = 16
	driveIDFormat   = "drid:%v"
	folderIDFormat  = "fdid:%v"
	fileIDFormat    = "flid:%v"
	contentIDFormat = "cid:%v"
	levelID         = "level"
	sharedID        = "shared"
	ownerKey        = "owner"
	timeKey         = "createdAt"
)

type Repository struct {
	next *cdb.Repository
	mux  *sync.RWMutex
}

func newID(format string) (string, error) {
	ID := make([]byte, idLen)
	if _, err := rand.Read(ID); err != nil {
		return "", err
	}
	encoded := base64.RawStdEncoding.EncodeToString(ID)
	value := fmt.Sprintf(format, encoded)
	return value, nil
}

func (r *Repository) CreateDrive(ctx context.Context, userID string) (driveID DriveID, err error) {
	idValue, err := newID(driveIDFormat)
	driveID = DriveID(idValue)
	if err != nil {
		return "", fmt.Errorf("failed to create drive id")
	}
	drive := schema.Folder{ID: idValue,
		ParentID:    "",
		OwnerID:     userID,
		Name:        "Drive",
		Created:     time.Now().String(),
		IsDeletable: false}
	r.mux.Lock()
	if _, err = r.next.Create(ctx, idValue, drive); err != nil {
		r.mux.Unlock()
		return "", err
	}
	r.mux.Unlock()
	arr := make([]string, 0)
	level := schema.FSLevel{RootID: idValue, FilesID: arr, FoldersID: arr}
	r.mux.Lock()

	if _, err = r.next.Create(ctx, levelID+idValue, level); err != nil {
		r.mux.Unlock()
		return "", err
	}
	userData := schema.UserData{DriveID: idValue}
	if _, err = r.next.Create(ctx, userID, userData); err != nil {
		r.mux.Unlock()
		return "", err
	}
	r.mux.Unlock()
	return
}

func (r *Repository) ComputeSize(ctx context.Context, userID string) (int64, int64, error) {
	data := new(schema.UserData)
	_, err := r.next.Read(ctx, userID, data)
	if err != nil {
		return -1, -1, fmt.Errorf("failed to read data for user %s: %v", userID, err.Error())
	}
	folderID := data.DriveID
	is, cs, err := r.computeSizes(ctx, folderID)
	if err != nil {
		return 0, 0, err
	}
	return is, cs, nil
}

func (r *Repository) computeSizes(ctx context.Context, folderID string) (int64, int64, error) {
	level := new(schema.FSLevel)
	_, err := r.next.Read(ctx, levelID+folderID, &level)
	if err != nil {
		fmt.Printf("failed to read folder content %s: %v", folderID, err.Error())
	}
	iSize := int64(0)
	cSize := int64(0)
	for _, fileID := range level.FilesID {
		file := new(schema.File)
		_, err = r.next.Read(ctx, fileID, file)
		if err != nil {
			fmt.Printf("failed to read file info %s: %v", fileID, err.Error())
		}
		iSize += file.InitialSize
		cSize += file.CompressedSize
	}
	for _, folderID1 := range level.FoldersID {
		folder := new(schema.Folder)
		_, err = r.next.Read(ctx, folderID1, folder)
		if err != nil {
			fmt.Printf("failed to read file info %s: %v", folderID1, err.Error())
		}
		is, cs, err := r.computeSizes(ctx, folder.ID)
		if err != nil {
			fmt.Printf("failed to read file info %s: %v", folderID1, err.Error())
		}
		iSize += is
		cSize += cs
	}

	return iSize, cSize, nil
}

func (r *Repository) GetMyDriveId(ctx context.Context, userID string) (string, error) {
	data := new(schema.UserData)
	_, err := r.next.Read(ctx, userID, data)
	if err != nil {
		return "", fmt.Errorf("failed to read data for user %s: %v", userID, err.Error())
	}
	return data.DriveID, nil
}

func (r *Repository) CreateNewFolder(ctx context.Context, userID, name, destinationID string) (folderID string, err error) {
	folderID, err = newID(folderIDFormat)
	if err != nil {
		return "", fmt.Errorf("failed to create folder id")
	}
	folder := schema.Folder{ID: folderID,
		ParentID:    destinationID,
		OwnerID:     userID,
		Name:        name,
		Created:     time.Now().String(),
		IsDeletable: true}
	if _, err = r.next.Create(ctx, folderID, folder); err != nil {
		return "", err
	}
	arr := make([]string, 0)
	level := schema.FSLevel{RootID: folderID, FilesID: arr, FoldersID: arr}

	r.mux.Lock()
	if _, err = r.next.Create(ctx, levelID+folderID, level); err != nil {
		r.mux.Unlock()
		return "", err
	}
	r.mux.Unlock()

	rLevel := new(schema.FSLevel)
	r.mux.RLock()
	cas, err := r.next.Read(ctx, levelID+destinationID, rLevel)
	r.mux.RUnlock()

	if err != nil {
		return "", fmt.Errorf("failed to read FSLevel")
	}
	rLevel.FoldersID = append(rLevel.FoldersID, folderID)

	r.mux.Lock()
	_, err = r.next.Update(ctx, levelID+destinationID, cas, rLevel)
	if err != nil {
		r.mux.Unlock()
		return "", fmt.Errorf("failed to read FSLevel")
	}
	r.mux.Unlock()
	return
}

func (r *Repository) CreateNewFile(ctx context.Context, userID string, name string, destinationID string, fileType string, writable bool, private bool, content []byte, iSize, cSize int64) (fileID string, err error) {
	fileID, err = newID(fileIDFormat)
	if err != nil {
		return "", fmt.Errorf("failed to create file id")
	}
	contentID, err := newID(contentIDFormat)
	if err != nil {
		return "", fmt.Errorf("failed to create content id")
	}
	file := schema.File{
		ID:             fileID,
		ParentID:       destinationID,
		OwnerID:        userID,
		Name:           name,
		Created:        time.Now().String(),
		Type:           fileType,
		IsWritable:     writable,
		IsPrivate:      private,
		ContentID:      contentID,
		InitialSize:    iSize,
		CompressedSize: cSize,
	}

	contentData := schema.Content{ID: contentID,
		Data: content}
	var cas uint
	if cas, err = r.next.Create(ctx, contentID, contentData); err != nil {
		return "", err
	}
	if _, err = r.next.Create(ctx, fileID, file); err != nil {
		r.next.Delete(ctx, contentID, cas)
		return "", err
	}

	rLevel := new(schema.FSLevel)

	r.mux.RLock()
	cas1, err := r.next.Read(ctx, levelID+destinationID, rLevel)
	r.mux.RUnlock()

	if err != nil {
		return "", fmt.Errorf("failed to read FSLevel")
	}
	rLevel.FilesID = append(rLevel.FilesID, fileID)

	r.mux.Lock()
	_, err = r.next.Update(ctx, levelID+destinationID, cas1, rLevel)
	if err != nil {
		r.mux.Unlock()
		return "", fmt.Errorf("failed to read FSLevel")
	}
	r.mux.Unlock()
	return
}
func (r *Repository) UploadFile(ctx context.Context, userID string, name string, destinationID string, fileType string, writable bool, private bool, content []byte, iSize, cSize int64) (fileID string, err error) {
	fileID, err = newID(fileIDFormat)
	if err != nil {
		return "", fmt.Errorf("failed to create file id")
	}
	contentID, err := newID(contentIDFormat)
	if err != nil {
		return "", fmt.Errorf("failed to create content id")
	}
	file := schema.File{ID: fileID,
		ParentID:       destinationID,
		OwnerID:        userID,
		Name:           name,
		Created:        time.Now().String(),
		Type:           fileType,
		IsWritable:     writable,
		IsPrivate:      private,
		ContentID:      contentID,
		InitialSize:    iSize,
		CompressedSize: cSize,
	}
	contentData := schema.Content{ID: contentID,
		Data: content}
	var cas uint
	if cas, err = r.next.Create(ctx, contentID, contentData); err != nil {
		return "", err
	}
	if _, err = r.next.Create(ctx, fileID, file); err != nil {
		r.next.Delete(ctx, contentID, cas)
		return "", err
	}
	rLevel := new(schema.FSLevel)

	r.mux.RLock()
	cas1, err := r.next.Read(ctx, levelID+destinationID, rLevel)
	r.mux.RUnlock()

	if err != nil {
		return "", fmt.Errorf("failed to read FSLevel")
	}
	rLevel.FilesID = append(rLevel.FilesID, fileID)

	r.mux.Lock()
	_, err = r.next.Update(ctx, levelID+destinationID, cas1, rLevel)
	if err != nil {
		r.mux.Unlock()
		return "", fmt.Errorf("failed to read FSLevel")
	}
	r.mux.Unlock()
	return
}

func (r *Repository) GetFileContent(ctx context.Context, userID, fileID string) (*schema.Content, error) {
	file := new(schema.File)
	_, err := r.next.Read(ctx, fileID, file)
	if err != nil {
		return nil, fmt.Errorf("failed to read File %s: %v", fileID, err.Error())
	}
	contentID := file.ContentID
	content := new(schema.Content)

	_, err = r.next.Read(ctx, contentID, &content)
	if err != nil {
		return nil, fmt.Errorf("failed to read content %s: %v", contentID, err.Error())
	}
	return content, nil
}
func (r *Repository) GetFolderContent(ctx context.Context, userID, folderID string) (*schema.FSLevelDetailed, error) {
	//levelID+
	level := new(schema.FSLevel)
	_, err := r.next.Read(ctx, levelID+folderID, level)
	if err != nil {
		return nil, fmt.Errorf("failed to read folder content %s: %v", folderID, err.Error())
	}
	folders := make([]schema.Folder, 0)
	files := make([]schema.File, 0)

	for _, fileID := range level.FilesID {
		file := new(schema.File)
		_, err := r.next.Read(ctx, fileID, file)
		if err == nil {
			files = append(files, *file)
		}
	}
	for _, folderID1 := range level.FoldersID {
		folder := new(schema.Folder)
		_, err := r.next.Read(ctx, folderID1, folder)
		if err == nil {
			folders = append(folders, *folder)
		}
	}

	return &schema.FSLevelDetailed{RootID: folderID, Folders: folders, Files: files}, nil
}
func (r *Repository) UpdateFileContent(ctx context.Context, userID string, fileID string, content []byte) (uint, error) {
	file := new(schema.File)
	r.mux.RLock()
	if _, err := r.next.Read(ctx, fileID, file); err != nil {
		r.mux.RUnlock()
		return 0, err
	}
	c := new(schema.Content)
	cas, err := r.next.Read(ctx, file.ContentID, c)
	if err != nil {
		r.mux.RUnlock()
		return 0, err
	}
	r.mux.RUnlock()
	contentData := schema.Content{ID: file.ContentID,
		Data: content}
	r.mux.Lock()
	if cas, err = r.next.Update(ctx, file.ContentID, cas, contentData); err != nil {
		r.mux.Unlock()
		return 0, err
	}
	r.mux.Unlock()
	return 0, nil
}
func (r *Repository) DeleteFile(ctx context.Context, userID string, fileID string) error {
	file := new(schema.File)
	r.mux.RLock()
	var fcas uint
	fcas, err := r.next.Read(ctx, fileID, file)
	if err != nil {
		r.mux.RUnlock()
		return err
	}
	c := new(schema.Content)
	var ccas uint
	ccas, err = r.next.Read(ctx, file.ContentID, c)
	if err != nil {
		r.mux.RUnlock()
		return err
	}
	r.mux.RUnlock()
	r.mux.Lock()
	_, err = r.next.Delete(ctx, fileID, fcas)
	if err != nil {
		r.mux.Unlock()
		return err
	}
	_, err = r.next.Delete(ctx, file.ContentID, ccas)
	if err != nil {
		r.mux.Unlock()
		return err
	}
	r.mux.Unlock()
	return nil
}

func (r *Repository) DownloadFile(ctx context.Context, userID, fileID string) (*schema.File, *schema.Content, error) {
	file := new(schema.File)
	_, err := r.next.Read(ctx, fileID, file)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read File %s: %v", fileID, err.Error())
	}
	contentID := file.ContentID
	content := new(schema.Content)

	_, err = r.next.Read(ctx, contentID, content)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to read content %s: %v", contentID, err.Error())
	}

	return file, content, nil
}

func (r *Repository) RenameFile(ctx context.Context, userID, fileID, newName string) error {
	file := new(schema.File)
	r.mux.RLock()
	cas, err := r.next.Read(ctx, fileID, file)
	if err != nil {
		r.mux.RUnlock()
		return err
	}
	r.mux.RUnlock()
	file.Name = newName
	r.mux.Lock()
	if _, err := r.next.Update(ctx, file.ID, cas, file); err != nil {
		r.mux.Unlock()
		return err
	}
	r.mux.Unlock()
	return nil
}

func (r *Repository) MoveFile(ctx context.Context, targetID, sourceID, destinationID string, copy bool) error {
	slevel := new(schema.FSLevel)
	dlevel := new(schema.FSLevel)
	r.mux.RLock()
	scas, err := r.next.Read(ctx, levelID+sourceID, &slevel)
	if err != nil {
		r.mux.RUnlock()
		return fmt.Errorf("failed to read folder content %s: %v", sourceID, err.Error())
	}
	dcas, err := r.next.Read(ctx, levelID+destinationID, &dlevel)
	if err != nil {
		r.mux.RUnlock()
		return fmt.Errorf("failed to read folder content %s: %v", sourceID, err.Error())
	}
	r.mux.RUnlock()

	var index int = -1
	for i := range slevel.FilesID {
		if slevel.FilesID[i] == targetID {
			index = i
		}
	}
	if index != -1 {
		dlevel.FilesID = append(dlevel.FilesID, slevel.FilesID[index])
		slevel.FilesID = append(slevel.FilesID[0:index], slevel.FilesID[index:]...)
	}
	index = -1
	for i := range slevel.FoldersID {
		if slevel.FoldersID[i] == targetID {
			index = i
		}
	}
	if index != -1 {
		dlevel.FoldersID = append(dlevel.FoldersID, slevel.FoldersID[index])
		slevel.FoldersID = append(slevel.FoldersID[0:index], slevel.FoldersID[index:]...)
	}

	r.mux.Lock()
	if !copy {
		_, err = r.next.Update(ctx, levelID+sourceID, scas, slevel)
		if err != nil {
			r.mux.Unlock()
			return fmt.Errorf("failed to read FSLevel")
		}
	}
	_, err = r.next.Update(ctx, levelID+destinationID, dcas, dlevel)
	if err != nil {
		r.mux.Unlock()
		return fmt.Errorf("failed to read FSLevel")
	}
	r.mux.Unlock()
	return nil
}

func NewFileRepository() (*Repository, error) {
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
	var m sync.RWMutex
	return &Repository{next: cdb.NewRepository(bucket), mux: &m}, nil
}
