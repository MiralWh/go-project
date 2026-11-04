package works

import (
	"bufio"
	"fmt"
	"go-project/app/helper"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func countErrors(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("открыть %s: %w", path, err)
	}
	defer f.Close()

	count := 0
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		if strings.Contains(line, "error") {
			count++
		}
	}
	if err := scanner.Err(); err != nil {
		return count, fmt.Errorf("чтение %s: %w", path, err)
	}
	return count, nil
}

//----------------------------------------------------------------

// textExtensions — белый список расширений, которые считаем текстовыми.
var textExtensions = map[string]bool{
	".txt": true, ".log": true, ".md": true, ".csv": true,
	".json": true, ".xml": true, ".yaml": true, ".yml": true,
	".go": true, ".html": true, ".css": true, ".js": true,
	".ini": true, ".conf": true, ".cfg": true, ".env": true,
}

func isTextFile(name string) bool {
	ext := strings.ToLower(filepath.Ext(name))
	return ext != "" && textExtensions[ext]
}

func collectFiles(dir string) ([]string, error) {

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("чтение директории: %w", err)
	}
	var files []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		files = append(files, dir+"\\"+e.Name())
	}
	sort.Strings(files)
	return files, nil
}

// func combine(files []string, outPath string) (written int, err error) {
// 	out, err := os.Create(outPath)
// 	if err != nil {
// 		return 0, fmt.Errorf("создать %s: %w", outPath, err)
// 	}

// 	defer func() {
// 		if cerr := out.Close(); err == nil && cerr != nil {
// 			err = fmt.Errorf("закрыть %s: %w", outPath, cerr)
// 		}
// 	}()

// 	w := bufio.NewWriter(out)

// 	for _, name := range files {

// 		if filepath.Clean(name) == outPath {
// 			continue
// 		}
// 		if !isTextFile(name) {
// 			fmt.Fprintf(os.Stderr, "пропуск (не текстовый): %s\n", name)
// 			continue
// 		}

// 		f, err := os.Open(name)
// 		if err != nil {
// 			fmt.Fprintf(os.Stderr, "пропуск %s: %v\n", name, err)
// 			continue
// 		}

// 		if _, err := fmt.Fprintf(w, "=== %s ===\n", name); err != nil {
// 			f.Close()
// 			return written, fmt.Errorf("запись заголовка: %w", err)
// 		}

// 		if _, err := io.Copy(w, f); err != nil {
// 			f.Close()
// 			return written, fmt.Errorf("копирование %s: %w", name, err)
// 		}
// 		fmt.Fprintln(w)

// 		if err := w.Flush(); err != nil {
// 			return written, fmt.Errorf("flush: %w", err)
// 		}

// 		f.Close()

// 		written++
// 	}

// 	// if err := w.Flush(); err != nil {
// 	// 	return written, fmt.Errorf("flush: %w", err)
// 	// }

// 	return written, nil
// }

// combine пишет содержимое files в outPath.
// Возвращает количество успешно обработанных файлов.
// Нетекстовые и нечитаемые файлы пропускаются с предупреждением в stderr.

func combine(files []string, outPath string) (written int, err error) {
	out, err := os.Create(outPath)
	if err != nil {
		return 0, fmt.Errorf("создать %s: %w", outPath, err)
	}

	w := bufio.NewWriter(out)

	defer out.Close()
	defer func() {
		if ferr := w.Flush(); ferr != nil && err == nil {
			err = fmt.Errorf("flush: %w", ferr)
		}
	}()

	for _, name := range files {
		if name == outPath {
			continue
		}
		if !isTextFile(name) {
			fmt.Fprintf(os.Stderr, "пропуск (не текстовый): %s\n", name)
			continue
		}

		f, err := os.Open(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "пропуск %s: %v\n", name, err)
			continue
		}

		if _, err := fmt.Fprintf(w, "=== %s ===\n", name); err != nil {
			f.Close()
			return written, fmt.Errorf("заголовок %s: %w", name, err)
		}

		if _, err := io.Copy(w, f); err != nil {
			f.Close()
			return written, fmt.Errorf("копирование %s: %w", name, err)
		}

		if err := w.WriteByte('\n'); err != nil {
			f.Close()
			return written, fmt.Errorf("разделитель: %w", err)
		}

		f.Close()
		written++
	}

	if err := w.Flush(); err != nil {
		return written, fmt.Errorf("flush: %w", err)
	}

	if err := out.Sync(); err != nil {
		return written, fmt.Errorf("sync: %w", err)
	}
	return written, nil
}

func Work8() {

	fmt.Println("Добро пожаловать на 8-e задание по языку \"Go\"")

	flag, task := helper.GetUserInputInt("Ввведите номер задачи 1 - 2")

	if flag {

		switch task {
		case 1:

			head := `Создайте программу, которая читает лог-файл server.log, подсчитывает количество строк, содержащих слово "error", и выводит это число.`
			fmt.Println(head)

			const path = "\\app\\files\\logfile.log"

			target := helper.GetRootPathCombine(path)
			if target == "" {
				fmt.Print("Что-то пошло не так")
			}
			fmt.Print(target)

			count, err := countErrors(target)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Ошибка:", err)
				os.Exit(1)
			}

			fmt.Println()

			fmt.Printf("Строк со словом \"error\": %d\n", count)

			fmt.Println()

		case 2:

			head := `Напишите программу, которая принимает список имен файлов в текущей директории, объединяет их содержимое и сохраняет результат в новый файл combined.txt (работать только с текстовыми файлами)`
			fmt.Println(head)

			const path = "\\app\\files\\newDirectory"

			const file = "\\app\\files\\newDirectory\\combined.txt"

			target := helper.GetRootPathCombine(path)

			output := helper.GetRootPathCombine(file)

			if target == "" {
				fmt.Print("Что-то пошло не так")
			}

			files, err := collectFiles(target)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ошибка:", err)
				os.Exit(1)
			}
			if len(files) == 0 {
				fmt.Fprintln(os.Stderr, "не найдено файлов для обработки")
				os.Exit(1)
			}

			n, err := combine(files, output)
			if err != nil {
				fmt.Fprintln(os.Stderr, "ошибка:", err)
				os.Exit(1)
			}

			fmt.Println()

			fmt.Printf("Готово: %d текстовых файлов объединены в %s\n", n, output)

			fmt.Println()

		default:
			fmt.Println("Неизвестная задача")
		}
	}
}
