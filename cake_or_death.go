package c

func CakeOrDeath(preference string) string {
	if preference == "death" {
		return "death it is"
	}
	return "sorry, all out of cake... your choice is 'or death'"
}
