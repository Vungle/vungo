// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package openrtb

import (
	json "encoding/json"
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

func easyjson9478868cDecodeGithubComVungleVungoOpenrtb(in *jlexer.Lexer, out *Metric) {
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
			out.Type = string(in.String())
		case "value":
			out.Value = float64(in.Float64())
		case "vendor":
			out.Vendor = string(in.String())
		case "ext":
			if m, ok := out.Ext.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.Ext.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.Ext = in.Interface()
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
func easyjson9478868cEncodeGithubComVungleVungoOpenrtb(out *jwriter.Writer, in Metric) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"type\":"
		out.RawString(prefix[1:])
		out.String(string(in.Type))
	}
	if in.Value != 0 {
		const prefix string = ",\"value\":"
		out.RawString(prefix)
		out.Float64(float64(in.Value))
	}
	if in.Vendor != "" {
		const prefix string = ",\"vendor\":"
		out.RawString(prefix)
		out.String(string(in.Vendor))
	}
	if in.Ext != nil {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		if m, ok := in.Ext.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.Ext.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.Ext))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Metric) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9478868cEncodeGithubComVungleVungoOpenrtb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Metric) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9478868cEncodeGithubComVungleVungoOpenrtb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Metric) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9478868cDecodeGithubComVungleVungoOpenrtb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Metric) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9478868cDecodeGithubComVungleVungoOpenrtb(l, v)
}
