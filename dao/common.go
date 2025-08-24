package dao

import (
    "errors"
    "fmt"

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
    if s.other.Offset >= 0 && s.other.Limit > 0 {
        err = DB.Count(&total).Error
        if err != nil {
            return
        }
    }
    DB = DB.Scopes(OtherScope(s.other))
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
    DB = DB.Scopes(OtherScope(s.other))
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
