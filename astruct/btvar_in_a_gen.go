package astruct

import "fmt"

type BtVarString [BtVarLen]byte

// BtVar Size Define
const (
	BtVarMaxLen = 12
	BtVarLen    = BtVarMaxLen + 1
	BtVarLenPos = BtVarLen - 1
)

// NewBtVarString
func NewBtVarString(btVar []byte) BtVarString {
	s := BtVarString{}
	s.FromBytes(btVar)
	return s
}

// IsValid
func (s *BtVarString) IsValid() bool {
	return s != nil && s[BtVarLenPos] != 0
}

// FromBytes
func (s *BtVarString) FromBytes(btVar []byte) {
	if s == nil {
		return
	}
	if len(btVar) > BtVarMaxLen {
		fmt.Printf("invalid len %v\n", len(btVar))
		return
	}
	copy(s[:len(btVar)], btVar)
	s[BtVarLenPos] = byte(len(btVar))
}

// ToBytes
func (s *BtVarString) ToBytes() []byte {
	if !s.IsValid() {
		return nil
	}
	ret := make([]byte, s[BtVarLenPos], s[BtVarLenPos])
	copy(ret, s[:s[BtVarLenPos]])
	return ret
}
