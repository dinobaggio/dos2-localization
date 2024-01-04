package libs

func VarUse(vals ...interface{}) {
	for _, val := range vals {
		_ = val
	}
}
