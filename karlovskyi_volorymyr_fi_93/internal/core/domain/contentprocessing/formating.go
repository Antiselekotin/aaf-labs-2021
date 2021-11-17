package contentprocessing

func ShiftAndNewLineString(swift_size int, s string) string {
	template := ""
	if swift_size < 0 {
		return s
	}
	for i := 0; i < swift_size; i++ {
		template += "  "
	}
	return template + s + "\n"
}

func ShiftString(swift_size int, s string) string {
	template := ""
	if swift_size < 0 {
		return s
	}
	for i := 0; i < swift_size; i++ {
		template += "  "
	}
	return template + s
}
