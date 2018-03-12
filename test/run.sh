#!/bin/bash
vegeta attack -rate=100 -duration=10s -targets=targets.txt > results.bin
vegeta report -inputs=results.bin -reporter=json > metrics.json
cat results.bin | vegeta report -reporter=plot > plot.html
cat results.bin | vegeta report -reporter="hist[0,100ms,200ms,300ms]"
xdg-pen plot.html
