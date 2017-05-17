package openrtb_test

import (
	"runtime"
	"strings"
	"sync"
	"testing"

	"github.com/Vungle/vungo/openrtb"
)

// BenchmarkMacroSubs10 runs benchmarkMacroSubs with using a string with 10 replacements.
func BenchmarkMacroSubs10(b *testing.B) {
	benchmarkMacroSubs(10, b)
}

// BenchmarkMacroSubs100 runs benchmarkMacroSubs with using a string with 100 replacements.
func BenchmarkMacroSubs100(b *testing.B) {
	benchmarkMacroSubs(100, b)
}

// BenchmarkMacroSubs1000 runs benchmarkMacroSubs with using a string with 1000 replacements.
func BenchmarkMacroSubs1000(b *testing.B) {
	benchmarkMacroSubs(1000, b)
}

// benchmarkMacroSubs is intended to benchmark the macro substitutions when run in a parallel setting,
// it creates a pool of workers and sends b.N substitution tasks through a channel.
func benchmarkMacroSubs(subs int, b *testing.B) {
	workers := runtime.GOMAXPROCS(0) * 3

	// Create expected input and output
	in := strings.Repeat(`abc${AUCTION_ID}defabc${AUCTION_ID:B64}def
        abc${AUCTION_BID_ID}defabc${AUCTION_IMP_ID}def
        abc${AUCTION_SEAT_ID}defabc${AUCTION_AD_ID}defabc${AUCTION_PRICE}def
        abc${AUCTION_CURRENCY}def${AUCTION_ID}${AUCTION_BID_ID}
        ${AUCTION_IMP_ID}${AUCTION_SEAT_ID}${AUCTION_AD_ID}${AUCTION_AD_ID:B64}
        ${AUCTION_PRICE}${AUCTION_CURRENCY}abc${AUCTION_ID}${AUCTION_ID}def/n`, subs)
	br := fakeBidResponse()
	expectedOut := openrtb.MacroSubs(in, br, br.SeatBids[0], br.SeatBids[0].Bids[0], testAuctionResult)
	inCh := make(chan *openrtb.BidResponse)
	outCh := make(chan string, b.N)

	var wg sync.WaitGroup
	wg.Add(workers)

	// Create workers.
	for i := 0; i < workers; i++ {
		go macroSubsWorker(in, inCh, outCh, &wg, b)
	}

	// Create consumer.
	go func() {
		for output := range outCh {
			if output != expectedOut {
				b.Fatalf("Expected output: %v. Got %v instead", expectedOut, output)
			}
		}
	}()
	b.ReportAllocs()
	b.ResetTimer()

	// Create producer.
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			inCh <- fakeBidResponse()
		}
	})

	close(inCh)
	wg.Wait()
	b.StopTimer()
	close(outCh)
}

// macroSubsWorker is a worker for performing openrtb.MacroSubs in BidResponses sent through its channel.
func macroSubsWorker(sub string, in <-chan *openrtb.BidResponse, out chan<- string, wg *sync.WaitGroup, b *testing.B) {
	for bidRes := range in {
		out <- openrtb.MacroSubs(sub, bidRes, bidRes.SeatBids[0], bidRes.SeatBids[0].Bids[0], testAuctionResult)
	}
	wg.Done()
}

// fakeBidResponse creates an example bid response for testing purposes.
func fakeBidResponse() *openrtb.BidResponse {
	var bid = openrtb.Bid{
		AdID:         "Some ad id goes here",
		ID:           "TheBidId!",
		ImpressionID: "ImpressionIdForBid",
		Price:        testPrice,
	}
	var seatbid = openrtb.SeatBid{
		Seat: "SeatBidIdentifier",
		Bids: []*openrtb.Bid{&bid},
	}
	var bidRes = &openrtb.BidResponse{
		ID:       "1234",
		SeatBids: []*openrtb.SeatBid{&seatbid},
		Currency: openrtb.CurrencyUSD,
	}
	return bidRes
}
