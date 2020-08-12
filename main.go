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

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

var version = "0.1.1"
var path = "./"

func main() {
	gin.SetMode(gin.ReleaseMode)
	log.SetFlags(0)

	hostFlg := flag.String("host", "127.0.0.1", "Server host")
	portFlg := flag.String("port", "8080", "Server port")
	pathFlg := flag.String("path", "./", "Path to serve and upload files to")
	tlsFlag := flag.Bool("tls", false, "Enables TLS. Cert and key must be in root 'cert.pem and key.pem'")
	safeFlg := flag.Bool("unsafe", false, "Removes the file upload limit of 8MB")
	noupFlg := flag.Bool("noupload", false, "Disables the upload endpoint")
	flag.Parse()

	errPrinter := color.New(color.FgHiRed)
	italics := color.New(color.Italic).SprintFunc()
	remarkText := color.New(color.FgMagenta, color.Bold).SprintFunc()

	log.Printf("Starting Zeppelin v%s\n", version)

	var schema string = "http://"
	if *tlsFlag {
		schema = "https://"
	}

	path = *pathFlg

	log.Printf("- Server:  %s\n", remarkText(schema, *hostFlg, ":", *portFlg))
	log.Printf("- Serving: %s\n", remarkText(*pathFlg))

	r := gin.New()
	if !*safeFlg {
		r.MaxMultipartMemory = 8 << 20 // 8MB
	} else {
		errPrinter.Println("- 8MB file upload limit DISABLED!")
	}

	r.Use(LoggerMW())
	r.Use(CustomizerMW())

	r.GET("/", IndexGet)
	if !*noupFlg {
		r.GET("/upl", UploadGet)
		r.POST("/upl", UploadPost)
	}
	r.GET("/cmd", CmdGet)
	r.StaticFS("/nav", gin.Dir(*pathFlg, true))

	log.Println("- Log:    ", italics("date - src:port - code - path"))

	if *tlsFlag {
		if err := r.RunTLS(*hostFlg+":"+*portFlg, "./cert.pem", "key.pem"); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := r.Run(*hostFlg + ":" + *portFlg); err != nil {
			log.Fatal(err)
		}
	}
}

// IndexGet returns the index page
func IndexGet(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.String(http.StatusOK, indexTemplate)
}

// UploadGet returns the upload page
func UploadGet(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.String(http.StatusOK, uploadTemplate)
}

// CmdGet returns the cmd page
func CmdGet(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.String(http.StatusOK, "NYI")
}

// UploadPost handles the upload form
func UploadPost(c *gin.Context) {
	file, err := c.FormFile("uploadfile")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: '%s'", err.Error()))
		return
	}
	log.Printf("- Uploaded file: '%s'", file.Filename)

	// FIXME: Sanitize path
	c.SaveUploadedFile(file, path+"/"+file.Filename)

	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.String(http.StatusOK, fmt.Sprintf(uploadedTemplate, file.Filename, path))
}

// =============
// Middlewares =
// =============

// LoggerMW is a custom logging middleware
func LoggerMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		path := c.Request.URL.Path
		src := c.Request.RemoteAddr

		if strings.Contains(path, "favicon") {
			return
		}

		// before request ^^^
		c.Next()
		// after request vvv

		status := c.Writer.Status()

		log.Printf("> %s - %s - %d - %s",
			now.Format(time.RFC1123Z),
			// delta,
			src,
			status,
			path,
		)
	}
}

// CustomizerMW sets some custom data on each request
func CustomizerMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache")
		c.Header("Server", "Zeppelin v"+version)
	}
}

// ===========
// = Helpers =
// ===========

func nyi(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, gin.H{
		"status":  http.StatusNotImplemented,
		"message": "NYI",
	})

	return
}

// =============
// = Templates =
// =============

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
