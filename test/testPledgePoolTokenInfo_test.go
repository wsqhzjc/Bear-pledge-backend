package test

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"pledge-backend/config"
	"pledge-backend/contract/bindings"
	"pledge-backend/log"
	serviceCommon "pledge-backend/schedule/common"
	"pledge-backend/utils"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var pledgePoolToken *bindings.PledgePoolToken
var ethereumConn *ethclient.Client

func setupTest(t *testing.T) {
	t.Helper()

	// Set up environment variables
	//deployer private key
	os.Setenv("plgr_admin_private_key", "70f16efc86d788ce1eeb8a000514569e2b9bf5a417c351ef341c97189d5a9507")
	//signer1 address
	// os.Setenv("plgr_admin_private_key", "db612dee6e7108560a08bee6686491cb354d04920ec4a157b104f16d83a2aab3")
	serviceCommon.GetEnv()

	var err error
	ethereumConn, err = ethclient.Dial(config.Config.TestNet.NetUrl)
	if nil != err {
		log.Logger.Error(err.Error())
		return
	}
	log.Logger.Info(fmt.Sprintf("pledgePoolToken: %v", config.Config.TestNet.PledgePoolToken))
	pledgePoolToken, err = bindings.NewPledgePoolToken(common.HexToAddress(config.Config.TestNet.PledgePoolToken), ethereumConn)
	if nil != err {
		log.Logger.Error(err.Error())
		return
	}
}

func TestGetSettleTimeByPoolID(t *testing.T) {
	setupTest(t)
	// You may want to mock models.NewPoolBases().PoolBaseInfo for a real unit test.
	poolId := 2

	isCurrentBlockTimeGreaterThanSettle, err := pledgePoolToken.PledgePoolTokenCaller.CheckoutSettle(nil, big.NewInt(int64(poolId)))

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	log.Logger.Info(fmt.Sprintf("isCurrentBlockTimeGreaterThanSettle: %v", isCurrentBlockTimeGreaterThanSettle))
}

func TestGetSettleTimeByPoolID_2(t *testing.T) {
	setupTest(t)
	poolId := 1

	out, err := pledgePoolToken.PledgePoolTokenCaller.PoolBaseInfo(nil, big.NewInt(int64(poolId)))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	log.Logger.Info(fmt.Sprintf("out: %v", out))
	settleTime := out.SettleTime

	// Get current block
	block, err := ethereumConn.BlockByNumber(context.Background(), nil)
	if err != nil {
		t.Fatalf("failed to get current block: %v", err)
	}
	currentBlockTime := block.Time()

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	log.Logger.Info(fmt.Sprintf("settleTime: %v", settleTime))
	log.Logger.Info(fmt.Sprintf("currentBlockTime: %v", currentBlockTime))

	if settleTime.Cmp(big.NewInt(int64(currentBlockTime))) > 0 {
		t.Fatalf("expected settleTime to be less than currentBlockTime")
	}
}

func TestCheckFinish(t *testing.T) {
	setupTest(t)
	poolId := 6

	out, err := pledgePoolToken.PledgePoolTokenCaller.CheckoutFinish(nil, big.NewInt(int64(poolId)))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	log.Logger.Info(fmt.Sprintf("out: %v", out))
}

func TestFinishPool(t *testing.T) {
	setupTest(t)
	poolId := 6

	// Create auth for transaction
	privateKeyEcdsa, err := crypto.HexToECDSA(serviceCommon.PlgrAdminPrivateKey)
	log.Logger.Info(fmt.Sprintf("privateKeyEcdsa: %v", serviceCommon.PlgrAdminPrivateKey))
	if err != nil {
		t.Fatalf("failed to create private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(utils.StringToInt64(config.Config.TestNet.ChainId)))
	if err != nil {
		t.Fatalf("failed to create auth: %v", err)
	}

	out, err := pledgePoolToken.PledgePoolTokenTransactor.Finish(auth, big.NewInt(int64(poolId)))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	log.Logger.Info(fmt.Sprintf("out: %v", out))
}

func TestGetMultiSignatureAddress(t *testing.T) {
	setupTest(t)
	multiSignatureAddress, err := pledgePoolToken.PledgePoolTokenCaller.GetMultiSignatureAddress(nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	log.Logger.Info(fmt.Sprintf("multiSignatureAddress: %v", multiSignatureAddress))
}

func TestPoolBaseInfo(t *testing.T) {
	setupTest(t)
	poolId := 6

	out, err := pledgePoolToken.PledgePoolTokenCaller.PoolBaseInfo(nil, big.NewInt(int64(poolId)))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	log.Logger.Info(fmt.Sprintf("out: %v", out))
}

func TestPoolDataInfo(t *testing.T) {

}
