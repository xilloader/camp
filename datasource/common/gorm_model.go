package common

type Model struct {
	ID        int64       `json:"id" mapstructure:"ID" gorm:"primary_key;"`
	CreatedAt camp.XTime  `json:"created_at"`
	UpdatedAt camp.XTime  `json:"updated_at"`
	DeletedAt *camp.XTime `json:"-" sql:"index"`
}

type QueryCond struct {
	Order string
	Filter
}

func NewQueryCond(order string, filter Filter) QueryCond {
	return QueryCond{
		Order:  order,
		Filter: filter,
	}
}

type Filter struct {
	Where string
	Value []interface{}
}

func NewWhere(where string, val ...interface{}) *Filter {
	return &Filter{
		Where: where,
		Value: val,
	}
}

func (w *Filter) And(where string, val ...interface{}) *Filter {
	if w.Where == "" {
		w.Where = where
	} else {
		w.Where += " and " + where
	}
	w.Value = append(w.Value, val...)
	return w
}

func (w *Filter) Or(where string, val ...interface{}) *Filter {
	w.Where += " or " + where
	w.Value = append(w.Value, val...)
	return w
}