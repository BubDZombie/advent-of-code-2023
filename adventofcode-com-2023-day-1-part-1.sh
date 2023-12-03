#!/bin/bash
s=0;while read j;do d="${j//[a-z]/}";f=${d:0:1};l=${d:0-1};let 's+=10*f+l';done</dev/stdin;echo $s
