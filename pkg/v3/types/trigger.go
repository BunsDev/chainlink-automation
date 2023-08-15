package types

import (
	"bytes"
	"fmt"
)

// Trigger represents a trigger for an upkeep.
// It contains an extension per trigger type, and the block number + hash
// in which the trigger was checked.
// NOTE: This struct is sent on the p2p network as part of observations to get quorum
// Any change here should be backwards compatible and should keep validation and
// quorum requirements in mind. Please ensure to get a proper review along with an
// upgrade plan before changing this
type Trigger struct {
	// BlockNumber is the block number in which the trigger was checked
	BlockNumber BlockNumber
	// BlockHash is the block hash in which the trigger was checked
	BlockHash [32]byte
	// LogTriggerExtension is the extension for log triggers
	LogTriggerExtension *LogTriggerExtension
}

// NewTrigger returns a new basic trigger w/o extension
func NewTrigger(blockNumber BlockNumber, blockHash [32]byte) Trigger {
	return Trigger{
		BlockNumber: blockNumber,
		BlockHash:   blockHash,
	}
}

func NewLogTrigger(blockNumber BlockNumber, blockHash [32]byte, logTriggerExtension *LogTriggerExtension) Trigger {
	return Trigger{
		BlockNumber:         blockNumber,
		BlockHash:           blockHash,
		LogTriggerExtension: logTriggerExtension,
	}
}

// LogTriggerExtension is the extension used for log triggers,
// It contains information of the log event that was triggered.
// NOTE: This struct is sent on the p2p network as part of observations to get quorum
// Any change here should be backwards compatible and should keep validation and
// quorum requirements in mind. Please ensure to get a proper review along with an
// upgrade plan before changing this
type LogTriggerExtension struct {
	// LogTxHash is the transaction hash of the log event
	TxHash [32]byte
	// Index is the index of the log event in the transaction
	Index uint32
	// BlockHash is the block hash in which the event occurred
	// NOTE: This field might be empty. If relying on this field check
	// it is non empty, if it's empty derive from txHash
	BlockHash [32]byte
	// BlockNumber is the block number in which the event occurred
	// NOTE: This field might be empty. If relying on this field check
	// it is non empty, if it's empty derive from txHash
	BlockNumber BlockNumber
}

// LogIdentifier returns a unique identifier for the log event,
// composed of the transaction hash and the log index bytes.
func (e LogTriggerExtension) LogIdentifier() []byte {
	return bytes.Join([][]byte{
		e.TxHash[:],
		[]byte(fmt.Sprintf("%d", e.Index)),
	}, []byte{})
}
