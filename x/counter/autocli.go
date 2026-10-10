package counter

import (	
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)


func (a AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return  &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: "example.counter.Query", 
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{RpcMethod: "Count", Use: "Count", Short: "Query the current counter value"},
			},
		},

	    Tx: &autocliv1.ServiceCommandDescriptor{
		Service: "example.counter.Msg", 
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{RpcMethod: "Add", Use: "add [amount]", Short: "Add to the counter",
				  PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "add"}}},
			},
		},
	}
}