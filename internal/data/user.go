package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"requestEth/internal/biz"
	"time"
)

type SwapTrade struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"type:bigint unsigned;not null;index:idx_block_number;index:idx_block_log,priority:1;comment:区块高度"`
	LogIndex    uint32 `gorm:"type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`
	BlockTime   uint64 `gorm:"type:bigint unsigned;not null;index:idx_block_number;index:idx_block_log,priority:1;comment:区块高度"`

	Sender string `gorm:"type:varchar(42);not null;index:idx_sender_block,priority:1;comment:事件sender(0x...)"`
	ToAddr string `gorm:"column:to_addr;type:varchar(42);not null;index:idx_to_block,priority:1;comment:事件to(0x...)"`

	Side uint8 `gorm:"type:tinyint unsigned;not null;default:0;index:idx_side_block,priority:1;comment:方向1=BUY(DL)2=SELL(DL)0=UNKNOWN"`

	Amount0In       float64 `gorm:"column:amount0_in;type:decimal(65,18);not null;default:0;comment:DL in"`
	Amount1In       float64 `gorm:"column:amount1_in;type:decimal(65,18);not null;default:0;comment:OTHER in"`
	Amount0OutGross float64 `gorm:"column:amount0_out_gross;type:decimal(65,18);not null;default:0;comment:DL out"`
	Amount0OutNet   float64 `gorm:"column:amount0_out_net;type:decimal(65,18);not null;default:0;comment:DL out net"`
	Amount1OutGross float64 `gorm:"column:amount1_out_gross;type:decimal(65,18);not null;default:0;comment:OTHER out"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

type PrimaryBuy struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"column:block_number;type:bigint unsigned;not null;index:idx_block_log,priority:1;comment:区块高度"`
	BlockTime   uint64 `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
	LogIndex    uint32 `gorm:"column:log_index;type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`

	Buyer  string `gorm:"column:buyer;type:varchar(42);not null;index:idx_buyer;comment:buyer"`
	ToAddr string `gorm:"column:to_addr;type:varchar(42);not null;index:idx_to;comment:to"`

	UsdtUsed     float64 `gorm:"column:usdt_used;type:decimal(65,18);not null;default:0;comment:usdtUsed"`
	AusdGrossOut float64 `gorm:"column:ausd_gross_out;type:decimal(65,18);not null;default:0;comment:ausdGrossOut"`
	AusdFee      float64 `gorm:"column:ausd_fee;type:decimal(65,18);not null;default:0;comment:ausdFee"`
	AusdNetOut   float64 `gorm:"column:ausd_net_out;type:decimal(65,18);not null;default:0;comment:ausdNetOut"`
	PriceBefore  float64 `gorm:"column:price_before;type:decimal(65,18);not null;default:0;comment:priceBefore"`
	PriceAfter   float64 `gorm:"column:price_after;type:decimal(65,18);not null;default:0;comment:priceAfter"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

type PrimarySell struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"column:block_number;type:bigint unsigned;not null;index:idx_block_log,priority:1;comment:区块高度"`
	BlockTime   uint64 `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
	LogIndex    uint32 `gorm:"column:log_index;type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`

	Seller string `gorm:"column:seller;type:varchar(42);not null;index:idx_seller;comment:seller"`
	ToAddr string `gorm:"column:to_addr;type:varchar(42);not null;index:idx_to;comment:to"`

	AusdGrossIn float64 `gorm:"column:ausd_gross_in;type:decimal(65,18);not null;default:0;comment:ausdGrossIn"`
	AusdFee     float64 `gorm:"column:ausd_fee;type:decimal(65,18);not null;default:0;comment:ausdFee"`
	AusdBurn    float64 `gorm:"column:ausd_burn;type:decimal(65,18);not null;default:0;comment:ausdBurn"`
	UsdtOut     float64 `gorm:"column:usdt_out;type:decimal(65,18);not null;default:0;comment:usdtOut"`
	PriceBefore float64 `gorm:"column:price_before;type:decimal(65,18);not null;default:0;comment:priceBefore"`
	PriceAfter  float64 `gorm:"column:price_after;type:decimal(65,18);not null;default:0;comment:priceAfter"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

type RewardNotified struct {
	ID uint64 `gorm:"primarykey;type:bigint unsigned;comment:主键"`

	BlockNumber uint64 `gorm:"column:block_number;type:bigint unsigned;not null;index:idx_block_log,priority:1;comment:区块高度"`
	BlockTime   uint64 `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
	LogIndex    uint32 `gorm:"column:log_index;type:int unsigned;not null;default:0;index:idx_block_log,priority:2;comment:logIndex(排序用)"`

	User string `gorm:"column:user;type:varchar(42);not null;index:idx_user_time,priority:1;comment:indexed user"`
	L1   string `gorm:"column:l1;type:varchar(42);not null;index:idx_l1_time,priority:1;comment:indexed l1"`
	L2   string `gorm:"column:l2;type:varchar(42);not null;index:idx_l2_time,priority:1;comment:indexed l2"`

	Profit    float64 `gorm:"column:profit;type:decimal(65,18);not null;default:0;comment:profit"`
	UserShare float64 `gorm:"column:user_share;type:decimal(65,18);not null;default:0;comment:userShare"`
	Top       string  `gorm:"column:top;type:varchar(42);not null;index:idx_top_time,priority:1;comment:top"`

	Pool             float64 `gorm:"column:pool;type:decimal(65,18);not null;default:0;comment:pool"`
	UplinePortionBps uint64  `gorm:"column:upline_portion_bps;type:bigint unsigned;not null;default:0;comment:uplinePortionBps"`

	ToL1      float64 `gorm:"column:to_l1;type:decimal(65,18);not null;default:0;comment:toL1"`
	ToL2      float64 `gorm:"column:to_l2;type:decimal(65,18);not null;default:0;comment:toL2"`
	ToTop     float64 `gorm:"column:to_top;type:decimal(65,18);not null;default:0;comment:toTop"`
	ToProject float64 `gorm:"column:to_project;type:decimal(65,18);not null;default:0;comment:toProject"`

	CreatedAt time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt time.Time `gorm:"type:datetime;not null;comment:更新时间"`
}

