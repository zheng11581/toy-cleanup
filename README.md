1. cleanup 删除指定目录下的指定后缀文件，并保留N天内的，目前只删除一层目录下文件，不会递归子目录
2. Usage of ./cleanup:
-before int
需要保留多少天内的文件 (default 30)
-directory string
需要删除文件的目录 (default "/This/is/a/dir")
-regexp string
需要删除文件的后缀 (default "pdf")
3. 放到服务器上
chmod +x cleanup
4. ./cleanup -directory /This/is/a/dir -regexp pdf -before 30
删除/This/is/a/dir目录下，pdf后缀的文件，并保留30天内的

