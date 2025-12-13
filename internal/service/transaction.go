package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
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
		err     error
		tx      string
		tmpUrl1 = "https://bsc-dataseed4.binance.org/"
		//tokenAddress = "0x337610d27c682E347C9cD60BD4b3b107C9d34dDd"
		//tmpUrl1      = "https://data-seed-prebsc-1-s3.binance.org:8545/"
	)

	if 0 == len(req.SendBody.ToAddr) {
		return &pb.SendTransactionReply{
			Tx: tx,
		}, nil
	}

	for i := 0; i <= 5; i++ {
		tx, err = toToken(req.SendBody.PrivateKey, req.SendBody.ToAddr, req.SendBody.Amount, req.SendBody.Token, tmpUrl1)
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
			fmt.Println("err info", req.SendBody.PrivateKey, req.SendBody.ToAddr, req.SendBody.Amount, req.SendBody.Token, tmpUrl1)
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

func (s *TransactionService) TokenBalance(ctx context.Context, req *pb.TokenBalanceRequest) (*pb.TokenBalanceReply, error) {
	//client, err := ethclient.Dial("https://data-seed-prebsc-1-s3.binance.org:8545/")
	client, err := ethclient.Dial("https://bsc-dataseed.binance.org/")
	if err != nil {
		return &pb.TokenBalanceReply{
			Balance: "",
			Err:     err.Error(),
		}, nil
	}

	tokenAddress := common.HexToAddress(req.Token)
	instance, err := NewUsdt(tokenAddress, client)
	if err != nil {
		return &pb.TokenBalanceReply{
			Balance: "",
			Err:     err.Error(),
		}, nil
	}

	address := common.HexToAddress(req.Address)
	bal, err := instance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return &pb.TokenBalanceReply{
			Balance: "",
			Err:     err.Error(),
		}, nil
	}

	return &pb.TokenBalanceReply{
		Balance: bal.String(),
		Err:     "",
	}, nil
}

func (s *TransactionService) Transaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionReply, error) {
	for i := 0; i <= 5; i++ {
		tmpUrl1 := ""
		if 0 == i {
			tmpUrl1 = "https://bsc-dataseed4.binance.org/"
		} else if 1 == i {
			tmpUrl1 = "https://bsc-dataseed3.binance.org/"
		} else if 2 == i {
			tmpUrl1 = "https://bsc-dataseed2.binance.org/"
		} else if 3 == i {
			tmpUrl1 = "https://bnb-bscnews.rpc.blxrbdn.com/"
		} else if 4 == i {
			tmpUrl1 = "https://bsc-dataseed.binance.org/"
		}

		// https://data-seed-prebsc-1-s3.binance.org:8545/
		// https://bsc-dataseed.binance.org/
		client, err := ethclient.Dial(tmpUrl1)
		if err != nil {
			continue
		}

		receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(req.GetTx()))
		if err != nil {
			continue
		}

		return &pb.TransactionReply{
			Status: receipt.Status,
		}, err
	}

	return &pb.TransactionReply{
		Status: 99,
	}, nil
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

