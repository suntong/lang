// Code generated by "stringer -type=MessageType -linecomment=true"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MsgtypeText-1]
	_ = x[MsgtypeImage-3]
	_ = x[MsgtypeVoice-34]
	_ = x[MsgtypeVerifymsg-37]
	_ = x[MsgtypePossiblefriendMsg-40]
	_ = x[MsgtypeSharecard-42]
	_ = x[MsgtypeVideo-43]
	_ = x[MsgtypeEmoticon-47]
	_ = x[MsgtypeLocation-48]
	_ = x[MsgtypeApp-49]
	_ = x[MsgtypeVoipmsg-50]
	_ = x[MsgtypeVoipnotify-52]
	_ = x[MsgtypeVoipinvite-53]
	_ = x[MsgtypeMicrovideo-62]
	_ = x[MsgtypeSys-10000]
	_ = x[MsgtypeRecalled-10002]
}

const _MessageType_name = "文本消息图片消息语音消息认证消息好友推荐消息名片消息视频消息表情消息地理位置消息APP消息VOIP消息VOIP结束消息VOIP邀请小视频消息系统消息消息撤回"

var _MessageType_map = map[MessageType]string{
	1:     _MessageType_name[0:12],
	3:     _MessageType_name[12:24],
	34:    _MessageType_name[24:36],
	37:    _MessageType_name[36:48],
	40:    _MessageType_name[48:66],
	42:    _MessageType_name[66:78],
	43:    _MessageType_name[78:90],
	47:    _MessageType_name[90:102],
	48:    _MessageType_name[102:120],
	49:    _MessageType_name[120:129],
	50:    _MessageType_name[129:139],
	52:    _MessageType_name[139:155],
	53:    _MessageType_name[155:165],
	62:    _MessageType_name[165:180],
	10000: _MessageType_name[180:192],
	10002: _MessageType_name[192:204],
}

func (i MessageType) String() string {
	if str, ok := _MessageType_map[i]; ok {
		return str
	}
	return "MessageType(" + strconv.FormatInt(int64(i), 10) + ")"
}
