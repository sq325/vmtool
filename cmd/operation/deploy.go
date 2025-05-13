package operation

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var DeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy vm-cluster",
	Long: `Deploy vm-cluster
	1. 检查 bin/ 中二进制文件是否存在
	2. 检查 conf 目录是否存在cluster.conf和daemon.conf，不存在生成默认配置，并退出
	3. 检查端口是否占用
	4. 创建 data/vmstorage-data 目录
	5. 创建 logs 目录
	6. 创建 tmp/vmselect 目录
	7. 创建 data/backup 目录
	8. 生成 conf/vmauth.conf`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// 需要检查的二进制文件列表
var requiredBinaries = []string{
	"vmselect",
	"vmstorage",
	"vminsert",
	"vmauth",
	"vmbackup",
	"vmrestore",
}

func check() error {
	// 检查 ./bin/ 中是否有vmselect, vmstorage, vminsert, vmauth, vmbackup, vmrestore
	binDir := "./bin"

	// 检查bin目录是否存在
	if _, err := os.Stat(binDir); os.IsNotExist(err) {
		return fmt.Errorf("bin 目录不存在，请先创建目录并放入必要的可执行文件")
	}

	// 检查每个二进制文件
	for _, binary := range requiredBinaries {
		binPath := filepath.Join(binDir, binary)
		info, err := os.Stat(binPath)
		if os.IsNotExist(err) {
			return fmt.Errorf("二进制文件 %s 不存在", binary)
		}

		// 检查文件是否可执行
		if info.Mode().Perm()&0111 == 0 {
			return fmt.Errorf("二进制文件 %s 不可执行，请添加执行权限", binary)
		}
	}

	fmt.Println("所有必要的二进制文件检查完成")
	return nil
}

func createDirectories() error {
	// 需要创建的目录列表
	dirs := []string{
		"./vmstorage-data", // 存储数据目录
		"./data",           // 数据目录
		"./logs",           // 日志目录
		"./tmp/vmselect",   // vmselect缓存目录
		"./backup",         // 备份目录
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("创建目录 %s 失败: %v", dir, err)
		}
		fmt.Printf("目录 %s 创建成功\n", dir)
	}

	return nil
}

func generateVmauthConfig() error {
	// 生成默认的 vmauth 配置文件
	const vmauthConfigPath = "./vmauth.conf"

	// 检查文件是否已存在，如果存在则不覆盖
	if _, err := os.Stat(vmauthConfigPath); err == nil {
		fmt.Printf("文件 %s 已存在，跳过生成\n", vmauthConfigPath)
		return nil
	}

	// 默认配置内容
	configContent := `
# Default vmauth configuration
unauthorized_user:
  url_map:
    - src_paths:
      - "/insert/.+"
      url_prefix:
        - "http://localhost:8480/"
    - src_paths:
      - "/select/.+"
      url_prefix:
        - "http://localhost:8481/"
    - src_paths:
      - "/admin/.+"
      url_prefix:
        - "http://localhost:8481/"
`

	// 写入配置文件
	if err := os.WriteFile(vmauthConfigPath, []byte(configContent), 0644); err != nil {
		return fmt.Errorf("写入配置文件失败: %v", err)
	}

	fmt.Printf("配置文件 %s 生成成功\n", vmauthConfigPath)
	return nil
}

func init() {
	DeployCmd.Flags()
}
