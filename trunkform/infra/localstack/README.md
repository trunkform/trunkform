# LOCALSTACK

🔴 Largely untested localstack setup

## USAGE

```sh
localstack start -d
TG_TF_PATH=tflocal terragrunt init
TG_TF_PATH=tflocal terragrunt apply
localstack stop
```

## CLI

just some notes on the localstack cli

```sh
cat > index.js <<'EOF'
exports.handler = async (event) => {
    return 'Hello, World!';
};
EOF
zip index.js.zip index.js

localstack start -d
awslocal lambda create-function --function-name test --runtime nodejs22.x --handler index.handler --role arn:aws:iam::000000000000:role/test --zip-file fileb://index.js.zip
{
    "FunctionName": "test",
    "FunctionArn": "arn:aws:lambda:us-east-1:000000000000:function:test",
    "Runtime": "nodejs22.x",
    "Role": "arn:aws:iam::000000000000:role/test",
    "Handler": "index.handler",
    "CodeSize": 232,
    "Description": "",
    "Timeout": 3,
    "MemorySize": 128,
    "LastModified": "2026-03-04T18:57:44.877390+0000",
    "CodeSha256": "nrWrb7LSymJgpzwb8QkSsu7bv+bHQOgGIa+t04nLEaI=",
    "Version": "$LATEST",
    "TracingConfig": {
        "Mode": "PassThrough"
    },
    "RevisionId": "7eed7895-26e1-4b98-8815-d9ffd6ccb70e",
    "State": "Pending",
    "StateReason": "The function is being created.",
    "StateReasonCode": "Creating",
    "PackageType": "Zip",
    "Architectures": [
        "x86_64"
    ],
    "EphemeralStorage": {
        "Size": 512
    }
awslocal lambda create-function-url-config --function-name test --auth-type NONE
{
    "FunctionUrl": "http://nf7s09zhslh0b2qv87diq8tlz3gngktp.lambda-url.us-east-1.localhost.localstack.cloud:4566/",
    "FunctionArn": "arn:aws:lambda:us-east-1:000000000000:function:test",
    "AuthType": "NONE",
    "CreationTime": "2026-03-04T18:59:10.658176+0000"
}
curl -X POST 'http://nf7s09zhslh0b2qv87diq8tlz3gngktp.lambda-url.us-east-1.localhost.localstack.cloud:4566/'
Hello, World!%


awslocal lambda delete-function-url-config --function-name test
awslocal lambda delete-function --function-name test

localstack stop
```
