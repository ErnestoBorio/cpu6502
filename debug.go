package cpu6502

import (
	"ansi"
	"fmt"
	"time"
)

type StepInfo struct {
	Operation
	PC      uint16
	Operand uint16
	Address uint16
	Value   byte
	Raw     [3]byte
}

// DbgFetchOp fetches the next operation pointed at by PC with metadata
func (cpu *CPU) DbgFetchOp() StepInfo {
	step := StepInfo{}
	step.PC = cpu.PC
	opcode := cpu.DbgReadMemory(cpu.PC)
	step.Raw[0] = opcode
	step.Operation = Opcodes[opcode]
	step.Address = step.AddressFunc(cpu, cpu.PC+1)
	step.Value = cpu.DbgReadMemory(step.Address)
	if step.Length >= 2 {
		step.Raw[1] = cpu.DbgReadMemory(cpu.PC + 1)
		step.Operand = uint16(step.Raw[1])
		if step.Length == 3 {
			step.Raw[2] = cpu.DbgReadMemory(cpu.PC + 2)
			step.Operand |= uint16(step.Raw[2]) << 8
		}
	}
	return step
}

// Get a little endian 16 bits value from 2 consecutive memory addresses
func (cpu *CPU) DbgGetUint16(address uint16) uint16 {
	value := uint16(cpu.DbgReadMemory(address))
	value |= uint16(cpu.DbgReadMemory(address+1)) << 8
	return value
}

/* Addressing modes */

func (cpu *CPU) DbgNoAddressing(uint16) uint16 {
	return 0
}

func (cpu *CPU) DbgImmediate(pc uint16) uint16 {
	return pc
}

// Returns zero page address from PC's following byte
func (cpu *CPU) DbgZeroPage(pc uint16) uint16 {
	return uint16(cpu.DbgReadMemory(pc))
}

func (cpu *CPU) DbgZeroPageX(pc uint16) uint16 {
	return uint16(cpu.DbgReadMemory(pc) + cpu.X)
}

func (cpu *CPU) DbgZeroPageY(pc uint16) uint16 {
	return uint16(cpu.DbgReadMemory(pc) + cpu.Y)
}

// Returns absolute address from PC's following 2 bytes
func (cpu *CPU) DbgAbsolute(pc uint16) uint16 {
	return cpu.DbgGetUint16(pc)
}

func (cpu *CPU) DbgAbsoluteX(pc uint16) uint16 {
	address := uint16(cpu.DbgReadMemory(pc)) + uint16(cpu.X)
	if address > 0xFF { // if crossed page boundary
		cpu.cycles++
	}
	address += (uint16(cpu.DbgReadMemory(pc+1)) << 8)
	return address
}

func (cpu *CPU) DbgAbsoluteY(pc uint16) uint16 {
	address := uint16(cpu.DbgReadMemory(pc)) + uint16(cpu.Y)
	if address > 0xFF { // if crossed page boundary
		cpu.cycles++
	}
	address += (uint16(cpu.DbgReadMemory(pc+1)) << 8)
	return address
}

func (cpu *CPU) DbgIndexedIndirectX(pc uint16) uint16 {
	pointer := cpu.DbgReadMemory(pc) + cpu.X
	var address uint16 = uint16(cpu.DbgReadMemory(uint16(pointer)))
	pointer++
	address |= (uint16(cpu.DbgReadMemory(uint16(pointer))) << 8)
	return address
}

func (cpu *CPU) DbgIndirectIndexedY(pc uint16) uint16 {
	base := cpu.DbgReadMemory(pc)
	address := uint16(cpu.DbgReadMemory(uint16(base))) + uint16(cpu.Y)
	if address > 0xFF {
		cpu.cycles++
	}
	base++
	return address + (uint16(cpu.DbgReadMemory(uint16(base))) << 8)
}

func (cpu *CPU) DbgIndirect(pc uint16) uint16 {
	pointer := cpu.DbgGetUint16(pc)
	address := uint16(cpu.DbgReadMemory(pointer)) // low byte
	if (pointer & 0xFF) == 0xFF {                 // address wraps around page
		pointer -= 0x100
	}
	pointer++
	return address | (uint16(cpu.DbgReadMemory(pointer)) << 8)
}

// Disassemble runs the c64 printing out the disassembled code that's ran
func (cpu *CPU) Disassemble() {
	for {
		step := cpu.DbgFetchOp()
		time.Sleep(50 * time.Microsecond)

		addrFmt := ""
		switch step.Addressing {
		case "IMM":
			addrFmt = "#$%02X"
		case "ZRP":
			addrFmt = "$%02X"
		case "ZPX":
			addrFmt = "$%02X, X"
		case "ZPY":
			addrFmt = "$%02X, Y"
		case "ABS":
			addrFmt = "$%04X"
		case "ABX":
			addrFmt = "$%04X, X"
		case "ABY":
			addrFmt = "$%04X, Y"
		case "IIX":
			addrFmt = "($%02X, X)"
		case "IIY":
			addrFmt = "($%02X), Y"
		case "IND":
			addrFmt = "($%04X)"
		case "REL":
			delta := int8(step.Operand)
			step.Address = uint16(int(step.PC) + 2 + int(delta))
			addrFmt = "$%04X"
		}
		addr := ""
		if step.Addressing == "REL" {
			addr = fmt.Sprintf(addrFmt, step.Address)
		} else if addrFmt != "" {
			addr = fmt.Sprintf(addrFmt, step.Operand)
		} else if step.Addressing == "ACU" {
			addr = "A"
		}

		var raw string
		raw = fmt.Sprintf("%02X", step.Raw[0])
		if step.Length >= 2 {
			raw += fmt.Sprintf(" %02X", step.Raw[1])
			if step.Length >= 3 {
				raw += fmt.Sprintf(" %02X", step.Raw[2])
			}
		}

		fmt.Printf("%v%04X%v %-8s %v%3s%v %-8s",
			ansi.Cyan, step.PC, ansi.Reset,
			raw,
			ansi.Magenta, step.Mnemonic, ansi.Reset,
			addr)

		if step.Length > 1 && step.Address != step.PC+1 && step.Addressing != "REL" {
			fmt.Printf(" %v%04X%v %02X", ansi.Green, step.Address, ansi.Reset, step.Value)
		}
		// fmt.Println()
		cpu.Step()

		if step.Length > 1 && step.Address != step.PC+1 && step.Addressing != "REL" {
			newVal := cpu.DbgReadMemory(step.Address)
			if step.Value != newVal {
				fmt.Printf(" %v->%v %02X", ansi.Yellow, ansi.Reset, newVal)
			}
		}
		fmt.Println()
	}
}
