package signing

import (
	"testing"
	"time"
)

const runingMax = 10

func TestPressure10s(t *testing.T) {
	cnt := 0
	runing := 0
	t.Log("Pressure test start!\n")
	for i := 0; i < runingMax; i++ {
		go signForTest(&runing, &cnt, t)
		runing++
	}
	startTime := time.Now()
	for {
		if runing < runingMax {
			for i := 0; i < runingMax-runing; i++ {
				go signForTest(&runing, &cnt, t)
				runing++
			}
		}

		if time.Since(startTime) >= time.Second*10 {
			t.Log("\n\n[---------signed ", cnt, " messages in 10s----------]\n\n")
			return
		}
	}
}

func signForTest(runing *int, cnt *int, t *testing.T) {
	TestE2EConcurrent(t)
	*runing--
	*cnt++
}
