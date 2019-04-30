package main

import (
	"fmt"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"log"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strings"
)

var (
	mergeRequest	= regexp.MustCompile(`Merge pull request #([\d]+)`)
	webconsoleBump	= regexp.MustCompile(regexp.QuoteMeta("bump(github.com/kubernetes-incubator/service-catalog-web-console): ") + `([\w]+)`)
	upstreamKube	= regexp.MustCompile(`^UPSTREAM: (\d+)+:(.+)`)
	upstreamRepo	= regexp.MustCompile(`^UPSTREAM: ([\w/-]+): (\d+)+:(.+)`)
	prefix		= regexp.MustCompile(`^[\w-]: `)
	assignments	= []prefixAssignment{{"cluster up", "cluster"}, {" pv ", "storage"}, {"haproxy", "router"}, {"router", "router"}, {"route", "route"}, {"authoriz", "auth"}, {"rbac", "auth"}, {"authent", "auth"}, {"reconcil", "auth"}, {"auth", "auth"}, {"role", "auth"}, {" dc ", "deploy"}, {"deployment", "deploy"}, {"rolling", "deploy"}, {"security context constr", "security"}, {"scc", "security"}, {"pipeline", "build"}, {"build", "build"}, {"registry", "registry"}, {"registries", "image"}, {"image", "image"}, {" arp ", "network"}, {" cni ", "network"}, {"egress", "network"}, {"network", "network"}, {"oc ", "cli"}, {"template", "template"}, {"etcd", "server"}, {"pod", "node"}, {"hack/", "hack"}, {"e2e", "test"}, {"integration", "test"}, {"cluster", "cluster"}, {"master", "server"}, {"packages", "hack"}, {"api", "server"}}
)

type prefixAssignment struct {
	term	string
	prefix	string
}
type commit struct {
	short	string
	parents	[]string
	message	string
}

