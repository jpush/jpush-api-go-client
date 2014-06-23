package main

import (
	"fmt"
	"github.com/jpush/jpush-api-go-client/push"
)

const (
	appKey       = "7d431e42dfa6a6d693ac2d04"
	masterSecret = "5e987ac6d2e04d95a9d8f0d1"
)

func push(po *jpush.PushPayload) {
	jpc := jpush.NewJPushClient(appKey, masterSecret)
	str, err := jpc.Push(po)
	fmt.Println("result : ", str, err)
}

func report() {
	jprc := NewJPushReportClient(appKey, masterSecret)
	r, e := jprc.GetReportObject(1613113584, 1229760629, 1174658841, 1174658641)
	if e == nil {
		for _, obj := range r {
			fmt.Println(obj.Msg_id, obj.Android_received, obj.Ios_apns_sent)
		}
	}
}

func main() {
	basic := jpush.NewPushPayload(jpush.NewPlatform(), jpush.NewAudience()).
		AddMessage(jpush.NewMessage("hello,jpush"))
	push(basic)

	/////////////////////////////////////////////////////

	et := jpush.NewExtras().Add("ex_num", 1).Add("ex_str", "str")

	and := jpush.NewAndroid("Android alert").AddTitle("android title").AddExtras(et)
	ios := jpush.NewIOS("iOS alert").AddBadge(123).AddExtras(et).AddContentAvaiable(true)
	wp := jpush.NewWinPhone("WP alert").AddExtras(et).AddOpenPage(234).AddTitle("wp")

	ntf := jpush.NewNotification("def-alert").AddIOS(ios).AddWinphone(wp).AddAndroid(and)
	opt := jpush.NewOptions().AddTimeToLive(1024).AddApnsProduction(false).AddSendNo(256)
	msg := jpush.NewMessage("message content").AddTitle("msg title").AddExtras(et)

	pf := jpush.NewPlatform()
	pf.AddPlatform(jpush.Platform_Android, jpush.Platform_IOS).AddPlatform(jpush.Platform_WP)
	ad := jpush.NewAudience()
	ad.AddTag("ShenZhen").AddTag("China").AddAlias("JPush")
	ad.AddRegistrationID("61").AddTadAnd("jpush.cn")

	po := jpush.NewPushPayload(pf, ad)
	po.AddNotification(ntf).AddOptions(opt).AddMessage(msg)
	s, e := po.ToJsonString()
	fmt.Println("PushPayload : ", s, e)

	push(po)

	////////////////////////////////////////////////////

	report()
}
