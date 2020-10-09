package files

var GlobalConfig *Config

func LoadConfig() {

	// user, err := user.Current()
	// if err != nil {
	// 	log.Fatalf(err.Error())
	// }

	// data, err := ioutil.ReadFile(user.HomeDir + "/search.json")
	// if err != nil {
	// 	log.Println("Could not find config file ...")
	// 	os.Exit(1)
	// }
	GlobalConfig = new(Config)
	// err = json.Unmarshal(data, GlobalConfig)
	// if err != nil {
	// 	log.Println("Could not read/parse the config file ...")
	// 	os.Exit(1)
	// }
	GlobalConfig.Ignore = []string{".gitignire", ".exe", ".git"}
	GlobalConfig.MaxFileSize = 2

}
