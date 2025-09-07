package dao

import (
	"errors"
	"fmt"
	"math/rand"

	"bosh-admin/global"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var NotFound = gorm.ErrRecordNotFound

// GormDB 获取gorm.DB
func GormDB() *gorm.DB {
	return global.GormDB
}

// Query 原生sql查询
func Query(sql string, values ...interface{}) error {
	return GormDB().Raw(sql, values...).Error
}

// Exec 原生sql执行
func Exec(sql string, values ...interface{}) error {
	return GormDB().Exec(sql, values...).Error
}

// Begin 开启事务
func Begin() *gorm.DB {
	return GormDB().Begin()
}

func Create(value interface{}, table ...string) error {
	DB := GormDB()
	if len(table) > 0 {
		DB = DB.Table(table[0])
	}
	return DB.Create(value).Error
}

func Save(value interface{}, table ...string) error {
	DB := GormDB()
	if len(table) > 0 {
		DB = DB.Table(table[0])
	}
	return DB.Save(value).Error
}

func Updates(value interface{}, table ...string) error {
	DB := GormDB()
	if len(table) > 0 {
		DB = DB.Table(table[0])
	}
	return DB.Select("*").Updates(value).Error
}

func Expr(query string, args ...interface{}) clause.Expr {
	return gorm.Expr(query, args...)
}

func QueryList[T any](s *Statement) (data []T, total int64, err error) {
	model := new(T)
	DB := s.Format().Model(model)
	if s.tableName != "" {
		DB = DB.Table(s.tableName)
	}
	s.fields.Range(func(query interface{}, args []interface{}) bool {
		DB = DB.Select(query, args...)
		return true
	})
	s.joins.Range(func(query string, args []interface{}) bool {
		DB = DB.Joins(query, args...)
		return true
	})
	if s.offset >= 0 && s.limit > 0 {
		err = DB.Count(&total).Error
		if err != nil {
			return
		}
	}
	DB = DB.Scopes(PageScope(s.limit, s.offset), OrderByScope(s.orderBy))
	s.preloads.Range(func(query string, args []interface{}) bool {
		DB = DB.Preload(query, args...)
		return true
	})
	err = DB.Find(&data).Error
	return
}

func QueryOne[T any](s *Statement) (data T, err error) {
	model := new(T)
	DB := s.Format().Model(model)
	if s.tableName != "" {
		DB = DB.Table(s.tableName)
	}
	s.fields.Range(func(query interface{}, args []interface{}) bool {
		DB = DB.Select(query, args...)
		return true
	})
	s.joins.Range(func(query string, args []interface{}) bool {
		DB = DB.Joins(query, args...)
		return true
	})
	DB = DB.Scopes(PageScope(1, 0), OrderByScope(s.orderBy))
	s.preloads.Range(func(query string, args []interface{}) bool {
		DB = DB.Preload(query, args...)
		return true
	})
	err = DB.First(&data).Error
	return
}

// QueryById 通过id查询
func QueryById[T any](id any) (data T, err error) {
	err = GormDB().First(&data, id).Error
	return
}

// DelById 通过id删除
func DelById[T any](id any) error {
	model := new(T)
	return GormDB().Delete(model, id).Error
}

// DelByIds 通过id批量删除
func DelByIds[T any](ids ...any) error {
	model := new(T)
	return GormDB().Delete(model, ids...).Error
}

// Count 统计数量
func Count[T any](s *Statement) (num int64, err error) {
	model := new(T)
	err = s.Format().Model(model).Count(&num).Error
	return
}

// Sum 求和
func Sum[T any](s *Statement) (num float64, err error) {
	model := new(T)
	fields := s.fields.Keys()
	if len(fields) == 0 || fields[0] == "" {
		err = errors.New("求和字段错误")
		return
	}
	field := fields[0]
	err = s.Format().Model(model).Pluck(fmt.Sprintf("COALESCE(SUM(%s), 0)", field), &num).Error
	return
}

// PageScope 分页作用域
func PageScope(limit, offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit <= 0 {
			limit = DefaultLimit
		}
		if offset >= 0 {
			db = db.Offset(offset).Limit(limit)
		}
		return db
	}
}

// OrderByScope 排序作用域
func OrderByScope(orderStr string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if orderStr != "" {
			db = db.Order(orderStr)
		} else {
			db = db.Order("id DESC")
		}
		return db
	}
}

// RandomOrderScope 随机作用域
func RandomOrderScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		switch db.Dialector.Name() {
		case "postgres", "sqlite":
			db = db.Order("RANDOM()")
		default: // mysql等
			db = db.Order("RAND()")
		}
		return db
	}
}

// SafeRandomOrderScope 事务安全随机作用域
func SafeRandomOrderScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Session(&gorm.Session{}).Scopes(RandomOrderScope())
	}
}

// OptimizedRandomOrderScope 大表随机作用域
func OptimizedRandomOrderScope(table interface{}, pkField ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// 确定主键字段名
		pk := "id"
		if len(pkField) > 0 && pkField[0] != "" {
			pk = pkField[0]
		}
		// 创建新会话避免污染原查询
		tx := db.Session(&gorm.Session{})
		// 动态设置表名或模型
		if name, ok := table.(string); ok {
			tx = tx.Table(name)
		} else {
			tx = tx.Model(table)
		}
		// 获取最大ID值
		var maxID int
		tx.Select("MAX(" + pk + ")").Scan(&maxID)
		if maxID <= 0 {
			// 如果表为空或主键非数字，回退到简单随机排序
			switch tx.Dialector.Name() {
			case "postgres", "sqlite":
				return tx.Order("RANDOM()")
			default: // mysql等
				return tx.Order("RAND()")
			}
		}
		// 生成随机ID范围查询
		randomID := rand.Intn(maxID)
		return tx.Where(pk+" >= ?", randomID).Order(pk).Limit(1)
	}
}
