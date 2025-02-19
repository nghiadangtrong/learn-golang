#

## Read Source code

- [7-open-source-projects-you-should-know-go-edition](https://dev.to/this-is-learning/7-open-source-projects-you-should-know-go-edition-3bo4)

## Tutorial build

- [Golang](https://www.youtube.com/watch?v=uCR_A-Bphl0&list=PLJbE2Yu2zumCe9cO3SIyragJ8pLmVv0z9)
- [Building a Blockchain](https://www.youtube.com/watch?v=mYlHT9bB6OE&list=PLJbE2Yu2zumC5QE39TQHBLYJDB2gfFE5Q)
- [Building a Basic RPC Server and client with go](https://www.youtube.com/watch?v=1MPWPq2N768&list=PLJbE2Yu2zumAixEws7gtptADSLmZ_pscP)

### Detail
## Install golang

```bash
LATEST_GO_VERSION=$(curl -sL https://go.dev/VERSION?m=text)
LATEST_GO_VERSION=$(echo $LATEST_GO_VERSION | awk '{ print $1 }')
GO_ARCH=$(if [ "$(uname -m)" = "x86_64" ]; then echo "amd64"; elif [ "$(uname -m)" = "aarch64" ] || [ "$(uname -m)" = "arm64" ]; then echo "arm64"; else echo "Kiến trúc hệ thống không được hỗ trợ: $(uname -m)" && exit 1; fi)
cd ~
curl -fsSL -O https://go.dev/dl/${LATEST_GO_VERSION}.linux-${GO_ARCH}.tar.gz
GO_FILE_INSTALL=$(echo ${LATEST_GO_VERSION}.linux-${GO_ARCH}.tar.gz)
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf ${GO_FILE_INSTALL}
grep -q "export PATH=\$PATH:/usr/local/go/bin" ~/.bashrc || echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
go version

```

## Create Project

Để khởi tạo và chạy một project trong Golang, bạn có thể làm theo các bước sau:

1. Cài đặt Golang từ trang chủ: https://golang.org/dl/  
2. Tạo một thư mục mới cho project của bạn và chuyển vào thư mục đó.

3. Khởi tạo một module Go bằng lệnh:
   ```sh
   go mod init my-go-project
   ```
   Lệnh này sẽ tạo file go.mod chứa thông tin về module.

4. Tạo file main.go với nội dung sau (dưới đây là file hoàn chỉnh):

```go title=main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

5. File go.mod được tạo ra sẽ có nội dung tương tự như sau (bạn có thể tạo file này nếu muốn tinh chỉnh hoặc xem thông tin version):

```go title=go.mod
module my-go-project

go 1.20
```

6. Chạy project bằng cách sử dụng lệnh:
   ```sh
   go run main.go
   ```
   Bạn sẽ thấy dòng chữ "Hello, World!" được in ra.

7. Nếu muốn biên dịch chương trình thành file thực thi, sử dụng lệnh:
   ```sh
   go build
   ```
   Sau khi biên dịch, file thực thi có tên là `my-go-project` (hoặc `my-go-project.exe` trên Windows) sẽ được tạo ra. Chạy file thực thi như sau:
   ```sh
   ./my-go-project
   ```

Ngoài ra, bạn có thể mở rộng cấu trúc project với các thư mục như `cmd/` và `pkg/` nhằm quản lý mã nguồn khi project phát triển. Ví dụ:

Cấu trúc thư mục:
```
my-go-project/
├── go.mod
├── cmd/
│   └── my-go-project/
│       └── main.go
└── pkg/
    └── somepackage/
        └── somefile.go
```

Ví dụ file main.go trong thư mục cmd/my-go-project/:

```go title=cmd/my-go-project/main.go
package main

import (
    "fmt"
    "my-go-project/pkg/somepackage"
)

func main() {
    fmt.Println("Hello, World!")
    somepackage.SomeFunction()
}
```

Và file somepackage/somefile.go:

```go title=pkg/somepackage/somefile.go
package somepackage

import "fmt"

func SomeFunction() {
    fmt.Println("Some function in package somepackage")
}
```

Như vậy, bạn có thể dễ dàng tổ chức và phát triển project Golang của mình. Chúc bạn thành công!
