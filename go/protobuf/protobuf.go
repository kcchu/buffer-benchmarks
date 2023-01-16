//go:generate protoc -I=../../schemas --go_out=.. ../../schemas/message.proto

package protobuf

import (
	"log"

	"github.com/kcchu/buffer-benchmarks/constants"
	"github.com/kcchu/buffer-benchmarks/protobuf/generated"
	"google.golang.org/protobuf/proto"
)

func Encode() []byte {
	md := &generated.MessageData{
		Body: &generated.MessageData_CastAddBody{
			CastAddBody: &generated.CastAddBody{
				Text: constants.SampleText,
			},
		},
		Type:      generated.MessageType_CastAdd,
		Timestamp: constants.SampleTimestamp,
		Fid:       constants.SampleFid,
		Network:   generated.FarcasterNetwork_Devnet,
	}
	mdBytes, err := proto.Marshal(md)
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
	mBytes, err := proto.Marshal(m)
	if err != nil {
		log.Fatalln("Failed to encode Message:", err)
	}
	return mBytes
}

func Decode(buf []byte) {
	m := new(generated.Message)
	if err := proto.Unmarshal(buf, m); err != nil {
		log.Fatalln("Failed to decode Message", err)
	}
	md := new(generated.MessageData)
	if err := proto.Unmarshal(m.Data, md); err != nil {
		log.Fatalln("Failed to decode MessageData", err)
	}
	castAddBody := md.Body.(*generated.MessageData_CastAddBody).CastAddBody
	if castAddBody.Text != constants.SampleText {
		log.Fatalln("Unexpected decoded text")
	}
}
