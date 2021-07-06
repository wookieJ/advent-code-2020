<p align="center">
 <img width="600" src="logo.jpg">
</p>

# Advent of Code 2020
https://adventofcode.com/2020

## Create new day template
There is a script, which creates new golang module for current day. It demands three input arguments:
1. Your AoC session cookie
2. Number of new day

Session cookie needs to be set for input and name of day downloading. You need to have aocdl go package:
```.env
go get -u github.com/GreenLightning/advent-of-code-downloader/aocdl
```

Example script running:

``` bash
cd tools
bash newDay.sh "<your_AoC_session_cookie>" 1
``` 

Remember to sync new go module (`go.mod [PPM]` -> `Sync Go Module`)

## Settings for IDE
### Running main
Try to run main func. If not working try to  edit configurations.
Set *Run kind* to `Package`, define package path (module name) and working 
directory to include module. For example:

* Package path: `Day-XX-template`
* Working directory: `<pwd>/advent-code-2020/Day-XX-template`

### Running tests
Make sure that you have `Go -> Go Modules -> Enable Go Modules integration` option enabled.
