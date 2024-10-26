// interfaces/client_interface.go
package interfaces

import (
	"context"
	"net/http"
	"oapi-codegen-test/client"
)

// ClientInterface defines the methods for the API client.
type ClientInterface interface {
	GetArtistList(ctx context.Context, params *client.GetArtistListParams) (*http.Response, error)
}
