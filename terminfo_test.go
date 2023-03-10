// Copyright © 2023 Timothy E. Peoples

package decor

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/xo/terminfo"
)

const (
	xt_sitm     = "\x1b[3m"
	xt_ritm     = "\x1b[23m"
	xt_defBG    = "\x1b[49m"
	xt_defFG    = "\x1b[39m"
	xt_setaf59  = "\x1b[38;5;59m"
	xt_setaf204 = "\x1b[38;5;204m"
	xt_setaf214 = "\x1b[38;5;214m"
	xt_sgr0     = "\x1b(B\x1b[m"
)

func decodeAttrString(in string) string {
	amap := map[string]string{
		"sitm":     "\x1b[3m",
		"ritm":     "\x1b[23m",
		"defBG":    "\x1b[49m",
		"defFG":    "\x1b[39m",
		"setaf59":  "\x1b[38;5;59m",
		"setaf204": "\x1b[38;5;204m",
		"setaf214": "\x1b[38;5;214m",
		"sgr0":     "\x1b(B\x1b[m",
	}

	var out string

	for in != "" {
		x := strings.IndexByte(in, 0x1b)
		if x == -1 {
			out += in
			break
		}

		out += in[:x]
		in = in[x:]

		for k, v := range amap {
			nin := strings.TrimPrefix(in, v)
			if in == nin {
				continue
			}

			out += fmt.Sprintf("<%s>", k)
			in = nin
			break
		}
	}

	return out
}

// xterm256Decorator returns a *Decorator of a known terminal type that
// is suitable for testing.
func xterm256Decorator() (*Decorator, error) {
	data, err := base64.StdEncoding.DecodeString(xterm256color)
	if err != nil {
		return nil, err
	}

	ti, err := terminfo.Decode(data)
	if err != nil {
		return nil, err
	}

	return newDecorator("xterm-256color", ti), nil
}

