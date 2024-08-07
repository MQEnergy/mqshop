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

func newProductGoodsSku(db *gorm.DB, opts ...gen.DOOption) productGoodsSku {
	_productGoodsSku := productGoodsSku{}

	_productGoodsSku.productGoodsSkuDo.UseDB(db, opts...)
	_productGoodsSku.productGoodsSkuDo.UseModel(&model.ProductGoodsSku{})

	tableName := _productGoodsSku.productGoodsSkuDo.TableName()
	_productGoodsSku.ALL = field.NewAsterisk(tableName)
	_productGoodsSku.ID = field.NewInt64(tableName, "id")
	_productGoodsSku.GoodsID = field.NewInt64(tableName, "goods_id")
	_productGoodsSku.Price = field.NewFloat64(tableName, "price")
	_productGoodsSku.Stock = field.NewInt(tableName, "stock")
	_productGoodsSku.StockWarning = field.NewInt(tableName, "stock_warning")
	_productGoodsSku.SkuNo = field.NewString(tableName, "sku_no")
	_productGoodsSku.ThumbURL = field.NewString(tableName, "thumb_url")
	_productGoodsSku.AttrValue = field.NewString(tableName, "attr_value")
	_productGoodsSku.AttrHash = field.NewString(tableName, "attr_hash")
	_productGoodsSku.CreatedAt = field.NewInt64(tableName, "created_at")
	_productGoodsSku.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_productGoodsSku.fillFieldMap()

	return _productGoodsSku
}

// productGoodsSku 商品sku配置表
type productGoodsSku struct {
	productGoodsSkuDo

	ALL          field.Asterisk
	ID           field.Int64
	GoodsID      field.Int64   // 商品ID
	Price        field.Float64 // 销售价格
	Stock        field.Int     // 商品库存
	StockWarning field.Int     // 库存预警值
	SkuNo        field.String  // sku编码
	ThumbURL     field.String  // 图片
	AttrValue    field.String  // 属性值 {"1":"金色", "2":"16G"}
	AttrHash     field.String  // 属性值对应的hash值
	CreatedAt    field.Int64
	UpdatedAt    field.Int64

	fieldMap map[string]field.Expr
}

func (p productGoodsSku) Table(newTableName string) *productGoodsSku {
	p.productGoodsSkuDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p productGoodsSku) As(alias string) *productGoodsSku {
	p.productGoodsSkuDo.DO = *(p.productGoodsSkuDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *productGoodsSku) updateTableName(table string) *productGoodsSku {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.GoodsID = field.NewInt64(table, "goods_id")
	p.Price = field.NewFloat64(table, "price")
	p.Stock = field.NewInt(table, "stock")
	p.StockWarning = field.NewInt(table, "stock_warning")
	p.SkuNo = field.NewString(table, "sku_no")
	p.ThumbURL = field.NewString(table, "thumb_url")
	p.AttrValue = field.NewString(table, "attr_value")
	p.AttrHash = field.NewString(table, "attr_hash")
	p.CreatedAt = field.NewInt64(table, "created_at")
	p.UpdatedAt = field.NewInt64(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *productGoodsSku) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *productGoodsSku) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 11)
	p.fieldMap["id"] = p.ID
	p.fieldMap["goods_id"] = p.GoodsID
	p.fieldMap["price"] = p.Price
	p.fieldMap["stock"] = p.Stock
	p.fieldMap["stock_warning"] = p.StockWarning
	p.fieldMap["sku_no"] = p.SkuNo
	p.fieldMap["thumb_url"] = p.ThumbURL
	p.fieldMap["attr_value"] = p.AttrValue
	p.fieldMap["attr_hash"] = p.AttrHash
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p productGoodsSku) clone(db *gorm.DB) productGoodsSku {
	p.productGoodsSkuDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p productGoodsSku) replaceDB(db *gorm.DB) productGoodsSku {
	p.productGoodsSkuDo.ReplaceDB(db)
	return p
}

type productGoodsSkuDo struct{ gen.DO }

type IProductGoodsSkuDo interface {
	gen.SubQuery
	Debug() IProductGoodsSkuDo
	WithContext(ctx context.Context) IProductGoodsSkuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProductGoodsSkuDo
	WriteDB() IProductGoodsSkuDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProductGoodsSkuDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProductGoodsSkuDo
	Not(conds ...gen.Condition) IProductGoodsSkuDo
	Or(conds ...gen.Condition) IProductGoodsSkuDo
	Select(conds ...field.Expr) IProductGoodsSkuDo
	Where(conds ...gen.Condition) IProductGoodsSkuDo
	Order(conds ...field.Expr) IProductGoodsSkuDo
	Distinct(cols ...field.Expr) IProductGoodsSkuDo
	Omit(cols ...field.Expr) IProductGoodsSkuDo
	Join(table schema.Tabler, on ...field.Expr) IProductGoodsSkuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProductGoodsSkuDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProductGoodsSkuDo
	Group(cols ...field.Expr) IProductGoodsSkuDo
	Having(conds ...gen.Condition) IProductGoodsSkuDo
	Limit(limit int) IProductGoodsSkuDo
	Offset(offset int) IProductGoodsSkuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProductGoodsSkuDo
	Unscoped() IProductGoodsSkuDo
	Create(values ...*model.ProductGoodsSku) error
	CreateInBatches(values []*model.ProductGoodsSku, batchSize int) error
	Save(values ...*model.ProductGoodsSku) error
	First() (*model.ProductGoodsSku, error)
	Take() (*model.ProductGoodsSku, error)
	Last() (*model.ProductGoodsSku, error)
	Find() ([]*model.ProductGoodsSku, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProductGoodsSku, err error)
	FindInBatches(result *[]*model.ProductGoodsSku, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProductGoodsSku) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProductGoodsSkuDo
	Assign(attrs ...field.AssignExpr) IProductGoodsSkuDo
	Joins(fields ...field.RelationField) IProductGoodsSkuDo
	Preload(fields ...field.RelationField) IProductGoodsSkuDo
	FirstOrInit() (*model.ProductGoodsSku, error)
	FirstOrCreate() (*model.ProductGoodsSku, error)
	FindByPage(offset int, limit int) (result []*model.ProductGoodsSku, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProductGoodsSkuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p productGoodsSkuDo) Debug() IProductGoodsSkuDo {
	return p.withDO(p.DO.Debug())
}

func (p productGoodsSkuDo) WithContext(ctx context.Context) IProductGoodsSkuDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p productGoodsSkuDo) ReadDB() IProductGoodsSkuDo {
	return p.Clauses(dbresolver.Read)
}

