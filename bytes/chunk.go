package bytes

import (
	"errors"
)

const (
	start int = 0
	//End denotes end of byte sequence
	End   int = -1
	on    int = 1
	off   int = 0
)

//Chunk defines composition of a given chunk
type Chunk struct {
	Value  []byte
	Marker int
	Error  error
}

//Join concatenates the chunks into a single slice of byte
func Join() (chan<- Chunk, <-chan Chunk) {
	var byteSeq []byte
	var state = off
	cin := make(chan Chunk, 4)
	cout := make(chan Chunk)
	go joiner(byteSeq, cin, cout, state)
	return cin, cout
}

func joiner(byteSeq []byte, cin chan Chunk, cout chan Chunk, state int) {
	for chunk := range cin {
		switch {
		case chunk.Marker == start:
			state = on
			byteSeq = append(byteSeq, chunk.Value...)
		case chunk.Marker > start && state == on:
			byteSeq = append(byteSeq, chunk.Value...)
		case chunk.Marker == End && state == on:
			byteSeq = append(byteSeq, chunk.Value...)
			cout <- Chunk{Value: byteSeq, Error: nil}
			byteSeq = nil
		case chunk.Marker == End && state == off:
			cout <- Chunk{Value: nil, Error: errors.New("Incomplete message")}
		default:
			continue
		}
	}
}

//Chunks breaks byteSeq into length of chunkSize. Shorter chunks are returned in the channel.
func Chunks(byteSeq []byte, chunkSize int) <-chan Chunk {
	c := make(chan Chunk, 4)
	go chunker(byteSeq, c, chunkSize)
	return c
}

func chunker(byteSeq []byte, c chan<- Chunk, chunkSize int) {
	totSize := len(byteSeq)
	chunkNum := totSize / chunkSize

	for ii := 0; ii < chunkNum; ii++ {
		c <- Chunk{Value: byteSeq[ii*chunkSize : (ii+1)*chunkSize], Marker: ii}
	}
	c <- Chunk{Value: byteSeq[chunkNum*chunkSize : totSize], Marker: End}
	close(c)
}