func (s *TransactionService) VerifySig(ctx context.Context, req *pb.VerifySigRequest) (*pb.VerifySigReply, error) {
	sig := hexutil.MustDecode(req.Sign)
	msg := accounts.TextHash([]byte(req.Content))

	if sig[crypto.RecoveryIDOffset] == 27 || sig[crypto.RecoveryIDOffset] == 28 {
		sig[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	recovered, err := crypto.SigToPub(msg, sig)
	if err != nil {
		return &pb.VerifySigReply{
			Res:     false,
			Address: "",
		}, err
	}

	recoveredAddr := crypto.PubkeyToAddress(*recovered)

	sigPublicKeyBytes := crypto.FromECDSAPub(recovered)
	signatureNoRecoverID := sig[:len(sig)-1] // remove recovery id
	verified := crypto.VerifySignature(sigPublicKeyBytes, msg, signatureNoRecoverID)
	return &pb.VerifySigReply{
		Res:     verified,
		Address: recoveredAddr.Hex(),
	}, nil
}

func (s *TransactionService) GetReserves(ctx context.Context, req *pb.GetReservesRequest) (*pb.GetReservesReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress(req.Pair) // LP Pair 地址
		instance, err := NewPair(tokenAddress, client)
		if err != nil {
			fmt.Println("NewPair error:", err)
			continue
		}

		// 获取储备量
		reserves, err := instance.GetReserves(&bind.CallOpts{})
		if err != nil {
			fmt.Println("GetReserves error:", err)
			continue
		}

		return &pb.GetReservesReply{
			ReservesOne: reserves.Reserve0.String(),
			ReservesTwo: reserves.Reserve1.String(),
		}, nil
	}

	return &pb.GetReservesReply{
		ReservesOne: "-1",
		ReservesTwo: "-1",
	}, nil
}

func (s *TransactionService) GetAll(ctx context.Context, req *pb.GetAllRequest) (*pb.GetAllReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	tmpOne := "-1"
	tmpTwo := "-1"
	tmpThree := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress("0x18Ac4F491F4A365B5a66F5c4e44C2b311e516bC7")
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("NewPair error:", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			one, errOne := instance.GetReqLpArrayLength(&bind.CallOpts{})
			if errOne != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpOne = one.String()
		}

		if "-1" == tmpTwo {
			two, errTwo := instance.GetBuyArrayLength(&bind.CallOpts{})
			if errTwo != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpTwo = two.String()
		}

		if "-1" == tmpThree {
			three, errThree := instance.GetReqWithdrawArrayLength(&bind.CallOpts{})
			if errThree != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpThree = three.String()
		}

		if "-1" != tmpOne && "-1" != tmpTwo && "-1" != tmpThree {
			break
		}
	}

	return &pb.GetAllReply{
		ReqLpArrayLength:    tmpOne,
		BuyArrayLength:      tmpTwo,
		WithdrawArrayLength: tmpThree,
	}, nil
}

func (s *TransactionService) GetAllTwo(ctx context.Context, req *pb.GetAllTwoRequest) (*pb.GetAllTwoReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	tmpTwo := "-1"
	if 20 >= len(req.Contract) {
		return &pb.GetAllTwoReply{
			BuyArrayLength: tmpTwo,
		}, nil
	}

	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress(req.Contract)
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("NewPair error:", err)
			continue
		}

		if "-1" == tmpTwo {
			two, errTwo := instance.GetBuyArrayLength(&bind.CallOpts{})
			if errTwo != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpTwo = two.String()
		}

		if "-1" != tmpTwo {
			break
		}
	}

	return &pb.GetAllTwoReply{
		BuyArrayLength: tmpTwo,
	}, nil
}

func (s *TransactionService) GetBuyByOrderId(ctx context.Context, req *pb.GetBuyByOrderIdRequest) (*pb.GetBuyByOrderIdReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	contractAddress := "0x18Ac4F491F4A365B5a66F5c4e44C2b311e516bC7"
	if 20 < len(req.Contract) {
		contractAddress = req.Contract
	}

	tmpOne := "-1"
	tmpTwo := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress(contractAddress)
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("GetBuyByOrderId error:", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			start, _ := new(big.Int).SetString(req.OrderId, 10)
			one, errOne := instance.UsdtArray(&bind.CallOpts{}, start)
			if errOne != nil {
				fmt.Println("GetBuyByOrderId error:", err)
				continue
			}

			tmpOne = one.String()
		}

		if "-1" == tmpTwo {
			start, _ := new(big.Int).SetString(req.OrderId, 10)
			two, errTwo := instance.CatArray(&bind.CallOpts{}, start)
			if errTwo != nil {
				fmt.Println("GetBuyByOrderId error:", err)
				continue
			}

			tmpTwo = two.String()
		}

		if "-1" != tmpOne && "-1" != tmpTwo {
			break
		}
	}

	return &pb.GetBuyByOrderIdReply{
		UsdtArray: tmpOne,
		CatArray:  tmpTwo,
	}, nil
}

