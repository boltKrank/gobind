#!/opt/homebrew/bin/bash
echo "Testing hosts.txt lookup"
echo "Testing ns1.lab"
cd ..
go run ./cmd/gobind hosts lookup ns1.lab
echo "Testing missing.lab"
go run ./cmd/gobind hosts lookup missing.lab
