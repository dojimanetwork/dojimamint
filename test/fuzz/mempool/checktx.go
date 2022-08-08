package checktx

import (
	"github.com/dojimanetwork/dojimamint/abci/example/kvstore"
	"github.com/dojimanetwork/dojimamint/config"
	mempl "github.com/dojimanetwork/dojimamint/mempool"
	"github.com/dojimanetwork/dojimamint/proxy"
)

var mempool mempl.Mempool

func init() {
	app := kvstore.NewApplication()
	cc := proxy.NewLocalClientCreator(app)
	appConnMem, _ := cc.NewABCIClient()
	err := appConnMem.Start()
	if err != nil {
		panic(err)
	}

	cfg := config.DefaultMempoolConfig()
	cfg.Broadcast = false

	mempool = mempl.NewCListMempool(cfg, appConnMem, 0)
}

func Fuzz(data []byte) int {
	err := mempool.CheckTx(data, nil, mempl.TxInfo{})
	if err != nil {
		return 0
	}

	return 1
}