func (s *TransactionService) GetLpByOrderId(ctx context.Context, req *pb.GetLpByOrderIdRequest) (*pb.GetLpByOrderIdReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	tmpOne := "-1"
	tmpTwo := "-1"
	tmpThree := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress("0x18Ac4F491F4A365B5a66F5c4e44C2b311e516bC7")
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("GetLpByOrderId error:", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			start, _ := new(big.Int).SetString(req.OrderId, 10)
			one, errOne := instance.LpArray(&bind.CallOpts{}, start)
			if errOne != nil {
				fmt.Println("GetLpByOrderId error:", err)
				continue
			}

			tmpOne = one.String()
		}

		if "-1" == tmpTwo {
			start, _ := new(big.Int).SetString(req.OrderId, 10)
			two, errTwo := instance.LpUsdtArray(&bind.CallOpts{}, start)
			if errTwo != nil {
				fmt.Println("GetLpByOrderId error:", err)
				continue
			}

			tmpTwo = two.String()
		}

		if "-1" == tmpThree {
			start, _ := new(big.Int).SetString(req.OrderId, 10)
			three, errThree := instance.LpCatArray(&bind.CallOpts{}, start)
			if errThree != nil {
				fmt.Println("GetLpByOrderId error:", err)
				continue
			}

			tmpThree = three.String()
		}

		if "-1" != tmpOne && "-1" != tmpTwo && "-1" != tmpThree {
			break
		}
	}

	return &pb.GetLpByOrderIdReply{
		LpArray:     tmpOne,
		LpUsdtArray: tmpTwo,
		LpCatArray:  tmpThree,
	}, nil
}

func (s *TransactionService) GetArray(ctx context.Context, req *pb.GetArrayRequest) (*pb.GetArrayReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	res := make([]*pb.GetArrayReply_List, 0)
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		contractAddress := "0x18Ac4F491F4A365B5a66F5c4e44C2b311e516bC7"
		if 20 < len(req.Contract) {
			contractAddress = req.Contract
		}

		tokenAddress := common.HexToAddress(contractAddress)
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("GetArray error:", err)
			continue
		}

		// 获取储备量
		start, _ := new(big.Int).SetString(req.Start, 10)
		end, _ := new(big.Int).SetString(req.End, 10)
		if "buy" == req.ReqType {
			tmp, errOne := instance.GetBuyArray(&bind.CallOpts{}, start, end)
			if errOne != nil {
				//fmt.Println("GetArray error:", err)
				continue
			}

			if nil != tmp && 0 < len(tmp) {
				for _, vTmp := range tmp {
					res = append(res, &pb.GetArrayReply_List{
						OrderId: vTmp.String(),
					})
				}

				break
			}
		} else if "withdraw" == req.ReqType {
			tmp, errOne := instance.GetReqWithdrawArray(&bind.CallOpts{}, start, end)
			if errOne != nil {
				//fmt.Println("GetArray error:", err)
				continue
			}

			if nil != tmp && 0 < len(tmp) {
				for _, vTmp := range tmp {
					res = append(res, &pb.GetArrayReply_List{
						OrderId: vTmp.String(),
					})
				}

				break
			}
		} else {
			tmp, errOne := instance.GetReqLpArray(&bind.CallOpts{}, start, end)
			if errOne != nil {
				//fmt.Println("GetArray error:", err)
				continue
			}

			if nil != tmp && 0 < len(tmp) {
				for _, vTmp := range tmp {
					res = append(res, &pb.GetArrayReply_List{
						OrderId: vTmp.String(),
					})
				}

				break
			}
		}
	}

	return &pb.GetArrayReply{
		List: res,
	}, nil
}

