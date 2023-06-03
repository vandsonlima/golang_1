package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func Gerar() *cli.App {
	/*app := cli.App{
		Name:  "Aplicação teste",
		Usage: "busca ips e nomes do servidor na internet",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "devbook.com.br",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "ip",
				Usage:  "Busca ips de enderecos",
				Action: buscarIps,
			},
		},
	}*/
	app := cli.NewApp()
	app.Name = "Aplicação teste"
	app.Usage = "Busca ips de enderecos"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPs",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIps,
		},
	}
	return app
}

func buscarIps(context *cli.Context) error {
	host := context.String("host")
	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}
