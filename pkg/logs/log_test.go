package logs

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestLog(t *testing.T) {
	Init(os.Stdout, os.Stderr, os.Stderr, LogField("tid"))

	ctx := context.WithValue(context.Background(), LogField("tid"), time.Now().Unix())
	Ctx(ctx).Info("kkkddd")
}
