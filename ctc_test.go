package ctc

import (
	"testing"
)

var repetitionString = map[string]bool{}
var repetitionName = map[string]bool{}
var repetitionInfo = map[string]bool{}

func testCheck(t *testing.T, beg, end, step Color) {
	for c := beg; c <= end; c += step {
		{
			s := c.String()
			if _, ok := repetitionString[s]; ok {
				t.Errorf("error %s", c.Name())
			}
			repetitionString[s] = true
		}

		{
			s := c.Name()
			if _, ok := repetitionName[s]; ok {
				t.Errorf("error %s", c.Name())
			}
			repetitionName[s] = true
		}

		{
			s := c.Info()
			if _, ok := repetitionInfo[s]; ok {
				t.Errorf("error %s", c.Info())
			}
			repetitionInfo[s] = true
		}
	}
}

func TestCheck(t *testing.T) {
	data := [...][3]Color{
		{
			ForegroundBlack,
			ForegroundWhite | ForegroundBright,
			1,
		},
		{
			BackgroundBlack,
			BackgroundWhite | BackgroundBright,
			1 << 4,
		},
		{
			ForegroundBlack | BackgroundBlack,
			ForegroundWhite | ForegroundBright | BackgroundWhite | BackgroundBright,
			1,
		},
	}
	for _, v := range data {
		testCheck(t, v[0], v[1], v[2])
	}
}
