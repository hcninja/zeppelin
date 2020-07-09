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
    HttpServer::new(|| {
        App::new()
            .service(fs::Files::new("/", ".").show_files_listing())
            .route("/upload", web::get().to(index))
    })
    .bind(format!("{}:{}", host, port))?
    .run()
    .await
}

async fn index() -> impl Responder {
    HttpResponse::Ok().body("Hello world!")
}
