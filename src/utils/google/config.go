package google

import (
	"context"
	"encoding/json"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/docs/v1"
	"google.golang.org/api/drive/v3"
	"net/http"
	"os"
)

const DocumentReadonlyScope = docs.DocumentsReadonlyScope
const DocumentFullAccessScope = docs.DocumentsScope

const DriveReadonlyScope = drive.DriveMetadataReadonlyScope
const DriveFullAccessScope = drive.DriveMetadataScope
const DriveFileScope = drive.DriveFileScope

type Client struct {
	Client  *http.Client
	Context context.Context
}

func getServiceAccountJson() map[string]string {
	return map[string]string{
		"type":                        os.Getenv("TYPE"),
		"project_id":                  os.Getenv("PROJECT_ID"),
		"private_key_id":              os.Getenv("PROJECT_KEY_ID"),
		"private_key":                 os.Getenv("PRIVATE_KEY"),
		"client_email":                os.Getenv("CLIENT_EMAIL"),
		"client_id":                   os.Getenv("CLIENT_ID"),
		"auth_uri":                    os.Getenv("AUTH_URI"),
		"token_uri":                   os.Getenv("TOKEN_URI"),
		"auth_provider_x509_cert_url": os.Getenv("AUTH_PROVIDER_X506_CERT_URI"),
		"client_x509_cert_url":        os.Getenv("CLIENT_X506_CERT_URI"),
		"universe_domain":             os.Getenv("UNIVERSE_DOMAIN"),
	}
}

func NewClient(scope ...string) (*Client, error) {
	serviceAccount := getServiceAccountJson()
	bytes, err := json.Marshal(serviceAccount)
	if err != nil {
		return &Client{}, err
	}

	config, err := google.JWTConfigFromJSON(bytes, scope...)
	if err != nil {
		return &Client{}, err
	}
	context.Background()
	ctx := context.Background()
	client := config.Client(ctx)
	return &Client{Client: client, Context: ctx}, nil
}
