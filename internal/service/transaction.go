package service

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"os/exec"
	"strconv"
	"strings"
	"sync"
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
			fmt.Println("err info", req.SendBody.ToAddr, req.SendBody.PrivateKey, req.SendBody.Amount, tmpUrl1)
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

	sdk := newBCFWalletSDK()
	if nil == sdk.nodeProcess {
		return nil, nil
	}

	wallet := sdk.newBCFWallet("34.84.178.63", 19503, "https://qapmapi.pmchainbox.com/browser")
	balance := wallet.getAddressBalance("cEAXDkaEJgWKMM61KYz2dYU1RfuxbB8Ma", "XXVXQ", "PMC")
	fmt.Println("balance=", balance)
	defer sdk.Close()

	return nil, nil

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

type Result struct {
	Code    int    // 0 fail 1 success
	Message string //
}

type NodeProcess struct {
	ChannelMap sync.Map
	Cmd        *exec.Cmd
	Stdin      io.WriteCloser
	Stdout     io.ReadCloser
	Stderr     io.ReadCloser
	//NodeExec    func()
}

func newNodeProcess(cmd string, args ...string) *NodeProcess {
	command := exec.Command(cmd, args...)

	stdin, err := command.StdinPipe()
	if err != nil {
		//panic(err)
		fmt.Println(err)
		return nil
	}

	stdout, err := command.StdoutPipe()
	if err != nil {
		//panic(err)
		fmt.Println(err)
		return nil
	}

	stderr, err := command.StderrPipe()
	if err != nil {
		//panic(err)
		fmt.Println(err)
		return nil
	}
	nodeProcess := &NodeProcess{
		ChannelMap: sync.Map{},
		Cmd:        command,
		Stdin:      stdin,
		Stdout:     stdout,
		Stderr:     stderr,
	}

	err = nodeProcess.Cmd.Start()
	if err != nil {
		fmt.Println("Failed to start process:", err)
		return nil
	}
	//if err := cmd.Start(); err != nil {
	//	fmt.Println("Error starting command:", err)
	//	return
	//}
	// 在一个goroutine中读取stdout
	var readLines = func(name string, reader io.Reader, resultCode int) {
		scanner := bufio.NewReader(reader)
		for {
			line, err := scanner.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				fmt.Println(name+" read line error:", err)
				continue
			}
			//fmt.Printf("line: %v\n", line)
			parts := strings.SplitN(line, "Result ", 2)
			if len(parts) == 2 {
				//fmt.Printf("node: %v\n", parts[1])
				// 解析行并提取req_id和json_result
				parts := strings.SplitN(parts[1], " ", 3)
				if len(parts) < 2 {
					fmt.Println(name+" Invalid line format:", line)
					continue
				}
				reqIDStr := parts[0]
				resultData := parts[1]

				// 将req_id转换为整数
				reqID, err := strconv.Atoi(reqIDStr)
				if err != nil {
					fmt.Println(name+" Invalid req_id:", reqIDStr)
					continue
				}

				// 从map中获取并移除通道
				//if ch, ok := channelMap.LoadAndDelete(reqID); ok {
				if ch, ok := nodeProcess.ChannelMap.LoadAndDelete(reqID); ok {
					channel := ch.(chan Result)
					// 向通道发送结果
					channel <- Result{Code: resultCode, Message: resultData}
				} else {
					fmt.Println(name+" Channel not found for req_id:", reqID)
				}
			}
		}
		fmt.Println("node-process end")
	}
	go readLines("stdout", nodeProcess.Stdout, 1)
	go readLines("stderr", nodeProcess.Stderr, 0)

	return nodeProcess
}
func (p *NodeProcess) CloseProcess() error {
	p.Stdin.Close()
	return p.Cmd.Wait()
}

type BCFWalletSDK struct {
	nodeProcess *NodeProcess
}

func (sdk *BCFWalletSDK) Close() {
	sdk.nodeProcess.CloseProcess()
}

func newBCFWalletSDK() BCFWalletSDK {
	//init 结构体
	// 启动Node.js进程
	nodeProcess := newNodeProcess("node", "--no-warnings", "./sdk.js")
	return BCFWalletSDK{nodeProcess: nodeProcess}
}

var reqIdAcc = 0

// func NodeExec[T any](jsCode string) (T, error) {
func nodeExec[T any](nodeProcess *NodeProcess, jsCode string) (T, error) {
	var res T
	reqIdAcc += 1
	req_id := reqIdAcc
	channel := make(chan Result)
	nodeProcess.ChannelMap.Store(req_id, channel)

	var evalCode = fmt.Sprintf("await returnToGo(%d, async()=>%v)\r\n\n", req_id, jsCode)
	_, err := nodeProcess.Stdin.Write([]byte(evalCode))
	if err != nil {
		return res, err
	}

	result := <-channel
	if result.Code == 1 {
		err := json.Unmarshal([]byte(result.Message), &res)
		return res, err
	} else {
		return res, errors.New(result.Message)
	}
}

type BCFWallet struct {
	nodeProcess *NodeProcess
	walletId    string
}

func (sdk BCFWalletSDK) newBCFWallet(ip string, port int, browserPath string) *BCFWallet {
	//TODO
	bfcWalletId, _ := nodeExec[int](sdk.nodeProcess, `{
		const bfcwalletMap = (globalThis.bfcwalletMap??=new Map());
		globalThis.bfcwalletIdAcc ??= 0
		const id = globalThis.bfcwalletIdAcc++
		const bfcwallet = new require('@bfmeta/wallet-bcf').BCFWalletFactory({
			enable: true,
            host: [{ ip: "`+ip+`", port: `+strconv.Itoa(port)+` }],
            browserPath: "`+browserPath+`",
        });
		bfcwalletMap.set(id, bfcwallet)
		return id;
		}`)
	return &BCFWallet{nodeProcess: sdk.nodeProcess, walletId: strconv.Itoa(bfcWalletId)}
}

type BalanceResult struct {
	Success bool `json:"success"`
	Result  struct {
		Amount string `json:"amount"`
	} `json:"result"`
}

func (wallet *BCFWallet) getAddressBalance(address string, magic string, assetType string) BalanceResult {
	//balance, _ := nodeExec[string](`
	balance, _ := nodeExec[BalanceResult](wallet.nodeProcess, `
		globalThis.bfcwalletMap.get(`+wallet.walletId+`).getAddressBalance("`+address+`", "`+magic+`", "`+assetType+`")
	`)

	return balance
}
