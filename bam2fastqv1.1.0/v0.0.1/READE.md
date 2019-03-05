# bam2fastq
* bam格式转fastq格式

# 参数
```
bam2fastq v1.1.0 - extract sequences from a BAM file

Usage: bam2fastq [options] <bam file>

Options:
  -o FILENAME, --output FILENAME
       Specifies the name of the FASTQ file(s) that will be generated.  May
       contain the special characters % (replaced with the lane number) and
       # (replaced with _1 or _2 to distinguish PE reads, removed for SE reads).
       [Default: s_%#_sequence.txt]
  -f, --force, --overwrite
       Create output files specified with --output, overwriting existing
       files if necessary [Default: exit program rather than overwrite files]
  --aligned
  --no-aligned
       Reads in the BAM that are aligned will (will not) be extracted.
       [Default: extract aligned reads]
  --unaligned
  --no-unaligned
       Reads in the BAM that are not aligned will (will not) be extracted.
       [Default: extract unaligned reads]
  --filtered
  --no-filtered
       Reads that are marked as failing QC checks will (will not) be extracted.
       [Default: extract filtered reads]
  -q, --quiet
       Suppress informational messages [Default: print messages]
  -s, --strict
       Keep bam2fastq's processing to a minimum, assuming that the BAM strictly       meets specifications. [Default: allow some errors in the BAM]
```
