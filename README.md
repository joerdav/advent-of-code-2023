# Advent of Code 2023

## Tasks

### run

Environment: GOEXPERIMENT=rangefunc

```sh
ls -d */ | grep day | xargs -I % sh -c 'cd %;go run .;'
```

### new

Inputs: DAY

```sh
cp -r ./template/ ./day$DAY
sed -i '' "s/Print.1,/Print($DAY,/g" ./day$DAY/main.go
```
