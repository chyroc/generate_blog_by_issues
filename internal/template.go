package internal

const headerTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<!-- Begin Jekyll SEO tag v2.3.0 -->
<title>{{ .Title }}</title>
<meta property="og:title" content="{{ .Host }}" />
<meta property="og:locale" content="en_US" />
<meta name="description" content="{{ .Name }}" />
<meta property="og:description" content="{{ .Name }}" />
<link rel="canonical" href="{{ .Host }}" />
<meta property="og:url" content="{{ .Host }}" />
<meta property="og:site_name" content="{{ .Host }}" />
<script type="application/ld+json">
{"name":"{{ .Host }}","description":"{{ .Name }}","author":"{{ .Author }}","@type":"WebSite","url":"{{ .Host }}","image":null,"publisher":null,"headline":"{{ .Host }}","dateModified":null,"datePublished":null,"sameAs":null,"mainEntityOfPage":null,"@context":"http://schema.org"}</script>
<!-- End Jekyll SEO tag -->
`

const articleTmpl = headerTmpl + `
<link href="http://{{ .Host }}/assets/css/style.css?v=305ca492b034089b2a2287dae4e9fa13ac15b666" rel="stylesheet">
</head>
<body>
<div class="container-lg px-3 my-5 markdown-body">

{{ .Content }}

<script src="http://{{ .Host }}/assets/javascript/anchor-js/anchor.min.js"></script>
<script>anchors.add();</script>

<div id="disqus_thread"></div>
<script>
(function () { // DON'T EDIT BELOW THIS LINE
var d = document, s = d.createElement('script')
s.src = 'https://chyroc.disqus.com/embed.js'
s.setAttribute('data-timestamp', +new Date());
(d.head || d.body).appendChild(s)
})()
</script>
<noscript>Please enable JavaScript to view the <a href="https://disqus.com/?ref_noscript">comments powered by
Disqus.</a></noscript>
<script id="dsq-count-scr" src="//chyroc.disqus.com/count.js" async></script>

</div>

</body>
</html>
`

const readmeTmpl = headerTmpl + `
<style>
.tab {
    overflow: hidden;
    background-color: #f1f1f1;
}
.tab button {
    background-color: inherit;
    float: left;
    border: none;
    outline: none;
    cursor: pointer;
    padding: 14px 16px;
    transition: 0.3s;
}
.tab button:hover {
    background-color: #ddd;
}
.tab button.active {
    background-color: #ccc;
}
.tabcontent {
    display: none;
    padding: 6px 12px;
    border-top: none;
}
</style>
<script>
function openCity(evt, cityName) {
    var i, tabcontent, tablinks;
    tabcontent = document.getElementsByClassName("tabcontent");
    for (i = 0; i < tabcontent.length; i++) {
        tabcontent[i].style.display = "none";
    }
    tablinks = document.getElementsByClassName("tablinks");
    for (i = 0; i < tablinks.length; i++) {
        tablinks[i].className = tablinks[i].className.replace(" active", "");
    }
    document.getElementById(cityName).style.display = "block";
    evt.currentTarget.className += " active";
}
</script>
<link href="http://{{ .Host }}/assets/css/style.css?v=305ca492b034089b2a2287dae4e9fa13ac15b666" rel="stylesheet">
</head>
<body>
<div class="container-lg px-3 my-5 markdown-body">

	<div class="tab">
	  <button class="tablinks" onclick="openCity(event, 'article')" id="defaultOpen">文章</button>
	  <button class="tablinks" onclick="openCity(event, 'blogroll')">链接</button>
	</div>

	<div id="article" class="tabcontent">
		{{ .Content }}
	</div>

	<div id="blogroll" class="tabcontent">
		{{ .Blogroll }}
	</div>
</div>
<script>document.getElementById("defaultOpen").click();</script>

<script src="http://{{ .Host }}/assets/javascript/anchor-js/anchor.min.js"></script>
<script>anchors.add();</script>

</body>
</html>
`
