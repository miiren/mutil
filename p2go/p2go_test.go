package p2go

import (
	"fmt"
	"testing"
	"time"
)

func TestDate(t *testing.T) {
	fmt.Println(Date("20060102150405", time.Now().Unix()))
}
