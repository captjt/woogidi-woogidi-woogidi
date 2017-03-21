FROM alpine 

LABEL name="woogidi-woogidi-woogidi" \
      maintainer="jtaylor007.jt@gmail.com"  

COPY woogidi-woogidi-woogidi woogidi-woogidi-woogidi

ENTRYPOINT ["./woogidi-woogidi-woogidi"]