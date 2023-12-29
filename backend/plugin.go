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
	//https://gitee.com/RandyWei/harmony-os-tool/raw/main/attachment/platform-tools_r34.0.5-windows.zip
	//https://gitee.com/RandyWei/harmony-os-tool/raw/main/attachment/platform-tools_r34.0.5-darwin.zip
	//获取当前系统类型
	downloadUrlPrefix := "https://docs.bughub.icu"
	envInfo := wailsRuntime.Environment(ctx)
	downloadUrl := fmt.Sprintf("%s/platform-tools_r34.0.5-%s.zip", downloadUrlPrefix, envInfo.Platform)

	fmt.Println(downloadUrl)
	cacheDir := appDir + string(os.PathSeparator) + "cache"
	//清空缓存目录
	os.RemoveAll(cacheDir)
	os.MkdirAll(cacheDir, 0755)
	fmt.Println(cacheDir)

	//下载zip
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return false, fmt.Errorf("下载失败了")
	}
	defer resp.Body.Close()

	out, err := os.Create(cacheDir + string(os.PathSeparator) + "platform-tools.zip")
	if err != nil {
		return false, fmt.Errorf("创建文件失败了")
	}
	defer out.Close()
	//然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return false, fmt.Errorf("拷贝文件失败了")
	}

	//解压platform-tools.zip
	if err := UnZip(appDir, cacheDir+string(os.PathSeparator)+"platform-tools.zip"); err != nil {
		return false, fmt.Errorf("解压失败了")
	}

	return true, nil
}

func UnZip(dst, src string) (err error) {
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
