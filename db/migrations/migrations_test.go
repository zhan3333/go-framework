package migrations_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/bootstrap"
	"go-framework/db"
	"go-framework/db/migrations"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap.SetInTest()
	bootstrap.Bootstrap()
	m.Run()
}

func TestGetAllMigrations(t *testing.T) {
	allMigrations := migrations.GetAllMigrations()
	assert.NotNil(t, allMigrations)
}

func TestGetNeedMigrateFiles(t *testing.T) {
	err := db.Conn.Exec("truncate migrations").Error
	assert.Nil(t, err)
	needMigrateFiles := migrations.GetNeedMigrateFiles(migrations.MigrateFiles)
	assert.NotNil(t, needMigrateFiles)
	assert.Equal(t, len(migrations.MigrateFiles), len(needMigrateFiles))
	m := migrations.Migration{
		Migration: migrations.MigrateFiles[0].Key(),
		Batch:     1,
	}
	db.Conn.Save(&m)

	needMigrateFiles = migrations.GetNeedMigrateFiles(migrations.MigrateFiles)
	assert.Equal(t, 0, len(needMigrateFiles))
}

func TestGetNeedRollbackKeys(t *testing.T) {
	err := db.Conn.Exec("truncate migrations").Error
	assert.Nil(t, err)
	var needRollbackMs []migrations.MigrateFile
	needRollbackMs = migrations.GetNeedRollbackKeys(1)
	assert.Equal(t, 0, len(needRollbackMs))

	m := migrations.Migration{
		Migration: migrations.MigrateFiles[0].Key(),
		Batch:     1,
	}
	db.Conn.Create(&m)

	needRollbackMs = migrations.GetNeedRollbackKeys(1)
	assert.Equal(t, 1, len(needRollbackMs))
	assert.Equal(t, migrations.MigrateFiles[0].Key(), needRollbackMs[0].Key())
}

func TestGetNextBatchNo(t *testing.T) {
	var nextBatch uint
	err := db.Conn.Exec("truncate migrations").Error
	assert.Nil(t, err)
	nextBatch = migrations.GetNextBatchNo()
	assert.Equal(t, uint(1), nextBatch)
	m := migrations.Migration{
		Migration: migrations.MigrateFiles[0].Key(),
		Batch:     nextBatch,
	}
	db.Conn.Create(&m)

	nextBatch = migrations.GetNextBatchNo()
	assert.Equal(t, uint(2), nextBatch)
}
