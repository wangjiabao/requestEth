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

// data/nft_market_purchase.go
type NftMarketPurchase struct {
	ID uint64 `gorm:"primaryKey;autoIncrement;column:id"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null"`

	Buyer   string `gorm:"column:buyer;type:varchar(42);not null"`
	Seller  string `gorm:"column:seller;type:varchar(42);not null"`
	TokenID uint64 `gorm:"column:token_id;not null"`

	PriceUSDT  float64 `gorm:"column:price_usdt;type:decimal(65,18);not null"`
	FeePaidInB uint8   `gorm:"column:fee_paid_in_b;not null"`
	FeeUSDT    float64 `gorm:"column:fee_usdt;type:decimal(65,18);not null"`
	FeeB       float64 `gorm:"column:fee_b;type:decimal(65,18);not null"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"` // 你表没有也无所谓
	CheckStatus uint64    `gorm:"not null"`
}

type NftMinted struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement" json:"id"`

	BlockNumber uint64 `gorm:"column:block_number;not null" json:"block_number"`
	BlockTime   uint64 `gorm:"column:block_time;not null" json:"block_time"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0" json:"log_index"`

	ToAddr  string `gorm:"column:to_addr;type:varchar(42);not null" json:"to_addr"`
	TokenID uint64 `gorm:"column:token_id;not null" json:"token_id"`

	Tier     uint64  `gorm:"column:tier;not null;default:0" json:"tier"`
	UsdtPaid float64 `gorm:"column:usdt_paid;not null;default:0" json:"usdt_paid"`

	Status   uint8  `gorm:"column:status;not null;default:0" json:"status"`
	ListedAt uint64 `gorm:"column:listed_at;not null;default:0" json:"listed_at"`

	OpenStatus uint8  `gorm:"column:open_status;not null;default:0" json:"open_status"`
	OpenedAt   uint64 `gorm:"column:opened_at;not null;default:0" json:"opened_at"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
	CheckStatus uint64    `gorm:"not null"`
	CheckTime   uint64    `gorm:"not null"`

	MintAddr string `gorm:"column:mint_addr;type:varchar(42);not null" json:"mint_addr"`
}

type NftMarketListed struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	Seller    string `gorm:"column:seller;type:varchar(42);not null"`
	TokenID   uint64 `gorm:"column:token_id;not null"`
	Timestamp uint64 `gorm:"column:timestamp;not null;default:0"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

type NftMarketUnlisted struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	Operator string `gorm:"column:operator;type:varchar(42);not null"`
	TokenID  uint64 `gorm:"column:token_id;not null"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

type NftOpened struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	UserAddr string `gorm:"column:user_addr;type:varchar(42);not null"`
	TokenID  uint64 `gorm:"column:token_id;not null" json:"token_id"`

	OpenedAt uint64  `gorm:"column:opened_at;not null;default:0"`
	Reward   float64 `gorm:"column:reward;type:decimal(65,18);not null;default:0"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

type NftTransfer struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	FromAddr string `gorm:"column:from_addr;type:varchar(42);not null"`
	ToAddr   string `gorm:"column:to_addr;type:varchar(42);not null"`
	TokenID  uint64 `gorm:"column:token_id;not null"`

	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	CheckStatus uint64    `gorm:"not null"`
}

