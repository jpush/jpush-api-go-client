package jpush

import (
	"encoding/json"
	"fmt"
)

type NamedValue map[string]interface{}

type TypePlatform string

const (
	Platform_Android      TypePlatform = "android"
	Platform_IOS          TypePlatform = "ios"
	Platform_WP           TypePlatform = "winphone"
	ALL                   string       = "all"
	ALIAS                 string       = "alias"
	ALERT                 string       = "alert"
	AUDIENCE              string       = "audience"
	APNS_PRODUCTION       string       = "apns_production"
	BUILDER_ID            string       = "builder_id"
	CONTENT_TYPE          string       = "content_type"
	EXTRAS                string       = "extras"
	IOS_BADGE             string       = "badge"
	IOS_SOUND             string       = "sound"
	IOS_CONTENT_AVAILABLE string       = "content-available"
	MAX_IOS_LENGTH        int          = 220
	MAX_CONTENT_LENGTH    int          = 1200
	MESSAGE               string       = "message"
	MSG_CONTENT           string       = "msg_content"
	NOTIFICATION          string       = "notification"
	OVERRIDE_MSG_ID       string       = "override_msg_id"
	OPTIONS               string       = "options"
	PLATFORM              string       = "platform"
	SENDNO                string       = "sendno"
	TAG                   string       = "tag"
	TAG_AND               string       = "tag_and"
	TITLE                 string       = "title"
	TTL                   string       = "time_to_live"
	REGISTRATION          string       = "registration_id"
	WP_OPEN_PAGE          string       = "_open_page"
)

type Platform []TypePlatform

func NewPlatform() *Platform {
	t := make(Platform, 0)
	return &t
}

func (vp *Platform) AddPlatform(platform ...TypePlatform) *Platform {
	*vp = append(*vp, platform...)
	return vp
}

func (vp *Platform) Value() interface{} {
	if len(*vp) == 0 {
		return ALL
	}

	return vp
}

type Audience map[string][]string

func NewAudience() *Audience {
	t := make(Audience)
	return &t
}

func (va *Audience) AddCond(key, value string) *Audience {
	if _, ok := (*va)[key]; !ok {
		(*va)[key] = make([]string, 0)
	}

	(*va)[key] = append((*va)[key], value)
	return va
}

func (va *Audience) AddTag(tag string) *Audience {
	return va.AddCond(TAG, tag)
}

func (va *Audience) AddTadAnd(tag_and string) *Audience {
	return va.AddCond(TAG_AND, tag_and)
}

func (va *Audience) AddAlias(alias string) *Audience {
	return va.AddCond(ALIAS, alias)
}

func (va *Audience) AddRegistrationID(id string) *Audience {
	return va.AddCond(REGISTRATION, id)
}

func (va *Audience) Value() interface{} {
	if len(*va) == 0 {
		return ALL
	}

	return va
}

type Extras NamedValue

func NewExtras() *Extras {
	e := make(Extras)
	return &e
}

func (ve *Extras) Add(name string, value interface{}) *Extras {
	(*ve)[name] = value
	return ve
}

type Android NamedValue

func NewAndroid(alert string) *Android {
	a := make(Android)
	a[ALERT] = alert
	return &a
}

func (va *Android) AddTitle(title string) *Android {
	(*va)[TITLE] = title
	return va
}

func (va *Android) Addbuilder_id(bid int) *Android {
	(*va)[BUILDER_ID] = bid
	return va
}

func (va *Android) AddExtras(e *Extras) *Android {
	(*va)[EXTRAS] = e
	return va
}

type iOS NamedValue

func NewIOS(alert string) *iOS {
	i := make(iOS)
	i[ALERT] = alert
	i[IOS_BADGE] = 1
	i[IOS_SOUND] = ""
	return &i
}

func (vi *iOS) AddSound(sd string) *iOS {
	(*vi)[IOS_SOUND] = sd
	return vi
}

func (vi *iOS) AddBadge(bd int) *iOS {
	(*vi)[IOS_BADGE] = bd
	return vi
}

func (vi *iOS) AddContentAvaiable(b bool) *iOS {
	(*vi)[IOS_CONTENT_AVAILABLE] = b
	return vi
}

func (vi *iOS) AddExtras(e *Extras) *iOS {
	(*vi)[EXTRAS] = e
	return vi
}

