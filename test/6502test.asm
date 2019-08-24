
;
; Cpu 6502 test ROM
; -----------------
; Tests all legal instructions, IRQ and NMI
; Doesn't test all addressing modes
; Doesn't test nor use decimal mode
; (c) 2014, Ernesto Borio, ernestoborio@gmail.com
;
; Use the following special addresses in your emulator:
; 
; $10: Number of the last test ran (1 byte)
; $500: If PC reaches here, all tests ran successfuly;
; $600: If PC reaches here, the last test ran failed
; $700: Reset or point PC here to start the tests
; $A000: When PC reaches here, your emulator should trigger an IRQ
; $B000: When PC reaches here, your emulator should trigger an NMI
;
; If your emulator doesn't trap $500 or $600, the program will just
; loop forever on those addresses.
; When a test fails no further tests will be ran and it will jump to $600
; If PC reaches $600, read zero page address $10 to see which test failed.
;
; Assemble this program with DASM http://dasm-dillon.sourceforge.net/
; 
; dasm "6502test.asm" -f3 -o"6502test.bin"

   processor 6502

; -----------------------------------------------------------
   ORG 0
indx_indr_hi_wrap:
indr_indx_ji_wrap:
   .byte >ind_ind_target_wrap

   ORG $10
; Store the current test number to know which one failed
current_test_no: .byte 0
zp_test_var: .byte $FF
   
   ORG $40
indx_indr:
   .byte 0, 0, 0
   .word indx_indr_target

   ORG $80
indr_indx:
   .word indr_indx_target
   
   ORG $FB
indx_indr_lo_wrap:
   .byte 0, 0, 0, 0
indr_indx_lo_wrap:
   .byte <ind_ind_target_wrap

; -----------------------------------------------------------
; This address is used by opcodes that don't use zero page
   ORG $200
test_address: .byte $FF

   ORG $220
indx_indr_target: .byte $BD

   ORG $240
ind_ind_target_wrap:
   .byte $AC ; For indexed indirect X
   .byte 0, 0, 0, 0, 0
   .byte $DE ; For indirect indeced Y

   ORG $2F0
indr_indx_target: 
   .byte 0, 0, 0, 0, 0
   .byte $ED
   
   ORG $310
indr_indx_target_cross_page:
   .byte $CB

; -----------------------------------------------------------
   ORG $500
; All tests ran correctly
Success:
   JMP Success

; -----------------------------------------------------------
   ORG $600
; current_test_no now holds the number of the test that failed
Fail:
   JMP Fail

; -----------------------------------------------------------
   ORG $700
; Tests begin here
Begin_tests:
   NOP
   NOP
   CLD ; Decimal mode won't be tested
   SEI ; Disable interrupts until they're tested
   LDX #$FF
   TXS ; Init the stack at $1FF
   NOP
   NOP
; -----------------------------------------------------------

test_1_LDA:
   LDA #1
   STA current_test_no
   CMP #1
   BEQ test_1_ok
   JMP Fail
test_1_ok:

; ----------------------------------
test_2_LDX:
   LDX #2
   STX current_test_no
   CPX #2
   BEQ test_2_ok
   JMP Fail
test_2_ok:

; ----------------------------------
test_3_LDY:
   LDY #3
   STY current_test_no
   CPY #3
   BEQ test_3_ok
   JMP Fail
test_3_ok:

; ----------------------------------
test_4_STAXY: ; STA, STX, STY
   LDA #4
   STA current_test_no
   STA test_address
   CMP test_address
   BNE test_4_fail
   LDX #$EA
   STX test_address
   CPX test_address
   BNE test_4_fail
   LDY #$AE
   STY test_address
   CPY test_address
   BNE test_4_fail
   JMP test_4_ok
test_4_fail:
   JMP Fail
test_4_ok:

; ----------------------------------
test_5_TAXY: ; TAX TXA TAY TYA TXS TSX
   LDA #5
   STA current_test_no
   LDX #0
   TAX
   CPX #5
   BNE test_5_fail
   LDY #0
   TAY
   CPY #5
   BNE test_5_fail
   LDA #0
   LDX #$EB
   TXA
   CMP #$EB
   BNE test_5_fail
   LDY #$CD
   TYA
   CMP #$CD
   BNE test_5_fail
   TXS
   LDX #$11
   TSX
   CPX #$EB
   BEQ test_5_ok
test_5_fail:
   JMP Fail
test_5_ok:
   LDX #$FF
   TXS ; restore sp
   
; ----------------------------------
test_6_AND_EOR_ORA_BIT:
   LDA #6
   STA current_test_no
   LDA #$F8
   AND #$8F
   CMP #$88
   BNE test_6_fail
   LDA #$F8
   EOR #$08
   CMP #$F0
   BNE test_6_fail
   LDA #$A0
   ORA #$0B
   CMP #$AB
   BNE test_6_fail
   LDA #05
   STA test_address
   LDA #04
   BIT test_address
   BNE test_6_ok
