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

func easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse(in *jlexer.Lexer, out *Asset) {
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
			out.ID = int64(in.Int64())
		case "required":
			out.Required = int8(in.Int8())
		case "title":
			if in.IsNull() {
				in.Skip()
				out.Title = nil
			} else {
				if out.Title == nil {
					out.Title = new(Title)
				}
				easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse1(in, out.Title)
			}
		case "img":
			if in.IsNull() {
				in.Skip()
				out.Img = nil
			} else {
				if out.Img == nil {
					out.Img = new(Image)
				}
				easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse2(in, out.Img)
			}
		case "video":
			if in.IsNull() {
				in.Skip()
				out.Video = nil
			} else {
				if out.Video == nil {
					out.Video = new(Video)
				}
				easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse3(in, out.Video)
			}
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				if out.Data == nil {
					out.Data = new(Data)
				}
				easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse4(in, out.Data)
			}
		case "link":
			if in.IsNull() {
				in.Skip()
				out.Link = nil
			} else {
				if out.Link == nil {
					out.Link = new(Link)
				}
				easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse5(in, out.Link)
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
func easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse(out *jwriter.Writer, in Asset) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != 0 {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	if in.Required != 0 {
		const prefix string = ",\"required\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int8(int8(in.Required))
	}
	if in.Title != nil {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse1(out, *in.Title)
	}
	if in.Img != nil {
		const prefix string = ",\"img\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse2(out, *in.Img)
	}
	if in.Video != nil {
		const prefix string = ",\"video\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse3(out, *in.Video)
	}
	if in.Data != nil {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse4(out, *in.Data)
	}
	if in.Link != nil {
		const prefix string = ",\"link\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse5(out, *in.Link)
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
func (v Asset) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Asset) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Asset) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Asset) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse(l, v)
}
func easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse5(in *jlexer.Lexer, out *Link) {
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
		case "url":
			out.URL = string(in.String())
		case "clicktrackers":
			if in.IsNull() {
				in.Skip()
				out.ClickTrackers = nil
			} else {
				in.Delim('[')
				if out.ClickTrackers == nil {
					if !in.IsDelim(']') {
						out.ClickTrackers = make([]string, 0, 4)
					} else {
						out.ClickTrackers = []string{}
					}
				} else {
					out.ClickTrackers = (out.ClickTrackers)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.ClickTrackers = append(out.ClickTrackers, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "fallback":
			out.Fallback = string(in.String())
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
func easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse5(out *jwriter.Writer, in Link) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix[1:])
		out.String(string(in.URL))
	}
	if len(in.ClickTrackers) != 0 {
		const prefix string = ",\"clicktrackers\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v2, v3 := range in.ClickTrackers {
				if v2 > 0 {
					out.RawByte(',')
				}
				out.String(string(v3))
			}
			out.RawByte(']')
		}
	}
	if in.Fallback != "" {
		const prefix string = ",\"fallback\":"
		out.RawString(prefix)
		out.String(string(in.Fallback))
	}
	if len(in.Extension) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Extension).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse4(in *jlexer.Lexer, out *Data) {
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
		case "type":
			out.Type = native.DataAssetType(in.Int64())
		case "len":
			out.Len = int64(in.Int64())
		case "value":
			out.Value = string(in.String())
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
func easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse4(out *jwriter.Writer, in Data) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Type != 0 {
		const prefix string = ",\"type\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.Type))
	}
	if in.Len != 0 {
		const prefix string = ",\"len\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.Len))
	}
	{
		const prefix string = ",\"value\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Value))
	}
	if len(in.Extension) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Extension).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse3(in *jlexer.Lexer, out *Video) {
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
		case "vasttag":
			out.VASTTag = string(in.String())
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
func easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse3(out *jwriter.Writer, in Video) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"vasttag\":"
		out.RawString(prefix[1:])
		out.String(string(in.VASTTag))
	}
	out.RawByte('}')
}
func easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse2(in *jlexer.Lexer, out *Image) {
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
		case "type":
			out.Type = native.ImageAssetType(in.Int64())
		case "url":
			out.URL = string(in.String())
		case "w":
			out.W = int64(in.Int64())
		case "h":
			out.H = int64(in.Int64())
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
func easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse2(out *jwriter.Writer, in Image) {
	out.RawByte('{')
	first := true
	_ = first
	if in.Type != 0 {
		const prefix string = ",\"type\":"
		first = false
		out.RawString(prefix[1:])
		out.Int64(int64(in.Type))
	}
	{
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	if in.W != 0 {
		const prefix string = ",\"w\":"
		out.RawString(prefix)
		out.Int64(int64(in.W))
	}
	if in.H != 0 {
		const prefix string = ",\"h\":"
		out.RawString(prefix)
		out.Int64(int64(in.H))
	}
	if len(in.Extension) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Extension).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson3b94576aDecodeGithubComVungleVungoOpenrtbNativeResponse1(in *jlexer.Lexer, out *Title) {
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
		case "text":
			out.Text = string(in.String())
		case "len":
			out.Len = int64(in.Int64())
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
func easyjson3b94576aEncodeGithubComVungleVungoOpenrtbNativeResponse1(out *jwriter.Writer, in Title) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix[1:])
		out.String(string(in.Text))
	}
	if in.Len != 0 {
		const prefix string = ",\"len\":"
		out.RawString(prefix)
		out.Int64(int64(in.Len))
	}
	if len(in.Extension) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Extension).MarshalJSON())
	}
	out.RawByte('}')
}
