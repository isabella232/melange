package message

import "code.google.com/p/goprotobuf/proto"

const (
	// Server message types.
	TypeTransaction   = "ZK-TX"
	TypeResolvedName  = "ZK-RS"
	TypeTransferFunds = "ZK-TF"

	// Client message types.
	TypeRequestFunds = "ZK-RF"
	TypeLookupName   = "ZK-LN"
	TypeRegisterName = "ZK-RN"
	TypeRenewName    = "ZK-RW"
)

func getMessageType(m proto.Message) string {
	switch m.(type) {
	// Server messages
	case *Transaction:
		return TypeTransaction
	case *ResolvedName:
		return TypeResolvedName
	case *TransferFunds:
		return TypeTransferFunds

		// Client messsage
	case *RequestFunds:
		return TypeRequestFunds
	case *LookupName:
		return TypeLookupName
	case *RegisterName:
		return TypeRegisterName
	case *RenewName:
		return TypeRenewName

	default:
		panic("zooko: cannot use getMessageType with a non-zooko message")
	}
}
