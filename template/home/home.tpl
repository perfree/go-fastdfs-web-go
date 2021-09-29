{{define "home/home.tpl"}}
<!DOCTYPE html>
<html lang="en" >
<head>
    <meta charset="UTF-8">
    <title>控制台</title>
    <link rel="Shortcut Icon" href="/static/img/favicon.ico" type="image/x-icon"/>
    <link rel="stylesheet" href="/static/libs/layui-v2.5.6/layui/css/layui.css">
    <link href="/static/libs/font-awesome-4.7.0/css/font-awesome.min.css" rel="stylesheet"/>
    <link href="/static/css/home.css" rel="stylesheet"/>
</head>
<body>
<div class="layui-row">
    <div class="layui-col-md3" style="padding:5px;">
        <div class="layui-card">
            <div class="layui-card-header">文件总数<i class="fa fa-calculator" style="font-size: 16px;position: absolute;right: 15px;top: 10px;"></i></div>
            <div class="layui-card-body">
                <p class="page-view-totla" id="totalFileCount"></p>
            </div>
        </div>
    </div>
    <div class="layui-col-md3" style="padding:5px;">
        <div class="layui-card">
            <div class="layui-card-header">文件总大小<i class="fa fa-file-archive-o" style="font-size: 16px;position: absolute;right: 15px;top: 10px;"></i></div>
            <div class="layui-card-body">
                <p class="page-view-totla" id="totalFileSize"></p>
            </div>
        </div>
    </div>
    <div class="layui-col-md3" style="padding:5px;">
        <div class="layui-card">
            <div class="layui-card-header">文件总数(30天)<i class="fa fa-calculator" style="font-size: 16px;position: absolute;right: 15px;top: 10px;"></i></div>
            <div class="layui-card-body">
                <p class="page-view-totla" id="dayFileCount"></p>
            </div>
        </div>
    </div>
    <div class="layui-col-md3" style="padding:5px;">
        <div class="layui-card">
            <div class="layui-card-header">文件大小(30天)<i class="fa fa-file-archive-o" style="font-size: 16px;position: absolute;right: 15px;top: 10px;"></i></div>
            <div class="layui-card-body">
                <p class="page-view-totla" id="dayFileSize"></p>
            </div>
        </div>
    </div>

    <div class="layui-col-md3" style="padding:5px;">
        <div class="layui-card">
            <div class="layui-card-header">磁盘总空间<i class="fa fa-cubes" style="font-size: 16px;position: absolute;right: 15px;top: 10px;"></i></div>
            <div class="layui-card-body">
                <p class="page-view-totla" id="diskTotalSize"></p>
            </div>
        </div>
    </div>
    <div class="layui-col-md3" style="padding:5px;">
        <div class="layui-card">
            <div class="layui-card-header">磁盘剩余空间<i class="fa fa-cube" style="font-size: 16px;position: absolute;right: 15px;top: 10px;"></i></div>
            <div class="layui-card-body">
                <p class="page-view-totla" id="diskFreeSize"></p>
            </div>
        </div>
    </div>
    <div class="layui-col-md3" style="padding:5px;">
        <div class="layui-card">
            <div class="layui-card-header">索引节点总数(Linux)<i class="fa fa-sitemap" style="font-size: 16px;position: absolute;right: 15px;top: 10px;"></i></div>
            <div class="layui-card-body">
                <p class="page-view-totla" id="inodesTotal"></p>
            </div>
        </div>
    </div>
    <div class="layui-col-md3" style="padding:5px;">
        <div class="layui-card">
            <div class="layui-card-header">剩余索引节点(Linux)<i class="fa fa-sitemap" style="font-size: 16px;position: absolute;right: 15px;top: 10px;"></i></div>
            <div class="layui-card-body">
                <p class="page-view-totla" id="inodesFree"></p>
            </div>
        </div>
    </div>



    <div class="layui-col-md6" style="padding:5px;">
        <div class="layui-card" style="overflow: auto;">
            <div class="layui-card-header">快捷操作<i class="layui-icon" style="float: right;font-size: 20px;">&#xe614;</i></div>
            <div class="layui-card-body" style="float: left;padding: 5px 10px 5px 5px;height: 200px;">
                <a class="shortcut-button" href="javascript:;" id="fileUpload">
                    <span class="shortcut-button-icon"><i class="layui-icon" style="font-size: 25px">&#xe681;</i></span>
                    <span>文件上传</span>
                </a>
                <a class="shortcut-button" href="javascript:;" id="fileList">
                    <span class="shortcut-button-icon"><i class="layui-icon" style="font-size: 25px">&#xe621;</i></span>
                    <span>文件列表</span>
                </a>
                <a class="shortcut-button" href="javascript:;" id="switchPeers">
                    <span class="shortcut-button-icon"><i class="layui-icon" style="font-size: 25px">&#xe674;</i></span>
                    <span>切换集群</span>
                </a>
                <a class="shortcut-button" href="javascript:;" id="repair_stat">
                    <span class="shortcut-button-icon"><i class="layui-icon" style="font-size: 25px">&#xe62c;</i></span>
                    <span>修正统计信息</span>
                </a>
                <a class="shortcut-button" href="javascript:;" id="remove_empty_dir">
                    <span class="shortcut-button-icon"><i class="layui-icon" style="font-size: 25px">&#xe640;</i></span>
                    <span>删除空目录</span>
                </a>
                <a class="shortcut-button" href="javascript:;" id="backup">
                    <span class="shortcut-button-icon"><i class="layui-icon" style="font-size: 25px">&#xe656;</i></span>
                    <span>备份元数据</span>
                </a>
                <a class="shortcut-button" href="javascript:;" id="repair">
                    <span class="shortcut-button-icon"><i class="layui-icon" style="font-size: 25px">&#xe631;</i></span>
                    <span>同步失败修复</span>
                </a>
            </div>
        </div>
    </div>
    <div class="layui-col-md6" style="padding:5px;">
        <div class="layui-card" style="overflow: auto;">
            <div class="layui-card-header">版本信息<i class="layui-icon" style="float: right;font-size: 20px;">&#xe60b;</i></div>
            <div class="layui-card-body" style="height: 200px;padding: 5px 10px 5px 5px;">
                <table class="layui-table">
                    <colgroup>
                        <col width="200">
                        <col>
                    </colgroup>
                    <tbody>
                    <tr>
                        <td>当前版本</td>
                        <td><span >{{.version}}</span>&nbsp;&nbsp;&nbsp;<a href="https://github.com/perfree/go-fastdfs-web/releases" style="color: #009688" target="_blank">更新日志</a></td>
                    </tr>
                    <tr>
                        <td>发布日期</td>
                        <td>{{.versionDate}}</td>
                    </tr>
                    <tr>
                        <td>操作系统</td>
                        <td >{{.osName}}</td>
                    </tr>
                    <tr>
                        <td>系统架构</td>
                        <td >{{.osArch}}</td>
                    </tr>
                    </tbody>
                </table>
            </div>
        </div>
    </div>
    <div class="layui-col-md12" style="padding:5px;">
        <div class="layui-card" style="overflow: auto;">
            <div class="layui-card-body">
                <div id="main" style="height:600px;"></div>
            </div>
        </div>
    </div>
</div>
<script src="/static/libs/jquery/jquery-3.5.1.min.js"></script>
<script src="/static/libs/layui-v2.5.6/layui/layui.all.js"></script>
<script src="/static/libs/echarts/echarts.common.min.js"></script>
<script src="/static/js/home.js"></script>
</body>
</html>
{{end}}