package str

import "unicode"

func ToSnake(camel string) (string, bool) {

	if len(camel) == 0 {
		return "", false
	}
	out := ""
	queue := ""
	for _, r := range camel {
		if unicode.IsUpper(rune(r)) {
			// Push caps to queue, append to output later
			queue += string(unicode.ToLower(rune(r)))
			continue
		}
		if len(queue) <= 0 {
			// no caps letters to write
			out += string(r)
			continue
		}

		if len(out) > 0 {
			// avoid writing '_' to front of output
			out += "_"
		}
		// convert queued caps and write to output
		if len(queue) > 1 {
			// case 'CONSTANTVar' -> 'constant_var'
			out += queue[:len(queue)-1]
			out += "_"
			out += string(queue[len(queue)-1])
		} else {
			// case 'Var' -> 'var'
			out += queue
		}
		queue = ""
		out += string(r)
	}
	return out, true
}
