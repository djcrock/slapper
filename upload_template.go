package main

const uploadTemplate = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<title>Slapper</title>
	</head>
	<body>
		<h1>Slapper</h1>
		<form action="/slap" method="POST" enctype="multipart/form-data">
			<label>
				Zipped site:
				<input type="file" name="site" required="required">
			</label>
			<input type="submit" value="Publish">
		</form>
	</body>
</html>
`
