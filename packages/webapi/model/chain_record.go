package model

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/registry_pkg/chain_record"
)

type ChainRecord struct {
	ChainID             ChainID `swagger:"desc(ChainID (base58-encoded))"`
	Active              bool    `swagger:"desc(Whether or not the chain is active)"`
	DedicatedDbInstance bool    `swagger:"desc(Whether or not the chain data is stored in a separate db instance)"`
}

func NewChainRecord(rec *chain_record.ChainRecord) *ChainRecord {
	return &ChainRecord{
		ChainID:             NewChainID(coretypes.NewChainID(rec.ChainIdAliasAddress)),
		Active:              rec.Active,
		DedicatedDbInstance: rec.DedicatedDbInstance,
	}
}

func (bd *ChainRecord) Record() *chain_record.ChainRecord {
	return &chain_record.ChainRecord{
		ChainIdAliasAddress: bd.ChainID.ChainID().AliasAddress,
		Active:              bd.Active,
	}
}
