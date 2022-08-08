package psql

import (
	"github.com/dojimanetwork/dojimamint/state/indexer"
	"github.com/dojimanetwork/dojimamint/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
