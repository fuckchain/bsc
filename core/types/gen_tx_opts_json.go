// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package types

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// MarshalJSON marshals as JSON.
func (t TransactionOpts) MarshalJSON() ([]byte, error) {
	type TransactionOpts struct {
		KnownAccounts  KnownAccounts   `json:"knownAccounts"`
		BlockNumberMin *hexutil.Uint64 `json:"blockNumberMin,omitempty"`
		BlockNumberMax *hexutil.Uint64 `json:"blockNumberMax,omitempty"`
		TimestampMin   *hexutil.Uint64 `json:"timestampMin,omitempty"`
		TimestampMax   *hexutil.Uint64 `json:"timestampMax,omitempty"`
	}
	var enc TransactionOpts
	enc.KnownAccounts = t.KnownAccounts
	enc.BlockNumberMin = t.BlockNumberMin
	enc.BlockNumberMax = t.BlockNumberMax
	enc.TimestampMin = t.TimestampMin
	enc.TimestampMax = t.TimestampMax
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (t *TransactionOpts) UnmarshalJSON(input []byte) error {
	type TransactionOpts struct {
		KnownAccounts  *KnownAccounts  `json:"knownAccounts"`
		BlockNumberMin *hexutil.Uint64 `json:"blockNumberMin,omitempty"`
		BlockNumberMax *hexutil.Uint64 `json:"blockNumberMax,omitempty"`
		TimestampMin   *hexutil.Uint64 `json:"timestampMin,omitempty"`
		TimestampMax   *hexutil.Uint64 `json:"timestampMax,omitempty"`
	}
	var dec TransactionOpts
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.KnownAccounts != nil {
		t.KnownAccounts = *dec.KnownAccounts
	}
	if dec.BlockNumberMin != nil {
		t.BlockNumberMin = dec.BlockNumberMin
	}
	if dec.BlockNumberMax != nil {
		t.BlockNumberMax = dec.BlockNumberMax
	}
	if dec.TimestampMin != nil {
		t.TimestampMin = dec.TimestampMin
	}
	if dec.TimestampMax != nil {
		t.TimestampMax = dec.TimestampMax
	}
	return nil
}