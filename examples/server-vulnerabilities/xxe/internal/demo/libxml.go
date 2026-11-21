package demo

/*
#cgo pkg-config: libxml-2.0
#include <stdlib.h>
#include <libxml/parser.h>
#include <libxml/tree.h>
// XML_PARSE_NOENT enables entity substitution. NONET does not block local files.
static xmlChar* vulnerable_parse(const char* data, int size) {
 xmlDocPtr doc = xmlReadMemory(data, size, "demo.xml", NULL,
     XML_PARSE_NOENT | XML_PARSE_NONET | XML_PARSE_NOERROR | XML_PARSE_NOWARNING);
 if (doc == NULL) return NULL;
 xmlNodePtr root = xmlDocGetRootElement(doc);
 xmlChar* text = root == NULL ? NULL : xmlNodeGetContent(root);
 xmlFreeDoc(doc);
 return text;
}
static void release_text(xmlChar* text) { xmlFree(text); }
*/
import "C"
import (
	"errors"
	"unsafe"
)

func parseVulnerable(data []byte) (string, error) {
	input := C.CBytes(data)
	defer C.free(input)
	output := C.vulnerable_parse((*C.char)(input), C.int(len(data)))
	if output == nil {
		return "", errors.New("invalid XML")
	}
	defer C.release_text(output)
	return C.GoString((*C.char)(unsafe.Pointer(output))), nil
}
