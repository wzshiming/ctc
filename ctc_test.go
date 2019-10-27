package ctc

import (
	"testing"
)

func TestCheck(t *testing.T) {
	var (
		repetitionString = map[string]bool{}
		repetitionName   = map[string]bool{}
		repetitionInfo   = map[string]bool{}
	)
	ForEach(func(c Color) {
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
				t.Errorf("error %s", c.Name())
			}
			repetitionInfo[s] = true
		}
	})
}
