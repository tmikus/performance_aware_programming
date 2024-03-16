package main

import "fmt"

type AsmParser struct {
	currentByte int
	data        []byte
}

func NewAsmParser(data []byte) AsmParser {
	return AsmParser{
		currentByte: 0,
		data:        data,
	}
}

func (p *AsmParser) getNextByte() byte {
	output := p.data[p.currentByte]
	p.currentByte++
	return output
}

func (p *AsmParser) isAtEnd() bool {
	return p.currentByte >= len(p.data)
}

func (p *AsmParser) Parse() string {
	output := "bits 16\n\n"
	for !p.isAtEnd() {
		output += p.parseMov() + "\n"
	}
	return output
}

func (p *AsmParser) parseMov() string {
	instruction := p.getNextByte()
	// Register to register
	if (instruction >> 2) == 0b100010 {
		// D flag is set to 1
		invertedOrderOfRegisters := ((instruction >> 1) & 1) == 1
		w := (instruction & 1) == 1
		modRegRM := p.getNextByte()
		rightRegister := p.parseRegisterName(w, modRegRM&0b111)
		leftRegister := p.parseRegisterName(w, (modRegRM>>3)&0b111)
		if invertedOrderOfRegisters {
			return fmt.Sprintf("mov %s, %s", leftRegister, rightRegister)
		}
		return fmt.Sprintf("mov %s, %s", rightRegister, leftRegister)
	}
	panic("Invalid instruction!!!!")
}

func (p *AsmParser) parseRegisterName(w bool, reg byte) string {
	switch reg {
	case 0b000:
		return pickRegister(w, "al", "ax")
	case 0b001:
		return pickRegister(w, "cl", "cx")
	case 0b010:
		return pickRegister(w, "dl", "dx")
	case 0b011:
		return pickRegister(w, "bl", "bx")
	case 0b100:
		return pickRegister(w, "ah", "sp")
	case 0b101:
		return pickRegister(w, "ch", "bp")
	case 0b110:
		return pickRegister(w, "dh", "si")
	case 0b111:
		return pickRegister(w, "bh", "di")
	}
	panic("Invalid reg")
}

func pickRegister(w bool, shortRegName string, longRegName string) string {
	if w {
		return longRegName
	}
	return shortRegName
}
