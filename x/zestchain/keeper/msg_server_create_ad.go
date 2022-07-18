package keeper

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	kk "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)

func (k kk.msgServer) CreateAd(goCtx context.Context, msg *types.MsgCreateAd) (*types.MsgCreateAdResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	adCount, found := k.Keeper.GetAdCount(ctx)
	if !found {
		panic("AdCount not found")
	}
	newIndex := strconv.FormatUint(adCount.counter, 10)
	ad := types.Ad{
		Index:   newIndex,
		Title:   msg.Title,
		Pot:     msg.Pot,
		Url:     msg.Url,
		Msg:     msg.Msg,
		Tags:    msg.Tags,
		Prefs:   msg.Prefs,
		Creator: msg.Creator,
	}
	err := ad.Validate()
	if err != nil {
		return nil, err
	}
	k.Keeper.SetAd(ctx, ad)

	adCount.counter++
	k.Keeper.SetAdCount(ctx, adCount)

	return &types.MsgCreateAdResponse{
		IdValue: newIndex,
	}, nil
}
