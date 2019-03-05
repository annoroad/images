# transpose
* 行列转置

# 参数
```
Description:     
 	This software is released under the GPL license
 	Reshapes delimited text data - amongst other things, it can transpose a matrix of plain text data.
 	The input file (optional, otherwise stdin is used) should be the last argument provided

 	Options:	
        -v or --version 	displays version information
        -h or --help    	displays this helpful information
        -b or --block <integer>		specify the blocksize (in bytes) for reading the input (default is 10kb)
        --fsep <character>		sets field separator to specified character. Defaults is \t
        --lsep <character>		sets line separator to specified character. Default is \n
 	--fieldmax or -f <integer>ine separator sets the number of characters to allow for in each field (default is 64)
 	--input or -i <integer>x<integer>[+<integer>+<integer>]		specificies size and cropping of the input. Dimension order is rows, columns
 	--output or -o <integer>x<integer>		specificies dimensions for the output. Default is determined from the input dimensions and whether we are transposing.
 	--limit or -l <integer>x<integer>		when not input dimensions, allow for this size (rowsxcolumns) matrix. Default is 1000x1000
 	--transpose or -t		indicates that the output matrix is to be filled in columnwise. When output dimensions are not specified (or equal the transpose of the input dimensions) the output is an exact transpose of the input.
```
