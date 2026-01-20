package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-kratos/kratos/v2/log"
	"math/big"
	"requestEth/internal/biz"
	"strings"
	"time"

	pb "requestEth/api/requestEth/v1"
)

type TransactionService struct {
	pb.UnimplementedTransactionServer
	log *log.Helper
	ac  *biz.AppUsecase
}

func NewTransactionService(ac *biz.AppUsecase, logger log.Logger) *TransactionService {
	return &TransactionService{
		ac:  ac,
		log: log.NewHelper(logger),
	}
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

		tokenAddress := common.HexToAddress("0xE447b0391d3F03befeC0dC09E25c049132618fd9")
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

		contractAddress := "0xE447b0391d3F03befeC0dC09E25c049132618fd9"

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

		for _, vTmp := range tmp {
			res = append(res, &pb.GetBoxNewReply_List{
				TokenId:  vTmp.TokenId.String(),
				Price:    vTmp.UsdtPaid.String(),
				OpenedAt: vTmp.OpenedAt,
				MintAt:   vTmp.MintedAt,
				Reward:   vTmp.Reward.String(),
				RewardAt: vTmp.RewardSetAt,
			})
		}

		break
	}

	return &pb.GetBoxNewReply{
		List: res,
	}, nil
}

func (s *TransactionService) GetExchangeList(ctx context.Context, req *pb.GetExchangeListRequest) (*pb.GetExchangeListReply, error) {
	return s.ac.GetExchangeList(ctx, req)
}

func (s *TransactionService) GetBuyList(ctx context.Context, req *pb.GetBuyListRequest) (*pb.GetBuyListReply, error) {
	return s.ac.GetBuyList(ctx, req)
}

func (s *TransactionService) GetRewardList(ctx context.Context, req *pb.GetRewardListRequest) (*pb.GetRewardListReply, error) {
	return s.ac.GetRewardList(ctx, req)
}

