package test

import (
	"fmt"
	"math/big"
	"os"
	"pledge-backend/config"
	"pledge-backend/contract/bindings"
	"pledge-backend/log"
	serviceCommon "pledge-backend/schedule/common"
	"pledge-backend/utils"
	"testing"
	"time"

	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var pancakeRouterToken *bindings.PancakeRouterToken

func setupSwapTest(t *testing.T) {
	t.Helper()

	// Set up environment variables
	//deployer private key
	os.Setenv("plgr_admin_private_key", "70f16efc86d788ce1eeb8a000514569e2b9bf5a417c351ef341c97189d5a9507") //0x344A6ba4e8aDa5A696444514273b9D428D6a041E
	//signer1 address //0x6517EA732f607BfeD2057a8E7f1B14E506E0007A
	// os.Setenv("plgr_admin_private_key", "db612dee6e7108560a08bee6686491cb354d04920ec4a157b104f16d83a2aab3")
	serviceCommon.GetEnv()

	var err error
	ethereumConn, err = ethclient.Dial(config.Config.TestNet.NetUrl)
	if nil != err {
		log.Logger.Error(err.Error())
		return
	}
	log.Logger.Info(fmt.Sprintf("pancakeRouterToken: %v", config.Config.TestNet.PancakeRouterToken))
	pancakeRouterToken, err = bindings.NewPancakeRouterToken(common.HexToAddress(config.Config.TestNet.PancakeRouterToken), ethereumConn)
	if nil != err {
		log.Logger.Error(err.Error())
		return
	}
}

func TestSwapBtcforBusd(t *testing.T) {
	setupSwapTest(t)

	// Create auth for transaction
	privateKeyEcdsa, err := crypto.HexToECDSA(serviceCommon.PlgrAdminPrivateKey)
	if err != nil {
		t.Fatalf("failed to create private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(utils.StringToInt64(config.Config.TestNet.ChainId)))
	if err != nil {
		t.Fatalf("failed to create auth: %v", err)
	}
	// Get token contract (BTC token)
	btcTokenAddress := common.HexToAddress("0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658") // BTC token
	btcToken, err := bindings.NewBindings(btcTokenAddress, ethereumConn)
	if err != nil {
		t.Fatalf("failed to create BTC token contract: %v", err)
	}

	// Get BUSD token contract
	busdTokenAddress := common.HexToAddress("0xE676Dcd74f44023b95E0E2C6436C97991A7497DA") // BUSD token
	busdToken, err := bindings.NewBindings(busdTokenAddress, ethereumConn)
	if err != nil {
		t.Fatalf("failed to create BUSD token contract: %v", err)
	}

	traderAddress := common.HexToAddress("0x344A6ba4e8aDa5A696444514273b9D428D6a041E")
	// routerAddress := common.HexToAddress(config.Config.TestNet.PancakeRouterToken)

	// Check token balance first
	balance, err := btcToken.BalanceOf(nil, traderAddress)
	if err != nil {
		t.Fatalf("failed to get balance: %v", err)
	}
	log.Logger.Info(fmt.Sprintf("Token balance: %v", balance))

	// Check if balance is sufficient
	amountIn := big.NewInt(1000000000000000)
	if balance.Cmp(amountIn) < 0 {
		t.Fatalf("insufficient balance: have %v, need %v", balance, amountIn)
	}

	//before swap
	balanceBefore, err := btcToken.BalanceOf(nil, traderAddress)
	if err != nil {
		t.Fatalf("failed to get balance: %v", err)
	}
	log.Logger.Info(fmt.Sprintf("BTC Token balance before swap: %v", balanceBefore))

	balanceBeforeBusd, err := busdToken.BalanceOf(nil, traderAddress)
	if err != nil {
		t.Fatalf("failed to get balance: %v", err)
	}
	log.Logger.Info(fmt.Sprintf("BUSD Token balance before swap: %v", balanceBeforeBusd))

	// Use SwapExactTokensForTokens from PancakeRouterTokenTransactor
	_, err = pancakeRouterToken.PancakeRouterTokenTransactor.SwapExactTokensForTokens(
		auth,
		amountIn,
		big.NewInt(0), // amountOutMin
		[]common.Address{
			btcTokenAddress,  // BTC token
			busdTokenAddress, // BUSD
		}, // path
		traderAddress, // to (your address)
		big.NewInt(time.Now().Add(20*time.Minute).Unix()), // deadline
	)
	if err != nil {
		t.Fatalf("failed to swap tokens: %v", err)
	}

	time.Sleep(10 * time.Second)

	//after swap
	balanceAfter, err := btcToken.BalanceOf(nil, traderAddress)
	if err != nil {
		t.Fatalf("failed to get balance: %v", err)
	}
	log.Logger.Info(fmt.Sprintf("BTC Token balance after swap: %v", balanceAfter))

	balanceAfterBusd, err := busdToken.BalanceOf(nil, traderAddress)
	if err != nil {
		t.Fatalf("failed to get balance: %v", err)
	}
	log.Logger.Info(fmt.Sprintf("BUSD Token balance after swap: %v", balanceAfterBusd))

}

func TestSwapBusdforBtc(t *testing.T) {
	setupSwapTest(t)

	// Create auth for transaction
	privateKeyEcdsa, err := crypto.HexToECDSA(serviceCommon.PlgrAdminPrivateKey)
	if err != nil {
		t.Fatalf("failed to create private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(utils.StringToInt64(config.Config.TestNet.ChainId)))
	if err != nil {
		t.Fatalf("failed to create auth: %v", err)
	}
	// Get token contract (BTC token)
	btcTokenAddress := common.HexToAddress("0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658") // BTC token
	btcToken, err := bindings.NewBindings(btcTokenAddress, ethereumConn)
	if err != nil {
		t.Fatalf("failed to create BTC token contract: %v", err)
	}

	// Get BUSD token contract
	busdTokenAddress := common.HexToAddress("0xE676Dcd74f44023b95E0E2C6436C97991A7497DA") // BUSD token
	busdToken, err := bindings.NewBindings(busdTokenAddress, ethereumConn)
	if err != nil {
		t.Fatalf("failed to create BUSD token contract: %v", err)
	}

	traderAddress := common.HexToAddress("0x344A6ba4e8aDa5A696444514273b9D428D6a041E")
	// routerAddress := common.HexToAddress(config.Config.TestNet.PancakeRouterToken)

	// Check token balance first
	busdBalance, err := busdToken.BalanceOf(nil, traderAddress)
	if err != nil {
		t.Fatalf("failed to get balance: %v", err)
	}
	log.Logger.Info(fmt.Sprintf("BUSD Token balance: %v", busdBalance))

	// Check if balance is sufficient
	amountIn := big.NewInt(1000000000000000000)
	if busdBalance.Cmp(amountIn) < 0 {
		t.Fatalf("insufficient balance: have %v, need %v", busdBalance, amountIn)
	}

	// Approve all BUSD balance for the router to spend
	routerAddress := common.HexToAddress(config.Config.TestNet.PancakeRouterToken)
	_, err = busdToken.Approve(auth, routerAddress, busdBalance)
	if err != nil {
		t.Fatalf("failed to approve BUSD token: %v", err)
	}

	// Use SwapExactTokensForTokens from PancakeRouterTokenTransactor
	_, err = pancakeRouterToken.PancakeRouterTokenTransactor.SwapExactTokensForTokens(
		auth,
		busdBalance,
		big.NewInt(0), // amountOutMin
		[]common.Address{
			busdTokenAddress, // BUSD
			btcTokenAddress,  // BTC token
		}, // path
		traderAddress, // to (your address)
		big.NewInt(time.Now().Add(20*time.Minute).Unix()), // deadline
	)
	if err != nil {
		t.Fatalf("failed to swap tokens: %v", err)
	}

	time.Sleep(10 * time.Second)

	//after swap
	balanceAfter, err := btcToken.BalanceOf(nil, traderAddress)
	if err != nil {
		t.Fatalf("failed to get balance: %v", err)
	}
	log.Logger.Info(fmt.Sprintf("BTC Token balance after swap: %v", balanceAfter))

	balanceAfterBusd, err := busdToken.BalanceOf(nil, traderAddress)
	if err != nil {
		t.Fatalf("failed to get balance: %v", err)
	}
	log.Logger.Info(fmt.Sprintf("BUSD Token balance after swap: %v", balanceAfterBusd))

}

func TestGetAmountsIn(t *testing.T) {
	setupSwapTest(t)

	btcTokenAddress := common.HexToAddress("0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658")
	busdTokenAddress := common.HexToAddress("0xE676Dcd74f44023b95E0E2C6436C97991A7497DA")

	amounts, err := pancakeRouterToken.GetAmountsIn(
		nil,
		new(big.Int).Mul(big.NewInt(1), big.NewInt(1000000000000000000)),
		[]common.Address{btcTokenAddress, busdTokenAddress},
	)

	if err != nil {
		t.Fatalf("failed to get amounts in: %v", err)
	}

	log.Logger.Info(fmt.Sprintf("amounts: %v", amounts))
}

func TestAddLiquidity(t *testing.T) {
	setupSwapTest(t)

	// Create auth for transaction
	//pledeg pool owner private key public key:  0xaDE481f68cb7fB5932de92Db1a214919CD09d89A
	privateKey := "ab6110cb1aca11bc7460843326af7d620138e69a9c5117700125e68bc8911bd5"
	privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	log.Logger.Info(fmt.Sprintf("privateKeyEcdsa: %v", privateKey))
	if err != nil {
		t.Fatalf("failed to create private key: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(utils.StringToInt64(config.Config.TestNet.ChainId)))
	if err != nil {
		t.Fatalf("failed to create auth: %v", err)
	}

	// Create PancakeSwap router instance
	router, err := bindings.NewPancakeRouterToken(common.HexToAddress(config.Config.TestNet.PancakeRouterToken), ethereumConn)
	if err != nil {
		t.Fatalf("failed to create router: %v", err)
	}

	btcTokenAddress := common.HexToAddress("0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658")  // BTC token
	busdTokenAddress := common.HexToAddress("0xE676Dcd74f44023b95E0E2C6436C97991A7497DA") // BUSD token

	// amountIn := new(big.Int).Mul(big.NewInt(1), big.NewInt(1000000000000000000)) // 1 token
	// amounts, err := pancakeRouterToken.GetAmountsOut(nil, amountIn, []common.Address{btcTokenAddress, busdTokenAddress})
	// if err != nil {
	// 	t.Fatal("")
	// }

	// // Convert big.Int to float64 for division
	// amountFloat := new(big.Float).SetInt(amounts[1])
	// amountFloat.Quo(amountFloat, big.NewFloat(1e18))
	// amountFloat64, _ := amountFloat.Float64()

	// log.Logger.Info(fmt.Sprintf("Price ratio: 1 BTC = %v BUSD", amountFloat64))

	// Create token instances
	btcToken, err := bindings.NewBindings(btcTokenAddress, ethereumConn)
	if err != nil {
		t.Fatalf("failed to create BTC token contract: %v", err)
	}
	busdToken, err := bindings.NewBindings(busdTokenAddress, ethereumConn)
	if err != nil {
		t.Fatalf("failed to create BUSD token contract: %v", err)
	}

	// Check balances
	btcBalance, err := btcToken.BalanceOf(nil, auth.From)
	if err != nil {
		t.Fatalf("failed to get BTC balance: %v", err)
	}
	busdBalance, err := busdToken.BalanceOf(nil, auth.From)
	if err != nil {
		t.Fatalf("failed to get BUSD balance: %v", err)
	}
	log.Logger.Info(fmt.Sprintf("BTC Balance: %v", btcBalance))
	log.Logger.Info(fmt.Sprintf("BUSD Balance: %v", busdBalance))

	// Approve tokens for router
	routerAddress := common.HexToAddress(config.Config.TestNet.PancakeRouterToken)
	_, err = btcToken.Approve(auth, routerAddress, new(big.Int).Mul(big.NewInt(1), big.NewInt(1000000000000000000)))
	if err != nil {
		t.Fatalf("failed to approve BTC token: %v", err)
	}
	_, err = busdToken.Approve(auth, routerAddress, new(big.Int).Mul(big.NewInt(100000), big.NewInt(1000000000000000000)))
	if err != nil {
		t.Fatalf("failed to approve BUSD token: %v", err)
	}

	// Add liquidity using PancakeSwap router
	_, err = router.PancakeRouterTokenTransactor.AddLiquidity(
		auth,
		btcTokenAddress,  // BTC token
		busdTokenAddress, // BUSD token
		new(big.Int).Mul(big.NewInt(1), big.NewInt(100000000000000000)),      // amountADesired
		new(big.Int).Mul(big.NewInt(100000), big.NewInt(100000000000000000)), // amountBDesired
		big.NewInt(0), // amountAMin
		big.NewInt(0), // amountBMin
		// common.HexToAddress("0xaDE481f68cb7fB5932de92Db1a214919CD09d89A"), // to
		auth.From,
		big.NewInt(time.Now().Add(20*time.Minute).Unix()), // deadline
	)
	if err != nil {
		t.Fatalf("failed to add liquidity: %v", err)
	}
}

func TestGetWETH(t *testing.T) {
	setupSwapTest(t)
	router, err := bindings.NewPancakeRouterToken(common.HexToAddress(config.Config.TestNet.PancakeRouterToken), ethereumConn)
	if err != nil {
		t.Fatalf("failed to create router: %v", err)
	}

	wethAddress, err := router.WETH(nil)
	if err != nil {
		t.Fatalf("failed to get WETH address: %v", err)
	}
	log.Logger.Info(fmt.Sprintf("WETH address: %v", wethAddress))

}

func TestGetInitCodeHash(t *testing.T) {
	setupSwapTest(t)

	factory, err := bindings.NewUniswapV2FactoryToken(common.HexToAddress(config.Config.TestNet.PancakeFactoryToken), ethereumConn)
	if err != nil {
		t.Fatalf("failed to create factory: %v", err)
	}

	initCodeHash, err := factory.InitCodeHash(nil)
	if err != nil {
		t.Fatalf("failed to get init code hash: %v", err)
	}
	log.Logger.Info(fmt.Sprintf("init code hash: %v", initCodeHash))

	initCodeHashHex := hex.EncodeToString(initCodeHash[:])
	log.Logger.Info(fmt.Sprintf("init code hash hex: 0x%v", initCodeHashHex))
}

func TestRemoveLiquidity(t *testing.T) {
	setupSwapTest(t)

	// // Create auth for transaction
	// //pledeg pool owner private key public key:  0xaDE481f68cb7fB5932de92Db1a214919CD09d89A
	// privateKey := "ab6110cb1aca11bc7460843326af7d620138e69a9c5117700125e68bc8911bd5"
	// privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)T
	// log.Logger.Info(fmt.Sprintf("privateKeyEcdsa: %v", privateKey))
	// if err != nil {
	// 	t.Fatalf("failed to create private key: %v", err)
	// }

	// auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(utils.StringToInt64(config.Config.TestNet.ChainId)))
	// if err != nil {
	// 	t.Fatalf("failed to create auth: %v", err)
	// }

	// // Create PancakeSwap router instance
	// router, err := bindings.NewPancakeRouterToken(common.HexToAddress(config.Config.TestNet.PancakeRouterToken), ethereumConn)
	// if err != nil {
	// 	t.Fatalf("failed to create router: %v", err)
	// }

	// btcTokenAddress := common.HexToAddress("0xB5514a4FA9dDBb48C3DE215Bc9e52d9fCe2D8658")  // BTC token
	// busdTokenAddress := common.HexToAddress("0xE676Dcd74f44023b95E0E2C6436C97991A7497DA") // BUSD token

	// amountIn := new(big.Int).Mul(big.NewInt(1), big.NewInt(1000000000000000000)) // 1 token
	// amounts, err := pancakeRouterToken.GetAmountsOut(nil, amountIn, []common.Address{btcTokenAddress, busdTokenAddress})
	// if err != nil {
	// 	t.Fatal("")
	// }

	// log.Logger.Info(fmt.Sprintf("amounts: %v", amounts))
}
