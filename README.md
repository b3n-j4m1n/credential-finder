# credential-finder

MITRE ATT&CK ID: T1081

This tool will search recursively from the current directory for files or folders containing the strings "password" or "credential", and read files of specified file type(s) and return the line number, path, and line containing the same strings.

```
Usage: go run main.go ...<extensions>
Example: go run main.go .txt .bat .py .sh .vb .vbs .ps1 .ps2 .csv .ini .env
```
