// For https://github.com/suntong/ffcvt
// https://go.dev/play/p/bH3Wlsj0DYj

package main

import (
	"encoding/json"
	"fmt"
)

type Encoding struct {
	VES string // video encoding method set
	AES string // audio encoding method set
	SES string // subtitle encoding method set
	VEP string // video encoding method prepend
	AEP string // audio encoding method prepend
	SEP string // subtitle encoding method prepend
	VEA string // video encoding method append
	AEA string // audio encoding method append
	ABR string // audio bitrate
	CRF string // the CRF value: 0-51. Higher CRF gives lower quality
	Ext string // extension for the output file
}

func main() {

	// String contains JSON rows.
	text := `
{
  "copy": {
    "AES": "copy",
    "VES": "copy",
    "SES": "-c:s copy",
    "ABR": "64k",
    "CRF": "42",
    "Ext": "_.mkv"
  },
  "webm": {
    "AES": "libopus",
    "VES": "libvpx-vp9",
    "SES": "-c:s copy",
    "ABR": "64k",
    "CRF": "42",
    "Ext": "_.mkv"
  },
  "x265-opus": {
    "AES": "libopus",
    "VES": "libx265",
    "ABR": "64k",
    "CRF": "28",
    "Ext": "_.mkv"
  },
  "wx": {
    "AES": "aac",
    "AEA": "-q:a 3",
    "VES": "libx264",
    "ABR": "48k",
    "CRF": "33",
    "Ext": "_.m4v"
  },
  "x264-mp3": {
    "AES": "libmp3lame",
    "AEA": "-q:a 3",
    "VES": "libx264",
    "ABR": "256k",
    "CRF": "23",
    "Ext": "_.mp4"
  },
  "youtube": {
    "AES": "libvorbis",
    "AEA": "-q:a 5",
    "VES": "libx264",
    "VEA": "-pix_fmt yuv420p",
    "ABR": "",
    "CRF": "20",
    "Ext": "_.avi"
  }
}
	`
	// Get byte slice from string.
	bytes := []byte(text)

	// Unmarshal string into structs.
	var Defaults = map[string]Encoding{}
	json.Unmarshal(bytes, &Defaults)

	fmt.Printf("Defaults: '%+v'\n", Defaults)
}

/*

$ go run json.Unmarshal.defaults.go
Defaults: 'map[copy:{VES:copy AES:copy SES:-c:s copy VEP: AEP: SEP: VEA: AEA: ABR:64k CRF:42 Ext:_.mkv} webm:{VES:libvpx-vp9 AES:libopus SES:-c:s copy VEP: AEP: SEP: VEA: AEA: ABR:64k CRF:42 Ext:_.mkv} wx:{VES:libx264 AES:aac SES: VEP: AEP: SEP: VEA: AEA:-q:a 3 ABR:48k CRF:33 Ext:_.m4v} x264-mp3:{VES:libx264 AES:libmp3lame SES: VEP: AEP: SEP: VEA: AEA:-q:a 3 ABR:256k CRF:23 Ext:_.mp4} x265-opus:{VES:libx265 AES:libopus SES: VEP: AEP: SEP: VEA: AEA: ABR:64k CRF:28 Ext:_.mkv} youtube:{VES:libx264 AES:libvorbis SES: VEP: AEP: SEP: VEA:-pix_fmt yuv420p AEA:-q:a 5 ABR: CRF:20 Ext:_.avi}]'

*/
