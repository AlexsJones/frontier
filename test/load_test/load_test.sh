#!/bin/bash
vegeta attack -rate=10000 -duration=10s -targets=targets.txt > results.bin
vegeta report -inputs=results.bin -reporter=json > metrics.json
cat results.bin | vegeta report -reporter=plot > plot.html
cat results.bin | vegeta report
xdg-open plot.html
