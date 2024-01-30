package util

import (
	"testing"
)

func TestGetDirAndFiles(t *testing.T) {
	str := `total 80
-rw-r--r--   1 root root    0 Jan 30 13:14 aa.txt
drwxr-xr-x   2 root root 4096 Apr  8  2021 bin
drwxr-xr-x   2 root root 4096 Mar 19  2021 boot
drwxr-xr-x   5 root root  360 Jan 20 12:51 dev
drwxr-xr-x   1 root root 4096 Apr 10  2021 docker-entrypoint.d
-rwxrwxr-x   1 root root 1202 Apr 10  2021 docker-entrypoint.sh
drwxr-xr-x   1 root root 4096 Jan 20 12:51 etc
drwxr-xr-x   2 root root 4096 Mar 19  2021 home
drwxr-xr-x   1 root root 4096 Apr 10  2021 lib
drwxr-xr-x   2 root root 4096 Apr  8  2021 lib64
drwxr-xr-x   2 root root 4096 Apr  8  2021 media
drwxr-xr-x   2 root root 4096 Apr  8  2021 mnt
drwxr-xr-x   2 root root 4096 Apr  8  2021 opt
dr-xr-xr-x 608 root root    0 Jan 20 12:51 proc
drwx------   2 root root 4096 Apr  8  2021 root
drwxr-xr-x   1 root root 4096 Jan 20 12:51 run
drwxr-xr-x   2 root root 4096 Apr  8  2021 sbin
drwxr-xr-x   2 root root 4096 Apr  8  2021 srv
dr-xr-xr-x  13 root root    0 Jan 20 12:51 sys
drwxrwxrwt   1 root root 4096 Apr 10  2021 tmp
drwxr-xr-x   1 root root 4096 Apr  8  2021 usr
drwxr-xr-x   1 root root 4096 Apr  8  2021 var
`

	files := GetDirAndFiles(str)
	for _, file := range files {
		t.Log(file)
	}
}
