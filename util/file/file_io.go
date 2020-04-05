package FileUtils

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type FileIO struct {
}

func (f *FileIO) ReadStrLines(filePath string) ([]string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")
	return lines, nil
}

func (f *FileIO) ReadByteLines(filePath string) ([][]byte, error) {
	var line bytes.Buffer
	fi, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fi.Close()
	buf := bufio.NewReader(fi)
	lines := make([][]byte, 0)
	for {
		data, err := buf.ReadBytes('\n')
		data = bytes.TrimSuffix(data, []byte("\n"))
		if err != nil {
			if bufio.ErrBufferFull == err {
				line.Write(data)
				continue
			}
		}
		if line.Len() > 0 {
			line.Write(data)
			lines = append(lines, line.Bytes())
			line.Reset()
		} else {
			lines = append(lines, data)
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}

	return lines, nil
}

func (f *FileIO) WriteFile(path string, content []byte, appendFlag bool) (int, error) {
	var err error
	var fi *os.File
	if appendFlag {
		fi, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	} else {
		fi, err = os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	}

	if err != nil {
		return 0, err
	}
	defer fi.Close()
	writerLength, err := fi.Write(content)
	if err != nil {
		return writerLength, err
	}
	return writerLength, nil
}

func (f *FileIO) WriteFileHead(path string, content []byte) error {

	var buf bytes.Buffer
	fileContent, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	buf.Write(content)
	buf.Write(fileContent)
	_, err = f.WriteFile(path, buf.Bytes(), false)
	if err != nil {
		return err
	}
	return nil

}

func (f *FileIO) UpdateFileByLine(path string, content []byte, lineNumber int) error {
	var err error
	var fi *os.File
	var buf bytes.Buffer
	byteSlice, err := f.ReadByteLines(path)
	if err != nil {
		return err
	}

	defer fi.Close()
	if (0 < lineNumber) && (lineNumber < len(content)-1) {
		byteSlice[lineNumber-1] = content
	} else {
		return errors.New("the number of line is wrong")
	}

	for _, b := range byteSlice {
		buf.Write(b)
		buf.WriteByte('\n')
	}

	fi, err = os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	_, err = fi.Write(buf.Bytes())
	if err != nil {
		return err
	}
	return nil

}

func (f *FileIO) FindWithPrefix(content []byte, prefix, end string) string {

	e := bytes.Index(content, []byte("package"))
	contentSub := content[:e]
	s := bytes.Index(contentSub, []byte("// +build"))
	if s <= 0 {
		return ""
	}
	contentSub = content[s:e]
	e = bytes.IndexByte(contentSub, '\n')
	return (string(contentSub[:e]))
}