type RewardDetail struct {
	ID         uint64    `gorm:"primarykey;type:bigint unsigned;comment:主键"`
	User       string    `gorm:"column:user;type:varchar(42);not null;index:idx_user_time,priority:1;comment:indexed user"`
	Amount     float64   `gorm:"column:amount;type:decimal(65,18);not null;default:0;comment:amount"`
	Reason     uint64    `gorm:"column:reason;type:bigint unsigned;not null;default:0;comment:reason"`
	NotifiedId uint64    `gorm:"column:notified_id;type:int unsigned;not null;default:0;comment:notified_id"`
	CreatedAt  time.Time `gorm:"type:datetime;not null;comment:创建时间"`
	UpdatedAt  time.Time `gorm:"type:datetime;not null;comment:更新时间"`
	BlockTime  uint64    `gorm:"column:block_time;type:bigint unsigned;not null;index:idx_block_time;comment:区块时间(秒)"`
}

type UserRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &UserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *UserRepo) GetSwapTradeLast(ctx context.Context) (*biz.SwapTrade, error) {
	var v *SwapTrade

	if err := u.data.DB(ctx).Table("swap_trade").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "SWAP_TRADE_ERROR", err.Error())
	}

	return &biz.SwapTrade{
		ID:              v.ID,
		BlockNumber:     v.BlockNumber,
		LogIndex:        v.LogIndex,
		BlockTime:       v.BlockTime,
		Sender:          v.Sender,
		ToAddr:          v.ToAddr,
		Side:            v.Side,
		Amount0In:       v.Amount0In,
		Amount1In:       v.Amount1In,
		Amount0OutNet:   v.Amount0OutNet,
		Amount1OutGross: v.Amount1OutGross,
		Amount0OutGross: v.Amount0OutGross,
		CreatedAt:       v.CreatedAt,
		UpdatedAt:       v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetSwapTrade(ctx context.Context, start, end uint64) ([]*biz.SwapTrade, error) {
	var s []*SwapTrade

	res := make([]*biz.SwapTrade, 0)
	if err := u.data.DB(ctx).Table("swap_trade").Where("block_time >=?", start).Where("block_time <=?", end).Find(&s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "SWAP_TRADE_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.SwapTrade{
			ID:              v.ID,
			BlockNumber:     v.BlockNumber,
			LogIndex:        v.LogIndex,
			BlockTime:       v.BlockTime,
			Sender:          v.Sender,
			ToAddr:          v.ToAddr,
			Side:            v.Side,
			Amount0In:       v.Amount0In,
			Amount1In:       v.Amount1In,
			Amount0OutNet:   v.Amount0OutNet,
			Amount1OutGross: v.Amount1OutGross,
			Amount0OutGross: v.Amount0OutGross,
			CreatedAt:       v.CreatedAt,
			UpdatedAt:       v.UpdatedAt,
		})
	}
	return res, nil
}

