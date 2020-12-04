package strto

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

/*    **************** ***** * ****** * ***** * ***** * ***** * ******************    */
/********************** *** *** *** **** *** *** *** *** *** *** **********************/
/*    ****************** * ***** * ****** * ***** * ***** * ***** ****************    */
//NAME 随机生成人名
// 姓
var lastName = []string{
	"赵", "钱", "孙", "李", "周", "吴", "郑", "王", "冯", "陈", "褚", "卫", "蒋",
	"沈", "韩", "杨", "朱", "秦", "尤", "许", "何", "吕", "施", "张", "孔", "曹", "严", "华", "金", "魏",
	"陶", "姜", "戚", "谢", "邹", "喻", "柏", "水", "窦", "章", "云", "苏", "潘", "葛", "奚", "范", "彭",
	"郎", "鲁", "韦", "昌", "马", "苗", "凤", "花", "方", "任", "袁", "柳", "鲍", "史", "唐", "费", "薛",
	"雷", "贺", "倪", "汤", "滕", "殷", "罗", "毕", "郝", "安", "常", "傅", "卞", "齐", "元", "顾", "孟",
	"平", "黄", "穆", "萧", "尹", "姚", "邵", "湛", "汪", "祁", "毛", "狄", "米", "伏", "成", "戴", "谈",
	"宋", "茅", "庞", "熊", "纪", "舒", "屈", "项", "祝", "董", "梁", "杜", "阮", "蓝", "闵", "季", "贾",
	"路", "娄", "江", "童", "颜", "郭", "梅", "盛", "林", "钟", "徐", "邱", "骆", "高", "夏", "蔡", "田",
	"樊", "胡", "凌", "霍", "虞", "万", "支", "柯", "管", "卢", "莫", "柯", "房", "裘", "缪", "解", "应",
	"宗", "丁", "宣", "邓", "单", "杭", "洪", "包", "诸", "左", "石", "崔", "吉", "龚", "程", "嵇", "邢",
	"裴", "陆", "荣", "翁", "荀", "于", "惠", "甄", "曲", "封", "储", "仲", "伊", "宁", "仇", "甘", "武",
	"符", "刘", "景", "詹", "龙", "叶", "幸", "司", "黎", "溥", "印", "怀", "蒲", "邰", "从", "索", "赖",
	"卓", "屠", "池", "乔", "胥", "闻", "莘", "党", "翟", "谭", "贡", "劳", "逄", "姬", "申", "扶", "堵",
	"冉", "宰", "雍", "桑", "寿", "通", "燕", "浦", "尚", "农", "温", "别", "庄", "晏", "柴", "瞿", "阎",
	"连", "习", "容", "向", "古", "易", "廖", "庾", "终", "步", "都", "耿", "满", "弘", "匡", "国", "文",
	"寇", "广", "禄", "阙", "东", "欧", "利", "师", "巩", "聂", "关", "荆", "司马", "上官", "欧阳", "夏侯",
	"诸葛", "闻人", "东方", "赫连", "皇甫", "尉迟", "公羊", "澹台", "公冶", "宗政", "濮阳", "淳于", "单于",
	"太叔", "申屠", "公孙", "仲孙", "轩辕", "令狐", "徐离", "宇文", "长孙", "慕容", "司徒", "司空"}
