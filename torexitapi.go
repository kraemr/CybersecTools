package main;
import "net/http"
//import "fmt"
import "io/ioutil"
import "strings"
import "io"
import "strconv"
var lines []string
var ipv4decs []uint32
func ipv4ToDec(ip string) uint32{
	splitted := strings.Split(ip,".")
	if len(splitted) != 4{
		return 1 // Invalid ipv4 address
	}
	var ipdec uint64=0
	for i:=0;i<4;i++{
		octetdec, err := strconv.ParseUint(splitted[i],10,32)
		if(err != nil){
			return 0
		}
		ipdec |= octetdec << (8*i)
	}
	return uint32(ipdec)
}

func binarysearch_ipv4(ipv4decs []uint32,ipstring string) bool{
	ip:=ipv4ToDec(ipstring)
	var left uint32 = 0
	var right uint32= uint32(len(ipv4decs)-1)
	for left <= right{
		middle := left + (right - left) 
		if ipv4decs[middle] == ip {
		return true
		}
		if ipv4decs[middle] > ip{
		right = middle- 1
		}else{
		left = middle + 1
		}
	}
	return false
}


func decToIpv4(ip uint32) string{
        octet1:=strconv.FormatUint(uint64(ip &  (0x000000ff)),      10)
        octet2:= strconv.FormatUint(uint64(ip & (0x0000ff00)) >> 8, 10)
        octet3:= strconv.FormatUint(uint64(ip & (0x00ff0000)) >> 16,10)
        octet4:= strconv.FormatUint(uint64(ip & (0xff000000)) >> 24,10)
	ipstr := octet1 + "." + octet2 + "." + octet3 + "." + octet4
//	fmt.Printf(str)
	return ipstr
}

func readLines() []string {
    bytesRead, _ := ioutil.ReadFile("exitnodes")
    fileContent := string(bytesRead)
    lines := strings.Split(fileContent,"\n")
    return lines
}


func readIpDecs() []uint32{
    lines := readLines()
    var ipdecs []uint32
    for i:=0;i<len(lines);i++{
	ipv4dec :=	ipv4ToDec(lines[i])
	if ipv4dec == 0 ||ipv4dec == 1 {
	}else{
		ipdecs = append(ipdecs,ipv4dec)
    	}

	}
        for n :=len(ipdecs); n > 1; n = n - 1 { // äußere Schleife
         for i := 0; i < n - 1; i = i + 1 { // innere Schleife
         if (ipdecs[i] > ipdecs[i + 1]) {
           b,a:= ipdecs[i],ipdecs[i+1]
	   ipdecs[i],ipdecs[i+1] = a,b
         }
       }
    }

    return ipdecs
}

func getTorIp(w http.ResponseWriter, r *http.Request){
	ip := strings.Trim( r.URL.Query().Get("ip")," ")
	exists:=binarysearch_ipv4(ipv4decs,ip)
	if exists{
		io.WriteString(w,"1")
	}else{
		io.WriteString(w,"0")
	}
}

func main(){
     ipv4decs=readIpDecs()
     http.HandleFunc("/", getTorIp)
     err := http.ListenAndServe(":4343", nil)
     _ = err
}

