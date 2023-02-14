package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sunjiangjun/xlog"
	"github.com/tidwall/gjson"
	"github.com/uduncloud/easynode_taskapi/config"
	"github.com/uduncloud/easynode_taskapi/service/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"io/ioutil"
	"time"
)

type Server struct {
	db         *gorm.DB
	log        *xlog.XLog
	blockChain []int64
}

func NewServer(dbConfig *config.TaskDb, blockChain []int64, log *xlog.XLog) *Server {
	g, err := db.Open(dbConfig.User, dbConfig.Password, dbConfig.Addr, dbConfig.DbName, dbConfig.Port, log)
	if err != nil {
		panic(err)
	}

	return &Server{
		log:        log,
		db:         g,
		blockChain: blockChain,
	}
}

func (s *Server) GetActiveNodes(c *gin.Context) {
	list, err := s.GetActiveNodesFromDB()
	if err != nil {
		s.Error(c, c.Request.URL.Path, err.Error())
		return
	}
	s.Success(c, list, c.Request.URL.Path)
}

func (s *Server) PushBlockTask(c *gin.Context) {

	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		s.Error(c, c.Request.URL.Path, err.Error())
		return
	}

	r := gjson.ParseBytes(bs)
	blockChain := r.Get("blockChain").Int()
	has := false
	for _, c := range s.blockChain {
		if c == blockChain {
			has = true
			break
		}
	}
	if !has {
		s.Error(c, c.Request.URL.Path, errors.New("blockchain is wrong").Error())
		return
	}

	blockHash := r.Get("blockHash").String()
	blockNumber := r.Get("blockNumber").String()

	if len(blockHash) < 10 && len(blockNumber) < 2 {
		s.Error(c, c.Request.URL.Path, errors.New("blockHash or blockNumber is wrong").Error())
		return
	}

	nodeSource := &NodeSource{BlockChain: blockChain, BlockHash: blockHash, BlockNumber: blockNumber, SourceType: 2}

	err = s.AddNodeSource(nodeSource)
	if err != nil {
		s.Error(c, c.Request.URL.Path, err.Error())
		return
	}

	s.Success(c, nil, c.Request.URL.Path)
}

func (s *Server) PushTxTask(c *gin.Context) {

	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		s.Error(c, c.Request.URL.Path, err.Error())
		return
	}

	r := gjson.ParseBytes(bs)
	blockChain := r.Get("blockChain").Int()
	has := false
	for _, c := range s.blockChain {
		if c == blockChain {
			has = true
			break
		}
	}
	if !has {
		s.Error(c, c.Request.URL.Path, errors.New("blockchain is wrong").Error())
		return
	}

	txHash := r.Get("txHash").String()

	if len(txHash) < 10 {
		s.Error(c, c.Request.URL.Path, errors.New("txHash is wrong").Error())
		return
	}

	nodeSource := &NodeSource{BlockChain: blockChain, TxHash: txHash, SourceType: 1}

	err = s.AddNodeSource(nodeSource)
	if err != nil {
		s.Error(c, c.Request.URL.Path, err.Error())
		return
	}

	s.Success(c, nil, c.Request.URL.Path)
}

func (s *Server) PushTxsTask(c *gin.Context) {

	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		s.Error(c, c.Request.URL.Path, err.Error())
		return
	}

	r := gjson.ParseBytes(bs)
	blockChain := r.Get("blockChain").Int()
	has := false
	for _, c := range s.blockChain {
		if c == blockChain {
			has = true
			break
		}
	}
	if !has {
		s.Error(c, c.Request.URL.Path, errors.New("blockchain is wrong").Error())
		return
	}

	blockHash := r.Get("blockHash").String()
	blockNumber := r.Get("blockNumber").String()

	if len(blockHash) < 10 && len(blockNumber) < 2 {
		s.Error(c, c.Request.URL.Path, errors.New("blockHash or blockNumber is wrong").Error())
		return
	}

	nodeSource := &NodeSource{BlockChain: blockChain, BlockHash: blockHash, BlockNumber: blockNumber, SourceType: 1}

	err = s.AddNodeSource(nodeSource)
	if err != nil {
		s.Error(c, c.Request.URL.Path, err.Error())
		return
	}

	s.Success(c, nil, c.Request.URL.Path)
}

