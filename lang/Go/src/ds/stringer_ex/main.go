package main

import "fmt"

//go:generate stringer -type=Pill
//go:generate stringer -type=Token -linecomment=true
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
type Token int

const (
	And Token = iota // &
	Or               // |
	Add              // +
	Sub              // -
	Period // .
)

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
	fmt.Printf("%d: %s\n", MsgtypeImage, MsgtypeImage)
	fmt.Printf("%d: %s\n", MsgtypeVoipmsg, MsgtypeVoipmsg)
	fmt.Printf("%d: %s\n", MsgtypeSys, MsgtypeSys)
}

/* 

1: Aspirin
3: Paracetamol
3: Paracetamol
2: +
3: 图片消息
50: VOIP消息
10000: 系统消息

*/
