package strings2

import "github.com/albertjin/ec"

func SplitSpaceBracket1(s string, space, bracketLeft, bracketRight rune) (ranges []int, err error) {
    const StateIdle = 0
    const StateWord = 1
    const StateInBracket = 2

    state := StateIdle
    for i, r := range s {
        switch state {
        case StateIdle:
            switch r {
            case bracketLeft:
                state = StateInBracket
                ranges = append(ranges, i+1)
            case bracketRight:
                err = ec.NewErrorf("unexcepted bracket right at %v", i)
                return
            case space:
            default:
                state = StateWord
                ranges = append(ranges, i)
            }
        case StateWord:
            switch r {
            case bracketLeft:
                fallthrough
            case bracketRight:
                err = ec.NewErrorf("unexcepted bracket at %v", i)
                return
            case space:
                ranges = append(ranges, i)
                state = StateIdle
            default:
            }
        case StateInBracket:
            switch r {
            case bracketLeft:
                err = ec.NewErrorf("unexcepted left bracket at %v", i)
                return
            case bracketRight:
                ranges = append(ranges, i)
                state = StateIdle
            default:
            }
        }
    }

    switch state {
    case StateWord:
        ranges = append(ranges, len(s))
    case StateInBracket:
        err = ec.NewError("right bracket is expected at the end of the string")
    }
    return
}
