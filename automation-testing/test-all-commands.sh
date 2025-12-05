#!/bin/bash
# SPDX-FileCopyrightText: 2023 Mercedes-Benz Tech Innovation GmbH
# SPDX-License-Identifier: MIT

##############################################################################
# Comprehensive test script for disclosure-cli commands
# Tests all available commands and generates a detailed report
##############################################################################

set -o pipefail

# Configuration
CONFIG_FILE="${1:-../config.yml}"
CLI_PATH="${2:-../disclosure-cli}"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
MAGENTA='\033[0;35m'
CYAN='\033[0;36m'
GRAY='\033[0;90m'
NC='\033[0m' # No Color

# Test results
TOTAL_TESTS=0
PASSED_TESTS=0
FAILED_TESTS=0
TEST_RESULTS=()
TIMESTAMP=$(date +"%Y%m%d-%H%M%S")

# Log file
LOG_FILE="test-results-${TIMESTAMP}.log"

##############################################################################
# Helper Functions
##############################################################################

print_header() {
    echo -e "${CYAN}=======================================================================${NC}"
    echo -e "${CYAN}           DISCLOSURE-CLI COMPREHENSIVE COMMAND TESTING                ${NC}"
    echo -e "${CYAN}=======================================================================${NC}"
    echo -e "${GRAY}Started at: $(date '+%Y-%m-%d %H:%M:%S')${NC}"
}

print_section() {
    echo -e "\n${MAGENTA}[$1]${NC}"
}

print_info() {
    echo -e "${CYAN}[INFO]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[PASS]${NC} $1"
}

