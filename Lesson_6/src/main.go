package main	

import (
	"fmt"
	"time"
	"io"
	"os"
)

func main() {
	sub := HwSubmitted {3456, "please take a look at my homework"}
	acc := HwAccepted {3456,5}

	LogOtusEvent(sub, os.Stdout)
	LogOtusEvent(acc, os.Stdout)
}

type HwAccepted struct {
	Id int
	Grade int
}

type HwSubmitted struct {
	Id int
	Comment string
}

type OtusEvent interface{
	WriteLog(w io.Writer)
}

func LogOtusEvent(e OtusEvent, w io.Writer) {
	e.WriteLog(w)
}

func (hw HwAccepted) WriteLog(w io.Writer) {
	fmt.Fprintf(w, "%+v Accepted %+v %+v\n", time.Now().Format("2006-01-02"), hw.Id, hw.Grade)
}

func (sub HwSubmitted) WriteLog(w io.Writer) {
	fmt.Fprintf(w,"%+v Submitted %+v \"%+v\"\n", time.Now().Format("2006-01-02"), sub.Id, sub.Comment)
}


