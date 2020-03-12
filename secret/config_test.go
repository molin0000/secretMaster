package secret

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	tm := time.Now()
	fmt.Println(tm.Hour(), tm.Minute())
}