func (s *TransactionService) AddLiquidity(ctx context.Context, req *pb.AddLiquidityRequest) (*pb.AddLiquidityReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress("0x18Ac4F491F4A365B5a66F5c4e44C2b311e516bC7")
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("AddLiquidity error:", err)
			continue
		}

		var (
			tx *types.Transaction
		)
		var authUser *bind.TransactOpts
		var privateKey *ecdsa.PrivateKey
		privateKey, err = crypto.HexToECDSA("")
		if err != nil {
			fmt.Println(err)
			continue
		}

		authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		orderId, _ := new(big.Int).SetString(req.SendBody.OrderId, 10)
		AmountUsdt, _ := new(big.Int).SetString(req.SendBody.AmountUsdt, 10)

		tx, err = instance.AddLiquidity(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		}, orderId, AmountUsdt)
		if err != nil {
			fmt.Println("AddLiquidity error:", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.AddLiquidityReply{
				Res: hashContent,
			}, nil
		}

		return &pb.AddLiquidityReply{
			Res: tx.Hash().Hex(),
		}, nil
	}

	return &pb.AddLiquidityReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) RemoveLiquidity(ctx context.Context, req *pb.RemoveLiquidityRequest) (*pb.RemoveLiquidityReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress("0x18Ac4F491F4A365B5a66F5c4e44C2b311e516bC7")
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("RemoveLiquidity error:", err)
			continue
		}

		var tx *types.Transaction
		var authUser *bind.TransactOpts

		var privateKey *ecdsa.PrivateKey
		privateKey, err = crypto.HexToECDSA("")
		if err != nil {
			fmt.Println(err)
			continue
		}

		authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		orderId, _ := new(big.Int).SetString(req.SendBody.OrderId, 10)
		lp, _ := new(big.Int).SetString(req.SendBody.LiquidityAmount, 10)

		tx, err = instance.RemoveLiquidity(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		}, orderId, lp)
		if err != nil {
			fmt.Println("RemoveLiquidity error:", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.RemoveLiquidityReply{
				Res: hashContent,
			}, nil
		}

		return &pb.RemoveLiquidityReply{
			Res: tx.Hash().Hex(),
		}, nil

	}

	return &pb.RemoveLiquidityReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) BuyAICAT(ctx context.Context, req *pb.BuyAICATRequest) (*pb.BuyAICATReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	contract := "0x18Ac4F491F4A365B5a66F5c4e44C2b311e516bC7"
	if 20 < len(req.SendBody.Contract) {
		contract = req.SendBody.Contract
	}

	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}
		//fmt.Println(req, req.SendBody.OrderId, req.SendBody.UsdtAmount)
		tokenAddress := common.HexToAddress(contract)
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("BuyAICAT error:", err)
			continue
		}

		var tx *types.Transaction
		var authUser *bind.TransactOpts

		var privateKey *ecdsa.PrivateKey
		privateKey, err = crypto.HexToECDSA("")
		if err != nil {
			fmt.Println(err)
			continue
		}

		authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		orderId, _ := new(big.Int).SetString(req.SendBody.OrderId, 10)
		amount, _ := new(big.Int).SetString(req.SendBody.UsdtAmount, 10)

		tx, err = instance.BuyAICAT(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		}, orderId, amount)
		if err != nil {
			fmt.Println("BuyAICAT error:", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.BuyAICATReply{
				Res: hashContent,
			}, nil
		}

		return &pb.BuyAICATReply{
			Res: tx.Hash().Hex(),
		}, nil

	}

	return &pb.BuyAICATReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) WithdrawAICAT(ctx context.Context, req *pb.WithdrawAICATRequest) (*pb.WithdrawAICATReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	contract := "0x18Ac4F491F4A365B5a66F5c4e44C2b311e516bC7"
	if 20 < len(req.SendBody.Contract) {
		contract = req.SendBody.Contract
	}

	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		//fmt.Println(req, req.SendBody.OrderId, req.SendBody.UsdtAmount)
		tokenAddress := common.HexToAddress(contract)
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("WithdrawAICAT error:", err)
			continue
		}

		var tx *types.Transaction
		var authUser *bind.TransactOpts

		var privateKey *ecdsa.PrivateKey
		privateKey, err = crypto.HexToECDSA("")
		if err != nil {
			fmt.Println(err)
			continue
		}

		authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		address := common.HexToAddress(req.SendBody.Address)
		//token := common.HexToAddress(req.SendBody.Token)
		token := common.HexToAddress("0xdA08bA041b901cd6F4744C924b6A111A2F57904e")
		amount, _ := new(big.Int).SetString(req.SendBody.Amount, 10)

		tx, err = instance.Withdraw(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		}, address, token, amount)
		if err != nil {
			fmt.Println("WithdrawAICAT error:", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.WithdrawAICATReply{
				Res: hashContent,
			}, nil
		}

		return &pb.WithdrawAICATReply{
			Res: tx.Hash().Hex(),
		}, nil

	}

	return &pb.WithdrawAICATReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) GetUserLp(ctx context.Context, req *pb.GetUserLpRequest) (*pb.GetUserLpReply, error) {
	//urls := []string{
	//	"https://bsc-dataseed4.binance.org/",
	//	"https://binance.llamarpc.com/",
	//	"https://bscrpc.com/",
	//	"https://bsc-pokt.nodies.app/",
	//	"https://data-seed-prebsc-1-s3.binance.org:8545/",
	//}

	tmpOne := "-1"
	tmpTwo := "-1"
	//for _, urlTmp := range urls {
	//	client, err := ethclient.Dial(urlTmp)
	//	if err != nil {
	//		fmt.Println("client error:", err)
	//		continue
	//	}
	//
	//	tokenAddress := common.HexToAddress("0x18Ac4F491F4A365B5a66F5c4e44C2b311e516bC7")
	//	instance, err := NewAdmin(tokenAddress, client)
	//	if err != nil {
	//		fmt.Println("GetUserLp error:", err)
	//		continue
	//	}
	//
	//	// 获取储备量
	//	if "-1" == tmpOne {
	//		address := common.HexToAddress(req.Address)
	//		one, errOne := instance.LpAmount(&bind.CallOpts{}, address)
	//		if errOne != nil {
	//			fmt.Println("GetUserLp error:", err)
	//			continue
	//		}
	//
	//		tmpOne = one.String()
	//	}
	//
	//	if "-1" == tmpTwo {
	//		two, errTwo := instance.LpAmountTotal(&bind.CallOpts{})
	//		if errTwo != nil {
	//			fmt.Println("GetUserLp error:", err)
	//			continue
	//		}
	//
	//		tmpTwo = two.String()
	//	}
	//
	//	if "-1" != tmpOne && "-1" != tmpTwo {
	//		break
	//	}
	//}

	return &pb.GetUserLpReply{
		LpAmount:      tmpOne,
		LpAmountTotal: tmpTwo,
	}, nil
}

