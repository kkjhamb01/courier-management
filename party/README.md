# Party

User Profile Management

### 1. Start Server
`courier-management start party
`

### 2. Protobuf
1. Install grpc-go
2. Enable go111module
   - `export GO111MODULE=on`
3. 
    - `go get google.golang.org/protobuf/cmd/protoc-gen-go google.golang.org/grpc/cmd/protoc-gen-go-grpc`
4. Generate structs and services 
```
cd party/proto
protoc   -I .   -I ${GOPATH}/src  --go_out=":."  party.proto
protoc   -I .   -I ${GOPATH}/src   -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate   --go_out=":."   --validate_out="lang=go:." --go_opt=paths=source_relative --go-grpc_out=./ --go-grpc_opt=paths=source_relative  party.proto
```

### 3. GRPC-UI
Run:
```shell
grpcui -plaintext -v localhost:8085
```

### 4. Request Samples
## 4.1. Create User
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "firstName": "Behnam",
   "lastName": "Nikbakht",
   "email": "behnam.nikbakht@gmail.com",
   "birthDate": "1985",
   "city": "Tehran"
}
```
Response:
```json
{}
```

## 4.2. Update User
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "firstName": "Beh",
   "lastName": "Nik",
   "email": "behnam2.nikbakht@gmail.com",
   "birthDate": "1988",
   "city": "London",
   "transportationType": "TRANSPORTATION_TYPE_CAR"
}
```

Response:
```json
{}
```

## 4.3. Update Profile Additional Info

In these requests, document_ids cannot be updated. They only are updated when user upload a new document, or remove an old one.

### 4.3.1. Update ID Card
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "idCard": {
      "documentIds": [],
      "firstName": "Behnam",
      "lastName": "Nikbakht",
      "number": "000111",
      "expirationDate": "2022",
      "issuePlace": "Iran"
   }
}
```

Response:
```json
{}
```

### 4.3.2. Update Drivers License
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "driversLicense": {
      "documentIds": [],
      "driversNumber": "n1",
      "dvlaCode": "c1"
   }
}
```

Response:
```json
{}
```

### 4.3.3. Update Driver Background
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "driverBackground": {
      "documentIds": [],
      "nationalInsuranceNumber": "098987987"
   }
}
```

Response:
```json
{}
```

### 4.3.4. Update Residence Card
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "residenceCard": {
      "documentIds": [],
      "number": "snnnnnn1",
      "expirationDate": "2030",
      "issueDate": "2020"
   }
}
```

Response:
```json
{}
```

### 4.3.5. Update Bank Account
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "bankAccount": {
      "documentIds": [],
      "bankCreditUnionName": "9999",
      "accountNumber": "8888",
      "accountHolderName": "Behnamnik",
      "sortCode": "987987"
   }
}
```

Response:
```json
{}
```

### 4.3.6. Update Address
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "address": {
      "documentIds": [],
      "street": "st1",
      "building": "bld1",
      "city": "city1",
      "country": "Iran",
      "postCode": "9988",
      "addressDetails": "west"
   }
}
```

Response:
```json
{}
```

## 4.4. Get Profile Additional Info
### 4.4.1. Get ID Card
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "type": "ADDITIONAL_INFO_TYPE_ID_CARD"
}
```

Response:
```json
{
   "id_card": {
      "first_name": "Behnam",
      "last_name": "Nikbakht",
      "number": "000111",
      "expiration_date": "2022",
      "issue_place": "Iran",
      "document_ids": [
         "12f920f2-71dc-9421-82ae-e7d77cbfff53"
      ]
   }
}
```
### 4.4.2. Get Drivers License
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "type": "ADDITIONAL_INFO_TYPE_DRIVERS_LICENSE"
}
```

Response:
```json
{
   "drivers_license": {
      "drivers_number": "n12",
      "dvla_code": "c12",
      "document_ids": []
   }
}
```
### 4.4.3. Get Driver Background
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "type": "ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND"
}
```

Response:
```json
{
  "driver_background": {
    "national_insurance_number": "09898798sss7",
    "document_ids": []
  }
}
```
### 4.4.4. Get Residence Card
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "type": "ADDITIONAL_INFO_TYPE_RESIDENCE_CARD"
}
```

Response:
```json
{
  "residence_card": {
    "number": "snnnnnn12",
    "expiration_date": "2031",
    "issue_date": "2021",
    "document_ids": []
  }
}
```
### 4.4.5. Get Bank Account
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "type": "ADDITIONAL_INFO_TYPE_BANK_ACCOUNT"
}
```

Response:
```json
{
  "bank_account": {
    "bank_credit_union_name": "99991",
    "account_number": "88881",
    "account_holder_name": "Behnamnik1",
    "sort_code": "9879871",
    "document_ids": []
  }
}
```
### 4.4.6. Get Address
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "type": "ADDITIONAL_INFO_TYPE_DRIVER_BACKGROUND"
}
```

