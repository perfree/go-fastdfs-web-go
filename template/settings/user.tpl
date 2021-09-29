{{define "settings/user.tpl"}}
<!DOCTYPE html>
<html lang="zh_CN">
<head>
<meta charset="UTF-8">
<title>个人资料</title>
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<link rel="stylesheet" href="/static/libs/layui-v2.5.6/layui/css/layui.css">
<link href="/static/libs/font-awesome-4.7.0/css/font-awesome.min.css" rel="stylesheet"/>
<style>
	.layui-form-label{
		width: 120px;
		color: #657180;
	}
	.editUser {
		height: 33px;
		line-height: 33px;
		background-color: #627aad;
	}
	.reset{
		height: 33px;
		line-height: 33px;
	}
</style>
</head>
<body>
<div class="layui-container" style="width: 100%;padding: 0">
	<div class="layui-row">
		<div class="layui-col-md12" style="padding:5px;">
			<div class="layui-card">
				<div class="layui-card-header">个人资料</div>
				<div class="layui-card-body">
					<form class="layui-form" action="">
						<input type="hidden" name="Id" value="{{.user.Id}}">
						<div class="layui-form-item">
							<label class="layui-form-label">账户:</label>
							<div class="layui-input-inline">
								<input type="text" name="Account" required  lay-verify="required" placeholder="请输入账户" autocomplete="off" class="layui-input" value="{{.user.Account}}" readonly>
							</div>
							<div class="layui-form-mid layui-word-aux">不可修改。用于登入名</div>
						</div>
						<div class="layui-form-item">
							<label class="layui-form-label">昵称:</label>
							<div class="layui-input-inline">
								<input type="text" name="Name"  required  lay-verify="required" placeholder="请输入昵称" autocomplete="off" class="layui-input" value="{{.user.Name}}">
							</div>
							<div class="layui-form-mid layui-word-aux">16位字符以内</div>
						</div>
						<div class="layui-form-item">
							<label class="layui-form-label">邮箱:</label>
							<div class="layui-input-inline">
								<input type="text" name="Email"  required  lay-verify="email" placeholder="请输入邮箱" autocomplete="off" class="layui-input" value="{{.user.Email}}">
							</div>
						</div>
						<div class="layui-form-item">
							<label class="layui-form-label">原密码:</label>
							<div class="layui-input-inline">
								<input type="password" name="Password"  required  lay-verify="required" placeholder="请输入原密码" autocomplete="off" class="layui-input">
							</div>
							<div class="layui-form-mid layui-word-aux">6-16位,不可包含特殊字符</div>
						</div>
						<div class="layui-form-item">
							<label class="layui-form-label">新密码:</label>
							<div class="layui-input-inline">
								<input type="password" name="NewPassword"  required  lay-verify="required" placeholder="请输入新密码" autocomplete="off" class="layui-input">
							</div>
							<div class="layui-form-mid layui-word-aux">6-16位,不可包含特殊字符</div>
						</div>
						<div class="layui-form-item">
							<div class="layui-input-block">
								<button class="layui-btn editUser" lay-submit="" lay-filter="editUser">立即提交</button>
								<button type="reset" class="reset layui-btn layui-btn-primary">重置</button>
							</div>
						</div>
					</form>
				</div>
			</div>
		</div>
	</div>
</div>
<script src="/static/libs/jquery/jquery-3.5.1.min.js"></script>
<script src="/static/libs/layui-v2.5.6/layui/layui.js"></script>
<script>
	layui.use('form', function() {
		var form = layui.form;
		form.on('submit(editUser)', function(data){
			$.post("/settings/editUser",data.field,function (data) {
				if(data.Code === 200){
					layer.msg("修改成功", {icon: 6});
					setTimeout(function(){
						window.location.reload();
					}, 1000);
				}else{
					layer.msg(data.Msg);
				}
			})
			return false;
		});
	})
</script>
</body>
</html>
{{end}}