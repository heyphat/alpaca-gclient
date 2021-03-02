// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package alpacaio

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

func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient(in *jlexer.Lexer, out *Trades) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Trades, 0, 0)
			} else {
				*out = Trades{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v1 Trade
			(v1).UnmarshalEasyJSON(in)
			*out = append(*out, v1)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient(out *jwriter.Writer, in Trades) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v2, v3 := range in {
			if v2 > 0 {
				out.RawByte(',')
			}
			(v3).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Trades) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Trades) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Trades) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Trades) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient1(in *jlexer.Lexer, out *Trade) {
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
		case "T":
			out.Event = string(in.String())
		case "S":
			out.Symbol = string(in.String())
		case "i":
			out.TradeID = int64(in.Int64())
		case "x":
			out.Exchange = string(in.String())
		case "p":
			out.Price = float32(in.Float32())
		case "s":
			out.Size = int32(in.Int32())
		case "t":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "c":
			if in.IsNull() {
				in.Skip()
				out.Conditions = nil
			} else {
				in.Delim('[')
				if out.Conditions == nil {
					if !in.IsDelim(']') {
						out.Conditions = make([]string, 0, 4)
					} else {
						out.Conditions = []string{}
					}
				} else {
					out.Conditions = (out.Conditions)[:0]
				}
				for !in.IsDelim(']') {
					var v4 string
					v4 = string(in.String())
					out.Conditions = append(out.Conditions, v4)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "z":
			out.Tape = string(in.String())
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
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient1(out *jwriter.Writer, in Trade) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"T\":"
		out.RawString(prefix[1:])
		out.String(string(in.Event))
	}
	{
		const prefix string = ",\"S\":"
		out.RawString(prefix)
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"i\":"
		out.RawString(prefix)
		out.Int64(int64(in.TradeID))
	}
	{
		const prefix string = ",\"x\":"
		out.RawString(prefix)
		out.String(string(in.Exchange))
	}
	{
		const prefix string = ",\"p\":"
		out.RawString(prefix)
		out.Float32(float32(in.Price))
	}
	{
		const prefix string = ",\"s\":"
		out.RawString(prefix)
		out.Int32(int32(in.Size))
	}
	{
		const prefix string = ",\"t\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix)
		if in.Conditions == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Conditions {
				if v5 > 0 {
					out.RawByte(',')
				}
				out.String(string(v6))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"z\":"
		out.RawString(prefix)
		out.String(string(in.Tape))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Trade) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Trade) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Trade) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Trade) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient1(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient2(in *jlexer.Lexer, out *StreamingServerMsges) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(StreamingServerMsges, 0, 0)
			} else {
				*out = StreamingServerMsges{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v7 StreamingServerMsg
			(v7).UnmarshalEasyJSON(in)
			*out = append(*out, v7)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient2(out *jwriter.Writer, in StreamingServerMsges) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v8, v9 := range in {
			if v8 > 0 {
				out.RawByte(',')
			}
			(v9).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v StreamingServerMsges) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StreamingServerMsges) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StreamingServerMsges) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StreamingServerMsges) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient2(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient3(in *jlexer.Lexer, out *StreamingServerMsg) {
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
		case "T":
			out.Event = string(in.String())
		case "S":
			out.Symbol = string(in.String())
		case "i":
			out.TradeID = int64(in.Int64())
		case "x":
			out.Exchange = string(in.String())
		case "p":
			out.Price = float32(in.Float32())
		case "s":
			out.Size = int32(in.Int32())
		case "t":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "z":
			out.Tape = string(in.String())
		case "bx":
			out.BidExchange = int32(in.Int32())
		case "ax":
			out.AskExchange = int32(in.Int32())
		case "bp":
			out.BidPrice = float32(in.Float32())
		case "ap":
			out.AskPrice = float32(in.Float32())
		case "bs":
			out.BidSize = int32(in.Int32())
		case "as":
			out.AskSize = int32(in.Int32())
		case "v":
			out.Volume = int32(in.Int32())
		case "o":
			out.Open = float32(in.Float32())
		case "h":
			out.High = float32(in.Float32())
		case "l":
			out.Low = float32(in.Float32())
		case "c":
			if m, ok := out.C.(easyjson.Unmarshaler); ok {
				m.UnmarshalEasyJSON(in)
			} else if m, ok := out.C.(json.Unmarshaler); ok {
				_ = m.UnmarshalJSON(in.Raw())
			} else {
				out.C = in.Interface()
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
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient3(out *jwriter.Writer, in StreamingServerMsg) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"T\":"
		out.RawString(prefix[1:])
		out.String(string(in.Event))
	}
	{
		const prefix string = ",\"S\":"
		out.RawString(prefix)
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"i\":"
		out.RawString(prefix)
		out.Int64(int64(in.TradeID))
	}
	{
		const prefix string = ",\"x\":"
		out.RawString(prefix)
		out.String(string(in.Exchange))
	}
	{
		const prefix string = ",\"p\":"
		out.RawString(prefix)
		out.Float32(float32(in.Price))
	}
	{
		const prefix string = ",\"s\":"
		out.RawString(prefix)
		out.Int32(int32(in.Size))
	}
	{
		const prefix string = ",\"t\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"z\":"
		out.RawString(prefix)
		out.String(string(in.Tape))
	}
	{
		const prefix string = ",\"bx\":"
		out.RawString(prefix)
		out.Int32(int32(in.BidExchange))
	}
	{
		const prefix string = ",\"ax\":"
		out.RawString(prefix)
		out.Int32(int32(in.AskExchange))
	}
	{
		const prefix string = ",\"bp\":"
		out.RawString(prefix)
		out.Float32(float32(in.BidPrice))
	}
	{
		const prefix string = ",\"ap\":"
		out.RawString(prefix)
		out.Float32(float32(in.AskPrice))
	}
	{
		const prefix string = ",\"bs\":"
		out.RawString(prefix)
		out.Int32(int32(in.BidSize))
	}
	{
		const prefix string = ",\"as\":"
		out.RawString(prefix)
		out.Int32(int32(in.AskSize))
	}
	{
		const prefix string = ",\"v\":"
		out.RawString(prefix)
		out.Int32(int32(in.Volume))
	}
	{
		const prefix string = ",\"o\":"
		out.RawString(prefix)
		out.Float32(float32(in.Open))
	}
	{
		const prefix string = ",\"h\":"
		out.RawString(prefix)
		out.Float32(float32(in.High))
	}
	{
		const prefix string = ",\"l\":"
		out.RawString(prefix)
		out.Float32(float32(in.Low))
	}
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix)
		if m, ok := in.C.(easyjson.Marshaler); ok {
			m.MarshalEasyJSON(out)
		} else if m, ok := in.C.(json.Marshaler); ok {
			out.Raw(m.MarshalJSON())
		} else {
			out.Raw(json.Marshal(in.C))
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StreamingServerMsg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StreamingServerMsg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StreamingServerMsg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StreamingServerMsg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient3(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient4(in *jlexer.Lexer, out *StockTradesResponse) {
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
		case "symbol":
			out.Symbol = string(in.String())
		case "next_page_token":
			out.PageToken = string(in.String())
		case "trades":
			(out.Trades).UnmarshalEasyJSON(in)
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
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient4(out *jwriter.Writer, in StockTradesResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"symbol\":"
		out.RawString(prefix[1:])
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"next_page_token\":"
		out.RawString(prefix)
		out.String(string(in.PageToken))
	}
	{
		const prefix string = ",\"trades\":"
		out.RawString(prefix)
		(in.Trades).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StockTradesResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StockTradesResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StockTradesResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StockTradesResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient4(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient5(in *jlexer.Lexer, out *StockQuotesResponse) {
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
		case "symbol":
			out.Symbol = string(in.String())
		case "next_page_token":
			out.PageToken = string(in.String())
		case "quotes":
			(out.Quotes).UnmarshalEasyJSON(in)
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
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient5(out *jwriter.Writer, in StockQuotesResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"symbol\":"
		out.RawString(prefix[1:])
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"next_page_token\":"
		out.RawString(prefix)
		out.String(string(in.PageToken))
	}
	{
		const prefix string = ",\"quotes\":"
		out.RawString(prefix)
		(in.Quotes).MarshalEasyJSON(out)
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StockQuotesResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StockQuotesResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StockQuotesResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StockQuotesResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient5(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient6(in *jlexer.Lexer, out *StockBarsResponse) {
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
		case "symbol":
			out.Symbol = string(in.String())
		case "next_page_token":
			out.PageToken = string(in.String())
		case "bars":
			if in.IsNull() {
				in.Skip()
				out.Bars = nil
			} else {
				in.Delim('[')
				if out.Bars == nil {
					if !in.IsDelim(']') {
						out.Bars = make(Bars, 0, 1)
					} else {
						out.Bars = Bars{}
					}
				} else {
					out.Bars = (out.Bars)[:0]
				}
				for !in.IsDelim(']') {
					var v10 Bar
					(v10).UnmarshalEasyJSON(in)
					out.Bars = append(out.Bars, v10)
					in.WantComma()
				}
				in.Delim(']')
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
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient6(out *jwriter.Writer, in StockBarsResponse) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"symbol\":"
		out.RawString(prefix[1:])
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"next_page_token\":"
		out.RawString(prefix)
		out.String(string(in.PageToken))
	}
	{
		const prefix string = ",\"bars\":"
		out.RawString(prefix)
		if in.Bars == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Bars {
				if v11 > 0 {
					out.RawByte(',')
				}
				(v12).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v StockBarsResponse) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v StockBarsResponse) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *StockBarsResponse) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *StockBarsResponse) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient6(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient7(in *jlexer.Lexer, out *RequestOptions) {
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
		case "PageToken":
			out.PageToken = string(in.String())
		case "Limit":
			out.Limit = int32(in.Int32())
		case "Start":
			out.Start = string(in.String())
		case "End":
			out.End = string(in.String())
		case "Timeframe":
			out.Timeframe = string(in.String())
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
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient7(out *jwriter.Writer, in RequestOptions) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"PageToken\":"
		out.RawString(prefix[1:])
		out.String(string(in.PageToken))
	}
	{
		const prefix string = ",\"Limit\":"
		out.RawString(prefix)
		out.Int32(int32(in.Limit))
	}
	{
		const prefix string = ",\"Start\":"
		out.RawString(prefix)
		out.String(string(in.Start))
	}
	{
		const prefix string = ",\"End\":"
		out.RawString(prefix)
		out.String(string(in.End))
	}
	{
		const prefix string = ",\"Timeframe\":"
		out.RawString(prefix)
		out.String(string(in.Timeframe))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v RequestOptions) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v RequestOptions) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *RequestOptions) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *RequestOptions) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient7(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient8(in *jlexer.Lexer, out *Quotes) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(Quotes, 0, 0)
			} else {
				*out = Quotes{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v13 Quote
			(v13).UnmarshalEasyJSON(in)
			*out = append(*out, v13)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient8(out *jwriter.Writer, in Quotes) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v14, v15 := range in {
			if v14 > 0 {
				out.RawByte(',')
			}
			(v15).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v Quotes) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Quotes) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Quotes) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Quotes) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient8(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient9(in *jlexer.Lexer, out *Quote) {
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
		case "T":
			out.Event = string(in.String())
		case "S":
			out.Symbol = string(in.String())
		case "c":
			if in.IsNull() {
				in.Skip()
				out.Condition = nil
			} else {
				in.Delim('[')
				if out.Condition == nil {
					if !in.IsDelim(']') {
						out.Condition = make([]string, 0, 4)
					} else {
						out.Condition = []string{}
					}
				} else {
					out.Condition = (out.Condition)[:0]
				}
				for !in.IsDelim(']') {
					var v16 string
					v16 = string(in.String())
					out.Condition = append(out.Condition, v16)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "bx":
			out.BidExchange = string(in.String())
		case "ax":
			out.AskExchange = string(in.String())
		case "bp":
			out.BidPrice = float32(in.Float32())
		case "ap":
			out.AskPrice = float32(in.Float32())
		case "bs":
			out.BidSize = int32(in.Int32())
		case "as":
			out.AskSize = int32(in.Int32())
		case "t":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "z":
			out.Tape = string(in.String())
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
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient9(out *jwriter.Writer, in Quote) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"T\":"
		out.RawString(prefix[1:])
		out.String(string(in.Event))
	}
	{
		const prefix string = ",\"S\":"
		out.RawString(prefix)
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix)
		if in.Condition == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Condition {
				if v17 > 0 {
					out.RawByte(',')
				}
				out.String(string(v18))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"bx\":"
		out.RawString(prefix)
		out.String(string(in.BidExchange))
	}
	{
		const prefix string = ",\"ax\":"
		out.RawString(prefix)
		out.String(string(in.AskExchange))
	}
	{
		const prefix string = ",\"bp\":"
		out.RawString(prefix)
		out.Float32(float32(in.BidPrice))
	}
	{
		const prefix string = ",\"ap\":"
		out.RawString(prefix)
		out.Float32(float32(in.AskPrice))
	}
	{
		const prefix string = ",\"bs\":"
		out.RawString(prefix)
		out.Int32(int32(in.BidSize))
	}
	{
		const prefix string = ",\"as\":"
		out.RawString(prefix)
		out.Int32(int32(in.AskSize))
	}
	{
		const prefix string = ",\"t\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"z\":"
		out.RawString(prefix)
		out.String(string(in.Tape))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Quote) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Quote) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Quote) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Quote) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient9(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient10(in *jlexer.Lexer, out *Bar) {
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
		case "t":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
			}
		case "o":
			out.Open = float32(in.Float32())
		case "h":
			out.High = float32(in.Float32())
		case "l":
			out.Low = float32(in.Float32())
		case "c":
			out.Close = float32(in.Float32())
		case "v":
			out.Volume = int64(in.Int64())
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
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient10(out *jwriter.Writer, in Bar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"t\":"
		out.RawString(prefix[1:])
		out.Raw((in.Timestamp).MarshalJSON())
	}
	{
		const prefix string = ",\"o\":"
		out.RawString(prefix)
		out.Float32(float32(in.Open))
	}
	{
		const prefix string = ",\"h\":"
		out.RawString(prefix)
		out.Float32(float32(in.High))
	}
	{
		const prefix string = ",\"l\":"
		out.RawString(prefix)
		out.Float32(float32(in.Low))
	}
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix)
		out.Float32(float32(in.Close))
	}
	{
		const prefix string = ",\"v\":"
		out.RawString(prefix)
		out.Int64(int64(in.Volume))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Bar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Bar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Bar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Bar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient10(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient11(in *jlexer.Lexer, out *AggregatedBars) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		in.Skip()
		*out = nil
	} else {
		in.Delim('[')
		if *out == nil {
			if !in.IsDelim(']') {
				*out = make(AggregatedBars, 0, 0)
			} else {
				*out = AggregatedBars{}
			}
		} else {
			*out = (*out)[:0]
		}
		for !in.IsDelim(']') {
			var v19 AggregatedBar
			(v19).UnmarshalEasyJSON(in)
			*out = append(*out, v19)
			in.WantComma()
		}
		in.Delim(']')
	}
	if isTopLevel {
		in.Consumed()
	}
}
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient11(out *jwriter.Writer, in AggregatedBars) {
	if in == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
		out.RawString("null")
	} else {
		out.RawByte('[')
		for v20, v21 := range in {
			if v20 > 0 {
				out.RawByte(',')
			}
			(v21).MarshalEasyJSON(out)
		}
		out.RawByte(']')
	}
}

// MarshalJSON supports json.Marshaler interface
func (v AggregatedBars) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AggregatedBars) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AggregatedBars) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AggregatedBars) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient11(l, v)
}
func easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient12(in *jlexer.Lexer, out *AggregatedBar) {
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
		case "T":
			out.Event = string(in.String())
		case "S":
			out.Symbol = string(in.String())
		case "v":
			out.Volume = int32(in.Int32())
		case "o":
			out.Open = float32(in.Float32())
		case "h":
			out.High = float32(in.Float32())
		case "l":
			out.Low = float32(in.Float32())
		case "c":
			out.Close = float32(in.Float32())
		case "s":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Timestamp).UnmarshalJSON(data))
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
func easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient12(out *jwriter.Writer, in AggregatedBar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"T\":"
		out.RawString(prefix[1:])
		out.String(string(in.Event))
	}
	{
		const prefix string = ",\"S\":"
		out.RawString(prefix)
		out.String(string(in.Symbol))
	}
	{
		const prefix string = ",\"v\":"
		out.RawString(prefix)
		out.Int32(int32(in.Volume))
	}
	{
		const prefix string = ",\"o\":"
		out.RawString(prefix)
		out.Float32(float32(in.Open))
	}
	{
		const prefix string = ",\"h\":"
		out.RawString(prefix)
		out.Float32(float32(in.High))
	}
	{
		const prefix string = ",\"l\":"
		out.RawString(prefix)
		out.Float32(float32(in.Low))
	}
	{
		const prefix string = ",\"c\":"
		out.RawString(prefix)
		out.Float32(float32(in.Close))
	}
	{
		const prefix string = ",\"s\":"
		out.RawString(prefix)
		out.Raw((in.Timestamp).MarshalJSON())
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v AggregatedBar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v AggregatedBar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjsonBc289ab0EncodeGithubComGtmkAlpacaGclient12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *AggregatedBar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *AggregatedBar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjsonBc289ab0DecodeGithubComGtmkAlpacaGclient12(l, v)
}
