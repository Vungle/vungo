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

func easyjson3c9d2b01DecodeGithubComVungleVungoOpenrtbNativeRequest(in *jlexer.Lexer, out *Request) {
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
		case "ver":
			out.Ver = string(in.String())
		case "context":
			out.Context = native.ContextType(in.Int64())
		case "contextsubtype":
			out.ContextSubType = native.ContextSubType(in.Int64())
		case "plcmttype":
			out.PlacementType = native.PlacementType(in.Int64())
		case "plcmtcnt":
			out.PlacementCnt = int64(in.Int64())
		case "seq":
			out.Seq = int64(in.Int64())
		case "assets":
			if in.IsNull() {
				in.Skip()
				out.Assets = nil
			} else {
				in.Delim('[')
				if out.Assets == nil {
					if !in.IsDelim(']') {
						out.Assets = make([]Asset, 0, 0)
					} else {
						out.Assets = []Asset{}
					}
				} else {
					out.Assets = (out.Assets)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Asset
					(v1).UnmarshalEasyJSON(in)
					out.Assets = append(out.Assets, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "aurlsupport":
			out.AURLSupport = int8(in.Int8())
		case "durlsupport":
			out.DURLSupport = int8(in.Int8())
		case "eventtrackers":
			if in.IsNull() {
				in.Skip()
				out.EventTrackers = nil
			} else {
				in.Delim('[')
				if out.EventTrackers == nil {
					if !in.IsDelim(']') {
						out.EventTrackers = make([]EventTracker, 0, 1)
					} else {
						out.EventTrackers = []EventTracker{}
					}
				} else {
					out.EventTrackers = (out.EventTrackers)[:0]
				}
				for !in.IsDelim(']') {
					var v2 EventTracker
					(v2).UnmarshalEasyJSON(in)
					out.EventTrackers = append(out.EventTrackers, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "privacy":
			out.Privacy = int8(in.Int8())
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
func easyjson3c9d2b01EncodeGithubComVungleVungoOpenrtbNativeRequest(out *jwriter.Writer, in Request) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Ver != "" {
		const prefix string = ",\"ver\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.Ver))
	}
	if in.Context != 0 {
		const prefix string = ",\"context\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Context))
	}
	if in.ContextSubType != 0 {
		const prefix string = ",\"contextsubtype\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.ContextSubType))
	}
	if in.PlacementType != 0 {
		const prefix string = ",\"plcmttype\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PlacementType))
	}
	if in.PlacementCnt != 0 {
		const prefix string = ",\"plcmtcnt\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.PlacementCnt))
	}
	if in.Seq != 0 {
		const prefix string = ",\"seq\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Seq))
	}
	{
		const prefix string = ",\"assets\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		if in.Assets == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v3, v4 := range in.Assets {
				if v3 > 0 {
					out.RawByte(',')
				}
				(v4).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.AURLSupport != 0 {
		const prefix string = ",\"aurlsupport\":"
		out.RawString(prefix)
		out.Int8(int8(in.AURLSupport))
	}
	if in.DURLSupport != 0 {
		const prefix string = ",\"durlsupport\":"
		out.RawString(prefix)
		out.Int8(int8(in.DURLSupport))
	}
	if len(in.EventTrackers) != 0 {
		const prefix string = ",\"eventtrackers\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.EventTrackers {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	if in.Privacy != 0 {
		const prefix string = ",\"privacy\":"
		out.RawString(prefix)
		out.Int8(int8(in.Privacy))
	}
	if len(in.Extension) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Extension).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Request) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3c9d2b01EncodeGithubComVungleVungoOpenrtbNativeRequest(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Request) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3c9d2b01EncodeGithubComVungleVungoOpenrtbNativeRequest(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Request) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3c9d2b01DecodeGithubComVungleVungoOpenrtbNativeRequest(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Request) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3c9d2b01DecodeGithubComVungleVungoOpenrtbNativeRequest(l, v)
}
