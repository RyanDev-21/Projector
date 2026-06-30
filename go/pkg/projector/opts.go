package projector

import "github.com/hellflame/argparse"

type Opts struct {
	Args   []string
	Config string
	Pwd    string
}

func GetOpts() (*Opts, error) {
	parse := argparse.NewParser("Projector", "gets all the values", &argparse.ParserConfig{DisableDefaultShowHelp: true})

	args := parse.Strings("a", "args", &argparse.Option{
		Positional: true,
		Required:   false,
		Default:    "",
	})

	config := parse.String("c", "config", &argparse.Option{
		Required: false,
		Default:  "",
	})

	pwd := parse.String("p", "pwd", &argparse.Option{
		Required: false,
		Default:  "",
	})

	err := parse.Parse(nil)
	if err != nil {
		return nil, err
	}
	return &Opts{
		Args:   *args,
		Config: *config,
		Pwd:    *pwd,
	}, nil
}
