package userdatabase

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
)

type User struct {
	Id         string   `json:"id"`
	Portnums   int      `json:"portnums"`
	Domainnums int      `json:"domainnums"`
	Ports      []Port   `json:"ports"`
	Domains    []Domain `json:"domains"`
	Ip         string   `json:"ip"`
}
type Port struct {
	Porto string `json:"porto"`
	Port  string `json:"port"`
}
type Domain struct {
	url string `json:"url"`
}

func findAll() []User {
	ps := make([]User, 0)
	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &ps) //反序列化
	return ps
}
func Create(id, natip string, portnums, domainnums int) { //保存容器ID和转发端口数量限制
	fields := make([]map[string]interface{}, 0)
	p1 := &User{
		Id:         id,
		Portnums:   portnums,
		Domainnums: domainnums,
		Ports:      []Port{},
		Domains:    []Domain{},
		Ip:         natip,
	}
	_, err := json.Marshal(p1) //序列化
	if err != nil {
		log.Fatal()
	}
	data, err := ioutil.ReadFile("./user.json") //讀取文件，存儲為map格式
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}
	//將結構體反射成map
	tp := reflect.TypeOf(p1).Elem()
	vp := reflect.ValueOf(p1).Elem()
	field := make(map[string]interface{}, 0)
	for i := 0; i < tp.NumField(); i++ {
		field1 := tp.Field(i)
		field2 := vp.Field(i)
		key := field1.Tag.Get("json")
		field[key] = field2.Interface()
	}
	fields = append(fields, field)
	out, err := json.Marshal(fields)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("./user.json", out, 0775)
}
func SearchUser(id string) *User { //取特定用戶信息
	product := &User{}
	p1 := &User{
		Id: id,
	}
	//反射转为map
	tp := reflect.TypeOf(p1).Elem()
	vp := reflect.ValueOf(p1).Elem()
	field := make(map[string]interface{}, 0)
	for i := 0; i < tp.NumField(); i++ {
		field1 := tp.Field(i)
		field2 := vp.Field(i)
		key := field1.Tag.Get("json")
		switch field2.Kind() {
		case reflect.Int:
			field[key] = float64(field2.Interface().(int))
		case reflect.Int8:
			field[key] = float64(field2.Interface().(int8))
		case reflect.Int16:
			field[key] = float64(field2.Interface().(int16))
		case reflect.Int32:
			field[key] = float64(field2.Interface().(int32))
		case reflect.Int64:
			field[key] = float64(field2.Interface().(int64))
		case reflect.Uint:
			field[key] = float64(field2.Interface().(uint))
		case reflect.Uint8:
			field[key] = float64(field2.Interface().(uint8))
		case reflect.Uint16:
			field[key] = float64(field2.Interface().(uint16))
		case reflect.Uint32:
			field[key] = float64(field2.Interface().(uint32))
		case reflect.Uint64:
			field[key] = float64(field2.Interface().(uint64))
		case reflect.Float32:
			field[key] = float64(field2.Interface().(float32))
		case reflect.Float64:
			field[key] = field2.Interface()
		default:
			field[key] = field2.Interface()
		}
	}
	_, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal(err)
	}
	fields := make([]map[string]interface{}, 0)
	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range fields {
		if item["user"] == field["user"] {
			field = item
			goto over
		}
	}
over:
	out, _ := json.Marshal(field)
	json.Unmarshal(out, &product)
	return product
}
func UpdateUserPortnums(id string, nums int) {
	p1 := &User{
		Id: id,
	}
	tp := reflect.TypeOf(p1).Elem()
	vp := reflect.ValueOf(p1).Elem()
	field := make(map[string]interface{}, 0)
	for i := 0; i < tp.NumField(); i++ {
		field1 := tp.Field(i)
		field2 := vp.Field(i)
		key := field1.Tag.Get("json")
		switch field2.Kind() {
		case reflect.Int:
			field[key] = float64(field2.Interface().(int))
		case reflect.Int8:
			field[key] = float64(field2.Interface().(int8))
		case reflect.Int16:
			field[key] = float64(field2.Interface().(int16))
		case reflect.Int32:
			field[key] = float64(field2.Interface().(int32))
		case reflect.Int64:
			field[key] = float64(field2.Interface().(int64))
		case reflect.Uint:
			field[key] = float64(field2.Interface().(uint))
		case reflect.Uint8:
			field[key] = float64(field2.Interface().(uint8))
		case reflect.Uint16:
			field[key] = float64(field2.Interface().(uint16))
		case reflect.Uint32:
			field[key] = float64(field2.Interface().(uint32))
		case reflect.Uint64:
			field[key] = float64(field2.Interface().(uint64))
		case reflect.Float32:
			field[key] = float64(field2.Interface().(float32))
		case reflect.Float64:
			field[key] = field2.Interface()
		default:
			field[key] = field2.Interface()
		}
	}
	_, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal(err)
	}
	fields := make([]map[string]interface{}, 0)
	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range fields {
		if v["id"] == field["id"] {
			v["portnums"] = nums
		}
	}
	out, _ := json.MarshalIndent(fields, "", " ")
	_ = ioutil.WriteFile("./user.json", out, 0755)
}

