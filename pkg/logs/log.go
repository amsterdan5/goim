package logs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"
)

type LogField string

const (
	levelInfo  = "INFO"
	levelWarn  = "WARN"
	levelError = "ERROR"
)

type logger struct {
	once        sync.Once
	infoWriter  io.Writer
	warnWriter  io.Writer
	errorWriter io.Writer
	ctxKeys     []LogField // 链路追踪信息
}

var logStd = logger{
	once:        sync.Once{},
	infoWriter:  os.Stdout,
	warnWriter:  os.Stdout,
	errorWriter: os.Stdout,
}

// 初始化
func Init(infoF, warnF, errorF io.Writer, ctxKyes ...LogField) {
	logStd.once.Do(func() {
		logStd.infoWriter = infoF
		logStd.warnWriter = warnF
		logStd.errorWriter = errorF
		logStd.ctxKeys = ctxKyes
	})
}

// 日志内容体
type entry struct {
	level    string
	time     string
	file     string
	line     int
	function string
	msg      string
	fields   [][2]interface{}
}

func Ctx(ctx context.Context) *entry {
	pc := make([]uintptr, 2)
	_ = runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc)

	file, function, line := "", "", 0
	if frame, more := frames.Next(); more {
		function = frame.Function[strings.LastIndex(frame.Function, "/")+1:]
		file = path.Base(frame.File)
		line = frame.Line
	}

	fields := make([][2]interface{}, 0, len(logStd.ctxKeys)+2)
	for _, k := range logStd.ctxKeys {
		// 获取上下文内存储的信息
		if v := ctx.Value(k); v != nil {
			fields = append(fields, [2]interface{}{k, v})
		}
	}

	return &entry{
		time:     time.Now().Format("2006-01-02 15:04:05"),
		file:     file,
		function: function,
		line:     line,
		fields:   fields,
	}
}

// 单字段
func (e *entry) WithField(k string, v interface{}) *entry {
	e.fields = append(e.fields, [2]interface{}{k, v})
	return e
}

// 多字段
func (e *entry) WithFields(k1 string, v1 interface{}, k2 string, v2 interface{}, kv ...interface{}) *entry {
	e.fields = append(e.fields, [2]interface{}{k1, v1}, [2]interface{}{k2, v2})

	kvLen := len(kv)
	if kvLen == 0 {
		return e
	}

	// k v不配对
	if kvLen%2 != 0 {
		for i := 0; i < kvLen; i += 2 {
			if i+1 < kvLen {
				e.fields = append(e.fields, [2]interface{}{fmt.Sprint(kv[i]), kv[i+1]})
			} else {
				e.fields = append(e.fields, [2]interface{}{fmt.Sprint(kv[i]), "未知"})
			}
		}
	} else {
		for i := 0; i < kvLen; i += 2 {
			e.fields = append(e.fields, [2]interface{}{fmt.Sprint(kv[i]), kv[i+1]})
		}
	}

	return e
}

// 输出内容
func (e *entry) writer(w io.Writer) (int, error) {
	// 缓冲
	outbuf := outBufPool.Get().(*bytes.Buffer)
	outbuf.Reset()

	// 重新放回池
	defer func() {
		outBufPool.Put(outbuf)
	}()

	// 记录错误信息所在文件/行
	outbuf.WriteString(fmt.Sprintf("[%s] %s %s:%d %s %s", e.level, e.time, e.file, e.line, e.function, e.msg))

	// 没有自定义字段 直接返回
	fieldLen := len(e.fields)
	if fieldLen == 0 {
		outbuf.WriteString("\n")
		return w.Write(outbuf.Bytes())
	}

	outbuf.WriteString(" {")
	for i := 0; i < fieldLen; i++ {
		if i != 0 {
			outbuf.WriteString(",")
		}

		key, _ := e.fields[i][0].(LogField)
		outbuf.WriteString(fmt.Sprintf(`"%s:"`, key))

		switch value := e.fields[i][1].(type) {
		case bool, int, int8, int16, int32, int64, float32, float64, uint8, uint16, uint32, uint64, uint:
			outbuf.WriteString(fmt.Sprintf("%+v", value))
		case string:
			outbuf.WriteString(value)
		case *string:
			outbuf.WriteString(*value)
		case []byte:
			outbuf.Write(value)
		case *[]byte:
			outbuf.Write(*value)
		default:
			var buf bytes.Buffer
			j := json.NewEncoder(&buf)
			j.SetEscapeHTML(false)

			if err := j.Encode(value); err != nil {
				outbuf.WriteString(fmt.Sprintf("%+v", value))
			} else {
				b := buf.Bytes()
				bLen := len(b)
				if bLen > 0 {
					outbuf.Write(b[:bLen-1])
				} else {
					outbuf.Write(nil)
				}
			}
		}
	}

	outbuf.WriteString("}\n")
	return w.Write(outbuf.Bytes())
}

// 根据级别输出信息
func (e *entry) output(l string, format string, a ...interface{}) error {
	msg := ""
	if format == "" {
		msg = fmt.Sprintln(a...)
		msg = msg[:len(msg)-1] // 去除结尾符
	} else {
		msg = fmt.Sprintf(format, a...)
	}

	e.level = l
	e.msg = msg

	var outWriter io.Writer
	switch l {
	case levelInfo:
		outWriter = logStd.infoWriter
	case levelWarn:
		outWriter = logStd.warnWriter
	case levelError:
		outWriter = logStd.errorWriter
	default:
		outWriter = logStd.errorWriter
	}

	_, err := e.writer(outWriter)
	return err
}

func (e *entry) outputLn(l string, a ...interface{}) error {
	return e.output(l, "", a...)
}

func (e *entry) outputFmtLn(l, format string, a ...interface{}) error {
	return e.output(l, format, a...)
}

// 普通
func (e *entry) Info(a ...interface{}) {
	_ = e.outputLn(levelInfo, a...)
}

// 错误
func (e *entry) Error(a ...interface{}) {
	_ = e.outputLn(levelError, a...)
}

// 提醒
func (e *entry) Warn(a ...interface{}) {
	_ = e.outputLn(levelWarn, a...)
}

// 格式化普通日志
func (e *entry) InfoF(format string, a ...interface{}) {
	_ = e.outputFmtLn(levelInfo, format, a...)
}

// 格式化提醒日志
func (e *entry) WarnF(format string, a ...interface{}) {
	_ = e.outputFmtLn(levelWarn, format, a...)
}

// 格式化错误日志
func (e *entry) ErrorF(format string, a ...interface{}) {
	_ = e.outputFmtLn(levelError, format, a...)
}

var outBufPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 4<<10))
	},
}
