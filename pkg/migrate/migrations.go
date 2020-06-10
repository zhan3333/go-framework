package migrate

import (
	db2 "go-framework/pkg/db"
)

type Migration struct {
	Id        uint   `json:"id" gorm:"primary_key"`
	Migration string `json:"migration" gorm:"type:varchar(255)"`
	Batch     uint   `json:"batch"`
}

func (Migration) TableName() string {
	return "migrations"
}

type File interface {
	Key() string
	Up() error
	Down() error
}

// 定义的迁移文件需要在这里注册
var Files []File

// 注册迁移文件
func Register(file File) {
	Files = append(Files, file)
}

// 获取需要迁移的 migrateFiles
// files 有, migrations 里没有的数据
func GetNeedMigrateFiles(migrateFiles []File) []File {
	var ans []File
	var ms = GetAllMigrations()
	diff := map[string]string{}
	for _, migrateFile := range migrateFiles {
		diff[migrateFile.Key()] = ""
	}
	for _, migration := range ms {
		delete(diff, migration.Migration)
	}
	for _, migrateFile := range migrateFiles {
		if _, ok := diff[migrateFile.Key()]; ok {
			ans = append(ans, migrateFile)
		}
	}
	return ans
}

// 获取需要回滚的 migrateFiles
func GetNeedRollbackKeys(step int) []File {
	var ans []File
	var ms = GetAllMigrations()
	var keyMigrateFile = map[string]File{}
	if step < 1 {
		return ans
	}
	for _, migrateFile := range Files {
		keyMigrateFile[migrateFile.Key()] = migrateFile
	}
	cur := 0
	for _, migrate := range ms {
		if step < 1 {
			break
		}
		if m, ok := keyMigrateFile[migrate.Migration]; ok {
			ans = append(ans, m)
		}
		if int(migrate.Batch) != cur {
			step--
		}
	}
	return ans
}

// 获取所有迁移记录
func GetAllMigrations() []Migration {
	var ms []Migration
	db2.Def().Order("id desc").Find(&ms)
	return ms
}

// 获取下一个迁移版本号
func GetNextBatchNo() uint {
	m := Migration{}
	batch := uint(0)
	db2.Def().Order("batch desc").Select("batch").First(&m)
	batch = m.Batch + 1
	return batch
}

func CreateMigrate(migration string, batch uint) (err error) {
	m := Migration{
		Migration: migration,
		Batch:     batch,
	}
	err = db2.Def().Create(&m).Error
	return
}

func DeleteMigrate(migration string) (err error) {
	m := Migration{
		Migration: migration,
	}
	err = db2.Def().Where(&m).Delete(Migration{}).Error
	return
}
