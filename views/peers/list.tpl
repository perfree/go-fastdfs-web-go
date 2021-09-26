<!DOCTYPE html>
<html lang="zh_CN" >
<head>
<meta charset="UTF-8">
<title>控制台</title>
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<link rel="stylesheet" href="/static/libs/layui-v2.5.6/layui/css/layui.css">
<link href="/static/libs/font-awesome-4.7.0/css/font-awesome.min.css" rel="stylesheet"/>
<style>
	.addPeers{
		background-color: #627aad;
	}
	.layui-laypage .layui-laypage-curr .layui-laypage-em {
		background-color: #627aad;
	}
</style>
</head>
<body>
<div class="layui-container" style="width: 100%;padding: 0">
	<div class="layui-row">
		<div class="layui-col-md12" style="padding:5px;">
			<div class="layui-card">
				<div class="layui-card-body">
					<table id="peersList" lay-filter="peersList"></table>
				</div>
			</div>
		</div>
	</div>
</div>
<script type="text/html" id="toolbar">
	<div class="layui-btn-container">
		<button class="layui-btn layui-btn-sm addPeers" lay-event="addPeers">添加集群</button>
	</div>
</script>
<script type="text/html" id="rightBar">
	<a class="layui-btn layui-btn-xs" lay-event="edit">编辑</a>
	<a class="layui-btn layui-btn-danger layui-btn-xs" lay-event="del">删除</a>
</script>
<script src="/static/libs/jquery/jquery-3.5.1.min.js"></script>
<script src="/static/libs/layui-v2.5.6/layui/layui.js"></script>
<script>
	layui.use(['table','laytpl'], function() {
		var table = layui.table;
		//第一个实例
		table.render({
			elem: '#peersList',
			url: '/peers/page',
			page: true,
			toolbar:'#toolbar',
			cols: [[
				{field: 'Id', title: 'ID', align:'center', sort: true, fixed: 'left'},
				{field: 'Name', title: '名称', align:'center'},
				{field: 'GroupName', title: '组名', align:'center'},
				{field: 'ServerAddress', title: '管理地址', align:'center'},
				{field: 'ShowAddress', title: '访问域名',align:'center'},
				{field: 'CreateTime', title: '添加时间', align:'center', templet : function (res) {
						return layui.util.toDateString(res.CreateTime, 'yyyy年MM月dd日 HH:mm:ss')
				}},
				{field: 'right', title: '操作', width: 150, align:'center', toolbar: '#rightBar'}
			]],
			response: {
				statusName: 'State',
				statusCode: 200,
				msgName: 'Msg',
				countName: 'Total',
				dataName: 'Data'
			}
		});

		//监听行工具事件
		table.on('tool(peersList)', function(obj){
			var data = obj.data;
			if(obj.event === 'del'){
				layer.confirm('确定要删除该集群吗?', function(index){
					$.post("/peers/del",{"id":obj.data.id},function (data) {
						if(data.code === 200){
							obj.del();
							layer.msg("删除成功");
						}else{
							layer.msg(data.msg);
						}
					})
					layer.close(index);
				});
			} else if(obj.event === 'edit'){
				layer.open({
					type: 2,
					area: ['650px','340px'],
					title :'编辑集群',
					shadeClose: true,
					maxmin: true,
					content: '/peers/edit?id='+obj.data.id
				});
			}
		});

		//头工具栏事件
		table.on('toolbar(peersList)', function(obj){
			switch(obj.event){
				case 'addPeers':
					layer.open({
						type: 2,
						area: ['650px','340px'],
						title :'添加集群',
						shadeClose: true,
						maxmin: true,
						content: '/peers/add'
					});
					break;
			};
		});
	})
</script>
</body>
</html>