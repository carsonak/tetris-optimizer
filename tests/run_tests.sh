#!/bin/bash

set -o nounset -o pipefail

# Configuration & Defaults
BAD_DIR="bad_examples"
GOOD_DIR="good_examples"
SAMPLES_DIR="samples"

# Flags
USE_COLOR=true
RUN_EXAMPLES=true
RUN_SAMPLES=false

EXECUTABLE=""

# ANSI Colors
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
RESET='\033[0m' # Reset Terminal

print_help() {
    cat << EOF
USAGE: ./run_tests [OPTIONS] <executable>
   Runs tests on the given executable.

OPTIONS
   -b, --bad-examples <dir_path>
       path to directory with bad examples, defaults to "bad_examples".
   --colour
       toggles colourisation, defaults to "on/true".
   -e, --examples
       toggles running of examples tests, defaults to "on/true".
   -g, --good-examples <dir_path>
       path to directory with good examples, defaults to "good_examples".
   -h, --help
       print this help message and exit.
   -s, --samples
       toggles running of samples tests, defaults to "off/false".
EOF
}

# Parse arguments
while [[ $# -gt 0 ]]; do
    case "$1" in
        -b|--bad-examples)
            BAD_DIR="$2"
            shift 2
            ;;
        --colour)
            # Toggle logic: if true, make false; if false, make true.
            if [[ "$USE_COLOR" = true ]]
            then USE_COLOR=false
            else USE_COLOR=true
            fi
            shift
            ;;
        -e|--examples)
            if [[ "$RUN_EXAMPLES" = true ]]
            then RUN_EXAMPLES=false
            else RUN_EXAMPLES=true
            fi
            shift
            ;;
        -g|--good-examples)
            GOOD_DIR="$2"
            shift 2
            ;;
        -h|--help)
            print_help
            exit 0
            ;;
        -s|--samples)
            if [[ "$RUN_SAMPLES" = true ]]
            then RUN_SAMPLES=false
            else RUN_SAMPLES=true
            fi
            shift
            ;;
        *)
            if [[ -z "$EXECUTABLE" ]]
            then
                EXECUTABLE="$1"
            else
                echo "Error: unknown argument: $1"
                exit 1
            fi
            shift
            ;;
    esac
done

log_pass() {
    if [[ "$USE_COLOR" = true ]]
    then echo -e "[${GREEN}PASS${RESET}] $1"
    else echo "[PASS] $1"
    fi
}

log_fail() {
    if [[ "$USE_COLOR" = true ]]
    then echo -e "[${RED}FAIL${RESET}] $1"
    else echo "[FAIL] $1"
    fi
}

log_info() {
    if [[ "$USE_COLOR" = true ]]
    then echo -e "[${YELLOW}INFO${RESET}] $1"
    else echo "[INFO] $1"
    fi
}

log_header() {
    if [[ "$USE_COLOR" = true ]]
    then echo -e "\n${BLUE}--- $1 ---${RESET}"
    else echo -e "\n--- $1 ---"
    fi
}

# Check Executable
if [[ -z "$EXECUTABLE" ]]
then
    echo "Error: No executable provided."
    print_help
    exit 1
fi

if ! command -v "$EXECUTABLE" >/dev/null 2>&1 && [[ ! -x "$EXECUTABLE" ]]
then
    echo "Error: Executable '$EXECUTABLE' not found or not executable."
    exit 1
fi

# Bad Tests (Expects ERROR)
run_bad_test() {
    local file="$1"
    local filename
    filename=$(basename "$file")

    # Capture combined stdout and stderr
    local output
    output=$("$EXECUTABLE" "$file" 2>&1)

    # Check if output starts with ERROR
    if [[ "$output" == ERROR* ]]
    then
        log_pass "$filename"
    else
        log_fail "$filename"
        echo "       Expected: ERROR..."
        echo "       Got:      $output"
    fi
}

# Good Tests (Expects success + space count check)
run_good_test() {
    local file="$1"
    local filename
    filename=$(basename "$file")

    # Parse expected dots from filename suffix (format: name-NN)
    local expected_spaces=""
    if [[ "$filename" =~ -([0-9]{2})$ ]]
    then
        # Force base 10 to avoid octal interpretation of 08/09
        expected_spaces=$((10#${BASH_REMATCH[1]}))
    fi

    local start_ts end_ts duration
    # Use nanoseconds if date supports it (%N), otherwise seconds
    if date +%N >/dev/null 2>&1
    then
        start_ts=$(date +%s%N)
        local output
        output=$("$EXECUTABLE" "$file")
        end_ts=$(date +%s%N)
        # Calculate seconds with decimals
        duration=$(echo "scale=3; ($end_ts - $start_ts) / 1000000000" | bc 2>/dev/null || echo "0")
    else
        start_ts=$(date +%s)
        local output
        output=$("$EXECUTABLE" "$file")
        end_ts=$(date +%s)
        duration=$((end_ts - start_ts))
    fi

    if [[ "$output" == ERROR* ]]
    then
        log_fail "$filename: Program returned ERROR"
        return
    fi

    local space_count
    # Count empty spaces
    space_count=$(echo "$output" | grep -o ' ' | wc -l)

    if [[ -n "$expected_spaces" ]]
    then
        if [[ "$space_count" -eq "$expected_spaces" ]]
        then
            log_pass "$filename: $space_count spaces (${duration}s)"
        else
            log_fail "$filename: Expected $expected_spaces spaces, got $space_count"
        fi
    else
        log_info "$filename: $space_count spaces (${duration}s)"
    fi
}

if [[ "$RUN_EXAMPLES" = true ]]
then
    if [[ -d "$BAD_DIR" ]]
    then
        log_header "Bad Examples"
        for f in "$BAD_DIR"/*; do
            [[ -e "$f" ]] || continue
            run_bad_test "$f"
        done
    else
        if [[ "$USE_COLOR" = true ]]
        then echo -e "${YELLOW}Warning: Directory '$BAD_DIR' not found.${RESET}"
        else echo "Warning: Directory '$BAD_DIR' not found."
        fi
    fi

    if [[ -d "$GOOD_DIR" ]]
    then
        log_header "Good Examples"
        for f in "$GOOD_DIR"/*; do
            [[ -e "$f" ]] || continue
            run_good_test "$f"
        done
    else
        if [[ "$USE_COLOR" = true ]]
        then echo -e "${YELLOW}Warning: Directory '$GOOD_DIR' not found.${RESET}"
        else echo "Warning: Directory '$GOOD_DIR' not found."
        fi
    fi
fi

if [[ "$RUN_SAMPLES" = true ]]
then
    if [[ -d "$SAMPLES_DIR" ]]
    then
        log_header "Samples"
        for f in "$SAMPLES_DIR"/*; do
            [[ -e "$f" ]] || continue
            run_good_test "$f"
        done
    else
        if [[ "$USE_COLOR" = true ]]
        then echo -e "${YELLOW}Warning: Directory '$SAMPLES_DIR' not found.${RESET}"
        else echo "Warning: Directory '$SAMPLES_DIR' not found."
        fi
    fi
fi
