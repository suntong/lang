package main

import "fmt"

//go:generate stringer -type=Pill
//go:generate stringer -type=Token -linecomment=true
//go:generate stringer -type=Note -linecomment=true
//go:generate stringer -type=MessageType -linecomment=true

// https://pkg.go.dev/golang.org/x/tools/cmd/stringer
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)

// https://www.reddit.com/r/golang/comments/7l05zx/stringer_now_supports_inline_comments/
// https://go.dev/play/p/t8n91_GF0Od
type Token int

const (
	And    Token = iota // &
	Or                  // |
	Add                 // +
	Sub                 // -
	Period              // .
)

type Note int

const (
	_ Note = iota
	C
	Db
	D
	Eb
	E
)

// type Pitch int

// const (
// 	_  Pitch = iota //
// 	C               // 261.63
// 	Db              // 277.18
// 	D               // 293.66
// 	Eb              // 311.13
// 	E               // 329.63
// )
// X: xx redeclared in this block

var Pitch = map[string]string{
	"C":  "261.63",
	"C#": "277.18",
	//	"Db": Pitch["C#"], // X: initialization loop for Pitch
	"Db": "277.18",
	"D":  "293.66",
	"Eb": "311.13",
	"E":  "329.63",
}

type MessageType int

const (
	// MSGTYPE_TEXT
	MsgtypeText MessageType = 1 // 文本消息
	// MSGTYPE_IMAGE
	MsgtypeImage MessageType = 3 // 图片消息
	// MSGTYPE_VOICE
	MsgtypeVoice MessageType = 34 // 语音消息
	// MSGTYPE_VERIFYMSG
	MsgtypeVerifymsg MessageType = 37 // 认证消息
	// MSGTYPE_POSSIBLEFRIEND_MSG
	MsgtypePossiblefriendMsg MessageType = 40 // 好友推荐消息
	// MSGTYPE_SHARECARD
	MsgtypeSharecard MessageType = 42 // 名片消息
	// MSGTYPE_VIDEO
	MsgtypeVideo MessageType = 43 // 视频消息
	// MSGTYPE_EMOTICON
	MsgtypeEmoticon MessageType = 47 // 表情消息
	// MSGTYPE_LOCATION
	MsgtypeLocation MessageType = 48 // 地理位置消息
	// MSGTYPE_APP
	MsgtypeApp MessageType = 49 // APP消息
	// MSGTYPE_VOIPMSG
	MsgtypeVoipmsg MessageType = 50 // VOIP消息
	// MSGTYPE_VOIPNOTIFY
	MsgtypeVoipnotify MessageType = 52 // VOIP结束消息
	// MSGTYPE_VOIPINVITE
	MsgtypeVoipinvite MessageType = 53 // VOIP邀请
	// MSGTYPE_MICROVIDEO
	MsgtypeMicrovideo MessageType = 62 // 小视频消息
	// MSGTYPE_SYS
	MsgtypeSys MessageType = 10000 // 系统消息
	// MSGTYPE_RECALLED
	MsgtypeRecalled MessageType = 10002 // 消息撤回
)

func main() {
	fmt.Printf("%d: %s\n", Aspirin, Aspirin)
	fmt.Printf("%d: %s\n", Paracetamol, Paracetamol)
	fmt.Printf("%d: %s\n", Acetaminophen, Acetaminophen)

	fmt.Printf("%d: %s\n", Add, Add)
	fmt.Printf("%d: %s\n", C, C)
	fmt.Printf("%d: %s\n", E, Eb)
	fmt.Printf("%v %s\n", mapNoteToEnum, Pitch["Eb"])

	fmt.Printf("%d: %s\n", MsgtypeImage, MsgtypeImage)
	fmt.Printf("%d: %s\n", MsgtypeVoipmsg, MsgtypeVoipmsg)
	fmt.Printf("%d: %s\n", MsgtypeSys, MsgtypeSys)
}

var mapNoteToEnum = func() map[string]Note {
	m := make(map[string]Note)
	for i := C; i <= E; i++ {
		m[i.String()] = i
	}
	return m
}()

// func getNotePitch(n string) string {
// 	return Pitch(mapNoteToEnum[n])
// }

/*

 go install golang.org/x/tools/cmd/stringer@latest
 or,
 go install golang.org/x/tools/cmd/stringer

# must build then run, not `go run main.go`
$ stringer_ex
1: Aspirin
3: Paracetamol
3: Paracetamol
2: +
1: 261.63
5: 311.13
3: 图片消息
50: VOIP消息
10000: 系统消息

*/
