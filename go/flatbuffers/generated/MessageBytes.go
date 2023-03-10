// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package generated

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type MessageBytes struct {
	_tab flatbuffers.Table
}

func GetRootAsMessageBytes(buf []byte, offset flatbuffers.UOffsetT) *MessageBytes {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &MessageBytes{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsMessageBytes(buf []byte, offset flatbuffers.UOffsetT) *MessageBytes {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &MessageBytes{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *MessageBytes) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *MessageBytes) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *MessageBytes) MessageBytes(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *MessageBytes) MessageBytesLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *MessageBytes) MessageBytesBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *MessageBytes) MutateMessageBytes(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func MessageBytesStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MessageBytesAddMessageBytes(builder *flatbuffers.Builder, messageBytes flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(messageBytes), 0)
}
func MessageBytesStartMessageBytesVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func MessageBytesEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
