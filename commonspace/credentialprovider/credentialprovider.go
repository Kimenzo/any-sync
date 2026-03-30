//go:generate mockgen -destination mock_credentialprovider/mock_credentialprovider.go github.com/Kimenzo/any-sync/commonspace/credentialprovider CredentialProvider
package credentialprovider

import (
	"context"
	"github.com/Kimenzo/any-sync/app"
	"github.com/Kimenzo/any-sync/commonspace/spacesyncproto"
)

const CName = "common.commonspace.credentialprovider"

func NewNoOp() CredentialProvider {
	return &noOpProvider{}
}

type CredentialProvider interface {
	app.Component
	GetCredential(ctx context.Context, spaceHeader *spacesyncproto.RawSpaceHeaderWithId) ([]byte, error)
}

type noOpProvider struct {
}

func (n noOpProvider) Init(a *app.App) (err error) {
	return nil
}

func (n noOpProvider) Name() (name string) {
	return CName
}

func (n noOpProvider) GetCredential(ctx context.Context, spaceHeader *spacesyncproto.RawSpaceHeaderWithId) ([]byte, error) {
	return nil, nil
}