type Winphone NamedValue

func NewWinPhone(alert string) *Winphone {
	w := make(Winphone)
	w[ALERT] = alert
	return &w
}

func (vw *Winphone) AddTitle(title string) *Winphone {
	(*vw)[TITLE] = title
	return vw
}

func (vw *Winphone) AddOpenPage(op int) *Winphone {
	(*vw)[WP_OPEN_PAGE] = op
	return vw
}

func (vw *Winphone) AddExtras(e *Extras) *Winphone {
	(*vw)[EXTRAS] = e
	return vw
}

type Notification NamedValue

func NewNotification(defaultAlert string) *Notification {
	n := make(Notification)
	n[ALERT] = defaultAlert
	return &n
}

func (vn *Notification) AddAndroid(android *Android) *Notification {
	(*vn)[string(Platform_Android)] = android
	return vn
}

func (vn *Notification) AddIOS(ios *iOS) *Notification {
	(*vn)[string(Platform_IOS)] = ios
	return vn
}

func (vn *Notification) AddWinphone(wp *Winphone) *Notification {
	(*vn)[string(Platform_WP)] = wp
	return vn
}

type Message NamedValue

func NewMessage(msg_content string) *Message {
	m := make(Message)
	m[MSG_CONTENT] = msg_content
	return &m
}

func (om *Message) AddTitle(t string) *Message {
	(*om)[TITLE] = t
	return om
}

func (om *Message) AddContentType(ct string) *Message {
	(*om)[CONTENT_TYPE] = ct
	return om
}

func (om *Message) AddExtras(e *Extras) *Message {
	(*om)[EXTRAS] = e
	return om
}

type Options NamedValue

func NewOptions() *Options {
	o := make(Options)
	return &o
}

func (vo *Options) AddSendNo(sn int) *Options {
	(*vo)[SENDNO] = sn
	return vo
}

func (vo *Options) AddTimeToLive(ttl int) *Options {
	(*vo)[TTL] = ttl
	return vo
}

func (vo *Options) AddOverrideMsgId(omi int) *Options {
	(*vo)[OVERRIDE_MSG_ID] = omi
	return vo
}

func (vo *Options) AddApnsProduction(ap bool) *Options {
	(*vo)[APNS_PRODUCTION] = ap
	return vo
}

type PushPayload NamedValue

func NewPushPayload(p *Platform, a *Audience) *PushPayload {
	msg := make(PushPayload)
	msg[PLATFORM] = p.Value()
	msg[AUDIENCE] = a.Value()
	return &msg
}

func (vm *PushPayload) Build(key string, comp interface{}) *PushPayload {
	(*vm)[key] = comp
	return vm
}

func (vm *PushPayload) AddNotification(n *Notification) *PushPayload {
	return vm.Build(NOTIFICATION, n)
}

func (vm *PushPayload) AddMessage(m *Message) *PushPayload {
	return vm.Build(MESSAGE, m)
}

func (vm *PushPayload) AddOptions(o *Options) *PushPayload {
	return vm.Build(OPTIONS, o)
}

func (vm *PushPayload) ToJsonString() (string, error) {
	for k, v := range *vm {
		if v == nil {
			return "", fmt.Errorf("%s without a value", k)
		}
	}

	if (*vm)[NOTIFICATION] == nil && (*vm)[MESSAGE] == nil {
		return "", fmt.Errorf("without Notification/Message")
	}

	var length int
	if ntf := (*vm)[NOTIFICATION]; ntf != nil {
		if n := ntf.(*Notification); n != nil {
			if ios := (*n)[string(Platform_IOS)]; ios != nil {
				if b, e := json.Marshal(ios); e != nil || len(b) > MAX_IOS_LENGTH {
					return "", fmt.Errorf("invalidate ios notification")
				}
			}
		}

		if b, e := json.Marshal(ntf); e != nil {
			return "", e
		} else {
			length += len(b)
		}
	}

	if msg := (*vm)[MESSAGE]; msg != nil {
		if b, e := json.Marshal(msg); e != nil {
			return "", e
		} else {
			length += len(b)
		}
	}

	if length > MAX_CONTENT_LENGTH {
		return "", fmt.Errorf("Notification/Message too large")
	}

	b, e := json.Marshal(vm)
	if e != nil {
		return "", e
	}

	return string(b), e
}
