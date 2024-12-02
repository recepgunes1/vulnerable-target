use colored::*;


pub fn search_service(keyword: &str) {
    let services = vec![
        ("juice-shop", vec!["owasp", "api", "xss", "top10"]),
        ("dvwa", vec!["sql", "xss", "csrf", "php"]),
        ("mutillidae", vec!["xss", "csrf", "php", "owasp"]),
        ("bwapp", vec!["sql", "xss", "top10", "buggy"]),
        ("nodegoat", vec!["nodejs", "api", "nosql", "owasp"]),
        ("metasploitable", vec!["network", "os", "services"]),
        ("broken-crystals", vec!["api", "dotnet", "top10"]),
        ("dvws", vec!["api", "rest", "xss"]),
    ];

    println!("{}", format!("Searching for services matching '{}':", keyword).yellow().bold());
    let mut found = false;

    for (service, tags) in services.iter() {
        if service.contains(keyword) || tags.iter().any(|tag| tag.contains(keyword)) {
            println!("- {} (tags: {:?})", service, tags);
            found = true;
        }
    }

    if !found {
        println!("{}", format!("No services found matching '{}:", keyword).red().bold());
    }
}
