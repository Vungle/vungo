// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package request

import (
	json "encoding/json"
	native "github.com/Vungle/vungo/openrtb/native"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson794297d0DecodeGithubComVungleVungoOpenrtbNativeRequest(in *jlexer.Lexer, out *Data) {
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
		case "type":
			out.Type = native.DataAssetType(in.Int64())
		case "len":
			out.Len = int64(in.Int64())
		case "ext":
			if m, ok := out.Extension.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Extension.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Extension = in.Interface()
			}
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
func easyjson794297d0EncodeGithubComVungleVungoOpenrtbNativeRequest(out *jwriter.Writer, in Data) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Type))
	}
	if in.Len != 0 {
		const prefix string = ",\"len\":"
		out.RawString(prefix)
		out.Int64(int64(in.Len))
	}
	if in.Extension != nil {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		if m, ok := in.Extension.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Extension.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Extension))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Data) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGithubComVungleVungoOpenrtbNativeRequest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Data) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGithubComVungleVungoOpenrtbNativeRequest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Data) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGithubComVungleVungoOpenrtbNativeRequest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Data) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGithubComVungleVungoOpenrtbNativeRequest(l, v)
}
