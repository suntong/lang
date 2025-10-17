package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	durationUsageDemo()
	durationOutputDemo()
	durationDiffingDemo()
	durationParsingDemo()
	customParsingDemo()
}

// --- Demos ---

// 1. Creation and Arithmetic
func durationUsageDemo() {
	fmt.Println("--- 1. Creation and Arithmetic ---")
	// Creation using constants
	d1 := 5 * time.Second
	d2 := 2*time.Minute + 30*time.Second
	d3 := time.Hour

	fmt.Printf("d1 (5s): %v\n", d1)
	fmt.Printf("d2 (2m30s): %v\n", d2)

	// Arithmetic (Addition, Subtraction)
	total := d1 + d2
	fmt.Printf("Addition (d1 + d2): %v\n", total) // 2m35s

	diff := d3 - d2                                 // 1h0m0s - 2m30s
	fmt.Printf("Subtraction (d3 - d2): %v\n", diff) // 57m30s

	// Negative Duration
	negativeDuration := -1 * time.Hour
	fmt.Printf("Negative duration: %v\n", negativeDuration) // -1h0m0s
}

// 2. Formatting and Conversion
func durationOutputDemo() {
	fmt.Println("\n--- 2. Formatting and Conversion ---")
	duration := 3*time.Hour + 25*time.Minute + 45*time.Second + 800*time.Millisecond

	// Default String() format
	fmt.Printf("Default format: %v\n", duration) // 3h25m45.8s

	// Converting to specific units (returns float64)
	fmt.Printf("Total Hours: %.2f\n", duration.Hours())     // 3.43
	fmt.Printf("Total Minutes: %.2f\n", duration.Minutes()) // 205.76
	fmt.Printf("Total Seconds: %.2f\n", duration.Seconds()) // 12345.80

	// Converting to integer units
	fmt.Printf("Total Milliseconds: %d\n", duration.Milliseconds())
}

// 3. Diffing
func durationDiffingDemo() {
	fmt.Println("\n--- 3. Diffing (Calculating Elapsed Time) ---")
	start := time.Now()
	// Simulate a long operation
	time.Sleep(3 * time.Millisecond) // Use a small sleep for quick demo

	// time.Since returns a time.Duration
	elapsed := time.Since(start)

	fmt.Printf("Start Time: %v\n", start.Format(time.RFC3339Nano))
	fmt.Printf("Elapsed Duration (time.Since): %v\n", elapsed)
}

// 4. Parsing
func durationParsingDemo() {
	fmt.Println("\n--- 4A. Standard Parsing (time.ParseDuration) ---")
	// Standard duration parsing uses suffixes (h, m, s, ms, etc.)
	d, err := time.ParseDuration("1h30m45.5s")
	if err != nil {
		fmt.Printf("Error parsing standard duration: %v\n", err)
		return
	}
	fmt.Printf("Input \"1h30m45.5s\": %v\n", d) // 1h30m45.5s

	d2, err := time.ParseDuration("-30m")
	if err == nil {
		fmt.Printf("Input \"-30m\": %v\n", d2) // -30m0s
	}
}

var d0, _ = time.Parse("15:04:05", "00:00:00")

// 5. Robust HH:MM:SS.mm Parsing Demo
func customParsingDemo() {
	fmt.Println("\n--- 4B. Custom HH:MM:SS.mm Format Parsing with Edge Cases ---")
	// Test cases for HH:MM:SS.mm format and its edge cases
	testCases := map[string]string{
		// Full format
		"1:5:3.58":     "1h5m3.58s",
		"1:05:3.58":    "1h5m3.58s",
		"0:15:13.58":   "15m13.58s",
		"00:15:13.58":  "15m13.58s",
		"25:10:05.123": "25h10m5.123s",
		// Missing .mm (milliseconds)
		"01:02:03": "1h2m3s",
		// Missing HH (MM:SS.mm)
		"15:30.500": "15m30.5s",
		// Missing HH and .mm (MM:SS)
		"05:01": "5m1s",
		// SS.mm (Pure seconds with fraction)
		"42.99": "42.99s",
		// SS (Pure seconds, check overflow to minutes)
		"75": "1m15s", // 75 seconds -> 1m15s
		// Edge Case: Hours > 24
		"100:00:00": "100h0m0s",
		// Edge Case: Milliseconds only
		"00:00:00.001": "1ms",
		// Invalid Cases
		"invalid":       "Error",
		"10:invalid:30": "Error",
	}

	fmt.Println("\n---- 4B1. Using time.Parse ---")
	//d0, _ := time.Parse("15:04:05", "00:00:00")
	for input, expected := range testCases {
		dur, err := time.Parse("15:04:05", input)
		if err != nil {
			// Check if the expected output is "Error"
			if expected != "Error" {
				fmt.Printf("❌ FAIL: Input '%s', Expected '%s', Got Error: %v\n", input, expected, err)
			} else {
				fmt.Printf("✅ PASS: Input '%s', Got expected error: %v\n", input, err)
			}
		} else {
			// Check if the formatted duration matches the expected string
			actual := dur.Sub(d0).String()
			if actual == expected {
				fmt.Printf("✅ PASS: Input '%s' -> %s\n", input, actual)
			} else {
				fmt.Printf("❌ FAIL: Input '%s', Expected '%s', Got %s\n", input, expected, actual)
			}
		}
	}

	fmt.Println("\n---- 4B2. Using ParseDuration ---")
	for input, expected := range testCases {
		dur, err := ParseDuration(input)
		if err != nil {
			// Check if the expected output is "Error"
			if expected != "Error" {
				fmt.Printf("❌ FAIL: Input '%s', Expected '%s', Got Error: %v\n", input, expected, err)
			} else {
				fmt.Printf("✅ PASS: Input '%s', Got expected error: %v\n", input, err)
			}
		} else {
			// Check if the formatted duration matches the expected string
			actual := dur.String()
			if actual == expected {
				fmt.Printf("✅ PASS: Input '%s' -> %s\n", input, actual)
			} else {
				fmt.Printf("❌ FAIL: Input '%s', Expected '%s', Got %s\n", input, expected, actual)
			}
		}
	}

	fmt.Println("\n---- 4B3. Using ParseHHMMSSmm ---")
	for input, expected := range testCases {
		dur, err := ParseHHMMSSmm(input)
		if err != nil {
			// Check if the expected output is "Error"
			if expected != "Error" {
				fmt.Printf("❌ FAIL: Input '%s', Expected '%s', Got Error: %v\n", input, expected, err)
			} else {
				fmt.Printf("✅ PASS: Input '%s', Got expected error: %v\n", input, err)
			}
		} else {
			// Check if the formatted duration matches the expected string
			actual := dur.String()
			if actual == expected {
				fmt.Printf("✅ PASS: Input '%s' -> %s\n", input, actual)
			} else {
				fmt.Printf("❌ FAIL: Input '%s', Expected '%s', Got %s\n", input, expected, actual)
			}
		}
	}

}

