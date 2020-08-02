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
use colored::*;

mod server;

// const VERSION: &str = "0.1.0";
// This takes directly the version number from cargo.toml
const VERSION: Option<&str> = option_env!("CARGO_PKG_VERSION");

#[allow(dead_code)]
fn main() {
    let host: String = String::from("127.0.0.1");
    let port: String = String::from("8080");
    let path: String = String::from("./");
    let tls: bool = false;

    println!("[=] Starting Zeppelin v{}", VERSION.unwrap_or("0.0.0_custom_build"));
    if tls {
        let addr = String::from(format!("https://{}:{}", host, port));
        println!("[*] Listening on: {}", addr.magenta().bold());
    } else {
        let addr = String::from(format!("http://{}:{}", host, port));
        println!("[*] Listening on: {}", addr.magenta().bold());
    }
    
    println!("[*] Serving path: {}", path.magenta().bold());

    std::env::set_var("RUST_LOG", "actix_server=debug,actix_web=debug");

    match server::start(host, port, tls, path) {
        Err(e) => println!("[!] Error: {},", e),
        _ => (),
    }
}
