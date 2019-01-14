package isa

// This virtual machine will never support the full T32 ISA and is intended only
// as an educational endeavor.

// Operation is a makeshift enum type
type Operation int

const (
	opAND Operation = iota // AND
	opEOR                  // EOR
	opSUB                  // Subtract
	opRSB                  // RSB
	opADD                  // Add
	opADC                  // Add w/ carry
	opSBC                  // SBC
	opRSC                  // RSC
	opTST                  // Test
	opTEQ                  // TEQ
	opCMP                  // Compare
	opCMN                  // CMN
	opORR                  // ORR
	opMOV                  // Move
	opBIC                  // BIC
	opMVN                  // MVN
)
