package cpu6502

func (cpu *Cpu) CpuStep() {
	opcode := cpu.readMemory[cpu.PC](cpu.PC)
	cpu.cycles = opcodeCycles[opcode]
	// TODO: byte following PC is read in advance. Test

	switch opcode {
		// LDA LDX LDY
		case LDA_Imm_A9: cpu.loadReg(&cpu.A, cpu.immediate())
		case LDX_Imm_A2: cpu.loadReg(&cpu.X, cpu.immediate())
		case LDY_Imm_A0: cpu.loadReg(&cpu.Y, cpu.immediate())
		case LDA_Zp_A5: cpu.loadReg(&cpu.A, cpu.zeroPage())
		case LDX_Zp_A6: cpu.loadReg(&cpu.X, cpu.zeroPage())
		case LDY_Zp_A4: cpu.loadReg(&cpu.Y, cpu.zeroPage())
		case LDY_Zp_X_B4: cpu.loadReg(&cpu.Y, cpu.zeroPageIndexed(cpu.X))
		case LDA_Zp_X_B5: cpu.loadReg(&cpu.A, cpu.zeroPageIndexed(cpu.X))
		case LDX_Zp_Y_B6: cpu.loadReg(&cpu.X, cpu.zeroPageIndexed(cpu.Y))
		case LDA_Abs_AD: cpu.loadReg(&cpu.A, cpu.absolute())
		case LDX_Abs_AE: cpu.loadReg(&cpu.X, cpu.absolute())
		case LDY_Abs_AC: cpu.loadReg(&cpu.Y, cpu.absolute())
		case LDA_Abs_X_BD: cpu.loadReg(&cpu.A, cpu.absoluteIndexed(cpu.X))
		case LDY_Abs_X_BC: cpu.loadReg(&cpu.Y, cpu.absoluteIndexed(cpu.X))
		case LDA_Abs_Y_B9: cpu.loadReg(&cpu.A, cpu.absoluteIndexed(cpu.Y))
		case LDX_Abs_Y_BE: cpu.loadReg(&cpu.X, cpu.absoluteIndexed(cpu.Y))
		case LDA_IndexIndirX_A1: cpu.loadReg(&cpu.A, cpu.indexedIndirectX())
		case LDA_IndirIndexY_B1: cpu.loadReg(&cpu.A, cpu.indirectIndexedY())
		// STA STX STY
		case STY_Zp_84: cpu.storeReg(cpu.Y, word(cpu.zeroPageAddress()))
		case STA_Zp_85: cpu.storeReg(cpu.A, word(cpu.zeroPageAddress()))
		case STX_Zp_86: cpu.storeReg(cpu.X, word(cpu.zeroPageAddress()))
		case STY_Zp_X_94: cpu.storeReg(cpu.Y, cpu.zeroPageIndexed(cpu.X))
		case STA_Zp_X_95: cpu.storeReg(cpu.A, cpu.zeroPageIndexed(cpu.X))
		case STX_Zp_Y_96: cpu.storeReg(cpu.X, cpu.zeroPageIndexed(cpu.Y))
		case STY_Abs_8C: cpu.storeReg(cpu.Y, cpu.absoluteAddress())
		case STA_Abs_8D: cpu.storeReg(cpu.A, cpu.absoluteAddress())
		case STX_Abs_8E: cpu.storeReg(cpu.X, cpu.absoluteAddress())
		case STA_Abs_X_9D: cpu.storeReg(cpu.A, cpu.absoluteIndexedAddress(cpu.X))
		case STA_Abs_Y_99: cpu.storeReg(cpu.A, cpu.absoluteIndexedAddress(cpu.Y))
		case STA_IndexIndirX_81: cpu.storeReg(cpu.A, cpu.indexedIndirectXaddress())
		case STA_IndirIndexY_91: cpu.storeReg(cpu.A, cpu.indirectIndexedYaddress())
		// INX INY INC DEX DEY DEC
		case INX_E8: cpu.incDecReg(&cpu.X, +1)
		case DEX_CA: cpu.incDecReg(&cpu.X, -1)
		case INY_C8: cpu.incDecReg(&cpu.Y, +1)
		case DEY_88: cpu.incDecReg(&cpu.Y, -1)
		case INC_Zp_E6: cpu.incDec(cpu.zeroPageAddress(), +1 )
		case DEC_Zp_C6: cpu.incDec(cpu.zeroPageAddress(), -1 )
		case INC_Zp_X_F6: cpu.incDec(cpu.zeroPageIndexedAddress(cpu.X), +1 )
		case DEC_Zp_X_D6: cpu.incDec(cpu.zeroPageIndexedAddress(cpu.X), -1 )
		case INC_Abs_EE: cpu.incDec(cpu.absoluteAddress(), +1 )
		case DEC_Abs_CE: cpu.incDec(cpu.absoluteAddress(), -1 )
		case INC_Abs_X_FE: cpu.incDec(cpu.absoluteIndexedAddress(cpu.X), +1 )
		case DEC_Abs_X_DE: cpu.incDec(cpu.absoluteIndexedAddress(cpu.X), -1 )
		// ADC SBC
		case ADC_Imm_69: cpu.adc(cpu.immediate())
		case SBC_Imm_E9: cpu.sbc(cpu.immediate())
		case ADC_Zp_65: cpu.adc(cpu.zeroPage())
		case SBC_Zp_E5: cpu.sbc(cpu.zeroPage())
		case ADC_Zp_X_75: cpu.adc(cpu.zeroPageIndexed(cpu.X))
		case SBC_Zp_X_F5: cpu.sbc(cpu.zeroPageIndexed(cpu.X))
		case ADC_Abs_6D: cpu.adc(cpu.absolute())
		case SBC_Abs_ED: cpu.sbc(cpu.absolute())
		case ADC_Abs_X_7D: cpu.adc(cpu.absoluteIndexed(cpu.X))
		case SBC_Abs_X_FD: cpu.sbc(cpu.absoluteIndexed(cpu.X))
		case ADC_Abs_Y_79: cpu.adc(cpu.absoluteIndexed(cpu.Y))
		case SBC_Abs_Y_F9: cpu.sbc(cpu.absoluteIndexed(cpu.Y))
		case ADC_IndexIndirX_61: cpu.adc(cpu.indexedIndirectX())
		case SBC_IndexIndirX_E1: cpu.sbc(cpu.indexedIndirectX())
		case ADC_IndirIndexY_71: cpu.adc(cpu.indirectIndexedY())
		case SBC_IndirIndexY_F1: cpu.sbc(cpu.indirectIndexedY())
		// CMP CPX CPY
		case CMP_Imm_C9: cpu.compare(cpu.A, cpu.immediate())
		case CPX_Imm_E0: cpu.compare(cpu.X, cpu.immediate())
		case CPY_Imm_C0: cpu.compare(cpu.Y, cpu.immediate())
		case CMP_Zp_C5: cpu.compare(cpu.A, cpu.zeroPage())
		case CPX_Zp_E4: cpu.compare(cpu.X, cpu.zeroPage())
		case CPY_Zp_C4: cpu.compare(cpu.Y, cpu.zeroPage())
		case CMP_Zp_X_D5: cpu.compare(cpu.A, cpu.zeroPageIndexed(cpu.X))
		case CMP_Abs_CD: cpu.compare(cpu.A, cpu.absolute())
		case CPX_Abs_EC: cpu.compare(cpu.X, cpu.absolute())
		case CPY_Abs_CC: cpu.compare(cpu.Y, cpu.absolute())
		case CMP_Abs_X_DD: cpu.compare(cpu.A, cpu.absoluteIndexed(cpu.X))
		case CMP_Abs_Y_D9: cpu.compare(cpu.A, cpu.absoluteIndexed(cpu.Y))
		case CMP_IndexIndirX_C1: cpu.compare(cpu.A, cpu.indexedIndirectX())
		case CMP_IndirIndexY_D1: cpu.compare(cpu.A, cpu.indirectIndexedY())
		// ASL LSR ROR ROL
		case ASL_Acu_0A:   cpu.asla()
		case ROL_Acu_2A:   cpu.rola()
		case LSR_Acu_4A:   cpu.lsra()
		case ROR_Acu_6A:   cpu.rora()
		case ASL_Zp_06:    cpu.asl(cpu.zeroPageAddress())
		case LSR_Zp_46:    cpu.lsr(cpu.zeroPageAddress())
		case ROL_Zp_26:    cpu.rol(cpu.zeroPageAddress())
		case ROR_Zp_66:    cpu.ror(cpu.zeroPageAddress())
		case ASL_Zp_X_16:  cpu.asl(cpu.zeroPageIndexedAddress())
		case LSR_Zp_X_56:  cpu.lsr(cpu.zeroPageIndexedAddress())
		case ROL_Zp_X_36:  cpu.rol(cpu.zeroPageIndexedAddress())
		case ROR_Zp_X_76:  cpu.ror(cpu.zeroPageIndexedAddress())
		case ASL_Abs_0E:   cpu.asl(cpu.absoluteAddress())
		case LSR_Abs_4E:   cpu.lsr(cpu.absoluteAddress())
		case ROL_Abs_2E:   cpu.rol(cpu.absoluteAddress())
		case ROR_Abs_6E:   cpu.ror(cpu.absoluteAddress())
		case ASL_Abs_X_1E: cpu.asl(cpu.absoluteIndexedAddress(cpu.X))
		case LSR_Abs_X_5E: cpu.lsr(cpu.absoluteIndexedAddress(cpu.X))
		case ROL_Abs_X_3E: cpu.rol(cpu.absoluteIndexedAddress(cpu.X))
		case ROR_Abs_X_7E: cpu.ror(cpu.absoluteIndexedAddress(cpu.X))
		// TAX TAY TXA TYA TSX TXS
		case TAX_AA: cpu.transferRegister(cpu.A, &cpu.X)
		case TAY_A8: cpu.transferRegister(cpu.A, &cpu.Y)
		case TXA_8A: cpu.transferRegister(cpu.X, &cpu.A)
		case TYA_98: cpu.transferRegister(cpu.Y, &cpu.A)
		case TSX_BA: cpu.transferRegister(cpu.Stack, &cpu.X)
		case TXS_9A: cpu.Stack = cpu.X; // TXS doesn't affect status flags
		// AND EOR ORA BIT
		case AND_Imm_29:   cpu.and(cpu.immediate())
		case EOR_Imm_49:   cpu.eor(cpu.immediate())
		case ORA_Imm_09:   cpu.ora(cpu.immediate())
		case AND_Zp_25:    cpu.and(cpu.zeroPage())
		case EOR_Zp_45:    cpu.eor(cpu.zeroPage())
		case ORA_Zp_05:    cpu.ora(cpu.zeroPage())
		case BIT_Zp_24:    cpu.bit(cpu.zeroPage())
		case AND_Zp_X_35:  cpu.and(cpu.zeroPageIndexed(cpu.X))
		case EOR_Zp_X_55:  cpu.eor(cpu.zeroPageIndexed(cpu.X))
		case ORA_Zp_X_15:  cpu.ora(cpu.zeroPageIndexed(cpu.X))
		case AND_Abs_2D:   cpu.and(cpu.absolute())
		case EOR_Abs_4D:   cpu.eor(cpu.absolute())
		case ORA_Abs_0D:   cpu.ora(cpu.absolute())
		case BIT_Abs_2C:   cpu.bit(cpu.absolute())
		case AND_Abs_X_3D: cpu.and(cpu.absoluteIndexed(cpu.X))
		case EOR_Abs_X_5D: cpu.eor(cpu.absoluteIndexed(cpu.X))
		case ORA_Abs_X_1D: cpu.ora(cpu.absoluteIndexed(cpu.X))
		case AND_Abs_Y_39: cpu.and(cpu.absoluteIndexed(cpu.Y))
		case EOR_Abs_Y_59: cpu.eor(cpu.absoluteIndexed(cpu.Y))
		case ORA_Abs_Y_19: cpu.ora(cpu.absoluteIndexed(cpu.Y))
		case AND_IndexIndirX_21: AND(cpu.indexedIndirectX())
		case EOR_IndexIndirX_41: EOR(cpu.indexedIndirectX())
		case ORA_IndexIndirX_01: ORA(cpu.indexedIndirectX())
		case AND_IndirIndexY_31: AND(cpu.indirectIndexedY())
		case EOR_IndirIndexY_51: EOR(cpu.indirectIndexedY())
		case ORA_IndirIndexY_11: ORA(cpu.indirectIndexedY())
		// Branches
		case BEQ_Relative_F0: Branch( cpu, cpu->status.zero, 1, operand );
		case BNE_Relative_D0: Branch( cpu, cpu->status.zero, 0, operand );
		case BMI_Relative_30: Branch( cpu, cpu->status.negative, 1, operand );
		case BPL_Relative_10: Branch( cpu, cpu->status.negative, 0, operand );
		case BCS_Relative_B0: Branch( cpu, cpu->status.carry, 1, operand );
		case BCC_Relative_90: Branch( cpu, cpu->status.carry, 0, operand );
		case BVS_Relative_70: Branch( cpu, cpu->status.overflow, 1, operand );
		case BVC_Relative_50: Branch( cpu, cpu->status.overflow, 0, operand );
		// Status flags
		case SEC_38: cpu->status.carry = 1;
		case CLC_18: cpu->status.carry = 0;
		case CLD_D8: cpu->status.decimal_mode = 0;
		case SED_F8: cpu->status.decimal_mode = 1;
		case SEI_78: cpu->status.interrupt_disable = 1;
		case CLI_58: cpu->status.interrupt_disable = 0;
		case CLV_B8: cpu->status.overflow = 0;
		// Stack
		case PHP_08: PHP( cpu );
		case PHA_48: push( cpu, cpu->a );
		case PLP_28: PLP( cpu );
		case PLA_68: PLA( cpu );
		// Misc
		case JMP_Abs_4C: JMPabs( cpu, operand );
		case JMP_Indirect_6C: JMPind( cpu, operand );
		case JSR_20: JSR( cpu, operand );
		case RTS_60: RTS( cpu );
		case RTI_40: RTI( cpu );
		case NOP_EA:
		case BRK_00: cpu->pc += 2; IRQ( cpu, 1 );
	}
}