generate openapi with 

npm install @openapitools/openapi-generator-cli -g
openapi-generator-cli generate -i .\openapi.yaml -g go-server -o ./api-generated