func (s *TransactionService) GetBuyBoxList(ctx context.Context, req *pb.GetBuyBoxListRequest) (*pb.GetBuyBoxListReply, error) {
	return s.ac.GetBuyBoxList(ctx, req)
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

		contractAddress := "0xE447b0391d3F03befeC0dC09E25c049132618fd9"

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

		for _, vTmp := range tmp {
			res = append(res, &pb.GetBoxOpenReply_List{
				TokenId:  vTmp.TokenId.String(),
				Price:    vTmp.UsdtPaid.String(),
				OpenedAt: vTmp.OpenedAt,
				MintAt:   vTmp.MintedAt,
				Reward:   vTmp.Reward.String(),
				RewardAt: vTmp.RewardSetAt,
			})
		}

		break
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

	contract := "0xE447b0391d3F03befeC0dC09E25c049132618fd9"
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

func (s *TransactionService) SetRewardTwo(ctx context.Context, req *pb.SetRewardRequest) (*pb.SetRewardReply, error) {
	urls := []string{
		"https://bsc-dataseed4.binance.org/",
		"https://binance.llamarpc.com/",
		"https://bscrpc.com/",
		"https://bsc-pokt.nodies.app/",
		"https://data-seed-prebsc-1-s3.binance.org:8545/",
	}

	contract := "0x47CC8586f34ecdbbcE595faf7A1402139A37eb23"
	hashContent := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}
		//fmt.Println(req, req.SendBody.OrderId, req.SendBody.UsdtAmount)
		tokenAddress := common.HexToAddress(contract)
		instance, err := NewStake(tokenAddress, client)
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

func (s *TransactionService) GetBoxRewardEvent(ctx context.Context, req *pb.GetBoxRewardEventRequest) (*pb.GetBoxRewardEventReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)
	for i := 1; i <= 19; i++ {
		urls := []string{
			"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
		}

		last := uint64(0)

		var (
			rLast *biz.RewardNotified
			errT  error
		)
		rLast, errT = s.ac.GetRewardLast(ctx)
		if nil != errT {
			return nil, errT
		}

		if nil != rLast {
			last = rLast.BlockNumber
		}

		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		var (
			events  []RewardNotifiedEvent
			newLast uint64
		)

		for _, url := range urls {
			client, err := ethclient.DialContext(ctx, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			events, newLast, err = PollRewardNotifiedIncremental(ctx, client, last)
			if err != nil {
				fmt.Println(err)
				// 换下一个 RPC
				continue
			}

			if 0 >= len(events) {
				time.Sleep(3 * time.Second)
				continue
			}

			if last >= newLast {
				continue
			}

			for _, v := range events {
				if last >= v.BlockNumber {
					break
				}

				err = s.ac.InsertReward(ctx, &biz.RewardNotified{
					BlockNumber:      v.BlockNumber,
					BlockTime:        v.BlockTime,
					LogIndex:         uint32(v.LogIndex),
					User:             v.User.String(),
					L1:               v.L1.String(),
					L2:               v.L2.String(),
					Profit:           BigIntToFloat64(v.Profit, 18),
					UserShare:        BigIntToFloat64(v.UserShare, 18),
					Top:              v.Top.String(),
					Pool:             BigIntToFloat64(v.Pool, 18),
					UplinePortionBps: v.UplinePortionBps.Uint64(),
					ToL1:             BigIntToFloat64(v.ToL1, 18),
					ToL2:             BigIntToFloat64(v.ToL2, 18),
					ToTop:            BigIntToFloat64(v.ToTop, 18),
					ToProject:        BigIntToFloat64(v.ToProject, 18),
				})
				if nil != err {
					fmt.Println("insert reward list err", err)
				}
			}
		}
	}

	return &pb.GetBoxRewardEventReply{}, nil
}

func (s *TransactionService) GetExchangeEvent(ctx context.Context, req *pb.GetExchangeEventRequest) (*pb.GetExchangeEventReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)
	for i := 1; i <= 19; i++ {
		urls := []string{
			"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
		}

		last := uint64(0)

		var (
			rLast *biz.SwapTrade
			errT  error
		)
		rLast, errT = s.ac.GetSwapTradeLast(ctx)
		if nil != errT {
			return nil, errT
		}

		if nil != rLast {
			last = rLast.BlockNumber
		}

		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		var (
			events  []SwapEvent
			newLast uint64
		)

		for _, url := range urls {
			client, err := ethclient.DialContext(ctx, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			events, newLast, err = PollSwapIncremental(ctx, client, last)
			if err != nil {
				fmt.Println(err)
				// 换下一个 RPC
				continue
			}

			if 0 >= len(events) {
				time.Sleep(3 * time.Second)
				continue
			}

			if last >= newLast {
				continue
			}

			for _, v := range events {
				if last >= v.BlockNumber {
					break
				}

				tmpSide := uint8(0)
				if 0 == v.Amount0In.Int64() { // 买
					tmpSide = 1
				} else if 0 == v.Amount1In.Int64() { // 卖
					tmpSide = 2
				}

				err = s.ac.InsertSwapTrade(ctx, &biz.SwapTrade{
					BlockNumber:     v.BlockNumber,
					LogIndex:        uint32(v.LogIndex),
					BlockTime:       v.BlockTime,
					Sender:          v.Sender.String(),
					ToAddr:          v.To.String(),
					Side:            tmpSide,
					Amount0In:       BigIntToFloat64(v.Amount0In, 18),
					Amount1In:       BigIntToFloat64(v.Amount1In, 18),
					Amount0OutNet:   BigIntToFloat64(v.Amount0OutNet, 18),
					Amount1OutGross: BigIntToFloat64(v.Amount1OutGross, 18),
					Amount0OutGross: BigIntToFloat64(v.Amount0OutGross, 18),
				})
				if nil != err {
					fmt.Println("insert swap trade err", err)
				}
			}

			return &pb.GetExchangeEventReply{}, nil
		}
	}

	return &pb.GetExchangeEventReply{}, nil
}

func (s *TransactionService) GetBuyEvent(ctx context.Context, req *pb.GetBuyEventRequest) (*pb.GetBuyEventReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)
	for i := 1; i <= 19; i++ {
		urls := []string{
			"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
		}

		last := uint64(0)

		var (
			rLast *biz.PrimaryBuy
			errT  error
		)
		rLast, errT = s.ac.GetBuyLast(ctx)
		if nil != errT {
			return nil, errT
		}

		if nil != rLast {
			last = rLast.BlockNumber
		}

		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		var (
			events  []BoughtEvent
			newLast uint64
		)

		for _, url := range urls {
			client, err := ethclient.DialContext(ctx, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			events, newLast, err = PollBoughtIncremental(ctx, client, last)
			if err != nil {
				fmt.Println(err)
				// 换下一个 RPC
				continue
			}

			if 0 >= len(events) {
				time.Sleep(3 * time.Second)
				continue
			}

			if last >= newLast {
				continue
			}

			for _, v := range events {
				if last >= v.BlockNumber {
					break
				}

				err = s.ac.InsertBuyTrade(ctx, &biz.PrimaryBuy{
					BlockNumber:  v.BlockNumber,
					BlockTime:    v.BlockTime,
					LogIndex:     uint32(v.LogIndex),
					Buyer:        v.Buyer.String(),
					ToAddr:       v.To.String(),
					UsdtUsed:     BigIntToFloat64(v.UsdtUsed, 18),
					AusdGrossOut: BigIntToFloat64(v.AusdGrossOut, 18),
					AusdFee:      BigIntToFloat64(v.AusdFee, 18),
					AusdNetOut:   BigIntToFloat64(v.AusdNetOut, 18),
					PriceBefore:  BigIntToFloat64(v.PriceBefore, 18),
					PriceAfter:   BigIntToFloat64(v.PriceAfter, 18),
				})
				if nil != err {
					fmt.Println("insert buy err", err)
				}
			}

			return &pb.GetBuyEventReply{}, nil
		}
	}

	return &pb.GetBuyEventReply{}, nil
}

func (s *TransactionService) GetSellEvent(ctx context.Context, req *pb.GetSellEventRequest) (*pb.GetSellEventReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)
	for i := 1; i <= 19; i++ {
		urls := []string{
			"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
		}

		last := uint64(0)

		var (
			rLast *biz.PrimarySell
			errT  error
		)
		rLast, errT = s.ac.GetSellLast(ctx)
		if nil != errT {
			return nil, errT
		}

		if nil != rLast {
			last = rLast.BlockNumber
		}

		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		var (
			events  []SoldEvent
			newLast uint64
		)

		for _, url := range urls {
			client, err := ethclient.DialContext(ctx, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			events, newLast, err = PollSoldIncremental(ctx, client, last)
			if err != nil {
				fmt.Println(err)
				// 换下一个 RPC
				continue
			}

			if 0 >= len(events) {
				time.Sleep(3 * time.Second)
				continue
			}

			if last >= newLast {
				continue
			}

			for _, v := range events {
				if last >= v.BlockNumber {
					break
				}

				err = s.ac.InsertSellTrade(ctx, &biz.PrimarySell{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    uint32(v.LogIndex),
					Seller:      v.Seller.String(),
					ToAddr:      v.To.String(),
					AusdGrossIn: BigIntToFloat64(v.AusdGrossIn, 18),
					AusdBurn:    BigIntToFloat64(v.AusdBurn, 18),
					AusdFee:     BigIntToFloat64(v.AusdFee, 18),
					UsdtOut:     BigIntToFloat64(v.UsdtOut, 18),
					PriceBefore: BigIntToFloat64(v.PriceBefore, 18),
					PriceAfter:  BigIntToFloat64(v.PriceAfter, 18),
				})
				if nil != err {
					fmt.Println("insert sell err", err)
				}
			}

			return &pb.GetSellEventReply{}, nil
		}
	}

	return &pb.GetSellEventReply{}, nil
}

func (s *TransactionService) GetBoxBuyEvent(ctx context.Context, req *pb.GetBoxBuyEventRequest) (*pb.GetBoxBuyEventReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)
	for i := 1; i <= 19; i++ {
		urls := []string{
			"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
		}

		last := uint64(0)

		var (
			rLast *biz.NftMarketPurchase
			errT  error
		)
		rLast, errT = s.ac.GetNftMarketPurchaseLast(ctx)
		if nil != errT {
			return nil, errT
		}

		if nil != rLast {
			last = rLast.BlockNumber
		}

		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		var (
			events  []PurchasedEvent
			newLast uint64
		)

		for _, url := range urls {
			client, err := ethclient.DialContext(ctx, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			events, newLast, err = PollPurchasedIncremental(ctx, client, last)
			if err != nil {
				fmt.Println(err)
				// 换下一个 RPC
				continue
			}

			if 0 >= len(events) {
				time.Sleep(3 * time.Second)
				continue
			}

			if last >= newLast {
				continue
			}

			for _, v := range events {
				if last >= v.BlockNumber {
					break
				}

				tmpPaid := uint8(1)
				if v.FeePaidInB {
					tmpPaid = 2
				}
				err = s.ac.InsertNftMarketPurchase(ctx, &biz.NftMarketPurchase{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					Buyer:       v.Buyer.String(),
					Seller:      v.Seller.String(),
					TokenID:     v.TokenID.Uint64(),
					PriceUSDT:   BigIntToFloat64(v.PriceUSDT, 18),
					FeePaidInB:  tmpPaid,
					FeeUSDT:     BigIntToFloat64(v.FeeUSDT, 18),
					FeeB:        BigIntToFloat64(v.FeeB, 18),
				})
				if nil != err {
					fmt.Println("insert sell err", err)
				}
			}

			return &pb.GetBoxBuyEventReply{}, nil
		}
	}

	return &pb.GetBoxBuyEventReply{}, nil
}

func (s *TransactionService) GetBoxMintEvent(ctx context.Context, req *pb.GetBoxMintEventRequest) (*pb.GetBoxMintEventReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)
	for i := 1; i <= 19; i++ {
		urls := []string{
			"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
		}

		last := uint64(0)

		var (
			rLast *biz.NftMinted
			errT  error
		)
		rLast, errT = s.ac.GetNftMintedLast(ctx)
		if nil != errT {
			return nil, errT
		}

		if nil != rLast {
			last = rLast.BlockNumber
		}

		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		var (
			events  []MintedEvent
			newLast uint64
		)

		for _, url := range urls {
			client, err := ethclient.DialContext(ctx, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			events, newLast, err = PollMintedIncremental(ctx, client, last)
			if err != nil {
				fmt.Println(err)
				// 换下一个 RPC
				continue
			}

			if 0 >= len(events) {
				time.Sleep(3 * time.Second)
				continue
			}

			if last >= newLast {
				continue
			}

			for _, v := range events {
				if last >= v.BlockNumber {
					break
				}

				err = s.ac.InsertNftMinted(ctx, &biz.NftMinted{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					ToAddr:      v.To.String(),
					TokenID:     v.TokenID.Uint64(),
					Tier:        uint64(v.Tier),
					UsdtPaid:    BigIntToFloat64(v.UsdtPaid, 18),
					Status:      0,
					ListedAt:    0,
					OpenStatus:  0,
					OpenedAt:    0,
				})
				if nil != err {
					fmt.Println("insert minted box err", err)
				}
			}

			return &pb.GetBoxMintEventReply{}, nil
		}
	}

	return &pb.GetBoxMintEventReply{}, nil
}

func (s *TransactionService) GetBoxListEvent(ctx context.Context, req *pb.GetBoxListEventRequest) (*pb.GetBoxListEventReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)
	for i := 1; i <= 19; i++ {
		urls := []string{
			"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
		}

		last := uint64(0)

		var (
			rLast *biz.NftMarketListed
			errT  error
		)
		rLast, errT = s.ac.GetNftMarketListedLast(ctx)
		if nil != errT {
			return nil, errT
		}

		if nil != rLast {
			last = rLast.BlockNumber
		}

		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		var (
			events  []ListedEvent
			newLast uint64
		)

		for _, url := range urls {
			client, err := ethclient.DialContext(ctx, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			events, newLast, err = PollListedIncremental(ctx, client, last)
			if err != nil {
				fmt.Println(err)
				// 换下一个 RPC
				continue
			}

			if 0 >= len(events) {
				time.Sleep(3 * time.Second)
				continue
			}

			if last >= newLast {
				continue
			}

			for _, v := range events {
				if last >= v.BlockNumber {
					break
				}

				err = s.ac.InsertNftMarketListed(ctx, &biz.NftMarketListed{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					Seller:      v.Seller.String(),
					TokenID:     v.TokenID.Uint64(),
					Timestamp:   v.Timestamp.Uint64(),
				})
				if nil != err {
					fmt.Println("insert list box err", err)
				}
			}

			return &pb.GetBoxListEventReply{}, nil
		}
	}

	return &pb.GetBoxListEventReply{}, nil
}

func (s *TransactionService) GetBoxUnListEvent(ctx context.Context, req *pb.GetBoxUnListEventRequest) (*pb.GetBoxUnListEventReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)
	for i := 1; i <= 19; i++ {
		urls := []string{
			"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
		}

		last := uint64(0)

		var (
			rLast *biz.NftMarketUnlisted
			errT  error
		)
		rLast, errT = s.ac.GetNftMarketUnlistedLast(ctx)
		if nil != errT {
			return nil, errT
		}

		if nil != rLast {
			last = rLast.BlockNumber
		}

		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		var (
			events  []UnlistedEvent
			newLast uint64
		)

		for _, url := range urls {
			client, err := ethclient.DialContext(ctx, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			events, newLast, err = PollUnlistedIncremental(ctx, client, last)
			if err != nil {
				fmt.Println(err)
				// 换下一个 RPC
				continue
			}

			if 0 >= len(events) {
				time.Sleep(3 * time.Second)
				continue
			}

			if last >= newLast {
				continue
			}

			for _, v := range events {
				if last >= v.BlockNumber {
					break
				}

				err = s.ac.InsertNftMarketUnlisted(ctx, &biz.NftMarketUnlisted{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					Operator:    v.Operator.String(),
					TokenID:     v.TokenID.Uint64(),
				})
				if nil != err {
					fmt.Println("insert un list box err", err)
				}
			}

			return &pb.GetBoxUnListEventReply{}, nil
		}
	}

	return &pb.GetBoxUnListEventReply{}, nil
}

func (s *TransactionService) GetBoxOpenEvent(ctx context.Context, req *pb.GetBoxOpenEventRequest) (*pb.GetBoxOpenEventReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)
	for i := 1; i <= 19; i++ {
		urls := []string{
			"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
		}

		last := uint64(0)

		var (
			rLast *biz.NftOpened
			errT  error
		)
		rLast, errT = s.ac.GetNftOpenedLast(ctx)
		if nil != errT {
			return nil, errT
		}

		if nil != rLast {
			last = rLast.BlockNumber
		}

		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		var (
			events  []OpenedEvent
			newLast uint64
		)

		for _, url := range urls {
			client, err := ethclient.DialContext(ctx, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			events, newLast, err = PollOpenedIncremental(ctx, client, last)
			if err != nil {
				fmt.Println(err)
				// 换下一个 RPC
				continue
			}

			if 0 >= len(events) {
				time.Sleep(3 * time.Second)
				continue
			}

			if last >= newLast {
				continue
			}

			for _, v := range events {
				if last >= v.BlockNumber {
					break
				}

				err = s.ac.InsertNftOpened(ctx, &biz.NftOpened{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					UserAddr:    v.Owner.String(),
					TokenID:     v.TokenID.Uint64(),
					OpenedAt:    v.Timestamp.Uint64(),
					Reward:      BigIntToFloat64(v.BnbPaid, 18),
				})
				if nil != err {
					fmt.Println("insert open box err", err)
				}
			}

			return &pb.GetBoxOpenEventReply{}, nil
		}
	}

	return &pb.GetBoxOpenEventReply{}, nil
}

func (s *TransactionService) GetBoxTransferEvent(ctx context.Context, req *pb.GetBoxTransferEventRequest) (*pb.GetBoxTransferEventReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)
	for i := 1; i <= 10; i++ {
		urls := []string{
			"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
		}

		last := uint64(0)

		var (
			rLast *biz.NftTransfer
			errT  error
		)
		rLast, errT = s.ac.GetNftTransferLast(ctx)
		if nil != errT {
			return nil, errT
		}

		if nil != rLast {
			last = rLast.BlockNumber
		}

		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		var (
			events  []TransferEvent
			newLast uint64
		)

		for _, url := range urls {
			client, err := ethclient.DialContext(ctx, url)
			if err != nil {
				fmt.Println(err)
				continue
			}
			events, newLast, err = PollTransferIncremental(ctx, client, last)
			if err != nil {
				fmt.Println(err)
				// 换下一个 RPC
				continue
			}

			if 0 >= len(events) {
				time.Sleep(3 * time.Second)
				continue
			}

			if last >= newLast {
				continue
			}

			for _, v := range events {
				if last >= v.BlockNumber {
					break
				}

				err = s.ac.InsertNftTransfer(ctx, &biz.NftTransfer{
					BlockNumber: v.BlockNumber,
					BlockTime:   v.BlockTime,
					LogIndex:    v.LogIndex,
					FromAddr:    v.From.String(),
					ToAddr:      v.To.String(),
					TokenID:     v.TokenID.Uint64(),
				})
				if nil != err {
					fmt.Println("insert tran box err", err)
				}
			}
		}

		time.Sleep(4 * time.Second)
	}

	return &pb.GetBoxTransferEventReply{}, nil
}

