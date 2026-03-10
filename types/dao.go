package types

// -------------------------------------------------------------------------------------------------------------------
// ---- DAO structs
// -------------------------------------------------------------------------------------------------------------------

type DAOID uint64

type DAOSymbol string

type DAO struct {
	// Unique ID of the DAO.
	ID DAOID `json:"id"`
	// The symbol of the DAO.
	Symbol DAOSymbol `json:"symbol"`
	// The name of the DAO.
	Name string `json:"name"`
}

func NewDAO(symbol DAOSymbol, name string) *DAO {
	return &DAO{
		// Set the ID to 0 to indicate that this object has been created
		// by the code and it may not exist in the database yet.
		ID:     0,
		Symbol: symbol,
		Name:   name,
	}
}

// WithID sets the ID of the DAO.
func (d *DAO) WithID(id DAOID) *DAO {
	d.ID = id
	return d
}
