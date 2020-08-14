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
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

// Server is the main server object
type Server struct {
	port     string
	host     string
	path     string
	safe     bool
	noUpload bool
	noCmd    bool
	tls      bool
}

// NewServer returns a new configured Server instance
func NewServer(host, port, path string, tls, safe, noUpload, noCmd bool) *Server {
	return &Server{
		host:     host,
		port:     port,
		path:     path,
		safe:     safe,
		tls:      tls,
		noCmd:    noCmd,
		noUpload: noUpload,
	}
}

// Run starts the server
func (s *Server) Run() error {
	errPrinter := color.New(color.FgHiRed)
	italics := color.New(color.Italic).SprintFunc()

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	if !s.safe {
		r.MaxMultipartMemory = 8 << 20 // 8MB
	} else {
		errPrinter.Println("- 8MB file upload limit DISABLED!")
	}

	r.Use(s.LoggerMW())
	r.Use(s.CustomizerMW())

	r.StaticFS("/nav", gin.Dir(s.path, true))
	r.GET("/", s.IndexGet)
	if !s.noUpload {
		r.GET("/upl", s.UploadGet)
		r.POST("/upl", s.UploadPost)
	}

	if !s.noCmd {
		r.GET("/cmd", s.CmdGet)
		r.POST("/cmd", s.CmdPost)
	}

	log.Println("- Log:    ", italics("date - src:port - code - path"))

	if s.tls {
		if err := r.RunTLS(s.host+":"+s.port, "./cert.pem", "key.pem"); err != nil {
			return err
		}
	} else {
		if err := r.Run(s.host + ":" + s.port); err != nil {
			return err
		}
	}
	return nil
}

// ============
// = Handlers =
// ============

// IndexGet returns the index page
func (s *Server) IndexGet(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.String(http.StatusOK, indexTemplate)
}

// UploadGet returns the upload page
func (s *Server) UploadGet(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.String(http.StatusOK, uploadTemplate)
}

// UploadPost handles the upload form
func (s *Server) UploadPost(c *gin.Context) {
	file, err := c.FormFile("uploadfile")
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: '%s'", err.Error()))
		return
	}
	log.Printf("- Uploaded file: '%s'", file.Filename)

	// FIXME: Sanitize path
	c.SaveUploadedFile(file, s.path+"/"+file.Filename)

	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.String(http.StatusOK, fmt.Sprintf(uploadedTemplate, file.Filename, path))
}

// CmdGet returns the cmd page
func (s *Server) CmdGet(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.String(http.StatusOK, cmdTemplate)
}

// CmdPost submits the command to execute and returns the result
func (s *Server) CmdPost(c *gin.Context) {
	var err error
	var out []byte

	cmd := c.PostForm("cmd")
	cmdTemplate := strings.Replace(cmdTemplate, "__CMD__", cmd, -1)

	if cmd == "" {
		out = []byte("Invalid command")
	} else {
		log.Printf("- cmd: %s", cmd)
		var os OS
		out, err = os.Exec(cmd)
		if err != nil {
			out = []byte(err.Error())
		}
	}

	cmdTemplate = strings.Replace(cmdTemplate, "__CODE__", string(out), -1)

	c.Header("Content-Type", "text/html; charset=UTF-8")
	c.String(http.StatusOK, cmdTemplate)
}

// =============
// Middlewares =
// =============

// LoggerMW is a custom logging middleware
func (s *Server) LoggerMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		path := c.Request.URL.Path
		src := c.Request.RemoteAddr

		if strings.Contains(path, "favicon") {
			return
		}

		// before request ^^^
		// c.Next()
		// after request vvv

		status := c.Writer.Status()

		log.Printf("> %s - %s - %d - %s",
			now.Format(time.RFC1123Z),
			src,
			status,
			path,
		)
	}
}

// CustomizerMW sets some custom data on each request
func (s *Server) CustomizerMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache")
		c.Header("Server", "Zeppelin v"+version)
	}
}
