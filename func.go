package gopinba

type dictionary []string

func inSlice(needle string, haystack dictionary) (int, bool) {

	for position, value := range haystack {

		if value == needle {

			return position, true
		}
	}

	return -1, false
}
