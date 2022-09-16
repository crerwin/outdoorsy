#!/bin/bash

platforms=(
    "linux/amd64" 
    "linux/arm64" 
    "darwin/amd64" 
    "darwin/arm64" 
    "windows/amd64"
)

for platform in ${platforms[@]}
do
    echo "Building outdoorsy for platform ${platform}"
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name="outdoorsy-${GOOS}-${GOARCH}"
    if [[ ${GOOS} = "windows" ]]; then
        output_name+=".exe"
    fi

    GOOS=${platform_split[0]} GOARCH=${platform_split[1]} go build -o bin/${output_name} cmd/outdoorsy/outdoorsy.go
done