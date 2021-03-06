package asm

/*
	Instructions

	The below instructions implement a simple Instruction interface that allows
	them to be combined, optimised and encoded into machine code.

	x86_64 is a bit interesting in that instructions with different sets of
	operands might require subtly different machine code opcodes, even though
	they do the same thing functionally speaking, so the below instructions
	make use of a thing called OpcodeMaps that will match the given arguments
	to an appropriate opcode. For more on OpcodeMaps see asm/opcodes/
*/

import (
	"github.com/bspaans/jit-compiler/asm/encoding"
	"github.com/bspaans/jit-compiler/asm/opcodes"
	"github.com/bspaans/jit-compiler/lib"
)

func ADD(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("add", opcodes.ADD, 2, dest, src)
}
func AND(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("and", opcodes.AND, 2, dest, src)
}
func CALL(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("call", opcodes.CALL, 1, dest)
}
func CMP(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("cmp", opcodes.CMP, 2, dest, src)
}
func CMP_immediate(v uint64, dest encoding.Operand) lib.Instruction {
	if reg, ok := dest.(*encoding.Register); ok && reg.Width() == lib.BYTE {
		return CMP(encoding.Uint8(v), dest)
	}
	return opcodes.OpcodesToInstruction("cmp", opcodes.CMP, 2, dest, encoding.Uint32(v))
}

// Convert signed integer to scalar double-precision floating point (float64)
func CVTSI2SD(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("cvtsi2sd", opcodes.CVTSI2SD, 2, dest, src)
}

// Convert double precision float to signed integer
func CVTTSD2SI(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("cvttsd2si", opcodes.CVTTSD2SI, 2, dest, src)
}

// Convert Byte to Word; al:ah = sign extend(ah)
func CBW() lib.Instruction {
	return opcodes.OpcodesToInstruction("cbw", []*encoding.Opcode{opcodes.CBW}, 0)
}

// Convert Word to Doubleword; dx:ax = sign extend(ax)
func CWD() lib.Instruction {
	return opcodes.OpcodesToInstruction("cwd", []*encoding.Opcode{opcodes.CWD}, 0)
}

// Convert Double word to Quadword; edx:eax = sign extend(eax)
func CDQ() lib.Instruction {
	return opcodes.OpcodesToInstruction("cdq", []*encoding.Opcode{opcodes.CDQ}, 0)
}

// Convert Quad word to double quad word; rdx:rax = sign extend(rax)
func CQO() lib.Instruction {
	return opcodes.OpcodesToInstruction("cqo", []*encoding.Opcode{opcodes.CQO}, 0)
}
func DEC(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("dec", opcodes.DEC, 1, dest)
}
func DIV(src encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("div", opcodes.DIV, 1, src)
}
func IDIV1(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("div", opcodes.IDIV1, 1, dest)
}
func IDIV2(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("div", opcodes.IDIV2, 2, dest, src)
}
func INC(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("inc", opcodes.INC, 1, dest)
}
func JA(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("ja", opcodes.JA, 1, dest)
}
func JAE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jae", opcodes.JAE, 1, dest)
}
func JB(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jb", opcodes.JB, 1, dest)
}
func JBE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jbe", opcodes.JBE, 1, dest)
}
func JE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("je", opcodes.JE, 1, dest)
}
func JG(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jg", opcodes.JG, 1, dest)
}
func JGE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jge", opcodes.JGE, 1, dest)
}
func JL(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jl", opcodes.JL, 1, dest)
}
func JLE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jle", opcodes.JLE, 1, dest)
}
func JNA(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jna", opcodes.JNA, 1, dest)
}
func JNAE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jnae", opcodes.JNAE, 1, dest)
}
func JNB(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jnb", opcodes.JNB, 1, dest)
}
func JNBE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jnbe", opcodes.JNBE, 1, dest)
}
func JNE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jne", opcodes.JNE, 1, dest)
}
func JNG(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jng", opcodes.JNG, 1, dest)
}
func JNGE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jnge", opcodes.JNGE, 1, dest)
}
func JNL(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jnl", opcodes.JNL, 1, dest)
}
func JNLE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jnle", opcodes.JNLE, 1, dest)
}
func JMP(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("jmp", opcodes.JMP, 1, dest)
}
func LEA(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("lea", opcodes.LEA, 2, dest, src)
}
func MOV(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("mov", opcodes.MOV, 2, dest, src)
}
func MOV_immediate(v uint64, dest encoding.Operand) lib.Instruction {
	if reg, ok := dest.(*encoding.Register); ok && reg.Width() == lib.BYTE {
		return MOV(encoding.Uint8(v), dest)
	}
	if reg, ok := dest.(*encoding.Register); ok && reg.Width() == lib.WORD {
		return MOV(encoding.Uint16(v), dest)
	}
	if v < (1 << 32) {
		return MOV(encoding.Uint32(v), dest)
	}
	return MOV(encoding.Uint64(v), dest)
}

// Move with sign-extend
func MOVSX(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("movsx", opcodes.MOVSX, 2, dest, src)
}

// Move with zero-extend
func MOVZX(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("movzx", opcodes.MOVZX, 2, dest, src)
}
func IMUL1(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("imul", opcodes.IMUL1, 1, dest)
}
func IMUL2(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("imul", opcodes.IMUL2, 2, dest, src)
}
func MUL(src encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("mul", opcodes.MUL, 1, src)
}
func OR(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("or", opcodes.OR, 2, dest, src)
}
func POP(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("pop", opcodes.POP, 1, dest)
}
func PUSH(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("push", opcodes.PUSH, 1, dest)
}
func PUSHFQ() lib.Instruction {
	return opcodes.OpcodeToInstruction("pushfq", opcodes.PUSHFQ, 0)
}
func RETURN() lib.Instruction {
	return opcodes.OpcodeToInstruction("return", opcodes.RETURN, 0)
}
func SETA(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("seta", opcodes.SETA, 1, dest)
}
func SETAE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("setae", opcodes.SETAE, 1, dest)
}
func SETB(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("setb", opcodes.SETB, 1, dest)
}
func SETBE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("setbe", opcodes.SETBE, 1, dest)
}
func SETC(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("setc", opcodes.SETC, 1, dest)
}
func SETE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("sete", opcodes.SETE, 1, dest)
}
func SETL(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("setl", opcodes.SETL, 1, dest)
}
func SETLE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("setle", opcodes.SETLE, 1, dest)
}
func SETG(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("setg", opcodes.SETG, 1, dest)
}
func SETGE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("setge", opcodes.SETGE, 1, dest)
}
func SETNE(dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("setne", opcodes.SETNE, 1, dest)
}
func SUB(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("sub", opcodes.SUB, 2, dest, src)
}
func SHL(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("shl", opcodes.SHL, 2, dest, src)
}
func SHR(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("shr", opcodes.SHR, 2, dest, src)
}
func SYSCALL() lib.Instruction {
	return opcodes.OpcodeToInstruction("syscall", opcodes.SYSCALL, 0)
}
func XOR(src, dest encoding.Operand) lib.Instruction {
	return opcodes.OpcodesToInstruction("xor", opcodes.XOR, 2, dest, src)
}
