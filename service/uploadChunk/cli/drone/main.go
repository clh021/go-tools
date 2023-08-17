package drone

import (
	"fmt"

	"gitee.com/linakesi/source-analysis-tools-ui/cli/config"
	"github.com/drone/drone-go/drone"
	"golang.org/x/oauth2"
)

type DroneService struct {
	conf *config.DroneConfig
}

func New(c config.DroneConfig) *DroneService {
	d := &DroneService{
		conf: &c,
		// client: 
	}
	return d
}

func (d DroneService) Client() {
	// create an http client with oauth authentication.
	config := new(oauth2.Config)
	auther := config.Client(
		oauth2.NoContext,
		&oauth2.Token{
			AccessToken: d.conf.Token,
		},
	)

	// create the drone client with authenticator
	client := drone.NewClient(d.conf.Host, auther)

	// gets the current user
	user, err := client.Self()
	fmt.Println(user, err)

	// gets the named repository information
	repo, err := client.Repo(d.conf.RepoNamespace, d.conf.RepoName)
	fmt.Println(repo, err)
}

func (d DroneService) Build() {
	fmt.Println("build")
}