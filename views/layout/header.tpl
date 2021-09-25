<div class="layui-header f-header" style="position: relative;" >
    <ul class="layui-nav layui-layout-left left-nav f-nav">
        <li class="layui-nav-item" lay-unselect>
            <a href="javascript:;" title="收缩菜单" class="f-switch-side-btn f-switch-btn-on">
                <i class="layui-icon layui-icon-shrink-right"></i>
            </a>
            <a href="javascript:;" title="展开菜单" class="f-switch-side-btn f-switch-btn-off" style="display: none;">
                <i class="layui-icon layui-icon-spread-left"></i>
            </a>
        </li>
        <li class="layui-nav-item" lay-unselect>
            <a href="javascript:;" class="f-refresh-btn" title="刷新">
                <i class="layui-icon layui-icon-refresh"></i>
            </a>
        </li>

        <li class="layui-nav-item" lay-unselect>
            <a href="javascript:;" class="f-support-btn" title="支持">
                <i class="layui-icon layui-icon-rmb"></i>
            </a>
        </li>
    </ul>

    <ul class="layui-nav layui-layout-right header-right f-nav">
        <li class="layui-nav-item" lay-unselect>
            <a href="javascript:;" class="f-screen-full-btn" title="全屏/退出全屏">
                <i class="layui-icon layui-icon-screen-full f-screen-full-btn-icon"></i>
                <i class="layui-icon layui-icon-screen-restore f-exit-full-btn-icon"></i>
            </a>
        </li>
        <li class="layui-nav-item" lay-unselect>
            <a href="javascript:;">
                 <img src="/static/img/user.png" class="layui-nav-img">
                 <span text="user.name"></span>
            </a>
            <dl class="layui-nav-child">
                <dd lay-unselect>
                    <a href="/logout" class="header-right-down-item">退出登录</a>
                </dd>
            </dl>
        </li>
    </ul>
</div>
