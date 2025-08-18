package src

func HelloWorld(name string) (string, error) {

	if name == "" {
		panic("Name Params Cannot Be Empty!")
	}

	return "Hi " + name, nil
}
