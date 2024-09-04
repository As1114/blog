package flags

import (
	"blog/global"
	"bytes"
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
	"time"
)

func MysqlExport(c *cli.Context) (err error) {
	mysql := global.Config.Mysql

	timer := time.Now().Format("20060102")

	sqlPath := fmt.Sprintf("%s_%s.sql", mysql.DB, timer)

	// 调用系统命令， 执行mysqldump进行数据库导出
	cmder := fmt.Sprintf("mysqldump -u%s -p%s %s > %s", mysql.User, mysql.Password, mysql.DB, sqlPath)
	cmd := exec.Command("sh", "-c", cmder)

	var out, stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err = cmd.Run()

	if err != nil {
		global.Log.Errorln(err.Error(), stderr.String())
		return err
	}
	global.Log.Infof("sql文件 %s 导出成功", sqlPath)
	return nil
}
