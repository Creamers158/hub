package vpn

import (
	"reflect"

	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

func NewHandler(nk Keeper, bk bank.Keeper) csdkTypes.Handler {
	return func(ctx csdkTypes.Context, msg csdkTypes.Msg) csdkTypes.Result {
		switch msg := msg.(type) {
		case MsgRegisterNode:
			return handleRegisterNode(ctx, nk, bk, msg)
		case MsgUpdateNode:
			return handleUpdateNode(ctx, nk, msg)
		case MsgDeregisterNode:
			return handleDeregisterNode(ctx, nk, bk, msg)
		default:
			return errorUnknownMsgType(reflect.TypeOf(msg).Name()).Result()
		}
	}
}

func handleRegisterNode(ctx csdkTypes.Context, nk Keeper, bk bank.Keeper, msg MsgRegisterNode) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	count, err := nk.GetNodesCount(ctx, msg.From)
	if err != nil {
		return err.Result()
	}

	id := NodeKey(msg.From, count)
	if details, err := nk.GetNodeDetails(ctx, id); true {
		if err != nil {
			return err.Result()
		}
		if details != nil {
			return errorNodeAlreadyExists().Result()
		}
	}

	lockAmount := csdkTypes.Coins{msg.AmountToLock}
	_, tags, err := bk.SubtractCoins(ctx, msg.From, lockAmount)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	details := NodeDetails{
		Owner:          msg.From,
		LockedAmount:   msg.AmountToLock,
		APIPort:        msg.APIPort,
		NetSpeed:       msg.NetSpeed,
		EncMethod:      msg.EncMethod,
		PerGBAmount:    msg.PerGBAmount,
		Version:        msg.Version,
		NodeType:       msg.NodeType,
		Status:         StatusRegistered,
		StatusAtHeight: ctx.BlockHeight(),
	}
	if err := nk.SetNodeDetails(ctx, id, &details); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", []byte(id))

	if err := nk.SetNodesCount(ctx, msg.From, count+1); err != nil {
		return err.Result()
	}

	if err := AddNodeOwner(ctx, nk, msg.From.String()); err != nil {
		return err.Result()
	}

	return csdkTypes.Result{Tags: allTags}
}

func handleUpdateNode(ctx csdkTypes.Context, nk Keeper, msg MsgUpdateNode) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	details, err := nk.GetNodeDetails(ctx, msg.ID)
	if err != nil {
		return err.Result()
	}
	if details == nil {
		return errorNodeNotExists().Result()
	}
	if !details.Owner.Equals(msg.From) {
		return errorUnauthorized().Result()
	}
	if details.Status != StatusRegistered &&
		details.Status != StatusActive &&
		details.Status != StatusInactive {
		return errorInvalidNodeStatus().Result()
	}

	if msg.APIPort != 0 {
		details.APIPort = msg.APIPort
	}
	if msg.NetSpeed.Download != 0 && msg.NetSpeed.Upload != 0 {
		details.NetSpeed = msg.NetSpeed
	}
	if len(msg.EncMethod) != 0 {
		details.EncMethod = msg.EncMethod
	}
	if msg.PerGBAmount != nil && msg.PerGBAmount.Len() != 0 {
		details.PerGBAmount = msg.PerGBAmount
	}
	if len(msg.Version) != 0 {
		details.Version = msg.Version
	}
	details.UpdatedAtHeight = ctx.BlockHeight()

	if err := nk.SetNodeDetails(ctx, msg.ID, details); err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTag("node_id", []byte(msg.ID))

	return csdkTypes.Result{Tags: allTags}
}

func handleDeregisterNode(ctx csdkTypes.Context, nk Keeper, bk bank.Keeper, msg MsgDeregisterNode) csdkTypes.Result {
	allTags := csdkTypes.EmptyTags()

	details, err := nk.GetNodeDetails(ctx, msg.ID)
	if err != nil {
		return err.Result()
	}
	if details == nil {
		return errorNodeNotExists().Result()
	}
	if !msg.From.Equals(details.Owner) {
		return errorUnauthorized().Result()
	}
	if details.Status != StatusRegistered &&
		details.Status != StatusInactive {
		return errorInvalidNodeStatus().Result()
	}

	details.Status = StatusDeregistered
	details.StatusAtHeight = ctx.BlockHeight()
	if err := nk.SetNodeDetails(ctx, msg.ID, details); err != nil {
		return err.Result()
	}
	allTags.AppendTag("node_id", []byte(msg.ID))

	releaseAmount := csdkTypes.Coins{details.LockedAmount}
	_, tags, err := bk.AddCoins(ctx, msg.From, releaseAmount)
	if err != nil {
		return err.Result()
	}
	allTags = allTags.AppendTags(tags)

	return csdkTypes.Result{Tags: allTags}
}
