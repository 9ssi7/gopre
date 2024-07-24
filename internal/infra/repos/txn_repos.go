package repos

import (
	"github.com/9ssi7/txn"
	"github.com/9ssi7/txn/txngorm"
	"gorm.io/gorm"
)

type txnGormRepo struct {
	adapter txngorm.GAdapter
}

func (r *txnGormRepo) GetTxnAdapter() txn.Adapter {
	return r.adapter
}

func newTxnGormRepo(db *gorm.DB) txnGormRepo {
	return txnGormRepo{
		adapter: txngorm.New(db),
	}
}
