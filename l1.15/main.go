package main

//Глобальные переменные часто являются плохой практикой, поэтому лучше делать билдер
//var justString string

// Если бы нам была нужна очень большая стринга, а используем мы только 100, то gc не смог бы очистить память память от v
func createHugeString(len int) string {

	return string(make([]byte, len))
}

func someFunc() string {
	v := createHugeString(1 << 10)

	return string([]byte(v[:100]))
}

// что происходит с juststring?
// в JustString передается указатель до 100 элементов
func main() {
	someFunc()
}
