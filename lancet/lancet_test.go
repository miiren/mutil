package lancet

import (
	"fmt"
	"testing"
	"time"
)

func TestToFormat(t *testing.T) {
	res := FormatTimeToStr(time.Now(), "yyyy-mm-dd hh:mm:ss")
	fmt.Println("res:", res)
}
