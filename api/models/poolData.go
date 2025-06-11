package models

import (
	"pledge-backend/db"
)

type PoolData struct {
	PoolID                 int    `json:"-"`
	SettleAmountLend       string `json:"settleAmountLend"`
	SettleAmountBorrow     string `json:"settleAmountBorrow"`
	FinishAmountLend       string `json:"finishAmountLend"`
	FinishAmountBorrow     string `json:"finishAmountBorrow"`
	LiquidationAmounLend   string `json:"liquidationAmounLend"`
	LiquidationAmounBorrow string `json:"liquidationAmounBorrow"`
	ChainID                int    `json:"chain_id"`
}

type PoolDataInfoRes struct {
	Index    int      `json:"index"`
	PoolData PoolData `json:"pool_data"`
}

func NewPoolData() *PoolData {
	return &PoolData{}
}

func (p *PoolData) PoolDataInfo(chainId int, res *[]PoolDataInfoRes) error {
	var poolData []PoolData

	err := db.Mysql.Table("pooldata").Where("chain_id=?", chainId).Order("CAST(pool_id AS SIGNED) ASC").Find(&poolData).Debug().Error
	if err != nil {
		return err
	}

	for _, v := range poolData {
		*res = append(*res, PoolDataInfoRes{
			Index:    v.PoolID - 1,
			PoolData: v,
		})
	}
	return nil
}
