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
import "os/exec"
import "runtime"

func main() {
  myos := runtime.GOOS;
  myarch := runtime.GOARCH;
  var chrome = "";
  var cmdopen *exec.Cmd;

  if myos == "windows" {
    if myarch == "386" {
      chrome = "C:/Program Files/Attendedbyhumans/Eid.jar";
    } else {
      chrome = "C:/Program Files (x86)/Attendedbyhumans/Eid.jar";
    }          

    cmdopen = exec.Command("java", "-cp", chrome, "eid.Eid");
    err := cmdopen.Start();
    if err != nil {
      println("Failed: ", err);
    } 

  } else {
    chrome = "/root/Eid.jar";       
    cmdopen = exec.Command("java", "-cp", chrome, "eid.Eid");
    err := cmdopen.Start();
    if err != nil {
      println("Failed: ", err);
    } 

  } 


}