# Advent of Code 2020

https://adventofcode.com/2020

## Create new day template
There is a script, which creates new golang module for current day. It demands three input arguments:
1. Module name
2. Short description
3. Your AoC session cookie

Session cookie needs to be set for input downloading

``` bash
cd tools
bash newDay.sh "Day-XX-name" "Day XX: description" "<your_AoC_session_cookie>"
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