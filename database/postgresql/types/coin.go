package types

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/dao-portal/extractor/types"
)

type DBCoin struct {
	Coin *types.Coin
}

var _ driver.Valuer = (*DBCoin)(nil)
var _ sql.Scanner = (*DBCoin)(nil)

func NewDBCoin(coin *types.Coin) DBCoin {
	return DBCoin{
		Coin: coin,
	}
}

// GetCoin returns the Coin value.
func (c *DBCoin) GetCoin() *types.Coin {
	return c.Coin
}

// -----------------------------------------------------------------------------
// ---- Database value and scan implementations
// -----------------------------------------------------------------------------

// Value implements driver.Valuer
func (c DBCoin) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%s)", c.Coin.Denom, c.Coin.Amount.String()), nil
}

// Scan implements sql.Scanner
func (c *DBCoin) Scan(src any) error {
	strValue := string(src.([]byte))
	strValue = strings.ReplaceAll(strValue, `"`, "")
	strValue = strings.ReplaceAll(strValue, "{", "")
	strValue = strings.ReplaceAll(strValue, "}", "")
	strValue = strings.ReplaceAll(strValue, "(", "")
	strValue = strings.ReplaceAll(strValue, ")", "")

	values := strings.Split(strValue, ",")
	denom := values[0]
	amount := values[1]
	coin, err := types.NewCoinFromString(denom, amount)
	if err != nil {
		return err
	}

	c.Coin = &coin
	return nil
}
