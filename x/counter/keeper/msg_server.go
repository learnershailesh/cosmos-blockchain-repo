package keeper 

import (
	"context"
	"github.com/cosmos/example/x/counter/types"
)

type msgServer struct {
	*Keeper
}

func NewMsgServerImpl(k *Keeper) types.MsgServer {
	return &msgServer{k}
}

func (m msgServer) Add(ctx context.Context, req *types.MsgAddRequest)(*types.MsgAddResponse, error) {
	
	//Delegate the state update to the keeper
	newCount, err := m.AddCount(ctx, req.GetAdd())
	if err != nil {
		return nil, err
	}

	//return the updated count to the caller.
	return &types.MsgAddResponse{UpdatedCount: newCount}, nil
}