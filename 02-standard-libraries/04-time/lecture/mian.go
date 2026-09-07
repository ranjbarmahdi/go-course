package main

import (
	"fmt"
	"time"
)

// ============================================================
// 1. Current Time
// ============================================================
// time.Now() returns the current local time.
//
// The returned value is a time.Time.

func main() {

	now := time.Now()

	fmt.Println(now)
	fmt.Printf("%T => %+v\n", now, now)

	// ============================================================
	// 2. UTC Time
	// ============================================================
	// UTC() converts a time.Time to UTC.

	now1 := time.Now().UTC()

	fmt.Println(now1)

	// ============================================================
	// 3. Extract Date Information
	// ============================================================
	// A time.Time contains individual date and time components.

	now2 := time.Now()

	fmt.Println("Year:", now2.Year())
	fmt.Println("Month:", now2.Month())
	fmt.Println("Day:", now2.Day())
	fmt.Println("Hour:", now2.Hour())
	fmt.Println("Minute:", now2.Minute())
	fmt.Println("Second:", now2.Second())

	// ============================================================
	// 4. Creating Custom Time
	// ============================================================
	// time.Date() creates a time.Time from individual components.
	//
	// Arguments:
	//
	// year
	// month
	// day
	// hour
	// minute
	// second
	// nanosecond
	// location

	t := time.Date(
		2026,
		time.January,
		1,
		18,
		12,
		0,
		0,
		time.UTC,
	)

	fmt.Println(t)

	// ============================================================
	// 5. Duration
	// ============================================================
	// time.Duration represents a period of time.
	//
	// Common duration constants:
	//
	// time.Nanosecond
	// time.Microsecond
	// time.Millisecond
	// time.Second
	// time.Minute
	// time.Hour

	fmt.Println(5 * time.Nanosecond)
	fmt.Println(5 * time.Microsecond)
	fmt.Println(5 * time.Millisecond)
	fmt.Println(5 * time.Second)
	fmt.Println(5 * time.Minute)
	fmt.Println(5 * time.Hour)

	// Convert Duration to numeric values.

	fmt.Println(5 * time.Nanosecond.Nanoseconds())
	fmt.Println(5 * time.Microsecond.Microseconds())
	fmt.Println(5 * time.Millisecond.Milliseconds())
	fmt.Println(5 * time.Second.Seconds())
	fmt.Println(5 * time.Minute.Minutes())
	fmt.Println(5 * time.Hour.Hours())

	// ============================================================
	// 6. Sleeping
	// ============================================================
	// time.Sleep() pauses the current goroutine for a duration.

	fmt.Println("start")

	time.Sleep(2 * time.Second)

	fmt.Println("finish")

	// ============================================================
	// 7. Calculate Difference Between Times
	// ============================================================
	// Sub() calculates the duration between two times.

	start := time.Now()

	time.Sleep(2 * time.Second)

	end := time.Now()

	duration := end.Sub(start)

	fmt.Println(duration)
	fmt.Println(duration.Seconds())

	// ============================================================
	// 8. time.Since()
	// ============================================================
	// time.Since(t) returns how much time has passed since t.
	//
	// It is commonly used to measure how long an operation takes.

	start1 := time.Now()

	time.Sleep(3 * time.Second)

	elapsed := time.Since(start1)

	fmt.Println(elapsed)
	fmt.Println(elapsed.Seconds())

	// ============================================================
	// 9. Add Time
	// ============================================================
	// Add() returns a new time with a duration added to it.

	time3 := time.Now()

	fmt.Println("Hour:", time3.Hour())

	time3 = time3.Add(5 * time.Hour)

	fmt.Println("Hour:", time3.Hour())

	fmt.Println("UTC Hour:", time3.UTC().Hour())

	// Example:
	// Useful for expiration times such as JWT tokens.

	expiredAt := time.Now().Add(5 * time.Minute)

	fmt.Println("Expires at:", expiredAt)

	// ============================================================
	// 10. Subtract Time
	// ============================================================
	// Use a negative duration with Add() to subtract time.

	today := time.Now()

	fmt.Println("Today:", today.Day())

	yesterday := today.Add(-24 * time.Hour)

	fmt.Println("Yesterday:", yesterday.Day())

	// ============================================================
	// 11. Compare Times
	// ============================================================
	// Before() -> true if the first time is before the second.
	// After()  -> true if the first time is after the second.
	// Equal()  -> true if both times represent the same instant.

	a := time.Now()

	b := a.Add(5 * time.Hour)

	fmt.Println(a.Before(b))
	fmt.Println(a.After(b))

	b = b.Add(-5 * time.Hour)

	fmt.Println(a.Equal(b))

	// ============================================================
	// 12. Formatting Time
	// ============================================================
	// Go does not use formats such as:
	//
	// YYYY-MM-DD
	// HH:mm:ss
	//
	// Go uses a special reference time:
	//
	// Mon Jan 2 15:04:05 MST 2006
	//
	// You build your format using these reference values.

	now3 := time.Now()

	fmt.Println("Now:", now3)

	// ------------------------------------------------------------
	// Date + Time
	// ------------------------------------------------------------

	formatted := now3.Format(
		"2006-01-02 15:04:05",
	)

	fmt.Println("Date time:", formatted)

	// ------------------------------------------------------------
	// Date
	// ------------------------------------------------------------

	formatted = now3.Format(
		"02/01/2006",
	)

	fmt.Println("Date:", formatted)

	// ------------------------------------------------------------
	// 12-hour clock
	// ------------------------------------------------------------

	formatted = now3.Format(
		"01/02 03:04:05 PM 2006",
	)

	fmt.Println("12-hour format:", formatted)

	// ------------------------------------------------------------
	// ISO 8601 / RFC3339
	// ------------------------------------------------------------
	// Commonly used in APIs and databases.

	formatted = now3.Format(time.RFC3339)

	fmt.Println("ISO format:", formatted)

	// ============================================================
	// 13. Parsing String to Time
	// ============================================================
	// time.Parse() converts a formatted string into time.Time.
	//
	// The layout must match the input string.

	date, err := time.Parse(
		"2006-01-02",
		"2026-08-19",
	)

	fmt.Println(date)
	fmt.Println(err)

	// RFC3339 input

	date1, err1 := time.Parse(
		time.RFC3339,
		"2026-08-19T15:30:00+03:30",
	)

	fmt.Println(date1)
	fmt.Println(err1)

	// ============================================================
	// 14. Convert Time to Unix Timestamp
	// ============================================================
	// Unix timestamp represents time relative to:
	//
	// 1970-01-01 00:00:00 UTC
	//
	// Common forms:
	//
	// Unix()      -> seconds
	// UnixMilli() -> milliseconds

	time5 := time.Now()

	fmt.Println("Seconds:", time5.Unix())
	fmt.Println("Milliseconds:", time5.UnixMilli())

	// ============================================================
	// 15. Convert Unix Timestamp Back to Time
	// ============================================================
	// time.Unix() converts Unix seconds into time.Time.

	timestamp := int64(1787143245)

	date2 := time.Unix(
		timestamp,
		0,
	)

	fmt.Println(date2)

	// ============================================================
	// 16. Location and Timezone
	// ============================================================
	// Location represents a timezone.
	//
	// time.Now() uses the local machine's timezone.

	now4 := time.Now()

	fmt.Println("Location:", now4.Location())

	// ============================================================
	// 17. Load a Specific Timezone
	// ============================================================
	// time.LoadLocation() loads an IANA timezone.
	//
	// Examples:
	//
	// Asia/Baku
	// Asia/Tehran
	// Europe/London
	// America/New_York

	location, err := time.LoadLocation(
		"Asia/Baku",
	)

	if err != nil {
		panic(err)
	}

	bakuTime := time.Now().In(location)

	fmt.Println("Baku:", bakuTime)

	// ============================================================
	// 18. Timer
	// ============================================================
	// A Timer sends a value on its channel after the duration.
	//
	// <-timer.C waits until the timer fires.

	timer := time.NewTimer(3 * time.Second)

	<-timer.C

	fmt.Println("Timer finished")

	// ============================================================
	// 19. Ticker
	// ============================================================
	// A Ticker sends a value repeatedly at a fixed interval.
	//
	// ticker.Stop() stops the ticker.

	ticker := time.NewTicker(
		2 * time.Second,
	)

	counter := 0

	for range ticker.C {

		fmt.Println(
			"time:",
			time.Now().Format("15:04:05"),
		)

		counter++

		if counter == 5 {
			ticker.Stop()
			break
		}
	}
}
