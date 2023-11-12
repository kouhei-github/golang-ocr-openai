package drives

import (
	"github.com/google/uuid"
	"google.golang.org/api/drive/v3"
	"google.golang.org/api/option"
	"io"
	"log"
	"net-http/myapp/utils/google"
)

type DriveFacade struct {
	Srv      *drive.Service
	FileName string
}

func NewDriveFacade(client *google.Client) (*DriveFacade, error) {
	srv, err := drive.NewService(client.Context, option.WithHTTPClient(client.Client))
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
		return &DriveFacade{}, err
	}
	uuid4, err := uuid.NewUUID()
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
		return &DriveFacade{}, err
	}
	return &DriveFacade{Srv: srv, FileName: uuid4.String()}, nil
}

func (r *DriveFacade) Upload(driveId string, file io.Reader) (string, error) {
	uploadedFile, err := r.Srv.Files.Create(&drive.File{
		Name:     r.FileName,
		MimeType: "application/vnd.google-apps.document",
		Parents:  []string{driveId},
	}).Media(file).Do()
	if err != nil {
		return "", err
	}
	return uploadedFile.Id, nil
}

func (r *DriveFacade) Delete(fileId string) error {
	return r.Srv.Files.Delete(fileId).Do()
}
