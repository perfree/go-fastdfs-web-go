<!DOCTYPE html>
<html lang="zh_CN">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<title>Go-FastDfs管理-安装</title>
<link rel="Shortcut Icon" href="/static/img/favicon.ico" type="image/x-icon"/>
<link rel="stylesheet" href="/static/libs/layui-v2.5.6/layui/css/layui.css">
<link href="/static/libs/step-lay/step.css" rel="stylesheet">
<link rel="stylesheet" href="/static/css/install.css" media="all">
</head>
<body>
<div class="layui-fluid">
    <div class="install-title">安装</div>
        <div class="layui-carousel" id="stepForm" lay-filter="stepForm" style="margin: 0 auto;">
            <div carousel-item>
                <div>
                    <form class="layui-form" style="margin: 0 auto;max-width: 460px;padding-top: 40px;">
                        <div class="layui-form-item">
                            <label class="layui-form-label">集群名称:</label>
                            <div class="layui-input-block">
                                <input name="Name" placeholder="请输入集群名称"  type="text" lay-verify="required" class="layui-input" value="集群1">
                                <span>填写自定义集群名称,如:集群1</span>
                            </div>
                        </div>
                        <div class="layui-form-item">
                            <label class="layui-form-label">组名称:</label>
                            <div class="layui-input-block">
                                <input name="GroupName" id="groupName" placeholder="请输入组名称" type="text" class="layui-input" >
                                <span>如果GoFastDfs开启了组支持(support_group_manage为true),则需要填写组名称(group),如:group1,未开启组支持请不要填写!</span>
                            </div>
                        </div>
                        <div class="layui-form-item">
                            <label class="layui-form-label">服务地址:</label>
                            <div class="layui-input-block">
                                <input name="ServerAddress" id="serverAddress" placeholder="请输入当前集群服务地址" lay-verify="required" type="text"  class="layui-input" >
                                <span>填写此集群的服务地址,如:http://153.153.234.32:8080,注意GoFastDfs需要配置(admin_ips)当前web服务主机的管理ip</span>
                            </div>
                        </div>
                        <div class="layui-form-item">
                            <label class="layui-form-label">访问域名:</label>
                            <div class="layui-input-block">
                                <input name="ShowAddress" id="showAddress" placeholder="资源访问域名"  type="text" class="layui-input" >
                                <span>配置能访问到已上传资源的域名,未做域名映射请不要填写!</span>
                            </div>
                        </div>
                        <div class="layui-form-item">
                            <div class="layui-input-block">
                                <button class="layui-btn" lay-submit lay-filter="formStep">
                                    &emsp;下一步&emsp;
                                </button>
                            </div>
                        </div>
                    </form>
                </div>
                <div>
                    <form class="layui-form" style="margin: 0 auto;max-width: 460px;padding-top: 40px;">
                        <div class="layui-form-item">
                            <label class="layui-form-label">用户名:</label>
                            <div class="layui-input-block">
                                <input name="UserName" placeholder="用户名" lay-verify="required" type="text" class="layui-input">
                            </div>
                        </div>
                        <div class="layui-form-item">
                            <label class="layui-form-label">登录账户:</label>
                            <div class="layui-input-block">
                                <input name="Account" placeholder="登录账户" lay-verify="required" type="text" class="layui-input">
                            </div>
                        </div>
                        <div class="layui-form-item">
                            <label class="layui-form-label">登录密码:</label>
                            <div class="layui-input-block">
                                <input name="Password" placeholder="登录密码" lay-verify="required"  type="password" class="layui-input">
                            </div>
                        </div>
                        <div class="layui-form-item">
                            <label class="layui-form-label">邮箱:</label>
                            <div class="layui-input-block">
                                <input name="Email"  placeholder="邮箱" lay-verify="required" type="text" class="layui-input">
                            </div>
                        </div>
                        <div class="layui-form-item">
                            <div class="layui-input-block">
                                <button type="button" class="layui-btn layui-btn-primary pre">上一步</button>
                                <button class="layui-btn" lay-submit lay-filter="formStep2">
                                    &emsp;安装&emsp;
                                </button>
                            </div>
                        </div>
                    </form>
                </div>
                <div>
                    <div style="text-align: center;margin-top: 90px;">
                        <i class="layui-icon layui-circle"
                           style="color: white;font-size:30px;font-weight:bold;background: #52C41A;padding: 20px;line-height: 80px;">&#xe605;</i>
                        <div style="font-size: 24px;color: #333;font-weight: 500;margin-top: 30px;">
                            安装成功
                        </div>
                        <div style="font-size: 14px;color: #666;margin-top: 20px;">欢迎使用Go-FastDfs 管理平台</div>
                    </div>
                    <div style="text-align: center;margin-top: 50px;">
                        <button class="layui-btn" onclick="toLogin()">账号登录</button>
                    </div>
                </div>
            </div>
    </div>
</div>

<script src="/static/libs/jquery/jquery-3.5.1.min.js"></script>
<script src="/static/libs/layui-v2.5.6/layui/layui.all.js"></script>
<script src="/static/libs/step-lay/step.js"></script>
<script type="text/javascript">
    layui.config({
        base:'/static/lib/step-lay/'
    }).use([ 'form', 'step'], function () {
        const $ = layui.$, form = layui.form, step = layui.step;
        // 检查本地是否安装了fastdfs
        const index = layer.load();
        $.get("/install/checkLocalServer", function(result){
            layer.close(index);
            if (result.Code === 200) {
                layer.msg("检测到本机(http://127.0.0.1:8080)存在运行的goFastDfs,已自动加载配置信息");
                let host = "";
                if (result.Data.data.host) {
                    host = result.Data.data.host.split("//")[1].split(":")[0];
                }
                if (result.Data.data.admin_ips.indexOf(host) >= 0) {
                    $("#serverAddress").val(result.Data.data.host);
                } else {
                    $("#serverAddress").val("http://127.0.0.1:8080");
                }
                $("#showAddress").val(result.Data.data.download_domain);
                if (result.Data.data.support_group_manage) {
                    $("#groupName").val(result.Data.data.group);
                }
            } else {
                layer.msg("未检测到本机(http://127.0.0.1:8080)运行goFastDfs,请手动配置goFastDfs信息");
            }
        });
        step.render({
            elem: '#stepForm',
            filter: 'stepForm',
            width: '100%', //设置容器宽度
            stepWidth: '750px',
            height: '560px',
            stepItems: [{
                title: '集群配置'
            }, {
                title: '账号配置'
            }, {
                title: '完成'
            }]
        });

        let peersFormValue = null;
        form.on('submit(formStep)', function (data) {
            peersFormValue = data.field;
            const index = layer.load();
            $.post("/install/checkServer",peersFormValue,function(result){
                if(result.Code === 200){
                    layer.close(index);
                    step.next('#stepForm');
                }else{
                    layer.close(index);
                    layer.msg(result.Msg);
                }
            });
            return false;
        });

        form.on('submit(formStep2)', function (data) {
            const index = layer.load();
            $.post("/install/doInstall",$.extend({}, peersFormValue,data.field),function(result){
                if(result.Code === 200){
                    layer.close(index);
                    step.next('#stepForm');
                }else{
                    layer.close(index);
                    layer.msg(result.Msg);
                }
            });
            return false;
        });

        $('.pre').click(function () {
            step.pre('#stepForm');
        });

        $('.next').click(function () {
            step.next('#stepForm');
        });
    })
    function toLogin() {
        window.location.href = "/login";
    }
</script>
</body>
</html>