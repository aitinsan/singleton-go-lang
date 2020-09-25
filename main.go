package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)
type Database interface {
	Getmark(name string) int
}
type singletonDatabase struct {
	students map[string] int
}
func (db *singletonDatabase) Getmark(name string) int  {
	return db.students[name]
}
func readData(path string) (map[string]int, error) {
	file, err := os.Open( path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := map[string]int{}
	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}
	return result, nil
}
var once sync.Once
var instance Database
func GetSingletonDatabase() Database  {
	once.Do(func() {
		caps,_ := readData("./students.txt")
		db := singletonDatabase{caps}
		instance = &db
	})
	return instance
}
// Dip
func GetMark(db Database,names []string) int {
	result := 0
	for _, name := range names {
		result += db.Getmark(name)
	}
	return result
}


func main() {
	db := GetSingletonDatabase()
	mark := db.Getmark("Insan")
	fmt.Println("Insan has a grade",mark)




}