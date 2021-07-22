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

func easyjson1db63384DecodeGithubComVungleVungoOpenrtb(in *jlexer.Lexer, out *Regulation) {
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
		case "coppa":
			if in.IsNull() {
				in.Skip()
				out.IsCoppaCompliant = nil
			} else {
				if out.IsCoppaCompliant == nil {
					out.IsCoppaCompliant = new(NumericBool)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.IsCoppaCompliant).UnmarshalJSON(data))
				}
			}
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Extension).UnmarshalJSON(data))
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
func easyjson1db63384EncodeGithubComVungleVungoOpenrtb(out *jwriter.Writer, in Regulation) {
	out.RawByte('{')
	first := true
	_ = first
	if in.IsCoppaCompliant != nil {
		const prefix string = ",\"coppa\":"
		first = false
		out.RawString(prefix[1:])
		out.Raw((*in.IsCoppaCompliant).MarshalJSON())
	}
	if len(in.Extension) != 0 {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Extension).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Regulation) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson1db63384EncodeGithubComVungleVungoOpenrtb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Regulation) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson1db63384EncodeGithubComVungleVungoOpenrtb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Regulation) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson1db63384DecodeGithubComVungleVungoOpenrtb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Regulation) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson1db63384DecodeGithubComVungleVungoOpenrtb(l, v)
}