func UpdateUserProt(id, proto, prot string) bool {
	p1 := &User{
		Id: id,
	}
	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal(err)
	}
	fields := make([]User, 0)
	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}
	var p2 []User
	for _, v := range fields { //检查对外端口
		for _, vv := range v.Ports {
			if vv.Porto == proto {
				return false
			}
		}
	}
	for _, v := range fields {
		if v.Id == p1.Id {
			if len(v.Ports) < v.Portnums {
				if len(v.Ports) == 0 {
					goto change
				}
				for _, vv := range v.Ports {
					if vv.Porto == proto {
						return false
					}
				}
			change:
				v.Ports = append(v.Ports, Port{proto, prot})
			}
		}
		p2 = append(p2, v)
	}
	out, err := json.MarshalIndent(p2, "", " ")
	_ = ioutil.WriteFile("./user.json", out, 0755)
	return true
}
func DeleteUserProt(id, proto string) bool {
	p1 := &User{
		Id: id,
	}
	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal(err)
	}
	fields := make([]User, 0)
	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}
	var p2 []User
	for _, v := range fields {
		if v.Id == p1.Id {
			length := len(v.Ports)
			for index, vv := range v.Ports {
				if vv.Porto == proto {
					if index == length-1 {
						v.Ports = v.Ports[0:index]
					} else {
						v.Ports = append(v.Ports[0:index], v.Ports[index+1:]...)
					}
				}
			}
		}
		p2 = append(p2, v)
	}
	out, err := json.MarshalIndent(p2, "", " ")
	_ = ioutil.WriteFile("./user.json", out, 0755)
	return true
}
func UpdateUserDomain(id, url string) bool {
	p1 := &User{
		Id: id,
	}
	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal(err)
	}
	fields := make([]User, 0)
	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}
	var p2 []User
	for _, v := range fields { //检查对外端口
		for _, vv := range v.Domains {
			if vv.url == url {
				return false
			}
		}
	}
	for _, v := range fields {
		if v.Id == p1.Id {
			if len(v.Domains) < v.Domainnums {
				if len(v.Domains) == 0 {
					goto change
				}
				for _, vv := range v.Domains {
					if vv.url == url {
						return false
					}
				}
			change:
				v.Domains = append(v.Domains, Domain{url: url})
			}
		}
		p2 = append(p2, v)
	}
	out, err := json.MarshalIndent(p2, "", " ")
	_ = ioutil.WriteFile("./user.json", out, 0755)
	return true
}
func DeleteUserDomain(id, url string) bool {
	p1 := &User{
		Id: id,
	}
	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal(err)
	}
	fields := make([]User, 0)
	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}
	var p2 []User
	for _, v := range fields {
		if v.Id == p1.Id {
			length := len(v.Domains)
			for index, vv := range v.Domains {
				if vv.url == url {
					if index == length-1 {
						v.Domains = v.Domains[0:index]
					} else {
						v.Domains = append(v.Domains[0:index], v.Domains[index+1:]...)
					}
				}
			}
		}
		p2 = append(p2, v)
	}
	out, err := json.MarshalIndent(p2, "", " ")
	_ = ioutil.WriteFile("./user.json", out, 0755)
	return true
}
func UserPort(id string) []Port {
	p1 := &User{
		Id: id,
	}
	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal(err)
	}
	fields := make([]User, 0)
	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range fields {
		if v.Id == p1.Id {
			return v.Ports
		}
	}
	return nil
}
func UserDoamin(id string) []Domain {
	p1 := &User{
		Id: id,
	}
	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal(err)
	}
	fields := make([]User, 0)
	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range fields {
		if v.Id == p1.Id {
			return v.Domains
		}
	}
	return nil
}

func Okip(ip_prefix string) string { //返回一个可用的ip
	data, err := ioutil.ReadFile("./user.json")
	if err != nil {
		log.Fatal(err)
	}
	fields := make([]User, 0)
	err = json.Unmarshal(data, &fields)
	if err != nil {
		log.Fatal(err)
	}
	for i := 2; i <= 254; i++ {
		if len(fields) == 0 {
			return ip_prefix + "." + strconv.Itoa(i)
		}
		for fieldslen, v := range fields {
			if ip_prefix+"."+strconv.Itoa(i) != v.Ip && fieldslen+1 == len(fields) {
				return ip_prefix + "." + strconv.Itoa(i)
			}
		}
	}
	return ""
}
