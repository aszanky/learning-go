# Chapter 01
### Downloading Go Packages
- If you want to call an external packages, before attach it to your code, you must:
  - Download the packages to your local with:
    ``` go get <the_package_link> ```
  - If you don't do that you will get error, when run your code
    ``` cannot find package .... ```
  - You can delete the packages
    ``` go clean -i -v -x <the_package_link> ```
    it is same with
    ``` rm -rf ~/go/src/<the_package_link> ```

### Printing the output
- You can use ```fmt```
- Or standart output ```io/os```

### About := and =
- ```:=``` is the *short assignment statement*
- ```=``` need var declaration, var is mostly used for declaring global variables

### Reading from standart input
- Using ```bufio``` or ```os```