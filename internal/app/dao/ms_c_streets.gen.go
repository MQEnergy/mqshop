// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/MQEnergy/mqshop/internal/app/model"
)

func newCStreets(db *gorm.DB, opts ...gen.DOOption) cStreets {
	_cStreets := cStreets{}

	_cStreets.cStreetsDo.UseDB(db, opts...)
	_cStreets.cStreetsDo.UseModel(&model.CStreets{})

	tableName := _cStreets.cStreetsDo.TableName()
	_cStreets.ALL = field.NewAsterisk(tableName)
	_cStreets.ID = field.NewInt64(tableName, "id")
	_cStreets.Code = field.NewString(tableName, "code")
	_cStreets.Name = field.NewString(tableName, "name")
	_cStreets.AreaCode = field.NewString(tableName, "area_code")
	_cStreets.ProvinceCode = field.NewString(tableName, "province_code")
	_cStreets.CityCode = field.NewString(tableName, "city_code")

	_cStreets.fillFieldMap()

	return _cStreets
}

type cStreets struct {
	cStreetsDo

	ALL          field.Asterisk
	ID           field.Int64
	Code         field.String
	Name         field.String
	AreaCode     field.String
	ProvinceCode field.String
	CityCode     field.String

	fieldMap map[string]field.Expr
}

func (c cStreets) Table(newTableName string) *cStreets {
	c.cStreetsDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cStreets) As(alias string) *cStreets {
	c.cStreetsDo.DO = *(c.cStreetsDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cStreets) updateTableName(table string) *cStreets {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.Code = field.NewString(table, "code")
	c.Name = field.NewString(table, "name")
	c.AreaCode = field.NewString(table, "area_code")
	c.ProvinceCode = field.NewString(table, "province_code")
	c.CityCode = field.NewString(table, "city_code")

	c.fillFieldMap()

	return c
}

func (c *cStreets) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cStreets) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 6)
	c.fieldMap["id"] = c.ID
	c.fieldMap["code"] = c.Code
	c.fieldMap["name"] = c.Name
	c.fieldMap["area_code"] = c.AreaCode
	c.fieldMap["province_code"] = c.ProvinceCode
	c.fieldMap["city_code"] = c.CityCode
}

func (c cStreets) clone(db *gorm.DB) cStreets {
	c.cStreetsDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cStreets) replaceDB(db *gorm.DB) cStreets {
	c.cStreetsDo.ReplaceDB(db)
	return c
}

type cStreetsDo struct{ gen.DO }

type ICStreetsDo interface {
	gen.SubQuery
	Debug() ICStreetsDo
	WithContext(ctx context.Context) ICStreetsDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICStreetsDo
	WriteDB() ICStreetsDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICStreetsDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICStreetsDo
	Not(conds ...gen.Condition) ICStreetsDo
	Or(conds ...gen.Condition) ICStreetsDo
	Select(conds ...field.Expr) ICStreetsDo
	Where(conds ...gen.Condition) ICStreetsDo
	Order(conds ...field.Expr) ICStreetsDo
	Distinct(cols ...field.Expr) ICStreetsDo
	Omit(cols ...field.Expr) ICStreetsDo
	Join(table schema.Tabler, on ...field.Expr) ICStreetsDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICStreetsDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICStreetsDo
	Group(cols ...field.Expr) ICStreetsDo
	Having(conds ...gen.Condition) ICStreetsDo
	Limit(limit int) ICStreetsDo
	Offset(offset int) ICStreetsDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICStreetsDo
	Unscoped() ICStreetsDo
	Create(values ...*model.CStreets) error
	CreateInBatches(values []*model.CStreets, batchSize int) error
	Save(values ...*model.CStreets) error
	First() (*model.CStreets, error)
	Take() (*model.CStreets, error)
	Last() (*model.CStreets, error)
	Find() ([]*model.CStreets, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CStreets, err error)
	FindInBatches(result *[]*model.CStreets, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CStreets) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICStreetsDo
	Assign(attrs ...field.AssignExpr) ICStreetsDo
	Joins(fields ...field.RelationField) ICStreetsDo
	Preload(fields ...field.RelationField) ICStreetsDo
	FirstOrInit() (*model.CStreets, error)
	FirstOrCreate() (*model.CStreets, error)
	FindByPage(offset int, limit int) (result []*model.CStreets, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICStreetsDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cStreetsDo) Debug() ICStreetsDo {
	return c.withDO(c.DO.Debug())
}

func (c cStreetsDo) WithContext(ctx context.Context) ICStreetsDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cStreetsDo) ReadDB() ICStreetsDo {
	return c.Clauses(dbresolver.Read)
}

