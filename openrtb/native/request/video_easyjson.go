// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package request

import (
	json "encoding/json"
	openrtb "github.com/Vungle/vungo/openrtb"
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

func easyjson3c9ce8c3DecodeGithubComVungleVungoOpenrtbNativeRequest(in *jlexer.Lexer, out *Video) {
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
		case "mimes":
			if in.IsNull() {
				in.Skip()
				out.MIMEs = nil
			} else {
				in.Delim('[')
				if out.MIMEs == nil {
					if !in.IsDelim(']') {
						out.MIMEs = make([]string, 0, 4)
					} else {
						out.MIMEs = []string{}
					}
				} else {
					out.MIMEs = (out.MIMEs)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.MIMEs = append(out.MIMEs, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "minduration":
			out.MinDuration = int64(in.Int64())
		case "maxduration":
			out.MaxDuration = int64(in.Int64())
		case "protocols":
			if in.IsNull() {
				in.Skip()
				out.Protocols = nil
			} else {
				in.Delim('[')
				if out.Protocols == nil {
					if !in.IsDelim(']') {
						out.Protocols = make([]openrtb.Protocol, 0, 8)
					} else {
						out.Protocols = []openrtb.Protocol{}
					}
				} else {
					out.Protocols = (out.Protocols)[:0]
				}
				for !in.IsDelim(']') {
					var v2 openrtb.Protocol
					v2 = openrtb.Protocol(in.Int())
					out.Protocols = append(out.Protocols, v2)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjson3c9ce8c3EncodeGithubComVungleVungoOpenrtbNativeRequest(out *jwriter.Writer, in Video) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"mimes\":"
		out.RawString(prefix[1:])
		if in.MIMEs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.MIMEs {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"minduration\":"
		out.RawString(prefix)
		out.Int64(int64(in.MinDuration))
	}
	{
		const prefix string = ",\"maxduration\":"
		out.RawString(prefix)
		out.Int64(int64(in.MaxDuration))
	}
	{
		const prefix string = ",\"protocols\":"
		out.RawString(prefix)
		if in.Protocols == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Protocols {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.Int(int(v6))
			}
			out.RawByte(']')
		}
	}
	if len(in.Extension) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Extension).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Video) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9ce8c3EncodeGithubComVungleVungoOpenrtbNativeRequest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Video) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9ce8c3EncodeGithubComVungleVungoOpenrtbNativeRequest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Video) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9ce8c3DecodeGithubComVungleVungoOpenrtbNativeRequest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Video) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9ce8c3DecodeGithubComVungleVungoOpenrtbNativeRequest(l, v)
}
