# sniffles
* 三代数据SV检测

# 参数
```
Usage: sniffles [options] -m <sorted.bam> -v <output.vcf> 
Version: 1.0.11
Contact: fritz.sedlazeck@gmail.com

Input/Output:
    -m <string>,  --mapped_reads <string>
        (required)  Sorted bam File
    -v <string>,  --vcf <string>
        VCF output file name []
    -b <string>,  --bedpe <string>
         bedpe output file name []
    --Ivcf <string>
        Input VCF file name. Enable force calling []
    --tmp_file <string>
        path to temporary file otherwise Sniffles will use the current directory. []

General:
    -s <int>,  --min_support <int>
        Minimum number of reads that support a SV. [10]
    --max_num_splits <int>
        Maximum number of splits per read to be still taken into account. [7]
    -d <int>,  --max_distance <int>
        Maximum distance to group SV together. [1000]
    -t <int>,  --threads <int>
        Number of threads to use. [3]
    -l <int>,  --min_length <int>
        Minimum length of SV to be reported. [30]
    -q <int>,  --minmapping_qual <int>
        Minimum Mapping Quality. [20]
    -n <int>,  --num_reads_report <int>
        Report up to N reads that support the SV in the vcf file. -1: report all. [0]
    -r <int>,  --min_seq_size <int>
        Discard read if non of its segment is larger then this. [2000]
    -z <int>,  --min_zmw <int>
        Discard SV that are not supported by at least x zmws. This applies only for PacBio recognizable reads. [0]
    --cs_string
        Enables the scan of CS string instead of Cigar and MD.  [false]

Clustering/phasing and genotyping:
    --genotype
        Enables Sniffles to compute the genotypes. [false]
    --cluster
        Enables Sniffles to phase SVs that occur on the same reads [false]
    --cluster_support <int>
        Minimum number of reads supporting clustering of SV. [1]
    -f <float>,  --allelefreq <float>
        Threshold on allele frequency (0-1).  [0]
    --min_homo_af <float>
        Threshold on allele frequency (0-1).  [0.8]
    --min_het_af <float>
        Threshold on allele frequency (0-1).  [0.3]

Advanced:
    --report_BND
        Dont report BND instead use Tra in vcf output.  [true]
    --report_seq
        Report sequences for indels in vcf output. (Beta version!)  [false]
    --ignore_sd
        Ignores the sd based filtering.  [false]
    --report_read_strands
        Enables the report of the strand categories per read. (Beta)  [false]
    --ccs_reads
        Preset CCS Pacbio setting. (Beta)  [false]

Parameter estimation:
    --skip_parameter_estimation
        Enables the scan if only very few reads are present.  [false]
    --del_ratio <float>
        Estimated ration of deletions per read (0-1).  [0.0458369]
    --ins_ratio <float>
        Estimated ratio of insertions per read (0-1).  [0.049379]
    --max_diff_per_window <int>
        Maximum differences per 100bp. [50]
    --max_dist_aln_events <int>
        Maximum distance between alignment (indel) events. [4]
```