func (s *TransactionService) UpdateBox(ctx context.Context, req *pb.UpdateBoxRequest) (*pb.UpdateBoxReply, error) {
	end := time.Now().UTC().Add(50 * time.Second)
	for i := 1; i <= 19; i++ {
		s.ac.UpdateBox(ctx, req)

		now := time.Now().UTC()
		if end.Before(now) {
			break
		}

		time.Sleep(3 * time.Second)
	}

	return &pb.UpdateBoxReply{}, nil
}

func (s *TransactionService) GetAddressBox(ctx context.Context, req *pb.GetAddressBoxRequest) (*pb.GetAddressBoxReply, error) {
	urls := []string{
		"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
	}

	var (
		count int64
	)
	res := make([]*pb.GetAddressBoxReply_List, 0)
	tmpTwo := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress("0x36e79a79829f3a0dad3f928d1ea8c28965b68237")
		instance, err := NewFactory(tokenAddress, client)
		if err != nil {
			fmt.Println("NewPair error:", err)
			continue
		}

		if "-1" == tmpTwo {
			two, errTwo := instance.VaultOf(&bind.CallOpts{}, common.HexToAddress(req.Address))
			if errTwo != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpTwo = two.String()
		}

		addresses := make([]string, 0)
		addresses = append(addresses, req.Address)
		if 10 < len(tmpTwo) {
			addresses = append(addresses, tmpTwo)
		}

		var (
			minted []*biz.NftMinted
		)
		minted, count, err = s.ac.GetAddressBox(ctx, addresses, req)
		for _, v := range minted {
			tmpPriceUsdtNow := float64(0)
			tokenAddressTmp := common.HexToAddress("0xE447b0391d3F03befeC0dC09E25c049132618fd9")
			instanceTmp, _ := NewNft(tokenAddressTmp, client)
			if nil != instanceTmp {
				var (
					tmpRes *big.Int
				)
				tmpRes, _ = instanceTmp.SecondaryBuyPrice(&bind.CallOpts{}, new(big.Int).SetUint64(v.TokenID))
				if nil != tmpRes {
					tmpPriceUsdtNow = BigIntToFloat64(tmpRes, 18)
				}
			}

			res = append(res, &pb.GetAddressBoxReply_List{
				TokenId:      v.TokenID,
				PriceUsdt:    v.UsdtPaid,
				PriceUsdtNow: tmpPriceUsdtNow,
				BlockTime:    v.BlockTime,
				ListTime:     v.ListedAt,
				Tier:         v.Tier,
				OpenStatus:   uint64(v.OpenStatus),
			})
		}
	}

	return &pb.GetAddressBoxReply{
		Count: uint64(count),
		List:  res,
	}, nil
}

func (s *TransactionService) GetMarketList(ctx context.Context, req *pb.GetMarketListRequest) (*pb.GetMarketListReply, error) {
	urls := []string{
		"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
	}

	var (
		count int64
	)
	res := make([]*pb.GetMarketListReply_List, 0)
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		var (
			minted []*biz.NftMinted
		)
		minted, count, err = s.ac.GetNftMintedPage(ctx, req)
		for _, v := range minted {
			tmpPriceUsdtNow := float64(0)
			tokenAddressTmp := common.HexToAddress("0xE447b0391d3F03befeC0dC09E25c049132618fd9")
			instanceTmp, _ := NewNft(tokenAddressTmp, client)
			if nil != instanceTmp {
				var (
					tmpRes *big.Int
				)
				tmpRes, _ = instanceTmp.SecondaryBuyPrice(&bind.CallOpts{}, new(big.Int).SetUint64(v.TokenID))
				if nil != tmpRes {
					tmpPriceUsdtNow = BigIntToFloat64(tmpRes, 18)
				}
			}

			res = append(res, &pb.GetMarketListReply_List{
				TokenId:      v.TokenID,
				PriceUsdt:    v.UsdtPaid,
				PriceUsdtNow: tmpPriceUsdtNow,
				BlockTime:    v.BlockTime,
				ListTime:     v.ListedAt,
				Tier:         v.Tier,
			})
		}
	}

	return &pb.GetMarketListReply{
		Count: uint64(count),
		List:  res,
	}, nil
}

