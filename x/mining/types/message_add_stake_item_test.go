package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/warmage-sports/katana/testutil/sample"
	"github.com/warmage-sports/katana/utils"
	"github.com/warmage-sports/katana/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestMsgAddStakeItem_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgAddStakeItem
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgAddStakeItem{
				Creator:         "invalid_address",
				PowerRewardRate: utils.MustNewDecFromStr("1.8"),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgAddStakeItem{
				Creator:         sample.AccAddress(),
				PowerRewardRate: utils.MustNewDecFromStr("0.5"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
