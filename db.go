package glow

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

// OpenDB opens gorm.DB connection
func OpenDB() (*gorm.DB, error) {
	dialect := os.Getenv("GORM_DIALECT")
	if dialect == "" {
		return nil, fmt.Errorf("GORM_DIALECT variable is not set")
	}
	connection := os.Getenv("GORM_CONNECTION")
	if connection == "" {
		return nil, fmt.Errorf("GORM_CONNECTION variable is not set")
	}

	return gorm.Open(dialect, connection)
}
