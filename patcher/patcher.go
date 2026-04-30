package patcher

// Patcher defines the standard behavior for all ROM patchers
type Patcher interface {
	Apply(baseFile, patchFile, outFile string) error
}
