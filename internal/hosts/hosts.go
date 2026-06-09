package hosts

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type HostsFile struct {
	Entries map[string][]net.IP
}

func Load(path string) (*HostsFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open hosts file: %w", err)
	}
	defer file.Close()

	h := &HostsFile{
		Entries: make(map[string][]net.IP),
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Remove comments.
		if idx := strings.Index(line, "#"); idx >= 0 {
			line = line[:idx]
		}

		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}

		ip := net.ParseIP(fields[0])
		if ip == nil {
			return nil, fmt.Errorf("invalid IP address %q", fields[0])
		}

		for _, name := range fields[1:] {
			name = NormalizeName(name)
			h.Entries[name] = append(h.Entries[name], ip)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan hosts file: %w", err)
	}

	return h, nil
}

func (h *HostsFile) LookupA(name string) ([]net.IP, bool) {
	name = NormalizeName(name)

	ips, ok := h.Entries[name]
	if !ok {
		return nil, false
	}

	var result []net.IP

	for _, ip := range ips {
		v4 := ip.To4()
		if v4 == nil {
			continue
		}

		result = append(result, v4)
	}

	return result, len(result) > 0
}

func NormalizeName(name string) string {
	name = strings.TrimSpace(name)
	name = strings.ToLower(name)
	name = strings.TrimSuffix(name, ".")
	return name + "."
}
