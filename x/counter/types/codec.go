package types 

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"

)


func RegisterInterfaces(registry codectypes.InterfaceRegistry){

	//Registe the MsgAddRequest as an sdk.Msg so the app can recognize it from the transactions
	registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgAddRequest{},
  )

  msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}  