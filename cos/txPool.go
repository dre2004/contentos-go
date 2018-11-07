package cos

import (
	"github.com/coschain/contentos-go/common/prototype"
	"github.com/coschain/contentos-go/p2p/depend/common"
	"github.com/coschain/contentos-go/p2p/depend/event"
	"sync"
)

type TxPoolConfig struct {
}

type TxPool struct {
	mu           sync.RWMutex
	scope        event.SubscriptionScope
	txFeed       event.Feed
}

// NewTxsEvent is posted when a batch of transactions enter the transaction pool.
type NewTxsEvent struct{
	Txs []*prototype.Transaction
}

func NewTxPool() *TxPool {
	return nil
}

// AddRemotes enqueues a batch of transactions into the pool if they are valid.
// If the senders are not among the locally tracked ones, full pricing constraints
// will apply.
func (pool *TxPool) AddRemotes(txs []*prototype.Transaction) []error {
	return pool.addTxs(txs, false)
}

// addTxs attempts to queue a batch of transactions if they are valid.
func (pool *TxPool) addTxs(txs []*prototype.Transaction, local bool) []error {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	return pool.addTxsLocked(txs, local)
}

// addTxsLocked attempts to queue a batch of transactions if they are valid,
// whilst assuming the transaction pool lock is already held.
func (pool *TxPool) addTxsLocked(txs []*prototype.Transaction, local bool) []error {
	errs := make([]error, len(txs))
	return errs
}

// Pending retrieves all currently processable transactions, grouped by origin
// account and sorted by nonce. The returned transaction set is a copy and can be
// freely modified by calling code.
func (pool *TxPool) Pending() (map[common.Address][]*prototype.Transaction, error) {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	pending := make(map[common.Address][]*prototype.Transaction)
	return pending, nil
}

// SubscribeNewTxsEvent registers a subscription of NewTxsEvent and
// starts sending event to the given channel.
func (pool *TxPool) SubscribeNewTxsEvent(ch chan<- NewTxsEvent) event.Subscription {
	return pool.scope.Track(pool.txFeed.Subscribe(ch))
}