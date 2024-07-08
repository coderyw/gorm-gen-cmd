// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.
// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/coderyw/easyjson"
	jlexer "github.com/coderyw/easyjson/jlexer"
	jwriter "github.com/coderyw/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjsonC80ae7adDecodeGithubComCoderywGormGenCmdModel(in *jlexer.Lexer, out *GenCfg) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "host":
			out.Host = string(in.String())
		case "port":
			out.Port = string(in.String())
		case "database":
			out.Database = string(in.String())
		case "auth":
			out.Auth = string(in.String())
		case "tables":
			if in.IsNull() {
				in.Skip()
				out.Tables = nil
			} else {
				in.Delim('[')
				if out.Tables == nil {
					if !in.IsDelim(']') {
						out.Tables = make([]string, 0, 4)
					} else {
						out.Tables = []string{}
					}
				} else {
					out.Tables = (out.Tables)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Tables = append(out.Tables, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "outpath":
			out.Outpath = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonC80ae7adEncodeGithubComCoderywGormGenCmdModel(out *jwriter.Writer, in GenCfg) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"host\":"
		out.RawString(prefix[1:])
		out.String(string(in.Host))
	}
	{
		const prefix string = ",\"port\":"
		out.RawString(prefix)
		out.String(string(in.Port))
	}
	{
		const prefix string = ",\"database\":"
		out.RawString(prefix)
		out.String(string(in.Database))
	}
	{
		const prefix string = ",\"auth\":"
		out.RawString(prefix)
		out.String(string(in.Auth))
	}
	{
		const prefix string = ",\"tables\":"
		out.RawString(prefix)
		if in.Tables == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Tables {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"outpath\":"
		out.RawString(prefix)
		out.String(string(in.Outpath))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v *GenCfg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonC80ae7adEncodeGithubComCoderywGormGenCmdModel(&w, *v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v *GenCfg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonC80ae7adEncodeGithubComCoderywGormGenCmdModel(w, *v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *GenCfg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonC80ae7adDecodeGithubComCoderywGormGenCmdModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *GenCfg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonC80ae7adDecodeGithubComCoderywGormGenCmdModel(l, v)
}
