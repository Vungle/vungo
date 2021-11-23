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

func easyjson220accf5DecodeGithubComVungleVungoOpenrtbNativeRequest(in *jlexer.Lexer, out *Image) {
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
			out.Type = native.ImageAssetType(in.Int64())
		case "w":
			out.W = int64(in.Int64())
		case "wmin":
			out.WMin = int64(in.Int64())
		case "h":
			out.H = int64(in.Int64())
		case "hmin":
			out.HMin = int64(in.Int64())
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
func easyjson220accf5EncodeGithubComVungleVungoOpenrtbNativeRequest(out *jwriter.Writer, in Image) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Type != 0 {
		const prefix string = ",\"type\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.Type))
	}
	if in.W != 0 {
		const prefix string = ",\"w\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.W))
	}
	if in.WMin != 0 {
		const prefix string = ",\"wmin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.WMin))
	}
	if in.H != 0 {
		const prefix string = ",\"h\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.H))
	}
	if in.HMin != 0 {
		const prefix string = ",\"hmin\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.HMin))
	}
	if len(in.MIMEs) != 0 {
		const prefix string = ",\"mimes\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v2, v3 := range in.MIMEs {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.Extension != nil {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
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

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Image) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson220accf5EncodeGithubComVungleVungoOpenrtbNativeRequest(w, v)
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Image) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson220accf5DecodeGithubComVungleVungoOpenrtbNativeRequest(l, v)
}
