# PassGuard

PassGuard is a lightweight password strength and exposure risk analyzer built in Go.  
It evaluates password security using rule-based complexity checks and common password detection, producing structured JSON output for security assessment.

---

## Features

- Analyzes password strength using rule-based checks  
- Detects weak passwords based on:
  - Length  
  - Uppercase letters  
  - Lowercase letters  
  - Numbers  
  - Symbols  
  - Common password patterns  
- Generates structured JSON reports  
- Provides actionable security recommendations  
- Ethical design: does not store or transmit passwords  
- Modular Go architecture for maintainability

---


### Requirements
- Go 1.20 or higher

### Installation Instructions

To install and run the tool, follow these steps:

1. Clone the repository:
    ```bash
   git clone https://github.com/FD785/PassGuard.git
    ```

2. Navigate to the project directory:
    ```bash
    cd PassGuard
    ```

3. Install the dependencies:
    ```bash
   go mod tidy
    ```

4. Run the tool:
    ```bash
    go run ./cmd/passguard password123
    ```

 Output will be like 👍
 ```bash
 {
  "score": 25,
  "risk_level": "high",                     <--- Shows Password Vulnerability 
  "issues": [
    "Password too short",                     |
    "Missing uppercase letters",              |--- Shows Password Vulnerability 
    "Missing symbols",                        |
    "Common password pattern"                 |
  ],
  "recommendations": [                        |
    "Use at least 12 characters",             |
    "Add uppercase characters",               |--- Shows How to create a Strong Password
    "Add special characters",                 |
    "Avoid dictionary words"                  |
  ]
}
```

    