func (s *TransactionService) GetSellBoxList(ctx context.Context, req *pb.GetSellBoxListRequest) (*pb.GetSellBoxListReply, error) {
	urls := []string{
		"https://bnb56743.allnodes.me:8545/hkrpfUWKCrv7Jio2",
	}

	var (
		count int64
	)
	res := make([]*pb.GetSellBoxListReply_List, 0)
	tmpTwo := "-1"
	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress("0x36e79a79829f3a0dad3f928d1ea8c28965b68237")
		instance, err := NewFactory(tokenAddress, client)
		if err != nil {
			fmt.Println("NewPair error:", err)
			continue
		}

		if "-1" == tmpTwo {
			two, errTwo := instance.VaultOf(&bind.CallOpts{}, common.HexToAddress(req.Address))
			if errTwo != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpTwo = two.String()
		}

		addresses := make([]string, 0)
		addresses = append(addresses, req.Address)
		if 10 < len(tmpTwo) {
			addresses = append(addresses, tmpTwo)
		}

		var (
			minted []*biz.NftMarketPurchase
		)
		minted, count, err = s.ac.GetSellBoxList(ctx, addresses, req)
		for _, v := range minted {
			res = append(res, &pb.GetSellBoxListReply_List{
				TokenId:   v.TokenID,
				PriceUsdt: v.PriceUSDT,
				FeeUsdt:   v.FeeUSDT,
				FeeB:      v.FeeB,
				BlockTime: v.BlockTime,
			})
		}
	}

	return &pb.GetSellBoxListReply{
		Count: uint64(count),
		List:  res,
	}, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	// DeployBlock TODO: 你把这里替换成实际部署块高（从 BscScan/OKLink 的 Contract Creation 里看）
	DeployBlock uint64 = 73406834

	// Confirmations 防回滚确认数：最新块往回退 6 个块
	Confirmations uint64 = 6

	// QueryStep 每次 FilterLogs 的分段大小（BSC 公共 RPC 经常会限制范围，建议 2k~5k）
	QueryStep uint64 = 2000
)

var (
	DividendContract = common.HexToAddress("0x135d05e0d1a8083525D64B9C9831579D4B6811d2")
)

// 仅包含 RewardNotified 事件（够用就行）
const umbrellaStakeDividendABI = `[
  {
		"anonymous": false,
		"inputs": [
			{
				"indexed": true,
				"internalType": "address",
				"name": "user",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "profit",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "userShare",
				"type": "uint256"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "l1",
				"type": "address"
			},
			{
				"indexed": true,
				"internalType": "address",
				"name": "l2",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "address",
				"name": "top",
				"type": "address"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "pool",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "uplinePortionBps",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "toL1",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "toL2",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "toTop",
				"type": "uint256"
			},
			{
				"indexed": false,
				"internalType": "uint256",
				"name": "toProject",
				"type": "uint256"
			}
		],
		"name": "RewardNotified",
		"type": "event"
	}
]`

type RewardNotifiedEvent struct {
	// 用于唯一定位/落库去重
	BlockNumber uint64
	BlockTime   uint64
	TxHash      common.Hash
	LogIndex    uint

	// indexed
	User common.Address
	L1   common.Address
	L2   common.Address

	// data（非 indexed）
	Profit           *big.Int
	UserShare        *big.Int
	Top              common.Address
	Pool             *big.Int
	UplinePortionBps *big.Int
	ToL1             *big.Int
	ToL2             *big.Int
	ToTop            *big.Int
	ToProject        *big.Int
}

func parseRewardNotified(a abi.ABI, lg types.Log) (RewardNotifiedEvent, error) {
	ev, ok := a.Events["RewardNotified"]
	if !ok {
		return RewardNotifiedEvent{}, fmt.Errorf("event RewardNotified not found in ABI")
	}
	if len(lg.Topics) != 4 {
		return RewardNotifiedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return RewardNotifiedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := RewardNotifiedEvent{
		BlockNumber: lg.BlockNumber,
		TxHash:      lg.TxHash,
		LogIndex:    lg.Index,

		User: common.BytesToAddress(lg.Topics[1].Bytes()),
		L1:   common.BytesToAddress(lg.Topics[2].Bytes()),
		L2:   common.BytesToAddress(lg.Topics[3].Bytes()),
	}

	// data 部分按事件里“非 indexed inputs”的顺序解
	values, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return RewardNotifiedEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	// 顺序：
	// profit, userShare, top, pool, uplinePortionBps, toL1, toL2, toTop, toProject
	if len(values) != 9 {
		return RewardNotifiedEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(values))
	}

	out.Profit = values[0].(*big.Int)
	out.UserShare = values[1].(*big.Int)
	out.Top = values[2].(common.Address)
	out.Pool = values[3].(*big.Int)
	out.UplinePortionBps = values[4].(*big.Int)
	out.ToL1 = values[5].(*big.Int)
	out.ToL2 = values[6].(*big.Int)
	out.ToTop = values[7].(*big.Int)
	out.ToProject = values[8].(*big.Int)

	return out, nil
}

func fillRewardNotifiedTimes(ctx context.Context, client *ethclient.Client, evs []RewardNotifiedEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}
		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

// FetchRewardNotifiedByRange 按 [fromBlock, toBlock]（包含边界）拉取 RewardNotified 事件，内部自动分段 QueryStep
func FetchRewardNotifiedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock uint64, toBlock uint64) ([]RewardNotifiedEvent, error) {

	if toBlock < fromBlock {
		return []RewardNotifiedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(umbrellaStakeDividendABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	eventID := parsedABI.Events["RewardNotified"].ID

	res := make([]RewardNotifiedEvent, 0, 64)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: []common.Address{contract},
			Topics:    [][]common.Hash{{eventID}},
		}

		logs, err := client.FilterLogs(ctx, q)
		if err != nil {
			return nil, fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
		}

		for _, lg := range logs {
			ev, err := parseRewardNotified(parsedABI, lg)
			if err != nil {
				return nil, fmt.Errorf("parse log tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
			}
			res = append(res, ev)
		}
	}

	// ✅ 统一补时间戳
	if err := fillRewardNotifiedTimes(ctx, client, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PollRewardNotifiedIncremental 你描述的“增量拉取 + 最新块往回 6”核心逻辑
// lastProcessedFromDB：DB 里记录的“已处理到的最高块”（建议含义：你上次成功跑完时的 safeTo）
func PollRewardNotifiedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []RewardNotifiedEvent, newLastProcessed uint64, err error) {

	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		// 链太新/异常情况
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	from := lastProcessedFromDB + 1
	if from < DeployBlock {
		from = DeployBlock
	}

	if safeTo < from {
		// 没有可确认的新块
		return []RewardNotifiedEvent{}, lastProcessedFromDB, nil
	}

	evs, err := FetchRewardNotifiedByRange(ctx, client, DividendContract, from, safeTo)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}

	// 只要你把 evs 全部落库成功（建议 txHash+logIndex 唯一键），再把 DB 的 lastProcessed 更新为 safeTo
	return evs, safeTo, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func BigIntToFloat64(x *big.Int, decimals int) float64 {
	if x == nil {
		return 0
	}
	r := new(big.Rat).SetInt(x)
	if decimals > 0 {
		r.Quo(r, new(big.Rat).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)))
	}
	f, _ := r.Float64()
	return f
}

const (
	// DeployBlockTwo TODO: 改成 SwapHrxUsdt 实际部署块高
	DeployBlockTwo uint64 = 72859368
)

var (
	// SwapContract TODO: 改成 SwapHrxUsdt 合约地址
	SwapContract = common.HexToAddress("0x3ff0acac62d1c5f74581d7a45b974bcce5f054e1")
)

const swapHrxUsdtABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "sender", "type": "address"},
      {"indexed": false, "internalType": "uint256", "name": "amount0In", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "amount1In", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "amount0OutGross", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "amount0OutNet", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "amount1OutGross", "type": "uint256"},
      {"indexed": true,  "internalType": "address", "name": "to", "type": "address"}
    ],
    "name": "Swap",
    "type": "event"
  }
]`

type SwapEvent struct {
	// 唯一定位（建议落库用 txHash+logIndex 去重）
	BlockNumber uint64
	TxHash      common.Hash
	LogIndex    uint

	// ✅ 秒级时间戳（区块级，同区块所有交易相同）
	BlockTime uint64

	// indexed
	Sender common.Address
	To     common.Address

	// data
	Amount0In       *big.Int
	Amount1In       *big.Int
	Amount0OutGross *big.Int
	Amount0OutNet   *big.Int
	Amount1OutGross *big.Int
}

func parseSwap(a abi.ABI, lg types.Log) (SwapEvent, error) {
	ev, ok := a.Events["Swap"]
	if !ok {
		return SwapEvent{}, fmt.Errorf("event Swap not found in ABI")
	}
	if len(lg.Topics) != 3 {
		return SwapEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return SwapEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := SwapEvent{
		BlockNumber: lg.BlockNumber,
		TxHash:      lg.TxHash,
		LogIndex:    lg.Index,

		Sender: common.BytesToAddress(lg.Topics[1].Bytes()),
		To:     common.BytesToAddress(lg.Topics[2].Bytes()),
	}

	vals, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return SwapEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	if len(vals) != 5 {
		return SwapEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(vals))
	}

	out.Amount0In = vals[0].(*big.Int)
	out.Amount1In = vals[1].(*big.Int)
	out.Amount0OutGross = vals[2].(*big.Int)
	out.Amount0OutNet = vals[3].(*big.Int)
	out.Amount1OutGross = vals[4].(*big.Int)
	return out, nil
}

// ✅ 批量填充 blockTime：缓存 blockNumber -> header.Time（秒）
func fillBlockTimes(ctx context.Context, client *ethclient.Client, evs []SwapEvent) error {
	cache := make(map[uint64]uint64, 256)

	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}
		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

// FetchSwapByRange 按 [fromBlock, toBlock] 拉取 Swap，并填充 BlockTime(秒)
func FetchSwapByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]SwapEvent, error) {
	if toBlock < fromBlock {
		return []SwapEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(swapHrxUsdtABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	swapID := parsedABI.Events["Swap"].ID

	res := make([]SwapEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: []common.Address{contract},
			Topics:    [][]common.Hash{{swapID}},
		}

		logs, err := client.FilterLogs(ctx, q)
		if err != nil {
			return nil, fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
		}

		for _, lg := range logs {
			ev, err := parseSwap(parsedABI, lg)
			if err != nil {
				return nil, fmt.Errorf("parse log tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
			}
			res = append(res, ev)
		}
	}

	// ✅ 统一补时间戳
	if err := fillBlockTimes(ctx, client, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PollSwapIncremental：增量拉取 + latest-6
func PollSwapIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []SwapEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	from := lastProcessedFromDB + 1
	if from < DeployBlockTwo {
		from = DeployBlockTwo
	}
	if safeTo < from {
		return []SwapEvent{}, lastProcessedFromDB, nil
	}

	evs, err := FetchSwapByRange(ctx, client, SwapContract, from, safeTo)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, safeTo, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DeployBlockThree uint64 = 72528236
	DeployBlockFour  uint64 = 72869154
)

var (
	// PrimaryMarketContract TODO: 改成 BondingCurvePrimaryMarket 合约地址
	PrimaryMarketContract = common.HexToAddress("0xe817f45dcd271dde3cac7d6ad9cb91de820c688a")
)

// 仅包含 Bought / Sold 事件（够用就行）
const primaryMarketABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "buyer", "type": "address"},
      {"indexed": true,  "internalType": "address", "name": "to", "type": "address"},
      {"indexed": false, "internalType": "uint256", "name": "usdtUsed", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "ausdGrossOut", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "ausdFee", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "ausdNetOut", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "priceBefore", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "priceAfter", "type": "uint256"}
    ],
    "name": "Bought",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "seller", "type": "address"},
      {"indexed": true,  "internalType": "address", "name": "to", "type": "address"},
      {"indexed": false, "internalType": "uint256", "name": "ausdGrossIn", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "ausdFee", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "ausdBurn", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "usdtOut", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "priceBefore", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "priceAfter", "type": "uint256"}
    ],
    "name": "Sold",
    "type": "event"
  }
]`

// -------------------- Bought --------------------

type BoughtEvent struct {
	BlockNumber uint64
	TxHash      common.Hash
	LogIndex    uint

	// ✅ 秒级时间戳（区块级）
	BlockTime uint64

	// indexed
	Buyer common.Address
	To    common.Address

	// data
	UsdtUsed     *big.Int
	AusdGrossOut *big.Int
	AusdFee      *big.Int
	AusdNetOut   *big.Int
	PriceBefore  *big.Int
	PriceAfter   *big.Int
}

func parseBought(a abi.ABI, lg types.Log) (BoughtEvent, error) {
	ev, ok := a.Events["Bought"]
	if !ok {
		return BoughtEvent{}, fmt.Errorf("event Bought not found in ABI")
	}
	// topics: [topic0=eventID, topic1=buyer, topic2=to]
	if len(lg.Topics) != 3 {
		return BoughtEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return BoughtEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := BoughtEvent{
		BlockNumber: lg.BlockNumber,
		TxHash:      lg.TxHash,
		LogIndex:    lg.Index,

		Buyer: common.BytesToAddress(lg.Topics[1].Bytes()),
		To:    common.BytesToAddress(lg.Topics[2].Bytes()),
	}

	// 非 indexed inputs 顺序：
	// usdtUsed, ausdGrossOut, ausdFee, ausdNetOut, priceBefore, priceAfter
	values, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return BoughtEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	if len(values) != 6 {
		return BoughtEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(values))
	}

	out.UsdtUsed = values[0].(*big.Int)
	out.AusdGrossOut = values[1].(*big.Int)
	out.AusdFee = values[2].(*big.Int)
	out.AusdNetOut = values[3].(*big.Int)
	out.PriceBefore = values[4].(*big.Int)
	out.PriceAfter = values[5].(*big.Int)

	return out, nil
}

