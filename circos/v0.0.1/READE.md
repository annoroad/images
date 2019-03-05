# circos
* 全局总览图

# 参数
```
Usage:
      # guess location of configuration file
      circos

      # use specific configuration file
      circos -conf circos.conf [-silent]

      # diagnose required modules
      circos -modules

      # man page
      circos -man

      # detailed debugging for code components
      # see http://www.circos.ca/documentation/tutorials/configuration/debugging
      circos -debug_group GROUP1,[GROUP2,...]

      # full debugging
      circos -debug_group _all

      # configuration dump of a block (or block tree) of
      # any parameters that match REGEXP (optional)
      circos -cdump [BLOCK1/[BLOCK2/...]]{:REGEXP}

      # overriding configuration parameters
      circos -param image/radius=2000p -param ideogram/show=no

      # for fun - randomize all colors in the image except for
      # COLOR1 and COLOR2
      circos -randomcolor COLOR1,[COLOR2,...]
```
