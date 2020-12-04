package interfaces


// 缓存接口
type Cache interface {
	// Set 设置缓存(键、值。缓存时间)
	Set(string, interface{}, int64) error
	// Get 获取缓存
	Get(string) (interface{}, error)
	//
	Del(key string) bool
	//
	Exist(key string) bool
}
