package postgres

import (
	"github.com/kartikmodi25/URL-Shortener/pkg/models"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type client struct {
	_db *gorm.DB
}

func (c *client) db() *gorm.DB {
	return c._db
}
// AutoMigrate create or update tables as when needed.
func (c *client) AutoMigrate() error {
	db := c.db()
	tables := []interface{}{
		&models.URL{},
	}
	err := db.AutoMigrate(tables...)
	return errors.Wrap(err, "AutoMigrate")
}
