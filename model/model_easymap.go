package model

import (
	fmt "fmt"
	unsafe "unsafe"
)

func str2Bytes_model_easymap(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
func (v *GenCfg) UnMarshalMap(m map[string]string) error {
	var (
		ok  bool
		val string
	)
	if val, ok = m["host"]; ok {
		{
			pv := val
			v.Host = string(pv)
		}
	}
	if val, ok = m["port"]; ok {
		{
			pv := val
			v.Port = string(pv)
		}
	}
	if val, ok = m["database"]; ok {
		{
			pv := val
			v.Database = string(pv)
		}
	}
	if val, ok = m["auth"]; ok {
		{
			pv := val
			v.Auth = string(pv)
		}
	}

	return nil
}
func (v *GenCfg) UnMarshalMapInterface(m map[string]interface{}) error {
	var (
		ok  bool
		val interface{}
	)
	if val, ok = m["host"]; ok {
		switch val.(type) {
		case string:
			v.Host = string(val.(string))
		}
	}
	if val, ok = m["port"]; ok {
		switch val.(type) {
		case string:
			v.Port = string(val.(string))
		}
	}
	if val, ok = m["database"]; ok {
		switch val.(type) {
		case string:
			v.Database = string(val.(string))
		}
	}
	if val, ok = m["auth"]; ok {
		switch val.(type) {
		case string:
			v.Auth = string(val.(string))
		}
	}
	if val, ok = m["outpath"]; ok {
		switch val.(type) {
		case string:
			v.Outpath = string(val.(string))
		}
	}

	return nil
}

func (v *GenCfg) MarshalMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	m["host"] = v.Host
	m["port"] = v.Port
	m["database"] = v.Database
	m["auth"] = v.Auth
	m["outpath"] = v.Outpath
	return m, nil
}

func (v *GenCfg) MarshalMapString() (map[string]string, error) {
	m := make(map[string]string)
	m["host"] = fmt.Sprint(v.Host)
	m["port"] = fmt.Sprint(v.Port)
	m["database"] = fmt.Sprint(v.Database)
	m["auth"] = fmt.Sprint(v.Auth)
	m["outpath"] = fmt.Sprint(v.Outpath)
	return m, nil
}

type GenCfgField string

func (v GenCfgField) MarshalBinary() (data []byte, err error) {
	return str2Bytes_model_easymap(string(v)), nil
}

const (
	GenCfg_Host     GenCfgField = "host"
	GenCfg_Port     GenCfgField = "port"
	GenCfg_Database GenCfgField = "database"
	GenCfg_Auth     GenCfgField = "auth"
	GenCfg_Tables   GenCfgField = "tables"
	GenCfg_Outpath  GenCfgField = "outpath"
)
