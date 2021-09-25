<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <meta charset="UTF-8">
  <title>Go-FastDfs管理</title>
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, shrink-to-fit=no"/>
  <meta name="renderer" content="webkit"/>
  <meta name="force-rendering" content="webkit"/>
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
  <link rel="Shortcut Icon" href="/static/img/favicon.ico" type="image/x-icon"/>
  <link rel="stylesheet" href="/static/libs/layui-v2.5.6/layui/css/layui.css">
  <link href="/static/libs/font-awesome-4.7.0/css/font-awesome.min.css" rel="stylesheet"/>
  <link href="/static/css/style.css" rel="stylesheet"/>
</head>
<body class="layui-layout-body">
<div class="layui-layout layui-layout-admin">
  {{template "layout/header.tpl"}}
  {{template "layout/sider.tpl"}}
  <!-- 内容start -->
  <div class="layui-body">
    <div class="layui-tab" lay-allowClose="true" lay-filter="tabNav">
      <ul class="layui-tab-title content-tab-title">
        <li class="layui-this" lay-id="1"><i class='fa fa-home' style='font-size: 16px;'></i></li>
      </ul>
      <div class="layui-tab-content f-tab-content">
        <div class="layui-tab-item layui-show">
          <!-- 内容主体区域 -->
          <iframe src="/home" scrolling='auto' width='100%' height='100%' frameborder='0'
                  allowfullscreen='true' webkitallowfullscreen='true' mozallowfullscreen='true'
                  class='f-ifram'></iframe>
        </div>
      </div>
    </div>
  </div>
  <!-- 内容ned -->
</div>
<script src="/static/libs/jquery/jquery-3.5.1.min.js"></script>
<script src="/static/libs/layui-v2.5.6/layui/layui.all.js"></script>
<script src="/static/js/main.js"></script>
</body>
</html>