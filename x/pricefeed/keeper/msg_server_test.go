package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/bandprotocol/oracle-consumer/testutil/keeper"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/keeper"
	"github.com/bandprotocol/oracle-consumer/x/pricefeed/types"
)

func TestUpdateParams(t *testing.T) {
	// Initialize the testing environment.
	k, ctx := testkeeper.PriceFeedKeeper(t)
	msgServer := keeper.NewMsgServerImpl(k)

	testCases := []struct {
		name      string
		request   *types.MsgUpdateParams
		expectErr bool
	}{
		{
			name: "set invalid authority",
			request: &types.MsgUpdateParams{
				Authority: "foo",
			},
			expectErr: true,
		},
		{
			name: "set invalid params",
			request: &types.MsgUpdateParams{
				Authority: k.GetAuthority(),
				Params:    types.Params{},
			},
			expectErr: true,
		},
		{
			name: "set full valid params",
			request: &types.MsgUpdateParams{
				Authority: k.GetAuthority(),
				Params:    types.DefaultParams(),
			},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		_, err := msgServer.UpdateParams(ctx, tc.request)
		if tc.expectErr {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
	}
}
