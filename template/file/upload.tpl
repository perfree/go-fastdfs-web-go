{{define "file/upload.tpl"}}
<!DOCTYPE html>
<html lang="zh_CN">
<head>
<meta charset="UTF-8">
<title>控制台</title>
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<link rel="stylesheet" href="/static/libs/layui-v2.5.6/layui/css/layui.css">
<link href="/static/libs/font-awesome-4.7.0/css/font-awesome.min.css" rel="stylesheet"/>
<style>
	.showUrl-href{color: #2F4056;font-size: 13px;text-decoration: underline;}
	.showUrl-href:hover{color:#01AAED;}
	.layui-container{padding: 0}
	#fileList{
		background-color: #627aad;
		height: 33px;
		line-height: 33px;
	}
	#fileListAction{
		height: 33px;
		line-height: 33px;
		background-color: #01AAED;
	}
</style>
</head>
<body>
<div class="layui-container" style="width: 100%;">
	<div class="layui-row">
		<div class="layui-col-md12" style="padding:5px;">
			<!-- 上传配置start -->
			<div class="layui-card">
				<div class="layui-card-header">上传配置</div>
				<div class="layui-card-body">
					<div class="layui-form-item">
						<label class="layui-form-label">场景</label>
						<div class="layui-input-block">
							<input id="scene" type="text" name="scene" value="default" placeholder="请输入场景" autocomplete="off" class="layui-input">
						</div>
					</div>
					<div class="layui-form-item">
						<label class="layui-form-label">路径</label>
						<div class="layui-input-block">
							<input id="path" type="text" name="path" value="" autocomplete="off" class="layui-input" placeholder="请输入上传路径,默认不填为default">
						</div>
					</div>
					<div class="layui-form-item">
						<label class="layui-form-label">回显url前缀</label>
						<div class="layui-input-block">
							<input id="showUrl" type="text" name="showUrl" autocomplete="off" class="layui-input" placeholder="默认为服务地址" value="{{.showAddress}}">
						</div>
					</div>
				</div>
			</div>
		</div>
		<div class="layui-col-md12" style="padding:5px;">
			<div class="layui-card">
				<div class="layui-card-body">
					<div class="layui-upload">
						<button type="button" class="layui-btn layui-btn-normal" id="fileList">选择文件</button>
						<button type="button" class="layui-btn" id="fileListAction">开始上传</button>
						<div class="layui-upload-list" style="overflow: auto;">
							<table class="layui-table">
								<thead>
								<tr>
									<th style="min-width: 100px">文件名</th>
									<th style="min-width: 60px">大小</th>
									<th style="min-width: 100px">状态</th>
									<th style="min-width: 200px">上传进度</th>
									<th style="min-width: 100px">操作</th>
								</tr>
								</thead>
								<tbody id="moreFileList">
								</tbody>
							</table>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</div>
<script src="/static/libs/jquery/jquery-3.5.1.min.js"></script>
<script src="/static/libs/layui-v2.5.6/layui/layui.js"></script>
<script src="/static/js/upload.js"></script>
</body>
</html>
{{end}}