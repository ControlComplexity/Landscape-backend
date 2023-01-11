package constant

var (
	jiangsuSlice  = []string{"苏州", "无锡", "常州", "镇江", "南京", "南通", "扬州", "泰州", "盐城", "淮安", "宿迁", "徐州", "连云港"}
	zhejiangSlice = []string{"杭州市", "宁波市", "温州市", "嘉兴市", "湖州市", "绍兴市", "金华市", "衢州市", "舟山市", "台州市", "丽水市"}
	shanghaiSlice = []string{}
)

func GetProvinceMap() map[string][]string {
	provinceMap := make(map[string][]string)
	provinceMap["上海"] = shanghaiSlice
	provinceMap["浙江"] = zhejiangSlice
	provinceMap["江苏"] = jiangsuSlice
	return provinceMap
}
