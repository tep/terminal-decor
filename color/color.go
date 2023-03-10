// Copyright © 2023 Timothy E. Peoples

// Package color provides functions for maping color names to numbers and vice
// versa.
//
// Color names are derived from the xterm color table with slight alterations
// (a suffix of 'a' or 'b') for names assigned to multiple color numbers.
//
// For all names having multiple color numbers, one will have no suffix at all
// (so the original name still resolves to a valid color number) but subsequent
// values will have a suffix of 'a' or 'b').
//
// Color numbers and their associated names are as specified below with the
// following distinctions:
//
//	Names assigned to 2 distinct color numbers (with one having a suffix of
//	'a') are marked with a dagger (†) while those having 3 numbers are marked
//	with a double dagger (‡).
//
//	#0    BLACK
//	#1    RED
//	#2    GREEN
//	#3    YELLOW
//	#4    BLUE
//	#5    MAGENTA
//	#6    CYAN
//	#7    WHITE
//	#8    BOLD_BLACK
//	#9    BOLD_RED
//	#10   BOLD_GREEN
//	#11   BOLD_YELLOW
//	#12   BOLD_BLUE
//	#13   BOLD_MAGENTA
//	#14   BOLD_CYAN
//	#15   BOLD_WHITE
//	#16   Grey0
//	#17   NavyBlue
//	#18   DarkBlue
//	#19   Blue3a †
//	#20   Blue3 †
//	#21   Blue1
//	#22   DarkGreen
//	#23   DeepSkyBlue4a ‡
//	#24   DeepSkyBlue4b ‡
//	#25   DeepSkyBlue4 ‡
//	#26   DodgerBlue3
//	#27   DodgerBlue2
//	#28   Green4
//	#29   SpringGreen4
//	#30   Turquoise4
//	#31   DeepSkyBlue3a †
//	#32   DeepSkyBlue3 †
//	#33   DodgerBlue1
//	#34   Green3a †
//	#35   SpringGreen3a †
//	#36   DarkCyan
//	#37   LightSeaGreen
//	#38   DeepSkyBlue2
//	#39   DeepSkyBlue1
//	#40   Green3 †
//	#41   SpringGreen3 †
//	#42   SpringGreen2a †
//	#43   Cyan3
//	#44   DarkTurquoise
//	#45   Turquoise2
//	#46   Green1
//	#47   SpringGreen2 †
//	#48   SpringGreen1
//	#49   MediumSpringGreen
//	#50   Cyan2
//	#51   Cyan1
//	#52   DarkRed1
//	#53   DeepPink4 ‡
//	#54   Purple4a †
//	#55   Purple4 †
//	#56   Purple3
//	#57   BlueViolet
//	#58   Orange4 †
//	#59   Grey37
//	#60   MediumPurple4
//	#61   SlateBlue3a †
//	#62   SlateBlue3 †
//	#63   RoyalBlue1
//	#64   Chartreuse4
//	#65   DarkSeaGreen4a †
//	#66   PaleTurquoise4
//	#67   SteelBlue
//	#68   SteelBlue3
//	#69   CornflowerBlue
//	#70   Chartreuse3a †
//	#71   DarkSeaGreen4 †
//	#72   CadetBlue2
//	#73   CadetBlue
//	#74   SkyBlue3
//	#75   SteelBlue1a †
//	#76   Chartreuse3 †
//	#77   PaleGreen3 †
//	#78   SeaGreen3
//	#79   Aquamarine3
//	#80   MediumTurquoise
//	#81   SteelBlue1 †
//	#82   Chartreuse2 †
//	#83   SeaGreen2
//	#84   SeaGreen1a †
//	#85   SeaGreen1 †
//	#86   Aquamarine1 †
//	#87   DarkSlateGray2
//	#88   DarkRed2
//	#89   DeepPink4a ‡
//	#90   DarkMagenta2
//	#91   DarkMagenta
//	#92   DarkViolet2
//	#93   Purple2
//	#94   Orange4a †
//	#95   LightPink4
//	#96   Plum4
//	#97   MediumPurple3a †
//	#98   MediumPurple3 †
//	#99   SlateBlue1
//	#100  Yellow4a †
//	#101  Wheat4
//	#102  Grey53
//	#103  LightSlateGrey
//	#104  MediumPurple
//	#105  LightSlateBlue
//	#106  Yellow4 †
//	#107  DarkOliveGreen3a ‡
//	#108  DarkSeaGreen
//	#109  LightSkyBlue3a †
//	#110  LightSkyBlue3 †
//	#111  SkyBlue2
//	#112  Chartreuse2a †
//	#113  DarkOliveGreen3b ‡
//	#114  PaleGreen3a †
//	#115  DarkSeaGreen3a †
//	#116  DarkSlateGray3
//	#117  SkyBlue1
//	#118  Chartreuse1
//	#119  LightGreen2
//	#120  LightGreen
//	#121  PaleGreen1a †
//	#122  Aquamarine1a †
//	#123  DarkSlateGray1
//	#124  Red3a †
//	#125  DeepPink4b ‡
//	#126  MediumVioletRed
//	#127  Magenta3a ‡
//	#128  DarkViolet
//	#129  Purple
//	#130  DarkOrange3a †
//	#131  IndianRed2
//	#132  HotPink3a †
//	#133  MediumOrchid3
//	#134  MediumOrchid
//	#135  MediumPurple2a †
//	#136  DarkGoldenrod
//	#137  LightSalmon3 †
//	#138  RosyBrown
//	#139  Grey63
//	#140  MediumPurple2 †
//	#141  MediumPurple1
//	#142  Gold3 †
//	#143  DarkKhaki
//	#144  NavajoWhite3
//	#145  Grey69
//	#146  LightSteelBlue3
//	#147  LightSteelBlue
//	#148  Yellow3 †
//	#149  DarkOliveGreen3 ‡
//	#150  DarkSeaGreen3 †
//	#151  DarkSeaGreen2a †
//	#152  LightCyan3
//	#153  LightSkyBlue1
//	#154  GreenYellow
//	#155  DarkOliveGreen2
//	#156  PaleGreen1 †
//	#157  DarkSeaGreen2 †
//	#158  DarkSeaGreen1 †
//	#159  PaleTurquoise1
//	#160  Red3 †
//	#161  DeepPink3a †
//	#162  DeepPink3 †
//	#163  Magenta3b ‡
//	#164  Magenta3 ‡
//	#165  Magenta2 †
//	#166  DarkOrange3 †
//	#167  IndianRed
//	#168  HotPink3 †
//	#169  HotPink2
//	#170  Orchid
//	#171  MediumOrchid1a †
//	#172  Orange3
//	#173  LightSalmon3a †
//	#174  LightPink3
//	#175  Pink3
//	#176  Plum3
//	#177  Violet
//	#178  Gold3a †
//	#179  LightGoldenrod3
//	#180  Tan
//	#181  MistyRose3
//	#182  Thistle3
//	#183  Plum2
//	#184  Yellow3a †
//	#185  Khaki3
//	#186  LightGoldenrod2a ‡
//	#187  LightYellow3
//	#188  Grey84
//	#189  LightSteelBlue1
//	#190  Yellow2
//	#191  DarkOliveGreen1a †
//	#192  DarkOliveGreen1 †
//	#193  DarkSeaGreen1a †
//	#194  Honeydew2
//	#195  LightCyan1
//	#196  Red1
//	#197  DeepPink2
//	#198  DeepPink1a †
//	#199  DeepPink1 †
//	#200  Magenta2a †
//	#201  Magenta1
//	#202  OrangeRed1
//	#203  IndianRed1a †
//	#204  IndianRed1 †
//	#205  HotPink4
//	#206  HotPink
//	#207  MediumOrchid1 †
//	#208  DarkOrange
//	#209  Salmon1
//	#210  LightCoral
//	#211  PaleVioletRed1
//	#212  Orchid2
//	#213  Orchid1
//	#214  Orange1
//	#215  SandyBrown
//	#216  LightSalmon1
//	#217  LightPink1
//	#218  Pink1
//	#219  Plum1
//	#220  Gold1
//	#221  LightGoldenrod2b ‡
//	#222  LightGoldenrod2 ‡
//	#223  NavajoWhite1
//	#224  MistyRose1
//	#225  Thistle1
//	#226  Yellow1
//	#227  LightGoldenrod1
//	#228  Khaki1
//	#229  Wheat1
//	#230  Cornsilk1
//	#231  Grey100
//	#232  Grey3
//	#233  Grey7
//	#234  Grey11
//	#235  Grey15
//	#236  Grey19
//	#237  Grey23
//	#238  Grey27
//	#239  Grey30
//	#240  Grey35
//	#241  Grey39
//	#242  Grey42
//	#243  Grey46
//	#244  Grey50
//	#245  Grey54
//	#246  Grey58
//	#247  Grey62
//	#248  Grey66
//	#249  Grey70
//	#250  Grey74
//	#251  Grey78
//	#252  Grey82
//	#253  Grey85
//	#254  Grey89
//	#255  Grey93
package color

import (
	"strings"

	"toolman.org/terminal/decor/internal/colors"
)

// Name returns the color name associated with the given number.
func Name(num uint8) string {
	return colors.Names[num]
}

// Number returns the color number assigned to the given name. If name does not
// match a known color number, -1 is returned. For this lookup, names are
// case-insensitive.
//
// See the package documentation for the full list of known color names.
func Number(name string) int {
	if n, ok := colors.Numbers[strings.ToLower(name)]; ok {
		return n
	}

	return -1
}
