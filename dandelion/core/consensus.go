package core

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/iservices"
	"github.com/coschain/contentos-go/node"
	"github.com/coschain/contentos-go/prototype"
)

const DummyConsensusName = "dandelion.DummyConsensus"
const BeforePreShuffleEvent = "dandelion.DummyConsensus.Event.BeforePreShuffle"
const AfterPreShuffleEvent = "dandelion.DummyConsensus.Event.AfterPreShuffle"

type DummyConsensus struct {
	ctx *node.ServiceContext
	evBus EventBus.Bus
	trxPool iservices.ITrxPool
	producers []string
}

func NewDummyConsensus(ctx *node.ServiceContext) (*DummyConsensus, error) {
	return &DummyConsensus{ ctx: ctx }, nil
}

func (c *DummyConsensus) Start(node *node.Node) error {
	if trxpool, err := c.ctx.Service(iservices.TxPoolServerName); err != nil {
		return fmt.Errorf("cannot get TrxPool service: %s", err.Error())
	} else {
		c.trxPool = trxpool.(iservices.ITrxPool)
	}
	c.evBus = node.EvBus
	c.trxPool.SetShuffle(c.shuffle)
	c.restoreProducers()
	return nil
}

func (c *DummyConsensus) Stop() error {
	return nil
}

func (c *DummyConsensus) shuffle(head common.ISignedBlock) (bool, []string) {
	blockNum := head.Id().BlockNum()
	if blockNum%constants.BlockProdRepetition != 0 ||
		blockNum/constants.BlockProdRepetition%uint64(len(c.producers)) != 0 {
		return false, []string{}
	}

	c.evBus.Publish(BeforePreShuffleEvent)
	_ = c.trxPool.PreShuffle()
	c.evBus.Publish(AfterPreShuffleEvent)

	prods, pubKeys := c.trxPool.GetBlockProducerTopN(constants.MaxBlockProducerCount)

	var seed uint64
	if head != nil {
		seed = head.Timestamp() << 32
	}
	c.updateProducers(seed, prods, pubKeys)

	return true, prods
}

func (c *DummyConsensus) updateProducers(seed uint64, prods []string, pubKeys []*prototype.PublicKeyType) int {
	prodNum := len(prods)
	for i := 0; i < prodNum; i++ {
		k := seed + uint64(i)*2695921657736338717
		k ^= k >> 12
		k ^= k << 25
		k ^= k >> 27
		k *= 2695921657736338717

		j := i + int(k%uint64(prodNum-i))
		prods[i], prods[j] = prods[j], prods[i]
		pubKeys[i], pubKeys[j] = pubKeys[j], pubKeys[i]
	}
	c.trxPool.SetShuffledBpList(prods, pubKeys, 0)
	c.producers = prods
	return prodNum
}

func (c *DummyConsensus) restoreProducers() {
	prods, _, _ := c.trxPool.GetShuffledBpList()
	c.producers = prods
}

func (c *DummyConsensus) GetProducer(num uint64) string {
	return c.producers[num/constants.BlockProdRepetition%uint64(len(c.producers))]
}
