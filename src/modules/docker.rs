use subprocess::Exec;

pub async fn start_service(service: &str) {
    if service == "all" {
        println!("Starting all services...");
        let result = Exec::shell("docker-compose up -d").join();
        match result {
            Ok(_) => println!("All services started."),
            Err(e) => eprintln!("Error starting services: {}", e),
        }
    } else {
        println!("Starting service: {}", service);
        let result = Exec::shell(format!("docker-compose up -d {}", service)).join();
        match result {
            Ok(_) => {
                println!("Service '{}' started.", service);
                inspect_service(service).await; 
            }
            Err(e) => eprintln!("Error starting service '{}': {}", service, e),
        }
    }
}

pub async fn inspect_service(service: &str) {
    println!("Inspecting service: {}", service);
    let cmd = format!("docker inspect {}", service);
    let result = Exec::shell(&cmd).capture();

    match result {
        Ok(output) => {
            let output_str = output.stdout_str();
            match parse_docker_inspect(&output_str) {
                Ok(info) => {
                    if info.ip_address.is_empty() {
                        println!("Warning: IP address is not available.");
                    }
                    println!(
                        "Access Info:\n- IP Address: {}\n- Web Port: {}",
                        info.ip_address,
                        info.web_port.unwrap_or_else(|| "N/A".to_string())
                    );
                }
                Err(e) => eprintln!("Failed to parse inspect output: {}", e),
            }
        }
        Err(e) => eprintln!("Error inspecting service '{}': {}", service, e),
    }
}

#[derive(Debug)]
struct ServiceInfo {
    ip_address: String,
    web_port: Option<String>,
}

fn parse_docker_inspect(output: &str) -> Result<ServiceInfo, &'static str> {
    let json: serde_json::Value = serde_json::from_str(output).map_err(|_| "Invalid JSON")?;

    let networks = json[0]["NetworkSettings"]["Networks"]
        .as_object()
        .ok_or("No networks found")?;

    let network = networks
        .get("vulnerabilitytargets_default")
        .ok_or("Network 'vulnerabilitytargets_default' not found")?;

    let ip_address = network["IPAddress"]
        .as_str()
        .unwrap_or("") // empty string if no ip address
        .to_string();

    let ports = json[0]["NetworkSettings"]["Ports"]
        .as_object()
        .ok_or("No ports found")?;
    let web_port = ports
        .get("80/tcp")
        .and_then(|v| v.as_array())
        .and_then(|arr| arr.get(0))
        .and_then(|port_info| port_info["HostPort"].as_str())
        .map(|s| s.to_string());

    Ok(ServiceInfo { ip_address, web_port })
}
