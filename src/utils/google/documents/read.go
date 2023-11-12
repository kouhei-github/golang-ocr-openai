package documents

import (
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/option"
	"net-http/myapp/utils/google"
)

type DocumentFacade struct {
	Srv         *docs.Service
	DriveFileId string
}

func NewDocumentFacade(client *google.Client, driveFileId string) (*DocumentFacade, error) {
	srv, err := docs.NewService(client.Context, option.WithHTTPClient(client.Client))
	if err != nil {
		return &DocumentFacade{}, err
	}
	return &DocumentFacade{Srv: srv, DriveFileId: driveFileId}, nil
}

func (r DocumentFacade) Read() (string, error) {
	doc, err := r.Srv.Documents.Get(r.DriveFileId).Do()
	if err != nil {
		return "", err
	}
	return getTextFromDocument(doc.Body.Content), nil
}

func getTextFromDocument(contents []*docs.StructuralElement) string {
	text := ""
	for _, content := range contents {
		if content.Paragraph != nil {
			for _, element := range content.Paragraph.Elements {
				if element.TextRun != nil {
					text = text + element.TextRun.Content
				}
			}
		}
	}

	return text
}
