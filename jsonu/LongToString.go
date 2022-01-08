package jsonu

import (
	"bytes"
	"encoding/json"
	"strconv"
)

var JsMaxNum int64 = 1 << 52

func StringToLong(body []byte) (bs []byte, err error) {
	//解析json
	m1 := make(map[string]interface{})
	d := json.NewDecoder(bytes.NewBuffer(body))
	d.UseNumber()
	err = d.Decode(&m1)
	if err != nil {
		return
	}
	err = stol(m1)
	if err != nil {
		return
	}
	bs, err = json.Marshal(m1)
	return
}

//map 就调用递归调用赋值
//array 就判断是否是 map,是map同上
//不是 map 就是判断是否是 str 是就 stol
//str 类型就调用 stol
func stol(m1 map[string]interface{}) (err error) {
	for k, v := range m1 {
		//map 就调用递归调用赋值
		m2, okMap := v.(map[string]interface{})
		if okMap {
			stol(m2)
			m1[k] = m2
			continue
		}
		//array 就判断是否是 map,是map递归转换,赋值转换的map,最后赋值 arr
		//不是map就是基本数据类型,赋值转换后的基本数据类型
		arr2, okArr := v.([]interface{})
		if okArr {
			for i2, v2 := range arr2 {
				m3, okMap := v2.(map[string]interface{})
				if okMap {
					stol(m3)
					arr2[i2] = m3
					continue
				}
				jns, ok := v2.(string)
				if ok {
					arr2[i2] = strToInt64(jns)
				}
			}
			m1[k] = arr2
			continue
		}
		jns, ok := v.(string)
		if ok {
			m1[k] = strToInt64(jns)
		}
	}
	return
}

//如果 str 是 int64 且比较大,就转为 int64 否则原样返回
func strToInt64(str string) (resp interface{}) {
	i, err := strconv.ParseInt(str, 10, 64)
	if err == nil && i > JsMaxNum {
		resp = i
	} else {
		resp = str
	}
	return
}

// //解析 json 将前端无法解析的 long 转换为 string 类型 多层
func LongToString(body []byte) (bs []byte, err error) {
	//解析json
	m1 := make(map[string]interface{})
	d := json.NewDecoder(bytes.NewBuffer(body))
	d.UseNumber()
	err = d.Decode(&m1)
	if err != nil {
		return
	}
	err = ltos(m1)
	if err != nil {
		return
	}
	bs, err = json.Marshal(m1)
	return
}

//map 就调用递归调用赋值
//array 就判断是否是 map,是map同上
//不是 map 就是基本数据类型 num 转 string
func ltos(m1 map[string]interface{}) (err error) {
	for k, v := range m1 {
		//map 就调用递归调用赋值
		m2, okMap := v.(map[string]interface{})
		if okMap {
			ltos(m2)
			m1[k] = m2
			continue
		}
		//array 就判断是否是 map,是map递归转换,赋值转换的map,最后赋值 arr
		//不是map就是基本数据类型,赋值转换后的基本数据类型
		arr2, okArr := v.([]interface{})
		if okArr {
			for i2, v2 := range arr2 {
				m3, okMap := v2.(map[string]interface{})
				if okMap {
					ltos(m3)
					arr2[i2] = m3
					continue
				}
				jn, ok := v2.(json.Number)
				if ok {
					arr2[i2] = int64ToStr(jn)
				}
			}
			m1[k] = arr2
			continue
		}
		jn, ok := v.(json.Number)
		if ok {
			m1[k] = int64ToStr(jn)
		}
	}
	return
}

// 过大的 int64 就转 string
// 小的 int64 保持 int64
// float64 保持 float64
func int64ToStr(jn json.Number) (resp interface{}) {
	jni, err := jn.Int64()
	if err == nil {
		// 过大的 int64 就转 string
		if jni > JsMaxNum {
			return jn.String()
		} else {
			// 小的 int64 保持 int64
			return jni
		}
	}
	// float64 保持 float64
	_, err = jn.Float64()
	if err == nil {
		jns := jn.String()
		jnf, err := strconv.ParseFloat(jns, 64)
		if err != nil {
			return jn
		}
		return jnf
	}
	return jn
}
