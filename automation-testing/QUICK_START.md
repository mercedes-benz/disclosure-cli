# Quick Start Guide - Disclosure-CLI Test Scripts


### Windows
```powershell
cd "c:\path\to\disclosure-cli\automation testing"
.\test-all-commands.ps1
```

### Linux/Mac
```bash
cd /path/to/disclosure-cli/automation\ testing
chmod +x test-all-commands.sh
./test-all-commands.sh
```

## 📋 What You Get

After running, you'll have:

1. **Real-time console output** with color-coded results
2. **Summary report** showing pass/fail statistics
3. **Log file** (Bash only) - `test-results-YYYYMMDD-HHMMSS.log`

## ✅ Sample Output

```
[PROJECT COMMANDS]
[Testing] Get project details
[PASS] Get project details

Summary:
  Total Tests: 17
  Passed:      17
  Failed:      0
  Success Rate: 100%
```

## 📊 Test Coverage

| Category | Tests | Commands |
|----------|-------|----------|
| Project Commands | 6 | details, policyrules, schema, status, children, sbomCheck |
| Version Commands | 5 | list, details, ccs, sboms, sbomDetails |
| SBOM Commands | 3 | help, search, tag |
| Utility | 2 | help, sha256 |
| Negative Tests | 1 | invalid config |
| **TOTAL** | **17** | **All major commands** |

## 🔧 Requirements

- [ ] disclosure-cli binary in parent directory (or Go to build)
- [ ] config.yml in parent directory with valid credentials
- [ ] Network access to disclosure host

## 📝 config.yml Template

Place this file in the disclosure-cli root directory (parent folder):

```yaml
projecttoken: "your-project-token-here"
projectuuid: "your-project-uuid-here"
projectversion: "dev"
host: "https://your-disclosure-host.com"
```

## 💡 Pro Tips

1. **Run before commits** - Ensure all commands work
2. **CI/CD integration** - Add to pipeline
3. **Check log file** - Detailed execution logs (Bash)
4. **Track trends** - Compare results over time

## 🎯 Exit Codes

- `0` = All tests passed ✅
- `1` = Some tests failed ❌

## 📖 More Info

- Full documentation: `TEST_SCRIPTS_README.md`
