package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"

	"github.com/spellshape/network/x/monitoringc/keeper"
	"github.com/spellshape/network/x/monitoringc/types"
)

func SimulateMsgCreateClient(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateClient{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateClient simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateClient simulation not implemented"), nil, nil
	}
}
