# Timestamp
 timestamp is a Go utility module for managing epoch timestamps. It provides easy-to-use functions for converting between standard time and epoch time, calculating future epochs, comparing timestamps, checking time ranges, and generating epochs for specific dates and times.

## Features

- Convert `time.Time` to epoch (seconds since Unix epoch).
- Convert epoch to `time.Time`.
- Calculate the epoch timestamp for a future date (e.g., 30 days from now).
- Compare two epoch timestamps (`<`, `>`, `<=`, `>=`).
- Check if an epoch timestamp is within a specific range.
- Get the epoch timestamp for a specific date and time.

## Installation

1. Install the module:
   ```bash
   go get github.com/scout-publishing-devops/timestamp
   ```

2. Import it in your project:
   ```go
   import "github.com/scout-publishing-devops/timestamp"
   ```

## Usage

### Examples

#### Convert `time.Time` to Epoch
```go
now := time.Now()
epochNow := timestamp.ToEpoch(now)
fmt.Println("Current time in epoch:", epochNow)
```

#### Convert Epoch to `time.Time`
```go
epoch := int64(1735121400)
convertedTime := timestamp.FromEpoch(epoch)
fmt.Println("Converted time:", convertedTime)
```

#### Calculate Epoch for 30 Days from Now
```go
epoch30Days := timestamp.EpochInDays(30)
fmt.Println("Epoch 30 days from now:", epoch30Days)
```

#### Compare Two Epoch Timestamps
```go
epoch1 := timestamp.ToEpoch(time.Now())
epoch2 := timestamp.EpochInDays(30) // Epoch for 30 days from now

isEarlier := timestamp.CompareEpoch(epoch1, epoch2, "<")
fmt.Println("Is now earlier than 30 days from now?:", isEarlier)
```

#### Check if Epoch is in a Time Range
```go
now := time.Now()
startRange := timestamp.ToEpoch(now.Add(-24 * time.Hour)) // 24 hours ago
endRange := timestamp.ToEpoch(now.Add(24 * time.Hour))   // 24 hours from now

epochNow := timestamp.ToEpoch(now)
isInRange := timestamp.TimeRange(epochNow, startRange, endRange)
fmt.Println("Is now within the range?:", isInRange)
```

#### Get Epoch for a Specific Date and Time
```go
date := "2024-12-25"
timeStr := "15:30:00" // 3:30 PM

epoch, err := timestamp.EpochForDate(date, timeStr)
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Epoch for", date, timeStr, "is:", epoch)
}
```

### Output Examples

1. **Convert `time.Time` to Epoch**:
   ```
   Current time in epoch: 1732399320
   ```

2. **Convert Epoch to `time.Time`**:
   ```
   Converted time: 2024-12-25 15:30:00 +0000 UTC
   ```

3. **Calculate Epoch for 30 Days from Now**:
   ```
   Epoch 30 days from now: 1735121400
   ```

4. **Compare Two Epoch Timestamps**:
   ```
   Is now earlier than 30 days from now?: true
   ```

5. **Check if Epoch is in a Time Range**:
   ```
   Is now within the range?: true
   ```

6. **Get Epoch for a Specific Date and Time**:
   ```
   Epoch for 2024-12-25 15:30:00 is: 1735121400
   ```

### Error Handling
The `EpochForDate` function will return an error if the date or time format is invalid:
```bash
Error: invalid date or time format, use 'YYYY-MM-DD' and 'HH:MM:SS'
```

## License

This module is licensed under the [MIT License](https://opensource.org/licenses/MIT).

## Contributing

Feel free to submit issues or pull requests for improvements and bug fixes.

---