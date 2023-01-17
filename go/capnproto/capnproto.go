//go:generate capnp compile -I$GOPATH/src/capnproto.org/go/capnp/std -ogo:generated --src-prefix=../../schemas ../../schemas/message.capnp

package capnproto

import (
	"log"

	"capnproto.org/go/capnp/v3"
	"github.com/kcchu/buffer-benchmarks/capnproto/generated"
	"github.com/kcchu/buffer-benchmarks/constants"
)

func Encode(packed bool) []byte {
	arena := capnp.SingleSegment(nil)

	msg, seg, err := capnp.NewMessage(arena)
	if err != nil {
		log.Fatalln("Failed to construct capnp Message:", err)
	}

	castId, err := generated.NewCastId(seg)
	if err != nil {
		log.Fatalln("Failed to construct CastId:", err)
	}
	castId.SetFid(constants.SampleFid)
	castId.SetTsHash(constants.SampleTsHash)

	castAddBody, err := generated.NewCastAddBody(seg)
	if err != nil {
		log.Fatalln("Failed to construct CastAddBody:", err)
	}
	castAddBody.Parent().SetCastId(castId)
	castAddBody.SetText(constants.SampleText)

	md, err := generated.NewRootMessageData(seg)
	if err != nil {
		log.Fatalln("Failed to construct MessageData:", err)
	}
	md.Body().SetCastAddBody(castAddBody)
	md.SetType(generated.MessageType_castAdd)
	md.SetTimestamp(constants.SampleTimestamp)
	md.SetFid(constants.SampleFid)
	md.SetNetwork(generated.FarcasterNetwork_devnet)
	var mdBytes []byte
	if packed {
		mdBytes, err = msg.MarshalPacked()
	} else {
		mdBytes, err = msg.Marshal()
	}
	if err != nil {
		log.Fatalln("Failed to encode MessageData:", err)
	}

	msg2, seg2, err := capnp.NewMessage(arena)
	if err != nil {
		log.Fatalln("Failed to construct capnp Message:", err)
	}

	m, err := generated.NewRootMessage(seg2)
	if err != nil {
		log.Fatalln("Failed to construct Message:", err)
	}
	m.SetData(mdBytes)
	m.SetHash(constants.SampleHash)
	m.SetHashScheme(generated.HashScheme_blake3)
	m.SetSignature(constants.SampleSignature)
	m.SetSignatureScheme(generated.SignatureScheme_ed25519)
	m.SetSigner(constants.SampleSigner)
	var mBytes []byte
	if packed {
		mBytes, err = msg2.MarshalPacked()
	} else {
		mBytes, err = msg2.Marshal()
	}
	if err != nil {
		log.Fatalln("Failed to encode Message:", err)
	}
	return mBytes
}

func Decode(packed bool, buf []byte) {
	var msg1 *capnp.Message
	var err error
	if packed {
		msg1, err = capnp.UnmarshalPacked(buf)
	} else {
		msg1, err = capnp.Unmarshal(buf)
	}
	if err != nil {
		log.Fatalln("Failed to decode capnp Message", err)
	}
	m, err := generated.ReadRootMessage(msg1)
	if err != nil {
		log.Fatalln("Failed to decode Message", err)
	}

	data, err := m.Data()
	if err != nil {
		log.Fatalln("Failed to decode Message.data", err)
	}

	var msg2 *capnp.Message
	if packed {
		msg2, err = capnp.UnmarshalPacked(data)
	} else {
		msg2, err = capnp.Unmarshal(data)
	}
	if err != nil {
		log.Fatalln("Failed to decode data buffer", err)
	}

	md, err := generated.ReadRootMessageData(msg2)
	if err != nil {
		log.Fatalln("Failed to decode MessageData", err)
	}

	castAddBody, err := md.Body().CastAddBody()
	if err != nil {
		log.Fatalln("Failed to decode CastAddBody", err)
	}

	text, err := castAddBody.Text()
	if err != nil {
		log.Fatalln("Failed to decode CastAddBody.text", err)
	}

	if text != constants.SampleText {
		log.Fatalln("Unexpected decoded text")
	}
}
