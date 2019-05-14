package lib

import (
	in "geek_time_go_puzzlers/3-lib/lib3/lib/internal"
	"os"
)

func SayHello(name string) {
	in.SayHello(os.Stdout, name)
}
