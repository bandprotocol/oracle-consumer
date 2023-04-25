package types

// IBC events
const (
	EventTypeEncodeCalldataFailed               = "encode_calldata_failed"
	EventTypeSendPacketSuccess                  = "request_bandchain_success"
	EventTypeRequestBandChainFailed             = "request_bandchain_failed"
	EventTypeDecodeBandChainResultFailed        = "decode_bandchain_result_failed"
	EventTypeBandChainAckSuccess                = "bandchain_ack_success"
	EventTypeBandChainAckError                  = "bandchain_ack_error"
	EventTypeBandChainStoreOracleResponsePacket = "bandchain_store_oracle_response_packet"

	AttributeKeyReason             = "reason"
	AttributeKeyData               = "data"
	AttributeKeySequence           = "sequence"
	AttributeKeyPortID             = "port_id"
	AttributeKeySourceChannel      = "source_channel"
	AttributeKeyDestinationPort    = "destination_port"
	AttributeKeyDestinationChannel = "destination_channel"
	AttributeKeyTimeoutHeight      = "timeout_height"
	AttributeKeyTimeoutTimestamp   = "timeout_timestamp"
	AttributeKeyRequestID          = "request_id"
	AttributeKeyResolveStatus      = "resolve_status"
	AttributeKeyResult             = "result"
	AttributeKeyResolveTime        = "resolve_time"
)
