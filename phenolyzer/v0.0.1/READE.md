# phenolyzer
* 表型相关性分析

# 参数
```
Usage:
     disease_annotation.pl [arguments] <disease_names or disease_filename>

     Optional arguments:
            -h, --help                      print help message
            -m, --man                       print complete documentation
            -v, --verbose                   use verbose output
            -out <string>                   output file name prefix (default:out)
            -d, --directory                 compiled database directory (default is ./lib/compiled_database)
            -f, --file                      the input will be treated as file names(both diseases and genes)
            -p, --prediction                Use the Protein interaction and Biosystem database to predict unreported gene 
                                            disease relations (like HPRD human protein interaction, Biosystem database and so on)
            -ph, --phenotype                the input term is also treated as a phenotype, the HPO annotation and OMIM description would be used      
            --bedfile                       the bed file as a genomic region used for selection and annotation of the genes
            --buildver                      the build version (hg18 or hg19) to annotate the bedfile
            --wordcloud                     generates a wordcloud of the interpretated diseases if used (not working if you input 'all diseases')
            --logistic                      uses the weight based on the logistic modeling with four different complex diseases
            --gene                          the genes used to select the results (file name if -f command is used)    
            --exact                         choose if you want only exact match but not just a word match
            --addon                         the name of a user-defined add-on gene-disease mapping file (has to be in the ./lib/compiled_database)
            --addon_gg                      the name of user-defined add-on gene-gene mapping file (has to be in the ./lib/compiled_database)
            --addon_weight                  the weight of add-on gene-disease mapping
            --addon_gg_weight               the weight of add-on gene-gene mapping
            --hprd_weight                   the weight for genes found in HPRD
            --biosystem_weight              the weight for genes found in Ncbi Biosystem 
            --gene_family_weight            the weight for genes found in HGNC Gene Family
            --htri_weight                   the weight for genes found in HTRI Transcription Interaction Database
            --gwas_weight                   the weight for gene disease pairs in Gwas Catalog
            --gene_reviews_weight           the weight for gene disease pairs in Gene Reviews  
            --clinvar_weight                the weight for gene disease pairs in Clinvar
            --omim_weight                   the weight for gene disease pairs in OMIM
            --orphanet_weight               the weight for gene disease pairs in Orphanet

    Function: automatically expand the input disease term to a list of
    professional disease names, get a prioritized genelist based on these
    disease names or phenotypes, score the genes.

    Notice: If you input 'all diseases' for disease name, then every item in
    the gene_disease database will be used and no disease expansion will be
    conducted. Addon Gene Gene file should be in the format "GENE A GENE B
    EVIDENCE SCORE PMID" Addon Gene Disease file should be in the format
    "GENE DISEASE DISEASE_ID SCORE SOURCE"

    Example: perl disease_annotation.pl sleep -p perl disease_annotation.pl
    disease -f -p -ph

    Version: 1.0.5 $Last Changed Date: 02-21-2015 by Hui Yang
```