test_6_fail:
   JMP Fail
test_6_ok:

; ----------------------------------
test_7_INX_DEX_INY_DEY:
   LDA #7
   STA current_test_no
   LDX #$10
   INX
   CPX #$11
   BNE test_7_fail
   LDX #$A5
   DEX
   CPX #$A4
   BNE test_7_fail
   LDY #$EA
   INY
   CPY #$EB
   BNE test_7_fail
   LDY #$DF
   DEY
   CPY #$DE
   BEQ test_7_ok
test_7_fail:
   JMP Fail
test_7_ok:

; ----------------------------------
test_8_ADC_SBC:
; ADC
   LDA #8
   STA current_test_no
   SEC
   LDA #$45
   ADC #$52
   CMP #$98
   BNE test_8_fail
   CLC
   LDA #$81
   ADC #$81
   BVC test_8_fail
;SBC
   CLC
   LDA #$20
   SBC #$21
   CMP #$FE
   BNE test_8_fail
   CLC
   LDA #$81
   SBC #$7B
   BVC test_8_fail
   JMP test_8_ok
test_8_fail:
   JMP Fail
test_8_ok:

; ----------------------------------
test_9_ASL_LSR_ROL_ROR:
   LDA #9
   STA current_test_no
; ASL
   CLC
   LDA #$88
   STA zp_test_var
   ASL zp_test_var
   LDA zp_test_var
   CMP #$10
   BNE test_9_fail
   BCC test_9_fail
; LSR
   CLC
   LDA #$11
   STA zp_test_var
   LSR zp_test_var
   LDA zp_test_var
   CMP #$8
   BNE test_9_fail
   BCC test_9_fail
; ROL
   SEC
   LDA #$88
   STA zp_test_var
   ROL zp_test_var
   LDA zp_test_var
   CMP #$11
   BNE test_9_fail
   BCC test_9_fail
; ROR
   SEC
   LDA #$11
   STA zp_test_var
   ROR zp_test_var
   LDA zp_test_var
   CMP #$88
   BNE test_9_fail
   BCC test_9_fail
   JMP test_9_ok
test_9_fail:
   JMP Fail
test_9_ok:

; ----------------------------------
test_10_Stack ; PHA PHP PLA PLP TSX TXS
   LDA #10
   STA current_test_no

   LDX #$FF
   TXS ; Set the stack at $FF

   LDA #$7F
   ADC #$10 ; Set the overflow flag
   SEC
   LDA #0 ; Set zero and clear negative flag

   PHP
   TSX
   CPX #$FE ; Check if the stack pointer updated ok
   BNE test_10_fail
   
   LDA #$FF ; Clear zero and set negative flag
   CLC
   SEI
   SED
   CLV
   ; Clear the flags and see if a pull sets them again
   
   PLP
   BCC test_10_fail
   BVC test_10_fail
   BNE test_10_fail
   BMI test_10_fail

   LDA #%11011111 ; set all but the unused status flag
   PHA
   LDA #0
   PLP ; cpu = a
   PHP
   PLA ; a = cpu
   CMP #$FF ; the unused bit shoud be set now
   BNE test_10_fail

   JMP test_10_ok
test_10_fail:
   JMP Fail
test_10_ok:

; ----------------------------------
test_11_JSR_RTS:
   LDA #11
   STA current_test_no
   CLC
   LDA #0 ; clear negative, set zero
   JSR subroutine

   BCC test_11_fail
   BPL test_11_fail
   BEQ test_11_fail
   CMP #$BC
   BNE test_11_fail

   JMP test_11_ok
test_11_fail:
   JMP Fail
test_11_ok:

; ----------------------------------
test_12_BRK:
   LDA #12
   STA current_test_no
   CLV
   CLC
   LDA #0
   
   BRK
   NOP ; BRK's dummy byte signature

   BCS test_12_fail
   BNE test_12_fail
   BMI test_12_fail
   CMP #$BE ; test that value BE was set in IRQ routine
   BNE test_12_fail
   CPX #0 ; x = pushed break flag, should be 1
   BEQ test_12_fail

   JMP test_12_ok
test_12_fail:
   JMP Fail
test_12_ok:

; ----------------------------------
test_13_IRQ:
   LDA #13
   STA current_test_no
   CLI
   LDA #0 ; set a = 0 and test if IRQ changes its value
   CLC
   JMP trigger_IRQ

; This address should trigger a hardware IRQ on the emulator
   ORG $A000
trigger_IRQ:
   NOP
   NOP
   SEI
   BCS test_13_fail
   CMP #0
   BEQ test_13_fail
   CPX #0 ; x = pushed break flag, should be 0
   BEQ test_13_ok

test_13_fail:
   JMP Fail
test_13_ok:

