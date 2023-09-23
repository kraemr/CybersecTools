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


/*
the flow of the program is as follows:
first we read file into mem
then we read our wordlists
then we analyze the lines:

i.e
p.loads("stuff" % var) 

detection from wordlist: "%s".loads(, pkgname   %s is alias, then we just split on ,
then we just print the line number and the vuln
*/

// check if vulnerable input is delivered to a function like pickle.loads() or os.system() either directly or indirectly --> "Some String %s" % uservar
func checkFuncArgs(){

}
// This gets called whenever an import is encountered and adds it to to a global map
func resolveImport(importstr  string) {
	splitted := strings.Split(importstr," ")
	if len(splitted) < 2 || !strings.Contains(importstr, "import") {
		return
	}
	if splitted[0] == "import" { // we are importing a package
		found_pkg_name := false // if this is a valid pkg import
		has_alias := false
		pkg_name := ""
		alias := ""
		for i:=1; i<len(splitted); i++{
			if splitted[i] != " " && found_pkg_name == false{
				pkg_name = splitted[i]
				found_pkg_name = true
				continue
			}
			if splitted[i] != " " && found_pkg_name && splitted[i] == "as" {
				has_alias = true
				for j:=i; j<len(splitted);j++{
					if splitted[j] != " "{
						alias = splitted[j]
					}
				}
			}
		}
		if has_alias==true{
			g_import_aliases[pkg_name] = alias
		}else {
			g_import_aliases[pkg_name] = pkg_name
		}
		fmt.Println(g_import_aliases[pkg_name])
	}
	
}

// regex maybe ??
// basically youll have a wordlist with format strings, where the formatted part gets replaced with the resolved pkg name
func checkForKeywords(str string,keywords []string,line int) bool{
	for i:=0; i< len(keywords); i++ {
		if strings.Contains(str,keywords[i]) {
			return true
		}
	}
	return false
}

// check if single quote string
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
		resolveImport(line)
		
		// First check for vuln funcs like os's .system, os.writefile pickle.loads, djangos template.render
	}
	//b := isFormatStringInjectable("\"test %s\" % variable",false)
	resolveImport("import pkg as p")
	s := fmt.Sprintf("%sloads(,pickle","p.")
	fmt.Println(s)
	return
}
