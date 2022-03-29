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

func easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb(in *jlexer.Lexer, out *Application) {
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
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "bundle":
			out.Bundle = string(in.String())
		case "domain":
			out.Domain = string(in.String())
		case "storeurl":
			out.StoreURL = string(in.String())
		case "cat":
			if in.IsNull() {
				in.Skip()
				out.Categories = nil
			} else {
				in.Delim('[')
				if out.Categories == nil {
					if !in.IsDelim(']') {
						out.Categories = make([]string, 0, 4)
					} else {
						out.Categories = []string{}
					}
				} else {
					out.Categories = (out.Categories)[:0]
				}
				for !in.IsDelim(']') {
					var v1 string
					v1 = string(in.String())
					out.Categories = append(out.Categories, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "sectioncat":
			if in.IsNull() {
				in.Skip()
				out.SectionCategories = nil
			} else {
				in.Delim('[')
				if out.SectionCategories == nil {
					if !in.IsDelim(']') {
						out.SectionCategories = make([]string, 0, 4)
					} else {
						out.SectionCategories = []string{}
					}
				} else {
					out.SectionCategories = (out.SectionCategories)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.SectionCategories = append(out.SectionCategories, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "pagecat":
			if in.IsNull() {
				in.Skip()
				out.PageCategories = nil
			} else {
				in.Delim('[')
				if out.PageCategories == nil {
					if !in.IsDelim(']') {
						out.PageCategories = make([]string, 0, 4)
					} else {
						out.PageCategories = []string{}
					}
				} else {
					out.PageCategories = (out.PageCategories)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.PageCategories = append(out.PageCategories, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ver":
			out.Version = string(in.String())
		case "privacypolicy":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.HasPrivacyPolicy).UnmarshalJSON(data))
			}
		case "paid":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.IsPaid).UnmarshalJSON(data))
			}
		case "publisher":
			if in.IsNull() {
				in.Skip()
				out.Publisher = nil
			} else {
				if out.Publisher == nil {
					out.Publisher = new(Publisher)
				}
				easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb1(in, out.Publisher)
			}
		case "content":
			if in.IsNull() {
				in.Skip()
				out.Content = nil
			} else {
				if out.Content == nil {
					out.Content = new(Content)
				}
				easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb2(in, out.Content)
			}
		case "keywords":
			out.Keywords = string(in.String())
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
func easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb(out *jwriter.Writer, in Application) {
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
	if in.Bundle != "" {
		const prefix string = ",\"bundle\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Bundle))
	}
	if in.Domain != "" {
		const prefix string = ",\"domain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Domain))
	}
	if in.StoreURL != "" {
		const prefix string = ",\"storeurl\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.StoreURL))
	}
	if len(in.Categories) != 0 {
		const prefix string = ",\"cat\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v4, v5 := range in.Categories {
				if v4 > 0 {
					out.RawByte(',')
				}
				out.String(string(v5))
			}
			out.RawByte(']')
		}
	}
	if len(in.SectionCategories) != 0 {
		const prefix string = ",\"sectioncat\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v6, v7 := range in.SectionCategories {
				if v6 > 0 {
					out.RawByte(',')
				}
				out.String(string(v7))
			}
			out.RawByte(']')
		}
	}
	if len(in.PageCategories) != 0 {
		const prefix string = ",\"pagecat\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v8, v9 := range in.PageCategories {
				if v8 > 0 {
					out.RawByte(',')
				}
				out.String(string(v9))
			}
			out.RawByte(']')
		}
	}
	if in.Version != "" {
		const prefix string = ",\"ver\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Version))
	}
	if in.HasPrivacyPolicy {
		const prefix string = ",\"privacypolicy\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.HasPrivacyPolicy).MarshalJSON())
	}
	if in.IsPaid {
		const prefix string = ",\"paid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.IsPaid).MarshalJSON())
	}
	if in.Publisher != nil {
		const prefix string = ",\"publisher\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb1(out, *in.Publisher)
	}
	if in.Content != nil {
		const prefix string = ",\"content\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb2(out, *in.Content)
	}
	if in.Keywords != "" {
		const prefix string = ",\"keywords\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Keywords))
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
func (v Application) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Application) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Application) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Application) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb(l, v)
}
func easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb2(in *jlexer.Lexer, out *Content) {
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
		case "id":
			out.ID = string(in.String())
		case "episode":
			if in.IsNull() {
				in.Skip()
				out.Episode = nil
			} else {
				if out.Episode == nil {
					out.Episode = new(int)
				}
				*out.Episode = int(in.Int())
			}
		case "title":
			out.Title = string(in.String())
		case "series":
			out.Series = string(in.String())
		case "season":
			out.Season = string(in.String())
		case "artist":
			out.Artist = string(in.String())
		case "genre":
			out.Genre = string(in.String())
		case "album":
			out.Album = string(in.String())
		case "isrc":
			out.ISRC = string(in.String())
		case "producer":
			if in.IsNull() {
				in.Skip()
				out.Producer = nil
			} else {
				if out.Producer == nil {
					out.Producer = new(Producer)
				}
				easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb3(in, out.Producer)
			}
		case "url":
			out.URL = string(in.String())
		case "cat":
			if in.IsNull() {
				in.Skip()
				out.Cat = nil
			} else {
				in.Delim('[')
				if out.Cat == nil {
					if !in.IsDelim(']') {
						out.Cat = make([]string, 0, 4)
					} else {
						out.Cat = []string{}
					}
				} else {
					out.Cat = (out.Cat)[:0]
				}
				for !in.IsDelim(']') {
					var v10 string
					v10 = string(in.String())
					out.Cat = append(out.Cat, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "prodq":
			out.ProdQ = ProductionQuality(in.Int())
		case "videoquality":
			out.VideoQuality = ProductionQuality(in.Int())
		case "context":
			out.Context = ContentContext(in.Int())
		case "contentrating":
			out.ContentRating = string(in.String())
		case "userrating":
			out.UserRating = string(in.String())
		case "qagmediarating":
			out.QAGMediaRating = IQGMediaRating(in.Int())
		case "keywords":
			out.Keywords = string(in.String())
		case "livestream":
			if in.IsNull() {
				in.Skip()
				out.LiveStream = nil
			} else {
				if out.LiveStream == nil {
					out.LiveStream = new(NumericBool)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.LiveStream).UnmarshalJSON(data))
				}
			}
		case "sourcerelationship":
			if in.IsNull() {
				in.Skip()
				out.SourceRelationship = nil
			} else {
				if out.SourceRelationship == nil {
					out.SourceRelationship = new(NumericBool)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.SourceRelationship).UnmarshalJSON(data))
				}
			}
		case "len":
			if in.IsNull() {
				in.Skip()
				out.Len = nil
			} else {
				if out.Len == nil {
					out.Len = new(int)
				}
				*out.Len = int(in.Int())
			}
		case "language":
			out.Language = string(in.String())
		case "embeddable":
			if in.IsNull() {
				in.Skip()
				out.Embeddable = nil
			} else {
				if out.Embeddable == nil {
					out.Embeddable = new(NumericBool)
				}
				if data := in.Raw(); in.Ok() {
					in.AddError((*out.Embeddable).UnmarshalJSON(data))
				}
			}
		case "data":
			if in.IsNull() {
				in.Skip()
				out.Data = nil
			} else {
				in.Delim('[')
				if out.Data == nil {
					if !in.IsDelim(']') {
						out.Data = make([]*Data, 0, 8)
					} else {
						out.Data = []*Data{}
					}
				} else {
					out.Data = (out.Data)[:0]
				}
				for !in.IsDelim(']') {
					var v11 *Data
					if in.IsNull() {
						in.Skip()
						v11 = nil
					} else {
						if v11 == nil {
							v11 = new(Data)
						}
						easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb4(in, v11)
					}
					out.Data = append(out.Data, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
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
func easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb2(out *jwriter.Writer, in Content) {
	out.RawByte('{')
	first := true
	_ = first
	if in.ID != "" {
		const prefix string = ",\"id\":"
		first = false
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Episode != nil {
		const prefix string = ",\"episode\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(*in.Episode))
	}
	if in.Title != "" {
		const prefix string = ",\"title\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Title))
	}
	if in.Series != "" {
		const prefix string = ",\"series\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Series))
	}
	if in.Season != "" {
		const prefix string = ",\"season\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Season))
	}
	if in.Artist != "" {
		const prefix string = ",\"artist\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Artist))
	}
	if in.Genre != "" {
		const prefix string = ",\"genre\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Genre))
	}
	if in.Album != "" {
		const prefix string = ",\"album\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Album))
	}
	if in.ISRC != "" {
		const prefix string = ",\"isrc\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ISRC))
	}
	if in.Producer != nil {
		const prefix string = ",\"producer\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb3(out, *in.Producer)
	}
	if in.URL != "" {
		const prefix string = ",\"url\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.URL))
	}
	if len(in.Cat) != 0 {
		const prefix string = ",\"cat\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v12, v13 := range in.Cat {
				if v12 > 0 {
					out.RawByte(',')
				}
				out.String(string(v13))
			}
			out.RawByte(']')
		}
	}
	if in.ProdQ != 0 {
		const prefix string = ",\"prodq\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.ProdQ))
	}
	if in.VideoQuality != 0 {
		const prefix string = ",\"videoquality\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.VideoQuality))
	}
	if in.Context != 0 {
		const prefix string = ",\"context\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.Context))
	}
	if in.ContentRating != "" {
		const prefix string = ",\"contentrating\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.ContentRating))
	}
	if in.UserRating != "" {
		const prefix string = ",\"userrating\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.UserRating))
	}
	if in.QAGMediaRating != 0 {
		const prefix string = ",\"qagmediarating\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(in.QAGMediaRating))
	}
	if in.Keywords != "" {
		const prefix string = ",\"keywords\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Keywords))
	}
	if in.LiveStream != nil {
		const prefix string = ",\"livestream\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.LiveStream).MarshalJSON())
	}
	if in.SourceRelationship != nil {
		const prefix string = ",\"sourcerelationship\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.SourceRelationship).MarshalJSON())
	}
	if in.Len != nil {
		const prefix string = ",\"len\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int(int(*in.Len))
	}
	if in.Language != "" {
		const prefix string = ",\"language\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Language))
	}
	if in.Embeddable != nil {
		const prefix string = ",\"embeddable\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((*in.Embeddable).MarshalJSON())
	}
	if len(in.Data) != 0 {
		const prefix string = ",\"data\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v14, v15 := range in.Data {
				if v14 > 0 {
					out.RawByte(',')
				}
				if v15 == nil {
					out.RawString("null")
				} else {
					easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb4(out, *v15)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb4(in *jlexer.Lexer, out *Data) {
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
					var v16 *Segment
					if in.IsNull() {
						in.Skip()
						v16 = nil
					} else {
						if v16 == nil {
							v16 = new(Segment)
						}
						easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb5(in, v16)
					}
					out.Segment = append(out.Segment, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
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
func easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb4(out *jwriter.Writer, in Data) {
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
			for v17, v18 := range in.Segment {
				if v17 > 0 {
					out.RawByte(',')
				}
				if v18 == nil {
					out.RawString("null")
				} else {
					easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb5(out, *v18)
				}
			}
			out.RawByte(']')
		}
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb5(in *jlexer.Lexer, out *Segment) {
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
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "value":
			out.Value = string(in.String())
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
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
func easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb5(out *jwriter.Writer, in Segment) {
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
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb3(in *jlexer.Lexer, out *Producer) {
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
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "cat":
			if in.IsNull() {
				in.Skip()
				out.Cat = nil
			} else {
				in.Delim('[')
				if out.Cat == nil {
					if !in.IsDelim(']') {
						out.Cat = make([]string, 0, 4)
					} else {
						out.Cat = []string{}
					}
				} else {
					out.Cat = (out.Cat)[:0]
				}
				for !in.IsDelim(']') {
					var v19 string
					v19 = string(in.String())
					out.Cat = append(out.Cat, v19)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "domain":
			out.Domain = string(in.String())
		case "ext":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Ext).UnmarshalJSON(data))
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
func easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb3(out *jwriter.Writer, in Producer) {
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
	if len(in.Cat) != 0 {
		const prefix string = ",\"cat\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		{
			out.RawByte('[')
			for v20, v21 := range in.Cat {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.String(string(v21))
			}
			out.RawByte(']')
		}
	}
	if in.Domain != "" {
		const prefix string = ",\"domain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.Domain))
	}
	if len(in.Ext) != 0 {
		const prefix string = ",\"ext\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Raw((in.Ext).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjsonB2e97d60DecodeGithubComVungleVungoOpenrtb1(in *jlexer.Lexer, out *Publisher) {
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
		case "id":
			out.ID = string(in.String())
		case "name":
			out.Name = string(in.String())
		case "cat":
			if in.IsNull() {
				in.Skip()
				out.Categories = nil
			} else {
				in.Delim('[')
				if out.Categories == nil {
					if !in.IsDelim(']') {
						out.Categories = make([]string, 0, 4)
					} else {
						out.Categories = []string{}
					}
				} else {
					out.Categories = (out.Categories)[:0]
				}
				for !in.IsDelim(']') {
					var v22 string
					v22 = string(in.String())
					out.Categories = append(out.Categories, v22)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "domain":
			out.Domain = string(in.String())
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
func easyjsonB2e97d60EncodeGithubComVungleVungoOpenrtb1(out *jwriter.Writer, in Publisher) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	if in.Name != "" {
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	if len(in.Categories) != 0 {
		const prefix string = ",\"cat\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v23, v24 := range in.Categories {
				if v23 > 0 {
					out.RawByte(',')
				}
				out.String(string(v24))
			}
			out.RawByte(']')
		}
	}
	if in.Domain != "" {
		const prefix string = ",\"domain\":"
		out.RawString(prefix)
		out.String(string(in.Domain))
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
