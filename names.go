// Copyright © 2023 Timothy E. Peoples

package termcolor

import "strings"

func Name(num int) string {
	return names[num]
}

func Number(name string) int {
	return numbers[strings.ToLower(name)]
}

func NumberOK(name string) (int, bool) {
	num, ok := numbers[strings.ToLower(name)]
	return num, ok
}

var names = []string{
	"BLACK",
	"RED",
	"GREEN",
	"YELLOW",
	"BLUE",
	"MAGENTA",
	"CYAN",
	"WHITE",
	"BOLD_BLACK",
	"BOLD_RED",
	"BOLD_GREEN",
	"BOLD_YELLOW",
	"BOLD_BLUE",
	"BOLD_MAGENTA",
	"BOLD_CYAN",
	"BOLD_WHITE",
	"Grey0",
	"NavyBlue",
	"DarkBlue",
	"Blue3a",
	"Blue3",
	"Blue1",
	"DarkGreen",
	"DeepSkyBlue4a",
	"DeepSkyBlue4b",
	"DeepSkyBlue4",
	"DodgerBlue3",
	"DodgerBlue2",
	"Green4",
	"SpringGreen4",
	"Turquoise4",
	"DeepSkyBlue3a",
	"DeepSkyBlue3",
	"DodgerBlue1",
	"Green3a",
	"SpringGreen3a",
	"DarkCyan",
	"LightSeaGreen",
	"DeepSkyBlue2",
	"DeepSkyBlue1",
	"Green3",
	"SpringGreen3",
	"SpringGreen2a",
	"Cyan3",
	"DarkTurquoise",
	"Turquoise2",
	"Green1",
	"SpringGreen2",
	"SpringGreen1",
	"MediumSpringGreen",
	"Cyan2",
	"Cyan1",
	"DarkRed1",
	"DeepPink4",
	"Purple4a",
	"Purple4",
	"Purple3",
	"BlueViolet",
	"Orange4",
	"Grey37",
	"MediumPurple4",
	"SlateBlue3a",
	"SlateBlue3",
	"RoyalBlue1",
	"Chartreuse4",
	"DarkSeaGreen4a",
	"PaleTurquoise4",
	"SteelBlue",
	"SteelBlue3",
	"CornflowerBlue",
	"Chartreuse3a",
	"DarkSeaGreen4",
	"CadetBlue2",
	"CadetBlue",
	"SkyBlue3",
	"SteelBlue1a",
	"Chartreuse3",
	"PaleGreen3",
	"SeaGreen3",
	"Aquamarine3",
	"MediumTurquoise",
	"SteelBlue1",
	"Chartreuse2",
	"SeaGreen2",
	"SeaGreen1a",
	"SeaGreen1",
	"Aquamarine1",
	"DarkSlateGray2",
	"DarkRed2",
	"DeepPink4a",
	"DarkMagenta2",
	"DarkMagenta",
	"DarkViolet2",
	"Purple2",
	"Orange4a",
	"LightPink4",
	"Plum4",
	"MediumPurple3a",
	"MediumPurple3",
	"SlateBlue1",
	"Yellow4a",
	"Wheat4",
	"Grey53",
	"LightSlateGrey",
	"MediumPurple",
	"LightSlateBlue",
	"Yellow4",
	"DarkOliveGreen3a",
	"DarkSeaGreen",
	"LightSkyBlue3a",
	"LightSkyBlue3",
	"SkyBlue2",
	"Chartreuse2a",
	"DarkOliveGreen3b",
	"PaleGreen3a",
	"DarkSeaGreen3a",
	"DarkSlateGray3",
	"SkyBlue1",
	"Chartreuse1",
	"LightGreen2",
	"LightGreen",
	"PaleGreen1a",
	"Aquamarine1a",
	"DarkSlateGray1",
	"Red3a",
	"DeepPink4b",
	"MediumVioletRed",
	"Magenta3a",
	"DarkViolet",
	"Purple",
	"DarkOrange3a",
	"IndianRed2",
	"HotPink3a",
	"MediumOrchid3",
	"MediumOrchid",
	"MediumPurple2a",
	"DarkGoldenrod",
	"LightSalmon3",
	"RosyBrown",
	"Grey63",
	"MediumPurple2",
	"MediumPurple1",
	"Gold3",
	"DarkKhaki",
	"NavajoWhite3",
	"Grey69",
	"LightSteelBlue3",
	"LightSteelBlue",
	"Yellow3",
	"DarkOliveGreen3",
	"DarkSeaGreen3",
	"DarkSeaGreen2a",
	"LightCyan3",
	"LightSkyBlue1",
	"GreenYellow",
	"DarkOliveGreen2",
	"PaleGreen1",
	"DarkSeaGreen2",
	"DarkSeaGreen1",
	"PaleTurquoise1",
	"Red3",
	"DeepPink3a",
	"DeepPink3",
	"Magenta3b",
	"Magenta3",
	"Magenta2",
	"DarkOrange3",
	"IndianRed",
	"HotPink3",
	"HotPink2",
	"Orchid",
	"MediumOrchid1a",
	"Orange3",
	"LightSalmon3a",
	"LightPink3",
	"Pink3",
	"Plum3",
	"Violet",
	"Gold3a",
	"LightGoldenrod3",
	"Tan",
	"MistyRose3",
	"Thistle3",
	"Plum2",
	"Yellow3a",
	"Khaki3",
	"LightGoldenrod2a",
	"LightYellow3",
	"Grey84",
	"LightSteelBlue1",
	"Yellow2",
	"DarkOliveGreen1a",
	"DarkOliveGreen1",
	"DarkSeaGreen1a",
	"Honeydew2",
	"LightCyan1",
	"Red1",
	"DeepPink2",
	"DeepPink1a",
	"DeepPink1",
	"Magenta2a",
	"Magenta1",
	"OrangeRed1",
	"IndianRed1a",
	"IndianRed1",
	"HotPink4",
	"HotPink",
	"MediumOrchid1",
	"DarkOrange",
	"Salmon1",
	"LightCoral",
	"PaleVioletRed1",
	"Orchid2",
	"Orchid1",
	"Orange1",
	"SandyBrown",
	"LightSalmon1",
	"LightPink1",
	"Pink1",
	"Plum1",
	"Gold1",
	"LightGoldenrod2b",
	"LightGoldenrod2",
	"NavajoWhite1",
	"MistyRose1",
	"Thistle1",
	"Yellow1",
	"LightGoldenrod1",
	"Khaki1",
	"Wheat1",
	"Cornsilk1",
	"Grey100",
	"Grey3",
	"Grey7",
	"Grey11",
	"Grey15",
	"Grey19",
	"Grey23",
	"Grey27",
	"Grey30",
	"Grey35",
	"Grey39",
	"Grey42",
	"Grey46",
	"Grey50",
	"Grey54",
	"Grey58",
	"Grey62",
	"Grey66",
	"Grey70",
	"Grey74",
	"Grey78",
	"Grey82",
	"Grey85",
	"Grey89",
	"Grey93",
}

