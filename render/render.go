package render

import (
	"encoding/json"
	"io"
)

func JSON(w io.Writer, data interface{}) (err error) {
	json, err := json.Marshal(data)
	if err != nil {
		return
	}

	w.Write(json)
	return
}