Response:
```json
{
  "address": {
    "street": "st1",
    "building": "bld1",
    "city": "London",
    "country": "Iran",
    "post_code": "9988",
    "address_details": "west",
    "document_ids": []
  }
}
```
## 4.5. Upload Document
### 4.5.1 Upload ID Card Document
Request:
```json
{
  "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
  "document": {
    "infoType": "DOCUMENT_INFO_TYPE_ID_CARD",
    "docType": "DOCUMENT_TYPE_PASSPORT",
    "fileType": "text",
    "data": "cGFzc3BvcnQgZG9jdW1lbnQK"
  }
}
```
Response:
```json
{
  "object_id": "12f920f2-71dc-9421-82ae-e7d77cbfff53"
}
```
### 4.5.2. Get Documents of User
Request:
```json
{
  "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ"
}
```
Response:
```json
{
   "documents": [
      {
         "info_type": "DOCUMENT_INFO_TYPE_ID_CARD",
         "doc_type": "DOCUMENT_TYPE_PASSPORT",
         "file_type": "text",
         "download_link_expiration": "1618882659",
         "download_link": "object=12f920f2-71dc-9421-82ae-e7d77cbfff53&exp=1618882659&h=513e8d80c526b8f606032e3bb6b0667c",
         "object_id": "12f920f2-71dc-9421-82ae-e7d77cbfff53"
      },
      {
         "info_type": "DOCUMENT_INFO_TYPE_DRIVERS_LICENSE",
         "doc_type": "DOCUMENT_TYPE_DRIVERS_LICENSE_PHOTO",
         "file_type": "png",
         "download_link_expiration": "1618882659",
         "download_link": "object=14a71c16-fd1a-b53e-e99e-636a5129031a&exp=1618882659&h=eecd8fea1f28a5e0d300ae5bde8d6767",
         "object_id": "14a71c16-fd1a-b53e-e99e-636a5129031a"
      }
   ]
}
```
### 4.5.3. Get Document
Request:
```json
{
  "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
  "objectId": "12f920f2-71dc-9421-82ae-e7d77cbfff53"
}
```
Response:
```json
{
  "download_link_expiration": "1618881757",
  "download_link": "object=12f920f2-71dc-9421-82ae-e7d77cbfff53&exp=1618881757&h=513e8d80c526b8f606032e3bb6b0667c"
}
```
### 4.5.4. Download Document
Request:
```json
{
  "downloadLink": "object=12f920f2-71dc-9421-82ae-e7d77cbfff53&exp=1618881757&h=513e8d80c526b8f606032e3bb6b0667c"
}
```
Resonse (data is in base64 format):
```json
{
  "data": "cGFzc3BvcnQgZG9jdW1lbnQK"
}
```
### 4.5.5. Direct Download Document
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "objectId": "12f920f2-71dc-9421-82ae-e7d77cbfff53"
}
```
Resonse (data is in base64 format):
```json
{
   "data": "cGFzc3BvcnQgZG9jdW1lbnQK"
}
```
### 4.5.6. Delete Document
Request:
```json
{
   "accessToken": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJjb3VyaWVyLW1hbmFnZW1lbnQtdWFhIiwiZGV2aWNlIjoicyIsImV4cCI6MTYxODk0ODc3MywiaWF0IjoxNjE4ODYyMzczLCJpZCI6ImY1NDBkMTEwLTlmNGQtNDVjOS1iYmZlLTEzNTUwYjQyOWFlNSIsImlzcyI6Imh0dHA6Ly8xMjcuMC4wLjE6ODA4MCIsImtpZCI6IjIwMjFqd3QiLCJwaG9uZV9udW1iZXIiOiI5ODkxMjYwMzE3MjQiLCJyb2xlcyI6WyJDT1VSSUVSIl0sInN1YiI6ImJjMGI2MmM3LTU1Y2MtMjY1Ny01NGNiLTcxYmU3ODUzODdmOCIsInR5cCI6ImJlYXJlciJ9.BQh4_2Mwb6-XF0z94lpSA6Ck_0ZvQcdvK1-tZFjjf7REiQ_GGhoOLC_Spj5INgmpARWIzA7rOoZWFV4Puf715aq5jYMTnRW1FPPztPyJyaR_5vr_4vQwjAOuQxCncjFzClu66CuYHFjM8ER6i1WQTjjf02ME1701spN_26uE_dnlgX5ZjogozfN3P_J8Hutux0LPnL3MU29L_Kv9T4frCszPCEVVP6hi229lO_p9MwPxB6zpEKbmGH2iPSw9vT4uxrWwI6G5RF1Eze9h1xz0ZBEBU0SGNnVHcxIT8sttBQmB3UhYL8GSJshoka1wSY60HmSMufvgl9TDNkIMm03ElQ",
   "objectId": "14a71c16-fd1a-b53e-e99e-636a5129031a"
}
```
Resonse (data is in base64 format):
```json
{}
```

Errors

- Canceled = 1
- Unknown = 2
- InvalidArgument = 3
- DeadlineExceeded = 4
- NotFound = 5
- AlreadyExists = 6
- PermissionDenied = 7
- ResourceExhausted = 8
- FailedPrecondition = 9
- Aborted = 10
- OutOfRange = 11
- Unimplemented = 12
- Internal = 13
- Unavailable = 14
- DataLoss = 15
- Unauthenticated = 16
- Others Code = 17