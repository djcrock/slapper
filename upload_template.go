package main

const uploadTemplate = `<!DOCTYPE html>
<html lang="en">
	<head>
		<meta charset="utf-8">
		<meta name=viewport content="width=device-width,initial-scale=1">
		<title>Slapper</title>
		<style>
			body {
				margin: 20px;
				padding: 0;
				color: white;
				background-color: #044BD9;
				text-align: center;
				font-family: "Lucida Console", Monaco, monospace;
			}

			main {
				max-width: 500px;
				margin: 0 auto;
				padding-bottom: 20px;
				background-color: white;
				box-shadow: 5px 10px 15px black;
				border-radius: 5px;
			}

			h1 {
				margin: 0;
				padding: 20px 5px;
				font-family:"ヒラギノ角ゴ Pro W3", "Hiragino Kaku Gothic Pro",Osaka, "メイリオ", Meiryo, "ＭＳ Ｐゴシック", "MS PGothic", sans-serif;
				color: #044BD9;
			}

			form {
				margin: 0 auto;
				display: flex;
				justify-content: space-evenly;
				flex-wrap: wrap;
			}

			input[type=file] {
				height: 0;
				width: 0;
				overflow: hidden;
				position: absolute;
				bottom: 0;
				left: 50%;
			}

			.btn {
				cursor: pointer;
				padding: 10px 20px;
				margin: 10px;
				font-size: 1.1em;
				border-radius: 5px;
				position: relative;
			}

			.btn:hover {
				top: 2px;
			}

			label.btn {
				background-color: #4D86F7;
				display: inline-block;
				box-shadow: 0 6px #1654A6;
			}

			label.btn:hover {
				box-shadow: 0 4px #1654A6;
			}

			button[type=submit] {
				border: 0;
				color: white;
				font-family: "Lucida Console", Monaco, monospace;
				background-color: #F2602B;
				box-shadow: 0 6px #C91905;
			}

			button[type=submit]:hover {
				box-shadow: 0 4px #C91905;
			}
		</style>
	</head>
	<body>
		<main>
			<h1>スラッパー</h1>
			<form action="/slap" method="POST" enctype="multipart/form-data">
				<label class="btn">
					pick<input type="file" name="site" required="required" accept=".zip">
				</label>
				<button type="submit" value="slap" class="btn">
					slap
				</button>
			</form>
		</main>
	</body>
</html>
`
