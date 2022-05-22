package LogicProcessors

import (
	"errors"
	"os"
)

// Validate функция, предназначена для валидации (проверки существования) корня в
// локальной системе, возвращает ошибку, если валидация не удалась, либо
// результат отрицательный, либо указан путь до файла
func Validate(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !info.IsDir() {
		return errors.New("ошибка. Указанный файл не является директорией")
	}
	return nil
}
