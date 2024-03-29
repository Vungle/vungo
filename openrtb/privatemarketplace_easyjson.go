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

func easyjsonE1675e4eDecodeGithubComVungleVungoOpenrtb(in *jlexer.Lexer, out *PrivateMarketplace) {
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
		case "private_auction":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.IsPrivateAuction).UnmarshalJSON(data))
			}
		case "deals":
			if in.IsNull() {
				in.Skip()
				out.Deals = nil
			} else {
				in.Delim('[')
				if out.Deals == nil {
					if !in.IsDelim(']') {
						out.Deals = make([]*Deal, 0, 8)
					} else {
						out.Deals = []*Deal{}
					}
				} else {
					out.Deals = (out.Deals)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Deal
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Deal)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Deals = append(out.Deals, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
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
func easyjsonE1675e4eEncodeGithubComVungleVungoOpenrtb(out *jwriter.Writer, in PrivateMarketplace) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"private_auction\":"
		out.RawString(prefix[1:])
		out.Raw((in.IsPrivateAuction).MarshalJSON())
	}
	{
		const prefix string = ",\"deals\":"
		out.RawString(prefix)
		if in.Deals == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Deals {
				if v2 > 0 {
					out.RawByte(',')
				}
				if v3 == nil {
					out.RawString("null")
				} else {
					(*v3).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
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
func (v PrivateMarketplace) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonE1675e4eEncodeGithubComVungleVungoOpenrtb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v PrivateMarketplace) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonE1675e4eEncodeGithubComVungleVungoOpenrtb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *PrivateMarketplace) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonE1675e4eDecodeGithubComVungleVungoOpenrtb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *PrivateMarketplace) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonE1675e4eDecodeGithubComVungleVungoOpenrtb(l, v)
}
