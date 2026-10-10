package keeper 


import (
	"context"
	"errors"
	"cosmossdk.io/collections"
	"cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/example/x/counter/types"
)

type Keeper struct {
	Schema collections.Schema
	counter collections.Item[uint64]
}

func NewKeeper(storeService store.KVStoreService, cdc codec.Codec) *Keeper {
	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{

		//Store the counter under prefix 0 in this module's KV store
	counter: collections.NewItem(sb, collections.NewPrefix(0), "counter", collections.Uint64Value),
	}
	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	//k.Schema = schema 
	k.Schema = schema
	return &k
}

func (k Keeper) GetCount(ctx context.Context) (uint64, error) {
	

	// Treat missing state as zero so a fresh chain starts cleanly.
	counter, err := k.counter.Get(ctx)

	if err != nil && !errors.Is(err, collections.ErrNotFound) {
		return 0, err
	}
	return counter, nil
}

func (k Keeper) AddCount(ctx context.Context, amount uint64) (uint64, error) {
	count, err := k.GetCount(ctx)
	if err != nil {
		return 0, err
	}

	//Increment the current count and write it back to the store 
	newCount := count + amount
	return newCount, k.counter.Set(ctx, newCount)
}

func (k *Keeper) InitGenesis(ctx context.Context, gs *types.GenesisState) error {
	return k.counter.Set(ctx, gs.Count)
}

func (k *Keeper) ExportGenesis(ctx context.Context) (*types.GenesisState, error){
	count, err := k.GetCount(ctx)
	if err != nil {
		return nil, err
	}
	return &types.GenesisState{Count: count}, nil 
}