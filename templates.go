/*
   Copyright 2020 - Jose Gonzalez Krause

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

var indexTemplate = `<!-- Index template -->
<html>
<head>
	<title>Zeppelin</title>
</head>

<body>
	<h1>Zeppelin index</h1>
	<ul>
		<li><a href="/upl">Upload</a></li>
		<li><a href="/nav/">Navigate</a></li>
		<li><a href="/cmd">Command line</a></li>
	</ul>
</body>
</html>
`

var uploadTemplate = `<!-- Upload form -->
<html>
<head>
	<title>Zeppelin</title>
</head>

<body>
	<h1>Zeppelin upload</h1>
	<form enctype="multipart/form-data" action="/upl" method="post">
		<input type="file" name="uploadfile" />
		<input type="submit" value="upload" />
	</form>
</body>
</html>
`

var uploadedTemplate = `<!-- Upload Ok -->
<html>
<head>
	<title>Zeppelin</title>
</head>

<body>
	<h1>Zeppelin upload</h1>
	<p>
		<h3>Uploaded '%s' to '%s'!</h1>
	</p>
	<ul>
		<li><a href="/upl">Upload</a></li>
		<li><a href="/nav/">Navigate</a></li>
		<li><a href="/cmd">Command line</a></li>
	</ul>
</body>
</html>
`
