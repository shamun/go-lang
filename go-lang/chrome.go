/*

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

$ go build -o id.exe source.go
$ go run source.go

*/

package main
import "os"
import "os/exec"
import "runtime"
import "encoding/json"

type Configuration struct {
  main []string
  name []string
  window []string
}

func main() {
  myos := runtime.GOOS;
  myarch := runtime.GOARCH;
  var chrome = "";
  var cmdopen *exec.Cmd;

  if myos == "windows" {
    if myarch == "386" {
      chrome = "C:/Program Files (x86)/Google/Chrome/Application/chrome.exe";
    } else {      
      //chrome = "C:/Program Files/Google/Chrome/Application/chrome.exe";
      chrome = "C:/Program Files (x86)/Google/Chrome/Application/chrome.exe";
    }          

    // Read config
    file, _ := os.Open("C:/Program Files (x86)/Attendedbyhumans/package.json");
    decoder := json.NewDecoder(file);
    configuration := Configuration{};
    err := decoder.Decode(&configuration);
    if err != nil {
      println("error: ", err);
    }

    println(configuration.main);

    // Execute
    cmdopen = exec.Command(chrome, "--app=http://icanhazip.com");
    err1 := cmdopen.Start();
    if err1 != nil {
      println("Failed: ", err1);
    } 

  } else {
    println("Incompatible");
  } 
}