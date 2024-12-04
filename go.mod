module github.com/SouthbankSoftware/proofable

go 1.13

require (
	github.com/SouthbankSoftware/provendb-verify v0.0.0-20200708050604-5e179a9af7d3
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/djherbis/times v1.2.0
	github.com/ethereum/go-ethereum v1.10.17
	github.com/fatih/color v1.7.0
	github.com/golang/protobuf v1.4.3
	github.com/google/go-cmp v0.5.4
	github.com/gookit/config/v2 v2.0.17
	github.com/karrick/godirwalk v1.15.3
	github.com/korovkin/limiter v0.0.0-20170610225302-6b837f5fd496
	github.com/manifoldco/promptui v0.6.0
	github.com/otiai10/copy v1.3.0
	github.com/phayes/freeport v0.0.0-20180830031419-95f893ade6f2
	github.com/skratchdot/open-golang v0.0.0-20200116055534-eef842397966
	github.com/spf13/afero v1.2.1 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.6.2
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/net v0.0.0-20210805182204-aaa1db679c0d
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.26.0
	src.techknowlogick.com/xgo v0.0.0-20200514233805-209a5cf70012
)

replace github.com/karrick/godirwalk => github.com/SouthbankSoftware/godirwalk v1.15.4-0.20200319040501-49f7c4f93f76
