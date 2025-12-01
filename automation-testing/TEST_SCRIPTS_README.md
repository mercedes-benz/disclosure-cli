# Disclosure-CLI Test Scripts

This directory contains comprehensive test scripts for the disclosure-cli tool that test all available commands and generate detailed reports.

## Available Scripts

### 1. PowerShell Script (Windows)
**File:** `test-all-commands.ps1`

**Usage:**
```powershell
# Run from the automation testing folder
cd "automation testing"
.\test-all-commands.ps1

# Run with custom config file (in parent directory)
.\test-all-commands.ps1 -ConfigFile "..\config.yml"

# Run with custom CLI path (in parent directory)
.\test-all-commands.ps1 -CliPath "..\disclosure-cli.exe"

# Run with both custom parameters
.\test-all-commands.ps1 -ConfigFile "..\myconfig.yml" -CliPath "..\disclosure-cli.exe"
```

### 2. Bash Script (Linux/Mac)
**File:** `test-all-commands.sh`

**Usage:**
```bash
# Make script executable
cd "automation testing"
chmod +x test-all-commands.sh

# Run with default settings (looks for config.yml in parent directory)
./test-all-commands.sh

# Run with custom config file
./test-all-commands.sh ../config.yml

# Run with custom config and CLI path
./test-all-commands.sh ../config.yml ../disclosure-cli
```

## What These Scripts Test

The scripts comprehensively test the following command categories:

### Basic Commands
- `--help` - Display help information
- `--version` - Display version information

### Project Commands
- `project details` - Get project details
- `project policyrules` - Get policy rules
- `project schema` - Get project schema
- `project status` - Get project status
- `project children` - Get project children
- `project sbomCheck` - On-demand SBOM check

### Version Commands
- `version list` - List all project versions
- `version details` - Get version details
- `version create` - Create new version (if applicable)
- `version delete` - Delete version (if applicable)
- `version ccs` - Get CCS status
- `version ccsAdd` - Add CCS reference
- `version sboms` - List all SBOMs
- `version sbomDetails` - Get SBOM details
- `version sbomStatus` - Get SBOM status
- `version sbomUpload` - Upload SBOM file
- `version sbomNotice` - Get SBOM notice

### SBOM Commands
- `sbom tag` - Tag SBOM
- `sbom lock` - Lock SBOM
- `sbom search` - Search SBOMs

### Utility Commands
- `sha256` - Generate SHA256 hash

### Negative Tests
- Invalid config file
- Invalid commands

## Output Formats

The scripts generate the following outputs:

### 1. Console Output
Real-time colored output showing:
- Test execution progress
- Pass/Fail status for each command
- Summary statistics with pass/fail counts
- Detailed failure information
- Success rate percentage

### 2. Log File (Bash only)
**File:** `test-results-YYYYMMDD-HHMMSS.log`

Detailed log containing:
- Full command output
- Error messages
- Timestamps
- Execution details

## Prerequisites

### For PowerShell Script
- Windows PowerShell 5.1 or higher
- Disclosure-CLI binary (will attempt to build if not found)
- Go installed (if building from source)
- Valid `config.yml` file

### For Bash Script
- Bash 4.0 or higher
- Disclosure-CLI binary (will attempt to build if not found)
- Go installed (if building from source)
- Valid `config.yml` file

## Configuration File

Both scripts require a `config.yml` file in the parent directory (root of disclosure-cli) with the following structure:

```yaml
projecttoken: "your-project-token"
projectuuid: "your-project-uuid"
projectversion: "dev"
host: "https://your-disclosure-host.com"
```

**Note:** The scripts look for `config.yml` in the parent directory by default. You can override this with the `-ConfigFile` parameter (PowerShell) or first argument (Bash).

## Exit Codes

- `0` - All tests passed
- `1` - One or more tests failed or script error

## Features

### Automatic CLI Detection
- Checks for existing CLI binary
- Attempts to build from source if not found
- Validates CLI availability before testing

### Comprehensive Testing
- Tests all documented commands
- Validates command output
- Checks exit codes
- Verifies expected patterns in output

### Error Handling
- Graceful failure handling
- Detailed error messages
- Continues testing even if individual tests fail

### Flexible Configuration
- Command-line parameters
- Environment variable support
- Config file validation

### Rich Reporting
- Multiple output formats
- Success/failure statistics
- Detailed failure analysis
- Timestamp tracking

## Example Output

```
═══════════════════════════════════════════════════════════════════════
           DISCLOSURE-CLI COMPREHENSIVE COMMAND TESTING                
═══════════════════════════════════════════════════════════════════════

[Pre-Flight Checks]
✓ CLI binary found at: ./disclosure-cli.exe
✓ Config file found: config.yml

[Starting Command Tests]

[Testing] Basic help command
Command: ./disclosure-cli --help
[PASS] Basic help command

[PROJECT COMMANDS]
[Testing] Get project details
Command: ./disclosure-cli project details -c config.yml
[PASS] Get project details

...

═══════════════════════════════════════════════════════════════════════
                    DISCLOSURE-CLI TEST REPORT                          
═══════════════════════════════════════════════════════════════════════

Summary:
  Total Tests: 25
  Passed:      23
  Failed:      2
  Success Rate: 92.00%
```

## Troubleshooting

### Script Won't Execute (PowerShell)
```powershell
# Check execution policy
Get-ExecutionPolicy

# If restricted, set to RemoteSigned for current user
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
```

### Script Won't Execute (Bash)
```bash
# Make sure script is executable
chmod +x test-all-commands.sh

# Check file permissions
ls -l test-all-commands.sh
```

### CLI Not Found
- Ensure `disclosure-cli` binary is in the same directory
- Or provide the path using `-CliPath` (PowerShell) or second argument (Bash)
- Or ensure Go is installed to build from source

### Config File Issues
- Verify `config.yml` exists in the current directory
- Check YAML syntax is correct
- Ensure all required fields are present
- Validate credentials are correct

### Command Failures
- Check network connectivity to the disclosure host
- Verify credentials in config.yml are valid
- Ensure project UUID and token are correct
- Check API endpoint is accessible

## Contributing

To add new tests:

1. Add the test command to the appropriate section in `run_all_tests` or `Start-Tests` function
2. Use the `test_cli_command` or `Test-CliCommand` function
3. Specify expected output patterns if needed
4. Set `should_succeed` appropriately for negative tests

## License

SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
SPDX-License-Identifier: MIT
