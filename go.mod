module github.com/AhhMonkeyDevs/discordbeat

go 1.15

replace (
	//github.com/AhhMonkeyDevs/discordgo-lite => ../discordgo-lite
	github.com/Microsoft/go-winio => github.com/bi-zone/go-winio v0.4.15
	github.com/Shopify/sarama => github.com/elastic/sarama v1.19.1-0.20210120173147-5c8cb347d877
	github.com/cucumber/godog => github.com/cucumber/godog v0.8.1
	github.com/docker/docker => github.com/docker/engine v0.0.0-20191113042239-ea84732a7725
	github.com/dop251/goja => github.com/andrewkroh/goja v0.0.0-20190128172624-dd2ac4456e20
	github.com/dop251/goja_nodejs => github.com/dop251/goja_nodejs v0.0.0-20171011081505-adff31b136e6
	github.com/fsnotify/fsnotify => github.com/adriansr/fsnotify v0.0.0-20180417234312-c9bbe1f46f1d
	golang.org/x/tools => golang.org/x/tools v0.0.0-20200602230032-c00d67ef29d0 // release 1.14
)

require (
	github.com/AhhMonkeyDevs/discordgo-lite v0.1.8
	github.com/akavel/rsrc v0.10.1 // indirect
	github.com/containerd/containerd v1.4.3 // indirect
	github.com/dlclark/regexp2 v1.4.0 // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/dop251/goja v0.0.0-20210225094849-f3cfc97811c0 // indirect
	github.com/dop251/goja_nodejs v0.0.0-20201222133159-1629e8d0b836 // indirect
	github.com/elastic/beats/v7 v7.0.0-alpha2.0.20210225210557-6adbccfdf1d6
	github.com/fatih/color v1.10.0 // indirect
	github.com/go-sourcemap/sourcemap v2.1.3+incompatible // indirect
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/josephspurrier/goversioninfo v1.2.0 // indirect
	github.com/magefile/mage v1.11.0
	github.com/mitchellh/gox v1.0.1
	github.com/mitchellh/hashstructure v1.1.0 // indirect
	github.com/pierrre/gotestcover v0.0.0-20160517101806-924dca7d15f0
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/tsg/go-daemon v0.0.0-20200207173439-e704b93fd89b
	go.elastic.co/apm v1.11.0 // indirect
	go.elastic.co/ecszap v1.0.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/lint v0.0.0-20201208152925-83fdc39ff7b5
	golang.org/x/net v0.0.0-20210224082022-3d97a244fca7 // indirect
	golang.org/x/sys v0.0.0-20210225134936-a50acf3fe073 // indirect
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools v0.1.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/tools v0.1.2 // indirect
	howett.net/plist v0.0.0-20201203080718-1454fab16a06 // indirect
)
