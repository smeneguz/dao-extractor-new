package types

// Default pagination configuration.
type DBPagination struct {
	Limit  int `yaml:"limit"`
	Offset int `yaml:"offset"`
}

// Custom database configuration.
type DBConfig struct {
	Pagination *DBPagination `yaml:"pagination"`
}
