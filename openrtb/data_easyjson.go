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

func easyjson794297d0DecodeGithubComVungleVungoOpenrtb(in *jlexer.Lexer, out *Data) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "segment":
			if in.IsNull() {
				in.Skip()
				out.Segment = nil
			} else {
				in.Delim('[')
				if out.Segment == nil {
					if !in.IsDelim(']') {
						out.Segment = make([]*Segment, 0, 8)
					} else {
						out.Segment = []*Segment{}
					}
				} else {
					out.Segment = (out.Segment)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Segment
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Segment)
						}
						easyjson794297d0DecodeGithubComVungleVungoOpenrtb1(in, v1)
					}
					out.Segment = append(out.Segment, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjson794297d0EncodeGithubComVungleVungoOpenrtb(out *jwriter.Writer, in Data) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if len(in.Segment) != 0 {
		const prefix string = ",\"segment\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.Segment {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					easyjson794297d0EncodeGithubComVungleVungoOpenrtb1(out, *v3)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Ext != nil {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
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
func (v Data) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson794297d0EncodeGithubComVungleVungoOpenrtb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Data) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson794297d0EncodeGithubComVungleVungoOpenrtb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Data) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson794297d0DecodeGithubComVungleVungoOpenrtb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Data) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson794297d0DecodeGithubComVungleVungoOpenrtb(l, v)
}
func easyjson794297d0DecodeGithubComVungleVungoOpenrtb1(in *jlexer.Lexer, out *Segment) {
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
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "value":
			out.Value = string(in.String())
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
func easyjson794297d0EncodeGithubComVungleVungoOpenrtb1(out *jwriter.Writer, in Segment) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Name))
	}
	if in.Value != "" {
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Value))
	}
	if in.Ext != nil {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
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
