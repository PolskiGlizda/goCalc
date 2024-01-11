# goCalc
Ask wolframalpha queries from your terminal
## Instalation
### Linux
#### Grab binary
1. Go to releases
2. Download the newest version
3. Run it from your shell
#### Build from source
``` bash
# Clone repository
git clone https://github.com/PolskiGlizda/goCalc.git
# Compile
cd goCalc
go build
# Run
./goCalc <query>
# Alternatively you can run it without compiling
go run goCalc <query>
```
### Windows
``` PowerShell
# Clone repository
git clone https://github.com/PolskiGlizda/goCalc.git
# Compile
cd goCalc
go build
# Run
./goCalc
# Alternatively you can run it without compiling
go run goCalc <query>
```
### macOS
Probably you can build from source the same way like on Linux but I can't test it

## Usage
``` 
goCalc <query>
```
Make sure your query is treated as single string in your shell. Using quotation marks or apostrophes is recomended for most users.
