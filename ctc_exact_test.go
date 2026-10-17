package ctc

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestExactSequences(t *testing.T) {
	allStyles := Bold | Dim | Italic | Underline | Blink | Negative | Hidden | Strike
	cases := []struct {
		expr string
		c    Color
		str  string
		name string
		info string
	}{
		{"ForegroundRed", ForegroundRed,
			"\x1b[0;31m", "ForegroundRed", "ForegroundRed"},
		{"ForegroundBright|ForegroundRed", ForegroundBright | ForegroundRed,
			"\x1b[0;91m", "ForegroundBrightRed", "ForegroundBright | ForegroundRed"},
		{"BackgroundBright|BackgroundWhite", BackgroundBright | BackgroundWhite,
			"\x1b[0;107m", "BackgroundBrightWhite", "BackgroundBright | BackgroundWhite"},
		{"ForegroundColor256(196)", ForegroundColor256(196),
			"\x1b[0;38;5;196m", "ForegroundColor256(196)", "ForegroundColor256(196)"},
		{"BackgroundColor256(16)", BackgroundColor256(16),
			"\x1b[0;48;5;16m", "BackgroundColor256(16)", "BackgroundColor256(16)"},
		{"ForegroundTrueColor(1,2,3)", ForegroundTrueColor(1, 2, 3),
			"\x1b[0;38;2;1;2;3m", "ForegroundTrueColor(1,2,3)", "ForegroundTrueColor(1,2,3)"},
		{"BackgroundTrueColor(255,128,0)", BackgroundTrueColor(255, 128, 0),
			"\x1b[0;48;2;255;128;0m", "BackgroundTrueColor(255,128,0)", "BackgroundTrueColor(255,128,0)"},
		{"BackgroundRed|ForegroundColor256(200)", BackgroundRed | ForegroundColor256(200),
			"\x1b[0;38;5;200;41m", "ForegroundColor256(200)BackgroundRed", "ForegroundColor256(200) | BackgroundRed"},
		{"ForegroundRed|BackgroundColor256(200)", ForegroundRed | BackgroundColor256(200),
			"\x1b[0;31;48;5;200m", "ForegroundRedBackgroundColor256(200)", "ForegroundRed | BackgroundColor256(200)"},
		{"ForegroundTrueColor(1,2,3)|BackgroundColor256(7)|Bold", ForegroundTrueColor(1, 2, 3) | BackgroundColor256(7) | Bold,
			"\x1b[0;38;2;1;2;3;48;5;7;1m", "ForegroundTrueColor(1,2,3)BackgroundColor256(7)Bold", "ForegroundTrueColor(1,2,3) | BackgroundColor256(7) | Bold"},
		{"BackgroundTrueColor(4,5,6)|ForegroundBright|ForegroundBlue", BackgroundTrueColor(4, 5, 6) | ForegroundBright | ForegroundBlue,
			"\x1b[0;94;48;2;4;5;6m", "ForegroundBrightBlueBackgroundTrueColor(4,5,6)", "ForegroundBright | ForegroundBlue | BackgroundTrueColor(4,5,6)"},
		{"Bold|Italic|Underline", Bold | Italic | Underline,
			"\x1b[0;1;3;4m", "BoldItalicUnderline", "Bold | Italic | Underline"},
		{"Negative", Negative,
			"\x1b[0;7m", "Negative", "Negative"},
		{"ForegroundTrueColor(255,255,255)|BackgroundTrueColor(255,255,255)|allStyles", ForegroundTrueColor(255, 255, 255) | BackgroundTrueColor(255, 255, 255) | allStyles,
			"\x1b[0;38;2;255;255;255;48;2;255;255;255;1;2;3;4;5;7;8;9m",
			"ForegroundTrueColor(255,255,255)BackgroundTrueColor(255,255,255)BoldDimItalicUnderlineBlinkNegativeHiddenStrike",
			"ForegroundTrueColor(255,255,255) | BackgroundTrueColor(255,255,255) | Bold | Dim | Italic | Underline | Blink | Negative | Hidden | Strike"},
	}
	for _, tc := range cases {
		if got := tc.c.String(); got != tc.str {
			t.Errorf("%s String(): got %q want %q", tc.expr, got, tc.str)
		}
		if got := string(tc.c.Bytes()); got != tc.c.String() {
			t.Errorf("%s Bytes(): got %q want %q", tc.expr, got, tc.c.String())
		}
		if got := tc.c.Name(); got != tc.name {
			t.Errorf("%s Name(): got %q want %q", tc.expr, got, tc.name)
		}
		if got := tc.c.Info(); got != tc.info {
			t.Errorf("%s Info(): got %q want %q", tc.expr, got, tc.info)
		}
	}
	longest := cases[len(cases)-1].c
	if got := len(longest.Bytes()); got != 54 {
		t.Errorf("len(Bytes()) of longest sequence: got %d want %d", got, 54)
	}
}

