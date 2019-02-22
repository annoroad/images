# bcl2fastq
* 经BCL数据转化成fastq

# 参数
```
  -i [ --input-dir ] arg (=<runfolder-dir>/Data/Intensities/BaseCalls/)
  -R [ --runfolder-dir ] arg (=./)                path to runfolder directory
  --intensities-dir arg (=<input-dir>/../)        path to intensities directory
  -o [ --output-dir ] arg (=<input-dir>)          path to demultiplexed output
  --interop-dir arg (=<runfolder-dir>/InterOp/)   path to demultiplexing statistics directory
  --stats-dir arg (=<output-dir>/Stats/)          path to human-readable demultiplexing statistics directory
  --reports-dir arg (=<output-dir>/Reports/)      path to reporting directory
  --sample-sheet arg (=<runfolder-dir>/SampleSheet.csv)
                                                  path to the sample sheet
  -r [ --loading-threads ] arg (=4)               number of threads used for loading BCL data
  -p [ --processing-threads ] arg                 number of threads used for processing demultiplexed data
  -w [ --writing-threads ] arg (=4)               number of threads used for writing FASTQ data.
                                                  This should not be set higher than the number of samples. If set =0 then these threads will be 
                                                  placed in the same pool as the loading threads, and the number of shared threads will be 
                                                  determined by --loading-threads.
  --tiles arg                                     comma-separated list of regular expressions to select only a subset of the tiles available in the
                                                  flow-cell. Multiple entries allowed, each applies to the corresponding base-calls.
                                                  For example:
                                                   * to select all the tiles ending with '5' in all lanes:
                                                       --tiles [0-9][0-9][0-9]5
                                                   * to select tile 2 in lane 1 and all the tiles in the other lanes:
                                                       --tiles s_1_0002,s_[2-8]
  --minimum-trimmed-read-length arg (=35)         minimum read length after adapter trimming
  --use-bases-mask arg                            specifies how to use each cycle.
  --mask-short-adapter-reads arg (=22)            smallest number of remaining bases (after masking bases below the minimum trimmed read length) 
                                                  below which whole read is masked
  --adapter-stringency arg (=0.9)                 adapter stringency
  --ignore-missing-bcls                           assume 'N'/'#' for missing calls
  --ignore-missing-filter                         assume 'true' for missing filters
  --ignore-missing-positions                      assume [0,i] for missing positions, where i is incremented starting from 0
  --ignore-missing-controls                       (deprecated) assume 0 for missing controls
  --write-fastq-reverse-complement                generate FASTQs containing reverse complements of actual data
  --with-failed-reads                             include non-PF clusters
  --create-fastq-for-index-reads                  create FASTQ files also for index reads
  --find-adapters-with-sliding-window             find adapters with simple sliding window algorithm
  --no-bgzf-compression                           turn off BGZF compression for FASTQ files
  --fastq-compression-level arg (=4)              zlib compression level (1-9) used for FASTQ files
  --barcode-mismatches arg (=1)                   number of allowed mismatches per index
                                                  Multiple, comma delimited, entries allowed. Each entry is applied to the corresponding index; 
                                                  last entry applies to all remaining indices.
                                                  Accepted values: 0, 1, 2.
  --no-lane-splitting                             do not split fastq files by lane
```
