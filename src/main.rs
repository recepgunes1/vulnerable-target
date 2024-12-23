use clap::{Arg, Command};
mod modules; // Include the modules folder

use crate::modules::docker::{start_service, inspect_service}; // Import Docker-related functions
use crate::modules::services::search_service; // Import search service function

#[tokio::main]
async fn main() {
    // banner 
    let banner = r#"
    @@@@@@@@&###&@@@@@@@
    @@@@@@@@&###&@@@@@@@
    @@@@@@@#^...7@@@@@@@
    @@@@@@@&:   !@@@@@@@
    @@@@@@@@BGGGB?777P@@
    @@@@@@@@@@@@#    !@@
    @@&BBBB&BBBB#!^^^5@@
    @@J   .?.   ?:.. ?@@
    @@Y ..:Y....J:...?@@
    @@&####&####&####&@@
    @@ Happy  Hacking @@
    "#;
    
    println!("{}", banner);
    let matches = Command::new("vulnerability target")
        .bin_name("vt")
        .about("CLI to manage Docker-based vulnerable labs")
        .subcommand(
            Command::new("start")
                .about("Start services")
                .arg(Arg::new("service").help("Service name or 'all'").required(true)),
        )
        .subcommand(
            Command::new("search")
                .about("Search services")
                .arg(Arg::new("keyword").help("Keyword to search").required(true)),
        )
        .subcommand(
            Command::new("inspect")
                .about("Inspect running service")
                .arg(Arg::new("service").help("Service name").required(true)),
        )
        .get_matches();

    match matches.subcommand() {
        Some(("start", sub_matches)) => {
            let service = sub_matches.get_one::<String>("service").unwrap();
            start_service(service).await;
        }
        Some(("search", sub_matches)) => {
            let keyword = sub_matches.get_one::<String>("keyword").unwrap();
            search_service(keyword);
        }
        Some(("inspect", sub_matches)) => {
            let service = sub_matches.get_one::<String>("service").unwrap();
            inspect_service(service).await;
        }
        _ => eprintln!("Invalid command. Use 'vt --help' for usage."),
    }
}