func TestTrueColorRoundTrip(t *testing.T) {
	fg := ForegroundTrueColor(10, 20, 30)
	if ok, r, g, b := fg.ForegroundTrueColor(); !ok || r != 10 || g != 20 || b != 30 {
		t.Errorf("ForegroundTrueColor(10,20,30).ForegroundTrueColor(): got (%v,%d,%d,%d) want (true,10,20,30)", ok, r, g, b)
	}
	if ok, _, _, _ := fg.BackgroundTrueColor(); ok {
		t.Errorf("ForegroundTrueColor(10,20,30).BackgroundTrueColor(): got ok=true want false")
	}
	if ok, _ := fg.ForegroundColor256(); ok {
		t.Errorf("ForegroundTrueColor(10,20,30).ForegroundColor256(): got ok=true want false")
	}
	if ok, _ := fg.ForegroundStandard(); ok {
		t.Errorf("ForegroundTrueColor(10,20,30).ForegroundStandard(): got ok=true want false")
	}

	bg := BackgroundTrueColor(200, 100, 50)
	if ok, r, g, b := bg.BackgroundTrueColor(); !ok || r != 200 || g != 100 || b != 50 {
		t.Errorf("BackgroundTrueColor(200,100,50).BackgroundTrueColor(): got (%v,%d,%d,%d) want (true,200,100,50)", ok, r, g, b)
	}
	if ok, _, _, _ := bg.ForegroundTrueColor(); ok {
		t.Errorf("BackgroundTrueColor(200,100,50).ForegroundTrueColor(): got ok=true want false")
	}
	if ok, _ := bg.BackgroundColor256(); ok {
		t.Errorf("BackgroundTrueColor(200,100,50).BackgroundColor256(): got ok=true want false")
	}
	if ok, _ := bg.BackgroundStandard(); ok {
		t.Errorf("BackgroundTrueColor(200,100,50).BackgroundStandard(): got ok=true want false")
	}

	mixed := ForegroundColor256(123) | BackgroundTrueColor(9, 8, 7)
	if ok, v := mixed.ForegroundColor256(); !ok || v != 123 {
		t.Errorf("mixed ForegroundColor256(): got (%v,%d) want (true,123)", ok, v)
	}
	if ok, r, g, b := mixed.BackgroundTrueColor(); !ok || r != 9 || g != 8 || b != 7 {
		t.Errorf("mixed BackgroundTrueColor(): got (%v,%d,%d,%d) want (true,9,8,7)", ok, r, g, b)
	}
}

func TestCubeLevelBoundaries(t *testing.T) {
	levels := []struct{ v, want uint8 }{
		{0, 0}, {47, 0}, {48, 1}, {95, 1}, {114, 1}, {115, 2}, {135, 2}, {154, 2},
		{155, 3}, {175, 3}, {194, 3}, {195, 4}, {215, 4}, {234, 4}, {235, 5}, {255, 5},
	}
	for _, tc := range levels {
		if got := rgbToCubeLevel(tc.v); got != tc.want {
			t.Errorf("rgbToCubeLevel(%d): got %d want %d", tc.v, got, tc.want)
		}
	}

	toCode := []struct{ r, g, b, want uint8 }{
		{0, 0, 0, 16},
		{255, 0, 0, 196},
		{0, 255, 0, 46},
		{0, 0, 255, 21},
		{255, 255, 255, 231},
		{115, 155, 195, 110},
	}
	for _, tc := range toCode {
		if got := RGBToColor256(tc.r, tc.g, tc.b); got != tc.want {
			t.Errorf("RGBToColor256(%d,%d,%d): got %d want %d", tc.r, tc.g, tc.b, got, tc.want)
		}
	}

	toRGB := []struct{ n, r, g, b uint8 }{
		{16, 0, 0, 0},
		{196, 255, 0, 0},
		{244, 128, 128, 128},
		{231, 255, 255, 255},
		{9, 255, 0, 0},
		{7, 192, 192, 192},
	}
	for _, tc := range toRGB {
		if r, g, b := Color256ToRGB(tc.n); r != tc.r || g != tc.g || b != tc.b {
			t.Errorf("Color256ToRGB(%d): got (%d,%d,%d) want (%d,%d,%d)", tc.n, r, g, b, tc.r, tc.g, tc.b)
		}
	}
}

func standardLists() (fgList, bgList [16]Color) {
	fg := []Color{ForegroundBlack, ForegroundRed, ForegroundGreen, ForegroundYellow,
		ForegroundBlue, ForegroundMagenta, ForegroundCyan, ForegroundWhite}
	bg := []Color{BackgroundBlack, BackgroundRed, BackgroundGreen, BackgroundYellow,
		BackgroundBlue, BackgroundMagenta, BackgroundCyan, BackgroundWhite}
	for k := 0; k < 8; k++ {
		fgList[k] = fg[k]
		fgList[8+k] = ForegroundBright | fg[k]
		bgList[k] = bg[k]
		bgList[8+k] = BackgroundBright | bg[k]
	}
	return fgList, bgList
}

func TestForEachOrder(t *testing.T) {
	var got []Color
	ForEach(func(c Color) { got = append(got, c) })

	fgList, bgList := standardLists()
	var want []Color
	want = append(want, fgList[:]...)
	want = append(want, bgList[:]...)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			want = append(want, fgList[j]|bgList[i])
		}
	}

	if len(got) != len(want) {
		t.Fatalf("ForEach count: got %d want %d", len(got), len(want))
	}
	for k := range want {
		if got[k] != want[k] {
			t.Fatalf("ForEach order at index %d: got %q want %q", k, got[k].Name(), want[k].Name())
		}
	}
}

func TestExtraGenerated(t *testing.T) {
	data, err := ioutil.ReadFile("ctc_extra.go")
	if err != nil {
		t.Fatal(err)
	}
	type pair struct{ name, info string }
	var got []pair
	for _, line := range strings.Split(string(data), "\n") {
		s := strings.TrimSpace(line)
		s = strings.TrimPrefix(s, "// ")
		if i := strings.Index(s, " = "); i >= 0 {
			got = append(got, pair{s[:i], s[i+len(" = "):]})
		}
	}

	var want []pair
	ForEach(func(c Color) { want = append(want, pair{c.Name(), c.Info()}) })

	if len(got) != len(want) {
		t.Fatalf("ctc_extra.go entries: got %d want %d", len(got), len(want))
	}
	for k := range want {
		if got[k] != want[k] {
			t.Fatalf("ctc_extra.go entry %d: got %q = %q want %q = %q", k, got[k].name, got[k].info, want[k].name, want[k].info)
		}
	}
}
