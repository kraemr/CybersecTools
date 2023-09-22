package main;
import "os"
import "fmt"
import "errors"
import "strings"
import "bufio"

var g_import_aliases map[string]string
var g_imported_func_alias map[string]string

/*
if there is no alias --> then key "pickle" --> "pickle"
if there is an alias for pickle called p then key "pickle" --> "p"
*/


// check if vulnerable input is delivered to a function like pickle.loads() or os.system() either directly or indirectly --> "Some String %s" % uservar 
func checkFuncArgs(){

}

// This gets called whenever an import is encountered and adds it to to a global obj
func resolveImports(){

}

// regex maybe ??
func checkForKeywords(str string,keywords []string,line int) bool{
	for i:=0; i< len(keywords); i++ {
		if strings.Contains(str,keywords[i]) {
			return true
		}
	}
	return false
}

func isFormatStringInjectable(str string, isDjango bool) bool {
	splitted := strings.Split(str,"\"")
	if len(splitted) < 3 {
		return false
	}
	if isDjango == true {
		return true
	}
	return strings.Contains(splitted[2],"%") && strings.Contains(splitted[1],"%s")
}

func readFileIntoStringBuf(filestr string) ([]string,error){
    f, err := os.OpenFile(filestr, os.O_RDONLY, os.ModePerm)
	var lines []string
	if err != nil {
        //log.Fatalf("open file error: %v", err)
        return lines,err
    }
    defer f.Close()
    sc := bufio.NewScanner(f)
    for sc.Scan() {
        lines = append(lines, sc.Text()) 
	}
	return lines,sc.Err()
}

func main(){
    argsWithoutProg := os.Args[1:]
	g_import_aliases = make(map[string]string)
	g_imported_func_alias = make(map[string]string)
	g_import_aliases["test"] = "test"
	if len(argsWithoutProg) == 0{
		fmt.Println("file name missing")
		return
	}
	buffer, err := readFileIntoStringBuf(argsWithoutProg[0]) // the file name
	if len(buffer) == 0{
        fmt.Println("empty file")
		return
	}
	if errors.Is(err,errors.New("bad input")){
        fmt.Println("bad input error")
	    return
	}

	//  read in list of function names and formatted strings like: "%sfuncname" where %s is either empty or class. 
	//  before that try resolving imports 
	for i:=0; i<len(buffer); i++{
		line := buffer[i]
		// First check for vuln funcs like os's .system, os.writefile pickle.loads, djangos template.render
	}

	b := isFormatStringInjectable("\"test %s\" % variable",false)
	fmt.Println(b)
	return
}