// So unit tests may execute against a known terminal definition, we
// include here the base64 encoding of the compiled terminfo entry for
// type "xterm-256color" as defined in "/lib/terminfo/x/xterm-256color"
// on Ubuntu 18.04.6 LTS.
const xterm256color = `
GgElACYADwCdAQIGeHRlcm0tMjU2Y29sb3J8eHRlcm0gd2l0aCAyNTYgY29sb3JzAAABAAABAAAA
AQAAAAABAQAAAAAAAAABAAABAAEBAAAAAAAAAAABAFAACAAYAP//////////////////////////
AAH/fwAABAAGAAgAGQAeACYAKgAuAP//OQBKAEwAUABXAP//WQBmAP//agBuAHgAfAD/////gACE
AIkAjgD//6AApQCqAP//rwC0ALkAvgDHAMsA0gD//+QA6QDvAPUA////////BwH///////8ZAf//
HQH///////8fAf//JAH//////////ygBLAEyATYBOgE+AUQBSgFQAVYBXAFgAf//ZQH//2kBbgFz
AXcBfgH//4UBiQGRAf////////////////////////////+ZAaIB/////6sBtAG9AcYBzwHYAeEB
6gHzAfwB////////BQIJAg4CEwInAjAC/////0ICRQJQAlMCVQJYArUC//+4Av//////////////
/7oC//////////++Av//8wL/////9wL9Av////////////////////////////8DAwcD////////
//////////////////////////////////////////////////////////8LA/////8SA///////
////GQMgAycD/////y4D//81A////////zwD/////////////0MDSQNPA1YDXQNkA2sDcwN7A4MD
iwOTA5sDowOrA7IDuQPAA8cDzwPXA98D5wPvA/cD/wMHBA4EFQQcBCMEKwQzBDsEQwRLBFMEWwRj
BGoEcQR4BH8EhwSPBJcEnwSnBK8EtwS/BMYEzQTUBP//////////////////////////////////
///////////////////////////ZBOQE6QT8BAAFCQUQBf////////////////////////////9u
Bf///////////////////////3MF////////////////////////////////////////////////
////////////////////////////////////////eQX///////99BbwF////////////////////
////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////wF/wUbW1oABwANABtbJWklcDElZDslcDIlZHIA
G1szZwAbW0gbWzJKABtbSwAbW0oAG1slaSVwMSVkRwAbWyVpJXAxJWQ7JXAyJWRIAAoAG1tIABtb
PzI1bAAIABtbPzEybBtbPzI1aAAbW0MAG1tBABtbPzEyOzI1aAAbW1AAG1tNABsoMAAbWzVtABtb
MW0AG1s/MTA0OWgbWzIyOzA7MHQAG1sybQAbWzRoABtbOG0AG1s3bQAbWzdtABtbNG0AG1slcDEl
ZFgAGyhCABsoQhtbbQAbWz8xMDQ5bBtbMjM7MDswdAAbWzRsABtbMjdtABtbMjRtABtbPzVoJDwx
MDAvPhtbPzVsABtbIXAbWz8zOzRsG1s0bBs+ABtbTAB/ABtbM34AG09CABtPUAAbWzIxfgAbT1EA
G09SABtPUwAbWzE1fgAbWzE3fgAbWzE4fgAbWzE5fgAbWzIwfgAbT0gAG1syfgAbT0QAG1s2fgAb
WzV+ABtPQwAbWzE7MkIAG1sxOzJBABtPQQAbWz8xbBs+ABtbPzFoGz0AG1s/MTAzNGwAG1s/MTAz
NGgAG1slcDElZFAAG1slcDElZE0AG1slcDElZEIAG1slcDElZEAAG1slcDElZFMAG1slcDElZEwA
G1slcDElZEQAG1slcDElZEMAG1slcDElZFQAG1slcDElZEEAG1tpABtbNGkAG1s1aQAlcDElYxtb
JXAyJXsxfSUtJWRiABtjG10xMDQHABtbIXAbWz8zOzRsG1s0bBs+ABs4ABtbJWklcDElZGQAGzcA
CgAbTQAlPyVwOSV0GygwJWUbKEIlOxtbMCU/JXA2JXQ7MSU7JT8lcDUldDsyJTslPyVwMiV0OzQl
OyU/JXAxJXAzJXwldDs3JTslPyVwNCV0OzUlOyU/JXA3JXQ7OCU7bQAbSAAJABtPRQBgYGFhZmZn
Z2lpampra2xsbW1ubm9vcHBxcXJyc3N0dHV1dnZ3d3h4eXl6ent7fHx9fX5+ABtbWgAbWz83aAAb
Wz83bAAbT0YAG09NABtbMzsyfgAbWzE7MkYAG1sxOzJIABtbMjsyfgAbWzE7MkQAG1s2OzJ+ABtb
NTsyfgAbWzE7MkMAG1syM34AG1syNH4AG1sxOzJQABtbMTsyUQAbWzE7MlIAG1sxOzJTABtbMTU7
Mn4AG1sxNzsyfgAbWzE4OzJ+ABtbMTk7Mn4AG1syMDsyfgAbWzIxOzJ+ABtbMjM7Mn4AG1syNDsy
fgAbWzE7NVAAG1sxOzVRABtbMTs1UgAbWzE7NVMAG1sxNTs1fgAbWzE3OzV+ABtbMTg7NX4AG1sx
OTs1fgAbWzIwOzV+ABtbMjE7NX4AG1syMzs1fgAbWzI0OzV+ABtbMTs2UAAbWzE7NlEAG1sxOzZS
ABtbMTs2UwAbWzE1OzZ+ABtbMTc7Nn4AG1sxODs2fgAbWzE5OzZ+ABtbMjA7Nn4AG1syMTs2fgAb
WzIzOzZ+ABtbMjQ7Nn4AG1sxOzNQABtbMTszUQAbWzE7M1IAG1sxOzNTABtbMTU7M34AG1sxNzsz
fgAbWzE4OzN+ABtbMTk7M34AG1syMDszfgAbWzIxOzN+ABtbMjM7M34AG1syNDszfgAbWzE7NFAA
G1sxOzRRABtbMTs0UgAbWzFLABtbJWklZDslZFIAG1s2bgAbWz8lWzswMTIzNDU2Nzg5XWMAG1tj
ABtbMzk7NDltABtdMTA0BwAbXTQ7JXAxJWQ7cmdiOiVwMiV7MjU1fSUqJXsxMDAwfSUvJTIuMlgv
JXAzJXsyNTV9JSolezEwMDB9JS8lMi4yWC8lcDQlezI1NX0lKiV7MTAwMH0lLyUyLjJYG1wAG1sz
bQAbWzIzbQAbW00AG1slPyVwMSV7OH0lPCV0MyVwMSVkJWUlcDElezE2fSU8JXQ5JXAxJXs4fSUt
JWQlZTM4OzU7JXAxJWQlO20AG1slPyVwMSV7OH0lPCV0NCVwMSVkJWUlcDElezE2fSU8JXQxMCVw
MSV7OH0lLSVkJWU0ODs1OyVwMSVkJTttABtsABttAAIAAABAAIIAAwMBAQAABwATABgAKgAwADoA
QQBIAE8AVgBdAGQAawByAHkAgACHAI4AlQCcAKMAqgCxALgAvwDGAM0A1ADbAOIA6QDwAPcA/gAF
AQwBEwEaASEBKAEvATYBPQFEAUsBUgFZAWABZwFuAXUBfAGDAYoBkQGYAZ8B//////////+mAawB
AAADAAYACQAMAA8AEgAVABgAHQAiACcALAAxADUAOgA/AEQASQBOAFQAWgBgAGYAbAByAHgAfgCE
AIoAjwCUAJkAngCjAKkArwC1ALsAwQDHAM0A0wDZAN8A5QDrAPEA9wD9AAMBCQEPARUBGwEfASQB
KQEuATMBOAE8AUABRAFIAU0BG10xMTIHABtdMTI7JXAxJXMHABtbM0oAG101MjslcDElczslcDIl
cwcAG1syIHEAG1slcDElZCBxABtbMzszfgAbWzM7NH4AG1szOzV+ABtbMzs2fgAbWzM7N34AG1sx
OzJCABtbMTszQgAbWzE7NEIAG1sxOzVCABtbMTs2QgAbWzE7N0IAG1sxOzNGABtbMTs0RgAbWzE7
NUYAG1sxOzZGABtbMTs3RgAbWzE7M0gAG1sxOzRIABtbMTs1SAAbWzE7NkgAG1sxOzdIABtbMjsz
fgAbWzI7NH4AG1syOzV+ABtbMjs2fgAbWzI7N34AG1sxOzNEABtbMTs0RAAbWzE7NUQAG1sxOzZE
ABtbMTs3RAAbWzY7M34AG1s2OzR+ABtbNjs1fgAbWzY7Nn4AG1s2Ozd+ABtbNTszfgAbWzU7NH4A
G1s1OzV+ABtbNTs2fgAbWzU7N34AG1sxOzNDABtbMTs0QwAbWzE7NUMAG1sxOzZDABtbMTs3QwAb
WzE7MkEAG1sxOzNBABtbMTs0QQAbWzE7NUEAG1sxOzZBABtbMTs3QQAbWzI5bQAbWzltAEFYAFhU
AENyAENzAEUzAE1zAFNlAFNzAGtEQzMAa0RDNABrREM1AGtEQzYAa0RDNwBrRE4Aa0ROMwBrRE40
AGtETjUAa0RONgBrRE43AGtFTkQzAGtFTkQ0AGtFTkQ1AGtFTkQ2AGtFTkQ3AGtIT00zAGtIT000
AGtIT001AGtIT002AGtIT003AGtJQzMAa0lDNABrSUM1AGtJQzYAa0lDNwBrTEZUMwBrTEZUNABr
TEZUNQBrTEZUNgBrTEZUNwBrTlhUMwBrTlhUNABrTlhUNQBrTlhUNgBrTlhUNwBrUFJWMwBrUFJW
NABrUFJWNQBrUFJWNgBrUFJWNwBrUklUMwBrUklUNABrUklUNQBrUklUNgBrUklUNwBrVVAAa1VQ
MwBrVVA0AGtVUDUAa1VQNgBrVVA3AGthMgBrYjEAa2IzAGtjMgBybXh4AHNteHgA
`
