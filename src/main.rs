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
use clap::{App, Arg};
use colored::*;

mod server;

// const VERSION: &str = "0.1.0";
// This takes directly the version number from cargo.toml
const VERSION: Option<&str> = option_env!("CARGO_PKG_VERSION");

fn main() {
    let app_version = VERSION.unwrap_or("0.0.0_custom_build");

    let matches = App::new("Zeppelin")
        .version(app_version)
        .author("Jose Gonzalez-Krause <contact@hackercat.ninja>")
        .about("Tiny file-server on steroids for pentesting.")
        .arg(Arg::with_name("HOST")
            .short("h")
            .long("host")
            .takes_value(true)
            .multiple(false)
            .help("Host address for the file-server [127.0.0.1]"))
        .arg(Arg::with_name("PORT")
            .short("p")
            .long("port")
            .takes_value(true)
            .multiple(false)
            .help("Port for the file-server [8080]"))
        .arg(Arg::with_name("PATH")
            .short("d")
            .long("dir")
            .help("Directory to serve [./]")
            .takes_value(true))
        .get_matches();

    let host: String = String::from(matches.value_of("HOST").unwrap_or("127.0.0.1"));
    let port: String = String::from(matches.value_of("PORT").unwrap_or("8080"));
    let path: String = String::from(matches.value_of("PATH").unwrap_or("./"));
    let tls: bool = false;

    println!("Starting Zeppelin v{}", app_version);

    if tls {
        let addr = String::from(format!("https://{}:{}", host, port));
        println!("Listening on: {}", addr.magenta().bold());
    } else {
        let addr = String::from(format!("http://{}:{}", host, port));
        println!("Listening on: {}", addr.magenta().bold());
    }
    
    println!("Serving path: {}", path.magenta().bold());

    std::env::set_var("RUST_LOG", "actix_server=debug,actix_web=debug");
    
    match server::start(host, port, tls, path) {
        Err(e) => println!("Error: {}", format!("{}", e).red().bold()),
        _ => (),
    }
}
