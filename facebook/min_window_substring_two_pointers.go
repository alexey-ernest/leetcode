import "fmt"

func minWindow(s string, t string) string {
    if len(t) == 0 {
        return ""
    }
    
    chars := make(map[byte]int)
    for i := 0; i < len(t); i += 1 {
        chars[t[i]]++
    }
    
    fstr := make([]charindex, 0)
    for i := 0; i < len(s); i += 1 {
        if _, ok := chars[s[i]]; ok {
            fstr = append(fstr, charindex{c: s[i], i: i})
        }
    }
    
    res := ""
    l, r := 0, 0
    counter := make(map[byte]int)
    formed := 0
    for ; r < len(fstr); r += 1 {
        ci := fstr[r]
        counter[ci.c]++
        
        if counter[ci.c] == chars[ci.c] {
            formed += 1
        }
        
        if counter[ci.c] >= chars[ci.c] {
            // trying to shrink the interval
            for ; l <= r; l += 1 {
                cl := fstr[l]
                if counter[cl.c] <= chars[cl.c] {
                    // we can't shrink
                    break
                }
                counter[cl.c]--
            }
            
            if formed < len(chars) {
                continue
            }
            
            if len(res) == 0 || fstr[r].i - fstr[l].i + 1 < len(res) {
                res = s[fstr[l].i:fstr[r].i+1]
            }
        }
    }
    
    return res
}

type charindex struct {
    c byte
    i int
}