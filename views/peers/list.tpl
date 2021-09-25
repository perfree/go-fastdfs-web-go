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
					<table class="layui-table" lay-data="{url:'/peers/page', page:true, id:'peersList',toolbar:'#toolbar'}" lay-filter="peersList">
						<thead>
						<tr>
							<th lay-data="{field:'Id', sort: true,align:'center'}">ID</th>
							<th lay-data="{field:'Name',align:'center'}">名称</th>
							<th lay-data="{field:'GroupName',align:'center'}">组名</th>
							<th lay-data="{field:'ServerAddress',align:'center'}">管理地址</th>
							<th lay-data="{field:'ShowAddress',align:'center'}">访问域名</th>
							<th lay-data="{field:'CreateTime',align:'center'}">添加时间</th>
							<th lay-data="{field:'right',toolbar: '#rightBar',width:150,align:'center'}">操作</th>
						</tr>
						</thead>
					</table>
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
		table.init('peersList', {
			limit: 15,
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