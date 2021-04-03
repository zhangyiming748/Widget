package rarArchive

import (
	"fmt"
	"os/exec"
)

func UnRar(rarpath string, pass string) {
	cmd := exec.Command("unrar", "e", "-p"+pass, rarpath) //解压到当前文件夹
	fmt.Printf("密码不是%s",pass)
	out, _ := cmd.Output()
	fmt.Printf("输出的长度是%")
	if len(out) == 203 { //len 203 为成功，每个人不同
		fmt.Printf("password is : \"%s\" \n", pass)
	}else {
		fmt.Println("尝试下一组密码")
	}
}
