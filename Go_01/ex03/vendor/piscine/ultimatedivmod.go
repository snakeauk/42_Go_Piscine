package piscine

func UltimateDivMod(a *int, b *int) {
	n := *a
	m := *b
	*a = n / m
	*b = n % m
}
