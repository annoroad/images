# annovarl
* vcf文件annovar注释

# 参数
```
Usage:
     table_annovar.pl [arguments] <query-file> <database-location>

     Optional arguments:
            -h, --help                      print help message
            -m, --man                       print complete documentation
            -v, --verbose                   use verbose output
                --protocol <string>         comma-delimited string specifying database protocol
                --operation <string>        comma-delimited string specifying type of operation
                --outfile <string>          output file name prefix
                --buildver <string>         genome build version (default: hg18)
                --remove                    remove all temporary files
                --(no)checkfile             check if database file exists (default: ON)
                --genericdbfile <files>     specify comma-delimited generic db files
                --gff3dbfile <files>        specify comma-delimited GFF3 files
                --bedfile <files>           specify comma-delimited BED files
                --vcfdbfile <files>         specify comma-delimited VCF files
                --otherinfo                 print out otherinfo (infomration after fifth column in queryfile)
                --onetranscript             print out only one transcript for exonic variants (default: all transcripts)
                --nastring <string>         string to display when a score is not available (default: null)
                --csvout                    generate comma-delimited CSV file (default: tab-delimited txt file)
                --argument <string>         comma-delimited strings as optional argument for each operation
                --tempdir <dir>             directory to store temporary files (default: --outfile)
                --vcfinput                  specify that input is in VCF format and output will be in VCF format
                --dot2underline             change dot in field name to underline (eg, Func.refGene to Func_refGene)
                --thread <int>              specify the number of threads to be used in annotation
                --maxgenethread <int>       specify the maximum number of threads allowed in gene annotation (default: 6)
            

     Function: automatically run a pipeline on a list of variants and summarize
     their functional effects in a comma-delimited file, or to an annotated VCF file
     if the original input is a VCF file
 
     Example: table_annovar.pl example/ex1.avinput humandb/ -buildver hg19 -out myanno -remove -protocol refGene,cytoBand,dbnsfp30a -operation g,r,f -nastring . -csvout
              table_annovar.pl example/ex2.vcf humandb/ -buildver hg19 -out myanno -remove -protocol refGene,cytoBand,esp6500siv2_all,1000g2015aug_all,1000g2015aug_afr,1000g2015aug_eas,1000g2015aug_eur,snp138,dbnsfp30a -operation g,r,f,f,f,f,f,f,f -nastring . -vcfinput
                  
     Version: $Date: 2016-02-01 00:11:18 -0800 (Mon,  1 Feb 2016) $
```
