package constants

import (
	"math/rand"
)

const SampleText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua." // 123 bytes
const SampleTimestamp = 1673848001
const SampleTsHashLen = 20
const SampleHashLen = 16
const SampleSignatureLen = 64
const SampleSignerLen = 32

var SampleTsHash []byte    // 20 bytes
var SampleFid []byte       // 4 bytes
var SampleHash []byte      // 16 bytes
var SampleSignature []byte // 64 bytes
var SampleSigner []byte    // 32 bytes

func init() {
	SampleFid = makeBytes(4)
	SampleTsHash = makeBytes(SampleTsHashLen)
	SampleHash = makeBytes(SampleHashLen)
	SampleSignature = makeBytes(SampleSignatureLen)
	SampleSigner = makeBytes(SampleSignerLen)
}

func makeBytes(len int) []byte {
	buf := make([]byte, len)
	rand.Read(buf)
	return buf
}
