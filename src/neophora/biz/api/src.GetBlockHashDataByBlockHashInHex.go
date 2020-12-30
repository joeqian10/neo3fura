package api

import (
	"neophora/lib/type/h256"
	"neophora/var/stderr"
)

// GetBlockHashDataByBlockHashInHex ...
// as an example:
//
// ```
// TODO
// ```
func (me *T) GetBlockHashDataByBlockHashInHex(args struct {
	BlockHash h256.T
}, ret *string) error {
	if args.BlockHash.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	return me.Data.GetArgsInHex(struct {
		Target string
		Index  string
		Keys   []string
	}{
		Target: "bins.hdt",
		Index:  "h256.blk",
		Keys:   []string{args.BlockHash.Val()},
	}, ret)
}
