package dir

import "strings"

// IsSubDirOf s が dir 配下のパスであるかをチェックする
func IsSubDirOf(s, dir string) bool {
	s = strings.TrimPrefix(s, "/")
	parts := strings.Split(s, "/")
	for _, v := range parts {
		if v == dir {
			return true
		}
	}
	return false
}
