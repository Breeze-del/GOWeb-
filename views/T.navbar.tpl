{{define "navbar"}}
<div class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <a class="navbar-brand" href="/">我的博客</a>
        <div>
            <ul class="nav navbar-nav">
                <li {{if .IsHome}} class="active" {{end}}>
                    <a href="/">首页</a>
                </li>
                <li {{if .IScategory}} class="active" {{end}}>
                    <a href="/category">分类</a>
                </li>
                <li {{if .IStopic}} class="active" {{end}}>
                    <a href="/topic">文章</a>
                </li>
            </ul>
        </div>

        <div class="pull-right">
            <ul class="nav navbar-nav">
                {{if .IsLogin}}
                    {{/*点击退出时候，get访问主页并附带了exit=true，主页通过获取到这个值*/}}
                    {{/*来确定是不是退出状态    */}}
                    <li><a href="/login?exit=true">退出</a> </li>
                {{else}}
                    <li><a href="/login">管理员登陆</a></li>
                {{end}}
            </ul>
        </div>
    </div>
</div>
{{end}}