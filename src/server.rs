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

// use std::future::Future;
use actix_multipart::Multipart;
use actix_web::{web, App, HttpResponse, HttpServer, Responder, Error};
use actix_files as fs;

#[allow(unused_variables)]
#[actix_rt::main]
pub async fn start(host: String, port: String, secure: bool, path: String) -> std::io::Result<()> {
    let host_addr = format!("{}:{}", host, port);
    let html_post = format_template(&host_addr);

    HttpServer::new(|| {
        App::new()
            .route("/", web::get().to(index))
            .route("/upl", web::get().to(get_upload))
            .route("/upl", web::post().to(post_upload))
            .route("/cmd", web::get().to(cmd))
            .service(fs::Files::new("/nav", ".").show_files_listing())
    })
    .bind(host_addr)?
    .run()
    .await
}

async fn index() -> impl Responder {
    HttpResponse::Ok().body(r#"
    <html>
    <h1>Zeppelin index</h1>
    <ul>
    <li><a href="/upl">Upload</a></li>
    <li><a href="/nav">Navigate</a></li>
    <li><a href="/cmd">Command line</a></li>
    </ul>
    </html>
    "#)
}

async fn get_upload() -> impl Responder {
    let html_post = format_template("localhost:8080");
    HttpResponse::Ok().body(html_post)
}

async fn post_upload(mut _payload: Multipart) -> Result<HttpResponse, Error> {
    // iterate over multipart stream
    // while let Ok(Some(mut field)) = payload.try_next().await {
    //     let content_type = field.content_disposition().unwrap();
    //     let filename = content_type.get_filename().unwrap();
    //     let filepath = format!("./tmp/{}", sanitize_filename::sanitize(&filename));
       
    //     // File::create is blocking operation, use threadpool
    //     let mut f = web::block(|| std::fs::File::create(filepath))
    //         .await
    //         .unwrap();
      
    //         // Field in turn is stream of *Bytes* object
    //     while let Some(chunk) = field.next().await {
    //         let data = chunk.unwrap();
         
    //         // filesystem operations are blocking, we have to use threadpool
    //         f = web::block(move || f.write_all(&data).map(|_| f)).await?;
    //     }
    // }

    Ok(HttpResponse::Ok().into())
}

async fn cmd() -> impl Responder {
    HttpResponse::Ok().body("Cmd client!")
}

fn format_template(host: &str) -> String{
    format!(r#"<!-- Upload form -->
    <html>
    <head>
    <title>Upload file</title>
    </head>
    <body>
    <h1>Zeppelin upload</h1>
    <form enctype="multipart/form-data" action="http://{}/upload" method="post">
        <input type="file" name="uploadfile" />
        <input type="submit" value="upload" />
    </form>
    </body>
    </html>"#, host)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn templating() {
        let host = "192.168.1.1:8080";
        let result = format_template(host);
        assert!(result.contains(host));
    }
}