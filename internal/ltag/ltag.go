package ltag

const (
	DefaultIPattern = "#"
	DefaultOPattern = "#$"
)

type Operation interface {
	Transform(line string) (string, bool)
}
