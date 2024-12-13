package initialize

import (
	"bufio"
	"os"
	"strings"

	"github.com/songzhibin97/gkit/cache/local_cache"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

func OtherInit() {
	// 解析JWT的过期时间
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		// 如果解析失败，抛出异常
		panic(err)
	}
	// 解析JWT的缓冲时间
	_, err = utils.ParseDuration(global.GVA_CONFIG.JWT.BufferTime)
	if err != nil {
		// 如果解析失败，抛出异常
		panic(err)
	}

	// 创建一个新的黑名单缓存，设置默认过期时间
	global.BlackCache = local_cache.NewCache(
		local_cache.SetDefaultExpire(dr),
	)
	// 打开go.mod文件
	file, err := os.Open("go.mod")
	if err == nil && global.GVA_CONFIG.AutoCode.Module == "" {
		// 如果文件打开成功且模块名为空
		scanner := bufio.NewScanner(file)
		scanner.Scan() // 读取文件的第一行
		// 从第一行中提取模块名并去掉前缀"module "
		global.GVA_CONFIG.AutoCode.Module = strings.TrimPrefix(scanner.Text(), "module ")
	}
}
