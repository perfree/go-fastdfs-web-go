{{define "login/login.tpl"}}
<!DOCTYPE html>
<html lang="zh_CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <title>Go-FastDfs管理-登录</title>
    <link rel="Shortcut Icon" href="/static/img/favicon.ico" type="image/x-icon"/>
    <link rel="stylesheet" href="/static/libs/layui-v2.5.6/layui/css/layui.css">
    <link rel="stylesheet" href="/static/css/login.css">
    <script type="text/javascript">
        if (window !== top)
            top.location.href = location.href;
    </script>
</head>
<body>
<div class="full">
    <div class="login-top-box"></div>
    <div class="login-box">
        <h2 class="login-title">登录</h2>
        <form class="layui-form" action="">
            <div class="layui-form-item">
                <div class="layui-input-block">
                    <input type="text" name="Account" required  lay-verify="required" placeholder="请输入账户" class="layui-input">
                </div>
            </div>
            <div class="layui-form-item">
                <div class="layui-input-block">
                    <input type="password" name="Password" required lay-verify="required" placeholder="请输入密码" class="layui-input">
                </div>
            </div>
            <div class="layui-form-item">
                <div class="layui-input-block">
                    <button class="layui-btn login-btn" lay-submit lay-filter="login">立即登录</button>
                </div>
            </div>
        </form>
    </div>
</div>

<script src="/static/libs/jquery/jquery-3.5.1.min.js"></script>
<script src="/static/libs/layui-v2.5.6/layui/layui.all.js"></script>
<script type="text/javascript">
    layui.use('form', function(){
        const form = layui.form;
        //监听提交
        form.on('submit(login)', function(data){
            const index = layer.load();
            $.post("/doLogin", data.field, function(result){
                if(result.Code === 200){
                    layer.close(index);
                    window.location.href = '/';
                }else{
                    layer.close(index);
                    layer.msg(result.Msg);
                }
            });
            return false;
        });
    });
</script>
</body>
</html>
{{end}}