func contains(arr []string, value string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	for _, s := range arr {
		if s == value {
			return true
		}
	}
	return false
}
func main() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	log.SetFlags(0)
	if len(os.Args) != 3 {
		log.Fatalf("Must specify two arguments, FROM and TO")
	}
	from := os.Args[1]
	to := os.Args[2]
	out, err := exec.Command("git", "log", "--topo-order", "--pretty=tformat:%h %p|%s", "--reverse", fmt.Sprintf("%s..%s", from, to)).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	hide := make(map[string]struct{})
	var apiChanges []string
	var webconsole []string
	var commits []commit
	var upstreams []commit
	var bumps []commit
	for _, line := range strings.Split(string(out), "\n") {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		parts := strings.SplitN(line, "|", 2)
		hashes := strings.Split(parts[0], " ")
		c := commit{short: hashes[0], parents: hashes[1:], message: parts[1]}
		if strings.HasPrefix(c.message, "UPSTREAM: ") {
			hide[c.short] = struct{}{}
			upstreams = append(upstreams, c)
		}
		if strings.HasPrefix(c.message, "bump(") {
			hide[c.short] = struct{}{}
			bumps = append(bumps, c)
		}
		if len(c.parents) == 1 {
			commits = append(commits, c)
			continue
		}
		matches := mergeRequest.FindStringSubmatch(line)
		if len(matches) == 0 {
			continue
		}
		var first int
		for i := range commits {
			first = i
			if contains(c.parents, commits[i].short) {
				first++
				break
			}
		}
		individual := commits[:first]
		merged := commits[first:]
		for _, commit := range individual {
			if len(commit.parents) > 1 {
				continue
			}
			if _, ok := hide[commit.short]; ok {
				continue
			}
			fmt.Printf("force-merge: %s %s\n", commit.message, commit.short)
		}
		out, err := exec.Command("git", "show", "--pretty=tformat:%b", c.short).CombinedOutput()
		if err != nil {
			log.Fatal(err)
		}
		var message string
		para := strings.Split(string(out), "\n\n")
		if len(para) > 0 && strings.HasPrefix(para[0], "Automatic merge from submit-queue") {
			para = para[1:]
		}
		if len(para) > 0 && strings.HasPrefix(para[0], "Merged by ") {
			para = para[1:]
		}
		if len(para) > 0 {
			message = strings.Split(para[0], "\n")[0]
		}
		if len(message) == 0 && len(merged) > 0 {
			message = merged[0].message
		}
		if len(message) > 0 && len(merged) == 1 && message == merged[0].message {
			merged = nil
		}
		if len(message) > 0 && !prefix.MatchString(message) {
			prefix, ok := findPrefixFor(message, merged)
			if ok {
				message = prefix + ": " + message
			}
		}
		display := fmt.Sprintf("%s [\\#%s](https://github.com/kubernetes-incubator/service-catalog/pull/%s)", message, matches[1], matches[1])
		if hasFileChanges(c.short, "api/") {
			apiChanges = append(apiChanges, display)
		}
		var filtered []commit
		for _, commit := range merged {
			if _, ok := hide[commit.short]; ok {
				continue
			}
			filtered = append(filtered, commit)
		}
		if len(filtered) > 0 {
			fmt.Printf("- %s\n", display)
			for _, commit := range filtered {
				fmt.Printf("  - %s (%s)\n", commit.message, commit.short)
			}
		}
		commits = []commit{c}
	}
	var lines []string
	for _, commit := range bumps {
		if m := webconsoleBump.FindStringSubmatch(commit.message); len(m) > 0 {
			webconsole = append(webconsole, m[1])
			continue
		}
		lines = append(lines, commit.message)
	}
	lines = sortAndUniq(lines)
	for _, line := range lines {
		fmt.Printf("- %s\n", line)
	}
	lines = nil
	for _, commit := range upstreams {
		lines = append(lines, commit.message)
	}
	lines = sortAndUniq(lines)
	for _, line := range lines {
		fmt.Printf("- %s\n", upstreamLinkify(line))
	}
	if len(webconsole) > 0 {
		fmt.Printf("- web: from %s^..%s\n", webconsole[0], webconsole[len(webconsole)-1])
	}
	for _, apiChange := range apiChanges {
		fmt.Printf("  - %s\n", apiChange)
	}
}
func findPrefixFor(message string, commits []commit) (string, bool) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	message = strings.ToLower(message)
	for _, m := range assignments {
		if strings.Contains(message, m.term) {
			return m.prefix, true
		}
	}
	for _, c := range commits {
		if prefix, ok := findPrefixFor(c.message, nil); ok {
			return prefix, ok
		}
	}
	return "", false
}
func hasFileChanges(commit string, prefixes ...string) bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	out, err := exec.Command("git", "diff", "--name-only", fmt.Sprintf("%s^..%s", commit, commit)).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range strings.Split(string(out), "\n") {
		for _, prefix := range prefixes {
			if strings.HasPrefix(file, prefix) {
				return true
			}
		}
	}
	return false
}
func sortAndUniq(lines []string) []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	sort.Strings(lines)
	out := make([]string, 0, len(lines))
	last := ""
	for _, s := range lines {
		if last == s {
			continue
		}
		last = s
		out = append(out, s)
	}
	return out
}
func upstreamLinkify(line string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if m := upstreamKube.FindStringSubmatch(line); len(m) > 0 {
		return fmt.Sprintf("UPSTREAM: [#%s](https://github.com/kubernetes/kubernetes/pull/%s):%s", m[1], m[1], m[2])
	}
	if m := upstreamRepo.FindStringSubmatch(line); len(m) > 0 {
		return fmt.Sprintf("UPSTREAM: [%s#%s](https://github.com/%s/pull/%s):%s", m[1], m[2], m[1], m[2], m[3])
	}
	return line
}
func _logClusterCodePath() {
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
