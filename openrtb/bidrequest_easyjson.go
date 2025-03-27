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

func easyjson89fe9b30DecodeGithubComVungleVungoOpenrtb(in *jlexer.Lexer, out *BidRequest) {
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
		case "imp":
			if in.IsNull() {
				in.Skip()
				out.Impressions = nil
			} else {
				in.Delim('[')
				if out.Impressions == nil {
					if !in.IsDelim(']') {
						out.Impressions = make([]*Impression, 0, 8)
					} else {
						out.Impressions = []*Impression{}
					}
				} else {
					out.Impressions = (out.Impressions)[:0]
				}
				for !in.IsDelim(']') {
					var v1 *Impression
					if in.IsNull() {
						in.Skip()
						v1 = nil
					} else {
						if v1 == nil {
							v1 = new(Impression)
						}
						(*v1).UnmarshalEasyJSON(in)
					}
					out.Impressions = append(out.Impressions, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "site":
			if in.IsNull() {
				in.Skip()
				out.Site = nil
			} else {
				if out.Site == nil {
					out.Site = new(Site)
				}
				(*out.Site).UnmarshalEasyJSON(in)
			}
		case "app":
			if in.IsNull() {
				in.Skip()
				out.Application = nil
			} else {
				if out.Application == nil {
					out.Application = new(Application)
				}
				(*out.Application).UnmarshalEasyJSON(in)
			}
		case "device":
			if in.IsNull() {
				in.Skip()
				out.Device = nil
			} else {
				if out.Device == nil {
					out.Device = new(Device)
				}
				(*out.Device).UnmarshalEasyJSON(in)
			}
		case "user":
			if in.IsNull() {
				in.Skip()
				out.User = nil
			} else {
				if out.User == nil {
					out.User = new(User)
				}
				(*out.User).UnmarshalEasyJSON(in)
			}
		case "test":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.IsTestMode).UnmarshalJSON(data))
			}
		case "at":
			out.AuctionType = BidAuctionType(in.Int())
		case "tmax":
			out.Timeout = int(in.Int())
		case "wseat":
			if in.IsNull() {
				in.Skip()
				out.WhitelistedSeats = nil
			} else {
				in.Delim('[')
				if out.WhitelistedSeats == nil {
					if !in.IsDelim(']') {
						out.WhitelistedSeats = make([]string, 0, 4)
					} else {
						out.WhitelistedSeats = []string{}
					}
				} else {
					out.WhitelistedSeats = (out.WhitelistedSeats)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.WhitelistedSeats = append(out.WhitelistedSeats, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "bseat":
			if in.IsNull() {
				in.Skip()
				out.BlocklistedSeats = nil
			} else {
				in.Delim('[')
				if out.BlocklistedSeats == nil {
					if !in.IsDelim(']') {
						out.BlocklistedSeats = make([]string, 0, 4)
					} else {
						out.BlocklistedSeats = []string{}
					}
				} else {
					out.BlocklistedSeats = (out.BlocklistedSeats)[:0]
				}
				for !in.IsDelim(']') {
					var v3 string
					v3 = string(in.String())
					out.BlocklistedSeats = append(out.BlocklistedSeats, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "allimps":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.HasAllImpressions).UnmarshalJSON(data))
			}
		case "cur":
			if in.IsNull() {
				in.Skip()
				out.Currencies = nil
			} else {
				in.Delim('[')
				if out.Currencies == nil {
					if !in.IsDelim(']') {
						out.Currencies = make([]Currency, 0, 4)
					} else {
						out.Currencies = []Currency{}
					}
				} else {
					out.Currencies = (out.Currencies)[:0]
				}
				for !in.IsDelim(']') {
					var v4 Currency
					v4 = Currency(in.String())
					out.Currencies = append(out.Currencies, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "wlang":
			if in.IsNull() {
				in.Skip()
				out.WhitelistLanguages = nil
			} else {
				in.Delim('[')
				if out.WhitelistLanguages == nil {
					if !in.IsDelim(']') {
						out.WhitelistLanguages = make([]string, 0, 4)
					} else {
						out.WhitelistLanguages = []string{}
					}
				} else {
					out.WhitelistLanguages = (out.WhitelistLanguages)[:0]
				}
				for !in.IsDelim(']') {
					var v5 string
					v5 = string(in.String())
					out.WhitelistLanguages = append(out.WhitelistLanguages, v5)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "bcat":
			if in.IsNull() {
				in.Skip()
				out.BlockedCategories = nil
			} else {
				in.Delim('[')
				if out.BlockedCategories == nil {
					if !in.IsDelim(']') {
						out.BlockedCategories = make([]string, 0, 4)
					} else {
						out.BlockedCategories = []string{}
					}
				} else {
					out.BlockedCategories = (out.BlockedCategories)[:0]
				}
				for !in.IsDelim(']') {
					var v6 string
					v6 = string(in.String())
					out.BlockedCategories = append(out.BlockedCategories, v6)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "badv":
			if in.IsNull() {
				in.Skip()
				out.BlockedAdvertisers = nil
			} else {
				in.Delim('[')
				if out.BlockedAdvertisers == nil {
					if !in.IsDelim(']') {
						out.BlockedAdvertisers = make([]string, 0, 4)
					} else {
						out.BlockedAdvertisers = []string{}
					}
				} else {
					out.BlockedAdvertisers = (out.BlockedAdvertisers)[:0]
				}
				for !in.IsDelim(']') {
					var v7 string
					v7 = string(in.String())
					out.BlockedAdvertisers = append(out.BlockedAdvertisers, v7)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "bapp":
			if in.IsNull() {
				in.Skip()
				out.BlockedAdvertisersByMarketID = nil
			} else {
				in.Delim('[')
				if out.BlockedAdvertisersByMarketID == nil {
					if !in.IsDelim(']') {
						out.BlockedAdvertisersByMarketID = make([]string, 0, 4)
					} else {
						out.BlockedAdvertisersByMarketID = []string{}
					}
				} else {
					out.BlockedAdvertisersByMarketID = (out.BlockedAdvertisersByMarketID)[:0]
				}
				for !in.IsDelim(']') {
					var v8 string
					v8 = string(in.String())
					out.BlockedAdvertisersByMarketID = append(out.BlockedAdvertisersByMarketID, v8)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "source":
			if in.IsNull() {
				in.Skip()
				out.Source = nil
			} else {
				if out.Source == nil {
					out.Source = new(Source)
				}
				easyjson89fe9b30DecodeGithubComVungleVungoOpenrtb1(in, out.Source)
			}
		case "regs":
			if in.IsNull() {
				in.Skip()
				out.Regulation = nil
			} else {
				if out.Regulation == nil {
					out.Regulation = new(Regulation)
				}
				(*out.Regulation).UnmarshalEasyJSON(in)
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
func easyjson89fe9b30EncodeGithubComVungleVungoOpenrtb(out *jwriter.Writer, in BidRequest) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"imp\":"
		out.RawString(prefix)
		if in.Impressions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v9, v10 := range in.Impressions {
				if v9 > 0 {
					out.RawByte(',')
				}
				if v10 == nil {
					out.RawString("null")
				} else {
					(*v10).MarshalEasyJSON(out)
				}
			}
			out.RawByte(']')
		}
	}
	if in.Site != nil {
		const prefix string = ",\"site\":"
		out.RawString(prefix)
		(*in.Site).MarshalEasyJSON(out)
	}
	if in.Application != nil {
		const prefix string = ",\"app\":"
		out.RawString(prefix)
		(*in.Application).MarshalEasyJSON(out)
	}
	if in.Device != nil {
		const prefix string = ",\"device\":"
		out.RawString(prefix)
		(*in.Device).MarshalEasyJSON(out)
	}
	if in.User != nil {
		const prefix string = ",\"user\":"
		out.RawString(prefix)
		(*in.User).MarshalEasyJSON(out)
	}
	if in.IsTestMode {
		const prefix string = ",\"test\":"
		out.RawString(prefix)
		out.Raw((in.IsTestMode).MarshalJSON())
	}
	{
		const prefix string = ",\"at\":"
		out.RawString(prefix)
		out.Int(int(in.AuctionType))
	}
	if in.Timeout != 0 {
		const prefix string = ",\"tmax\":"
		out.RawString(prefix)
		out.Int(int(in.Timeout))
	}
	if len(in.WhitelistedSeats) != 0 {
		const prefix string = ",\"wseat\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v11, v12 := range in.WhitelistedSeats {
				if v11 > 0 {
					out.RawByte(',')
				}
				out.String(string(v12))
			}
			out.RawByte(']')
		}
	}
	if len(in.BlocklistedSeats) != 0 {
		const prefix string = ",\"bseat\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v13, v14 := range in.BlocklistedSeats {
				if v13 > 0 {
					out.RawByte(',')
				}
				out.String(string(v14))
			}
			out.RawByte(']')
		}
	}
	if in.HasAllImpressions {
		const prefix string = ",\"allimps\":"
		out.RawString(prefix)
		out.Raw((in.HasAllImpressions).MarshalJSON())
	}
	if len(in.Currencies) != 0 {
		const prefix string = ",\"cur\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v15, v16 := range in.Currencies {
				if v15 > 0 {
					out.RawByte(',')
				}
				out.String(string(v16))
			}
			out.RawByte(']')
		}
	}
	if len(in.WhitelistLanguages) != 0 {
		const prefix string = ",\"wlang\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v17, v18 := range in.WhitelistLanguages {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.String(string(v18))
			}
			out.RawByte(']')
		}
	}
	if len(in.BlockedCategories) != 0 {
		const prefix string = ",\"bcat\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v19, v20 := range in.BlockedCategories {
				if v19 > 0 {
					out.RawByte(',')
				}
				out.String(string(v20))
			}
			out.RawByte(']')
		}
	}
	if len(in.BlockedAdvertisers) != 0 {
		const prefix string = ",\"badv\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v21, v22 := range in.BlockedAdvertisers {
				if v21 > 0 {
					out.RawByte(',')
				}
				out.String(string(v22))
			}
			out.RawByte(']')
		}
	}
	if len(in.BlockedAdvertisersByMarketID) != 0 {
		const prefix string = ",\"bapp\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v23, v24 := range in.BlockedAdvertisersByMarketID {
				if v23 > 0 {
					out.RawByte(',')
				}
				out.String(string(v24))
			}
			out.RawByte(']')
		}
	}
	if in.Source != nil {
		const prefix string = ",\"source\":"
		out.RawString(prefix)
		easyjson89fe9b30EncodeGithubComVungleVungoOpenrtb1(out, *in.Source)
	}
	if in.Regulation != nil {
		const prefix string = ",\"regs\":"
		out.RawString(prefix)
		(*in.Regulation).MarshalEasyJSON(out)
	}
	if len(in.Extension) != 0 {
		const prefix string = ",\"ext\":"
		out.RawString(prefix)
		out.Raw((in.Extension).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BidRequest) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson89fe9b30EncodeGithubComVungleVungoOpenrtb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BidRequest) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson89fe9b30EncodeGithubComVungleVungoOpenrtb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BidRequest) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson89fe9b30DecodeGithubComVungleVungoOpenrtb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BidRequest) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson89fe9b30DecodeGithubComVungleVungoOpenrtb(l, v)
}
func easyjson89fe9b30DecodeGithubComVungleVungoOpenrtb1(in *jlexer.Lexer, out *Source) {
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
		case "fd":
			out.FD = int8(in.Int8())
		case "tid":
			out.TID = string(in.String())
		case "pchain":
			out.PChain = string(in.String())
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
func easyjson89fe9b30EncodeGithubComVungleVungoOpenrtb1(out *jwriter.Writer, in Source) {
	out.RawByte('{')
	first := true
	_ = first
	if in.FD != 0 {
		const prefix string = ",\"fd\":"
		first = false
		out.RawString(prefix[1:])
		out.Int8(int8(in.FD))
	}
	if in.TID != "" {
		const prefix string = ",\"tid\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.TID))
	}
	if in.PChain != "" {
		const prefix string = ",\"pchain\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.PChain))
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