func (s *TransactionService) GetDailyFee(ctx context.Context, req *pb.GetDailyFeeRequest) (*pb.GetDailyFeeReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	tmpOne := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress("0xdA08bA041b901cd6F4744C924b6A111A2F57904e")
		instance, err := NewEarth(tokenAddress, client)
		if err != nil {
			fmt.Println("GetDailyFee error:", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			start, _ := new(big.Int).SetString(req.Ts, 10)
			one, errOne := instance.DailyFee(&bind.CallOpts{}, start)
			if errOne != nil {
				fmt.Println("GetDailyFee error:", err)
				continue
			}

			tmpOne = one.String()
		}

		if "-1" != tmpOne {
			break
		}
	}

	return &pb.GetDailyFeeReply{Ammount: tmpOne}, nil
}

func (s *TransactionService) AddWhite(ctx context.Context, req *pb.AddWhiteRequest) (*pb.AddWhiteReply, error) {
	//f, err := excelize.OpenFile("../../abi/a.xlsx")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//
	//// 指定读取的工作表
	//sheets := []string{"工作表 1 - 25.31白地址", "要设置的白名单"}
	//addressSet := make(map[string]struct{})
	//
	//for _, sheet := range sheets {
	//	rows, err := f.GetRows(sheet)
	//	if err != nil {
	//		log.Printf("读取工作表 %s 出错: %v\n", sheet, err)
	//		continue
	//	}
	//	for i, row := range rows {
	//		for j, col := range row {
	//			if col != "" && len(col) == 42 && col[:2] == "0x" {
	//				addressSet[col] = struct{}{}
	//			}
	//			// 额外提取第一列和第一行
	//			if i == 0 || j == 0 {
	//				if col != "" && len(col) == 42 && col[:2] == "0x" {
	//					addressSet[col] = struct{}{}
	//				}
	//			}
	//		}
	//	}
	//}
	//
	//// 打印成 Go 的字面量数组形式
	//fmt.Println("var whitelist = []string{")
	//for addr := range addressSet {
	//	fmt.Printf("\t\"%s\",\n", addr)
	//}
	//fmt.Println("}")

	//urls := []string{
	//	"https://bsc-dataseed4.binance.org/",
	//	"https://binance.llamarpc.com/",
	//	"https://bscrpc.com/",
	//	"https://bsc-pokt.nodies.app/",
	//	"https://data-seed-prebsc-1-s3.binance.org:8545/",
	//}
	//for _, urlTmp := range urls {
	//	client, err := ethclient.Dial(urlTmp)
	//	if err != nil {
	//		fmt.Println("client error:", err)
	//		continue
	//	}
	//
	//	tokenAddress := common.HexToAddress("0xdA08bA041b901cd6F4744C924b6A111A2F57904e")
	//	instance, err := NewCat(tokenAddress, client)
	//	if err != nil {
	//		fmt.Println("GetDailyFee error:", err)
	//		continue
	//	}
	//
	//	var tx *types.Transaction
	//	var authUser *bind.TransactOpts
	//
	//	var privateKey *ecdsa.PrivateKey
	//	privateKey, err = crypto.HexToECDSA("")
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//
	//	authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
	//	if err != nil {
	//		fmt.Println(err)
	//		continue
	//	}
	//
	//	users := make([]common.Address, 0)
	//	for _, v := range whitelist {
	//		users = append(users, common.HexToAddress(v))
	//	}
	//
	//	tx, err = instance.SetWhiteMap(&bind.TransactOpts{
	//		From:     authUser.From,
	//		Signer:   authUser.Signer,
	//		GasLimit: 0,
	//	}, users, true)
	//	if err != nil {
	//		fmt.Println("BuyAICAT error:", err)
	//		continue
	//	}
	//
	//	if 0 >= len(tx.Hash().Hex()) {
	//		return &pb.AddWhiteReply{Res: tx.Hash().Hex()}, nil
	//	}
	//}

	return &pb.AddWhiteReply{}, nil
}

