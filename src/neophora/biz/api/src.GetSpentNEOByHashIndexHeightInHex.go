package api

import (
	"neophora/lib/type/h256"
	"neophora/lib/type/uintval"
	"neophora/var/stderr"
)

// GetSpentNEOByHashIndexHeightInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetSpentNEOByHashIndexHeightInHex(args struct {
	Hash   h256.T
	Index  uintval.T
	Height uintval.T
}, ret *string) error {
	if args.Hash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.Index.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	if args.Height.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetLastValInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "u128.spt",
		Index:  "h256.trx-uint.num-uint.hgt",
		Keys:   []string{args.Hash.Val(), args.Index.Hex(), args.Height.Hex()},
	}, ret)
}
