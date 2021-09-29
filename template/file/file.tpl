{{define "file/file.tpl"}}
<!DOCTYPE html>
<html lang="zh_CN">
<head>
    <meta charset="UTF-8">
    <title>文件列表</title>
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
    <link rel="stylesheet" href="/static/libs/layui-v2.5.6/layui/css/layui.css">
    <link href="/static/libs/font-awesome-4.7.0/css/font-awesome.min.css" rel="stylesheet"/>
    <link rel="stylesheet" href="/static/css/file.css" media="all">
    <script type="text/javascript" src="/static/libs/iconfont/iconfont.js"></script>
</head>
<body>
<div class="layui-container" style="width: 100%;">
    <div class="layui-row">
        <div class="layui-col-md12" style="padding:5px;">
            <div class="layui-card">
                <div class="layui-card-header" style="padding: 5px;">
                    <button class="layui-btn" id="fileUpload">
                        <i class="fa fa-cloud-upload"></i>&nbsp;上传
                    </button>
                </div>
                <div class="layui-card-body">
                    <div class="file-list-side">
						<span class="layui-breadcrumb" id="path-side">

                        </span>
                    </div>
                    <hr class="layui-bg-gray">
                    <div class="file-list">
                        <div class="file-list-header">
                            <ul class="file-list-header-ul">
                                <li class="file-list-header-li file-list-name">文件名</li>
                                <li class="file-list-header-li file-list-size">大小</li>
                                <li class="file-list-header-li file-list-date">创建日期</li>
                                <li class="file-list-header-li file-list-handle">操作</li>
                            </ul>
                        </div>
                        <div class="file-list-result" id="file-result">

                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>

<script src="/static/libs/jquery/jquery-3.5.1.min.js"></script>
<script src="/static/libs/layui-v2.5.6/layui/layui.all.js"></script>
<script src="/static/libs/template-web/template-web.js"></script>
<script src="/static/js/kit.js"></script>
<script src="/static/js/file.js"></script>
</body>
</html>
{{end}}