package keeper

import (
	"context"
	"github.com/cosmos/example/x/counter/types"
)

type queryServer struct {
	*Keeper
}

func NewQueryServer(k *Keeper) types.QueryServer {
	return &queryServer{k}
}


func (q queryServer) Count(ctx context.Context, _ *types.QueryCountRequest) (*types.QueryCountResponse, error) {
	count, err := q.GetCount(ctx)

	if err != nil  {
		return nil, err
	}

	return &types.QueryCountResponse{Count: count}, nil
}