package bar

import (
	"encoding/json"
	"fmt"
	"io"
)

type Config struct {
	String string
	Float  float64
	Bool   bool
}

func Run(out io.Writer, conf Config) error {
	b, err := json.MarshalIndent(conf, "", " ")
	if err != nil {
		return err
	}

	fmt.Fprintf(out, "Bar config:\n%s", b)

	return nil
}
