package main

import (
	"fmt"
	"strings"
)

func main() {	
	PrintPath("/a/./b/../../c/")
	PrintPath("/home/")
	PrintPath("/../")
	PrintPath("/dir1/dir2/../dir3/")
}

func PrintPath(path string) {
	sanitizedPath := simplifyPath(path)
	fmt.Printf("---\ninput: %v\noutput: %v\n---\n", path, sanitizedPath)
}

func simplifyPath(path string) string {
		strs := strings.Split(path, "/")
    sanitized := make([]string, 0)
    
    for _, v := range strs {
        word := strings.TrimSpace(v)
        
        if word != "." && word != "/" && word != " " && word != "" && word != ".." {
            sanitized = append(sanitized, fmt.Sprintf("/%v", word)) 
        }
        
        if word == ".." && len(sanitized) > 0 {
            length := len(sanitized) - 1
            sanitized = sanitized[:length]
        } 
    }
    
    if len(strs) != 0 && len(sanitized) == 0 {
        sanitized = append(sanitized, "/")
    }

    return strings.Join(sanitized, "")
}
