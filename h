func ParseInput(input string) []string {
	// Expecting input like: "title~2025-06-10 21:21:13.370097 +0530 IST m=+62.407111835.md"
	parts := strings.SplitN(input, "~", 2)
	if len(parts) != 2 {
		return []string{"", ""}
	}
	title := parts[0]
	dateStr := strings.TrimSuffix(parts[1], ".md")

	// Split dateStr into fields
	dateFields := strings.Fields(dateStr)
	if len(dateFields) < 5 {
		fmt.Println("Date string format is invalid")
		return []string{title, ""}
	}

	// Parse the datetime with timezone offset
	layout := "2006-01-02 15:04:05.999999 -0700"
	datetimeStr := fmt.Sprintf("%s %s %s", dateFields[0], dateFields[1], dateFields[2])
	t, err := time.Parse(layout, datetimeStr)
	if err != nil {
		fmt.Println("Parse error:", err)
		return []string{title, ""}
	}

	// Extract fields
	day := fmt.Sprintf("%02d", t.Day())
	month := t.Month().String()
	year := fmt.Sprintf("%d", t.Year())
	hour := t.Format("03") // 12-hour format
	minute := t.Format("04")
	second := t.Format("05")
	ampm := t.Format("PM")
	timeStr := fmt.Sprintf("%s:%s:%s", hour, minute, second)
	timezone := dateFields[3]
	offset := dateFields[2]

	result := []string{day, month, year, timeStr, ampm, timezone, offset}
	dateString := strings.Join(result, " ")
	return []string{title, dateString}
}
