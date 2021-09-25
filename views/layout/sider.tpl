<div class="layui-side f-side">
    <div class="layui-logo f-logo-text">
        Go-FastDfs 管理
    </div>
    <div class="layui-logo f-logo-img">
         <img src="/static/img/logo.png">
    </div>
    <div class="layui-side-scroll">
        <!-- 左侧导航区域（可配合layui已有的垂直导航） -->
        <ul class="layui-nav layui-nav-tree f-side-nav" lay-filter="side">
            <li class="layui-nav-item f-nav-item">
                <a href="javascript:;" onclick="clickMenu(this)" class="p-menu-item"
                   data-url="/home" id="1" data-icon="fa-home" data-name="控制台">
                    <i class="fa fa-home" aria-hidden="true"></i>
                    <span class="f-nav-content">控制台</span>
                </a>
            </li>
            <li class="layui-nav-item f-nav-item">
                <a class="" href="javascript:;">
                    <i class="fa fa-files-o" aria-hidden="true"></i>
                    <span class="f-nav-content">文件管理</span>
                </a>
                <dl class="layui-nav-child">
                    <dd class="f-child-side">
                        <a href="javascript:;" class="p-menu-item"
                           data-url="/file/upload" id="2" data-icon="fa-cloud-upload"
                           data-name="文件上传"
                           onclick="clickMenu(this);">
                            文件上传
                        </a>
                    </dd>
                </dl>
                <dl class="layui-nav-child">
                    <dd class="f-child-side">
                        <a href="javascript:;" class="p-menu-item"
                           data-url="/file" id="3" data-icon="fa-file-o"
                           data-name="文件列表"
                           onclick="clickMenu(this);">
                            文件列表
                        </a>
                    </dd>
                </dl>
            </li>
            <li class="layui-nav-item f-nav-item">
                <a class="" href="javascript:;">
                    <i class="fa fa-server" aria-hidden="true"></i>
                    <span class="f-nav-content">集群管理</span>
                </a>
                <dl class="layui-nav-child">
                    <dd class="f-child-side">
                        <a href="javascript:;" class="p-menu-item"
                           data-url="/peers" id="4" data-icon="fa-server"
                           data-name="集群列表"
                           onclick="clickMenu(this);">
                           集群列表
                        </a>
                    </dd>
                </dl>
            </li>
            <li class="layui-nav-item f-nav-item">
                <a href="https://sjqzhang.github.io/go-fastdfs/#character" class="p-menu-item"
                   data-url="/settings/user" id="6" data-icon="fa-book" data-name="go-fastdfs文档" target="_blank">
                    <i class="fa fa-book" aria-hidden="true"></i>
                    <span class="f-nav-content">go-fastdfs文档</span>
                </a>
            </li>

            <li class="layui-nav-item f-nav-item">
                <a href="javascript:;" onclick="clickMenu(this)" class="p-menu-item"
                   data-url="/settings/user" id="5" data-icon="fa-sliders" data-name="系统设置">
                    <i class="fa fa-sliders" aria-hidden="true"></i>
                    <span class="f-nav-content">系统设置</span>
                </a>
            </li>
        </ul>
    </div>
</div>