// --- Custom Parsing Function ---

var durationFormat = []string{"", "05", "04:05", "15:04:05"}

// ParseHHMMSSmm parses strings in format [HH:]MM:SS[.mm] into a time.Duration.
// It handles HH:MM:SS.mm, MM:SS.mm, SS.mm, and their integer counterparts.
func ParseDuration(durationStr string) (time.Duration, error) {

	// 1. Split the string by ':'
	parts := strings.Split(durationStr, ":")
	numParts := len(parts)

	if numParts == 0 || numParts > 3 {
		return 0, fmt.Errorf("invalid duration string format '%s'. Expected [HH:]MM:SS[.mm]", durationStr)
	}

	dur, err := time.Parse(durationFormat[numParts], durationStr)
	if err != nil {
		return 0, err
	}

	//d0, _ := time.Parse("15:04:05", "00:00:00")
	return dur.Sub(d0), nil

}

// ParseHHMMSSmm parses strings in format [HH:]MM:SS[.mm] into a time.Duration.
// It handles HH:MM:SS.mm, MM:SS.mm, SS.mm, and their integer counterparts.
func ParseHHMMSSmm(durationStr string) (time.Duration, error) {
	var totalDuration time.Duration

	// 1. Split the string by ':'
	parts := strings.Split(durationStr, ":")
	numParts := len(parts)

	if numParts == 0 || numParts > 3 {
		return 0, fmt.Errorf("invalid duration string format '%s'. Expected [HH:]MM:SS[.mm]", durationStr)
	}

	var h, m int
	var s float64 // Use float for seconds to handle the .mm fractional part

	// Determine components based on the number of colon-separated parts
	switch numParts {
	case 1:
		// Case: SS or SS.mm (no colons) - assume it's seconds
		sStr := parts[0]
		v, err := strconv.ParseFloat(sStr, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid seconds part '%s': %w", sStr, err)
		}
		s = v
	case 2:
		// Case: MM:SS or MM:SS.mm
		mStr, sStr := parts[0], parts[1]
		var err error

		// Parse Minutes
		m, err = strconv.Atoi(mStr)
		if err != nil {
			return 0, fmt.Errorf("invalid minutes part '%s': %w", mStr, err)
		}

		// Parse Seconds (and milliseconds)
		s, err = strconv.ParseFloat(sStr, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid seconds part '%s': %w", sStr, err)
		}
	case 3:
		// Case: HH:MM:SS or HH:MM:SS.mm
		hStr, mStr, sStr := parts[0], parts[1], parts[2]
		var err error

		// Parse Hours
		h, err = strconv.Atoi(hStr)
		if err != nil {
			return 0, fmt.Errorf("invalid hours part '%s': %w", hStr, err)
		}

		// Parse Minutes
		m, err = strconv.Atoi(mStr)
		if err != nil {
			return 0, fmt.Errorf("invalid minutes part '%s': %w", mStr, err)
		}

		// Parse Seconds (and milliseconds)
		s, err = strconv.ParseFloat(sStr, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid seconds part '%s': %w", sStr, err)
		}
	}

	// Accumulate total duration
	totalDuration += time.Duration(h) * time.Hour
	totalDuration += time.Duration(m) * time.Minute
	// Convert the float seconds (with its fraction) to a Duration in nanoseconds
	totalDuration += time.Duration(s * float64(time.Second))

	return totalDuration, nil
}
