### Requirements
#### Redis
- set the following config (to enable pubsub on expiration): `config set notify-keyspace-events Ex`

### Project Structure
- `api` To handle API
- `business` To handle business rules
- `storage` To handle storage data on DB/File/GCS/...
- `db` To handle connections to sql/nosql databases
- `pubsub` To handle a service-internal eventbus

#### Models
`offering-grpc/proto/models.proto` contains all offering models.

#### Proto
To generate the proto files for a new language: please add the generated files under `offering-grpc/{language_name}`. For instance, `offering-grpc/java`

### Test API
Please refer to [test data file](./TEST_DATA.md).