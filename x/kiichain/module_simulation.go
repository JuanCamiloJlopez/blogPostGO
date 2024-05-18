package kiichain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"kiichain/testutil/sample"
	kiichainsimulation "kiichain/x/kiichain/simulation"
	"kiichain/x/kiichain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = kiichainsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateTile = "op_weight_msg_tile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTile int = 100

	opWeightMsgUpdateTile = "op_weight_msg_tile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateTile int = 100

	opWeightMsgDeleteTile = "op_weight_msg_tile"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteTile int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	kiichainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		TileList: []types.Tile{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		TileCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&kiichainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateTile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTile, &weightMsgCreateTile, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTile = defaultWeightMsgCreateTile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTile,
		kiichainsimulation.SimulateMsgCreateTile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateTile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateTile, &weightMsgUpdateTile, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateTile = defaultWeightMsgUpdateTile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateTile,
		kiichainsimulation.SimulateMsgUpdateTile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteTile int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteTile, &weightMsgDeleteTile, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteTile = defaultWeightMsgDeleteTile
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteTile,
		kiichainsimulation.SimulateMsgDeleteTile(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateTile,
			defaultWeightMsgCreateTile,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				kiichainsimulation.SimulateMsgCreateTile(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateTile,
			defaultWeightMsgUpdateTile,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				kiichainsimulation.SimulateMsgUpdateTile(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteTile,
			defaultWeightMsgDeleteTile,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				kiichainsimulation.SimulateMsgDeleteTile(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