// InsertSwapTrade .
func (u *UserRepo) InsertSwapTrade(ctx context.Context, iData *biz.SwapTrade) error {
	var (
		err error
		s   SwapTrade
	)

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.Side = iData.Side
	s.Sender = iData.Sender
	s.Amount1In = iData.Amount1In
	s.Amount0OutGross = iData.Amount0OutGross
	s.Amount0OutNet = iData.Amount0OutNet
	s.Amount1OutGross = iData.Amount1OutGross
	s.Amount0In = iData.Amount0In
	s.ToAddr = iData.ToAddr
	s.LogIndex = iData.LogIndex

	err = u.data.DB(ctx).Table("swap_trade").Create(&s).Error
	if err != nil {
		return errors.New(500, "CREATE_SWAP_TRADE_ERROR", "信息创建失败")
	}

	return nil
}

func (u *UserRepo) GetPrimaryBuyLast(ctx context.Context) (*biz.PrimaryBuy, error) {
	var v PrimaryBuy

	if err := u.data.DB(ctx).Table("primary_buy").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "PRIMARY_BUY_ERROR", err.Error())
	}

	return &biz.PrimaryBuy{
		ID:           v.ID,
		BlockNumber:  v.BlockNumber,
		BlockTime:    v.BlockTime,
		LogIndex:     v.LogIndex,
		Buyer:        v.Buyer,
		ToAddr:       v.ToAddr,
		UsdtUsed:     v.UsdtUsed,
		AusdGrossOut: v.AusdGrossOut,
		AusdFee:      v.AusdFee,
		AusdNetOut:   v.AusdNetOut,
		PriceBefore:  v.PriceBefore,
		PriceAfter:   v.PriceAfter,
		CreatedAt:    v.CreatedAt,
		UpdatedAt:    v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetPrimaryBuy(ctx context.Context, start, end uint64) ([]*biz.PrimaryBuy, error) {
	var s []PrimaryBuy
	res := make([]*biz.PrimaryBuy, 0)

	if err := u.data.DB(ctx).
		Table("primary_buy").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Order("block_time asc, id asc").
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "PRIMARY_BUY_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.PrimaryBuy{
			ID:           v.ID,
			BlockNumber:  v.BlockNumber,
			BlockTime:    v.BlockTime,
			LogIndex:     v.LogIndex,
			Buyer:        v.Buyer,
			ToAddr:       v.ToAddr,
			UsdtUsed:     v.UsdtUsed,
			AusdGrossOut: v.AusdGrossOut,
			AusdFee:      v.AusdFee,
			AusdNetOut:   v.AusdNetOut,
			PriceBefore:  v.PriceBefore,
			PriceAfter:   v.PriceAfter,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}
	return res, nil
}

