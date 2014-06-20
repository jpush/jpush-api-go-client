package jpush

import "testing"
import "strings"

func Test_Basic(t *testing.T) {
	pf := NewPlatform()
	ad := NewAudience()
	ntf := NewNotification("notification")
	po := NewPushPayload(pf, ad).AddNotification(ntf)

	if s, e := po.ToJsonString(); e == nil {
		if !strings.Contains(s, `"audience":"all"`) {
			t.Errorf("test audience failed")
		}
		if !strings.Contains(s, `"platform":"all"`) {
			t.Errorf("test platform failed")
		}
		if !strings.Contains(s, `"notification":{"alert":"notification"}`) {
			t.Errorf("test notification failed")
		}
	} else {
		t.Error(e)
	}
}

func Test_Platform(t *testing.T) {
	pf := NewPlatform().AddPlatform(Platform_WP).AddPlatform(Platform_IOS)
	pf.AddPlatform(Platform_Android)
	ntf := NewNotification("notification")
	po := NewPushPayload(pf, NewAudience()).AddNotification(ntf)

	if s, e := po.ToJsonString(); e == nil {
		if !strings.Contains(s, `"platform":["winphone","ios","android"]`) {
			t.Errorf("test platform failed")
		}
	} else {
		t.Error(e)
	}
}

func Test_Audience(t *testing.T) {
	ad := NewAudience().AddAlias("alias").AddTag("tag").AddTadAnd("ta")
	ad.AddCond("cond", "true").AddRegistrationID("61")
	ntf := NewNotification("notification")
	po := NewPushPayload(NewPlatform(), ad).AddNotification(ntf)

	if s, e := po.ToJsonString(); e == nil {
		if !strings.Contains(s, `"alias":["alias"]`) ||
			!strings.Contains(s, `"cond":["true"]`) ||
			!strings.Contains(s, `"registration_id":["61"]`) ||
			!strings.Contains(s, `"tag":["tag"]`) ||
			!strings.Contains(s, `"tag_and":["ta"]`) {
			t.Errorf(s + "test audience failed")
		}
	} else {
		t.Error(e)
	}
}

func Test_Notification(t *testing.T) {
	ntf := NewNotification("notification")
	po := NewPushPayload(NewPlatform(), NewAudience()).AddNotification(ntf)

	if s, e := po.ToJsonString(); e == nil {
		if !strings.Contains(s, `"notification":{"alert":"notification"}`) {
			t.Errorf("test notification failed")
		}
	} else {
		t.Error(e)
	}

	ntf.AddAndroid(NewAndroid("android")).AddIOS(NewIOS("IOS"))
	ntf.AddWinphone(NewWinPhone("WP"))
	if s, e := po.ToJsonString(); e == nil {
		if !strings.Contains(s, `"android":{"alert":"android"}`) ||
			!strings.Contains(s, `"ios":{"alert":"IOS","badge":1,"sound":""}`) ||
			!strings.Contains(s, `"winphone":{"alert":"WP"}`) {
			t.Errorf("test notification failed")
		}
	} else {
		t.Error(e)
	}
}

func Test_Message(t *testing.T) {
	msg := NewMessage("message")
	po := NewPushPayload(NewPlatform(), NewAudience()).AddMessage(msg)

	if s, e := po.ToJsonString(); e == nil {
		if !strings.Contains(s, `"message":{"msg_content":"message"}`) {
			t.Errorf("test message failed")
		}
	} else {
		t.Error(e)
	}

	msg.AddTitle("title")
	if s, e := po.ToJsonString(); e == nil {
		if !strings.Contains(s, `"title":"title"`) {
			t.Errorf("test message failed")
		}
	} else {
		t.Error(e)
	}
}

func Test_Options(t *testing.T) {
	msg := NewMessage("message")
	po := NewPushPayload(NewPlatform(), NewAudience()).AddMessage(msg)
	opt := NewOptions()
	po.AddOptions(opt)

	if s, e := po.ToJsonString(); e == nil {
		if !strings.Contains(s, `"options":{}`) {
			t.Errorf("test options failed")
		}
	} else {
		t.Error(e)
	}

	opt.AddApnsProduction(false).AddOverrideMsgId(123).AddSendNo(321).AddTimeToLive(16)
	if s, e := po.ToJsonString(); e == nil {
		if !strings.Contains(s, `"apns_production":false`) ||
			!strings.Contains(s, `"override_msg_id":123`) ||
			!strings.Contains(s, `"sendno":321`) ||
			!strings.Contains(s, `"time_to_live":16`) {
			t.Errorf("test options failed")
		}
	} else {
		t.Error(e)
	}
}

func Test_Length(t *testing.T) {
	ios := NewIOS(strings.Repeat("r", 220))
	ntf := NewNotification("ntf").AddIOS(ios)
	po := NewPushPayload(NewPlatform(), NewAudience()).AddNotification(ntf)

	if _, e := po.ToJsonString(); e == nil {
		t.Errorf("ios maybe too large")
	}

	msg := NewMessage(strings.Repeat("r", 1200))
	po2 := NewPushPayload(NewPlatform(), NewAudience()).AddMessage(msg)

	if _, e := po2.ToJsonString(); e == nil {
		t.Errorf("message maybe too large")
	}
}