func (s *TransactionService) SendAICAT(ctx context.Context, req *pb.SendAICATRequest) (*pb.SendAICATReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}
		//fmt.Println(req, req.SendBody.OrderId, req.SendBody.UsdtAmount)
		tokenAddress := common.HexToAddress("0x4984BE11B25b22615f20B96d997AD64fbC7a28D3")
		//tokenAddress := common.HexToAddress("0x58B6F82Ce0f1a948dDD58824b16b878a0Bcd3259")
		instance, err := NewAdminNew(tokenAddress, client)
		if err != nil {
			fmt.Println("BuyAICAT error:", err)
			continue
		}

		var tx *types.Transaction
		var authUser *bind.TransactOpts

		var privateKey *ecdsa.PrivateKey
		privateKey, err = crypto.HexToECDSA("")
		if err != nil {
			fmt.Println(err)
			continue
		}

		authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		tx, err = instance.SendToDead(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		})
		if err != nil {
			fmt.Println("BuyAICAT error:", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.SendAICATReply{
				Res: hashContent,
			}, nil
		}

		return &pb.SendAICATReply{
			Res: tx.Hash().Hex(),
		}, nil

	}

	return &pb.SendAICATReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) GetBuyAICATByOrderId(ctx context.Context, req *pb.GetBuyAICATByOrderIdRequest) (*pb.GetBuyAICATByOrderIdReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	tmpOne := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress("0x4984BE11B25b22615f20B96d997AD64fbC7a28D3")
		//tokenAddress := common.HexToAddress("0x58B6F82Ce0f1a948dDD58824b16b878a0Bcd3259")
		instance, err := NewAdminNew(tokenAddress, client)
		if err != nil {
			fmt.Println("GetLpByOrderId error:", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			start, _ := new(big.Int).SetString(req.OrderId, 10)
			one, errOne := instance.ReqOrderIdArray(&bind.CallOpts{}, start)
			if errOne != nil {
				fmt.Println("GetLpByOrderId error:", err)
				continue
			}

			tmpOne = one.String()
		}
	}

	return &pb.GetBuyAICATByOrderIdReply{Amount: tmpOne}, nil
}

func (s *TransactionService) GetBoxAllLength(ctx context.Context, req *pb.GetBoxAllRequest) (*pb.GetBoxAllReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	tmpOne := "-1"
	tmpTwo := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress("0x290feBA6d39cA8ad4A77d40B6a1Cea0878E313Dd")
		instance, err := NewNft(tokenAddress, client)
		if err != nil {
			fmt.Println("GetBoxAll error:", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			one, errOne := instance.OpenActionTokenIdsLength(&bind.CallOpts{})
			if errOne != nil {
				fmt.Println("GetBoxAll error:", err)
				continue
			}

			tmpOne = one.String()
		}

		if "-1" == tmpTwo {
			two, errTwo := instance.MintedTokenIdsLength(&bind.CallOpts{})
			if errTwo != nil {
				fmt.Println("GetBoxAll error:", err)
				continue
			}

			tmpTwo = two.String()
		}

		if "-1" != tmpOne && "-1" != tmpTwo {
			break
		}
	}

	return &pb.GetBoxAllReply{
		NewLength:  tmpTwo,
		OpenLength: tmpOne,
	}, nil
}

