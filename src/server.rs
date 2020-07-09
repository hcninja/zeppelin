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

// Module for handling the web server related functions

use actix_web::{web, App, HttpResponse, HttpServer, Responder};
use actix_files as fs;

#[allow(unused_variables)]
#[actix_rt::main]
pub async fn start(host: String, port: String, secure: bool, path: String) -> std::io::Result<()> {
    let host_addr = format!("{}:{}", host, port);
    let html_post = format_template(&host_addr);

    HttpServer::new(|| {
        App::new()
            .route("/", web::get().to(index))
            .route("/upl", web::get().to(upload))
            .route("/cmd", web::get().to(cmd))
            .service(fs::Files::new("/nav", ".").show_files_listing())
    })
    .bind(host_addr)?
    .run()
    .await
}

async fn index() -> impl Responder {
    HttpResponse::Ok().body(r#"
    <p><a href="/upl">Upload</a></p>
    <p><a href="/nav">Navigate</a></p>
    <p><a href="/cmd">Command line</a></p>
    "#)
}

async fn upload() -> impl Responder {
    HttpResponse::Ok().body("Upload Template!")
}

async fn cmd() -> impl Responder {
    HttpResponse::Ok().body("Cmd client!")
}

fn format_template(host: &str) -> String{
    format!(r#"<html>
    <head>
           <title>Upload file</title>
    </head>
    <body>
    <form enctype="multipart/form-data" action="http://{}/upload" method="post">
        <input type="file" name="uploadfile" />
        <input type="submit" value="upload" />
    </form>
    </body>
    </html>"#, host)
}