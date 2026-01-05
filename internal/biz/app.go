package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "requestEth/api/requestEth/v1"
	"time"
)

type SwapTrade struct {
	ID              uint64
	BlockNumber     uint64
	LogIndex        uint32
	BlockTime       uint64
	Sender          string
	ToAddr          string
	Side            uint8
	Amount0In       float64
	Amount1In       float64
	Amount0OutNet   float64
	Amount1OutGross float64
	Amount0OutGross float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type PrimaryBuy struct {
	ID           uint64
	BlockNumber  uint64
	BlockTime    uint64
	LogIndex     uint32
	Buyer        string
	ToAddr       string
	UsdtUsed     float64
	AusdGrossOut float64
	AusdFee      float64
	AusdNetOut   float64
	PriceBefore  float64
	PriceAfter   float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PrimarySell struct {
	ID          uint64
	BlockNumber uint64
	BlockTime   uint64
	LogIndex    uint32
	Seller      string
	ToAddr      string
	AusdGrossIn float64
	AusdFee     float64
	AusdBurn    float64
	UsdtOut     float64
	PriceBefore float64
	PriceAfter  float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (PrimaryBuy) TableName() string { return "primary_buy" }

type UserRepo interface {
	GetSwapTradeLast(ctx context.Context) (*SwapTrade, error)
	GetSwapTrade(ctx context.Context, start, end uint64) ([]*SwapTrade, error)
	InsertSwapTrade(ctx context.Context, iData *SwapTrade) error
	GetPrimaryBuyLast(ctx context.Context) (*PrimaryBuy, error)
	GetPrimaryBuy(ctx context.Context, start, end uint64) ([]*PrimaryBuy, error)
	InsertPrimaryBuy(ctx context.Context, iData *PrimaryBuy) error
	GetPrimarySellLast(ctx context.Context) (*PrimarySell, error)
	GetPrimarySell(ctx context.Context, start, end uint64) ([]*PrimarySell, error)
	InsertPrimarySell(ctx context.Context, iData *PrimarySell) error
}

// AppUsecase is an app usecase.
type AppUsecase struct {
	userRepo UserRepo
	tx       Transaction
	log      *log.Helper
}

// NewAppUsecase new a app usecase.
func NewAppUsecase(userRepo UserRepo, tx Transaction, logger log.Logger) *AppUsecase {
	return &AppUsecase{userRepo: userRepo, tx: tx, log: log.NewHelper(logger)}
}

func (ac *AppUsecase) GetSwapTradeLast(ctx context.Context) (*SwapTrade, error) {
	var (
		rLast *SwapTrade
		err   error
	)
	rLast, err = ac.userRepo.GetSwapTradeLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertSwapTrade(ctx context.Context, trade *SwapTrade) error {
	var (
		err error
	)

	err = ac.userRepo.InsertSwapTrade(ctx, trade)

	return err
}

func (ac *AppUsecase) GetBuyLast(ctx context.Context) (*PrimaryBuy, error) {
	var (
		rLast *PrimaryBuy
		err   error
	)
	rLast, err = ac.userRepo.GetPrimaryBuyLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertBuyTrade(ctx context.Context, trade *PrimaryBuy) error {
	var (
		err error
	)

	err = ac.userRepo.InsertPrimaryBuy(ctx, trade)

	return err
}

func (ac *AppUsecase) GetSellLast(ctx context.Context) (*PrimarySell, error) {
	var (
		rLast *PrimarySell
		err   error
	)
	rLast, err = ac.userRepo.GetPrimarySellLast(ctx)
	if nil != err || nil == rLast {
		return nil, err
	}

	return rLast, nil
}

func (ac *AppUsecase) InsertSellTrade(ctx context.Context, trade *PrimarySell) error {
	var (
		err error
	)

	err = ac.userRepo.InsertPrimarySell(ctx, trade)

	return err
}

func (ac *AppUsecase) GetExchangeList(ctx context.Context, req *pb.GetExchangeListRequest) (*pb.GetExchangeListReply, error) {
	res := make([]*pb.GetExchangeListReply_List, 0)

	var (
		swapTrades []*SwapTrade
		err        error
	)

	swapTrades, err = ac.userRepo.GetSwapTrade(ctx, req.Start, req.End)
	if err != nil {
		return nil, err
	}

	for _, v := range swapTrades {

		var tmpPrice float64
		if 1 == v.Side && 0 < v.Amount0OutGross {
			tmpPrice = v.Amount1In / v.Amount0OutGross
		} else if 2 == v.Side && 0 < v.Amount0In {
			tmpPrice = v.Amount1OutGross / v.Amount0In
		} else {
			continue
		}

		res = append(res, &pb.GetExchangeListReply_List{
			BlockTime: v.BlockTime,
			Price:     tmpPrice,
		})
	}

	return &pb.GetExchangeListReply{List: res}, nil
}

func (ac *AppUsecase) GetBuyList(ctx context.Context, req *pb.GetBuyListRequest) (*pb.GetBuyListReply, error) {
	res := make([]*pb.GetBuyListReply_List, 0)

	var (
		buyList []*PrimaryBuy
		err     error
	)

	buyList, err = ac.userRepo.GetPrimaryBuy(ctx, req.Start, req.End)
	if err != nil {
		return nil, err
	}

	for _, v := range buyList {
		res = append(res, &pb.GetBuyListReply_List{
			BlockTime:    v.BlockTime,
			Price:        v.PriceAfter,
			UsdtUse:      v.UsdtUsed,
			AusdGrossOut: v.AusdGrossOut,
		})
	}

	return &pb.GetBuyListReply{List: res}, nil
}
