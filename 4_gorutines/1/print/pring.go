package printDa

import "os"

func WriteString(msg string) {
	os.Stdout.Write([]byte(msg))
	os.Stdout.Write([]byte("\n"))
}
