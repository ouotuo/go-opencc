#include <stdlib.h>

#include "opencc.h"

void* Opencc_New(const char *configFile);
void Opencc_Delete(void *id);
const char *Opencc_Convert(void *id, const char *input);
void Opencc_Free_String(char *p);
const char *Convert(const char *input, const char *config);
