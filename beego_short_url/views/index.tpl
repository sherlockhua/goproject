<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
</head>

<body>
<table border="1px">
 {{range .url_list}}
        <tr>
        <td> shorturl</td>
        <td> <a href={{.ShortUrl}}>{{.ShortUrl}}</a></td>
        </tr>
{{end}}
</table>
</body>
</html>
