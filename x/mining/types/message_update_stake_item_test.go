package types_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/warmage-sports/katana/testutil/sample"
	"github.com/warmage-sports/katana/utils"
	"github.com/warmage-sports/katana/x/mining/types"
	"github.com/stretchr/testify/require"
)

func TestMsgUpdateStakeItem_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  types.MsgUpdateStakeItem
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgUpdateStakeItem{
				Creator:         "invalid_address",
				PowerRewardRate: utils.MustNewDecFromStr("0.1"),
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: types.MsgUpdateStakeItem{
				Creator:         sample.AccAddress(),
				PowerRewardRate: utils.MustNewDecFromStr("3.1"),
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
