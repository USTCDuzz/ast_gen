package astruct

import "fmt"

type SsString [SsLen]byte

// Ss Size Define
const (
	SsMaxLen = 10
	SsLen    = SsMaxLen + 1
	SsLenPos = SsLen - 1
)

// NewSsString
func NewSsString(ss string) SsString {
	s := SsString{}
	s.FromString(ss)
	return s
}

// IsValid
func (s *SsString) IsValid() bool {
	return s != nil && s[SsLenPos] != 0
}

// FromBytes
func (s *SsString) FromString(ss string) {
	if s == nil {
		return
	}
	if len(ss) > SsMaxLen {
		fmt.Printf("invalid len %v\n", len(ss))
		return
	}
	copy(s[:len(ss)], ss)
	s[SsLenPos] = byte(len(ss))
}

// ToBytes
func (s *SsString) ToString() string {
	if !s.IsValid() {
		return ""
	}
	return string(s[:s[SsLenPos]])
}
