package dao

import (
	"github.com/duke-git/lancet/v2/maputil"
	"gorm.io/gorm"
)

const DefaultLimit = 10

// Statement 查询构造器
type Statement struct {
	db        *gorm.DB
	tableName string
	where     *maputil.OrderedMap[interface{}, []interface{}]
	fields    *maputil.OrderedMap[interface{}, []interface{}]
	joins     *maputil.OrderedMap[string, []interface{}]
	preloads  *maputil.OrderedMap[string, []interface{}]
	other     Other
}

// NewStatement 创建查询构造器
func NewStatement(other ...Other) *Statement {
	s := new(Statement)
	s.db = GormDB()
	s.tableName = ""
	s.where = maputil.NewOrderedMap[any, []any]()
	s.fields = maputil.NewOrderedMap[any, []any]()
	s.joins = maputil.NewOrderedMap[string, []any]()
	s.preloads = maputil.NewOrderedMap[string, []any]()
	if len(other) > 0 {
		s.other = other[0]
	} else {
		s.other = Other{}
	}
	return s
}

// Init 重置查询构造器
func (s *Statement) Init() {
	s.db = GormDB()
	s.tableName = ""
	s.where = maputil.NewOrderedMap[any, []any]()
	s.fields = maputil.NewOrderedMap[any, []any]()
	s.joins = maputil.NewOrderedMap[string, []any]()
	s.preloads = maputil.NewOrderedMap[string, []any]()
	s.other = Other{}
}

func (s *Statement) Table(tableName string) {
	s.tableName = tableName
}

// InitWhere 重置查询构造器where
func (s *Statement) InitWhere() {
	s.db = GormDB()
	s.where = maputil.NewOrderedMap[any, []any]()
}

// Where 查询条件
func (s *Statement) Where(query any, args ...any) {
	s.where.Set(query, args)
}

// DelWhere 删除查询条件
func (s *Statement) DelWhere(query any) {
	if _, ok := s.where.Get(query); ok {
		s.where.Delete(query)
	}
}

// Format 格式化查询条件
func (s *Statement) Format() *gorm.DB {
	s.where.Range(func(query any, args []any) bool {
		s.db = s.db.Where(query, args...)
		return true
	})
	return s.db
}

// Select 查询字段
func (s *Statement) Select(query any, args ...any) {
	s.fields.Set(query, args)
}

// Join 查询关联
func (s *Statement) Join(query string, args ...any) {
	if _, ok := s.joins.Get(query); !ok {
		s.joins.Set(query, args)
	}
}

// Preload 查询预加载
func (s *Statement) Preload(query string, args ...any) {
	if _, ok := s.preloads.Get(query); !ok {
		s.preloads.Set(query, args)
	}
}

// Pagination 分页
func (s *Statement) Pagination(pageNo, pageSize int) {
	if pageNo == -1 {
		s.other.Offset = -1
	} else {
		if pageSize <= 0 {
			s.other.Limit = DefaultLimit
		} else {
			s.other.Limit = pageSize
		}
		s.other.Offset = pageSize * (pageNo - 1)
	}
}

// OrderBy 排序
func (s *Statement) OrderBy(orderStr string) {
	if orderStr != "" {
		s.other.OrderBy = orderStr
	}
}

// Other 其他条件
func (s *Statement) Other(other Other) {
	s.other = other
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
		}
		return db
	}
}

// OtherScope 作用域
func OtherScope(other Other) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if other.OrderBy != "" {
			db = db.Order(other.OrderBy)
		}
		if other.Limit > 0 && other.Offset >= 0 {
			db = db.Offset(other.Offset).Limit(other.Limit)
		}
		return db
	}
}
