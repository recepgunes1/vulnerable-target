use subprocess::Exec;
use colored::*;
use std::fmt;
use serde_json::Value;


#[derive(Debug)]
struct ServiceInfo {
    ip_address: String,
    web_port: Option<String>,
}

impl fmt::Display for ServiceInfo {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            f,
            "{}\n{}",
            format!("- IP Address: {}", self.ip_address).blue(),
            format!("- Web Port: {}", self.web_port.as_deref().unwrap_or("N/A")).cyan()
        )
    }
}


pub async fn start_service(service: &str) {
    match service {
        "all" => {
            println!("Starting all services...");
            if let Err(e) = Exec::shell("docker-compose up -d").join() {
                eprintln!("{}", format!("Error starting all services: {}", e).red());
            } else {
                println!("{}", "All services started successfully.".green());
            }
        }
        _ => {
            println!("{}", format!("Starting service: {}", service).yellow());
            let command = format!("docker-compose up -d {}", service);
            if let Err(e) = Exec::shell(&command).join() {
                eprintln!("{}", format!("Error starting service '{}': {}", service, e).red());
            } else {
                println!("{}", format!("Service '{}' started successfully.", service).green());
                inspect_service(service).await;
            }
        }
    }
}

pub async fn inspect_service(service: &str) {
    println!("{}", format!("Inspecting service: {}", service).yellow());

    let command = format!("docker inspect {}", service);
    match Exec::shell(&command).capture() {
        Ok(output) => {
            if let Err(e) = parse_and_display_inspect_output(&output.stdout_str()) {
                eprintln!("{}", format!("Failed to process inspection output for '{}': {}", service, e).red());
            }
        }
        Err(e) => eprintln!("{}", format!("Error inspecting service '{}': {}", service, e).red()),
    }
}

fn parse_and_display_inspect_output(output: &str) -> Result<(), String> {
    match parse_docker_inspect(output) {
        Ok(info) => {
            if info.ip_address.is_empty() {
                println!("{}", "Warning: IP address is not available.".yellow());
            }
            println!("{}\n{}", "Access Info:".bold(), info);
            Ok(())
        }
        Err(e) => Err(format!("Parsing error: {}", e)),
    }
}

fn parse_docker_inspect(output: &str) -> Result<ServiceInfo, &'static str> {
    let json: Value = serde_json::from_str(output).map_err(|_| "Invalid JSON format")?;

    let networks = json[0]["NetworkSettings"]["Networks"]
        .as_object()
        .ok_or("No network configuration found")?;

    let network = networks
        .get("vt_vulnerabilitytargets_default")
        .ok_or("Network 'vulnerabilitytargets_default' not found")?;

    let ip_address = network["IPAddress"]
        .as_str()
        .unwrap_or_default()
        .to_string();

    let ports = json[0]["NetworkSettings"]["Ports"]
        .as_object()
        .ok_or("No port configuration found")?;

    let web_port = ports
        .get("80/tcp")
        .and_then(|v| v.as_array())
        .and_then(|arr| arr.get(0))
        .and_then(|port_info| port_info["HostPort"].as_str())
        .map(|s| s.to_string());

    Ok(ServiceInfo { ip_address, web_port })
}
