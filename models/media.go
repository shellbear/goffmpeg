package models

import (
	"fmt"
	"io"
	"reflect"
)

type Mediafile struct {
	threadSet        bool
	threads          int
	inputPath        string
	inputPipe        bool
	metadata         Metadata
	inputPipeReader  *io.PipeReader
	inputPipeWriter  *io.PipeWriter
	outputPipe       bool
	outputPipeReader *io.PipeReader
	outputPipeWriter *io.PipeWriter
	outputFiles      []*OutputFile
}

/*** SETTERS ***/
func (m *Mediafile) SetThreads(v int) {
	m.threadSet = true
	m.threads = v
}

func (m *Mediafile) SetInputPath(val string) {
	m.inputPath = val
}

func (m *Mediafile) SetInputPipe(val bool) {
	m.inputPipe = val
}

func (m *Mediafile) SetInputPipeReader(r *io.PipeReader) {
	m.inputPipeReader = r
}

func (m *Mediafile) SetInputPipeWriter(w *io.PipeWriter) {
	m.inputPipeWriter = w
}

/*** GETTERS ***/

func (m *Mediafile) InputPath() string {
	return m.inputPath
}

func (m *Mediafile) InputPipe() bool {
	return m.inputPipe
}

func (m *Mediafile) InputPipeReader() *io.PipeReader {
	return m.inputPipeReader
}

func (m *Mediafile) InputPipeWriter() *io.PipeWriter {
	return m.inputPipeWriter
}

func (m *Mediafile) Threads() int {
	return m.threads
}

/** OPTS **/
func (m *Mediafile) ToStrCommand() []string {
	var strCommand []string

	opts := []string{
		"InputPath",
		"InputPipe",
		"Threads",
		"OutputPipe",
	}

	for _, name := range opts {
		opt := reflect.ValueOf(m).MethodByName(fmt.Sprintf("Obtain%s", name))
		if (opt != reflect.Value{}) {
			result := opt.Call([]reflect.Value{})

			if val, ok := result[0].Interface().([]string); ok {
				strCommand = append(strCommand, val...)
			}
		}
	}

	for _, outputFile := range m.outputFiles {
		strCommand = append(strCommand, outputFile.ToStrCommand()...)
	}

	return strCommand
}

func (m *Mediafile) ObtainInputPath() []string {
	if m.inputPath != "" {
		return []string{"-i", m.inputPath}
	}
	return nil
}

func (m *Mediafile) ObtainInputPipe() []string {
	if m.inputPipe {
		return []string{"-i", "pipe:0"}
	}
	return nil
}

func (m *Mediafile) ObtainThreads() []string {
	if m.threadSet {
		return []string{"-threads", fmt.Sprintf("%d", m.threads)}
	}
	return nil
}

func (m *Mediafile) AddOutputFile(outputFile *OutputFile) {
	m.outputFiles = append(m.outputFiles, outputFile)
}

func (m *Mediafile) SetMetadata(v Metadata) {
	m.metadata = v
}

func (m *Mediafile) Close() {
	if m.outputPipeWriter != nil {
		m.outputPipeWriter.Close()
	}

	if m.outputPipeReader != nil {
		m.outputPipeReader.Close()
	}
}

func (m *Mediafile) SetOutputPipe(val bool) {
	m.outputPipe = val
}

func (m *Mediafile) SetOutputPipeReader(r *io.PipeReader) {
	m.outputPipeReader = r
}

func (m *Mediafile) SetOutputPipeWriter(w *io.PipeWriter) {
	m.outputPipeWriter = w
}

func (m *Mediafile) ObtainOutputPipe() []string {
	if m.outputPipe {
		return []string{"pipe:1"}
	}
	return nil
}

func (m *Mediafile) OutputPipe() bool {
	return m.outputPipe
}

func (m *Mediafile) OutputPipeReader() *io.PipeReader {
	return m.outputPipeReader
}

func (m *Mediafile) OutputPipeWriter() *io.PipeWriter {
	return m.outputPipeWriter
}

func (m *Mediafile) OutputFiles() []*OutputFile {
	return m.outputFiles
}

func (m *Mediafile) Metadata() Metadata {
	return m.metadata
}
