//go:generate flatc --go ../../schemas/message.fbs --go-namespace generated

package flatbuffers

import (
	"log"

	flatbuffers "github.com/google/flatbuffers/go"
	"github.com/kcchu/buffer-benchmarks/constants"
	"github.com/kcchu/buffer-benchmarks/flatbuffers/generated"
)

func Encode() []byte {
	mdb := flatbuffers.NewBuilder(1024)

	text := mdb.CreateString(constants.SampleText)
	fid := mdb.CreateByteVector(constants.SampleFid)

	// CastAddBody
	generated.CastAddBodyStart(mdb)
	generated.CastAddBodyAddText(mdb, text)
	castAddBody := generated.CastAddBodyEnd(mdb)

	// MessageData
	generated.MessageDataStart(mdb)
	generated.MessageDataAddBodyType(mdb, generated.MessageBodyCastAddBody)
	generated.MessageDataAddBody(mdb, castAddBody)
	generated.MessageDataAddType(mdb, generated.MessageTypeCastAdd)
	generated.MessageDataAddTimestamp(mdb, constants.SampleTimestamp)
	generated.MessageDataAddFid(mdb, fid)
	generated.MessageDataAddNetwork(mdb, generated.FarcasterNetworkDevnet)
	messageData := generated.MessageDataEnd(mdb)

	mdb.Finish(messageData)

	mb := flatbuffers.NewBuilder(1024)
	data := mb.CreateByteVector(mdb.FinishedBytes())
	hash := mb.CreateByteVector(constants.SampleHash)
	signature := mb.CreateByteVector(constants.SampleSignature)
	signer := mb.CreateByteVector(constants.SampleSigner)

	// Message
	generated.MessageStart(mb)
	generated.MessageAddData(mb, data)
	generated.MessageAddHash(mb, hash)
	generated.MessageAddHashScheme(mb, generated.HashSchemeBlake3)
	generated.MessageAddSignature(mb, signature)
	generated.MessageAddSignatureScheme(mb, generated.SignatureSchemeEd25519)
	generated.MessageAddSigner(mb, signer)
	message := generated.MessageEnd(mb)

	mb.Finish(message)
	return mb.FinishedBytes()
}

func Decode(buf []byte) {
	m := generated.GetRootAsMessage(buf, 0)
	data := m.DataBytes()
	md := generated.GetRootAsMessageData(data, 0)
	bodyTable := new(flatbuffers.Table)
	if !md.Body(bodyTable) {
		log.Fatalln("Cannot decode body")
		castAddBody := new(generated.CastAddBody)
		castAddBody.Init(bodyTable.Bytes, bodyTable.Pos)
		if string(castAddBody.Text()) != constants.SampleText {
			log.Fatalln("Unexpected decoded text")
		}
	}
}
