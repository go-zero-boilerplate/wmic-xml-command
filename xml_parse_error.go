package wmic_command

import "fmt"

type XmlParseError struct {
	XmlContent   []byte
	UnmarshalErr error
}

func (x *XmlParseError) Error() string {
	return fmt.Sprintf("Failed to parse WMIC xml. Error: %s. Output: %s", x.UnmarshalErr.Error(), string(x.XmlContent))
}
