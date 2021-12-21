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

func parseLiteralPkg(buf *bytes.Buffer, c *cursor) int {
	return int(decodeLiteralValue(buf, c))
}

func parseOperatorPkg(buf *bytes.Buffer, c *cursor, myTypeId byte) int {
	operands := []int{}
	lengthTypeId := getLengthTypeId(buf, c)
	if lengthTypeId == 0 {
		totalLength := getTotalBitLength(buf, c)
		startPos := c.consumed
		for c.consumed < startPos+int(totalLength) {
			operands = append(operands, getDeepResult(buf, c))
		}
	} else {
		numSub := getNumberSubPackets(buf, c)
		for cnt := 0; cnt < int(numSub); cnt++ {
			operands = append(operands, getDeepResult(buf, c))
		}
	}

	if len(operands) == 0 {
		fmt.Println("wtf")
	}

	switch myTypeId {
	case 0:
		res := 0
		for _, op := range operands {
			res += op
		}
		fmt.Printf("sum of %v is %v\n", operands, res)
		return res
	case 1:
		res := 1
		for _, op := range operands {
			res *= op
		}
		fmt.Printf("mul of %v is %v\n", operands, res)
		return res
	case 2:
		res := math.MaxInt
		for _, op := range operands {
			if op < res {
				res = op
			}
		}
		fmt.Printf("max of %v is %v\n", operands, res)
		return res
	case 3:
		res := math.MinInt
		for _, op := range operands {
			if op > res {
				res = op
			}
		}
		fmt.Printf("min of %v is %v\n", operands, res)
		return res
	case 5:
		if operands[0] > operands[1] {
			fmt.Printf("greater of %v is %v\n", operands, 1)
			return 1
		}
		fmt.Printf("greater of %v is %v\n", operands, 0)
		return 0
	case 6:
		if operands[0] < operands[1] {
			fmt.Printf("lesser of %v is %v\n", operands, 1)
			return 1
		}
		fmt.Printf("lesser of %v is %v\n", operands, 0)
		return 0
	case 7:
		if operands[0] == operands[1] {
			fmt.Printf("equal of %v is %v\n", operands, 1)
			return 1
		}
		fmt.Printf("equal of %v is %v\n", operands, 0)
		return 0
	default:
		return 0
	}

}

func getDeepResult(buf *bytes.Buffer, c *cursor) int {
	getPacketVersion(buf, c)
	typeId := getPacketTypeId(buf, c)
	switch typeId {
	case 4:
		return parseLiteralPkg(buf, c)
	default:
		return parseOperatorPkg(buf, c, typeId)
	}
}

func getResult(buf *bytes.Buffer) int {
	b, _ := buf.ReadByte()
	c := cursor{pos: 0, b: b}
	return getDeepResult(buf, &c)
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

	fmt.Printf("Result is %v\n", getResult(&buf))
}