var numbers = map[string]int{
	"black":             0,
	"red":               1,
	"green":             2,
	"yellow":            3,
	"blue":              4,
	"magenta":           5,
	"cyan":              6,
	"white":             7,
	"bold_black":        8,
	"bold_red":          9,
	"bold_green":        10,
	"bold_yellow":       11,
	"bold_blue":         12,
	"bold_magenta":      13,
	"bold_cyan":         14,
	"bold_white":        15,
	"grey0":             16,
	"navyblue":          17,
	"darkblue":          18,
	"blue3a":            19,
	"blue3":             20,
	"blue1":             21,
	"darkgreen":         22,
	"deepskyblue4a":     23,
	"deepskyblue4b":     24,
	"deepskyblue4":      25,
	"dodgerblue3":       26,
	"dodgerblue2":       27,
	"green4":            28,
	"springgreen4":      29,
	"turquoise4":        30,
	"deepskyblue3a":     31,
	"deepskyblue3":      32,
	"dodgerblue1":       33,
	"green3a":           34,
	"springgreen3a":     35,
	"darkcyan":          36,
	"lightseagreen":     37,
	"deepskyblue2":      38,
	"deepskyblue1":      39,
	"green3":            40,
	"springgreen3":      41,
	"springgreen2a":     42,
	"cyan3":             43,
	"darkturquoise":     44,
	"turquoise2":        45,
	"green1":            46,
	"springgreen2":      47,
	"springgreen1":      48,
	"mediumspringgreen": 49,
	"cyan2":             50,
	"cyan1":             51,
	"darkred1":          52,
	"deeppink4":         53,
	"purple4a":          54,
	"purple4":           55,
	"purple3":           56,
	"blueviolet":        57,
	"orange4":           58,
	"grey37":            59,
	"mediumpurple4":     60,
	"slateblue3a":       61,
	"slateblue3":        62,
	"royalblue1":        63,
	"chartreuse4":       64,
	"darkseagreen4a":    65,
	"paleturquoise4":    66,
	"steelblue":         67,
	"steelblue3":        68,
	"cornflowerblue":    69,
	"chartreuse3a":      70,
	"darkseagreen4":     71,
	"cadetblue2":        72,
	"cadetblue":         73,
	"skyblue3":          74,
	"steelblue1a":       75,
	"chartreuse3":       76,
	"palegreen3":        77,
	"seagreen3":         78,
	"aquamarine3":       79,
	"mediumturquoise":   80,
	"steelblue1":        81,
	"chartreuse2":       82,
	"seagreen2":         83,
	"seagreen1a":        84,
	"seagreen1":         85,
	"aquamarine1":       86,
	"darkslategray2":    87,
	"darkred2":          88,
	"deeppink4a":        89,
	"darkmagenta2":      90,
	"darkmagenta":       91,
	"darkviolet2":       92,
	"purple2":           93,
	"orange4a":          94,
	"lightpink4":        95,
	"plum4":             96,
	"mediumpurple3a":    97,
	"mediumpurple3":     98,
	"slateblue1":        99,
	"yellow4a":          100,
	"wheat4":            101,
	"grey53":            102,
	"lightslategrey":    103,
	"mediumpurple":      104,
	"lightslateblue":    105,
	"yellow4":           106,
	"darkolivegreen3a":  107,
	"darkseagreen":      108,
	"lightskyblue3a":    109,
	"lightskyblue3":     110,
	"skyblue2":          111,
	"chartreuse2a":      112,
	"darkolivegreen3b":  113,
	"palegreen3a":       114,
	"darkseagreen3a":    115,
	"darkslategray3":    116,
	"skyblue1":          117,
	"chartreuse1":       118,
	"lightgreen2":       119,
	"lightgreen":        120,
	"palegreen1a":       121,
	"aquamarine1a":      122,
	"darkslategray1":    123,
	"red3a":             124,
	"deeppink4b":        125,
	"mediumvioletred":   126,
	"magenta3a":         127,
	"darkviolet":        128,
	"purple":            129,
	"darkorange3a":      130,
	"indianred2":        131,
	"hotpink3a":         132,
	"mediumorchid3":     133,
	"mediumorchid":      134,
	"mediumpurple2a":    135,
	"darkgoldenrod":     136,
	"lightsalmon3":      137,
	"rosybrown":         138,
	"grey63":            139,
	"mediumpurple2":     140,
	"mediumpurple1":     141,
	"gold3":             142,
	"darkkhaki":         143,
	"navajowhite3":      144,
	"grey69":            145,
	"lightsteelblue3":   146,
	"lightsteelblue":    147,
	"yellow3":           148,
	"darkolivegreen3":   149,
	"darkseagreen3":     150,
	"darkseagreen2a":    151,
	"lightcyan3":        152,
	"lightskyblue1":     153,
	"greenyellow":       154,
	"darkolivegreen2":   155,
	"palegreen1":        156,
	"darkseagreen2":     157,
	"darkseagreen1":     158,
	"paleturquoise1":    159,
	"red3":              160,
	"deeppink3a":        161,
	"deeppink3":         162,
	"magenta3b":         163,
	"magenta3":          164,
	"magenta2":          165,
	"darkorange3":       166,
	"indianred":         167,
	"hotpink3":          168,
	"hotpink2":          169,
	"orchid":            170,
	"mediumorchid1a":    171,
	"orange3":           172,
	"lightsalmon3a":     173,
	"lightpink3":        174,
	"pink3":             175,
	"plum3":             176,
	"violet":            177,
	"gold3a":            178,
	"lightgoldenrod3":   179,
	"tan":               180,
	"mistyrose3":        181,
	"thistle3":          182,
	"plum2":             183,
	"yellow3a":          184,
	"khaki3":            185,
	"lightgoldenrod2a":  186,
	"lightyellow3":      187,
	"grey84":            188,
	"lightsteelblue1":   189,
	"yellow2":           190,
	"darkolivegreen1a":  191,
	"darkolivegreen1":   192,
	"darkseagreen1a":    193,
	"honeydew2":         194,
	"lightcyan1":        195,
	"red1":              196,
	"deeppink2":         197,
	"deeppink1a":        198,
	"deeppink1":         199,
	"magenta2a":         200,
	"magenta1":          201,
	"orangered1":        202,
	"indianred1a":       203,
	"indianred1":        204,
	"hotpink4":          205,
	"hotpink":           206,
	"mediumorchid1":     207,
	"darkorange":        208,
	"salmon1":           209,
	"lightcoral":        210,
	"palevioletred1":    211,
	"orchid2":           212,
	"orchid1":           213,
	"orange1":           214,
	"sandybrown":        215,
	"lightsalmon1":      216,
	"lightpink1":        217,
	"pink1":             218,
	"plum1":             219,
	"gold1":             220,
	"lightgoldenrod2b":  221,
	"lightgoldenrod2":   222,
	"navajowhite1":      223,
	"mistyrose1":        224,
	"thistle1":          225,
	"yellow1":           226,
	"lightgoldenrod1":   227,
	"khaki1":            228,
	"wheat1":            229,
	"cornsilk1":         230,
	"grey100":           231,
	"grey3":             232,
	"grey7":             233,
	"grey11":            234,
	"grey15":            235,
	"grey19":            236,
	"grey23":            237,
	"grey27":            238,
	"grey30":            239,
	"grey35":            240,
	"grey39":            241,
	"grey42":            242,
	"grey46":            243,
	"grey50":            244,
	"grey54":            245,
	"grey58":            246,
	"grey62":            247,
	"grey66":            248,
	"grey70":            249,
	"grey74":            250,
	"grey78":            251,
	"grey82":            252,
	"grey85":            253,
	"grey89":            254,
	"grey93":            255,
}
