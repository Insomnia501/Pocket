package service

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"os"
	"pocket-serv/global"
	"strings"

	log "github.com/cihub/seelog"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	err                    error
	client                 *ethclient.Client
	erc6551RegistryAddress common.Address
	vipcardAccountAddress  common.Address
	fromAddress            common.Address
	privateKey             *ecdsa.PrivateKey
	chainId                *big.Int
)

func DeployERC721Contract(name string, address string, descripe string) (string, error) {
	// address是用户钱包地址，即部署这个721合约的地址。目前先实现了由admin账户进行部署，再将owner设为用户钱包地址的部署方式
	//1.准备工作
	fromAddress = common.HexToAddress(address)
	erc6551RegistryAddress = common.HexToAddress(global.App.Config.Contract.Erc6551RegistryAddress)
	vipcardAccountAddress = common.HexToAddress(global.App.Config.Contract.VipcardAccountAddress)
	client, err := ethclient.Dial(global.App.Config.Contract.AlchemyRpcUrl)
	if err != nil {
		log.Debugf("Failed to create eth client: %v", err)
	}
	chainId, err = client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	//2.读取合约的ABI和字节码
	contractJSON, err := os.ReadFile(global.App.Config.Contract.VipcardContractAbi)
	if err != nil {
		log.Debugf("Failed to read ERC721 contract file: %v", contractJSON)
	}

	var contractData struct {
		ABI      string
		Bytecode string
	}
	//`json:"abi"``json:"bytecode"`

	err = json.Unmarshal(contractJSON, &contractData)
	log.Debugf("ABI: %v", contractData.ABI)
	if err != nil {
		log.Debugf("Failed to unmarshal ERC721 contract JSON: %v", err)
	}

	//3.构建部署合约的ABI
	//contractAbi, err := abi.JSON(bytes.NewReader(abiData))
	parsedABI, err := abi.JSON(strings.NewReader(contractData.ABI))
	if err != nil {
		log.Debugf("Failed to parse ABI: %v", err)
	}

	//4.创建一个新的交易，然后使用您的私钥对交易进行签名
	privateKey, err := crypto.HexToECDSA(global.App.Config.Contract.PrivateKey)
	if err != nil {
		log.Debug(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Debugf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Debug(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Debugf("Failed to create authorized transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	contractAddress, tx, _, err := bind.DeployContract(
		auth,
		parsedABI,
		common.FromHex(contractData.Bytecode),
		client,
		name,
		global.App.Config.Contract.VipcardSymbol,
		global.App.Config.Contract.Erc6551RegistryAddress,
		global.App.Config.Contract.VipcardAccountAddress,
		fromAddress)
	if err != nil {
		log.Debug(err)
	}

	_ = tx // handle or log the transaction as per your requirement

	return contractAddress.Hex(), nil
}
