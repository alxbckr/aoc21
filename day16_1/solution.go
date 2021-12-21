package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
	"math/bits"
	"os"
)

type cursor struct {
	pos      int
	b        byte
	consumed int
}

func getPacketVersion(buf *bytes.Buffer, c *cursor) byte {
	return getBitVal(buf, 3, c)
}

func getPacketTypeId(buf *bytes.Buffer, c *cursor) byte {
	return getBitVal(buf, 3, c)
}

func getLengthTypeId(buf *bytes.Buffer, c *cursor) byte {
	return getBitVal(buf, 1, c)
}

func getTotalBitLength(buf *bytes.Buffer, c *cursor) uint16 {
	var val uint16
	val = bits.RotateLeft16(uint16(getBitVal(buf, 8, c)), 7)
	val = val + uint16(getBitVal(buf, 7, c))
	return val
}

func getNumberSubPackets(buf *bytes.Buffer, c *cursor) uint16 {
	var val uint16
	val = bits.RotateLeft16(uint16(getBitVal(buf, 8, c)), 3)
	val = val + uint16(getBitVal(buf, 3, c))
	return val
}

func getBitVal(buf *bytes.Buffer, bits int, c *cursor) byte {
	rem := 8 - c.pos
	var val byte
	if bits <= rem {
		val = c.b & (byte(math.Pow(2, float64(rem))) - 1)
		val = val >> (rem - bits)
		if (c.pos+bits)/8 > 0 {
			c.b, _ = buf.ReadByte()
		}
	} else {
		b1, _ := buf.ReadByte()
		delta := bits - rem
		val = (c.b&(byte(math.Pow(2, float64(rem)))-1))<<delta + b1>>(8-delta)
		c.b = b1
	}
	c.pos = (c.pos + bits) % 8
	c.consumed += int(bits)
	return val
}

func decodeLiteralValue(buf *bytes.Buffer, c *cursor) uint {
	var flag byte
	var part byte
	var res uint

	for {
		flag = getBitVal(buf, 1, c)
		res = bits.RotateLeft(res, 4)
		part = getBitVal(buf, 4, c)
		res += uint(part)
		if flag == 0 {
			// last group
			return res
		}
	}
}

func parseLiteralPkg(buf *bytes.Buffer, c *cursor) {
	fmt.Printf("literal value %v\n", decodeLiteralValue(buf, c))
}

func parseOperatorPkg(buf *bytes.Buffer, c *cursor) int {
	sum := 0
	lengthTypeId := getLengthTypeId(buf, c)
	if lengthTypeId == 0 {
		totalLength := getTotalBitLength(buf, c)
		startPos := c.consumed
		for c.consumed < startPos+int(totalLength) {
			sum = sum + getDeepPkgVersions(buf, c)
		}
	} else {
		numSub := getNumberSubPackets(buf, c)
		for cnt := 0; cnt < int(numSub); cnt++ {
			sum = sum + getDeepPkgVersions(buf, c)
		}
	}
	return sum
}

func getDeepPkgVersions(buf *bytes.Buffer, c *cursor) int {
	sum := 0
	ver := getPacketVersion(buf, c)
	sum += int(ver)
	fmt.Printf("Packet version %v\n", ver)
	typeId := getPacketTypeId(buf, c)
	if typeId == 4 {
		parseLiteralPkg(buf, c)
	} else {
		sum += parseOperatorPkg(buf, c)
	}
	return sum
}

func getSumVersions(buf *bytes.Buffer) int {
	b, _ := buf.ReadByte()
	c := cursor{pos: 0, b: b}
	return getDeepPkgVersions(buf, &c)
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer f.Close()

	buf := bytes.Buffer{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		input := scanner.Text()
		bytes, _ := hex.DecodeString(input)
		buf.Write(bytes)
	}

	fmt.Printf("Versions sum %v\n", getSumVersions(&buf))
}
