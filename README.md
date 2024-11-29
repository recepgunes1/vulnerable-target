# Vulnerable Target


## Features
- Community-Curated List of Vulnerable Targets
- Interactive Vulnerability Playground (TODO)
- CLI (TODO)


## How to Contribute

We welcome contributions from everyone! Here's how you can contribute:


### Adding a New Vulnerable Target

1. **Fork the Repository**  
   Start by forking our [vulnerable-target repository](https://github.com/HappyHackingSpace/vulnerable-target).

2. **Create a New Target Entry**  
   - Add a new entry to the `vulnerable-target-list.json` file.  
   - Each entry must follow this structure:
     ```json
     {
       "name": "example-vulnerable-app",
       "description": "An intentionally vulnerable web app for SQL injection testing.",
       "url": "https://vulnerabletarget.com",
       "technologies": ["php", "mysql"]
       "tags": ["sql-injection", "web", "beginner"]
     }
     ```

3. **Submit a Pull Request**  
   Open a pull request with your changes.


### Contribution File Format

#### Required Fields:
- `name` (string): Unique identifier for the target.
- `description` (string): A brief description of the target and its vulnerabilities.
- `url` (string): Link to the repository or resource.
- `tags` (array of strings): Keywords that describe the vulnerabilities (e.g., `xss`, `api`, `authentication`).

#### Example
```json
{
  "name": "vulnerable-blog-app",
  "description": "A vulnerable blog application for testing XSS and CSRF attacks.",
  "url": "https://blog.vulnerabletarget.com",
  "technologies": ["php", "mysql"]
  "tags": ["xss", "csrf", "web"]
}
```


Happy hacking!
