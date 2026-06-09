package main

import (
	"fmt"
	"net"
	"os"

	"github.com/alecthomas/kong"

	"gobind/internal/hosts"
)

type CLI struct {
	Hosts HostsCmd `cmd:"" help:"Work with hosts.txt-style name lookup."`
}

type HostsCmd struct {
	Lookup HostsLookupCmd `cmd:"" help:"Look up a name in a hosts file."`
}

type HostsLookupCmd struct {
	HostsPath string `name:"hosts" default:"testdata/hosts.txt" help:"Path to hosts file."`
	Name      string `arg:"" help:"Name to look up."`
}

func (cmd *HostsLookupCmd) Run() error {
	h, err := hosts.Load(cmd.HostsPath)
	if err != nil {
		return err
	}

	ips, ok := h.LookupA(cmd.Name)
	if !ok {
		return fmt.Errorf("no A record found for %s", hosts.NormalizeName(cmd.Name))
	}

	for _, ip := range ips {
		fmt.Printf("%s A %s\n", hosts.NormalizeName(cmd.Name), formatIPv4(ip))
	}

	return nil
}

func formatIPv4(ip net.IP) string {
	v4 := ip.To4()
	if v4 == nil {
		return ip.String()
	}

	return net.IP(v4).String()
}

func main() {
	var cli CLI

	ctx := kong.Parse(
		&cli,
		kong.Name("gobind"),
		kong.Description("A tiny history-shaped DNS learning project in Go."),
		kong.UsageOnError(),
	)

	if err := ctx.Run(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
