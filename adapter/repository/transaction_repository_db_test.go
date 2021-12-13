package repository

import (
	"testing"
	"os"
	"github.com/jvcalassio/fc-payment-gateway/adapter/repository/fixture"
	"github.com/jvcalassio/fc-payment-gateway/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestTransactionRepository_DbInsert(t *testing.T) {
	migrationsDir := os.DirFS("fixture/sql")
	db := fixture.Up(migrationsDir, true)
	defer fixture.Down(db, migrationsDir)

	repository := NewTransactionRepositoryDb(db)
	err := repository.Insert("1", "1", 12.1, entity.APPROVED, "")
	assert.Nil(t, err)
}