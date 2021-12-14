// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package response

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

func easyjson88abeb4cDecodeGithubComVungleVungoOpenrtbNativeResponse(in *jlexer.Lexer, out *EventTracker) {
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
		case "event":
			out.Event = native.EventType(in.Int64())
		case "method":
			out.Method = native.EventTrackingMethod(in.Int64())
		case "url":
			out.URL = string(in.String())
		case "customdata":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.CustomData).UnmarshalJSON(data))
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
func easyjson88abeb4cEncodeGithubComVungleVungoOpenrtbNativeResponse(out *jwriter.Writer, in EventTracker) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"event\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.Event))
	}
	{
		const prefix string = ",\"method\":"
		out.RawString(prefix)
		out.Int64(int64(in.Method))
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		out.RawString(prefix)
		out.String(string(in.URL))
	}
	if len(in.CustomData) != 0 {
		const prefix string = ",\"customdata\":"
		out.RawString(prefix)
		out.Raw((in.CustomData).MarshalJSON())
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
func (v EventTracker) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson88abeb4cEncodeGithubComVungleVungoOpenrtbNativeResponse(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v EventTracker) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson88abeb4cEncodeGithubComVungleVungoOpenrtbNativeResponse(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *EventTracker) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson88abeb4cDecodeGithubComVungleVungoOpenrtbNativeResponse(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *EventTracker) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson88abeb4cDecodeGithubComVungleVungoOpenrtbNativeResponse(l, v)
}