print_failure() {
    echo -e "${RED}[FAIL]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

##############################################################################
# Test Execution Functions
##############################################################################

add_test_result() {
    local command="$1"
    local description="$2"
    local status="$3"
    local exit_code="$4"
    local output="$5"
    local error="$6"
    local timestamp=$(date '+%Y-%m-%d %H:%M:%S')
    
    TOTAL_TESTS=$((TOTAL_TESTS + 1))
    
    if [ "$status" = "PASS" ]; then
        PASSED_TESTS=$((PASSED_TESTS + 1))
        status_symbol="✓"
    else
        FAILED_TESTS=$((FAILED_TESTS + 1))
        status_symbol="✗"
    fi
    
    # Store result
    TEST_RESULTS+=("$status_symbol|$status|$command|$description|$exit_code|$error|$timestamp")
    
    # Write to log
    {
        echo "----------------------------------------"
        echo "Status: $status_symbol $status"
        echo "Command: $command"
        echo "Description: $description"
        echo "Exit Code: $exit_code"
        echo "Timestamp: $timestamp"
        if [ -n "$error" ]; then
            echo "Error: $error"
        fi
        if [ -n "$output" ]; then
            echo "Output:"
            echo "$output"
        fi
        echo ""
    } >> "$LOG_FILE"
}

test_cli_command() {
    local cmd="$1"
    local description="$2"
    local expected_patterns="$3"
    local should_succeed="${4:-true}"
    
    echo -e "\n${BLUE}[Testing]${NC} $description"
    echo -e "${GRAY}Command: $cmd${NC}"
    
    # Execute command and capture output
    local output
    local exit_code
    output=$(eval "$cmd 2>&1")
    exit_code=$?
    
    local passed=false
    local error_msg=""
    
    if [ "$should_succeed" = "true" ]; then
        if [ $exit_code -eq 0 ]; then
            # Check for expected patterns
            if [ -z "$expected_patterns" ]; then
                passed=true
            else
                passed=true
                IFS='|' read -ra PATTERNS <<< "$expected_patterns"
                for pattern in "${PATTERNS[@]}"; do
                    if ! echo "$output" | grep -qi "$pattern"; then
                        passed=false
                        error_msg="Expected pattern not found: $pattern"
                        break
                    fi
                done
            fi
        else
            error_msg="Command returned non-zero exit code: $exit_code"
        fi
    else
        # For commands expected to fail
        if [ $exit_code -ne 0 ]; then
            passed=true
        else
            error_msg="Command was expected to fail but succeeded"
        fi
    fi
    
    if [ "$passed" = true ]; then
        print_success "$description"
        add_test_result "$cmd" "$description" "PASS" "$exit_code" "$output" ""
    else
        print_failure "$description - $error_msg"
        add_test_result "$cmd" "$description" "FAIL" "$exit_code" "$output" "$error_msg"
    fi
}

##############################################################################
# Pre-flight Checks
##############################################################################

check_cli_exists() {
    print_info "Checking for CLI binary..."
    
    if [ -f "$CLI_PATH" ]; then
        print_success "CLI binary found at: $CLI_PATH"
        return 0
    else
        print_warning "CLI binary not found at: $CLI_PATH"
        print_info "Attempting to build CLI..."
        
        if command -v go &> /dev/null; then
            if go build -o disclosure-cli; then
                CLI_PATH="./disclosure-cli"
                print_success "CLI built successfully"
                return 0
            else
                print_failure "Failed to build CLI"
                return 1
            fi
        else
            print_failure "Go is not installed. Cannot build CLI."
            return 1
        fi
    fi
}

check_config_exists() {
    print_info "Checking for config file..."
    
    if [ -f "$CONFIG_FILE" ]; then
        print_success "Config file found: $CONFIG_FILE"
        echo -e "${GRAY}Config contents:${NC}"
        cat "$CONFIG_FILE" | while read -r line; do
            echo -e "${GRAY}  $line${NC}"
        done
        return 0
    else
        print_warning "Config file not found: $CONFIG_FILE"
        return 1
    fi
}

##############################################################################
# Test Suites
##############################################################################

run_all_tests() {
    print_section "PRE-FLIGHT CHECKS"
    
    if ! check_cli_exists; then
        print_failure "Cannot proceed without CLI binary"
        exit 1
    fi
    
    check_config_exists
    
    print_section "STARTING COMMAND TESTS"
    
    # Basic help commands
    test_cli_command "$CLI_PATH --help" "Basic help command" "Available Commands|Usage" true
    test_cli_command "$CLI_PATH version --help" "Version help command" "Available Commands" true
    
    # Project commands
    print_section "PROJECT COMMANDS"
    test_cli_command "$CLI_PATH project --help" "Project help" "Available Commands" true
    test_cli_command "$CLI_PATH project details -c $CONFIG_FILE" "Get project details" "uuid|name" true
    test_cli_command "$CLI_PATH project policyrules -c $CONFIG_FILE" "Get policy rules" "" true
    test_cli_command "$CLI_PATH project schema -c $CONFIG_FILE" "Get project schema" "" true
    test_cli_command "$CLI_PATH project status -c $CONFIG_FILE" "Get project status" "" true
    test_cli_command "$CLI_PATH project children -c $CONFIG_FILE" "Get project children" "" true
    test_cli_command "$CLI_PATH project sbomCheck -c $CONFIG_FILE" "On-demand SBOM check" "" true
    
    # Version commands
    print_section "VERSION COMMANDS"
    test_cli_command "$CLI_PATH version --help" "Version command help" "Available Commands" true
    test_cli_command "$CLI_PATH version list -c $CONFIG_FILE" "List project versions" "" true
    test_cli_command "$CLI_PATH version details -c $CONFIG_FILE" "Get version details" "" true
    test_cli_command "$CLI_PATH version ccs -c $CONFIG_FILE" "Get CCS status" "" true
    test_cli_command "$CLI_PATH version sboms -c $CONFIG_FILE" "List SBOMs" "" true
    test_cli_command "$CLI_PATH version sbomDetails -c $CONFIG_FILE" "Get SBOM details" "" true
    test_cli_command "$CLI_PATH version sbomStatus -c $CONFIG_FILE" "Get SBOM status" "" true
    
    # SBOM commands
    print_section "SBOM COMMANDS"
    test_cli_command "$CLI_PATH sbom --help" "SBOM help" "Available Commands" true
    test_cli_command "$CLI_PATH sbom search -c $CONFIG_FILE" "SBOM search" "" true
    
    # Utility commands
    print_section "UTILITY COMMANDS"
    test_cli_command "$CLI_PATH sha256 --help" "SHA256 help" "" true
    
    # Negative tests
    print_section "NEGATIVE TESTS"
    test_cli_command "$CLI_PATH project details -c nonexistent.yml" "Invalid config file" "" false
    test_cli_command "$CLI_PATH invalid-command" "Invalid command" "" false
}

##############################################################################
# Report Generation
##############################################################################

show_summary_report() {
    local success_rate=0
    if [ $TOTAL_TESTS -gt 0 ]; then
        success_rate=$(awk "BEGIN {printf \"%.2f\", ($PASSED_TESTS / $TOTAL_TESTS) * 100}")
    fi
    
    echo ""
    echo -e "${CYAN}=======================================================================${NC}"
    echo -e "${CYAN}                    DISCLOSURE-CLI TEST REPORT                          ${NC}"
    echo -e "${CYAN}=======================================================================${NC}"
    
    echo -e "\n${YELLOW}Summary:${NC}"
    echo -e "  Total Tests: $TOTAL_TESTS"
    echo -e "  ${GREEN}Passed:      $PASSED_TESTS${NC}"
    echo -e "  ${RED}Failed:      $FAILED_TESTS${NC}"
    echo -e "  ${YELLOW}Success Rate: ${success_rate}%${NC}"
    
    echo -e "\n${GRAY}---------------------------------------------------------------------${NC}"
    echo -e "\n${YELLOW}Detailed Results:${NC}"
    echo -e "${GRAY}---------------------------------------------------------------------${NC}"
    
    printf "%-10s %-50s %-30s %-10s\n" "Status" "Command" "Description" "Exit Code"
    echo -e "${GRAY}---------------------------------------------------------------------${NC}"
    
    for result in "${TEST_RESULTS[@]}"; do
        IFS='|' read -r symbol status cmd desc exit_code error timestamp <<< "$result"
        # Truncate long strings for display
        cmd_short=$(echo "$cmd" | cut -c1-48)
        desc_short=$(echo "$desc" | cut -c1-28)
        
        if [ "$status" = "PASS" ]; then
            printf "${GREEN}%-10s${NC} %-50s %-30s %-10s\n" "$symbol $status" "$cmd_short" "$desc_short" "$exit_code"
        else
            printf "${RED}%-10s${NC} %-50s %-30s %-10s\n" "$symbol $status" "$cmd_short" "$desc_short" "$exit_code"
        fi
    done
    
    echo -e "\n${YELLOW}Failed Tests Details:${NC}"
    echo -e "${GRAY}---------------------------------------------------------------------${NC}"
    
    local has_failures=false
    for result in "${TEST_RESULTS[@]}"; do
        IFS='|' read -r symbol status cmd desc exit_code error timestamp <<< "$result"
        if [ "$status" = "FAIL" ]; then
            has_failures=true
            echo -e "\n${RED}Command:${NC} $cmd"
            echo -e "${RED}Description:${NC} $desc"
            echo -e "${RED}Error:${NC} $error"
            echo -e "${GRAY}---------------------------------------------------------------------${NC}"
        fi
    done
    
    if [ "$has_failures" = false ]; then
        echo -e "  ${GREEN}No failed tests! 🎉${NC}"
    fi
    
    echo -e "\n${CYAN}=======================================================================${NC}"
}

##############################################################################
# Main Execution
##############################################################################

main() {
    # Initialize log file
    echo "Disclosure-CLI Test Execution Log" > "$LOG_FILE"
    echo "Started: $(date '+%Y-%m-%d %H:%M:%S')" >> "$LOG_FILE"
    echo "========================================" >> "$LOG_FILE"
    echo "" >> "$LOG_FILE"
    
    print_header
    
    # Run all tests
    run_all_tests
    
    # Generate report
    echo ""
    print_section "GENERATING REPORT"
    print_success "Log file generated: $LOG_FILE"
    
    # Show summary
    show_summary_report
    
    # Exit with appropriate code
    if [ $FAILED_TESTS -eq 0 ]; then
        exit 0
    else
        exit 1
    fi
}

# Run main function
main
