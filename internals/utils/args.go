package utils

func GetDescription(args []string, idx int) string {
	if len(args) < idx+1 {
		return ""
	}

	return args[idx]
}