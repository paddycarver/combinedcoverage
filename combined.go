package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/tools/cover"
)

// percentCovered returns, as a percentage, the fraction of the statements in
// the profiles passed in.
// In effect, it reports the coverage of a given set of source files.
func percentCovered(p []*cover.Profile) float64 {
	var total, covered int64
	for _, prof := range p {
		for _, b := range prof.Blocks {
			total += int64(b.NumStmt)
			if b.Count > 0 {
				covered += int64(b.NumStmt)
			}
		}
	}
	if total == 0 {
		return 0
	}
	return float64(covered) / float64(total) * 100
}

func main() {
	flag.Parse()

	var profiles []*cover.Profile

	for _, file := range flag.Args() {
		profs, err := cover.ParseProfiles(file)
		if err != nil {
			log.Fatalf("failed to parse profiles: %v", err)
		}
		profiles = append(profiles, profs...)
	}
	fmt.Printf("coverage: %.1f%%\n", percentCovered(profiles))
}
