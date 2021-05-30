## go_utils

### quick start

```bash
go get github.com/dengjiawen8955/go_utils
```
see the **_test file to learn how to use.

### utils

#### 1. testu
[test_util](test_util/) loop time +  time cost + ops (QPS)

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

[throw_util](throw_util) try catch the error can print stack error to help you find where is error happen.

#### 3. string_util

3. [string_util](string_util) (1) get random string,(2) md5 ...
