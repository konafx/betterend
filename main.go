package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/kelseyhightower/envconfig"
	"github.com/joho/godotenv"
	"github.com/konafx/natalya/loop"
	log "github.com/sirupsen/logrus"
)

type Command *discordgo.ApplicationCommand

type CommandHandler func(*discordgo.Session, *discordgo.InteractionCreate)

var (
	s *discordgo.Session
	commands	[]Command
	commandHandlers		map[string]CommandHandler = map[string]CommandHandler{}
	handlers	[]interface{}
)

func addCommand(c Command, ch CommandHandler) {
	commands = append(commands, c)
	commandHandlers[c.Name] = ch
}

func addHandler(h ...interface{}){
	handlers = append(handlers, h...)
}

type Env struct {
	Guilds			[]string
	BotToken		string `split_words:"true"`
	RemoveCommand	bool `split_word:"true" default:"true"`
}
var env Env

func init() {
	log.SetLevel(log.DebugLevel)
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Infoln(".env file not found")
	}
}

func init() {
	if err := envconfig.Process("discord", &env); err != nil {
		log.Fatal(err.Error())
	}
}

func init() {
	var err error
	s, err = discordgo.New("Bot " + env.BotToken)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	s.AddHandler(ready)

	s.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		if h, ok := commandHandlers[i.Data.Name]; ok {
			h(s, i)
		}
	})

	if err := s.Open(); err != nil {
		fmt.Println(err)
		return
	}
	defer s.Close()

	log.Printf("%#v", s.State)

	if len(env.Guilds) == 0 {
		log.Debugf("Create global command")
		for _, v := range commands {
			log.Println(v)
			cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
			if err != nil {
				log.Fatalf("Cannot create '%+v' command: %v", err, v.Name)
			}
			v.ID = cmd.ID
		}
	}

	type key struct {
		k1, k2 string
	}
	appcmds := make(map[key]string)

	for _, x := range env.Guilds {
		log.Debugf("Create guild command")
		for _, v := range commands {
			log.Println(x, v)
			cmd, err := s.ApplicationCommandCreate(s.State.User.ID, x, v)
			if err != nil {
				log.Fatalf("Cannot create '%+v' command: %v", err, v.Name)
			}
			v.ID = cmd.ID
			appcmds[key{x, v.Name}] = cmd.ID
		}
	}

	for _, v := range handlers {
		s.AddHandler(v)
	}

	fmt.Println("bot is running now. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	fmt.Println("Goodbye!")
	if !env.RemoveCommand { return }

	if len(env.Guilds) == 0 {
		for _, v := range commands {
			if err := s.ApplicationCommandDelete(s.State.User.ID, "", v.ID); err != nil {
				log.Errorf("Skip delete cmd: %s (ID: %s)", v.Name, v.ID)
				log.Error(err)
			}
		}
	}
	for _, x := range env.Guilds {
		for _, v := range commands {
			cmdID := appcmds[key{x, v.Name}]
			if err := s.ApplicationCommandDelete(s.State.User.ID, x, cmdID); err != nil {
				log.Errorf("Skip delete cmd: %s (ID: %s) on guild %s", v.Name, cmdID, x)
				log.Error(err)
			}
		}
	}

	return
}

func ready(s *discordgo.Session, e *discordgo.Ready) {
	// s.UpdateGameStatus(0, "Dancing!")
}
