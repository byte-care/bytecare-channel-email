go build -o main
zip function.zip main
aws lambda update-function-code --function-name bytecare-channel-email --zip-file fileb://function.zip