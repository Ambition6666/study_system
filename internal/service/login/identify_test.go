package login

import (
	"testing"
)

// 测试密码能否加密
func TestEncrypt(t *testing.T) {
	s := Encrypt("756863")
	t.Log(s)
}
