package cpu6502

func (cpu *CPU) Step() uint8 {
	opcode := cpu.ReadMemory(cpu.PC)
	cpu.PC++
	cpu.cycles = opcodeCycles[opcode]
	// TODO: byte following PC is read in advance. Test it

	switch opcode {
		// LDA LDX LDY
		case LDY_Imm_A0: cpu.loadReg(&cpu.Y, cpu.immediate())
		case LDX_Imm_A2: cpu.loadReg(&cpu.X, cpu.immediate())
		case LDA_Imm_A9: cpu.loadReg(&cpu.A, cpu.immediate())
		case LDY_Zp_A4: cpu.loadReg(&cpu.Y, cpu.zeroPage())
		case LDA_Zp_A5: cpu.loadReg(&cpu.A, cpu.zeroPage())
		case LDX_Zp_A6: cpu.loadReg(&cpu.X, cpu.zeroPage())
		case LDY_ZpX_B4: cpu.loadReg(&cpu.Y, cpu.zeroPageIndexed(cpu.X))
		case LDA_ZpX_B5: cpu.loadReg(&cpu.A, cpu.zeroPageIndexed(cpu.X))
		case LDX_ZpY_B6: cpu.loadReg(&cpu.X, cpu.zeroPageIndexed(cpu.Y))
		case LDY_Abs_AC: cpu.loadReg(&cpu.Y, cpu.absolute())
		case LDA_Abs_AD: cpu.loadReg(&cpu.A, cpu.absolute())
		case LDX_Abs_AE: cpu.loadReg(&cpu.X, cpu.absolute())
		case LDA_AbsY_B9: cpu.loadReg(&cpu.A, cpu.absoluteIndexed(cpu.Y))
		case LDY_AbsX_BC: cpu.loadReg(&cpu.Y, cpu.absoluteIndexed(cpu.X))
		case LDA_AbsX_BD: cpu.loadReg(&cpu.A, cpu.absoluteIndexed(cpu.X))
		case LDX_AbsY_BE: cpu.loadReg(&cpu.X, cpu.absoluteIndexed(cpu.Y))
		case LDA_IndexIndirX_A1: cpu.loadReg(&cpu.A, cpu.indexedIndirectX())
		case LDA_IndirIndexY_B1: cpu.loadReg(&cpu.A, cpu.indirectIndexedY())
		// STA STX STY
		case STY_Zp_84: cpu.storeReg(cpu.Y, cpu.zeroPageAddress())
		case STA_Zp_85: cpu.storeReg(cpu.A, cpu.zeroPageAddress())
		case STX_Zp_86: cpu.storeReg(cpu.X, cpu.zeroPageAddress())
		case STY_ZpX_94: cpu.storeReg(cpu.Y, cpu.zeroPageIndexedAddress(cpu.X))
		case STA_ZpX_95: cpu.storeReg(cpu.A, cpu.zeroPageIndexedAddress(cpu.X))
		case STX_ZpY_96: cpu.storeReg(cpu.X, cpu.zeroPageIndexedAddress(cpu.Y))
		case STY_Abs_8C: cpu.storeReg(cpu.Y, cpu.absoluteAddress())
		case STA_Abs_8D: cpu.storeReg(cpu.A, cpu.absoluteAddress())
		case STX_Abs_8E: cpu.storeReg(cpu.X, cpu.absoluteAddress())
		case STA_AbsY_99: cpu.storeReg(cpu.A, cpu.absoluteIndexedAddress(cpu.Y))
		case STA_AbsX_9D: cpu.storeReg(cpu.A, cpu.absoluteIndexedAddress(cpu.X))
		case STA_IndexIndirX_81: cpu.storeReg(cpu.A, cpu.indexedIndirectXaddress())
		case STA_IndirIndexY_91: cpu.storeReg(cpu.A, cpu.indirectIndexedYaddress())
		// INX INY INC DEX DEY DEC
		case INX_E8: cpu.incDecReg(&cpu.X, +1)
		case DEX_CA: cpu.incDecReg(&cpu.X, -1)
		case INY_C8: cpu.incDecReg(&cpu.Y, +1)
		case DEY_88: cpu.incDecReg(&cpu.Y, -1)
		case INC_Zp_E6: cpu.incDec(cpu.zeroPageAddress(), +1 )
		case DEC_Zp_C6: cpu.incDec(cpu.zeroPageAddress(), -1 )
		case INC_ZpX_F6: cpu.incDec(cpu.zeroPageIndexedAddress(cpu.X), +1 )
		case DEC_ZpX_D6: cpu.incDec(cpu.zeroPageIndexedAddress(cpu.X), -1 )
		case INC_Abs_EE: cpu.incDec(cpu.absoluteAddress(), +1 )
		case DEC_Abs_CE: cpu.incDec(cpu.absoluteAddress(), -1 )
		case INC_AbsX_FE: cpu.incDec(cpu.absoluteIndexedAddress(cpu.X), +1 )
		case DEC_AbsX_DE: cpu.incDec(cpu.absoluteIndexedAddress(cpu.X), -1 )
		// ADC SBC
		case ADC_Imm_69: cpu.adc(cpu.immediate())
		case SBC_Imm_E9: cpu.sbc(cpu.immediate())
		case ADC_Zp_65: cpu.adc(cpu.zeroPage())
		case SBC_Zp_E5: cpu.sbc(cpu.zeroPage())
		case ADC_ZpX_75: cpu.adc(cpu.zeroPageIndexed(cpu.X))
		case SBC_ZpX_F5: cpu.sbc(cpu.zeroPageIndexed(cpu.X))
		case ADC_Abs_6D: cpu.adc(cpu.absolute())
		case SBC_Abs_ED: cpu.sbc(cpu.absolute())
		case ADC_AbsX_7D: cpu.adc(cpu.absoluteIndexed(cpu.X))
		case SBC_AbsX_FD: cpu.sbc(cpu.absoluteIndexed(cpu.X))
		case ADC_AbsY_79: cpu.adc(cpu.absoluteIndexed(cpu.Y))
		case SBC_AbsY_F9: cpu.sbc(cpu.absoluteIndexed(cpu.Y))
		case ADC_IndexIndirX_61: cpu.adc(cpu.indexedIndirectX())
		case SBC_IndexIndirX_E1: cpu.sbc(cpu.indexedIndirectX())
		case ADC_IndirIndexY_71: cpu.adc(cpu.indirectIndexedY())
		case SBC_IndirIndexY_F1: cpu.sbc(cpu.indirectIndexedY())
		// CMP CPX CPY
		case CMP_Imm_C9:  cpu.compare(cpu.A, cpu.immediate())
		case CPX_Imm_E0:  cpu.compare(cpu.X, cpu.immediate())
		case CPY_Imm_C0:  cpu.compare(cpu.Y, cpu.immediate())
		case CMP_Zp_C5:   cpu.compare(cpu.A, cpu.zeroPage())
		case CPX_Zp_E4:   cpu.compare(cpu.X, cpu.zeroPage())
		case CPY_Zp_C4:   cpu.compare(cpu.Y, cpu.zeroPage())
		case CMP_ZpX_D5:  cpu.compare(cpu.A, cpu.zeroPageIndexed(cpu.X))
		case CMP_Abs_CD:  cpu.compare(cpu.A, cpu.absolute())
		case CPX_Abs_EC:  cpu.compare(cpu.X, cpu.absolute())
		case CPY_Abs_CC:  cpu.compare(cpu.Y, cpu.absolute())
		case CMP_AbsX_DD: cpu.compare(cpu.A, cpu.absoluteIndexed(cpu.X))
		case CMP_AbsY_D9: cpu.compare(cpu.A, cpu.absoluteIndexed(cpu.Y))
		case CMP_IndexIndirX_C1: cpu.compare(cpu.A, cpu.indexedIndirectX())
		case CMP_IndirIndexY_D1: cpu.compare(cpu.A, cpu.indirectIndexedY())
		// ASL LSR ROR ROL
		case ASL_Acu_0A:  cpu.asla()
		case ROL_Acu_2A:  cpu.rola()
		case LSR_Acu_4A:  cpu.lsra()
		case ROR_Acu_6A:  cpu.rora()
		case ASL_Zp_06:   cpu.asl(cpu.zeroPageAddress())
		case LSR_Zp_46:   cpu.lsr(cpu.zeroPageAddress())
		case ROL_Zp_26:   cpu.rol(cpu.zeroPageAddress())
		case ROR_Zp_66:   cpu.ror(cpu.zeroPageAddress())
		case ASL_ZpX_16:  cpu.asl(cpu.zeroPageIndexedAddress(cpu.X))
		case LSR_ZpX_56:  cpu.lsr(cpu.zeroPageIndexedAddress(cpu.X))
		case ROL_ZpX_36:  cpu.rol(cpu.zeroPageIndexedAddress(cpu.X))
		case ROR_ZpX_76:  cpu.ror(cpu.zeroPageIndexedAddress(cpu.X))
		case ASL_Abs_0E:  cpu.asl(cpu.absoluteAddress())
		case LSR_Abs_4E:  cpu.lsr(cpu.absoluteAddress())
		case ROL_Abs_2E:  cpu.rol(cpu.absoluteAddress())
		case ROR_Abs_6E:  cpu.ror(cpu.absoluteAddress())
		case ASL_AbsX_1E: cpu.asl(cpu.absoluteIndexedAddress(cpu.X))
		case LSR_AbsX_5E: cpu.lsr(cpu.absoluteIndexedAddress(cpu.X))
		case ROL_AbsX_3E: cpu.rol(cpu.absoluteIndexedAddress(cpu.X))
		case ROR_AbsX_7E: cpu.ror(cpu.absoluteIndexedAddress(cpu.X))
		// TAX TAY TXA TYA TSX TXS
		case TAX_AA: cpu.transferRegister(cpu.A, &cpu.X)
		case TAY_A8: cpu.transferRegister(cpu.A, &cpu.Y)
		case TXA_8A: cpu.transferRegister(cpu.X, &cpu.A)
		case TYA_98: cpu.transferRegister(cpu.Y, &cpu.A)
		case TSX_BA: cpu.transferRegister(cpu.Stack, &cpu.X)
		case TXS_9A: cpu.Stack = cpu.X; // TXS doesn't affect status flags
		// AND EOR ORA BIT
		case AND_Imm_29:  cpu.and(cpu.immediate())
		case EOR_Imm_49:  cpu.eor(cpu.immediate())
		case ORA_Imm_09:  cpu.ora(cpu.immediate())
		case AND_Zp_25:   cpu.and(cpu.zeroPage())
		case EOR_Zp_45:   cpu.eor(cpu.zeroPage())
		case ORA_Zp_05:   cpu.ora(cpu.zeroPage())
		case BIT_Zp_24:   cpu.bit(cpu.zeroPage())
		case AND_ZpX_35:  cpu.and(cpu.zeroPageIndexed(cpu.X))
		case EOR_ZpX_55:  cpu.eor(cpu.zeroPageIndexed(cpu.X))
		case ORA_ZpX_15:  cpu.ora(cpu.zeroPageIndexed(cpu.X))
		case AND_Abs_2D:  cpu.and(cpu.absolute())
		case EOR_Abs_4D:  cpu.eor(cpu.absolute())
		case ORA_Abs_0D:  cpu.ora(cpu.absolute())
		case BIT_Abs_2C:  cpu.bit(cpu.absolute())
		case AND_AbsX_3D: cpu.and(cpu.absoluteIndexed(cpu.X))
		case EOR_AbsX_5D: cpu.eor(cpu.absoluteIndexed(cpu.X))
		case ORA_AbsX_1D: cpu.ora(cpu.absoluteIndexed(cpu.X))
		case AND_AbsY_39: cpu.and(cpu.absoluteIndexed(cpu.Y))
		case EOR_AbsY_59: cpu.eor(cpu.absoluteIndexed(cpu.Y))
		case ORA_AbsY_19: cpu.ora(cpu.absoluteIndexed(cpu.Y))
		case AND_IndexIndirX_21: cpu.and(cpu.indexedIndirectX())
		case EOR_IndexIndirX_41: cpu.eor(cpu.indexedIndirectX())
		case ORA_IndexIndirX_01: cpu.ora(cpu.indexedIndirectX())
		case AND_IndirIndexY_31: cpu.and(cpu.indirectIndexedY())
		case EOR_IndirIndexY_51: cpu.eor(cpu.indirectIndexedY())
		case ORA_IndirIndexY_11: cpu.ora(cpu.indirectIndexedY())
		// Branches
		case BEQ_Relative_F0: cpu.branch(cpu.Status.Zero, true, cpu.immediate())
		case BNE_Relative_D0: cpu.branch(cpu.Status.Zero, false, cpu.immediate())
		case BMI_Relative_30: cpu.branch(cpu.Status.Negative, true, cpu.immediate())
		case BPL_Relative_10: cpu.branch(cpu.Status.Negative, false, cpu.immediate())
		case BCS_Relative_B0: cpu.branch(cpu.Status.Carry, true, cpu.immediate())
		case BCC_Relative_90: cpu.branch(cpu.Status.Carry, false, cpu.immediate())
		case BVS_Relative_70: cpu.branch(cpu.Status.Overflow, true, cpu.immediate())
		case BVC_Relative_50: cpu.branch(cpu.Status.Overflow, false, cpu.immediate())
		// Status flags
		case CLC_18: cpu.Status.Carry       = false
		case SEC_38: cpu.Status.Carry       = true
		case CLI_58: cpu.Status.NoInterrupt = false
		case SEI_78: cpu.Status.NoInterrupt = true
		case CLD_D8: cpu.Status.Decimal     = false
		case SED_F8: cpu.Status.Decimal     = true
		case CLV_B8: cpu.Status.Overflow    = false
		// Stack
		case PHP_08: cpu.php()
		case PLP_28: cpu.plp()
		case PHA_48: cpu.push(cpu.A)
		case PLA_68: cpu.pla()
		// Jumps
		case JMP_Abs_4C: cpu.jumpAbsolute()
		case JMP_Ind_6C: cpu.jumpIndirect()
		case JSR_20: cpu.jsr()
		case RTI_40: cpu.rti()
		case RTS_60: cpu.rts()
		// The byte following BRK is skipped, so it's like a 2 byte instruction
		case BRK_00: cpu.PC++; cpu.irq(true)
		case NOP_EA, 0x1A, 0x3A, 0x5A, 0x7A, 0xDA, 0xFA: // implied undocumented NOPs
		case 0x0C, 0x1C, 0x3C, 0x5C, 0x7C, 0xDC, 0xFC: cpu.PC += 2 // absolute undoc NOPs
		case 0x80, 0x82, 0x89, 0xC2, 0xE2, 0x04, 0x14, 0x34, 0x44, 0x54, 0x64, 
			 0x74, 0xD4, 0xF4: cpu.PC++ // immediate and zeropage undoc NOPs
		default: panic("Undocumented opcode")
	}
	return cpu.cycles
}

var opcodeCycles = [0x100] byte {
//  0 1 2 3 4 5 6 7 8 9 A B C D E F
	0,6,2,8,3,3,5,5,3,2,2,2,4,4,6,6, //00
	2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7, //10
	6,6,2,8,3,3,5,5,4,2,2,2,4,4,6,6, //20
	2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7, //30
	6,6,2,8,3,3,5,5,3,2,2,2,3,4,6,6, //40
	2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7, //50
	6,6,2,8,3,3,5,5,4,2,2,2,5,4,6,6, //60
	2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7, //70
	2,6,2,6,3,3,3,3,2,2,2,2,4,4,4,4, //80
	2,6,2,6,4,4,4,4,2,5,2,5,5,5,5,5, //90
	2,6,2,6,3,3,3,3,2,2,2,2,4,4,4,4, //A0
	2,5,2,5,4,4,4,4,2,4,2,4,4,4,4,4, //B0
	2,6,2,8,3,3,5,5,2,2,2,2,4,4,6,6, //C0
	2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7, //D0
	2,6,3,8,3,3,5,5,2,2,2,2,4,4,6,6, //E0
	2,5,2,8,4,4,6,6,2,4,2,7,4,4,7,7} //F0