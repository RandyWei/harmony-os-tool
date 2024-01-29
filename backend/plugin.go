package backend

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"icu.bughub.app/harmonyos-tool/backend/utils"
)

func TryRunAdb() bool {
	cmd := exec.Command("adb", "version")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	execResult := string(output)

	fmt.Println(execResult)

	return strings.Contains(execResult, "Android Debug Bridge")
}

/**
 * 下载安装adb
 */
func InstallAdb(ctx context.Context, appDir string) (bool, error) {
	//https://look.yun.chinahrt.com/u/customer/appPackage/20241/8/wu_1hjjdp9n51n1do231otb54v1fh6c.apk
	//https://gitee.com/RandyWei/harmony-os-tool/raw/main/attachment/platform-tools_r34.0.5-windows.zip

	//https://look.yun.chinahrt.com/u/customer/appPackage/20241/8/wu_1hjjdln602sma6n7j58113j54.apk
	//https://gitee.com/RandyWei/harmony-os-tool/raw/main/attachment/platform-tools_r34.0.5-darwin.zip
	//获取当前系统类型
	downloadUrlPrefix := "https://look.yun.chinahrt.com/u/customer/appPackage/20241/8/"
	envInfo := wailsRuntime.Environment(ctx)
	fileName := "wu_1hjjdp9n51n1do231otb54v1fh6c"
	if envInfo.Platform == "darwin" {
		fileName = "wu_1hjjdln602sma6n7j58113j54"
	}
	downloadUrl := fmt.Sprintf("%s%s.apk", downloadUrlPrefix, fileName)

	cacheDir := appDir + string(os.PathSeparator) + "cache"
	//清空缓存目录
	os.RemoveAll(cacheDir)
	os.MkdirAll(cacheDir, 0755)

	//下载zip
	resp, err := http.Get(downloadUrl)
	if err != nil {
		utils.LogE(ctx, err.Error())
		return false, fmt.Errorf("下载失败了")
	}
	defer resp.Body.Close()

	out, err := os.Create(cacheDir + string(os.PathSeparator) + "platform-tools.zip")
	if err != nil {
		utils.LogE(ctx, err.Error())
		return false, fmt.Errorf("创建文件失败了")
	}
	defer out.Close()
	//然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		utils.LogE(ctx, err.Error())
		return false, fmt.Errorf("拷贝文件失败了")
	}

	//解压platform-tools.zip
	if err := unzip(appDir, cacheDir+string(os.PathSeparator)+"platform-tools.zip"); err != nil {
		utils.LogE(ctx, err.Error())
		return false, fmt.Errorf("解压失败了")
	}

	//删除zip文件
	if err := os.Remove(cacheDir + string(os.PathSeparator) + "platform-tools.zip"); err != nil {
		utils.LogE(ctx, err.Error())
		return false, fmt.Errorf("删除文件失败了")
	}

	return true, nil
}

func unzip(dst, src string) (err error) {
	// 打开压缩文件，这个 zip 包有个方便的 ReadCloser 类型
	// 这个里面有个方便的 OpenReader 函数，可以比 tar 的时候省去一个打开文件的步骤
	zr, err := zip.OpenReader(src)

	if err != nil {
		return
	}
	defer zr.Close()
	// 如果解压后不是放在当前目录就按照保存目录去创建目录
	if dst != "" {
		if err := os.MkdirAll(dst, 0755); err != nil {
			return err
		}
	}

	// 遍历 zr ，将文件写入到磁盘
	for _, file := range zr.File {
		path := filepath.Join(dst, file.Name)

		// 如果是目录，就创建目录
		if file.FileInfo().IsDir() {
			if err := os.MkdirAll(path, file.Mode()); err != nil {
				return err
			}
			// 因为是目录，跳过当前循环，因为后面都是文件的处理
			continue
		}

		//使用匿名函数，解决无法defer
		err := func() error {
			rc, err := file.Open()
			if err != nil {
				return err
			}
			defer rc.Close()

			fw, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
			if err != nil {
				return err
			}
			defer fw.Close()
			n, err := io.Copy(fw, rc)
			if err != nil {
				return err
			}
			// 将解压的结果输出
			fmt.Printf("成功解压 %s ，共写入了 %d 个字符的数据\n", path, n)
			return nil
		}()
		if err != nil {
			return err
		}

	}
	return nil
}