var firstName = []string{
	"伟", "刚", "勇", "毅", "俊", "峰", "强", "军", "平", "保", "东", "文", "辉", "力", "明", "永", "健", "世", "广", "志", "义",
	"兴", "良", "海", "山", "仁", "波", "宁", "贵", "福", "生", "龙", "元", "全", "国", "胜", "学", "祥", "才", "发", "武", "新",
	"利", "清", "飞", "彬", "富", "顺", "信", "子", "杰", "涛", "昌", "成", "康", "星", "光", "天", "达", "安", "岩", "中", "茂",
	"进", "林", "有", "坚", "和", "彪", "博", "诚", "先", "敬", "震", "振", "壮", "会", "思", "群", "豪", "心", "邦", "承", "乐",
	"绍", "功", "松", "善", "厚", "庆", "磊", "民", "友", "裕", "河", "哲", "江", "超", "浩", "亮", "政", "谦", "亨", "奇", "固",
	"之", "轮", "翰", "朗", "伯", "宏", "言", "若", "鸣", "朋", "斌", "梁", "栋", "维", "启", "克", "伦", "翔", "旭", "鹏", "泽",
	"晨", "辰", "士", "以", "建", "家", "致", "树", "炎", "德", "行", "时", "泰", "盛", "雄", "琛", "钧", "冠", "策", "腾", "楠",
	"榕", "风", "航", "弘", "秀", "娟", "英", "华", "慧", "巧", "美", "娜", "静", "淑", "惠", "珠", "翠", "雅", "芝", "玉", "萍",
	"红", "娥", "玲", "芬", "芳", "燕", "彩", "春", "菊", "兰", "凤", "洁", "梅", "琳", "素", "云", "莲", "真", "环", "雪", "荣",
	"爱", "妹", "霞", "香", "月", "莺", "媛", "艳", "瑞", "凡", "佳", "嘉", "琼", "勤", "珍", "贞", "莉", "桂", "娣", "叶", "璧",
	"璐", "娅", "琦", "晶", "妍", "茜", "秋", "珊", "莎", "锦", "黛", "青", "倩", "婷", "姣", "婉", "娴", "瑾", "颖", "露", "瑶",
	"怡", "婵", "雁", "蓓", "纨", "仪", "荷", "丹", "蓉", "眉", "君", "琴", "蕊", "薇", "菁", "梦", "岚", "苑", "婕", "馨", "瑗",
	"琰", "韵", "融", "园", "艺", "咏", "卿", "聪", "澜", "纯", "毓", "悦", "昭", "冰", "爽", "琬", "茗", "羽", "希", "欣", "飘",
	"育", "滢", "馥", "筠", "柔", "竹", "霭", "凝", "晓", "欢", "霄", "枫", "芸", "菲", "寒", "伊", "亚", "宜", "可", "姬", "舒",
	"影", "荔", "枝", "丽", "阳", "妮", "宝", "贝", "初", "程", "梵", "罡", "恒", "鸿", "桦", "骅", "剑", "娇", "纪", "宽", "苛",
	"灵", "玛", "媚", "琪", "晴", "容", "睿", "烁", "堂", "唯", "威", "韦", "雯", "苇", "萱", "阅", "彦", "宇", "雨", "洋", "忠",
	"宗", "曼", "紫", "逸", "贤", "蝶", "菡", "绿", "蓝", "儿", "翠", "烟", "小", "轩"}
var lastNameLen = len(lastName)
var firstNameLen = len(firstName)

func GetFullName() string {
	rand.Seed(time.Now().UnixNano()) //设置随机数种子
	var first string
	times := rand.Intn(2)
	for i := 0; i <= times; i++ {
		first += fmt.Sprint(firstName[rand.Intn(firstNameLen-1)])
	}
	//返回姓名
	return fmt.Sprintf("%s%s", fmt.Sprint(lastName[rand.Intn(lastNameLen-1)]), first)
}

/*    **************** ***** * ****** * ***** * ***** * ***** * ******************    */
/********************** *** *** *** **** *** *** *** *** *** *** **********************/
/*    ****************** * ***** * ****** * ***** * ***** * ***** ****************    */
// Check string
// 校验密码
//字符串简单应用
const (
	//密码验证选项 只能含有
	PWD_OPT_Number  uint16 = 1 << iota //数字 1
	PWD_OPT_Lower                      //小写 2
	PWD_OPT_Upper                      //大写 4
	PWD_OPT_Special                    //特殊 8
)

//N：待判断的二进制数
//B：待判断的位（右往左）
//
//结果：((N>>(B-1))&1
type CheckProcess struct {
	min, max        uint16
	options, result uint16
	mustOpts        uint16 // 必选项
	expelRunes      string // 设置自定义禁用字符 $#!
	ErrMsg          string // 错误信息
}

// 设置密码检查选项
// 参数 min, max, options 有默认选项
// 参数 mustOpts, expel 为额外校验选项 不检查可设置为空值
func CheckPwd(pwd string, min, max, options, mustOpts uint16, expel string) bool {
	return NewCheckProcess(min, max, options, mustOpts, expel).VerifyPassword(pwd)
}

func NewCheckProcess(min, max, options, mustOpts uint16, expel string) *CheckProcess {
	if options == 0 {
		options = PWD_OPT_Number | PWD_OPT_Lower
	}
	if min < 6 {
		min = 6 // default minimum
	}
	if max < min {
		max = 20 // default maximum
	}
	if min > max {
		min, max = max, min
	}
	return &CheckProcess{
		min:        min,
		max:        max,
		options:    options | mustOpts,
		mustOpts:   mustOpts,
		expelRunes: expel,
	}
}

