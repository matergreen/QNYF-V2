package utils

import (
	"QnyfV2/config"
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadFile() []string {
	var content []string
	fileObj, _ := os.OpenFile(config.FILENAME, os.O_RDONLY, 0666)
	defer fileObj.Close()

	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		content = append(content, line)
		if err == io.EOF {
			break
		}
	}

	return content
}
