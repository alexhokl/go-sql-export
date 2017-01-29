package model

// ExportConfig struct
type ExportConfig struct {
	connectionString string `yaml:"connection_string"`
	documentName     string `yaml:"document_name"`
	sheets           []SheetConfig
}

// SheetConfig struct
type SheetConfig struct {
	name    string
	query   string
	columns []ColumnConfig
}

// ColumnConfig struct
type ColumnConfig struct {
	index    int
	dataType string `yaml:"data_type"`
	format   string
}

// ConnectionString returns connection string
func (c *ExportConfig) ConnectionString() string {
	return c.connectionString
}

// DocumentName return export document name
func (c *ExportConfig) DocumentName() string {
	return c.documentName
}

// Sheets returns configuration of export query (and sheet)
func (c *ExportConfig) Sheets() []SheetConfig {
	return c.sheets
}

// Name returns name of this export sheet
func (s *SheetConfig) Name() string {
	return s.name
}

// Query returns the T-SQL query statement of this sheet
func (s *SheetConfig) Query() string {
	return s.query
}

// Columns returns columns of this sheet
func (s *SheetConfig) Columns() []ColumnConfig {
	return s.columns
}

// Index returns the column index (zero-based) in this sheet
func (c *ColumnConfig) Index() int {
	return c.index
}

// DataType returns data type of this column
func (c *ColumnConfig) DataType() string {
	return c.dataType
}

// Format returns the format string of this column
func (c *ColumnConfig) Format() string {
	return c.format
}