func (p productGoodsSkuDo) WriteDB() IProductGoodsSkuDo {
	return p.Clauses(dbresolver.Write)
}

func (p productGoodsSkuDo) Session(config *gorm.Session) IProductGoodsSkuDo {
	return p.withDO(p.DO.Session(config))
}

func (p productGoodsSkuDo) Clauses(conds ...clause.Expression) IProductGoodsSkuDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p productGoodsSkuDo) Returning(value interface{}, columns ...string) IProductGoodsSkuDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p productGoodsSkuDo) Not(conds ...gen.Condition) IProductGoodsSkuDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p productGoodsSkuDo) Or(conds ...gen.Condition) IProductGoodsSkuDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p productGoodsSkuDo) Select(conds ...field.Expr) IProductGoodsSkuDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p productGoodsSkuDo) Where(conds ...gen.Condition) IProductGoodsSkuDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p productGoodsSkuDo) Order(conds ...field.Expr) IProductGoodsSkuDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p productGoodsSkuDo) Distinct(cols ...field.Expr) IProductGoodsSkuDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p productGoodsSkuDo) Omit(cols ...field.Expr) IProductGoodsSkuDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p productGoodsSkuDo) Join(table schema.Tabler, on ...field.Expr) IProductGoodsSkuDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p productGoodsSkuDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProductGoodsSkuDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p productGoodsSkuDo) RightJoin(table schema.Tabler, on ...field.Expr) IProductGoodsSkuDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p productGoodsSkuDo) Group(cols ...field.Expr) IProductGoodsSkuDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p productGoodsSkuDo) Having(conds ...gen.Condition) IProductGoodsSkuDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p productGoodsSkuDo) Limit(limit int) IProductGoodsSkuDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p productGoodsSkuDo) Offset(offset int) IProductGoodsSkuDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p productGoodsSkuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProductGoodsSkuDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p productGoodsSkuDo) Unscoped() IProductGoodsSkuDo {
	return p.withDO(p.DO.Unscoped())
}

func (p productGoodsSkuDo) Create(values ...*model.ProductGoodsSku) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p productGoodsSkuDo) CreateInBatches(values []*model.ProductGoodsSku, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p productGoodsSkuDo) Save(values ...*model.ProductGoodsSku) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p productGoodsSkuDo) First() (*model.ProductGoodsSku, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductGoodsSku), nil
	}
}

func (p productGoodsSkuDo) Take() (*model.ProductGoodsSku, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductGoodsSku), nil
	}
}

func (p productGoodsSkuDo) Last() (*model.ProductGoodsSku, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductGoodsSku), nil
	}
}

func (p productGoodsSkuDo) Find() ([]*model.ProductGoodsSku, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProductGoodsSku), err
}

func (p productGoodsSkuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProductGoodsSku, err error) {
	buf := make([]*model.ProductGoodsSku, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p productGoodsSkuDo) FindInBatches(result *[]*model.ProductGoodsSku, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p productGoodsSkuDo) Attrs(attrs ...field.AssignExpr) IProductGoodsSkuDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p productGoodsSkuDo) Assign(attrs ...field.AssignExpr) IProductGoodsSkuDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p productGoodsSkuDo) Joins(fields ...field.RelationField) IProductGoodsSkuDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p productGoodsSkuDo) Preload(fields ...field.RelationField) IProductGoodsSkuDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p productGoodsSkuDo) FirstOrInit() (*model.ProductGoodsSku, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductGoodsSku), nil
	}
}

func (p productGoodsSkuDo) FirstOrCreate() (*model.ProductGoodsSku, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductGoodsSku), nil
	}
}

func (p productGoodsSkuDo) FindByPage(offset int, limit int) (result []*model.ProductGoodsSku, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p productGoodsSkuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p productGoodsSkuDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p productGoodsSkuDo) Delete(models ...*model.ProductGoodsSku) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *productGoodsSkuDo) withDO(do gen.Dao) *productGoodsSkuDo {
	p.DO = *do.(*gen.DO)
	return p
}
