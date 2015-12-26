<!DOCTYPE html>
<html>
<head>
<title>Template</title>
</head>
<body>
<h1>Header</h1>
<h1>{{.Path}}</h1>
{{range .Items}}
<a href="{{.URL}}">{{.Name}}</a><br>
{{end}}
</body>
</html>
