package vpn

import (
	"fmt"
	"time"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"

	"github.com/ironman0x7b2/sentinel-sdk/types"
)

type NodeDetails struct {
	Owner           csdkTypes.AccAddress
	LockedAmount    csdkTypes.Coin
	APIPort         uint16
	NetSpeed        types.Bandwidth
	EncMethod       string
	PerGBAmount     csdkTypes.Coins
	NodeType        string
	Version         string
	Status          string
	StatusAtHeight  int64
	UpdatedAtHeight int64
}

type SessionDetails struct {
	VPNOwnerAddress    csdkTypes.AccAddress
	ClientAddress      csdkTypes.AccAddress
	PerGBAmount        csdkTypes.Coins
	BandwidthToProvide types.Bandwidth
	BandwidthConsumed  types.Bandwidth
	StartTime          *time.Time
	EndTime            *time.Time
	Status             string
}

func NodesCountKey(accAddress csdkTypes.AccAddress) string {
	return fmt.Sprintf("nodes_count/%s", accAddress.String())
}

func SessionsCountKey(accAddress csdkTypes.AccAddress) string {
	return fmt.Sprintf("sessions_count/%s", accAddress.String())
}

func NodeKey(accAddress csdkTypes.AccAddress, count uint64) string {
	return fmt.Sprintf("%s/%d", accAddress.String(), count)
}

const (
	KeyActiveNodeIDs    = "ACTIVE_NODE_IDS"
	KeyActiveSessionIDs = "ACTIVE_SESSION_IDS"
	KeyNodeOwners       = "NODE_OWNERS"

	StoreKeySession = "vpn_session"
	StoreKeyNode    = "vpn_node"

	RouteKey = "vpn"

	StatusRegistered   = "REGISTERED"
	StatusActive       = "ACTIVE"
	StatusInactive     = "INACTIVE"
	StatusDeregistered = "DEREGISTERED"
	StatusStart        = "STARTED"
	StatusEnd          = "ENDED"
)