func (s *TransactionService) GetBoxNew(ctx context.Context, req *pb.GetBoxNewRequest) (*pb.GetBoxNewReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	res := make([]*pb.GetBoxNewReply_List, 0)
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		contractAddress := "0x290feBA6d39cA8ad4A77d40B6a1Cea0878E313Dd"

		tokenAddress := common.HexToAddress(contractAddress)
		instance, err := NewNft(tokenAddress, client)
		if err != nil {
			fmt.Println("GetBoxNew error:", err)
			continue
		}

		// 获取认购盲盒记录
		start, _ := new(big.Int).SetString(req.Start, 10)
		end, _ := new(big.Int).SetString(req.End, 10)
		tmp, errOne := instance.GetMintedByPage(&bind.CallOpts{}, start, end)
		if errOne != nil {
			//fmt.Println("GetArray error:", err)
			continue
		}

		if len(tmp.Infos) != len(tmp.TokenIds) {

		}

		for k, vTmp := range tmp.TokenIds {
			res = append(res, &pb.GetBoxNewReply_List{
				TokenId:  vTmp.String(),
				Price:    tmp.Infos[k].UsdtPaid.String(),
				OpenedAt: tmp.Infos[k].OpenedAt,
				MintAt:   tmp.Infos[k].MintedAt,
				Reward:   tmp.Infos[k].Reward.String(),
				RewardAt: tmp.Infos[k].RewardSetAt,
			})
		}
	}

	return &pb.GetBoxNewReply{
		List: res,
	}, nil
}

func (s *TransactionService) GetBoxOpen(ctx context.Context, req *pb.GetBoxOpenRequest) (*pb.GetBoxOpenReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	res := make([]*pb.GetBoxOpenReply_List, 0)
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		contractAddress := "0x290feBA6d39cA8ad4A77d40B6a1Cea0878E313Dd"

		tokenAddress := common.HexToAddress(contractAddress)
		instance, err := NewNft(tokenAddress, client)
		if err != nil {
			fmt.Println("GetBoxNew error:", err)
			continue
		}

		// 获取认购盲盒记录
		start, _ := new(big.Int).SetString(req.Start, 10)
		end, _ := new(big.Int).SetString(req.End, 10)
		tmp, errOne := instance.GetOpenedByPage(&bind.CallOpts{}, start, end)
		if errOne != nil {
			//fmt.Println("GetArray error:", err)
			continue
		}

		if len(tmp.Infos) != len(tmp.TokenIds) {

		}

		for k, vTmp := range tmp.TokenIds {
			res = append(res, &pb.GetBoxOpenReply_List{
				TokenId:  vTmp.String(),
				Price:    tmp.Infos[k].UsdtPaid.String(),
				OpenedAt: tmp.Infos[k].OpenedAt,
				MintAt:   tmp.Infos[k].MintedAt,
				Reward:   tmp.Infos[k].Reward.String(),
				RewardAt: tmp.Infos[k].RewardSetAt,
			})
		}
	}

	return &pb.GetBoxOpenReply{
		List: res,
	}, nil
}

func (s *TransactionService) SetReward(ctx context.Context, req *pb.SetRewardRequest) (*pb.SetRewardReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	contract := "0x290feBA6d39cA8ad4A77d40B6a1Cea0878E313Dd"
	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}
		//fmt.Println(req, req.SendBody.OrderId, req.SendBody.UsdtAmount)
		tokenAddress := common.HexToAddress(contract)
		instance, err := NewNft(tokenAddress, client)
		if err != nil {
			fmt.Println("SetReward error:", err)
			continue
		}

		var tx *types.Transaction
		var authUser *bind.TransactOpts

		var privateKey *ecdsa.PrivateKey
		privateKey, err = crypto.HexToECDSA("")
		if err != nil {
			fmt.Println(err)
			continue
		}

		authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			continue
		}

		orderId, _ := new(big.Int).SetString(req.SendBody.TokenId, 10)
		amount, _ := new(big.Int).SetString(req.SendBody.Reward, 10)

		tx, err = instance.SetRewardOnce(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 2000000, // 固定 200万 gas limit
		}, orderId, amount)
		if err != nil {
			fmt.Println("SetReward error:", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.SetRewardReply{
				Res: hashContent,
			}, nil
		}

		return &pb.SetRewardReply{
			Res: tx.Hash().Hex(),
		}, nil

	}

	return &pb.SetRewardReply{
		Res: hashContent,
	}, nil
}
