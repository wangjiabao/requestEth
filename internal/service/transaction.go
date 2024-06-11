package service

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	sdk "github.com/BioforestChain/go-bfmeta-wallet-sdk"
	"github.com/BioforestChain/go-bfmeta-wallet-sdk/entity/req/address"
	"github.com/BioforestChain/go-bfmeta-wallet-sdk/entity/req/broadcastTra"
	"github.com/BioforestChain/go-bfmeta-wallet-sdk/entity/req/createTransferAsset"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"

	pb "requestEth/api/requestEth/v1"
)

type TransactionService struct {
	pb.UnimplementedTransactionServer
}

func NewTransactionService() *TransactionService {
	return &TransactionService{}
}

func (s *TransactionService) SendTransaction(ctx context.Context, req *pb.SendTransactionRequest) (*pb.SendTransactionReply, error) {

	var (
		err          error
		tx           string
		tokenAddress = "0x55d398326f99059fF775485246999027B3197955"
		tmpUrl1      = "https://bsc-dataseed4.binance.org/"
		//tokenAddress = "0x337610d27c682E347C9cD60BD4b3b107C9d34dDd"
		//tmpUrl1      = "https://data-seed-prebsc-1-s3.binance.org:8545/"
	)

	if 0 == len(req.SendBody.ToAddr) {
		return &pb.SendTransactionReply{
			Tx: tx,
		}, nil
	}

	for i := 0; i <= 5; i++ {
		tx, err = toToken(req.SendBody.PrivateKey, req.SendBody.ToAddr, req.SendBody.Amount, tokenAddress, tmpUrl1)
		if err == nil {
			break
		} else if "insufficient funds for gas * price + value" == err.Error() {
			return &pb.SendTransactionReply{
				Tx:  "",
				Err: err.Error(),
			}, nil
		} else {
			fmt.Println(err)
			if 0 == i {
				tmpUrl1 = "https://bsc-dataseed1.binance.org"
			} else if 1 == i {
				tmpUrl1 = "https://bsc-dataseed3.binance.org"
			} else if 2 == i {
				tmpUrl1 = "https://bsc-dataseed2.binance.org"
			} else if 3 == i {
				tmpUrl1 = "https://bnb-bscnews.rpc.blxrbdn.com"
			} else if 4 == i {
				tmpUrl1 = "https://bsc-dataseed.binance.org"
			}
			fmt.Println("err info", req.SendBody.PrivateKey, req.SendBody.ToAddr, req.SendBody.Amount, tokenAddress, tmpUrl1)
			time.Sleep(3 * time.Second)
		}
	}

	if "" != tx {
		return &pb.SendTransactionReply{
			Tx: tx,
		}, nil
	}

	return &pb.SendTransactionReply{
		Tx:  tx,
		Err: err.Error(),
	}, nil
}

func toToken(userPrivateKey string, toAccount string, withdrawAmount string, withdrawTokenAddress string, url1 string) (string, error) {
	client, err := ethclient.Dial(url1)
	//client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return "", err
	}

	tokenAddress := common.HexToAddress(withdrawTokenAddress)
	instance, err := NewUsdt(tokenAddress, client)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var authUser *bind.TransactOpts

	var privateKey *ecdsa.PrivateKey
	privateKey, err = crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	//gasPrice, err := client.SuggestGasPrice(context.Background())
	//if err != nil {
	//	fmt.Println(err)
	//	return "", err
	//}

	authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	tmpWithdrawAmount, _ := new(big.Int).SetString(withdrawAmount, 10)
	var tx *types.Transaction
	tx, err = instance.Transfer(&bind.TransactOpts{
		From:     authUser.From,
		Signer:   authUser.Signer,
		GasLimit: 0,
	}, common.HexToAddress(toAccount), tmpWithdrawAmount)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *TransactionService) SendTransactionEth(ctx context.Context, req *pb.SendTransactionEthRequest) (*pb.SendTransactionEthReply, error) {

	var (
		err     error
		tx      string
		tmpUrl1 = "https://bsc-dataseed4.binance.org/"
		//tokenAddress = "0x337610d27c682E347C9cD60BD4b3b107C9d34dDd"
		//tmpUrl1      = "https://data-seed-prebsc-1-s3.binance.org:8545/"
	)

	if 0 == len(req.SendBody.ToAddr) {
		return &pb.SendTransactionEthReply{
			Tx: tx,
		}, nil
	}

	for i := 0; i <= 5; i++ {
		_, tx, err = toBnB(req.SendBody.ToAddr, req.SendBody.PrivateKey, req.SendBody.Amount, tmpUrl1)
		if err == nil {
			break
		} else if "insufficient funds for gas * price + value" == err.Error() {
			return &pb.SendTransactionEthReply{
				Tx:  "",
				Err: err.Error(),
			}, nil
		} else {
			fmt.Println(err)
			if 0 == i {
				tmpUrl1 = "https://bsc-dataseed1.binance.org"
			} else if 1 == i {
				tmpUrl1 = "https://bsc-dataseed3.binance.org"
			} else if 2 == i {
				tmpUrl1 = "https://bsc-dataseed2.binance.org"
			} else if 3 == i {
				tmpUrl1 = "https://bnb-bscnews.rpc.blxrbdn.com"
			} else if 4 == i {
				tmpUrl1 = "https://bsc-dataseed.binance.org"
			}
			fmt.Println("err info eth", req.SendBody.ToAddr, req.SendBody.PrivateKey, req.SendBody.Amount, tmpUrl1)
			time.Sleep(3 * time.Second)
		}
	}

	if "" != tx {
		return &pb.SendTransactionEthReply{
			Tx: tx,
		}, nil
	}

	return &pb.SendTransactionEthReply{
		Tx:  tx,
		Err: err.Error(),
	}, nil
}

