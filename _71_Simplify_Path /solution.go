package _71_Simplify_Path_

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	var stack []string
	parts := strings.Split(path, "/")

	for _, p := range parts {
		if p == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if p == "." || len(p) == 0 {
			continue
		} else {
			stack = append(stack, p)
		}
	}

	joined := strings.Join(stack, "/")
	if len(joined) == 0 || joined[0] != '/' {
		return fmt.Sprintf("/%s", joined)
	}
	return joined
}
