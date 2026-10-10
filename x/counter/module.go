// package counter

// import  (
// 	"context"
// 	"encoding/json"
// 	"cosmossdk.io/core/appmodule"
// 	"github.com/cosmos/cosmos-sdk/client"
// 	"github.com/cosmos/cosmos-sdk/codec"
// 	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/cosmos/cosmos-sdk/types/module"
// 	"github.com/grpc-ecosystem/grpc-gateway/runtime"

// 	"github.com/cosmos/example/x/counter/keeper"
// 	countertypes "github.com/cosmos/example/x/counter/types"

// )

// var (

// 	//compile time checks that AppModule implements the required module interfaces
// 	_ appmodule.AppModule   = AppModule{}
// 	_ module.HasConsensusVersion = AppModule{}
// 	_ module.HasGenesis = AppModule{}
// 	_ module.HasServices = AppModule{}
// )

// type AppModule struct {
// 	cdc codec.Codec
// }

// func (a AppModuleBasic) Name() string {return countertypes.ModuleName}

// func (a AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}

// func (a AppModuleBasic) RegisterInterfaces(registry codecTypes.InterfaceRegistry) {
// 	countertypes.RegisterInterfaces(registry)
// }


// func (a AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
// 	return cdc.MustMarshalJSON(&countertypes.GenesisState{Count: 0})
// }



// func (a AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
// 	gs := countertypes.GenesisState{}
// 	return cdc.UnmarshalJSON(bz, &gs)
// }


// func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux){

// 	if err := countertypes.RegisterQueryHandlerClient(context.Background(), mux, countertypes.NewQueryClient(clientCtx)); err != nil {
// 		panic(err)
// 	}
// }

// type AppModule struct {
// 	AppModuleBasic
// 	keeper *keeper.Keeper
// }

// func NewAppModule(cdc codec.Codec, k *keeper.Keeper) AppModule {
// 	return AppModule{AppModuleBasic: AppModuleBasic{cdc: cdc}, keeper: k}
// }

// func (a AppModule ) IsOnePerModuleType() {}
// func (a AppModule) IsAppModule() {}


// func (a AppModule) ConsensusVersion() uint64 {
// 	return 1
// }

// func (a AppModule) RegisterServices(cfg module.Configurator) {


// 	//connect the generate sercvice interfaces to your keeper-backend implementations.
// 	countertypes.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImp(a.keeper))
// 	countertypes.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(a.keeper))
// }

// func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) {
// 	gs := &countertypes.GenesisState{}
// 	cdc.MustUnmarshalJSON(bz, gs)


// 	//Load the initial counter Value into state at chain start
// 	if err := a.keeper.InitGenesis(ctx, gs); err != nil {
// 		panic(err)
// 	}
// }

// func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RAWMessage {
// 	gs, err := a.keeper.ExportGenesis(ctx)
// 	if err != nil {
// 		panic(err)
// 	}


// 	//Write the current counter value back out for exports.
// 	return cdc.MustMarshalJSON(gs)
// }


// x/counter/module.go
package counter

import (
    "context"
    "encoding/json"

    "cosmossdk.io/core/appmodule"
    "github.com/cosmos/cosmos-sdk/client"
    "github.com/cosmos/cosmos-sdk/codec"
    codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
    sdk "github.com/cosmos/cosmos-sdk/types"
    "github.com/cosmos/cosmos-sdk/types/module"
    "github.com/grpc-ecosystem/grpc-gateway/runtime"

    "github.com/cosmos/example/x/counter/keeper"
    countertypes "github.com/cosmos/example/x/counter/types"
)

var (
    // Compile-time checks that AppModule implements the required module interfaces.
    _ appmodule.AppModule        = AppModule{}
    _ module.HasConsensusVersion = AppModule{}
    _ module.HasGenesis          = AppModule{}
    _ module.HasServices         = AppModule{}
)

type AppModuleBasic struct {
    cdc codec.Codec
}

func (a AppModuleBasic) Name() string { return countertypes.ModuleName }

func (a AppModuleBasic) RegisterLegacyAminoCodec(*codec.LegacyAmino) {}

func (a AppModuleBasic) RegisterInterfaces(registry codecTypes.InterfaceRegistry) {
    countertypes.RegisterInterfaces(registry)
}

func (a AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
    // Start the module with a zero counter by default.
    return cdc.MustMarshalJSON(&countertypes.GenesisState{Count: 0})
}

func (a AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
    gs := countertypes.GenesisState{}
    return cdc.UnmarshalJSON(bz, &gs)
}

func (a AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
    // Expose the Query service through the HTTP gateway.
    if err := countertypes.RegisterQueryHandlerClient(context.Background(), mux, countertypes.NewQueryClient(clientCtx)); err != nil {
        panic(err)
    }
}

type AppModule struct {
    AppModuleBasic
    keeper *keeper.Keeper
}

func NewAppModule(cdc codec.Codec, k *keeper.Keeper) AppModule {
    return AppModule{AppModuleBasic: AppModuleBasic{cdc: cdc}, keeper: k}
}

func (a AppModule) IsOnePerModuleType() {}
func (a AppModule) IsAppModule()        {}

func (a AppModule) ConsensusVersion() uint64 { return 1 }

func (a AppModule) RegisterServices(cfg module.Configurator) {
    // Connect the generated service interfaces to your keeper-backed implementations.
    countertypes.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(a.keeper))
    countertypes.RegisterQueryServer(cfg.QueryServer(), keeper.NewQueryServer(a.keeper))
}

func (a AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) {
    gs := &countertypes.GenesisState{}
    cdc.MustUnmarshalJSON(bz, gs)
    // Load the initial counter value into state at chain start.
    if err := a.keeper.InitGenesis(ctx, gs); err != nil {
        panic(err)
    }
}

func (a AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
    gs, err := a.keeper.ExportGenesis(ctx)
    if err != nil {
        panic(err)
    }
    // Write the current counter value back out for exports.
    return cdc.MustMarshalJSON(gs)
}