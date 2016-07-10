{{define "navbar"}}
<a class="navbar-brand" href="/">LogServer</a>
<div>
	<ul class="nav navbar-nav">
		<li {{if .IsGoCTLog}}class="active"{{end}}><a href="/goctlog">GoCT日志</a></li>
		<li {{if .IsMSLog}}class="active"{{end}}><a href="/mslog">微服务日志</a></li>
		<li {{if .IsAllLog}}class="active"{{end}}><a href="/alllog">所有日志</a></li>
	</ul>
</div>

<form class="navbar-form navbar-right" role="search">
  <div class="form-group">
    <input id="sendbox" type="text" class="form-control" placeholder="过滤">
  </div>
  <button id="sendbtn" type="button" class="btn btn-default">过滤</button>
</form>

{{end}}