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

// dumpsys meminfo com.huawei.android.hwouc

var Apps0 = [...]models.App{
	{
		Id:   "com.huawei.android.hwouc",
		Name: "系统更新服务",
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
		Id:          "com.huawei.himovie",
		RelatedIds:  []string{"com.huawei.himovie.partner1", "com.huawei.himovie.partner2", "com.tencent.qqlivehuawei"},
		Description: "卸载华为视频后图库内无法打开视频",
		Name:        "华为视频",
	},
	{
		Id:          "com.huawei.music",
		Description: "",
		RelatedIds:  []string{"com.android.mediacenter"},
		Name:        "华为音乐",
	},
	{
		Id:          "com.huawei.pcassistant",
		Description: "",
		RelatedIds:  []string{"com.huawei.android.instantshare"},
		Name:        "华为分享",
	},
	{
		Id:          "com.huawei.hwireader",
		Description: "",
		RelatedIds:  []string{"com.huawei.hwreader.al", "com.huawei.hnreader", "com.huawei.hwread.al", "com.huawei.hwireader"},
		Name:        "华为阅读",
	},
	{
		Id:          "com.huawei.browser",
		Description: "",
		RelatedIds:  []string{"com.android.browser"},
		Name:        "华为浏览器",
	},
	{
		Id:          "com.huawei.appmarket",
		Description: "",
		Name:        "应用商店",
	},
	{
		Id:          "com.huawei.camera",
		Description: "",
		Name:        "华为相机",
	},
	{
		Id:          "com.huawei.hwid",
		Description: "",
		Name:        "HMS core ",
	},
	{
		Id:   "com.huawei.phoneservice",
		Name: "服务/我的华为",
	},
	{
		Id:          "com.huawei.ohos.famanager",
		Description: "请最后卸载服务中心，据说这样能够避免弹出“服务请求失败”",
		Name:        "服务中心",
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
		Id:          "com.huawei.scanner",
		Description: "已释放【扫一扫】相机插件,也叫【智慧视觉】，扫码功能建议使用华为浏览器或其它，打开相机的功耗较大",
		Name:        "智慧视觉(扫一扫)",
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
	{
		Id:          "com.android.emergency",
		Description: "",
		Name:        "紧急信息",
	},
	{
		Id:          "com.huawei.gameassistant",
		Description: "",
		Name:        "游戏空间",
	},
	{
		Id:          "com.android.managedprovisioning",
		Description: "",
		Name:        "工作资料设置",
	},
	{
		Id:          "com.android.htmlviewer",
		Description: "",
		Name:        "html查看器",
	},
	{
		Id:          "com.baidu.input_huawei",
		Description: "据反馈，需要先关闭安全输入法输入密码，否则卸载后不会弹出输入法\n因此请您在设置中将 【输入密码时，启用安全输入键盘】 关闭,  然后继续执行卸载\n不同机型不同情况，希望您谨慎考虑卸载自带输入法，【【小白不建议操作】】\n【无法解锁补救措施】\n在卸载前备份数据！！或提前开启仅充电模式下连接ADB!\n外接键盘解锁！",
		RelatedIds:  []string{"com.huawei.secime"},
		Name:        "百度输入法",
	},
	{
		Id:          "com.huawei.motionservice",
		Description: "可能影响到智慧感知、隔空操作、注视不息屏，有这些功能的机子不建议删",
		Name:        "悬浮球",
	},
	{
		Id:          "com.huawei.HwMultiScreenShot",
		Description: "",
		Name:        "滚动截屏",
	},
	{
		Id:          "com.huawei.android.FloatTasks",
		Description: "",
		Name:        "悬浮导航",
	},
	{
		Id:          "com.huawei.videoeditor",
		Description: "",
		Name:        "花瓣剪辑",
	},
	{
		Id:          "com.huawei.controlcenter",
		Description: "",
		Name:        "多屏协同",
	},
	{
		Id:          "com.example.android.notepad",
		Description: "",
		Name:        "备忘录",
	},
	{
		Id:          "com.huawei.aod",
		Description: "",
		Name:        "息屏显示",
	},
	{
		Id:          "com.huawei.meetime",
		Description: "",
		RelatedIds:  []string{"com.huawei.hwvoipservice"},
		Name:        "畅连",
	},
	{
		Id:          "com.iflytek.speechsuite",
		Description: "",
		RelatedIds:  []string{"com.sohu.sohuvideo.emplayer"},
		Name:        "讯飞语音引擎",
	},
	{
		Id:          "com.huawei.android.findmyphone",
		Description: "",
		Name:        "查找手机",
	},
	{
		Id:          "com.android.mms.service",
		Description: "",
		Name:        "彩信服务",
	},
	{
		Id:          "com.huawei.magazine",
		Description: "",
		Name:        "杂志锁屏",
	},
	{
		Id:          "com.huawei.localBackup",
		Description: "",
		Name:        "备份",
	},
	{
		Id:          "com.huawei.android.thememanager",
		Description: "",
		Name:        "主题",
	},
	{
		Id:          "com.huawei.hwdetectrepair",
		Description: "",
		Name:        "智能检测",
	},
	{
		Id:          "com.android.calculator2",
		Description: "",
		Name:        "计算器",
	},
	{
		Id:          "com.android.emergency",
		Description: "",
		Name:        "个人紧急信息",
	},
	{
		Id:          "com.android.deskclock",
		Description: "",
		Name:        "时 钟",
	},
	{
		Id:          "com.android.printspooler",
		Description: "",
		RelatedIds:  []string{"com.huawei.printservice"},
		Name:        "打印服务",
	},
	{
		Id:          "com.huawei.hicar",
		Description: "",
		Name:        "HiCar智行",
	},
	{
		Id:          "com.huawei.hiai",
		Description: "",
		Name:        "华为智慧引擎",
	},
	{
		Id:          "com.huawei.hidisk",
		Description: "云空间可能删不掉，被系统保护\n文件管理可能会一同被卸载",
		RelatedIds:  []string{"com.huawei.hicloud"},
		Name:        "云空间",
	},
	{
		Id:   "com.huawei.android.karaoke",
		Name: "k歌特效",
	},
	{
		Id:   "com.huawei.android.projectmenu",
		Name: "工程菜单",
	},
	{
		Id:   "com.huawei.scenepack",
		Name: "旅行助手",
	},
	{
		Id:          "com.huawei.android.totemweather",
		Description: "建议使用pure天气替代，更小更简洁更流\n或者可以使用魅族天气，简洁流畅\n并且桌面的天气时间小工具更加美观实用，款式更丰富",
		Name:        "天气",
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
		Name:        "冗余用户体验计划1",
	},
	{
		Id:          "com.huawei.hiview",
		Description: "com.huawei.hiview",
		RelatedIds:  []string{"com.huawei.hiviewtunnel"},
		Name:        "冗余用户体验计划2",
	},
	{
		Id:          "com.huawei.android.UEInfoCheck",
		Description: "com.huawei.android.UEInfoCheck",
		Name:        "冗余用户体验计划3",
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

// 性能优化
var Apps5 = [...]models.App{
	{
		Id:          "com.huawei.powergenie",
		Description: "就是powergenie，解除后能拉满手机功耗",
		Name:        "功耗管理",
	},
	{
		Id:          "com.huawei.iaware",
		Description: "就是iaware，解除后不杀主动后台",
		Name:        "后台管理",
	},
	{
		Id:          "com.huawei.android.hwaps",
		Description: "解除后不降分辨率",
		Name:        "HWAPS",
	},
}

// 华为功能模块
var HWModules = [...]models.Module{
	{
		Id:   "0",
		Type: "disable",
		Apps: Apps0[:],
	},
	{
		Id:   "1",
		Type: "uninstall",
		Name: "华为全家桶",
		Apps: Apps1[:],
	},
	{
		Id:   "2",
		Type: "uninstall",
		Name: "智慧增值服务",
		Apps: Apps2[:],
	},
	{
		Id:   "3",
		Type: "uninstall",
		Name: "系统功能",
		Apps: Apps3[:],
	},
	{
		Id:   "4",
		Type: "uninstall",
		Name: "冗余服务",
		Apps: Apps4[:],
	},
	{
		Id:   "5",
		Type: "uninstall",
		Name: "性能优化",
		Apps: Apps5[:],
	},
}
