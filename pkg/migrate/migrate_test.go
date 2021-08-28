package migrate

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go-framework/pkg/gdb"
	"go-framework/pkg/migrate/testdata"
	"gorm.io/gorm"
	"math"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gdb.ConnConfigs = map[string]gdb.MySQLConf{
		gdb.DefaultName: {
			Host:     "127.0.0.1",
			Port:     "3306",
			Username: "root",
			Password: "123456",
			Database: "test",
		},
	}
	DB = gdb.Def()
	err := DelAll()
	if err != nil {
		panic(err)
	}
	exitCode := m.Run()
	err = DelAll()
	if err != nil {
		panic(err)
	}
	os.Exit(exitCode)
}

func TestGetAllMigrations(t *testing.T) {
	assert.Nil(t, DelAll())
	assert.Nil(t, InitMigrationTable())
	allMigrations := getAllMigrations()
	assert.NotNil(t, allMigrations)
}

func TestGetNeedMigrateFiles(t *testing.T) {
	assert.Nil(t, DelAll())
	assert.Nil(t, InitMigrationTable())
	Register(&testdata.TestFile{})
	needMigrateFiles := getNeedMigrateFiles(files, math.MaxInt64)
	assert.NotNil(t, needMigrateFiles)
	assert.Equal(t, len(files), len(needMigrateFiles))
	// 手动插入一条已经迁移的数据
	m := Migration{
		Migration: files[0].Key(),
		Batch:     1,
	}
	DB.Save(&m)

	// 此时没有需要迁移的数据了
	needMigrateFiles = getNeedMigrateFiles(files, math.MaxInt64)
	assert.Equal(t, 0, len(needMigrateFiles))
}

func TestGetNeedRollbackKeys(t *testing.T) {
	assert.Nil(t, DelAll())
	assert.Nil(t, InitMigrationTable())
	Register(&testdata.TestFile{})
	var needRollbackMs []File
	needRollbackMs = getNeedRollbackKeys(1)
	assert.Equal(t, 0, len(needRollbackMs))

	m := Migration{
		Migration: files[0].Key(),
		Batch:     1,
	}
	gdb.Def().Create(&m)

	needRollbackMs = getNeedRollbackKeys(1)
	assert.Equal(t, 1, len(needRollbackMs))
	assert.Equal(t, files[0].Key(), needRollbackMs[0].Key())
}

func TestGetNextBatchNo(t *testing.T) {
	assert.Nil(t, DelAll())
	assert.Nil(t, InitMigrationTable())
	Register(&testdata.TestFile{})
	var nextBatch uint
	nextBatch = getNextBatchNo()
	assert.Equal(t, uint(1), nextBatch)
	m := Migration{
		Migration: files[0].Key(),
		Batch:     nextBatch,
	}
	gdb.Def().Create(&m)

	nextBatch = getNextBatchNo()
	assert.Equal(t, uint(2), nextBatch)
}

func TestTables(t *testing.T) {
	assert.Nil(t, DelAll())
	assert.Nil(t, InitMigrationTable())
	tables, err := Tables()
	assert.Nil(t, err)
	assert.True(t, len(tables) > 0)
}

func TestDel(t *testing.T) {
	assert.Nil(t, DelAll())
	assert.Nil(t, InitMigrationTable())
	err := DB.Exec("create table test (id int)").Error
	assert.Nil(t, err)
	assert.Nil(t, err)
	exist, err := TableExist("test")
	assert.Nil(t, err)
	assert.True(t, exist)
	err = Del("test")
	assert.Nil(t, err)
	exist, err = TableExist("test")
	assert.Nil(t, err)
	assert.False(t, exist)
}

func TestMigrate(t *testing.T) {
	assert.Nil(t, DelAll())
	assert.Nil(t, InitMigrationTable())
	exist, err := TableExist("test")
	assert.Nil(t, err)
	assert.False(t, exist)
	Register(&testdata.TestFile{})
	// Migrate
	assert.Nil(t, Migrate(1))
	exist, err = TableExist("test")
	assert.Nil(t, err)
	assert.True(t, exist)

	// Rollback
	assert.Nil(t, Rollback(1))
	exist, err = TableExist("test")
	assert.Nil(t, err)
	assert.False(t, exist)

	assert.True(t, len(getNeedRollbackKeys(1)) == 0)
}

func TestFresh(t *testing.T) {
	assert.Nil(t, DelAll())
	assert.Nil(t, InitMigrationTable())
	exist, err := TableExist("test")
	assert.Nil(t, err)
	assert.False(t, exist)
	Register(&testdata.TestFile{})

	// Migrate
	assert.Nil(t, Migrate(1))
	exist, err = TableExist("test")
	assert.Nil(t, err)
	assert.True(t, exist)

	// Fresh

	assert.Nil(t, Fresh())
	exist, err = TableExist("test")
	assert.Nil(t, err)
	assert.True(t, exist)
}

func TestTruncate(t *testing.T) {
	assert.Nil(t, DelAll())
	assert.Nil(t, InitMigrationTable())
	exist, err := TableExist("test")
	assert.Nil(t, err)
	assert.False(t, exist)
	Register(&testdata.TestFile{})

	// Migrate
	assert.Nil(t, Migrate(1))
	exist, err = TableExist("test")
	assert.Nil(t, err)
	assert.True(t, exist)

	test := testdata.Test{
		ID: 1,
	}
	assert.Nil(t, DB.Create(&test).Error)
	test2 := testdata.Test{}
	assert.Nil(t, DB.First(&test2).Error)
	assert.True(t, test2.ID != 0)

	// 清空表数据
	assert.Nil(t, Truncate("test"))
	test3 := testdata.Test{}
	err = DB.First(&test3).Error
	t.Logf("%+v", err)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
	assert.True(t, test3.ID == 0)
}
