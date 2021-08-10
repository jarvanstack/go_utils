## go_utils

### quick start

```bash
go get -u github.com/dengjiawen8955/go_utils
```
see the **_test file to learn how to use.

### utils

#### 1. testu
[test_util](testu/) loop time +  time cost + ops (QPS)

How to use?
```go
package main
import (
"github.com/dengjiawen8955/go_utils/test_util"
"testing"
"time"
)
func Test_test_util(t *testing.T) {
    tu := test_util.NewTestUtil(1000)
    tu.Start()
    for i := 0; i < 1000; i++ {
    time.Sleep(time.Millisecond)
    }
    tu.End()
}
```

#### 2. throw_util

[throw_util](erru) try catch the error can print stack error to help you find where is error happen.

#### 3. string_util

3. [string_util](stringu) (1) get random string,(2) md5 ...

### 4. syscall_util 

expose some unsafe pointer or fd to help you to do syscall call.

```bash
import (
	"fmt"
	"github.com/dengjiawen8955/go_utils/syscall_util"
	"github.com/dengjiawen8955/go_utils/throw_util"
)

func main() {
	var err error
	listen, err := sysu.Listen(9999)
	erru.Throw(err)
	for {
		conn, err := listen.Accept()
		erru.Throw(err)
		go func() {
			defer conn.Close()
			fmt.Printf("conn.ClientFd=%#v\n", conn.ClientFd)
			fmt.Printf("conn.ServerFd=%#v\n", conn.ServerFd)
			conn.Write([]byte("hi,syscall_util"))
		}()
	}
}

```