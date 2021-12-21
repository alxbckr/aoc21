package main

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestLiteralPkg(t *testing.T) {

	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("D2FE28")
	buf.Write(bytes)

	b, _ := buf.ReadByte()
	c := cursor{pos: 0, b: b}

	ver := getPacketVersion(&buf, &c)
	if ver != 6 {
		t.Errorf("packet version is incorrect, got %v, want %v", ver, 6)
	}
	typeId := getPacketTypeId(&buf, &c)
	if typeId != 4 {
		t.Errorf("type id is incorrect, got %v, want %v", typeId, 4)
	}
	literal := decodeLiteralValue(&buf, &c)
	if literal != 2021 {
		t.Errorf("literal is incorrect, got %v, want %v", literal, 2021)
	}
}

func TestOperatorPkgLen(t *testing.T) {

	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("38006F45291200")
	buf.Write(bytes)

	b, _ := buf.ReadByte()
	c := cursor{pos: 0, b: b}

	ver := getPacketVersion(&buf, &c)
	if ver != 1 {
		t.Errorf("packet version is incorrect, got %v, want %v", ver, 1)
	}
	typeId := getPacketTypeId(&buf, &c)
	if typeId != 6 {
		t.Errorf("type id is incorrect, got %v, want %v", typeId, 6)
	}
	lengthTypeId := getLengthTypeId(&buf, &c)
	if lengthTypeId != 0 {
		t.Errorf("length type id is incorrect, got %v, want %v", lengthTypeId, 0)
	}
	totalLength := getTotalBitLength(&buf, &c)
	if totalLength != 27 {
		t.Errorf("total bit legnth is incorrect, got %v, want %v", totalLength, 27)
	}

	startPos := c.consumed
	literals := []uint{}
	for c.consumed < startPos+int(totalLength) {
		vSub := getPacketVersion(&buf, &c)
		if vSub == 0 {
			t.Errorf("subpacket version is zero")
		}
		tSub := getPacketTypeId(&buf, &c)
		if tSub != 4 {
			t.Errorf("subpacket type id is incorrect, got %v, want %v", typeId, 4)
		}
		literal := decodeLiteralValue(&buf, &c)
		literals = append(literals, literal)
	}

	if len(literals) != 2 {
		t.Errorf("literals count is wrong, got %v, want %v", len(literals), 2)
	}

	expLit := []uint{10, 20}
	if !reflect.DeepEqual(literals, expLit) {
		t.Errorf("literals are wrong, got %v, want %v", literals, expLit)
	}
}

func TestOperatorPkgCnt(t *testing.T) {

	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("EE00D40C823060")
	buf.Write(bytes)

	b, _ := buf.ReadByte()
	c := cursor{pos: 0, b: b}

	ver := getPacketVersion(&buf, &c)
	if ver != 7 {
		t.Errorf("packet version is incorrect, got %v, want %v", ver, 7)
	}
	typeId := getPacketTypeId(&buf, &c)
	if typeId != 3 {
		t.Errorf("type id is incorrect, got %v, want %v", typeId, 3)
	}
	lengthTypeId := getLengthTypeId(&buf, &c)
	if lengthTypeId != 1 {
		t.Errorf("length type id is incorrect, got %v, want %v", lengthTypeId, 1)
	}
	numSub := getNumberSubPackets(&buf, &c)
	if numSub != 3 {
		t.Errorf("number subpackets incorrect, got %v, want %v", numSub, 3)
	}

	literals := []uint{}
	for cnt := 0; cnt < int(numSub); cnt++ {
		vSub := getPacketVersion(&buf, &c)
		if vSub == 0 {
			t.Errorf("subpacket version is zero")
		}
		tSub := getPacketTypeId(&buf, &c)
		if tSub != 4 {
			t.Errorf("subpacket type id is incorrect, got %v, want %v", typeId, 4)
		}
		literal := decodeLiteralValue(&buf, &c)
		literals = append(literals, literal)
	}

	if len(literals) != 3 {
		t.Errorf("literals count is wrong, got %v, want %v", len(literals), 3)
	}

	expLit := []uint{1, 2, 3}
	if !reflect.DeepEqual(literals, expLit) {
		t.Errorf("literals are wrong, got %v, want %v", literals, expLit)
	}
}

func TestSumm(t *testing.T) {

	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("C200B40A82")
	buf.Write(bytes)

	res := getResult(&buf)
	if res != 3 {
		t.Errorf("result is, got %v, want %v", res, 3)
	}
}

func TestMult(t *testing.T) {

	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("04005AC33890")
	buf.Write(bytes)

	res := getResult(&buf)
	if res != 54 {
		t.Errorf("result is, got %v, want %v", res, 54)
	}
}

func TestMin(t *testing.T) {

	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("880086C3E88112")
	buf.Write(bytes)

	res := getResult(&buf)
	if res != 7 {
		t.Errorf("result is, got %v, want %v", res, 7)
	}
}

func TestMax(t *testing.T) {

	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("CE00C43D881120")
	buf.Write(bytes)

	res := getResult(&buf)
	if res != 9 {
		t.Errorf("result is, got %v, want %v", res, 9)
	}
}
func TestLessLess(t *testing.T) {

	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("D8005AC2A8F0")
	buf.Write(bytes)

	res := getResult(&buf)
	if res != 1 {
		t.Errorf("result is, got %v, want %v", res, 1)
	}
}

func TestLessGreater(t *testing.T) {

	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("F600BC2D8F")
	buf.Write(bytes)

	res := getResult(&buf)
	if res != 0 {
		t.Errorf("result is, got %v, want %v", res, 0)
	}
}

func TestEqualNot(t *testing.T) {

	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("9C005AC2F8F0")
	buf.Write(bytes)

	res := getResult(&buf)
	if res != 0 {
		t.Errorf("result is, got %v, want %v", res, 0)
	}
}

func TestEqualCompl(t *testing.T) {

	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("9C0141080250320F1802104A08")
	buf.Write(bytes)

	res := getResult(&buf)
	if res != 1 {
		t.Errorf("result is, got %v, want %v", res, 1)
	}
}
