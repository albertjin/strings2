package strings2

import (
    "strings"
)

func GetTextBlock(text, begin, end string, optionalEnding, includeBeginning, includeEnding bool) (block, remaining string, matched bool) {
    remaining = text
    a := strings.Index(text, begin)
    if a >= 0 {
        if !includeBeginning {
            a += len(begin)
        }
        u := text[a:]
        b := strings.Index(u, end)
        if b >= 0 {
            c := b
            d := b + len(end)
            if includeEnding {
                c = d
            }
            matched = true
            block = u[:c]
            remaining = u[d:]
        } else if optionalEnding {
            matched = true
            block = u
            remaining = ""
        }
    }
    return
}

func ProcessBlocks(text, begin, end string, optionalEnding, includeBeginning, includeEnding bool, handle func(string) bool) {
    var s string
    var matched bool
    for r := text; ; {
        s, r, matched = GetTextBlock(r, begin, end, optionalEnding, includeBeginning, includeEnding)
        if !matched || !handle(s) {
            break
        }
    }
}

func GetKeyValueMap(s string) (m map[string]string) {
    m = map[string]string{}
    for _, item := range strings.Split(s, ";") {
        kv := strings.SplitN(item, "=", 2)
        if len(kv) == 2 {
            m[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
        }
    }
    return
}
