package api

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"neophora/lib/type/h160"
	"neophora/var/stderr"
)

func (me *T) GetAssetsHeldByAddress(args struct {
	Address h160.T
	Limit   int64
	Skip    int64
}, ret *json.RawMessage) error {
	if args.Address.Valid() == false {
		return stderr.ErrInvalidArgs
	}
	// Step1
	var r1 map[string]interface{}
	r1, err := me.Data.Client.QueryOne(struct {
		Collection string
		Index      string
		Sort       bson.M
		Filter     bson.M
		Query      []string
	}{
		Collection: "Address",
		Index:      "someIndex",
		Sort:       bson.M{},
		Filter:     bson.M{"address": args.Address.Val()},
		Query:      []string{"_id"},
	}, ret)
	if err != nil {
		return err
	}
	// Step2
	r2, err := me.Data.Client.QueryAll(
		struct {
			Collection string
			Index      string
			Sort       bson.M
			Filter     bson.M
			Query      []string
			Limit      int64
			Skip       int64
		}{Collection: "[Asset~Address(Addresses)]", Index: "someIndex", Sort: bson.M{}, Filter: bson.M{"ChildID": r1["_id"]}, Query: []string{"ParentID"}, Limit: args.Limit, Skip: args.Skip}, ret)
	if err != nil {
		return err
	}
	// Step 3
	r3 := make([]map[string]interface{}, 0)
	for _, item := range r2 {
		r, err := me.Data.Client.QueryOne(struct {
			Collection string
			Index      string
			Sort       bson.M
			Filter     bson.M
			Query      []string
		}{Collection: "Asset", Index: "someIndex", Sort: bson.M{}, Filter: bson.M{"_id": item["ParentID"]}}, ret)
		if err != nil {
			return err
		}
		balance, err := me.GetBalanceByContractHashAddress(struct {
			ContractHash h160.T
			Address      h160.T
		}{ContractHash: h160.T(fmt.Sprint(r["hash"])), Address: args.Address}, ret)
		if err != nil {
			return err
		}
		r["balanceinfo"] = balance
		r3 = append(r3, r)
	}
	r, err := json.Marshal(r3)
	if err != nil {
		return err
	}
	*ret = json.RawMessage(r)
	return nil
}
