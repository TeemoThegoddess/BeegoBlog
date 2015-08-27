<nav class="navbar navbar-inverse navbar-fixed-top">
	<div class="container">
		<a class="navbar-brand logo" href="/">Blog</a>
		 	<ul class="nav navbar-nav">
		 		<li {{if .isHome}}class="active"{{end}}><a href="/">Home<span class="sr-only"></span></a></li>
		 		<li {{if .isCategory}}class="active"{{end}}><a href="/category">Category</a></li>
		 		<li {{if .isTopic}}class="active"{{end}}><a href="/topic">Topic</a></li>
		 	</ul>
		 	<ul class="nav navbar-nav navbar-right ">
		 			{{if .isLogin}}
		 				<li><a class="link" href="/login?exit=true">
		 					Logout
		 				</a></li>
		 			{{else}}
		 				<li><a class="link" href="/login">
		 					Login
		 				</a></li>
		 			{{end}}
			</ul>
	</div>
</nav>