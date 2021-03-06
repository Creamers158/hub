// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/sentinel-official/hub/x/plan/types
// ALIASGEN: github.com/sentinel-official/hub/x/plan/keeper
// ALIASGEN: github.com/sentinel-official/hub/x/plan/querier
package plan

import (
	"github.com/sentinel-official/hub/x/plan/keeper"
	"github.com/sentinel-official/hub/x/plan/querier"
	"github.com/sentinel-official/hub/x/plan/types"
)

const (
	Codespace             = types.Codespace
	EventTypeSetCount     = types.EventTypeSetCount
	EventTypeSet          = types.EventTypeSet
	EventTypeSetStatus    = types.EventTypeSetStatus
	EventTypeAddNode      = types.EventTypeAddNode
	EventTypeRemoveNode   = types.EventTypeRemoveNode
	AttributeKeyAddress   = types.AttributeKeyAddress
	AttributeKeyID        = types.AttributeKeyID
	AttributeKeyCount     = types.AttributeKeyCount
	AttributeKeyStatus    = types.AttributeKeyStatus
	ModuleName            = types.ModuleName
	QuerierRoute          = types.QuerierRoute
	QueryPlan             = types.QueryPlan
	QueryPlans            = types.QueryPlans
	QueryPlansForProvider = types.QueryPlansForProvider
	QueryNodesForPlan     = types.QueryNodesForPlan
)

var (
	// functions aliases
	RegisterCodec                  = types.RegisterCodec
	ErrorMarshal                   = types.ErrorMarshal
	ErrorUnmarshal                 = types.ErrorUnmarshal
	ErrorUnknownMsgType            = types.ErrorUnknownMsgType
	ErrorUnknownQueryType          = types.ErrorUnknownQueryType
	ErrorInvalidField              = types.ErrorInvalidField
	ErrorProviderDoesNotExist      = types.ErrorProviderDoesNotExist
	ErrorPlanDoesNotExist          = types.ErrorPlanDoesNotExist
	ErrorNodeDoesNotExist          = types.ErrorNodeDoesNotExist
	ErrorUnauthorized              = types.ErrorUnauthorized
	NewGenesisState                = types.NewGenesisState
	DefaultGenesisState            = types.DefaultGenesisState
	PlanKey                        = types.PlanKey
	GetPlanForProviderKeyPrefix    = types.GetPlanForProviderKeyPrefix
	PlanForProviderKey             = types.PlanForProviderKey
	GetNodeForPlanKeyPrefix        = types.GetNodeForPlanKeyPrefix
	NodeForPlanKey                 = types.NodeForPlanKey
	NewMsgAdd                      = types.NewMsgAdd
	NewMsgSetStatus                = types.NewMsgSetStatus
	NewMsgAddNode                  = types.NewMsgAddNode
	NewMsgRemoveNode               = types.NewMsgRemoveNode
	NewQueryPlanParams             = types.NewQueryPlanParams
	NewQueryPlansParams            = types.NewQueryPlansParams
	NewQueryPlansForProviderParams = types.NewQueryPlansForProviderParams
	NewQueryNodesForPlanParams     = types.NewQueryNodesForPlanParams
	NewKeeper                      = keeper.NewKeeper
	Querier                        = querier.Querier

	// variable aliases
	ModuleCdc                = types.ModuleCdc
	RouterKey                = types.RouterKey
	StoreKey                 = types.StoreKey
	EventModuleName          = types.EventModuleName
	CountKey                 = types.CountKey
	PlanKeyPrefix            = types.PlanKeyPrefix
	PlanForProviderKeyPrefix = types.PlanForProviderKeyPrefix
	NodeForPlanKeyPrefix     = types.NodeForPlanKeyPrefix
)

type (
	GenesisPlan                 = types.GenesisPlan
	GenesisPlans                = types.GenesisPlans
	GenesisState                = types.GenesisState
	MsgAdd                      = types.MsgAdd
	MsgSetStatus                = types.MsgSetStatus
	MsgAddNode                  = types.MsgAddNode
	MsgRemoveNode               = types.MsgRemoveNode
	Plan                        = types.Plan
	Plans                       = types.Plans
	QueryPlanParams             = types.QueryPlanParams
	QueryPlansParams            = types.QueryPlansParams
	QueryPlansForProviderParams = types.QueryPlansForProviderParams
	QueryNodesForPlanParams     = types.QueryNodesForPlanParams
	Keeper                      = keeper.Keeper
)
