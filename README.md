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

## Installation

### Requirements
- Go 1.20 or higher

### Clone the repository
Step 1
```bash
git clone https://github.com/FD785/PassGuard.git
cd PassGuard
