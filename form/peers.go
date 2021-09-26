package form

import (
	"github.com/astaxie/beego/validation"
	"go-fastdfs-web-go/models"
	"regexp"
)

type PeersForm struct {
	models.Peers
}

// Valid 自定义校验
func (install *PeersForm) Valid(v *validation.Validation) {
	v.Required(install.Name, "Name").Message("集群名称不能为空且在50字以内")
	v.MaxSize(install.Name, 50, "NameMax").Message("集群名称不能为空且在50字以内")

	v.MaxSize(install.GroupName, 50, "GroupNameMax").Message("组名称应在50字以内")

	v.Required(install.ServerAddress, "ServerAddress").Message("集群服务地址不能为空且在100字以内")
	v.MaxSize(install.ServerAddress, 100, "ServerAddressMax").Message("集群服务地址不能为空且在100字以内")

	v.MaxSize(install.ShowAddress, 100, "ShowAddressMax").Message("访问域名应在50字以内")

	urlRegexp := regexp.MustCompile(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)
	v.Match(install.ServerAddress, urlRegexp, "ServerAddressUrl").Message("请正确填写集群服务地址")

	if install.ShowAddress != "" {
		v.Match(install.ShowAddress, urlRegexp, "ShowAddressUrl").Message("请正确填写访问域名")
	}
}
