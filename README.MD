# Instructions

- ### Clone repository
- ### Create DynamoDB Table

```bash
aws dynamodb create-table --table-name DNA --attribute-definitions AttributeName=id,AttributeType=S  AttributeName=dna_type,AttributeType=S --key-schema AttributeName=id,KeyType=HASH --provisioned-throughput  ReadCapacityUnits=10,WriteCapacityUnits=5 --global-secondary-indexes \
        "[
            {
                \"IndexName\": \"Dna-index\",
                \"KeySchema\":
                    [{\"AttributeName\":\"dna_type\",\"KeyType\":\"HASH\"}],
                \"Projection\":{
                    \"ProjectionType\":\"ALL\"
                },
                \"ProvisionedThroughput\": {
                    \"ReadCapacityUnits\": 10,
                    \"WriteCapacityUnits\": 5
                }
            }
        ]"
```

- For local dyanamoDB add

```bash
--endpoint-url http://localhost:8080
```

- ### Get dependencies

```bash
go get
```

- ### Run Test

```bash
go test -v ./... -cover
```

## Endpoints

- POST: /mutant
  <br>
  Body

```json
{
  "dna": ["ATGCGA", "CAGTGC", "TTATTT", "AGACGG", "GCGTCA", "TCACTG"]
}
```

- Response
  <br>

```json
{
  "status": "ok"
}
```

- GET: /stats
- Response
  <br>

```json
{
  "count_mutant_dna": 1,
  "count_human_dna": 1,
  "ratio": 1
}
```
