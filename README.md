# Advent of Code 2020

https://adventofcode.com

## Create new day template
``` bash
cd tools
bash newDay.sh "Day-XX-name" "Day XX: description"
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