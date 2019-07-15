package internal

import (
	"github.com/cheggaaa/pb/v3"
)

func progress() {
	bar := pb.StartNew(count)
}
