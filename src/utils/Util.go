package utils

import (
	"fmt"
	"sort"
	"time"
)

/*****************************************************************************
*  function name : RemoveDuplicatesAndEmpty
*  params :
*  return :
*
*  description : Term去重去空
*
******************************************************************************/
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	sort.Strings(a)
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

/*****************************************************************************
*  function name : Merge
*  params :
*  return :
*
*  description : 求并集
*
******************************************************************************/
func Merge(a []DocIdInfo, b []DocIdInfo) ([]DocIdInfo, bool) {
	lena := len(a)
	lenb := len(b)
	c := make([]DocIdInfo, 0)

	ia := 0
	ib := 0
	//fmt.Printf("Lena : %v ======== Lenb : %v \n",lena,lenb)
	if lena == 0 && lenb == 0 {
		return nil, false
	}

	for ia < lena && ib < lenb {

		if a[ia].DocId == b[ib].DocId {
			c = append(c, a[ia])
			ia++
			ib++
			continue
		}

		if a[ia].DocId < b[ib].DocId {
			//	fmt.Printf("ia : %v ======== ib : %v \n",ia,ib)
			c = append(c, a[ia])
			ia++
		} else {
			c = append(c, b[ib])
			ib++
		}
	}

	if ia < lena {
		for ; ia < lena; ia++ {
			c = append(c, a[ia])
		}

	} else {
		for ; ib < lenb; ib++ {
			c = append(c, b[ib])
		}
	}

	if len(c) == 0 {
		return nil, false
	} else {
		return c, true
	}

}

/*****************************************************************************
*  function name : Interaction
*  params :
*  return :
*
*  description : 求交集
*
******************************************************************************/
func Interaction(a []DocIdInfo, b []DocIdInfo) ([]DocIdInfo, bool) {

	lena := len(a)
	lenb := len(b)
	c := make([]DocIdInfo, 0)

	ia := 0
	ib := 0
	for ia < lena && ib < lenb {

		if a[ia].DocId == b[ib].DocId {
			c = append(c, a[ia])
		}

		if a[ia].DocId < b[ib].DocId {
			ia++
		} else {
			ib++
		}
	}

	if len(c) == 0 {
		return nil, false
	} else {
		return c, true
	}

}


func InitTime() func(string) string {
	init_time := time.Now()

	return func(description string) string {
		now := time.Now()
		cost := fmt.Sprintf("[ %v ] cost:%v", description, now.Sub(init_time))
		init_time = now
		return cost

	}

}

//自定义接口..用于外部写查询插件
type CustomInterface interface {
	CustomeFunction(v1, v2 interface{}) bool
	Init() bool 
	SetRules(rules interface{}) func(value_byte interface{}) bool
	//插件分词函数,返回string数组,bool参数表示是建立索引的时候还是查询的调用,STYPE = 9 调用
	SegmentFunc(value interface{},isSearch bool) []string
	//数字分词函数,返回string数组,bool参数表示是建立索引的时候还是查询的调用,STYPE = 9 调用
	SplitNum(value interface{}) int64
	//插件正排处理函数，建立索引的时候调用，stype =9 调用 ,返回byte数组
	BuildByteProfile(value []byte) []byte 
	//插件正排处理函数，建立索引的时候调用，stype =9 调用 ,返回string,定长！！！！
	BuildStringProfile(value interface{}) string
	//插件正排处理函数，建立索引的时候调用，stype =9 调用 ,返回int64
	BuildIntProfile(value interface{}) int64
}


type Condition struct {
	Key   string `json:"key"`
	Op    string `json:"operate"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
	Range string `json:"daterange"`
}
