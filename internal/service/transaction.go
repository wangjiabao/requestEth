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
	tmpFour := "-1"
	tmpFive := "-1"
	tmpSix := "-1"
	tmpOneLength := "-1"
	tmpTwoLength := "-1"
	tmpThreeLength := "-1"
	tmpFourLength := "-1"

	for _, urlTmp := range urls {
		client, err := ethclient.Dial(urlTmp)
		if err != nil {
			fmt.Println("client error:", err)
			continue
		}

		tokenAddress := common.HexToAddress("0x7d3482934977EE703F9D7B14b6D158691AacBae7")
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("NewPair error:", err)
			continue
		}

		// 获取储备量
		if "-1" == tmpOne {
			one, errOne := instance.One(&bind.CallOpts{}, common.HexToAddress(req.Address))
			if errOne != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpOne = one.String()
		}

		if "-1" == tmpTwo {
			two, errTwo := instance.Two(&bind.CallOpts{}, common.HexToAddress(req.Address))
			if errTwo != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpTwo = two.String()
		}

		if "-1" == tmpThree {
			three, errThree := instance.Three(&bind.CallOpts{}, common.HexToAddress(req.Address))
			if errThree != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpThree = three.String()
		}

		if "-1" == tmpFour {
			four, errFour := instance.Four(&bind.CallOpts{}, common.HexToAddress(req.Address))
			if errFour != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpFour = four.String()
		}

		if "-1" == tmpFive {
			five, errFive := instance.LpAmount(&bind.CallOpts{}, common.HexToAddress(req.Address))
			if errFive != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpFive = five.String()
		}

		if "-1" == tmpSix {
			six, errSix := instance.LpAmountTotal(&bind.CallOpts{})
			if errSix != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpSix = six.String()
		}

		if "-1" == tmpOneLength {
			six, errSix := instance.GetOneArrayLength(&bind.CallOpts{}, common.HexToAddress(req.Address))
			if errSix != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpOneLength = six.String()
		}

		if "-1" == tmpTwoLength {
			six, errSix := instance.GetTwoArrayLength(&bind.CallOpts{}, common.HexToAddress(req.Address))
			if errSix != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpTwoLength = six.String()
		}

		if "-1" == tmpThreeLength {
			six, errSix := instance.GetThreeArrayLength(&bind.CallOpts{}, common.HexToAddress(req.Address))
			if errSix != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpThreeLength = six.String()
		}

		if "-1" == tmpFourLength {
			six, errSix := instance.GetFourArrayLength(&bind.CallOpts{}, common.HexToAddress(req.Address))
			if errSix != nil {
				fmt.Println("GetAll error:", err)
				continue
			}

			tmpFourLength = six.String()
		}

		if "-1" != tmpOne && "-1" != tmpTwo && "-1" != tmpThree && "-1" != tmpFour &&
			"-1" != tmpFive && "-1" != tmpSix &&
			"-1" != tmpOneLength && "-1" != tmpTwoLength && "-1" != tmpThreeLength && "-1" != tmpFourLength {
			break
		}
	}

	return &pb.GetAllReply{
		One:           tmpOne,
		Two:           tmpTwo,
		Three:         tmpThree,
		Four:          tmpFour,
		LpAmount:      tmpFive,
		LpAmountTotal: tmpSix,
		OneLength:     tmpOneLength,
		ThreeLength:   tmpThreeLength,
		TwoLength:     tmpTwoLength,
		FourLength:    tmpFourLength,
	}, nil
}

func (s *TransactionService) PushOne(ctx context.Context, req *pb.PushOneRequest) (*pb.PushOneReply, error) {
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

		tokenAddress := common.HexToAddress("0x7d3482934977EE703F9D7B14b6D158691AacBae7")
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("PushOne error:", err)
			continue
		}

		var (
			tx *types.Transaction
		)
		address := make([]common.Address, 0)
		one := make([]*big.Int, 0)
		for _, v := range req.SendBody.Address {
			address = append(address, common.HexToAddress(v.Address))
			tmpAmount, _ := new(big.Int).SetString(v.One, 10)
			if nil == tmpAmount {
				return &pb.PushOneReply{
					Res: hashContent,
				}, nil
			}
			one = append(one, tmpAmount)
		}

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

		tx, err = instance.SetOneTwo(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 0,
		}, address, one)
		if err != nil {
			fmt.Println("PushOne error:", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.PushOneReply{
				Res: hashContent,
			}, nil
		}

		return &pb.PushOneReply{
			Res: tx.Hash().Hex(),
		}, nil
	}

	return &pb.PushOneReply{
		Res: hashContent,
	}, nil
}

func (s *TransactionService) PushThreeFour(ctx context.Context, req *pb.PushThreeFourRequest) (*pb.PushThreeFourReply, error) {
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

		tokenAddress := common.HexToAddress("0x7d3482934977EE703F9D7B14b6D158691AacBae7")
		instance, err := NewAdmin(tokenAddress, client)
		if err != nil {
			fmt.Println("NewPair error:", err)
			continue
		}

		var (
			tx *types.Transaction
		)
		address := make([]common.Address, 0)
		three := make([]*big.Int, 0)
		four := make([]*big.Int, 0)
		for _, v := range req.SendBody.Address {
			address = append(address, common.HexToAddress(v.Address))
			tmpAmount, _ := new(big.Int).SetString(v.Three, 10)
			tmpAmountTwo, _ := new(big.Int).SetString(v.Four, 10)
			three = append(three, tmpAmount)
			four = append(four, tmpAmountTwo)
		}

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

		tx, err = instance.SetThreeFour(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 0,
		}, address, three, four)
		if err != nil {
			fmt.Println("PushThreeFour error:", err)
			continue
		}

		if 0 >= len(tx.Hash().Hex()) {
			return &pb.PushThreeFourReply{
				Res: "-1",
			}, nil
		}

	}

	return &pb.PushThreeFourReply{
		Res: "ok",
	}, nil
}
