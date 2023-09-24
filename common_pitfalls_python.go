package main;
import "os"
import "fmt"
import "errors"
import "strings"
import "bufio"
var g_import_aliases map[string]string
var g_imported_func_cwe_mapped map[string]string // this maps a vulnerable function name to a cwe 

/*
if there is no alias --> then key "pickle" --> "pickle"
if there is an alias for pickle called p then key "pickle" --> "p"
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
// i fucking hate the python import system, jesus
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
	}
}

// regex maybe ??
// structure message returned so its parsable with jq
// include CWE number and link
// give short description
// line and file that it occurs and why (format string injection, possible user input injection ...)
func checkForKeywords(str string,keywords []string,line int) (bool,string){
	for i:=0; i< len(keywords); i++ {
		if strings.Contains(str,keywords[i]) { // also check here if user input is supplied
			return true,g_imported_func_cwe_mapped[keywords[i]]
		}
	}
	return false,""
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

// This parses a given function detection file
func parseFunctionDetection(path string) ([]string,error){
	lines,err := readFileIntoStringBuf(path)
	if err != nil{
		return nil , errors.New("couldnt parse functions file")
	}
	var parsed_lines []string 
	for i:=0; i<len(lines); i++{
		splitted := strings.Split(lines[i],",")
		alias := ""
		if g_import_aliases[splitted[1]] == "" {
			alias = splitted[1]
		}else{
			alias = g_import_aliases[splitted[1]]	
		}
		fmt_func_str := strings.TrimSpace(fmt.Sprintf(splitted[0],alias))
		if len(splitted) < 3{

		}else{
			g_imported_func_cwe_mapped[fmt_func_str] = splitted[2]
		}
		parsed_lines = append(parsed_lines,fmt_func_str)
	}
	return parsed_lines,nil
}

func readFileIntoStringBuf(filestr string) ([]string,error){
    f, err := os.OpenFile(filestr, os.O_RDONLY, os.ModePerm)
	var lines []string
	if err != nil {
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
	g_imported_func_cwe_mapped = make(map[string]string)
	if len(argsWithoutProg) == 0 {
		fmt.Println("file name missing")
		return
	}
	buffer, err := readFileIntoStringBuf(argsWithoutProg[0]) // the file name
	if len(buffer) == 0 {
        fmt.Println("empty file")
		return
	}
	if errors.Is(err,errors.New("bad input")) {
        fmt.Println("bad input error")
	    return
	}
	parsed_funcs,err := parseFunctionDetection("dangerous_python_funcs.txt")
	found,message := checkForKeywords("os.system",parsed_funcs,0)
	for i:=0; i<len(buffer); i++{
		line := buffer[i]
		resolveImport(line)
		found,message = checkForKeywords(line,parsed_funcs,i)
		if(found){
			fmt.Printf("%s line:%d file:%s\n",message,i,argsWithoutProg[0]) // try if this output is parseable with jq
		}
	}
}
