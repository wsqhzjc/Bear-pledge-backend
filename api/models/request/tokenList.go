package request

type TokenList struct {
	ChainId int `form:"chainId" json:"chain_id" binding:"required"`
}
