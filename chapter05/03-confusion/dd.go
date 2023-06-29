package main

type downloadFromNetDisk struct {
	secret   DynamicSecret
	filepath string
}

func (dd *downloadFromNetDisk) Download() {
	if err := dd.logincheak(); err != nil {
		//todo 重新登录
	}
	dd.downloadFromAliYun(dd.filepath)
}

func (dd *downloadFromNetDisk) logincheak() error {
	dd.cheakSecret(dd.secret.GetSecret())
	return nil
}

func (dd *downloadFromNetDisk) downloadFromAliYun(file string) {

}

func (dd *downloadFromNetDisk) cheakSecret(secret string) {
	//todo 调用阿里云验证接口去验证密码是否有效
}