func toBnB(toAccount string, fromPrivateKey string, toAmount string, url1 string) (bool, string, error) {
	//client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	client, err := ethclient.Dial(url1)
	if err != nil {
		return false, "", err
	}

	privateKey, err := crypto.HexToECDSA(fromPrivateKey)
	if err != nil {
		return false, "", err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return false, "", err
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return false, "", err
	}

	value := big.NewInt(0) // in wei (1 eth) 最低0.03bnb才能转账
	value.SetString(toAmount, 10)
	gasLimit := uint64(210000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return false, "", err
	}
	toAddress := common.HexToAddress(toAccount)
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		return false, "", err
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return false, "", err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return false, "", err
	}
	return true, signedTx.Hash().Hex(), nil
}

func (s *TransactionService) EthBalance(ctx context.Context, req *pb.EthBalanceRequest) (*pb.EthBalanceReply, error) {
	//client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	client, err := ethclient.Dial("https://bsc-dataseed4.binance.org/")
	if err != nil {
		return &pb.EthBalanceReply{
			Balance: "",
			Err:     err.Error(),
		}, nil
	}

	account := common.HexToAddress(req.Address)
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return &pb.EthBalanceReply{
			Balance: "",
			Err:     err.Error(),
		}, nil
	}

	return &pb.EthBalanceReply{
		Balance: balance.String(),
		Err:     "",
	}, nil
}

func (s *TransactionService) UsdtBalance(ctx context.Context, req *pb.UsdtBalanceRequest) (*pb.UsdtBalanceReply, error) {
	//client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return &pb.UsdtBalanceReply{
			Balance: "",
			Err:     err.Error(),
		}, nil
	}

	tokenAddress := common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")
	instance, err := NewUsdt(tokenAddress, client)
	if err != nil {
		return &pb.UsdtBalanceReply{
			Balance: "",
			Err:     err.Error(),
		}, nil
	}

	address := common.HexToAddress(req.Address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return &pb.UsdtBalanceReply{
			Balance: "",
			Err:     err.Error(),
		}, nil
	}

	return &pb.UsdtBalanceReply{
		Balance: bal.String(),
		Err:     "",
	}, nil
}

func (s *TransactionService) UsdtBalanceBiw(ctx context.Context, req *pb.UsdtBalanceBiwRequest) (*pb.UsdtBalanceBiwReply, error) {
	sdkClient := sdk.NewBCFWalletSDK()
	wallet := sdkClient.NewBCFWallet("35.213.66.234", 30003, "https://tracker.biw-meta.info/browser")
	p := address.Params{
		req.Address,
		"JWWWB",
		"BIW",
	}
	balance := wallet.GetAddressBalance(p)
	defer sdkClient.Close()

	return &pb.UsdtBalanceBiwReply{
		Balance: balance.Result.Amount,
	}, nil
	//client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	//client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	//if err != nil {
	//	return &pb.UsdtBalanceBiwReply{
	//		Balance: "",
	//		Err:     err.Error(),
	//	}, nil
	//}
	//
	//tokenAddress := common.HexToAddress("0x55d398326f99059fF775485246999027B3197955")
	//instance, err := NewUsdt(tokenAddress, client)
	//if err != nil {
	//	return &pb.UsdtBalanceBiwReply{
	//		Balance: "",
	//		Err:     err.Error(),
	//	}, nil
	//}
	//
	//address := common.HexToAddress(req.Address)
	//bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	//if err != nil {
	//	return &pb.UsdtBalanceBiwReply{
	//		Balance: "",
	//		Err:     err.Error(),
	//	}, nil
	//}
	//
	//return &pb.UsdtBalanceBiwReply{
	//	Balance: bal.String(),
	//	Err:     "",
	//}, nil
}

