package constants

import "icu.bughub.app/harmonyos-tool/backend/models"

// com.huawei.powergenie
// com.huawei.iaware
// com.huawei.android.hwaps

//系统禁更
//adb shell pm disable-user com.huawei.android.hwouc

// 谷歌服务 HMS服务
// com.google.android.gms
// com.google.android.gsf
// com.android.vending
// com.google.android.onetimeinitializer
// com.google.android.partnersetup
// com.google.android.marvin.talkback
// com.android.ext.services
// com.google.android.backuptransport
// com.google.android.gsf.login
// com.google.android.printservice.recommendation
// com.google.android.feedback
// com.google.android.syncadapters.calendar
// com.google.android.syncadapters.contacts
// com.google.android.gsf.login
var Apps0 = [...]models.App{
	{
		Id:   "com.huawei.fastapp",
		Name: "谷歌服务",
	},
}

// 华为全家桶
var Apps1 = [...]models.App{
	{
		Id:          "com.huawei.wallet",
		Description: "不用NFC建议删",
		Name:        "华为钱包",
	},
	{
		Id:          "com.huawei.search",
		Description: "下拉搜索：误解+广告",
		Name:        "华为视频",
	},
	{
		Id:          "com.huawei.music",
		Description: "",
		RelatedIds:  []string{"com.android.mediacenter"},
		Name:        "华为音乐",
	},
	{
		Id:          "com.huawei.ohos.famanager",
		Description: "请最后卸载服务中心，据说这样能够避免弹出“服务请求失败”",
		Name:        "服务/我的华为",
	},
}

// 智慧增值服务
var Apps2 = [...]models.App{
	{
		Id:          "com.huawei.search",
		Description: "下拉搜索：误解+广告",
		RelatedIds:  []string{"com.huawei.searchservice"},
		Name:        "智慧搜索",
	},
	{
		Id:          "com.huawei.vassistant",
		Description: "后台占运存特大",
		Name:        "智慧语音",
	},
	{
		Id:          "com.huawei.hitouch",
		Description: "双指按屏扫文字",
		RelatedIds:  []string{"com.huawei.hiaction", "com.huawei.contentsensor"},
		Name:        "智慧识屏",
	},
	{
		Id:          "com.huawei.ohos.suggestion",
		Description: "桌面的智慧卡片",
		Name:        "智慧建议",
	},
	{
		Id:          "com.huawei.intelligent",
		Description: "负一屏，关闭也占运存",
		Name:        "智慧助手-今天",
	},
	{
		Id:          "com.huawei.fastapp",
		Description: "类似小程序中心",
		Name:        "快应用中心",
	},
	{
		Id:         "com.huawei.health",
		RelatedIds: []string{"com.huawei.ohos.health"},
		Name:       "运行健康",
	},
}

// 系统功能
var Apps3 = [...]models.App{
	{
		Id:          "com.huawei.skytone",
		Description: "",
		RelatedIds:  []string{"com.huawei.hiskytone"},
		Name:        "天际通",
	},
}

// 冗余服务
var Apps4 = [...]models.App{
	{
		Id:          "com.android.stk",
		Description: " 不是sim卡管理，早年运营商要求加的功能",
		Name:        "sim卡应用",
	},
	{
		Id:          "com.huawei.hifolder",
		Description: "桌面文件夹内推荐应用的毒瘤玩意",
		Name:        "精品推荐",
	},
	{
		Id:          "com.huawei.bd",
		Description: "com.huawei.bd",
		Name:        "com.huawei.bd",
	},
	{
		Id:          "com.huawei.hiview",
		Description: "com.huawei.hiview",
		RelatedIds:  []string{"com.huawei.hiviewtunnel"},
		Name:        "com.huawei.hiview",
	},
	{
		Id:          "com.huawei.android.UEInfoCheck",
		Description: "com.huawei.android.UEInfoCheck",
		Name:        "com.huawei.android.UEInfoCheck",
	},
	{
		Id:          "com.android.cellbroadcastreceiver",
		Description: "安卓早期遗留",
		Name:        "小区广播",
	},
	{
		Id:          "com.huawei.spaceservice",
		Description: "用于管理地理禁区的，注意是你被管理",
		Name:        "地理围栏服务",
	},
	{
		Id:          "com.huawei.tips",
		Description: "早年安卓2.3推出的桌面提醒小机器人服务",
		Name:        "tips",
	},
	{
		Id:          "com.google.android.backuptransport",
		Description: "与框架和服务无关",
		Name:        "谷歌备份传输",
	},
	{
		Id:          "com.huawei.pengine",
		Description: "现在被智慧建议代替了，但服务依然留在手机里",
		Name:        "智能建议服务",
	},
	//猜测是早期的程序，HarmonyOS 4.0已经没有了
	{
		Id:          "com.google.android.marvin.talkback",
		Description: "残障人士以听代视的服务，卸载影响跳过广告的app",
		Name:        "无障碍服务",
	},
	{
		Id:          "com.huawei.arengine.service",
		Description: "大多AR软件自带AR组件，但如果您依赖华为【AR测量】功能，您可以保留",
		RelatedIds:  []string{"com.huawei.vrservice"},
		Name:        "AR引擎",
	},
	{
		Id:          "com.huawei.bonevoiceui",
		Description: "若您购买了有骨声纹功能的华为产品，建议您保留",
		Name:        "骨声纹插件",
	},
	{
		Id:          "com.huawei.rcsserviceapplication",
		Description: "早年中国移动推出的移动数字短信，给移动用户发短信用流量，资费0.01元",
		Name:        "RCS服务",
	},
	//猜测是早期的程序，HarmonyOS 4.0已经没有了
	{
		Id:          "com.android.dreams.basic",
		Description: "可能影响到一镜到底，没有一镜到底的机型可以卸载",
		RelatedIds:  []string{"com.android.dreams.phototable"},
		Name:        "屏保程序",
	},
	{
		Id:          "com.huawei.livewallpaper.paradise",
		Description: "可能影响到一镜到底，没有一镜到底的机型可以卸载",
		RelatedIds:  []string{"com.huawei.livewallpaper.mountaincloud"},
		Name:        "动态壁纸程序",
	},
}
