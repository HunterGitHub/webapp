package main

import (
	"reflect"
	"net/url"
	"strconv"
	"os"
	"bufio"
)

func HostMntr(host string) string {
	var allstr string
	f, err := os.OpenFile("/proc/stat", os.O_RDONLY, 0x755)
	checkErr(err, "open /proc/stat error")
	r := bufio.NewReader(f)
	line, err := r.ReadString('\n')
	allstr += line

	f, err = os.OpenFile("/proc/meminfo", os.O_RDONLY, 0x755)
	checkErr(err, "open /proc/stat error")
	r = bufio.NewReader(f)
	line, err = r.ReadString('\n')
	allstr += line

	f, err = os.OpenFile("/proc/uptime", os.O_RDONLY, 0x755)
	checkErr(err, "open /proc/stat error")
	r = bufio.NewReader(f)
	line, err = r.ReadString('\n')
	allstr += line
	return allstr
}

func structToMap(i interface{}) (values url.Values) {
	values = url.Values{}
	iVal := reflect.ValueOf(i).Elem()
	typ := iVal.Type()
	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		// You ca use tags here...
		// tag := typ.Field(i).Tag.Get("tagname")
		// Convert each type into a string for the url.Values string map
		var v string
		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}
		values.Set(typ.Field(i).Name, v)
	}
	return
}
