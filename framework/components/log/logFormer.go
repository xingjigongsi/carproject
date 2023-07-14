package log

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

func JsonFormer(level LogLevel, message string, files map[string]interface{}) ([]byte, error) {
	buffer := bytes.Buffer{}
	segment := "\t"
	t := &LogService{}
	buffer.WriteString(t.Transtoint(level))
	buffer.WriteString(segment)
	buffer.WriteString(time.Now().Format("2006-01-02 03:04:05"))
	buffer.WriteString(segment)
	buffer.WriteString(message)
	buffer.WriteString(segment)
	if marshal, err := json.Marshal(files); err == nil {
		buffer.Write(marshal)
	} else {
		return nil, errors.Wrap(err, "files 解析出错")
	}
	buffer.WriteString("\r\n")
	return buffer.Bytes(), nil
}

func StrFormer(level LogLevel, message string, files map[string]interface{}) ([]byte, error) {
	builder := bytes.NewBuffer([]byte{})
	t := &LogService{}
	builder.WriteString(t.Transtoint(level))
	segment := "\t"
	builder.WriteString(segment)
	builder.WriteString(time.Now().Format("2006-01-02 03:04:05"))
	builder.WriteString(segment)
	builder.WriteString(message)
	builder.WriteString(segment)
	builder.WriteString(fmt.Sprint(files))
	builder.WriteString(segment)
	builder.WriteString("\r\n")
	return builder.Bytes(), nil

}
