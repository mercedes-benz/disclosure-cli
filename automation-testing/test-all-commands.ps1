#!/usr/bin/env pwsh
# SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
# SPDX-License-Identifier: MIT

param(
    [string]$ConfigFile = "..\config.yml",
    [string]$CliPath = "..\disclosure-cli.exe"
)

# Test results storage
$script:TestResults = @()
$script:TotalTests = 0
$script:PassedTests = 0
$script:FailedTests = 0

function Add-TestResult {
    param(
        [string]$Command,
        [string]$Description,
        [bool]$Passed,
        [string]$Output,
        [string]$Error,
        [int]$ExitCode
    )
    
    $script:TotalTests++
    if ($Passed) {
        $script:PassedTests++
        $Status = "PASS"
    } else {
        $script:FailedTests++
        $Status = "FAIL"
    }
    
    $script:TestResults += [PSCustomObject]@{
        Command = $Command
        Description = $Description
        Status = $Status
        ExitCode = $ExitCode
        Output = $Output
        Error = $Error
        Timestamp = Get-Date -Format "yyyy-MM-dd HH:mm:ss"
    }
}

function Test-CliCommand {
    param(
        [string]$Command,
        [string]$Description,
        [string[]]$ExpectedInOutput = @(),
        [bool]$ShouldSucceed = $true
    )
    
    Write-Host "`n[Testing] $Description" -ForegroundColor Cyan
    Write-Host "Command: $Command" -ForegroundColor Gray
    
    try {
        $output = Invoke-Expression "$Command 2>&1" | Out-String
        $exitCode = $LASTEXITCODE
        
        $passed = $false
        $errorMsg = ""
        
        if ($ShouldSucceed) {
            if ($exitCode -eq 0) {
                if ($ExpectedInOutput.Count -eq 0) {
                    $passed = $true
                } else {
                    $allFound = $true
                    foreach ($expected in $ExpectedInOutput) {
                        if ($output -notmatch $expected) {
                            $allFound = $false
                            $errorMsg = "Expected pattern not found: $expected"
                            break
                        }
                    }
                    $passed = $allFound
                }
            } else {
                $errorMsg = "Command returned non-zero exit code: $exitCode"
            }
        } else {
            $passed = ($exitCode -ne 0)
            if (-not $passed) {
                $errorMsg = "Command was expected to fail but succeeded"
            }
        }
        
        if ($passed) {
            Write-Host "[PASS] $Description" -ForegroundColor Green
        } else {
            Write-Host "[FAIL] $Description - $errorMsg" -ForegroundColor Red
        }
        
        Add-TestResult -Command $Command -Description $Description -Passed $passed `
                       -Output $output -Error $errorMsg -ExitCode $exitCode
        
    } catch {
        Write-Host "[ERROR] $Description - $_" -ForegroundColor Red
        Add-TestResult -Command $Command -Description $Description -Passed $false `
                       -Output "" -Error $_.Exception.Message -ExitCode -1
    }
}

function Test-CliExists {
    if (Test-Path $CliPath) {
        Write-Host "[OK] CLI binary found at: $CliPath" -ForegroundColor Green
        return $true
    } else {
        Write-Host "[!] CLI binary not found at: $CliPath" -ForegroundColor Red
        Write-Host "Building CLI..." -ForegroundColor Yellow
        
        try {
            go build -o disclosure-cli.exe
            if (Test-Path $CliPath) {
                Write-Host "[OK] CLI built successfully" -ForegroundColor Green
                return $true
            }
        } catch {
            Write-Host "[!] Failed to build CLI: $_" -ForegroundColor Red
            return $false
        }
    }
    return $false
}

function Test-ConfigExists {
    if (Test-Path $ConfigFile) {
        Write-Host "[OK] Config file found: $ConfigFile" -ForegroundColor Green
        $content = Get-Content $ConfigFile -Raw
        Write-Host "Config contents:" -ForegroundColor Gray
        Write-Host $content -ForegroundColor Gray
        return $true
    } else {
        Write-Host "[!] Config file not found: $ConfigFile" -ForegroundColor Red
        return $false
    }
}

