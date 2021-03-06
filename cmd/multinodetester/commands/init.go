package commands

import (
	"fmt"
	"github.com/coschain/cobra"
	"github.com/coschain/contentos-go/common"
	"github.com/coschain/contentos-go/common/constants"
	"github.com/coschain/contentos-go/config"
	"github.com/coschain/contentos-go/node"
	"github.com/coschain/contentos-go/prototype"
	"os"
	"path/filepath"
	"strconv"
)

var InitCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init count(default 4)",
		Short: "Initialize configuration files for multi cosd",
		Run:   initConf,
	}
	cmd.Flags().StringVarP(&chainName, "chain", "c", "", "chain name [main/test/dev], default is main")
	return cmd
}

var p2pPortStart = 20200
var maxSeedNodeCount = 4
var seeds []string

func addConf(confdir string, cfg node.Config, index int) {
	var err error

	fmt.Println("config Dir: ", confdir)
	if _, err = os.Stat(confdir); os.IsNotExist(err) {
		if err = os.MkdirAll(confdir, 0700); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	//logging.CLog()
	cfg.GRPC.RPCListen = fmt.Sprintf("127.0.0.1:%d", 8888+index)
	cfg.GRPC.HTTPListen = fmt.Sprintf("127.0.0.1:%d", 8080+index)

	cfg.P2P.NodePort = uint(p2pPortStart + index*2 + 0)
	cfg.P2P.NodeConsensusPort = uint(p2pPortStart + index*2 + 1)
	cfg.P2P.Genesis.SeedList = seeds
	cfg.HealthCheck.Port = fmt.Sprintf("%d", 9090+index)
	cfg.Database.Driver = "mysql"
	cfg.Database.Db = "contentosdb"
	cfg.Database.User = "contentos"
	cfg.Database.Password = "123456"

	if index > 0 {
		cfg.Consensus.LocalBpName = fmt.Sprintf("%s%d", constants.COSInitMiner, index)
		key, err := prototype.FixBytesToPrivateKey([]byte(cfg.Consensus.LocalBpName))
		if err != nil {
			panic(err)
		}
		cfg.Consensus.LocalBpPrivateKey = key.ToWIF()
	}

	err = config.WriteNodeConfigFile(confdir, "config.toml", cfg, 0600)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
func initConf(cmd *cobra.Command, args []string) {
	var nodeCount int = 4
	if len(args) > 0 {
		cnt, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		nodeCount = cnt
	}

	initConfByCount(nodeCount)
}

func initConfByCount(nodeCount int) {
	seedCount := nodeCount
	if seedCount > maxSeedNodeCount {
		seedCount = maxSeedNodeCount
	}

	for i := 0; i < seedCount; i++ {
			seeds = append(seeds, fmt.Sprintf("127.0.0.1:%d", i*2+p2pPortStart))
	}

	fmt.Println("Seed nodes list: ", seeds)

	if len(chainName) == 0 {
		chainName = common.ChainNameMainNet
	}

	for i := 0; i < nodeCount; i++ {
		cfg := config.DefaultNodeConfig
		cfg.MinDiskSpaceInGB = 0
		cfg.Name = fmt.Sprintf("%s_%d", TesterClientIdentifier, i)
		cfg.ChainId = chainName
		if i > 0 {
			cfg.Consensus.BootStrap = false
		}
		confDir := filepath.Join(cfg.DataDir, cfg.Name)
		addConf(confDir, cfg, i)
	}
}
