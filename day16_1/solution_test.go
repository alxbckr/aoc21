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

func TestSumLiteralOnly(t *testing.T) {
	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("D2FE28")
	buf.Write(bytes)

	sum := getSumVersions(&buf)
	if sum != 6 {
		t.Errorf("packet version sum is incorrect, got %v, want %v", sum, 6)
	}

}

func TestSumSingleOpLen(t *testing.T) {
	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("38006F45291200")
	buf.Write(bytes)

	sum := getSumVersions(&buf)
	if sum != 9 {
		t.Errorf("packet version sum is incorrect, got %v, want %v", sum, 9)
	}
}

func TestSumSingleOpCnt(t *testing.T) {
	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("EE00D40C823060")
	buf.Write(bytes)

	sum := getSumVersions(&buf)
	if sum != 14 {
		t.Errorf("packet version sum is incorrect, got %v, want %v", sum, 14)
	}
}

func TestSumDeep1(t *testing.T) {
	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("8A004A801A8002F478")
	buf.Write(bytes)

	sum := getSumVersions(&buf)
	if sum != 16 {
		t.Errorf("packet version sum is incorrect, got %v, want %v", sum, 16)
	}
}

func TestSumDeep2(t *testing.T) {
	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("620080001611562C8802118E34")
	buf.Write(bytes)

	sum := getSumVersions(&buf)
	if sum != 12 {
		t.Errorf("packet version sum is incorrect, got %v, want %v", sum, 12)
	}
}

func TestSumDeep3(t *testing.T) {
	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("C0015000016115A2E0802F182340")
	buf.Write(bytes)

	sum := getSumVersions(&buf)
	if sum != 23 {
		t.Errorf("packet version sum is incorrect, got %v, want %v", sum, 23)
	}
}

func TestSumDeep4(t *testing.T) {
	buf := bytes.Buffer{}
	bytes, _ := hex.DecodeString("A0016C880162017C3686B18A3D4780")
	buf.Write(bytes)

	sum := getSumVersions(&buf)
	if sum != 31 {
		t.Errorf("packet version sum is incorrect, got %v, want %v", sum, 31)
	}
}