func (s *TransactionService) SendTransactionBiw(ctx context.Context, req *pb.SendTransactionBiwRequest) (*pb.SendTransactionBiwReply, error) {
	sdkClient := sdk.NewBCFWalletSDK()
	bCFSignUtil := sdkClient.NewBCFSignUtil("b")
	wallet := sdkClient.NewBCFWallet("35.213.66.234", 30003, "https://tracker.biw-meta.info/browser")
	bCFSignUtilCreateKeypair, _ := bCFSignUtil.CreateKeypair("")

	reqCreateTransferAsset := createTransferAsset.TransferAssetTransactionParams{
		TransactionCommonParamsWithRecipientId: createTransferAsset.TransactionCommonParamsWithRecipientId{
			TransactionCommonParams: createTransferAsset.TransactionCommonParams{
				PublicKey:        bCFSignUtilCreateKeypair.PublicKey,
				Fee:              "1000",
				ApplyBlockHeight: wallet.GetLastBlock().Result.Height,
				//Remark: map[string]string{
				//	"note": "example transaction",
				//},
				//BinaryInfos: []createTransferAsset.KVStorageInfo{
				//	{
				//		Key: "exampleKey",
				//		FileInfo: createTransferAsset.FileInfo{
				//			Name: "exampleFile",
				//			Size: 1234,
				//		},
				//	},
				//},
				//Timestamp: 1622732931,
			},
			RecipientId: req.SendBody.ToAddr, //钱包地址
		},
		//SourceChainMagic: "exampleSourceChainMagic",
		//SourceChainName:  "exampleSourceChainName",
		//AssetType:        "exampleAssetType",
		Amount: req.SendBody.Amount,
	}
	createTransferAssetResp, _ := wallet.CreateTransferAsset(reqCreateTransferAsset)

	//// 3.3 生成签名
	var s1 = []byte(createTransferAssetResp.Result.Buffer)
	var ss = []byte(bCFSignUtilCreateKeypair.SecretKey)
	detachedSign, _ := bCFSignUtil.DetachedSign(s1, ss)

	//// 3.4 wallet.BroadcastTransferAsset()
	req1 := broadcastTra.BroadcastTransactionParams{
		Signature: hex.EncodeToString(detachedSign.Data),
		//SignSignature: "exampleSignSignature", //非必传
		Buffer:    createTransferAssetResp.Result.Buffer, //3.2 上面取得的buffer
		IsOnChain: true,
	}
	success, _ := wallet.BroadcastTransferAsset(req1)

	return &pb.SendTransactionBiwReply{
		Tx:      hex.EncodeToString(detachedSign.Data),
		Success: success.Success,
	}, nil
}

func (s *TransactionService) Transaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionReply, error) {
	// https://data-seed-prebsc-1-s3.binance.org:8545/
	// https://bsc-dataseed.binance.org/
	client, err := ethclient.Dial("https://bsc-dataseed4.binance.org/")
	if err != nil {
		return &pb.TransactionReply{
			Status: 99,
		}, nil
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(req.GetTx()))
	if err != nil {
		return &pb.TransactionReply{
			Status: 99,
		}, nil
	}

	return &pb.TransactionReply{
		Status: receipt.Status,
	}, err
}

func (s *TransactionService) GenerateKey(ctx context.Context, req *pb.GenerateKeyRequest) (*pb.GenerateKeyReply, error) {
	address, privateKey, err := generateKey()
	if nil != err {
		return &pb.GenerateKeyReply{
			Address:    address,
			PrivateKey: privateKey,
		}, err
	}

	return &pb.GenerateKeyReply{
		Address:    address,
		PrivateKey: privateKey,
		Err:        "",
	}, err
}

func generateKey() (string, string, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	//fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", nil
	}

	//publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	//fmt.Println(address)

	return address, hexutil.Encode(privateKeyBytes)[2:], nil
}
