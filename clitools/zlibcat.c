#include <zlib.h>
#include <stdio.h>
#include <stdlib.h>

// i know this could be less hacky but i dont give a shit :)
int main(int argc,char * argv[]){
FILE * fp = fopen(argv[1],"rb");
fseek(fp, 0, SEEK_END);
long file_size = ftell(fp); 
fseek(fp, 0, SEEK_SET);  


unsigned char * buffer = (unsigned char * )malloc(file_size);
fread(buffer,file_size,1,fp);
fclose(fp);
unsigned char * buf2 = (unsigned char *)malloc(file_size * 3 + 1);
z_stream stream;
stream.zalloc = Z_NULL;
stream.zfree = Z_NULL;
stream.opaque = Z_NULL;
stream.avail_in = file_size;
stream.next_in = buffer;
stream.next_out = buf2;
inflateInit(&stream);
inflate(&stream,Z_NO_FLUSH);
inflateEnd(&stream);

buf2[file_size * 3] = '\0'; // so we can print
printf("%s",buf2);

free(buffer);
free(buf2);
return 0;
}
