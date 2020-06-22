package astruct

import "gogen/bstruct"

// ClearBb 禁用Bb
func (s *A) ClearBb() {
	if s == nil {
		return
	}
	s.Bb.Reset()
	s.Bb.Disable()
}

// EnableBb 使能Bb
func (s *A) EnableBb() {
	if s == nil {
		return
	}
	s.Bb.Enable()
}

// GetBb 获取Bb的指针
func (s *A) GetBb() *bstruct.B {
	if s != nil && s.Bb.Valid() {
		return &s.Bb
	}
	return nil
}
