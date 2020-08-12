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
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

var (
	version   = "dev_build"
	buildTime = "N/A"
	path      = "./"
)

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

	remarkText := color.New(color.FgMagenta, color.Bold).SprintFunc()

	log.Printf("Starting Zeppelin %s (%s)\n", version, buildTime)

	var schema string = "http://"
	if *tlsFlag {
		schema = "https://"
	}

	path = *pathFlg

	log.Printf("- Server:  %s\n", remarkText(schema, *hostFlg, ":", *portFlg))
	log.Printf("- Serving: %s\n", remarkText(*pathFlg))

	server := NewServer(*hostFlg, *portFlg, *pathFlg, *tlsFlag, *safeFlg, *noupFlg)
	if err := server.Run(); err != nil {
		log.Fatal(err)
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
