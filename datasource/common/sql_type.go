package common

import (
	"database/sql/driver"
	"fmt"
	"time"
)

/*    **************** ***** * ****** * ***** * ***** * ***** * ******************    */
/********************** *** *** *** **** *** *** *** *** *** *** **********************/
/*    ****************** * ***** * ****** * ***** * ***** * ***** ****************    */

/*
 * 定义新的时间类型
 *  MarshalJSON 控制该时间类型的json格式输出
 *  Value 写入数据库的类型和数值
 *  Scan 从数据库取出
 */

// 1. 创建 time.Time 类型的副本 XTime；
type XTime struct {
	time.Time
}

// 2. 为 Xtime 重写 MarshaJSON 方法，在此方法中实现自定义格式的转换；
func (t XTime) MarshalJSON() ([]byte, error) {
	output := fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))
	return []byte(output), nil
}

// 3. 为 Xtime 实现 Value 方法，写入数据库时会调用该方法将自定义时间类型转换并写入数据库；
func (t XTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// 4. 为 Xtime 实现 Scan 方法，读取数据库时会调用该方法将时间数据转换成自定义时间类型；
func (t *XTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = XTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

/*    **************** ***** * ****** * ***** * ***** * ***** * ******************    */
/********************** *** *** *** **** *** *** *** *** *** *** **********************/
/*    ****************** * ***** * ****** * ***** * ***** * ***** ****************    */
