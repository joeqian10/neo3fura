package api

import (
	"encoding/json"
	"neophora/lib/type/bins"
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetSpentNEOByTransactionHashOutputIndexInJSON ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetSpentNEOByTransactionHashOutputIndexInJSON(args struct {
	TransactionHash h256.T
	OutputIndex     uintval.T
}, ret *json.RawMessage) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	var result bins.T
	if err := me.Data.GetLatestUint64ValInBins(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "u128.spt",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.TransactionHash.Val(), args.OutputIndex.Hex()},
	}, &result); err != nil {
		return err
	}
	if result.Valid() == false {
		return stderr.ErrNotFound
	}
	js, err := result.JSONViaSpentNEO()
	if err != nil {
		return stderr.ErrNotFound
	}
	*ret = js
	return nil
}