function Show-TestReport {
    Write-Host "`n=======================================================================" -ForegroundColor Cyan
    Write-Host "                    DISCLOSURE-CLI TEST REPORT                          " -ForegroundColor Cyan
    Write-Host "=======================================================================" -ForegroundColor Cyan
    
    Write-Host "`nSummary:" -ForegroundColor Yellow
    Write-Host "  Total Tests: $script:TotalTests" -ForegroundColor White
    Write-Host "  Passed:      $script:PassedTests" -ForegroundColor Green
    Write-Host "  Failed:      $script:FailedTests" -ForegroundColor Red
    
    $successRate = if ($script:TotalTests -gt 0) { [math]::Round(($script:PassedTests / $script:TotalTests) * 100, 2) } else { 0 }
    Write-Host "  Success Rate: $successRate%" -ForegroundColor Yellow
    
    Write-Host "`n---------------------------------------------------------------------" -ForegroundColor Gray
    Write-Host "`nDetailed Results:" -ForegroundColor Yellow
    Write-Host "---------------------------------------------------------------------" -ForegroundColor Gray
    
    $script:TestResults | Format-Table -AutoSize -Property `
        @{Label="Status"; Expression={$_.Status}; Width=10},
        @{Label="Command"; Expression={if ($_.Command.Length -gt 45) {$_.Command.Substring(0,42) + "..."} else {$_.Command}}; Width=45},
        @{Label="Description"; Expression={$_.Description}; Width=35},
        @{Label="Exit Code"; Expression={$_.ExitCode}; Width=10}
    
    Write-Host "`nFailed Tests Details:" -ForegroundColor Yellow
    Write-Host "---------------------------------------------------------------------" -ForegroundColor Gray
    
    $failedTests = $script:TestResults | Where-Object { $_.Status -eq "FAIL" }
    if ($failedTests.Count -eq 0) {
        Write-Host "  No failed tests!" -ForegroundColor Green
    } else {
        foreach ($test in $failedTests) {
            Write-Host "`nCommand: $($test.Command)" -ForegroundColor Red
            Write-Host "Description: $($test.Description)" -ForegroundColor White
            Write-Host "Error: $($test.Error)" -ForegroundColor Red
            if ($test.Output) {
                Write-Host "Output:" -ForegroundColor Gray
                Write-Host $test.Output -ForegroundColor DarkGray
            }
            Write-Host "---------------------------------------------------------------------" -ForegroundColor Gray
        }
    }
    
    Write-Host "`n=======================================================================" -ForegroundColor Cyan
}

function Start-Tests {
    Write-Host "=======================================================================" -ForegroundColor Cyan
    Write-Host "           DISCLOSURE-CLI COMPREHENSIVE COMMAND TESTING                " -ForegroundColor Cyan
    Write-Host "=======================================================================" -ForegroundColor Cyan
    Write-Host "Started at: $(Get-Date -Format 'yyyy-MM-dd HH:mm:ss')" -ForegroundColor Gray
    
    Write-Host "`n[Pre-Flight Checks]" -ForegroundColor Yellow
    if (-not (Test-CliExists)) {
        Write-Host "Cannot proceed without CLI binary" -ForegroundColor Red
        return
    }
    
    Test-ConfigExists
    
    Write-Host "`n[Starting Command Tests]" -ForegroundColor Yellow
    
    # Basic help commands
    Test-CliCommand "$CliPath --help" "Basic help command" @("Available Commands", "Usage") $true
    
    # Project commands
    Write-Host "`n[PROJECT COMMANDS]" -ForegroundColor Magenta
    Test-CliCommand "$CliPath project --help" "Project help" @("Available Commands") $true
    Test-CliCommand "$CliPath project details -c $ConfigFile" "Get project details" @("uuid|name") $true
    Test-CliCommand "$CliPath project policyrules -c $ConfigFile" "Get policy rules" @() $true
    Test-CliCommand "$CliPath project schema -c $ConfigFile" "Get project schema" @() $true
    Test-CliCommand "$CliPath project status -c $ConfigFile" "Get project status" @() $true
    Test-CliCommand "$CliPath project children -c $ConfigFile" "Get project children" @() $true
    
    # Version commands
    Write-Host "`n[VERSION COMMANDS]" -ForegroundColor Magenta
    Test-CliCommand "$CliPath version --help" "Version help" @("Available Commands") $true
    Test-CliCommand "$CliPath version list -c $ConfigFile" "List project versions" @() $true
    Test-CliCommand "$CliPath version details -c $ConfigFile" "Get version details" @() $true
    Test-CliCommand "$CliPath version ccs -c $ConfigFile" "Get CCS status" @() $true
    Test-CliCommand "$CliPath version sboms -c $ConfigFile" "List SBOMs" @() $true
    
    # SBOM commands
    Write-Host "`n[SBOM COMMANDS]" -ForegroundColor Magenta
    Test-CliCommand "$CliPath sbom --help" "SBOM help" @("Available Commands") $true
    Test-CliCommand "$CliPath version sbomDetails -c $ConfigFile" "Get SBOM details" @() $true
    Test-CliCommand "$CliPath version sbomStatus -c $ConfigFile" "Get SBOM status" @() $true
    
    # Utility commands
    Write-Host "`n[UTILITY COMMANDS]" -ForegroundColor Magenta
    Test-CliCommand "$CliPath sha256 --help" "SHA256 help" @() $true
    
    # Negative tests
    Write-Host "`n[NEGATIVE TESTS]" -ForegroundColor Magenta
    Test-CliCommand "$CliPath project details -c nonexistent.yml" "Invalid config file" @() $false
    
    Show-TestReport
}

Start-Tests
