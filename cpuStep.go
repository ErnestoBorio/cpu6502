package Cpu6502

func (cpu *Cpu) step() byte {
	cpu.pageCross = false
	opcode := cpu.readMemory[cpu.PC](cpu.PC)
	cpu.PC++

	// this can be incremented by special cases like branch page crossing
	cpu.cycles = opcodeCycles[opcode]

	// The 6502 reads the next byte in advance to gain time, this could have side effects, so it's not trivial
	// TODO: Ignore this for now, then test with and without it and see if it impacts
	// var operand = cpu.ReadMemory[cpu.PC+1](cpu.PC+1)

	switch opcode {

	}

	return cpu.cycles
}
