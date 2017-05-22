package strings2

import "github.com/albertjin/ec"

func SplitSpaceBracket1(s string, space, bracketLeft, bracketRight rune) (ranges []int, err error) {
    const StateIdle = 0
    const StateWord = 1
    const StateInBracket = 2

    state := StateIdle
    for i, ch := range s {
        switch state {
        case StateIdle:
            switch ch {
            case bracketLeft:
                state = StateInBracket
                ranges = append(ranges, i + 1)
            case bracketRight:
                err = ec.NewErrorf("unexpected right bracket at %v", i)
                return
            case space:
            default:
                state = StateWord
                ranges = append(ranges, i)
            }
        case StateWord:
            switch ch {
            case space:
                state = StateIdle
                ranges = append(ranges, i)
            case bracketLeft:
                fallthrough
            case bracketRight:
                err = ec.NewErrorf("unexpected bracket at %v", i)
                return
            }
        case StateInBracket:
            switch ch {
            case bracketRight:
                state = StateIdle
                ranges = append(ranges, i)
            case bracketLeft:
                err = ec.NewErrorf("unexpected left bracket at %v", i)
                return
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