func (cp *CheckProcess) verify(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		switch {
		case strings.Contains(cp.expelRunes, string(r)):
			cp.ErrMsg = "contains expulsion rune"
			return false
		case unicode.IsNumber(r):
			cp.result = cp.result | PWD_OPT_Number
		case unicode.IsLower(r):
			cp.result = cp.result | PWD_OPT_Lower
		case unicode.IsUpper(r):
			cp.result = cp.result | PWD_OPT_Upper
		case unicode.IsPunct(r) || unicode.IsSymbol(r): //标点符号 和 字符
			cp.result = cp.result | PWD_OPT_Special
		default:
			cp.ErrMsg = "unknown character type"
			return false
		}
		// 当 cp.options&cp.result != cp.result 表示密码字符串超出 options 范围
		if cp.options&cp.result != cp.result {
			cp.ErrMsg = "actual option overflow"
			fmt.Printf("options: %4b\n result: %4b\n", cp.options, cp.result)
			return false
		}
	}
	// options 选项满足
	fmt.Printf("options: %4b\n result: %4b\n   must: %4b\n", cp.options, cp.result, cp.mustOpts)
	// 最后检查必选项
	end := cp.mustOpts&cp.result == cp.mustOpts
	if !end {
		cp.ErrMsg = "missing must options"
	}
	return end
}

func (cp *CheckProcess) VerifyPassword(s string) bool {
	l := len(s)
	if l < int(cp.min) || l > int(cp.max) {
		return false
	}
	return cp.verify(s)
}

func VerifyPwd(pwd string, options, must uint16) bool {
	if pwd == "" {
		return false
	}
	var result uint16
	for _, r := range pwd {
		switch {
		case unicode.IsNumber(r):
			result = result | PWD_OPT_Number
		case unicode.IsLower(r):
			result = result | PWD_OPT_Lower
		case unicode.IsUpper(r):
			result = result | PWD_OPT_Upper
		case unicode.IsPunct(r) || unicode.IsSymbol(r): //标点符号 和 字符
			result = result | PWD_OPT_Special
		default:
			return false
		}
		// 当 cp.options&cp.result != cp.result 表示密码字符串超出 options 范围
		if options&result != result {
			return false
		}
	}
	return must&result == must
}

/*    **************** ***** * ****** * ***** * ***** * ***** * ******************    */
/********************** *** *** *** **** *** *** *** *** *** *** **********************/
/*    ****************** * ***** * ****** * ***** * ***** * ***** ****************    */
var (
	ZHUnit               = []string{"个", "万", "亿"}
	SmallZHUnit          = []string{"千", "百", "十", "个"}
	UppercaseSmallZHUnit = []string{"仟", "佰", "什", "个"}
	UppercaseNumMap      = map[rune]string{'0': "零", '1': "壹", '2': "贰", '3': "叁", '4': "肆", '5': "肆", '6': "陆", '7': "柒", '8': "捌", '9': "玖"}
	LowercaseNumMap      = map[rune]string{'0': "零", '1': "一", '2': "二", '3': "三", '4': "四", '5': "四", '6': "六", '7': "七", '8': "八", '9': "九"}
)

func transfer(num int, upper bool) string {
	NumGroup := []int{}
	for num > 0 {
		group := num % 10000
		NumGroup = append(NumGroup, group)
		num = num / 10000
	}
	n := len(NumGroup)
	if n > 3 {
		return fmt.Sprintf("%d", num)
	}

	var smallUint = SmallZHUnit
	var CHNum = LowercaseNumMap
	if upper {
		smallUint = UppercaseSmallZHUnit
		CHNum = UppercaseNumMap
	}

	var ZHNum string
	for i, g := range NumGroup {
		if g == 0 {
			ZHNum = "零" + ZHNum
			continue
		}
		var chinese string
		//注意这里是倒序的
		for j, rg := range strconv.Itoa(g) {
			chinese += CHNum[rg] + smallUint[j]
		}
		//注意替换顺序
		chinese = strings.Replace(chinese, "零十", "零", 1)
		chinese = strings.Replace(chinese, "零百", "零", 1)
		chinese = strings.Replace(chinese, "零千", "零", 1)
		chinese = strings.Replace(chinese, "零零", "零", 1)
		chinese = strings.Replace(chinese, "零个", "", 1)

		ZHNum = chinese + ZHUnit[i] + ZHNum
		if ZHUnit[i] == "亿" {
			ZHNum = strings.Replace(ZHNum, "零亿", "亿", 1)
		} else if ZHUnit[i] == "万" {
			ZHNum = strings.Replace(ZHNum, "零万", "万", 1)
		} else {
			ZHNum = strings.Replace(ZHNum, "个个", "", 1)
			ZHNum = strings.Replace(ZHNum, "个", "", 1)
		}
	}

	return ZHNum
}
