package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetUTXOByHashIndexInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetUTXOByHashIndexInHex(args struct {
	TransactionHash h256.T
	OutputIndex     uintval.T
}, ret *string) error {
	if args.TransactionHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.OutputIndex.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.utx",
		Index:  "h256.trx-uint.num",
		Keys:   []string{args.TransactionHash.Val(), args.OutputIndex.Hex()},
	}, ret)
}
