package service

import "time"

const (
	NodeInfoTable   = "node_info"
	NodeSourceTable = "node_source"
	NodeTaskTable   = "node_task_%v"
)

const (
	TimeFormat = "2006-01-02 15:04:05"
)

type NodeSource struct {
	Id          int64     `json:"id" gorm:"column:id"`
	BlockChain  int64     `json:"blockChain" gorm:"column:block_chain"` // 公链code
	TxHash      string    `json:"txHash" gorm:"column:tx_hash"`
	BlockHash   string    `json:"blockHash" gorm:"column:block_hash"`
	BlockNumber string    `json:"blockNumber" gorm:"column:block_number"`
	SourceType  int       `json:"sourceType" gorm:"column:source_type"` // 任务类型 1: 交易 2:区块 3.收据
	CreateTime  time.Time `json:"createTime" gorm:"column:create_time"`
}