// FetchBoughtByRange 按 [fromBlock, toBlock] 拉取 Bought，自动分段 QueryStep，并填 BlockTime
func FetchBoughtByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]BoughtEvent, error) {
	if toBlock < fromBlock {
		return []BoughtEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(primaryMarketABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	boughtID := parsedABI.Events["Bought"].ID

	res := make([]BoughtEvent, 0, 128)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: []common.Address{contract},
			Topics:    [][]common.Hash{{boughtID}},
		}

		logs, err := client.FilterLogs(ctx, q)
		if err != nil {
			return nil, fmt.Errorf("FilterLogs Bought [%d,%d]: %w", start, end, err)
		}

		for _, lg := range logs {
			ev, err := parseBought(parsedABI, lg)
			if err != nil {
				return nil, fmt.Errorf("parse Bought tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
			}
			res = append(res, ev)
		}
	}

	if err := fillBoughtTimes(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func fillBoughtTimes(ctx context.Context, client *ethclient.Client, evs []BoughtEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}
		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

// PollBoughtIncremental 增量拉取 + latest-6
func PollBoughtIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []BoughtEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	from := lastProcessedFromDB + 1
	if from < DeployBlockThree {
		from = DeployBlockThree
	}
	if safeTo < from {
		return []BoughtEvent{}, lastProcessedFromDB, nil
	}

	evs, err := FetchBoughtByRange(ctx, client, PrimaryMarketContract, from, safeTo)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, safeTo, nil
}

// -------------------- Sold --------------------

type SoldEvent struct {
	BlockNumber uint64
	TxHash      common.Hash
	LogIndex    uint

	// ✅ 秒级时间戳（区块级）
	BlockTime uint64

	// indexed
	Seller common.Address
	To     common.Address

	// data
	AusdGrossIn *big.Int
	AusdFee     *big.Int
	AusdBurn    *big.Int
	UsdtOut     *big.Int
	PriceBefore *big.Int
	PriceAfter  *big.Int
}

func parseSold(a abi.ABI, lg types.Log) (SoldEvent, error) {
	ev, ok := a.Events["Sold"]
	if !ok {
		return SoldEvent{}, fmt.Errorf("event Sold not found in ABI")
	}
	// topics: [topic0=eventID, topic1=seller, topic2=to]
	if len(lg.Topics) != 3 {
		return SoldEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return SoldEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := SoldEvent{
		BlockNumber: lg.BlockNumber,
		TxHash:      lg.TxHash,
		LogIndex:    lg.Index,

		Seller: common.BytesToAddress(lg.Topics[1].Bytes()),
		To:     common.BytesToAddress(lg.Topics[2].Bytes()),
	}

	// 非 indexed inputs 顺序：
	// ausdGrossIn, ausdFee, ausdBurn, usdtOut, priceBefore, priceAfter
	values, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return SoldEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	if len(values) != 6 {
		return SoldEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(values))
	}

	out.AusdGrossIn = values[0].(*big.Int)
	out.AusdFee = values[1].(*big.Int)
	out.AusdBurn = values[2].(*big.Int)
	out.UsdtOut = values[3].(*big.Int)
	out.PriceBefore = values[4].(*big.Int)
	out.PriceAfter = values[5].(*big.Int)

	return out, nil
}

// FetchSoldByRange 按 [fromBlock, toBlock] 拉取 Sold，自动分段 QueryStep，并填 BlockTime
func FetchSoldByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]SoldEvent, error) {
	if toBlock < fromBlock {
		return []SoldEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(primaryMarketABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	soldID := parsedABI.Events["Sold"].ID

	res := make([]SoldEvent, 0, 128)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: []common.Address{contract},
			Topics:    [][]common.Hash{{soldID}},
		}

		logs, err := client.FilterLogs(ctx, q)
		if err != nil {
			return nil, fmt.Errorf("FilterLogs Sold [%d,%d]: %w", start, end, err)
		}

		for _, lg := range logs {
			ev, err := parseSold(parsedABI, lg)
			if err != nil {
				return nil, fmt.Errorf("parse Sold tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
			}
			res = append(res, ev)
		}
	}

	if err := fillSoldTimes(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func fillSoldTimes(ctx context.Context, client *ethclient.Client, evs []SoldEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}
		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

// PollSoldIncremental 增量拉取 + latest-6
func PollSoldIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []SoldEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	from := lastProcessedFromDB + 1
	if from < DeployBlockFour {
		from = DeployBlockFour
	}
	if safeTo < from {
		return []SoldEvent{}, lastProcessedFromDB, nil
	}

	evs, err := FetchSoldByRange(ctx, client, PrimaryMarketContract, from, safeTo)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, safeTo, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	// DeployBlockNFT TODO: 改成 BlindBoxNFT 实际部署块高
	DeployBlockNFT uint64 = 72530801
)

var (
	// NftContract TODO: 改成 BlindBoxNFT 合约地址
	NftContract = common.HexToAddress("0xE447b0391d3F03befeC0dC09E25c049132618fd9")
)

const blindBoxPurchasedABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "buyer", "type": "address"},
      {"indexed": true,  "internalType": "address", "name": "seller", "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "priceUSDT", "type": "uint256"},
      {"indexed": false, "internalType": "bool",    "name": "feePaidInB", "type": "bool"},
      {"indexed": false, "internalType": "uint256", "name": "feeUSDT", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "feeB", "type": "uint256"}
    ],
    "name": "Purchased",
    "type": "event"
  }
]`

type PurchasedEvent struct {
	// 唯一定位（建议落库用 blockNumber+logIndex 去重）
	BlockNumber uint64
	LogIndex    uint

	// ✅ 秒级时间戳（区块级，同区块所有交易相同）
	BlockTime uint64

	// indexed
	Buyer   common.Address
	Seller  common.Address
	TokenID *big.Int

	// data（原样 bigint）
	PriceUSDT  *big.Int
	FeePaidInB bool
	FeeUSDT    *big.Int
	FeeB       *big.Int
}

func parsePurchased(a abi.ABI, lg types.Log) (PurchasedEvent, error) {
	ev, ok := a.Events["Purchased"]
	if !ok {
		return PurchasedEvent{}, fmt.Errorf("event Purchased not found in ABI")
	}
	// topics: [eventId, buyer, seller, tokenId]
	if len(lg.Topics) != 4 {
		return PurchasedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return PurchasedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := PurchasedEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,

		Buyer:   common.BytesToAddress(lg.Topics[1].Bytes()),
		Seller:  common.BytesToAddress(lg.Topics[2].Bytes()),
		TokenID: new(big.Int).SetBytes(lg.Topics[3].Bytes()),
	}

	vals, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return PurchasedEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	// data: priceUSDT, feePaidInB, feeUSDT, feeB
	if len(vals) != 4 {
		return PurchasedEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(vals))
	}

	price, ok := vals[0].(*big.Int)
	if !ok {
		return PurchasedEvent{}, fmt.Errorf("bad type priceUSDT: %T", vals[0])
	}
	paidInB, ok := vals[1].(bool)
	if !ok {
		return PurchasedEvent{}, fmt.Errorf("bad type feePaidInB: %T", vals[1])
	}
	feeU, ok := vals[2].(*big.Int)
	if !ok {
		return PurchasedEvent{}, fmt.Errorf("bad type feeUSDT: %T", vals[2])
	}
	feeB, ok := vals[3].(*big.Int)
	if !ok {
		return PurchasedEvent{}, fmt.Errorf("bad type feeB: %T", vals[3])
	}

	out.PriceUSDT = price
	out.FeePaidInB = paidInB
	out.FeeUSDT = feeU
	out.FeeB = feeB

	return out, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetryPurchased(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetryPurchased(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 批量填充 blockTime：缓存 blockNumber -> header.Time（秒）
func fillBlockTimesPurchased(ctx context.Context, client *ethclient.Client, evs []PurchasedEvent) error {
	cache := make(map[uint64]uint64, 256)

	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetryPurchased(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

// FetchPurchasedByRange 按 [fromBlock, toBlock] 拉取 Purchased，并填充 BlockTime(秒)
func FetchPurchasedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]PurchasedEvent, error) {
	if toBlock < fromBlock {
		return []PurchasedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(blindBoxPurchasedABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	purchasedID := parsedABI.Events["Purchased"].ID

	res := make([]PurchasedEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: []common.Address{contract},
			Topics:    [][]common.Hash{{purchasedID}},
		}

		// ✅ 失败不直接 return：最多重试 5 次
		logs, err := filterLogsRetryPurchased(ctx, client, q)
		if err != nil {
			return nil, fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
		}

		for _, lg := range logs {
			ev, err := parsePurchased(parsedABI, lg)
			if err != nil {
				return nil, fmt.Errorf("parse log tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
			}
			res = append(res, ev)
		}
	}

	// ✅ 统一补时间戳（HeaderByNumber 也最多重试 5 次）
	if err := fillBlockTimesPurchased(ctx, client, res); err != nil {
		return nil, err
	}

	return res, nil
}

// PollPurchasedIncremental：增量拉取 + latest-Confirmations
func PollPurchasedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []PurchasedEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	from := lastProcessedFromDB + 1
	if from < DeployBlockNFT {
		from = DeployBlockNFT
	}
	if safeTo < from {
		return []PurchasedEvent{}, lastProcessedFromDB, nil
	}

	evs, err := FetchPurchasedByRange(ctx, client, NftContract, from, safeTo)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, safeTo, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DeployBlockNFTListed uint64 = 73402672 // TODO: 改成你想开始监听的块高
)

const blindBoxListedABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "seller", "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "timestamp", "type": "uint256"}
    ],
    "name": "Listed",
    "type": "event"
  }
]`

type ListedEvent struct {
	BlockNumber uint64
	LogIndex    uint
	BlockTime   uint64 // 秒级（区块时间）

	// indexed
	Seller  common.Address
	TokenID *big.Int

	// data
	Timestamp *big.Int
}