// InsertPrimaryBuy .
func (u *UserRepo) InsertPrimaryBuy(ctx context.Context, iData *biz.PrimaryBuy) error {
	var s PrimaryBuy

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex
	s.Buyer = iData.Buyer
	s.ToAddr = iData.ToAddr

	s.UsdtUsed = iData.UsdtUsed
	s.AusdGrossOut = iData.AusdGrossOut
	s.AusdFee = iData.AusdFee
	s.AusdNetOut = iData.AusdNetOut
	s.PriceBefore = iData.PriceBefore
	s.PriceAfter = iData.PriceAfter

	if err := u.data.DB(ctx).Table("primary_buy").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_PRIMARY_BUY_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetPrimarySellLast(ctx context.Context) (*biz.PrimarySell, error) {
	var v PrimarySell

	if err := u.data.DB(ctx).Table("primary_sell").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "PRIMARY_SELL_ERROR", err.Error())
	}

	return &biz.PrimarySell{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,
		Seller:      v.Seller,
		ToAddr:      v.ToAddr,
		AusdGrossIn: v.AusdGrossIn,
		AusdFee:     v.AusdFee,
		AusdBurn:    v.AusdBurn,
		UsdtOut:     v.UsdtOut,
		PriceBefore: v.PriceBefore,
		PriceAfter:  v.PriceAfter,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetPrimarySell(ctx context.Context, start, end uint64) ([]*biz.PrimarySell, error) {
	var s []PrimarySell
	res := make([]*biz.PrimarySell, 0)

	if err := u.data.DB(ctx).
		Table("primary_sell").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Order("block_time asc, id asc").
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "PRIMARY_SELL_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.PrimarySell{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,
			Seller:      v.Seller,
			ToAddr:      v.ToAddr,
			AusdGrossIn: v.AusdGrossIn,
			AusdFee:     v.AusdFee,
			AusdBurn:    v.AusdBurn,
			UsdtOut:     v.UsdtOut,
			PriceBefore: v.PriceBefore,
			PriceAfter:  v.PriceAfter,
			CreatedAt:   v.CreatedAt,
			UpdatedAt:   v.UpdatedAt,
		})
	}
	return res, nil
}

// InsertPrimarySell .
func (u *UserRepo) InsertPrimarySell(ctx context.Context, iData *biz.PrimarySell) error {
	var s PrimarySell

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex
	s.Seller = iData.Seller
	s.ToAddr = iData.ToAddr

	s.AusdGrossIn = iData.AusdGrossIn
	s.AusdFee = iData.AusdFee
	s.AusdBurn = iData.AusdBurn
	s.UsdtOut = iData.UsdtOut
	s.PriceBefore = iData.PriceBefore
	s.PriceAfter = iData.PriceAfter

	if err := u.data.DB(ctx).Table("primary_sell").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_PRIMARY_SELL_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetRewardNotifiedLast(ctx context.Context) (*biz.RewardNotified, error) {
	var v RewardNotified

	if err := u.data.DB(ctx).Table("reward_notified").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "REWARD_NOTIFIED_ERROR", err.Error())
	}

	return &biz.RewardNotified{
		ID:               v.ID,
		BlockNumber:      v.BlockNumber,
		BlockTime:        v.BlockTime,
		LogIndex:         v.LogIndex,
		User:             v.User,
		L1:               v.L1,
		L2:               v.L2,
		Profit:           v.Profit,
		UserShare:        v.UserShare,
		Top:              v.Top,
		Pool:             v.Pool,
		UplinePortionBps: v.UplinePortionBps,
		ToL1:             v.ToL1,
		ToL2:             v.ToL2,
		ToTop:            v.ToTop,
		ToProject:        v.ToProject,
		CreatedAt:        v.CreatedAt,
		UpdatedAt:        v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetRewardNotified(ctx context.Context, start, end uint64) ([]*biz.RewardNotified, error) {
	var s []RewardNotified
	res := make([]*biz.RewardNotified, 0)

	if err := u.data.DB(ctx).
		Table("reward_notified").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Order("block_time asc, id asc").
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "REWARD_NOTIFIED_ERROR", err.Error())
	}

	for _, v := range s {
		res = append(res, &biz.RewardNotified{
			ID:               v.ID,
			BlockNumber:      v.BlockNumber,
			BlockTime:        v.BlockTime,
			LogIndex:         v.LogIndex,
			User:             v.User,
			L1:               v.L1,
			L2:               v.L2,
			Profit:           v.Profit,
			UserShare:        v.UserShare,
			Top:              v.Top,
			Pool:             v.Pool,
			UplinePortionBps: v.UplinePortionBps,
			ToL1:             v.ToL1,
			ToL2:             v.ToL2,
			ToTop:            v.ToTop,
			ToProject:        v.ToProject,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		})
	}
	return res, nil
}

