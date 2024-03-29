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

func easyjson8a221a72DecodeGithubComVungleVungoOpenrtb(in *jlexer.Lexer, out *Deal) {
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
		case "bidfloor":
			out.BidFloorPrice = float64(in.Float64())
		case "bidfloorcur":
			out.BidFloorCurrency = Currency(in.String())
		case "at":
			out.AuctionType = DealAuctionType(in.Int())
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
					var v1 string
					v1 = string(in.String())
					out.WhitelistedSeats = append(out.WhitelistedSeats, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "wadomain":
			if in.IsNull() {
				in.Skip()
				out.AdvertiserDomains = nil
			} else {
				in.Delim('[')
				if out.AdvertiserDomains == nil {
					if !in.IsDelim(']') {
						out.AdvertiserDomains = make([]string, 0, 4)
					} else {
						out.AdvertiserDomains = []string{}
					}
				} else {
					out.AdvertiserDomains = (out.AdvertiserDomains)[:0]
				}
				for !in.IsDelim(']') {
					var v2 string
					v2 = string(in.String())
					out.AdvertiserDomains = append(out.AdvertiserDomains, v2)
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
func easyjson8a221a72EncodeGithubComVungleVungoOpenrtb(out *jwriter.Writer, in Deal) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.String(string(in.ID))
	}
	{
		const prefix string = ",\"bidfloor\":"
		out.RawString(prefix)
		out.Float64(float64(in.BidFloorPrice))
	}
	if in.BidFloorCurrency != "" {
		const prefix string = ",\"bidfloorcur\":"
		out.RawString(prefix)
		out.String(string(in.BidFloorCurrency))
	}
	if in.AuctionType != 0 {
		const prefix string = ",\"at\":"
		out.RawString(prefix)
		out.Int(int(in.AuctionType))
	}
	if len(in.WhitelistedSeats) != 0 {
		const prefix string = ",\"wseat\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v3, v4 := range in.WhitelistedSeats {
				if v3 > 0 {
					out.RawByte(',')
				}
				out.String(string(v4))
			}
			out.RawByte(']')
		}
	}
	if len(in.AdvertiserDomains) != 0 {
		const prefix string = ",\"wadomain\":"
		out.RawString(prefix)
		{
			out.RawByte('[')
			for v5, v6 := range in.AdvertiserDomains {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
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
func (v Deal) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson8a221a72EncodeGithubComVungleVungoOpenrtb(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Deal) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson8a221a72EncodeGithubComVungleVungoOpenrtb(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Deal) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson8a221a72DecodeGithubComVungleVungoOpenrtb(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Deal) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson8a221a72DecodeGithubComVungleVungoOpenrtb(l, v)
}
