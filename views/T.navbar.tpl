{{define "navbar"}}
<div class="navbar navbar-default navbar-fixed-top">
    <div class="container">
        <a class="navbar-brand" href="/">我的博客</a>
        <ul class="nav navbar-nav">
            <li {{if .IsHome}} class="active" {{end}}>
                <a href="/">首页</a>
            </li>
            <li {{if .IsHome}} class="active" {{end}}>
                <a href="/category">分类</a>
            </li>
            <li {{if .IsHome}} class="active" {{end}}>
                <a href="/topic">文章</a>
            </li>
        </ul>
    </div>
</div>
{{end}}