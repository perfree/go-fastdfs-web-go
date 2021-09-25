package form

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/astaxie/beego/validation"
	uuid "github.com/satori/go.uuid"
	"go-fastdfs-web-go/models"
	"regexp"
)

type Install struct {
	Name          string
	GroupName     string
	ServerAddress string
	ShowAddress   string
	Account       string
	Password      string
	UserName      string
	Email         string
}

// GetUser 获取user
func (install *Install) GetUser() models.User {
	user := models.User{}
	user.Email = install.Email
	user.Account = install.Account
	user.Name = install.UserName
	user.Password = install.Password
	user.CredentialsSalt = uuid.NewV4().String()

	m5 := md5.New()
	m5.Write([]byte(user.Password))
	m5.Write([]byte(user.CredentialsSalt))
	st := m5.Sum(nil)
	user.Password = hex.EncodeToString(st)
	return user
}

// GetPeers 获取Peers
func (install *Install) GetPeers() models.Peers {
	peers := models.Peers{}
	peers.GroupName = install.GroupName
	peers.ServerAddress = install.ServerAddress
	peers.ShowAddress = install.ShowAddress
	peers.Name = install.Name
	return peers
}

// Valid 自定义校验
func (install *Install) Valid(v *validation.Validation) {
	v.Required(install.Name, "Name").Message("集群名称不能为空且在50字以内")
	v.MaxSize(install.Name, 50, "NameMax").Message("集群名称不能为空且在50字以内")

	v.MaxSize(install.GroupName, 50, "GroupNameMax").Message("组名称应在50字以内")

	v.Required(install.ServerAddress, "ServerAddress").Message("集群服务地址不能为空且在100字以内")
	v.MaxSize(install.ServerAddress, 100, "ServerAddressMax").Message("集群服务地址不能为空且在100字以内")

	v.MaxSize(install.ShowAddress, 100, "ShowAddressMax").Message("访问域名应在50字以内")

	v.Required(install.Account, "Account").Message("账户不能为空且在30字以内")
	v.MaxSize(install.Account, 30, "AccountMax").Message("账户不能为空且在30字以内")

	v.Required(install.Password, "Password").Message("密码不能为空且在30字以内")
	v.MaxSize(install.Password, 30, "PasswordMax").Message("密码不能为空且在30字以内")

	v.Required(install.UserName, "UserName").Message("用户名不能为空且在30字以内")
	v.MaxSize(install.UserName, 30, "UserNameMax").Message("用户名不能为空且在30字以内")

	v.Required(install.Email, "Email").Message("邮箱不能为空")
	v.Email(install.Email, "EmailValid").Message("请检查邮箱格式是否正确")

	urlRegexp := regexp.MustCompile(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)
	v.Match(install.ServerAddress, urlRegexp, "ServerAddressUrl").Message("请正确填写集群服务地址")

	if install.ShowAddress != "" {
		v.Match(install.ShowAddress, urlRegexp, "ShowAddressUrl").Message("请正确填写访问域名")
	}
}
