package constants

const SampleText = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua." // 123 bytes
const SampleTimestamp = 1673848001
const SampleHashLen = 16
const SampleSignatureLen = 64
const SampleSignerLen = 32

var SampleFid []byte       // 4 bytes
var SampleHash []byte      // 16 bytes
var SampleSignature []byte // 64 bytes
var SampleSigner []byte    // 32 bytes

func init() {
	SampleFid = []byte{0x31, 0x32, 0x33, 0x34}
	SampleHash = make([]byte, SampleHashLen)
	for i := 0; i < SampleHashLen; i++ {
		SampleHash[i] = 0xFF
	}
	SampleSignature = make([]byte, SampleSignatureLen)
	for i := 0; i < SampleSignatureLen; i++ {
		SampleSignature[i] = 0xFF
	}
	SampleSigner = make([]byte, SampleSignerLen)
	for i := 0; i < SampleSignerLen; i++ {
		SampleSigner[i] = 0xFF
	}
}
