//go:generate mockgen -destination mock_kvinterfaces/mock_kvinterfaces.go github.com/Kimenzo/any-sync/commonspace/object/keyvalue/kvinterfaces KeyValueService
package kvinterfaces

import (
	"context"

	"storj.io/drpc"

	"github.com/Kimenzo/any-sync/app"
	"github.com/Kimenzo/any-sync/commonspace/object/keyvalue/keyvaluestorage"
	"github.com/Kimenzo/any-sync/commonspace/spacesyncproto"
	"github.com/Kimenzo/any-sync/net/peer"
)

const CName = "common.object.keyvalue"

type KeyValueService interface {
	app.ComponentRunnable
	DefaultStore() keyvaluestorage.Storage
	HandleMessage(ctx context.Context, msg drpc.Message) (err error)
	SyncWithPeer(p peer.Peer) (err error)
	HandleStoreDiffRequest(ctx context.Context, req *spacesyncproto.StoreDiffRequest) (resp *spacesyncproto.StoreDiffResponse, err error)
	HandleStoreElementsRequest(ctx context.Context, stream spacesyncproto.DRPCSpaceSync_StoreElementsStream) (err error)
}