; ----------------------------------
test_14_NMI:
   LDA #14
   STA current_test_no
   CLC
   LDA #0 ; set a = 0 and test if NMI changes its value
   JMP trigger_NMI
   
; This address should trigger an NMI on the emulator
   ORG $B000
trigger_NMI:
   NOP
   NOP
   BCS test_14_fail
   CMP #$BC
   BEQ test_14_ok

test_14_fail:
   JMP Fail
test_14_ok:

; ----------------------------------
; Oops I forgot to test INC & DEC
test_15_INC_DEC:
; INC
   LDA #15
   STA current_test_no
   LDA #$FF
   STA zp_test_var
   INC zp_test_var
   BMI test_15_fail
   BNE test_15_fail
   LDA zp_test_var   
   CMP #0
   BNE test_15_fail
; DEC 
   LDA #$80
   STA zp_test_var
   DEC zp_test_var
   BMI test_15_fail
   LDA zp_test_var
   CMP #$7F
   BNE test_15_fail
   JMP test_15_ok
   
test_15_fail:
   JMP Fail
test_15_ok

; ----------------------------------
test_16_JMP_Indirect:
   LDA #16
   STA current_test_no
   LDA 0
   JMP (test_16_pointer)
return2test_16:
   CMP #$FA
   BEQ test_16_ok

   JMP Fail
test_16_ok:

; ----------------------------------
test_17_Indexed_indirect_X:
   LDA #17
   STA current_test_no
   LDA #0
   LDX #3
   LDA (indx_indr,X)
   CMP #$BD
   BNE test_17_fail
   
   LDA #0
   LDX #4
   LDA (indx_indr_lo_wrap,X)
   CMP #$AC
   BEQ test_17_ok

test_17_fail:  
   JMP Fail
test_17_ok:

; ----------------------------------
test_18_Indirect_indexed_Y:
   LDA #18
   STA current_test_no
   LDA #0
   LDY #5
   LDA (indr_indx),Y
   CMP #$ED
   BNE test_18_fail

; Test crossing page 
   LDA #0
   LDY #$20
   LDA (indr_indx),Y
   CMP #$CB
   BNE test_18_fail

; Test with Zero page wrap around   
   LDA #0
   LDY #6
   LDA (indr_indx_lo_wrap),Y
   CMP #$DE
   BEQ test_18_ok
      
test_18_fail:  
   JMP Fail
test_18_ok:

; ----------------------------------
; Test JMP (indirect) with a pointer across page boundary
test_19:
   LDA #19
   STA current_test_no
   LDA #0
   JMP (expected_test_19_pointer)
return_test_19:
   CMP $FB
   BEQ test_19_ok
   
   JMP Fail
test_19_ok:

; -----------------------------------------------------------
; tests ended
   JMP Success
   STA current_test_no

; -----------------------------------------------------------
   ORG $C000
test_16_pointer:
   .word test_16_target
   
   ORG $C100
test_16_target:
   LDA #$FA
   JMP return2test_16

; -----------------------------------------------------------
; 6502 can't handle a page crossing in this situation, so although the pointer is at C2FF-C300,
; it reads the pointer's low byte from C2FF but wraps around the page and reads the pointer's high byte
; from C200 instead of the expected C300.
; The result is that the effective address will have the expected low byte, but a totally unexpected
; page, i.e. high byte. Programmers wouldn't expect this behavior, but it's what 6502 actually does.
; I guess they had to make sure their indirect pointers wouldn't fall in a page boundary.

   ORG $C200
actual_test_19_pointer: ; only the pointer's high byte is here
   .byte >actual_test_19_target ; $C5
   
   ORG $C2FF
expected_test_19_pointer:
   .word expected_test_19_target

   ORG $C433
expected_test_19_target:
   LDA #$AA
   JMP return_test_19

   ORG $C533
actual_test_19_target:
   LDA $FB
   JMP return_test_19
   
; -----------------------------------------------------------
   ORG $D000
subroutine:
   LDA #$BC ; set negative, clear zero
   SEC
   RTS

; -----------------------------------------------------------
   ORG $E000
IRQ_BRK_routine:
   PLA
   PHA
   AND #%10000
   TAX ; X holds whether the break flag was pushed as 1 or 0
   LDA #$BE
   SEC
   RTI

; -----------------------------------------------------------
   ORG $F000
NMI_routine:
   PLA
   PHA
   AND #%10000
   BEQ break_clear
   JMP Fail ; the break flag should've been pushed as 0

break_clear:
   LDA #$BC ; This A change should be kept
   SEC ; This flag change shouln't be kept
   RTI

; -----------------------------------------------------------
    ORG $FFFA
; reset vectors so the 6502 knows at what address to start running
Reset:
   .word NMI_routine     ; NMI interrupt, just ignore it
   .word Begin_tests     ; Reset interrupt, go to tests start
   .word IRQ_BRK_routine ; BRK/IRQ interrupt