func (s *Server) PushReceiptTask(c *gin.Context) {

	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		s.Error(c, c.Request.URL.Path, err.Error())
		return
	}

	r := gjson.ParseBytes(bs)
	blockChain := r.Get("blockChain").Int()
	has := false
	for _, c := range s.blockChain {
		if c == blockChain {
			has = true
			break
		}
	}
	if !has {
		s.Error(c, c.Request.URL.Path, errors.New("blockchain is wrong").Error())
		return
	}

	txHash := r.Get("txHash").String()

	if len(txHash) < 10 {
		s.Error(c, c.Request.URL.Path, errors.New("txHash  is wrong").Error())
		return
	}

	nodeSource := &NodeSource{BlockChain: blockChain, TxHash: txHash, SourceType: 3}

	err = s.AddNodeSource(nodeSource)
	if err != nil {
		s.Error(c, c.Request.URL.Path, err.Error())
		return
	}

	s.Success(c, nil, c.Request.URL.Path)
}

func (s *Server) PushReceiptsTask(c *gin.Context) {

	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		s.Error(c, c.Request.URL.Path, err.Error())
		return
	}

	r := gjson.ParseBytes(bs)
	blockChain := r.Get("blockChain").Int()
	has := false
	for _, c := range s.blockChain {
		if c == blockChain {
			has = true
			break
		}
	}
	if !has {
		s.Error(c, c.Request.URL.Path, errors.New("blockchain is wrong").Error())
		return
	}

	blockHash := r.Get("blockHash").String()
	blockNumber := r.Get("blockNumber").String()

	if len(blockHash) < 10 && len(blockNumber) < 2 {
		s.Error(c, c.Request.URL.Path, errors.New("blockHash or blockNumber is wrong").Error())
		return
	}

	nodeSource := &NodeSource{BlockChain: blockChain, BlockHash: blockHash, BlockNumber: blockNumber, SourceType: 3}

	err = s.AddNodeSource(nodeSource)
	if err != nil {
		s.Error(c, c.Request.URL.Path, err.Error())
		return
	}

	s.Success(c, nil, c.Request.URL.Path)
}

func (s *Server) AddNodeSourceList(sources []*NodeSource) error {
	return s.db.Table(NodeSourceTable).Omit("create_time,id").Clauses(clause.Insert{Modifier: "IGNORE"}).CreateInBatches(sources, 10).Error
}

func (s *Server) AddNodeSource(source *NodeSource) error {
	return s.db.Table(NodeSourceTable).Omit("create_time,id").Clauses(clause.Insert{Modifier: "IGNORE"}).Create(source).Error
}

func (s *Server) GetActiveNodesFromDB() ([]string, error) {
	t := time.Now().Add(-5 * time.Minute).Format(TimeFormat)
	var nodeList []string
	err := s.db.Table(NodeInfoTable).Select("node_id").Where("create_time>?", t).Pluck("node_id", &nodeList).Error

	if err != nil {
		return nil, err
	}

	return nodeList, nil
}

const (
	SUCCESS = 0
	FAIL    = 1
)

func (s *Server) Success(c *gin.Context, resp interface{}, path string) {
	s.log.Printf("path=%v,body=%v\n", path, resp)
	mp := make(map[string]interface{})
	mp["code"] = SUCCESS
	mp["data"] = resp
	c.JSON(200, mp)
}

func (s *Server) Error(c *gin.Context, path string, err string) {
	s.log.Printf("path=%v,err=%v\n", path, err)
	mp := make(map[string]interface{})
	mp["code"] = FAIL
	mp["data"] = err
	c.JSON(200, mp)
}