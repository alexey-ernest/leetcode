import "strings"
import "fmt"

func braceExpansionII(expression string) []string {
    stack := make([]string, 0, len(expression))
    
    // processing chars
    str := ""
    for i := 0; i < len(expression); i += 1 {
        ch := expression[i]
        
        if ch == '}' || ch == '{' || ch == ',' {
            if str != "" {
                stack = append(stack, str)
                stack = mergeTop(stack)
                str = ""
            }
        } else {
            str += string(ch)
        }
        
        if ch == '{' || ch == ',' {
            stack = append(stack, string(ch))
        }
        
        if ch == '}' {
                        
            // building a set enclosed in braces
            set := make(map[string]struct{})
            for j := len(stack)-1; j >= 0; j -= 1 {
                if stack[j] == "{" {
                    stack = stack[:j]
                    break
                }
                if stack[j] == "," {
                    continue
                }
                
                // check if char or set
                mset := deserialize(stack[j])
                for k := range mset {
                    set[k] = struct{}{}    
                }
            }
            
            stack = append(stack, serialize(set))
            
            // merging top (multiply with previous)
            stack = mergeTop(stack)
        }
    }
    
    // processing sets
    if len(stack) > 1 {
        panic("wrong input")
    }
    
    var fields []string
    if len(stack) == 0 {
        // just one field in expression
        fields = append(fields, str)
    } else {
        resset := deserialize(stack[0])
    
        fields = make([]string, 0, len(resset))
        for k := range resset {
            if str != "" {
                k += str
            }
            fields = append(fields, k)
        }
    }
    
    // sorting fields in ascending order
    sort.Strings(fields)
    return fields
}

func mergeTop(stack []string) []string {
    if len(stack) == 0 {
        return stack
    }
    
    //fmt.Printf("stack to merge %+v\n", stack)
    
    // multiply with any previous char or set down on the stack
    top := stack[len(stack)-1]
    set := deserialize(top)
    c := 0
    for j := len(stack)-2; j >= 0; j -= 1 {
        if stack[j] == "," || stack[j] == "{" {
            if c > 0 {
                stack = stack[:j+1]    
            }
            break
        }

        // check if it is set or char
        mset := deserialize(stack[j])
        c += 1

        // building cartesian product
        newset := make(map[string]struct{})
        for s1 := range mset {
            for s2 := range set {
                newset[s1+s2] = struct{}{}
            }
        }
        set = newset

        if j == 0 {
            // we've processed all the items in the stack
            stack = stack[:0]
        }
    }

    if c > 0 {
        stack = append(stack, serialize(set))
    }
    
    return stack
}

func serialize(m map[string]struct{}) string {
    keys := make([]string, 0)
    for k := range m {
        keys = append(keys, k)
    }

    return strings.Join(keys, ",")
}

func deserialize(s string) map[string]struct{} {
    keys := strings.Split(s, ",")
    m := make(map[string]struct{})
    for _, k := range keys {
        m[k] = struct{}{}
    }
    
    return m
}