func parseListed(a abi.ABI, lg types.Log) (ListedEvent, error) {
	ev, ok := a.Events["Listed"]
	if !ok {
		return ListedEvent{}, fmt.Errorf("event Listed not found in ABI")
	}
	// topics: [eventId, seller, tokenId]
	if len(lg.Topics) != 3 {
		return ListedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return ListedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := ListedEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,
		Seller:      common.BytesToAddress(lg.Topics[1].Bytes()),
		TokenID:     new(big.Int).SetBytes(lg.Topics[2].Bytes()),
	}

	vals, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return ListedEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	// data: timestamp
	if len(vals) != 1 {
		return ListedEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(vals))
	}

	ts, ok := vals[0].(*big.Int)
	if !ok {
		return ListedEvent{}, fmt.Errorf("bad type timestamp: %T", vals[0])
	}
	out.Timestamp = ts

	return out, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetryListed(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetryListed(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func fillBlockTimesListed(ctx context.Context, client *ethclient.Client, evs []ListedEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetryListed(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

func FetchListedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]ListedEvent, error) {
	if toBlock < fromBlock {
		return []ListedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(blindBoxListedABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	listedID := parsedABI.Events["Listed"].ID

	res := make([]ListedEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: []common.Address{contract},
			Topics:    [][]common.Hash{{listedID}},
		}

		// ✅ 失败不直接 return：最多重试 5 次
		logs, err := filterLogsRetryListed(ctx, client, q)
		if err != nil {
			return nil, fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
		}

		for _, lg := range logs {
			ev, err := parseListed(parsedABI, lg)
			if err != nil {
				return nil, fmt.Errorf("parse log tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
			}
			res = append(res, ev)
		}
	}

	// ✅ HeaderByNumber 也最多重试 5 次
	if err := fillBlockTimesListed(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func PollListedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []ListedEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	from := lastProcessedFromDB + 1
	if from < DeployBlockNFTListed {
		from = DeployBlockNFTListed
	}
	if safeTo < from {
		return []ListedEvent{}, lastProcessedFromDB, nil
	}

	evs, err := FetchListedByRange(ctx, client, NftContract, from, safeTo)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, safeTo, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DeployBlockNFTUnlisted uint64 = 73530659 // TODO: 改成你想开始监听的块高
)

const blindBoxUnlistedABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "operator", "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"}
    ],
    "name": "Unlisted",
    "type": "event"
  }
]`

type UnlistedEvent struct {
	BlockNumber uint64
	LogIndex    uint
	BlockTime   uint64 // 秒级

	// indexed
	Operator common.Address
	TokenID  *big.Int
}

func parseUnlisted(a abi.ABI, lg types.Log) (UnlistedEvent, error) {
	ev, ok := a.Events["Unlisted"]
	if !ok {
		return UnlistedEvent{}, fmt.Errorf("event Unlisted not found in ABI")
	}
	// topics: [eventId, operator, tokenId]
	if len(lg.Topics) != 3 {
		return UnlistedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return UnlistedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	return UnlistedEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,
		Operator:    common.BytesToAddress(lg.Topics[1].Bytes()),
		TokenID:     new(big.Int).SetBytes(lg.Topics[2].Bytes()),
	}, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetryUnlisted(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetryUnlisted(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func fillBlockTimesUnlisted(ctx context.Context, client *ethclient.Client, evs []UnlistedEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetryUnlisted(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

func FetchUnlistedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]UnlistedEvent, error) {
	if toBlock < fromBlock {
		return []UnlistedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(blindBoxUnlistedABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	unlistedID := parsedABI.Events["Unlisted"].ID

	res := make([]UnlistedEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: []common.Address{contract},
			Topics:    [][]common.Hash{{unlistedID}},
		}

		// ✅ 失败不直接 return：最多重试 5 次
		logs, err := filterLogsRetryUnlisted(ctx, client, q)
		if err != nil {
			return nil, fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
		}

		for _, lg := range logs {
			ev, err := parseUnlisted(parsedABI, lg)
			if err != nil {
				return nil, fmt.Errorf("parse log tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
			}
			res = append(res, ev)
		}
	}

	// ✅ HeaderByNumber 也最多重试 5 次
	if err := fillBlockTimesUnlisted(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func PollUnlistedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []UnlistedEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	from := lastProcessedFromDB + 1
	if from < DeployBlockNFTUnlisted {
		from = DeployBlockNFTUnlisted
	}
	if safeTo < from {
		return []UnlistedEvent{}, lastProcessedFromDB, nil
	}

	evs, err := FetchUnlistedByRange(ctx, client, NftContract, from, safeTo)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, safeTo, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

/* =========================
   Minted (铸造) 事件
   ========================= */

const (
	DeployBlockNFTMint uint64 = 72539344
)

const blindBoxMintedABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "to", "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"},
      {"indexed": true,  "internalType": "uint8",   "name": "tier", "type": "uint8"},
      {"indexed": false, "internalType": "uint256", "name": "usdtPaid", "type": "uint256"}
    ],
    "name": "Minted",
    "type": "event"
  }
]`

type MintedEvent struct {
	BlockNumber uint64
	LogIndex    uint
	BlockTime   uint64 // 秒级

	// indexed
	To      common.Address
	TokenID *big.Int

	// data
	Tier     uint8
	UsdtPaid *big.Int
}

func parseMinted(a abi.ABI, lg types.Log) (MintedEvent, error) {
	ev, ok := a.Events["Minted"]
	if !ok {
		return MintedEvent{}, fmt.Errorf("event Minted not found in ABI")
	}
	// topics: [eventId, to, tokenId, tier]
	if len(lg.Topics) != 4 {
		return MintedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return MintedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := MintedEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,

		To:      common.BytesToAddress(lg.Topics[1].Bytes()),
		TokenID: new(big.Int).SetBytes(lg.Topics[2].Bytes()),
	}

	// tier 是 indexed，topic 里是 32 bytes uint
	tierBI := new(big.Int).SetBytes(lg.Topics[3].Bytes())
	if tierBI.Sign() < 0 || tierBI.BitLen() > 8 {
		return MintedEvent{}, fmt.Errorf("tier out of uint8: %s", tierBI.String())
	}
	out.Tier = uint8(tierBI.Uint64())

	// 非 indexed data：usdtPaid
	vals, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return MintedEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	if len(vals) != 1 {
		return MintedEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(vals))
	}

	usdtPaid, ok := vals[0].(*big.Int)
	if !ok {
		return MintedEvent{}, fmt.Errorf("bad type usdtPaid: %T", vals[0])
	}
	out.UsdtPaid = usdtPaid

	return out, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetryMinted(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetryMinted(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func fillBlockTimesMinted(ctx context.Context, client *ethclient.Client, evs []MintedEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetryMinted(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

// FetchMintedByRange 按 [fromBlock, toBlock] 拉取 Minted，并填充 BlockTime(秒)
func FetchMintedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]MintedEvent, error) {
	if toBlock < fromBlock {
		return []MintedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(blindBoxMintedABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	mintedID := parsedABI.Events["Minted"].ID

	res := make([]MintedEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: []common.Address{contract},
			Topics:    [][]common.Hash{{mintedID}},
		}

		// ✅ 失败不直接 return：最多重试 5 次
		logs, err := filterLogsRetryMinted(ctx, client, q)
		if err != nil {
			return nil, fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
		}

		for _, lg := range logs {
			ev, err := parseMinted(parsedABI, lg)
			if err != nil {
				return nil, fmt.Errorf("parse log tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
			}
			res = append(res, ev)
		}
	}

	// ✅ HeaderByNumber 也最多重试 5 次
	if err := fillBlockTimesMinted(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

// PollMintedIncremental：增量拉取 + latest-Confirmations
func PollMintedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []MintedEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	from := lastProcessedFromDB + 1
	if from < DeployBlockNFTMint {
		from = DeployBlockNFTMint
	}
	if safeTo < from {
		return []MintedEvent{}, lastProcessedFromDB, nil
	}

	evs, err := FetchMintedByRange(ctx, client, NftContract, from, safeTo)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, safeTo, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DeployBlockNFTOpened uint64 = 73406511 // TODO: 按你的实际部署/开启监听块高改
)

const blindBoxOpenedABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "owner", "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "timestamp", "type": "uint256"},
      {"indexed": false, "internalType": "uint256", "name": "bnbPaid", "type": "uint256"}
    ],
    "name": "Opened",
    "type": "event"
  }
]`

type OpenedEvent struct {
	BlockNumber uint64
	LogIndex    uint
	BlockTime   uint64 // 秒级（区块时间）

	// indexed
	Owner   common.Address
	TokenID *big.Int

	// data
	Timestamp *big.Int
	BnbPaid   *big.Int
}

func parseOpened(a abi.ABI, lg types.Log) (OpenedEvent, error) {
	ev, ok := a.Events["Opened"]
	if !ok {
		return OpenedEvent{}, fmt.Errorf("event Opened not found in ABI")
	}
	// topics: [eventId, owner, tokenId]
	if len(lg.Topics) != 3 {
		return OpenedEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return OpenedEvent{}, fmt.Errorf("topic0 mismatch")
	}

	out := OpenedEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,

		Owner:   common.BytesToAddress(lg.Topics[1].Bytes()),
		TokenID: new(big.Int).SetBytes(lg.Topics[2].Bytes()),
	}

	vals, err := ev.Inputs.NonIndexed().Unpack(lg.Data)
	if err != nil {
		return OpenedEvent{}, fmt.Errorf("unpack data: %w", err)
	}
	// data: timestamp, bnbPaid
	if len(vals) != 2 {
		return OpenedEvent{}, fmt.Errorf("unexpected decoded values len=%d", len(vals))
	}

	ts, ok := vals[0].(*big.Int)
	if !ok {
		return OpenedEvent{}, fmt.Errorf("bad type timestamp: %T", vals[0])
	}
	bnbPaid, ok := vals[1].(*big.Int)
	if !ok {
		return OpenedEvent{}, fmt.Errorf("bad type bnbPaid: %T", vals[1])
	}

	out.Timestamp = ts
	out.BnbPaid = bnbPaid
	return out, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetryOpened(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetryOpened(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func fillBlockTimesOpened(ctx context.Context, client *ethclient.Client, evs []OpenedEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetryOpened(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

func FetchOpenedByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]OpenedEvent, error) {
	if toBlock < fromBlock {
		return []OpenedEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(blindBoxOpenedABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	openedID := parsedABI.Events["Opened"].ID

	res := make([]OpenedEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: []common.Address{contract},
			Topics:    [][]common.Hash{{openedID}},
		}

		// ✅ 失败不直接 return：最多重试 5 次
		logs, err := filterLogsRetryOpened(ctx, client, q)
		if err != nil {
			return nil, fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
		}

		for _, lg := range logs {
			ev, err := parseOpened(parsedABI, lg)
			if err != nil {
				return nil, fmt.Errorf("parse log tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
			}
			res = append(res, ev)
		}
	}

	// ✅ HeaderByNumber 也最多重试 5 次
	if err := fillBlockTimesOpened(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func PollOpenedIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []OpenedEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	from := lastProcessedFromDB + 1
	if from < DeployBlockNFTOpened {
		from = DeployBlockNFTOpened
	}
	if safeTo < from {
		return []OpenedEvent{}, lastProcessedFromDB, nil
	}

	evs, err := FetchOpenedByRange(ctx, client, NftContract, from, safeTo)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, safeTo, nil
}

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	DeployBlockNFTTransfer uint64 = 72530801 // TODO: 按你的实际部署/开启监听块高改
)

const erc721TransferABI = `[
  {
    "anonymous": false,
    "inputs": [
      {"indexed": true,  "internalType": "address", "name": "from", "type": "address"},
      {"indexed": true,  "internalType": "address", "name": "to",   "type": "address"},
      {"indexed": true,  "internalType": "uint256", "name": "tokenId", "type": "uint256"}
    ],
    "name": "Transfer",
    "type": "event"
  }
]`

type TransferEvent struct {
	BlockNumber uint64
	LogIndex    uint
	BlockTime   uint64 // 秒级（区块时间）

	// indexed
	From    common.Address
	To      common.Address
	TokenID *big.Int
}

func parseTransfer(a abi.ABI, lg types.Log) (TransferEvent, error) {
	ev, ok := a.Events["Transfer"]
	if !ok {
		return TransferEvent{}, fmt.Errorf("event Transfer not found in ABI")
	}
	// topics: [eventId, from, to, tokenId]
	if len(lg.Topics) != 4 {
		return TransferEvent{}, fmt.Errorf("bad topics len=%d", len(lg.Topics))
	}
	if lg.Topics[0] != ev.ID {
		return TransferEvent{}, fmt.Errorf("topic0 mismatch")
	}

	return TransferEvent{
		BlockNumber: lg.BlockNumber,
		LogIndex:    lg.Index,

		From:    common.BytesToAddress(lg.Topics[1].Bytes()),
		To:      common.BytesToAddress(lg.Topics[2].Bytes()),
		TokenID: new(big.Int).SetBytes(lg.Topics[3].Bytes()),
	}, nil
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func filterLogsRetry(ctx context.Context, client *ethclient.Client, q ethereum.FilterQuery) ([]types.Log, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		logs, err := client.FilterLogs(ctx, q)
		if err == nil {
			return logs, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

// ✅ 简单重试：最多 5 次（局部变量，不放全局）
func headerByNumberRetry(ctx context.Context, client *ethclient.Client, bn uint64) (*types.Header, error) {
	maxTry := 5
	delay := 400 * time.Millisecond

	var lastErr error
	for i := 0; i < maxTry; i++ {
		h, err := client.HeaderByNumber(ctx, new(big.Int).SetUint64(bn))
		if err == nil {
			return h, nil
		}
		lastErr = err
		if i != maxTry-1 {
			time.Sleep(delay)
		}
	}
	return nil, lastErr
}

func fillBlockTimesTransfer(ctx context.Context, client *ethclient.Client, evs []TransferEvent) error {
	cache := make(map[uint64]uint64, 256)
	for i := range evs {
		bn := evs[i].BlockNumber
		if ts, ok := cache[bn]; ok {
			evs[i].BlockTime = ts
			continue
		}

		h, err := headerByNumberRetry(ctx, client, bn)
		if err != nil {
			return fmt.Errorf("HeaderByNumber(%d): %w", bn, err)
		}

		cache[bn] = h.Time
		evs[i].BlockTime = h.Time
	}
	return nil
}

func FetchTransferByRange(ctx context.Context, client *ethclient.Client, contract common.Address, fromBlock, toBlock uint64) ([]TransferEvent, error) {
	if toBlock < fromBlock {
		return []TransferEvent{}, nil
	}

	parsedABI, err := abi.JSON(strings.NewReader(erc721TransferABI))
	if err != nil {
		return nil, fmt.Errorf("parse abi: %w", err)
	}
	transferID := parsedABI.Events["Transfer"].ID

	res := make([]TransferEvent, 0, 256)

	for start := fromBlock; start <= toBlock; start += QueryStep {
		end := start + QueryStep - 1
		if end > toBlock {
			end = toBlock
		}

		q := ethereum.FilterQuery{
			FromBlock: new(big.Int).SetUint64(start),
			ToBlock:   new(big.Int).SetUint64(end),
			Addresses: []common.Address{contract},
			Topics:    [][]common.Hash{{transferID}},
		}

		// ✅ 失败不直接 return：最多重试 5 次
		logs, err := filterLogsRetry(ctx, client, q)
		if err != nil {
			return nil, fmt.Errorf("FilterLogs [%d,%d]: %w", start, end, err)
		}

		for _, lg := range logs {
			ev, err := parseTransfer(parsedABI, lg)
			if err != nil {
				return nil, fmt.Errorf("parse log tx=%s idx=%d: %w", lg.TxHash.Hex(), lg.Index, err)
			}
			res = append(res, ev)
		}
	}

	// ✅ HeaderByNumber 也最多重试 5 次
	if err := fillBlockTimesTransfer(ctx, client, res); err != nil {
		return nil, err
	}
	return res, nil
}

func PollTransferIncremental(ctx context.Context, client *ethclient.Client, lastProcessedFromDB uint64) (events []TransferEvent, newLastProcessed uint64, err error) {
	head, err := client.BlockNumber(ctx)
	if err != nil {
		return nil, lastProcessedFromDB, fmt.Errorf("BlockNumber: %w", err)
	}
	if head <= Confirmations {
		return nil, lastProcessedFromDB, nil
	}
	safeTo := head - Confirmations

	from := lastProcessedFromDB + 1
	if from < DeployBlockNFTTransfer {
		from = DeployBlockNFTTransfer
	}
	if safeTo < from {
		return []TransferEvent{}, lastProcessedFromDB, nil
	}

	evs, err := FetchTransferByRange(ctx, client, NftContract, from, safeTo)
	if err != nil {
		return nil, lastProcessedFromDB, err
	}
	return evs, safeTo, nil
}
