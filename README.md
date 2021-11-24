# Shrimping-client
A client cli for accessing shrimping platform with GRPC protocol, see our [website](https://get-shrimping.footprint-ai.com) for more details.

#### Terminalogy
- [Shrimping for Shopee](https://github.com/FootprintAI/shrimping-client/tree/main/shrimping)
- [Shritgram for Instagram](https://github.com/FootprintAI/shrimping-client/tree/main/shritagram)


### Run with docker

```
docker run \
-it footprintai/shrimping-grpcclient:v1 \
/out/shrimping.linux help

```

### Run with Binary

We now support Linux/Windows/Darwin platform with Golang's Cross-Compiled binary, you cna download them from folders.

```
./shrimping.linux help

Usage:
  shrimping [command]

Available Commands:
  category    lookup items by category name
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  itemids     lookup items by item ids
  itemurls    lookup items by item urls
  login       login with email & password
  version     show client veriosn

Flags:
      --cfg string           config file (default is ./shrimpingcfg) (default "./shrimpingcfg")
      --format string        output format (default: table), possible value: csv / table (default "table")
  -h, --help                 help for shrimping
      --svr_address string   data platform server address (default is 172.17.0.1:50050) (default "172.17.0.1:50050")


```

##### Login
Please use email/password that we provided to you to access our platform, the access toke (with one day expiry duration) will be kept in local configuration file

```
./shrimping.linux login --email hello@example.com --password thisisnotarealpassword

```

#### Feedbacks

Please open issue to report any feedbacks or features you wanted.

