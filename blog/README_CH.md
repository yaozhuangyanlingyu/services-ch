＃ 博客
这是编写多服务博客应用程序的完整的端到端示例

＃＃ 用法
在开发人员文档中查看[博客教程]（https://m3o.dev/tutorials/building-a-blog）。

＃＃ 怎么运行的
＃＃＃ 当下
博客服务的设计使用户可以将其部署到自己的微型名称空间，并使用诸如以下命令的微型帐户写入内容

```sh
micro posts save --id=7 --tags=News,Finance --title="Breaking News" --content="The stock market has just crashed"
```

并通过使用API​​在前端显示内容：

```sh
curl -H "Authorization: Bearer $MICRO_API_TOKEN" "Micro-Namespace: $NAMESPACE" https://api.m3o.com/tags/list
{
	"tags": [
		{
			"type": "post-tag",
			"slug": "news",
			"title": "News",
			"count": "3"
		}
    ]
]
```

目前还没有提供评论，只有帖子和标签。
访问受身份验证规则（即）约束。帖子列表，标签列表已打开，帖子保存需要微型登录。

###未来的可能性

####允许非Micro用户撰写帖子，评论
如果我们提供用户/登录服务（与auth明显不同，它可以是基于会话的简单身份验证）以允许非Micro用户注册，则可以执行以下操作：
-用户（从这一点起，我们将其称为用户Alice）在其名称空间中启动帖子，标签和登录服务。
-爱丽丝打开了端点
-在Netlify或Github Pages上托管JS和HTML的人们（我们称它们为Yoga Pants Co和Drone Inc）可以在Alice托管的服务中创建帐户。这样，通过拥有一个微型帐户，爱丽丝成为了无头的CMS提供者。可以在Alice的服务实例之上创建多个博客。

问题：
-Yoga Pants Co或Drone Inc如何向Alice或M3O支付后端托管费用？