// Code generated by "stringer -type=Opcode"; DO NOT EDIT.

package vm

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ADD-24]
	_ = x[ADDF-88]
	_ = x[ADDR-144]
	_ = x[AND-64]
	_ = x[CLEAR-180]
	_ = x[COMP-40]
	_ = x[COMPF-136]
	_ = x[COMPR-160]
	_ = x[DIV-36]
	_ = x[DIVF-100]
	_ = x[DIVR-156]
	_ = x[FIX-196]
	_ = x[FLOAT-192]
	_ = x[HIO-244]
	_ = x[J-60]
	_ = x[JEQ-48]
	_ = x[JGT-52]
	_ = x[JLT-56]
	_ = x[JSUB-72]
	_ = x[LDA-0]
	_ = x[LDB-104]
	_ = x[LDCH-80]
	_ = x[LDF-112]
	_ = x[LDL-8]
	_ = x[LDS-108]
	_ = x[LDT-116]
	_ = x[LDX-4]
	_ = x[LPS-208]
	_ = x[MUL-32]
	_ = x[MULF-96]
	_ = x[MULR-152]
	_ = x[NORM-200]
	_ = x[OR-68]
	_ = x[RD-216]
	_ = x[RMO-172]
	_ = x[RSUB-76]
	_ = x[SHIFTL-164]
	_ = x[SHIFTR-168]
	_ = x[SIO-240]
	_ = x[SSK-236]
	_ = x[STA-12]
	_ = x[STB-120]
	_ = x[STCH-84]
	_ = x[STF-128]
	_ = x[STI-212]
	_ = x[STL-20]
	_ = x[STS-124]
	_ = x[STSW-232]
	_ = x[STT-132]
	_ = x[STX-16]
	_ = x[SUB-28]
	_ = x[SUBF-92]
	_ = x[SUBR-148]
	_ = x[SVC-176]
	_ = x[TD-224]
	_ = x[TIO-248]
	_ = x[TIX-44]
	_ = x[TIXR-184]
	_ = x[WD-220]
}

const _Opcode_name = "LDALDXLDLSTASTXSTLADDSUBMULDIVCOMPTIXJEQJGTJLTJANDORJSUBRSUBLDCHSTCHADDFSUBFMULFDIVFLDBLDSLDFLDTSTBSTSSTFSTTCOMPFADDRSUBRMULRDIVRCOMPRSHIFTLSHIFTRRMOSVCCLEARTIXRFLOATFIXNORMLPSSTIRDWDTDSTSWSSKSIOHIOTIO"

var _Opcode_map = map[Opcode]string{
	0:   _Opcode_name[0:3],
	4:   _Opcode_name[3:6],
	8:   _Opcode_name[6:9],
	12:  _Opcode_name[9:12],
	16:  _Opcode_name[12:15],
	20:  _Opcode_name[15:18],
	24:  _Opcode_name[18:21],
	28:  _Opcode_name[21:24],
	32:  _Opcode_name[24:27],
	36:  _Opcode_name[27:30],
	40:  _Opcode_name[30:34],
	44:  _Opcode_name[34:37],
	48:  _Opcode_name[37:40],
	52:  _Opcode_name[40:43],
	56:  _Opcode_name[43:46],
	60:  _Opcode_name[46:47],
	64:  _Opcode_name[47:50],
	68:  _Opcode_name[50:52],
	72:  _Opcode_name[52:56],
	76:  _Opcode_name[56:60],
	80:  _Opcode_name[60:64],
	84:  _Opcode_name[64:68],
	88:  _Opcode_name[68:72],
	92:  _Opcode_name[72:76],
	96:  _Opcode_name[76:80],
	100: _Opcode_name[80:84],
	104: _Opcode_name[84:87],
	108: _Opcode_name[87:90],
	112: _Opcode_name[90:93],
	116: _Opcode_name[93:96],
	120: _Opcode_name[96:99],
	124: _Opcode_name[99:102],
	128: _Opcode_name[102:105],
	132: _Opcode_name[105:108],
	136: _Opcode_name[108:113],
	144: _Opcode_name[113:117],
	148: _Opcode_name[117:121],
	152: _Opcode_name[121:125],
	156: _Opcode_name[125:129],
	160: _Opcode_name[129:134],
	164: _Opcode_name[134:140],
	168: _Opcode_name[140:146],
	172: _Opcode_name[146:149],
	176: _Opcode_name[149:152],
	180: _Opcode_name[152:157],
	184: _Opcode_name[157:161],
	192: _Opcode_name[161:166],
	196: _Opcode_name[166:169],
	200: _Opcode_name[169:173],
	208: _Opcode_name[173:176],
	212: _Opcode_name[176:179],
	216: _Opcode_name[179:181],
	220: _Opcode_name[181:183],
	224: _Opcode_name[183:185],
	232: _Opcode_name[185:189],
	236: _Opcode_name[189:192],
	240: _Opcode_name[192:195],
	244: _Opcode_name[195:198],
	248: _Opcode_name[198:201],
}

func (i Opcode) String() string {
	if str, ok := _Opcode_map[i]; ok {
		return str
	}
	return "Opcode(" + strconv.FormatInt(int64(i), 10) + ")"
}