//go:generate protoc -I=../../schemas --gogofaster_out=paths=source_relative:generated ../../schemas/message.proto

package gogoprotobuf

import (
	"log"

	"github.com/kcchu/buffer-benchmarks/constants"
	"github.com/kcchu/buffer-benchmarks/gogoprotobuf/generated"
)

func Encode() []byte {
	md := &generated.MessageData{
		Body: &generated.MessageData_CastAddBody{
			CastAddBody: &generated.CastAddBody{
				Parent: &generated.CastAddBody_CastId{
					CastId: &generated.CastId{
						Fid:    constants.SampleFid,
						TsHash: constants.SampleTsHash,
					},
				},
				Text: constants.SampleText,
			},
		},
		Type:      generated.MessageType_CastAdd,
		Timestamp: constants.SampleTimestamp,
		Fid:       constants.SampleFid,
		Network:   generated.FarcasterNetwork_Devnet,
	}
	mdBytes, err := md.Marshal()
	if err != nil {
		log.Fatalln("Failed to encode MessageData:", err)
	}

	m := &generated.Message{
		Data:            mdBytes,
		Hash:            constants.SampleHash,
		HashScheme:      generated.HashScheme_Blake3,
		Signature:       constants.SampleSignature,
		SignatureScheme: generated.SignatureScheme_Ed25519,
		Signer:          constants.SampleSigner,
	}
	mBytes, err := m.Marshal()
	if err != nil {
		log.Fatalln("Failed to encode Message:", err)
	}
	return mBytes
}

func Decode(buf []byte) {
	m := new(generated.Message)
	if err := m.Unmarshal(buf); err != nil {
		log.Fatalln("Failed to decode Message", err)
	}
	md := new(generated.MessageData)
	if err := md.Unmarshal(m.Data); err != nil {
		log.Fatalln("Failed to decode MessageData", err)
	}
	castAddBody := md.Body.(*generated.MessageData_CastAddBody).CastAddBody
	if castAddBody.Text != constants.SampleText {
		log.Fatalln("Unexpected decoded text")
	}
}