func (c cStreetsDo) WriteDB() ICStreetsDo {
	return c.Clauses(dbresolver.Write)
}

func (c cStreetsDo) Session(config *gorm.Session) ICStreetsDo {
	return c.withDO(c.DO.Session(config))
}

func (c cStreetsDo) Clauses(conds ...clause.Expression) ICStreetsDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cStreetsDo) Returning(value interface{}, columns ...string) ICStreetsDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cStreetsDo) Not(conds ...gen.Condition) ICStreetsDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cStreetsDo) Or(conds ...gen.Condition) ICStreetsDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cStreetsDo) Select(conds ...field.Expr) ICStreetsDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cStreetsDo) Where(conds ...gen.Condition) ICStreetsDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cStreetsDo) Order(conds ...field.Expr) ICStreetsDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cStreetsDo) Distinct(cols ...field.Expr) ICStreetsDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cStreetsDo) Omit(cols ...field.Expr) ICStreetsDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cStreetsDo) Join(table schema.Tabler, on ...field.Expr) ICStreetsDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cStreetsDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICStreetsDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cStreetsDo) RightJoin(table schema.Tabler, on ...field.Expr) ICStreetsDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cStreetsDo) Group(cols ...field.Expr) ICStreetsDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cStreetsDo) Having(conds ...gen.Condition) ICStreetsDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cStreetsDo) Limit(limit int) ICStreetsDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cStreetsDo) Offset(offset int) ICStreetsDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cStreetsDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICStreetsDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cStreetsDo) Unscoped() ICStreetsDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cStreetsDo) Create(values ...*model.CStreets) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cStreetsDo) CreateInBatches(values []*model.CStreets, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cStreetsDo) Save(values ...*model.CStreets) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cStreetsDo) First() (*model.CStreets, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CStreets), nil
	}
}

func (c cStreetsDo) Take() (*model.CStreets, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CStreets), nil
	}
}

func (c cStreetsDo) Last() (*model.CStreets, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CStreets), nil
	}
}

func (c cStreetsDo) Find() ([]*model.CStreets, error) {
	result, err := c.DO.Find()
	return result.([]*model.CStreets), err
}

func (c cStreetsDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CStreets, err error) {
	buf := make([]*model.CStreets, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cStreetsDo) FindInBatches(result *[]*model.CStreets, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cStreetsDo) Attrs(attrs ...field.AssignExpr) ICStreetsDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cStreetsDo) Assign(attrs ...field.AssignExpr) ICStreetsDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cStreetsDo) Joins(fields ...field.RelationField) ICStreetsDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cStreetsDo) Preload(fields ...field.RelationField) ICStreetsDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cStreetsDo) FirstOrInit() (*model.CStreets, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CStreets), nil
	}
}

func (c cStreetsDo) FirstOrCreate() (*model.CStreets, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CStreets), nil
	}
}

func (c cStreetsDo) FindByPage(offset int, limit int) (result []*model.CStreets, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cStreetsDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cStreetsDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cStreetsDo) Delete(models ...*model.CStreets) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cStreetsDo) withDO(do gen.Dao) *cStreetsDo {
	c.DO = *do.(*gen.DO)
	return c
}