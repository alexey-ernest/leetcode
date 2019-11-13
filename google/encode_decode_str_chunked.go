import (
    "strings"
    "fmt"
)

type Codec struct {}

// Encodes a list of strings to a single string.
func (codec *Codec) Encode(strs []string) string {
    var sb strings.Builder
    for _, s := range strs {
        sb.WriteString(lentostr(len(s)))
        sb.WriteString(s)
    }
    
    return sb.String()
}

// Decodes a single string to a list of strings.
func (codec *Codec) Decode(strs string) []string {
    res := make([]string, 0)
    for i := 0; i < len(strs); i += 1 {
        l := strtolen(strs[i:i+4])
        i += 4
        res = append(res, strs[i:i+l])
        i += l-1
    }
    
    return res
}

func lentostr(i int) string {
    bytes := make([]byte, 4)
    for k := 3; k >= 0; k -= 1 {
        bytes[k] = byte(i >> (uint(k-3) * 8) & 255)
    }
    
    return string(bytes)
}

func strtolen(str string) int {
    res := 0
    for i := 0; i < 4; i += 1 {
        res = res * 256 + int(str[i])
    }
    return res
}

// Your Codec object will be instantiated and called as such:
// var codec Codec
// codec.Decode(codec.Encode(strs));