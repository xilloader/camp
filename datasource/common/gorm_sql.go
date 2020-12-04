package common

import (
	"errors"
	"github.com/xilloader/camp/datasource"
	"gorm.io/gorm"
)

var db *gorm.DB

func NewCommonDB() {
	db = datasource.GetDB()
}

// Create
func Create(value interface{}) error {
	return db.Create(value).Error
}

// Save
func Save(value interface{}) error {
	return db.Save(value).Error
}

// Updates
func Updates(where interface{}, value interface{}) error {
	return db.Model(where).Updates(value).Error
}

// Delete
func DeleteByModel(model interface{}) (count int64, err error) {
	db := db.Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByWhere(model, where interface{}) (count int64, err error) {
	db := db.Where(where).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByID(model interface{}, id uint64) (count int64, err error) {
	db := db.Where("id=?", id).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// Delete
func DeleteByIDS(model interface{}, ids []uint64) (count int64, err error) {
	db := db.Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

// First
func FirstByID(out interface{}, id int) (notFound bool, err error) {
	err = db.First(out, id).Error
	if err != nil {
		notFound = errors.Is(err, gorm.ErrRecordNotFound)
	}
	return
}

// First
func First(where interface{}, out interface{}) (notFound bool, err error) {
	err = db.Where(where).First(out).Error
	if err != nil {
		notFound = errors.Is(err, gorm.ErrRecordNotFound)
	}
	return
}

// Find
func Find(where interface{}, out interface{}, orders ...string) error {
	db := db.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

// Scan
func Scan(model, where interface{}, out interface{}) (notFound bool, err error) {
	err = db.Model(model).Where(where).Scan(out).Error
	if err != nil {
		notFound = errors.Is(err, gorm.ErrRecordNotFound)
	}
	return
}

// ScanList
func ScanList(model, where interface{}, out interface{}, orders ...string) error {
	db := db.Model(model).Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Scan(out).Error
}

// GetPage
func GetPage(model, where interface{}, out interface{}, pageIndex, pageSize int, totalCount *int64, whereOrder ...QueryCond) error {
	db := db.Model(model).Where(where)
	if len(whereOrder) > 0 {
		for _, wo := range whereOrder {
			if wo.Order != "" {
				db = db.Order(wo.Order)
			}
			if wo.Where != "" {
				db = db.Where(wo.Where, wo.Value...)
			}
		}
	}
	err := db.Count(totalCount).Error
	if err != nil {
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	return db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
}

// PluckList
func PluckList(model, where interface{}, out interface{}, fieldName string) error {
	return db.Model(model).Where(where).Pluck(fieldName, out).Error
}

func ResetErrRecordNotFound(err error) error {
	if errors.Is(gorm.ErrRecordNotFound, err) {
		return errors.New("查询结果不存在")
	}
	return err
}
