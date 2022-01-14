package foo

import (
	"encoding/json"
	"fmt"
	"github.com/corverroos/stingoftheviper/bar"
	"io"
)

type Config struct {
	Bar    bar.Config
	String string
	Float  float64
	Bool   bool
}

func Run(out io.Writer, conf Config) error {
	b, err := json.MarshalIndent(conf, "", " ")
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "Foo config:\n%s", b)

	return nil
}
