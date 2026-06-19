package hosts

import (
	"os"
	"path/filepath"
	"testing"
)

// Test the normalization function here (add long examples later)
func TestNormalizeName(t *testing.T) {
	tests := map[string]string{
		"ns1.lab":   "ns1.lab.",
		"NS1.LAB":   "ns1.lab.",
		"ns1.lab.":  "ns1.lab.",
		" ns1.lab ": "ns1.lab.",
	}

	for input, want := range tests {
		got := NormalizeName(input)
		if got != want {
			t.Fatalf("NormalizeName(%q) = %q, want %q", input, got, want)
		}
	}
}

// Check hosts file loading
func TestLoadHostsFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "hosts.txt")

	content := `
# comment
127.0.0.1 localhost
192.168.10.10 ns1.lab ns1
2001:db8::1 ipv6.lab
`

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	h, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}

	ips, ok := h.LookupA("ns1.lab")
	if !ok {
		t.Fatal("expected ns1.lab lookup to succeed")
	}

	if got, want := ips[0].String(), "192.168.10.10"; got != want {
		t.Fatalf("got %s, want %s", got, want)
	}

	if _, ok := h.LookupA("ipv6.lab"); ok {
		t.Fatal("expected IPv6-only host not to return A record")
	}
}
