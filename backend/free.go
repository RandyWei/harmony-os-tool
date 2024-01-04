package backend

import (
	"context"
	"strings"

	"icu.bughub.app/harmonyos-tool/backend/models"
)

// 查看当前设备内存情况
// ```free -h```
func Free(ctx context.Context) (string, error) {
	return AdbShellCommand(ctx, "free", "-h")
}

// top -n 1 -s 3 -o PID,VIRT,RES,SHR,S,%CPU,%MEM,TIME+,ARGS
// 打印出来前后出现乱码使用-b解决
func Top(ctx context.Context) (models.TopInfo, error) {

	result, err := AdbShellCommand(ctx, "top", "-n", "1", "-s", "3", "-o", "PID,VIRT,RES,SHR,%CPU,%MEM,TIME+,ARGS", "-b", "-m", "40")

	tasks := ""
	mem := ""
	processList := []models.Process{}
	//按行分割result
	//然后遍历结果
	//提取第一行，移除掉Tasks:、total、sleeping、running、stopped、zombie，然后按照,分割
	//提取第二行，移除掉Mem:、total、used、free、buffres，然后按照,分割
	//跳过第三、第四、第五行
	//剩下的行按空格分割
	//按行分割
	for index, line := range strings.Split(result, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		if index == 0 {
			tasks = strings.TrimSpace(line)
			tasks = strings.ReplaceAll(tasks, "Tasks:", "任务：")
			tasks = strings.ReplaceAll(tasks, "total", "个(总量)")
			tasks = strings.ReplaceAll(tasks, "sleeping", "个睡眠")
			tasks = strings.ReplaceAll(tasks, "running", "个正在运行")
			tasks = strings.ReplaceAll(tasks, "stopped", "个停止")
			tasks = strings.ReplaceAll(tasks, "zombie", "个僵死")
			continue
		}
		if index == 1 {
			mem = strings.TrimSpace(line)
			mem = strings.ReplaceAll(mem, "Mem:", "内存：")
			mem = strings.ReplaceAll(mem, "total", "(总量)")
			mem = strings.ReplaceAll(mem, "used", "已被使用")
			mem = strings.ReplaceAll(mem, "free", "空闲")
			mem = strings.ReplaceAll(mem, "buffers", "内核缓存内存量")
			continue
		}
		if index == 2 || index == 3 || index == 4 {
			continue
		}

		processInfos := strings.Split(line, " ")

		//	去除空格
		for i := 0; i < len(processInfos); i++ {
			if strings.TrimSpace(processInfos[i]) == "" {
				processInfos = append(processInfos[:i], processInfos[i+1:]...)
				i--
			}
		}

		if len(processInfos) < 7 {
			continue
		}

		args := processInfos[7]
		desc := ""
		if strings.Contains(args, "com.tencent.mm") {
			desc += "微信"
		}
		if strings.Contains(args, ":push") {
			desc += "、推送服务"
		}
		if strings.Contains(args, "com.tencent.tim") {
			desc += "TIM"
		}
		if strings.Contains(args, "com.tencent.mobileqq") {
			desc += "QQ"
		}

		process := models.Process{
			PID:  processInfos[0],
			VIRT: processInfos[1],
			RES:  processInfos[2],
			SHR:  processInfos[3],
			CPU:  processInfos[4],
			MEM:  processInfos[5],
			TIME: processInfos[6],
			ARGS: args,
			DESC: desc,
		}

		processList = append(processList, process)

	}

	return models.TopInfo{
		Mem:       mem,
		Processes: processList,
		Tasks:     tasks,
	}, err
}
