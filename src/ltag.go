package ltag

type Operation interface {
	Transform(line string) (string, bool)
}
