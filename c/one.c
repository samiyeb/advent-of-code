#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* readFile(const char* filename) {
    FILE* fp = fopen(filename, "rb");
    if (!fp) {
        perror("Failed to open file");
        return NULL;
    }

    // Initial buffer size
    size_t size = 128;
    char* buffer = malloc(size);
    if (!buffer) {
        perror("Failed to allocate memory");
        fclose(fp);
        return NULL;
    }

    size_t len = 0;
    int ch;
    while ((ch = fgetc(fp)) != EOF) {
        // Reallocate buffer if necessary
        if (len + 1 >= size) {
            size *= 2;
            char* new_buffer = realloc(buffer, size);
            if (!new_buffer) {
                perror("Failed to reallocate memory");
                free(buffer);
                fclose(fp);
                return NULL;
            }
            buffer = new_buffer;
        }
        buffer[len++] = ch;
    }
    buffer[len] = '\0';  // Null-terminate the string

    fclose(fp);
    return buffer;
}

int main() {
    char* content = readFile("text.txt");

    char *token = strtok(content, ",");
    while(token) {
        token = strtok(NULL, ",");
    }

    // if (content) {
    //     printf("File Content:\n%s\n", content);
    //     free(content);
    // }
    return 0;
}
