package bot

import (
	"time"

	"github.com/TradlyLabs/tradly-common/pkg/services/db"
	"gorm.io/gorm"
)

// BotBot represents the bot.bot_bots table structure
type BotBot struct {
	BotID         string    `gorm:"column:bot_id;primaryKey;not null"`
	UserID        string    `gorm:"column:user_id;not null"`
	Name          string    `gorm:"column:name;type:varchar(255);not null"`
	SourceBotID   *int64    `gorm:"column:source_bot_id"`
	StrategyType  int32     `gorm:"column:strategy_type;not null"`
	StrategyID    int64     `gorm:"column:strategy_id;not null"`
	PairID        int32     `gorm:"column:pair_id;not null"`
	ChainID       int64     `gorm:"column:chain_id;not null"`
	DexID         int32     `gorm:"column:dex_id;not null"`
	InvestToken   string    `gorm:"column:invest_token;type:char(42);not null"`
	InvestAmount  string    `gorm:"column:invest_amount;type:numeric(78,18);not null"`
	StartUSD      float64   `gorm:"column:start_usd;type:numeric(18,2);not null"`
	EndUSD        float64   `gorm:"column:end_usd;type:numeric(18,2);not null"`
	Investment    float64   `gorm:"column:investment;type:numeric(18,2);not null"`
	RealizedPNL   float64   `gorm:"column:realized_pnl;type:numeric(18,2);not null"`
	UnrealizedPNL float64   `gorm:"column:unrealized_pnl;type:numeric(18,2);not null"`
	Copiers       int32     `gorm:"column:copiers;not null"`
	Runtime       *int64    `gorm:"column:runtime"`
	Status        *int32    `gorm:"column:status;default:1"`
	CreatedAt     time.Time `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt     time.Time `gorm:"column:updated_at;not null;default:now()"`
}

// TableName sets the table name for the BotBot struct
func (BotBot) TableName() string {
	return "bot.bot_bots"
}

func (b *BotBot) Save() error {
	return db.Get().Save(b).Error
}

// Create inserts a new bot into the database
func (b *BotBot) Create() error {
	return db.Get().Create(b).Error
}

// Update updates the bot in the database
func (b *BotBot) Update() error {
	return db.Get().Save(b).Error
}

// Delete removes the bot from the database
func (b *BotBot) Delete() error {
	return db.Get().Delete(b).Error
}

// GetBotByID retrieves a bot by its ID
func GetBotByID(botID string) (*BotBot, error) {
	var bot BotBot
	err := db.Get().Where("bot_id = ?", botID).First(&bot).Error
	if err != nil {
		return nil, err
	}
	return &bot, nil
}

// GetBotByOrderID retrieves a bot by its OrderID
func GetBotByOrderID(orderID string) (*BotBot, error) {
	var bot BotBot
	err := db.Get().Where("order_id = ?", orderID).First(&bot).Error
	if err != nil {
		return nil, err
	}
	return &bot, nil
}

// GetAllBots retrieves all bots with pagination
func GetAllBots(page, pageSize int) ([]BotBot, int64, error) {
	var bots []BotBot
	var total int64

	dbInstance := db.Get()
	dbInstance.Model(&BotBot{}).Count(&total)

	offset := (page - 1) * pageSize
	err := dbInstance.Offset(offset).Limit(pageSize).Find(&bots).Error
	if err != nil {
		return nil, 0, err
	}

	return bots, total, nil
}

// UpdateBotStatus updates the status of a bot
func UpdateBotStatus(botID string, status BotStatus) error {
	return db.Get().Model(&BotBot{}).Where("bot_id = ?", botID).Update("status", status).Error

}

// IncrementRetryCount increments the retry count for a bot
func IncrementRetryCount(botID string) error {
	return db.Get().Model(&BotBot{}).Where("bot_id = ?", botID).Update("retry_count", gorm.Expr("retry_count + 1")).Error
}
