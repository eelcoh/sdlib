package sdlib

import "os"

var serviceDiscoveryURL = os.Getenv("SERVICE_DISCOVERY_SERVICE_HOST")
var ip = os.Getenv("MY_POD_IP")
