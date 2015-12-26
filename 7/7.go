package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lookup map[string]Evaluable = make(map[string]Evaluable)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	check(err)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var op1, op2, target string
		var intval uint16
		line := scanner.Text()
		if strings.Contains(line, " AND ") {
			_, err := fmt.Sscanf(line, "%s AND %s -> %s", &op1, &op2, &target)
			check(err)
			lookup[target] = &And{WireReference{op1}, WireReference{op2}}
		} else if strings.Contains(line, " OR ") {
			_, err := fmt.Sscanf(line, "%s OR %s -> %s", &op1, &op2, &target)
			check(err)
			lookup[target] = &Or{WireReference{op1}, WireReference{op2}}
		} else if strings.Contains(line, "NOT ") {
			_, err := fmt.Sscanf(line, "NOT %s -> %s", &op1, &target)
			check(err)
			lookup[target] = &Not{WireReference{op1}}
		} else if strings.Contains(line, " LSHIFT ") {
			_, err := fmt.Sscanf(line, "%s LSHIFT %d -> %s", &op1, &intval, &target)
			check(err)
			lookup[target] = &LeftShift{WireReference{op1}, intval}
		} else if strings.Contains(line, " RSHIFT ") {
			_, err := fmt.Sscanf(line, "%s RSHIFT %d -> %s", &op1, &intval, &target)
			check(err)
			lookup[target] = &RightShift{WireReference{op1}, intval}
		} else {
			_, err := fmt.Sscanf(line, "%s -> %s", &op1, &target)
			check(err)
			intlit, err := strconv.ParseUint(op1, 10, 16)
			if err != nil {
				lookup[target] = &WireReference{op1}
			} else {
				lookup[target] = &LiteralValue{uint16(intlit)}
			}
		}
	}

	fmt.Printf("Wire %s: %d\n", os.Args[2], lookup[os.Args[2]].eval())
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Evaluable interface {
	eval() uint16
}

type LiteralValue struct {
	value uint16
}

func (Value *LiteralValue) eval() uint16 {
	return Value.value
}

type WireReference struct {
	ref string
}

func (Wire *WireReference) eval() uint16 {
	lit, err := strconv.ParseUint(Wire.ref, 10, 16)
	if err == nil {
		return uint16(lit)
	} else {
		value := lookup[Wire.ref].eval()
		lookup[Wire.ref] = &LiteralValue{value} // update lookup to avoid re-evaluation
		return value
	}
}

type And struct {
	op1, op2 WireReference
}

func (And *And) eval() uint16 {
	return And.op1.eval() & And.op2.eval()
}

type Or struct {
	op1, op2 WireReference
}

func (Or *Or) eval() uint16 {
	return Or.op1.eval() | Or.op2.eval()
}

type Not struct {
	op WireReference
}

func (Not *Not) eval() uint16 {
	return ^Not.op.eval()
}

type LeftShift struct {
	op     WireReference
	amount uint16
}

func (Left *LeftShift) eval() uint16 {
	return Left.op.eval() << Left.amount
}

type RightShift struct {
	op     WireReference
	amount uint16
}

func (Right *RightShift) eval() uint16 {
	return Right.op.eval() >> Right.amount
}
