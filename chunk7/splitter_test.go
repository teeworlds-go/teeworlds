package chunk7

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSplitterMotdSettingsReady(t *testing.T) {
	t.Parallel()
	// real map change packet from a vanilla teeworlds 0.7 server
	// dumped with the go client
	payload := []byte{
		0x40, 0x02, 0x02, 0x02, 0x00, 0x40, 0x07, 0x03, 0x22, 0x01, 0x00, 0x01, 0x00, 0x01, 0x08, 0x40, 0x01, 0x04, 0x0b,
	}

	chunks := UnpackChunks(payload)

	{
		got := len(chunks)
		want := 3

		if want != got {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

func TestSplitterMapChange(t *testing.T) {
	t.Parallel()
	// real map change packet from a vanilla teeworlds 0.7 server
	// dumped with the go client

	// header := []byte{
	// 	0x00, 0x01, 0x01, 0x01, 0x02, 0x03, 0x04,
	// }

	payload := []byte{
		0x40, 0x3a, 0x01, 0x05, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x6b, 0x75, 0x70,
		0x73, 0x00, 0x91, 0xa7, 0xf2, 0xa0, 0x05, 0x8b, 0x11, 0x08, 0xa8, 0x15, 0xa9, 0x19, 0x38, 0xaf, 0xfd,
		0xb6, 0xcb, 0xf1, 0xdb, 0x5a, 0xfc, 0x73, 0x34, 0xae, 0xa3, 0x68, 0xa7, 0xfa, 0x35, 0x54, 0x37, 0xf9,
		0x39, 0x96, 0xd6, 0x05, 0x25, 0xef, 0x7b, 0x22, 0x40, 0x9d,
	}

	chunks := UnpackChunks(payload)

	{
		got := len(chunks)
		want := 1

		if want != got {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}
}

func TestSplitter(t *testing.T) {
	t.Parallel()
	// packet payload of real traffic
	//
	// Teeworlds 0.7 Protocol packet
	//     Flags: none (..00 00..)
	//     Acknowledged sequence number: 3 (.... ..00 0000 0011)
	//     Number of chunks: 3
	//     Token: 560baebb
	//     Payload (80 bytes)
	// Teeworlds 0.7 Protocol chunk: game.sv_vote_clear_options
	//	Header (vital: 5)
	// 	Flags: vital (01.. ....)
	// 	Size: 1 byte (..00 0000 ..00 0001)
	// 	Sequence number: 5 (00.. .... 0000 0101)
	// Teeworlds 0.7 Protocol chunk: game.sv_tune_params
	//	Header (vital: 6)
	//	Flags: vital (01.. ....)
	//	Size: 69 bytes (..00 0001 ..00 0101)
	//	Sequence number: 6 (00.. .... 0000 0110)
	// Teeworlds 0.7 Protocol chunk: game.sv_ready_to_enter
	//	Header (vital: 7)
	//	Flags: vital (01.. ....)
	//	Size: 1 byte (..00 0000 ..00 0001)
	//	Sequence number: 7 (00.. .... 0000 0111)
	data := []byte{
		0x40, 0x01, 0x05, 0x16, 0x41, 0x05, 0x06, 0x0c, 0xa8, 0x0f, 0x88, 0x03, 0x32, 0xa8, 0x14, 0xb0,
		0x12, 0xb4, 0x07, 0x96, 0x02, 0x9f, 0x01, 0xb0, 0xd1, 0x04, 0x80, 0x7d, 0xac, 0x04, 0x9c, 0x17,
		0x32, 0x98, 0xdb, 0x06, 0x80, 0xb5, 0x18, 0x8c, 0x02, 0xbd, 0x01, 0xa0, 0xed, 0x1a, 0x88, 0x03,
		0xbd, 0x01, 0xb8, 0xc8, 0x21, 0x90, 0x01, 0x14, 0xbc, 0x0a, 0xa0, 0x9a, 0x0c, 0x88, 0x03, 0x80,
		0xe2, 0x09, 0x98, 0xea, 0x01, 0xa4, 0x01, 0x00, 0xa4, 0x01, 0xa4, 0x01, 0x40, 0x01, 0x07, 0x10,
	}

	chunks := UnpackChunks(data)

	{
		got := len(chunks)
		want := 3

		if want != got {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	{
		fmt.Printf("%v\n", chunks)

		got := len(chunks[0].Data)
		want := 1

		if want != got {
			t.Errorf("got %v, wanted %v", got, want)
		}

		want = 0x16
		got = int(chunks[0].Data[0])

		if want != got {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	{
		got := len(chunks[1].Data)
		want := 69

		if want != got {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	{
		want := ChunkHeader{
			Flags: ChunkFlags{
				Vital:  true,
				Resend: false,
			},
			Size: 1,
			Seq:  5,
		}

		if !reflect.DeepEqual(chunks[0].Header, want) {
			t.Errorf("got %v, wanted %v", chunks[0].Header, want)
		}
	}
}