type UserRegistered struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement"`

	BlockNumber uint64 `gorm:"column:block_number;not null"`
	BlockTime   uint64 `gorm:"column:block_time;not null"`
	LogIndex    uint   `gorm:"column:log_index;not null;default:0"`

	UserAddr   string `gorm:"column:user_addr;type:varchar(42);not null"`
	ParentAddr string `gorm:"column:parent_addr;type:varchar(42);not null"`
	TopAddr    string `gorm:"column:top_addr;type:varchar(42);not null"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
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

// GetNftMarketPurchaseByAddressPage
// side: 0=buyer或seller  1=buyer  2=seller
func (u *UserRepo) GetNftMarketPurchaseByAddressPage(
	ctx context.Context,
	b *biz.Pagination,
	address string,
	addressTwo []string,
	side uint64,
) ([]*biz.NftMarketPurchase, error, int64) {

	var (
		count int64
		rows  []*NftMarketPurchase
	)

	res := make([]*biz.NftMarketPurchase, 0)

	instance := u.data.DB(ctx).Table("nft_market_purchase").Order("id desc")

	if address != "" {
		if side == 1 {
			instance = instance.Where("buyer = ?", address)
		} else if side == 2 {
			instance = instance.Where("seller = ?", address)
		} else if side == 3 {
			instance = instance.Where("seller in (?)", addressTwo)
		} else {
			instance = instance.Where("(buyer = ? OR seller = ?)", address, address)
		}
	}

	instance = instance.Count(&count)

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("NFT_PURCHASE_NOT_FOUND", "purchase not found"), 0
		}
		return nil, errors.New(500, "NFT_PURCHASE_PAGE_ERROR", err.Error()), 0
	}

	for _, r := range rows {
		res = append(res, &biz.NftMarketPurchase{
			ID:          r.ID,
			BlockNumber: r.BlockNumber,
			BlockTime:   r.BlockTime,
			LogIndex:    r.LogIndex,

			Buyer:   r.Buyer,
			Seller:  r.Seller,
			TokenID: r.TokenID,

			PriceUSDT:  r.PriceUSDT,
			FeePaidInB: r.FeePaidInB,
			FeeUSDT:    r.FeeUSDT,
			FeeB:       r.FeeB,

			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}

	return res, nil, count
}

func (u *UserRepo) InsertNftMarketPurchase(ctx context.Context, iData *biz.NftMarketPurchase) error {
	var s NftMarketPurchase

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.Buyer = iData.Buyer
	s.Seller = iData.Seller
	s.TokenID = iData.TokenID

	s.PriceUSDT = iData.PriceUSDT
	s.FeePaidInB = iData.FeePaidInB
	s.FeeUSDT = iData.FeeUSDT
	s.FeeB = iData.FeeB

	if err := u.data.DB(ctx).Table("nft_market_purchase").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MARKET_PURCHASE_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftMarketPurchaseLast(ctx context.Context) (*biz.NftMarketPurchase, error) {
	var v NftMarketPurchase

	if err := u.data.DB(ctx).Table("nft_market_purchase").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MARKET_PURCHASE_ERROR", err.Error())
	}

	return &biz.NftMarketPurchase{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		Buyer:   v.Buyer,
		Seller:  v.Seller,
		TokenID: v.TokenID,

		PriceUSDT:  v.PriceUSDT,
		FeePaidInB: v.FeePaidInB,
		FeeUSDT:    v.FeeUSDT,
		FeeB:       v.FeeB,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) UpdateNftMinted(ctx context.Context, tokenId uint64, mintAddr string) error {
	res := u.data.DB(ctx).Table("nft_minted").Where("token_id=?", tokenId).
		Updates(map[string]interface{}{
			"mint_addr":  mintAddr,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_MT_ERROR", "修改失败")
	}
	return nil
}

func (u *UserRepo) InsertNftMinted(ctx context.Context, iData *biz.NftMinted) error {
	var s NftMinted

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.ToAddr = iData.ToAddr
	s.TokenID = iData.TokenID

	s.Tier = iData.Tier
	s.UsdtPaid = iData.UsdtPaid
	s.MintAddr = iData.MintAddr

	// ✅ 默认值交给 DB：status/open_status/listed_at/opened_at
	// 如果你就是要插入时强行写，也可以从 iData 赋值

	if err := u.data.DB(ctx).Table("nft_minted").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MINTED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftMintedLast(ctx context.Context) (*biz.NftMinted, error) {
	var v NftMinted

	if err := u.data.DB(ctx).Table("nft_minted").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error())
	}

	return &biz.NftMinted{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		ToAddr:  v.ToAddr,
		TokenID: v.TokenID,

		Tier:     v.Tier,
		UsdtPaid: v.UsdtPaid,

		Status:   v.Status,
		ListedAt: v.ListedAt,

		OpenStatus: v.OpenStatus,
		OpenedAt:   v.OpenedAt,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) InsertNftMarketListed(ctx context.Context, iData *biz.NftMarketListed) error {
	var s NftMarketListed

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.Seller = iData.Seller
	s.TokenID = iData.TokenID
	s.Timestamp = iData.Timestamp

	if err := u.data.DB(ctx).Table("nft_market_listed").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MARKET_LISTED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftMarketListedLast(ctx context.Context) (*biz.NftMarketListed, error) {
	var v NftMarketListed

	if err := u.data.DB(ctx).Table("nft_market_listed").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MARKET_LISTED_ERROR", err.Error())
	}

	return &biz.NftMarketListed{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		Seller:    v.Seller,
		TokenID:   v.TokenID,
		Timestamp: v.Timestamp,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) InsertNftMarketUnlisted(ctx context.Context, iData *biz.NftMarketUnlisted) error {
	var s NftMarketUnlisted

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.Operator = iData.Operator
	s.TokenID = iData.TokenID

	if err := u.data.DB(ctx).Table("nft_market_unlisted").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_MARKET_UNLISTED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftMarketUnlistedLast(ctx context.Context) (*biz.NftMarketUnlisted, error) {
	var v NftMarketUnlisted

	if err := u.data.DB(ctx).Table("nft_market_unlisted").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_MARKET_UNLISTED_ERROR", err.Error())
	}

	return &biz.NftMarketUnlisted{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		Operator: v.Operator,
		TokenID:  v.TokenID,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) InsertNftOpened(ctx context.Context, iData *biz.NftOpened) error {
	var s NftOpened

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.UserAddr = iData.UserAddr
	s.TokenID = iData.TokenID

	s.OpenedAt = iData.OpenedAt
	s.Reward = iData.Reward

	if err := u.data.DB(ctx).Table("nft_opened").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_OPENED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftOpenedLast(ctx context.Context) (*biz.NftOpened, error) {
	var v NftOpened

	if err := u.data.DB(ctx).Table("nft_opened").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_OPENED_ERROR", err.Error())
	}

	return &biz.NftOpened{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		UserAddr: v.UserAddr,
		TokenID:  v.TokenID,

		OpenedAt: v.OpenedAt,
		Reward:   v.Reward,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) InsertNftTransfer(ctx context.Context, iData *biz.NftTransfer) error {
	var s NftTransfer

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.FromAddr = iData.FromAddr
	s.ToAddr = iData.ToAddr
	s.TokenID = iData.TokenID

	if err := u.data.DB(ctx).Table("nft_transfer").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_NFT_TRANSFER_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetNftTransferLast(ctx context.Context) (*biz.NftTransfer, error) {
	var v NftTransfer

	if err := u.data.DB(ctx).Table("nft_transfer").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	return &biz.NftTransfer{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		FromAddr: v.FromAddr,
		ToAddr:   v.ToAddr,
		TokenID:  v.TokenID,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

func (u *UserRepo) GetNftTransferLastNoCheck(ctx context.Context) ([]*biz.NftTransfer, error) {
	var vL []*NftTransfer

	res := make([]*biz.NftTransfer, 0)
	if err := u.data.DB(ctx).Table("nft_transfer").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftTransfer{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			FromAddr: v.FromAddr,
			ToAddr:   v.ToAddr,
			TokenID:  v.TokenID,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftListLastNoCheck(ctx context.Context) ([]*biz.NftMarketListed, error) {
	var vL []*NftMarketListed

	res := make([]*biz.NftMarketListed, 0)
	if err := u.data.DB(ctx).Table("nft_market_listed").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftMarketListed{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			Seller:    v.Seller,
			TokenID:   v.TokenID,
			Timestamp: v.Timestamp,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftUnListLastNoCheck(ctx context.Context) ([]*biz.NftMarketUnlisted, error) {
	var vL []*NftMarketUnlisted

	res := make([]*biz.NftMarketUnlisted, 0)
	if err := u.data.DB(ctx).Table("nft_market_unlisted").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftMarketUnlisted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			Operator: v.Operator,
			TokenID:  v.TokenID,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftBuyLastNoCheck(ctx context.Context) ([]*biz.NftMarketPurchase, error) {
	var vL []*NftMarketPurchase

	res := make([]*biz.NftMarketPurchase, 0)
	if err := u.data.DB(ctx).Table("nft_market_purchase").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftMarketPurchase{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			Buyer:   v.Buyer,
			Seller:  v.Seller,
			TokenID: v.TokenID,

			PriceUSDT:  v.PriceUSDT,
			FeePaidInB: v.FeePaidInB,
			FeeUSDT:    v.FeeUSDT,
			FeeB:       v.FeeB,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftOpenLastNoCheck(ctx context.Context) ([]*biz.NftOpened, error) {
	var vL []*NftOpened

	res := make([]*biz.NftOpened, 0)
	if err := u.data.DB(ctx).Table("nft_opened").Where("check_status=?", 0).Order("id asc").Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res = append(res, &biz.NftOpened{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			UserAddr: v.UserAddr,
			TokenID:  v.TokenID,

			OpenedAt: v.OpenedAt,
			Reward:   v.Reward,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

func (u *UserRepo) GetNftMintedByTokenIds(ctx context.Context, tokenIds []uint64) (map[uint64]*biz.NftMinted, error) {
	var vL []*NftMinted

	res := make(map[uint64]*biz.NftMinted, 0)
	if err := u.data.DB(ctx).Table("nft_minted").Where("token_id in(?)", tokenIds).Find(&vL).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, nil
		}
		return nil, errors.New(500, "NFT_TRANSFER_ERROR", err.Error())
	}

	for _, v := range vL {
		res[v.TokenID] = &biz.NftMinted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			ToAddr:  v.ToAddr,
			TokenID: v.TokenID,

			Tier:     v.Tier,
			UsdtPaid: v.UsdtPaid,

			Status:   v.Status,
			ListedAt: v.ListedAt,

			OpenStatus: v.OpenStatus,
			OpenedAt:   v.OpenedAt,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			CheckTime: v.CheckTime,
		}
	}

	return res, nil
}

// UpdateNftMintedToAddress .
func (u *UserRepo) UpdateNftMintedToAddress(ctx context.Context, id, idT, check uint64, toAddr string) error {
	res := u.data.DB(ctx).Table("nft_transfer").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_T_ERROR", "修改失败")
	}

	if 1 == check {
		//.Where("check_time<=?", checkTime).
		resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).
			Updates(map[string]interface{}{
				"to_addr": toAddr,
				//"check_time": checkTime,
				"updated_at": time.Now().Format("2006-01-02 15:04:05"),
			})
		if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
			return errors.New(500, "UPDATE_MINTED_ERROR", "修改失败")
		}
	}

	return nil
}

// UpdateNftMintedListStatus .
func (u *UserRepo) UpdateNftMintedListStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_listed").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	//.Where("check_time<=?", checkTime).
	resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).Where("check_time<?", checkTime).
		Updates(map[string]interface{}{
			"check_time": checkTime,
			"status":     1,
			"listed_at":  checkTime,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_L_ERROR", "修改失败")
	}

	return nil
}

// UpdateNftMintedUnListStatus .
func (u *UserRepo) UpdateNftMintedUnListStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_unlisted").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	//.Where("check_time<=?", checkTime).
	resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).Where("check_time<?", checkTime).
		Updates(map[string]interface{}{
			"check_time": checkTime,
			"status":     0,
			"listed_at":  0,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_U_L_ERROR", "修改失败")
	}

	return nil
}

func (u *UserRepo) UpdateUnlistedCheckStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_unlisted").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 2,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_CUL_ERROR", "修改失败")
	}

	return nil
}

func (u *UserRepo) UpdateListedCheckStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_listed").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 2,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_CL_ERROR", "修改失败")
	}

	return nil
}

func (u *UserRepo) UpdateBuyCheckStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_purchase").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 2,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_CB_ERROR", "修改失败")
	}

	return nil
}

// UpdateNftMintedBuyStatus .
func (u *UserRepo) UpdateNftMintedBuyStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_market_purchase").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	//.Where("check_time<=?", checkTime).
	resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).Where("check_time<?", checkTime).
		Updates(map[string]interface{}{
			"check_time": checkTime,
			"status":     0,
			"listed_at":  0,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_B_ERROR", "修改失败")
	}

	return nil
}

// UpdateNftMintedOpenStatus .
func (u *UserRepo) UpdateNftMintedOpenStatus(ctx context.Context, id, idT, checkTime uint64) error {
	res := u.data.DB(ctx).Table("nft_opened").Where("id=?", idT).
		Updates(map[string]interface{}{
			"check_status": 1,
			"updated_at":   time.Now().Format("2006-01-02 15:04:05"),
		})
	if res.Error != nil || 0 >= res.RowsAffected {
		return errors.New(500, "UPDATE_L_ERROR", "修改失败")
	}

	//.Where("check_time<=?", checkTime).
	resTwo := u.data.DB(ctx).Table("nft_minted").Where("id=?", id).
		Updates(map[string]interface{}{
			"open_status": 1,
			"opened_at":   checkTime,
			"updated_at":  time.Now().Format("2006-01-02 15:04:05"),
		})
	if resTwo.Error != nil || 0 >= resTwo.RowsAffected {
		return errors.New(500, "UPDATE_MINTED_O_ERROR", "修改失败")
	}

	return nil
}

// GetNftMintedPage .
func (u *UserRepo) GetNftMintedPage(
	ctx context.Context,
	b *biz.Pagination,
	order uint64,
	orderTwo uint64,
	tier uint64,
) ([]*biz.NftMinted, error, int64) {

	var (
		count int64
		rows  []*NftMinted
	)

	res := make([]*biz.NftMinted, 0)

	instance := u.data.DB(ctx).Table("nft_minted").Where("status=?", 1).Where("open_status=?", 0)

	if 0 < tier {
		instance = instance.Where("tier=?", tier)
	}

	instance = instance.Count(&count)

	if 0 < order {
		instance = instance.Order("id asc")
	} else {
		instance = instance.Order("id desc")
	}

	if 0 < orderTwo {
		instance = instance.Order("tier asc")
	} else {
		instance = instance.Order("tier desc")
	}

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("NFT_MINTED_NOT_FOUND", "minted not found"), 0
		}
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error()), 0
	}

	for _, v := range rows {
		res = append(res, &biz.NftMinted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			ToAddr:  v.ToAddr,
			TokenID: v.TokenID,

			Tier:     v.Tier,
			UsdtPaid: v.UsdtPaid,

			Status:   v.Status,
			ListedAt: v.ListedAt,

			OpenStatus: v.OpenStatus,
			OpenedAt:   v.OpenedAt,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil, count
}

// GetNftMintedByAddressPage .
func (u *UserRepo) GetNftMintedByAddressPage(
	ctx context.Context,
	b *biz.Pagination,
	start uint64,
	end uint64,
	address []string,
	status uint64,
	order uint64,
	tier uint64,
	openStatus uint64,
	openAtOrder uint64,
) ([]*biz.NftMinted, error, int64) {

	var (
		count int64
		rows  []*NftMinted
	)

	res := make([]*biz.NftMinted, 0)

	instance := u.data.DB(ctx).Table("nft_minted").Where("to_addr in(?)", address)

	if 0 < start && 0 < end {
		instance = instance.Where("block_time >= ?", start).Where("block_time <= ?", end)
	}

	if 2 > openStatus {
		if 1 == openStatus {
			instance = instance.Where("open_status=?", 1)
		} else {
			instance = instance.Where("open_status=?", 0)
		}
	}

	if 2 > status {
		if 1 == status {
			instance = instance.Where("status=?", 1)
		} else {
			instance = instance.Where("status=?", 0)
		}
	}

	if 0 < tier {
		instance = instance.Where("tier=?", tier)
	}

	instance = instance.Count(&count)

	if 0 < order {
		instance = instance.Order("id asc")
	} else {
		instance = instance.Order("id desc")
	}

	if 2 > openAtOrder {
		if 0 < openAtOrder {
			instance = instance.Order("opened_at asc")
		} else {
			instance = instance.Order("opened_at desc")
		}
	}

	if err := instance.Scopes(Paginate(b.PageNum, b.PageSize)).Find(&rows).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return res, errors.NotFound("NFT_MINTED_NOT_FOUND", "minted not found"), 0
		}
		return nil, errors.New(500, "NFT_MINTED_ERROR", err.Error()), 0
	}

	for _, v := range rows {
		res = append(res, &biz.NftMinted{
			ID:          v.ID,
			BlockNumber: v.BlockNumber,
			BlockTime:   v.BlockTime,
			LogIndex:    v.LogIndex,

			ToAddr:  v.ToAddr,
			TokenID: v.TokenID,

			Tier:     v.Tier,
			UsdtPaid: v.UsdtPaid,

			Status:   v.Status,
			ListedAt: v.ListedAt,

			OpenStatus: v.OpenStatus,
			OpenedAt:   v.OpenedAt,

			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			MintAddr:  v.MintAddr,
		})
	}

	return res, nil, count
}

func (u *UserRepo) InsertUserRegistered(ctx context.Context, iData *biz.UserRegistered) error {
	var s UserRegistered

	s.BlockNumber = iData.BlockNumber
	s.BlockTime = iData.BlockTime
	s.LogIndex = iData.LogIndex

	s.UserAddr = iData.UserAddr
	s.ParentAddr = iData.ParentAddr
	s.TopAddr = iData.TopAddr

	if err := u.data.DB(ctx).Table("user_registered").Create(&s).Error; err != nil {
		return errors.New(500, "CREATE_USER_REGISTERED_ERROR", "信息创建失败")
	}
	return nil
}

func (u *UserRepo) GetUserRegisteredLast(ctx context.Context) (*biz.UserRegistered, error) {
	var v UserRegistered

	if err := u.data.DB(ctx).Table("user_registered").Order("id desc").First(&v).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.New(500, "USER_REGISTERED_ERROR", err.Error())
	}

	return &biz.UserRegistered{
		ID:          v.ID,
		BlockNumber: v.BlockNumber,
		BlockTime:   v.BlockTime,
		LogIndex:    v.LogIndex,

		UserAddr:   v.UserAddr,
		ParentAddr: v.ParentAddr,
		TopAddr:    v.TopAddr,

		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
	}, nil
}

// GetUserRCount .
func (u *UserRepo) GetUserRCount(ctx context.Context) int64 {
	var (
		count int64
	)

	u.data.db.Table("user_registered").Count(&count)
	return count
}

// GetUserRCountBySe .
func (u *UserRepo) GetUserRCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("user_registered").Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

// GetMintNftCountBySe .
func (u *UserRepo) GetMintNftCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_minted").Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

// GetMintNftUsdtPaidSumBySe 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetMintNftUsdtPaidSumBySe(ctx context.Context, start, end uint64) string {
	var sum string
	u.data.db.Table("nft_minted").
		Select("COALESCE(SUM(usdt_paid), 0)").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Scan(&sum)
	return sum
}

// GetMintNftCount .
func (u *UserRepo) GetMintNftCount(paidType uint64) int64 {
	var (
		count int64
	)

	instance := u.data.db.Table("nft_minted")
	if 0 < paidType {
		instance = instance.Where("tier", paidType)
	}
	instance.Count(&count)
	return count
}

// GetMintNftUsdtPaidSum 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetMintNftUsdtPaidSum(paidType uint64) string {
	var sum string
	instance := u.data.db.Table("nft_minted")
	if 0 < paidType {
		instance = instance.Where("tier", paidType)
	}

	instance.Select("COALESCE(SUM(usdt_paid), 0)").Scan(&sum)
	return sum
}

// GetNftBuyCountBySe .
func (u *UserRepo) GetNftBuyCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_market_purchase").Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

// GetNftBuySumBySe 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetNftBuySumBySe(ctx context.Context, start, end uint64) string {
	var sum string
	u.data.db.Table("nft_market_purchase").
		Select("COALESCE(SUM(price_usdt), 0)").
		Where("block_time >= ?", start).
		Where("block_time <= ?", end).
		Scan(&sum)
	return sum
}

// GetNftBuyCount .
func (u *UserRepo) GetNftBuyCount() int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_market_purchase").Count(&count)
	return count
}

// GetNftBuySum 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetNftBuySum() string {
	var sum string
	u.data.db.Table("nft_market_purchase").
		Select("COALESCE(SUM(price_usdt), 0)").
		Scan(&sum)
	return sum
}

// GetNftOpenCountBySe .
func (u *UserRepo) GetNftOpenCountBySe(ctx context.Context, start, end uint64) int64 {
	var (
		count int64
	)

	u.data.db.Table("nft_opened").Where("block_time >= ?", start).Where("block_time <= ?", end).Count(&count)
	return count
}

// GetNftOpenSumB 统计时间段内 mint 的 usdt_paid 总和（返回字符串，避免精度问题）
func (u *UserRepo) GetNftOpenSum() string {
	var sum string
	u.data.db.Table("nft_opened").
		Select("COALESCE(SUM(reward), 0)").
		Scan(&sum)
	return sum
}
