# Advent of Code 2023

## Tasks

### run

Environment: GOEXPERIMENT=rangefunc

```sh
ls -d */ | grep -v template | xargs -I % sh -c 'cd %;go run .;'
```

### new

Inputs: DAY

```sh
cp -r ./template/ ./day$DAY
sed -i '' "s/1\./$DAY./g" ./day$DAY/main.go
```