func (u *UserRepo) GetRewardNotifiedByIds(ctx context.Context, ids []uint64) (map[uint64]*biz.RewardNotified, error) {
	var s []RewardNotified
	res := make(map[uint64]*biz.RewardNotified, 0)

	if err := u.data.DB(ctx).
		Table("reward_notified").
		Where("id in(?)", ids).
		Find(&s).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "REWARD_NOTIFIED_ERROR", err.Error())
	}

	for _, v := range s {
		res[v.ID] = &biz.RewardNotified{
			ID:               v.ID,
			BlockNumber:      v.BlockNumber,
			BlockTime:        v.BlockTime,
			LogIndex:         v.LogIndex,
			User:             v.User,
			L1:               v.L1,
			L2:               v.L2,
			Profit:           v.Profit,
			UserShare:        v.UserShare,
			Top:              v.Top,
			Pool:             v.Pool,
			UplinePortionBps: v.UplinePortionBps,
			ToL1:             v.ToL1,
			ToL2:             v.ToL2,
			ToTop:            v.ToTop,
			ToProject:        v.ToProject,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		}
	}
	return res, nil
}

func (u *UserRepo) InsertRewardNotified(ctx context.Context, iData *biz.RewardNotified) error {
	var s RewardNotified

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.User = iData.User
	s.L1 = iData.L1
	s.L2 = iData.L2

	s.Profit = iData.Profit
	s.UserShare = iData.UserShare
	s.Top = iData.Top

	s.Pool = iData.Pool
	s.UplinePortionBps = iData.UplinePortionBps

	s.ToL1 = iData.ToL1
	s.ToL2 = iData.ToL2
	s.ToTop = iData.ToTop
	s.ToProject = iData.ToProject

	if err := u.data.DB(ctx).Table("reward_notified").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_NOTIFIED_ERROR", "信息创建失败")
	}

	var one RewardDetail
	one.User = s.User
	one.Reason = 1
	one.NotifiedId = s.ID
	one.Amount = s.UserShare
	one.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&one).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var two RewardDetail
	two.User = s.L1
	two.Reason = 2
	two.NotifiedId = s.ID
	two.Amount = s.ToL1
	two.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&two).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var three RewardDetail
	three.User = s.L2
	three.Reason = 3
	three.NotifiedId = s.ID
	three.Amount = s.ToL2
	three.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&three).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var four RewardDetail
	four.User = s.Top
	four.Reason = 4
	four.NotifiedId = s.ID
	four.Amount = s.ToTop
	four.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&four).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	var five RewardDetail
	five.User = "0x144351b1a5Af4538f53556633D8f10495aA564A8"
	five.Reason = 5
	five.NotifiedId = s.ID
	five.Amount = s.ToProject
	five.BlockTime = s.BlockTime
	if err := u.data.DB(ctx).Table("reward_detail").Create(&five).Error; err != nil {
		return errors.New(500, "CREATE_REWARD_DETAIL_ERROR", "信息创建失败")
	}

	return nil
}

// GetUserRewardByUserIdPage .
func (u *UserRepo) GetUserRewardByUserIdPage(ctx context.Context, b *biz.Pagination, address string, reason uint64) ([]*biz.RewardDetail, error, int64) {
	var (
		count   int64
		rewards []*RewardDetail
	)

	res := make([]*biz.RewardDetail, 0)

	instance := u.data.db.Where("user", address).Table("reward_detail").Order("id desc")
	if 0 < reason {
		instance = instance.Where("reason=?", reason)
	}

	instance = instance.Count(&count)

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rewards).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("REWARD_NOT_FOUND", "reward not found"), 0
		}

		return nil, errors.New(500, "REWARD ERROR", err.Error()), 0
	}

	for _, reward := range rewards {
		res = append(res, &biz.RewardDetail{
			ID:         reward.ID,
			User:       reward.User,
			Amount:     reward.Amount,
			Reason:     reward.Reason,
			BlockTime:  reward.BlockTime,
			NotifiedId: reward.NotifiedId,
			CreatedAt:  reward.CreatedAt,
			UpdatedAt:  reward.UpdatedAt,
		})
	}

	return res, nil, count
}
